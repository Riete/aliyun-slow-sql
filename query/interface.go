package query

import (
	"github.com/riete/aliyun-slow-sql/query/polardb"
	"github.com/riete/aliyun-slow-sql/query/rds"
	"github.com/riete/aliyun-slow-sql/send"
)

type DBType string

const (
	Rds     DBType = "RDS"
	PolarDB DBType = "PolarDB"
)

type Client interface {
	NewClient(regionId, accessKeyId, accessKeySecret string) error
	SlowSql(instanceId string) (send.Records, error)
	InstanceName(instanceId string) string
}

func NewClient(t DBType) Client {
	switch t {
	case Rds:
		return &rds.RdsClient{}
	case PolarDB:
		return &polardb.PolardbClient{}
	}
	return nil
}
