package polardb

import (
	"errors"
	"fmt"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/polardb"
	"github.com/riete/aliyun-slow-sql/send"
)

const TimeFormat = "2006-01-02T15:04Z"

type PolardbClient struct {
	client *polardb.Client
}

func (p *PolardbClient) NewClient(regionId, accessKeyId, accessKeySecret string) error {
	var err error
	p.client, err = polardb.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		return errors.New(fmt.Sprintf("[PolarDB]: create client failed, %s", err))
	}
	return nil
}

func (p *PolardbClient) SlowSql(instanceId string) (send.Records, error) {
	var records send.Records
	utcNow := time.Now().UTC()
	request := polardb.CreateDescribeSlowLogRecordsRequest()
	request.DBClusterId = instanceId
	request.EndTime = utcNow.Format(TimeFormat)
	request.StartTime = utcNow.Add(-5 * time.Minute).Format(TimeFormat)
	request.PageNumber = requests.NewInteger(1)
	request.PageSize = requests.NewInteger(100)
	response, err := p.client.DescribeSlowLogRecords(request)
	if err != nil {
		return records, err
	}
	for _, i := range response.Items.SQLSlowRecord {
		records = append(
			records,
			send.Record{
				HostAddress:        i.HostAddress,
				DBName:             i.DBName,
				SQLText:            i.SQLText,
				QueryTimes:         i.QueryTimes,
				LockTimes:          i.LockTimes,
				ParseRowCounts:     i.ParseRowCounts,
				ReturnRowCounts:    i.ReturnRowCounts,
				ExecutionStartTime: i.ExecutionStartTime,
			},
		)
	}
	return records, nil
}

func (p *PolardbClient) InstanceName(instanceId string) string {
	request := polardb.CreateDescribeDBClusterAttributeRequest()
	request.DBClusterId = instanceId
	response, err := p.client.DescribeDBClusterAttribute(request)
	if err != nil {
		return instanceId
	}
	instanceName := response.DBClusterDescription
	if instanceName == "" {
		instanceName = instanceId
	}
	return instanceName
}
