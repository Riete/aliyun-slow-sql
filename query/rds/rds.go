package rds

import (
	"errors"
	"fmt"
	"time"

	"github.com/riete/aliyun-slow-sql/send"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
)

const TimeFormat = "2006-01-02T15:04Z"

type RdsClient struct {
	client *rds.Client
}

func (r *RdsClient) NewClient(regionId, accessKeyId, accessKeySecret string) error {
	var err error
	r.client, err = rds.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		return errors.New(fmt.Sprintf("[RDS]: create client failed, %s", err))
	}
	return nil
}

func (r RdsClient) SlowSql(instanceId string) (send.Records, error) {
	var records send.Records
	utcNow := time.Now().UTC()
	request := rds.CreateDescribeSlowLogRecordsRequest()
	request.DBInstanceId = instanceId
	request.EndTime = utcNow.Format(TimeFormat)
	request.StartTime = utcNow.Add(-5 * time.Minute).Format(TimeFormat)
	request.PageNumber = requests.NewInteger(1)
	request.PageSize = requests.NewInteger(100)
	response, err := r.client.DescribeSlowLogRecords(request)
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

func (r RdsClient) InstanceName(instanceId string) string {
	request := rds.CreateDescribeDBInstanceAttributeRequest()
	request.DBInstanceId = instanceId
	response, err := r.client.DescribeDBInstanceAttribute(request)
	if err != nil {
		return instanceId
	}
	instanceName := response.Items.DBInstanceAttribute[0].DBInstanceDescription
	if instanceName == "" {
		instanceName = instanceId
	}
	return instanceName
}
