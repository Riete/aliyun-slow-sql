package rdsquery

import (
	"fmt"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
)

const TimeFormat = "2006-01-02T15:04Z"

func NewClient(regionId, accessKeyId, accessKeySecret string) *rds.Client {
	client, err := rds.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		panic(fmt.Sprintf("[RDS]: create client failed, %s", err))
	}
	return client
}

func QuerySlowSQL(client *rds.Client, instanceId string) ([]rds.SQLSlowRecord, error) {
	utcNow := time.Now().UTC()
	request := rds.CreateDescribeSlowLogRecordsRequest()
	request.DBInstanceId = instanceId
	request.EndTime = utcNow.Format(TimeFormat)
	request.StartTime = utcNow.Add(-5 * time.Minute).Format(TimeFormat)
	request.PageNumber = requests.NewInteger(1)
	request.PageSize = requests.NewInteger(100)
	response, err := client.DescribeSlowLogRecords(request)
	if err != nil {
		return nil, err
	}
	records := response.Items.SQLSlowRecord
	if len(records) <= 20 {
		return records, nil
	}
	return records[0:20], nil
}

func GetNameById(client *rds.Client, instanceId string) (string, error) {
	request := rds.CreateDescribeDBInstanceAttributeRequest()
	request.DBInstanceId = instanceId
	response, err := client.DescribeDBInstanceAttribute(request)
	if err != nil {
		return "", err
	}
	return response.Items.DBInstanceAttribute[0].DBInstanceDescription, nil
}
