package send

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/riete/go-tools/notify"

	"github.com/riete/aliyun-slow-sql/rdsquery"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
)

func newMessage(title string, record rds.SQLSlowRecord, exclude map[string]bool, ch chan<- string) {
	executeTimeStr := record.ExecutionStartTime
	executeTime, _ := time.Parse("2006-01-02T15:04:05Z", executeTimeStr)
	if !exclude[record.DBName] {
		message := fmt.Sprintf(
			"> 执行时间：%s\n\n> 客户端IP：%s\n\n> 数据库名：%s\n\n> 执行时长：%ds\n\n"+
				"> 锁定时长：%ds\n\n> 解析行数：%d\n\n> 返回行数：%d\n\n> SQL语句：%s\n\n",
			executeTime.UTC().Add(8*time.Hour).Format("2006-01-02 15:04:05"),
			record.HostAddress,
			record.DBName,
			record.QueryTimes,
			record.LockTimes,
			record.ParseRowCounts,
			record.ReturnRowCounts,
			record.SQLText,
		)
		ch <- fmt.Sprintf("%s=====%s", title, message)
	}
}

func NewMessage(instanceId string, client *rds.Client, exclude map[string]bool, ch chan<- string) error {
	name, err := rdsquery.GetNameById(client, instanceId)
	if err != nil {
		return err
	}
	title := fmt.Sprintf("RDS数据库[%s]新增慢SQL信息", name)

	records, err := rdsquery.QuerySlowSQL(client, instanceId)
	if err != nil {
		return err
	}
	for _, record := range records {
		newMessage(title, record, exclude, ch)
	}
	return nil
}

func DoSend(url, secret string, ch <-chan string) {
	var err error
	var result string

	for m := range ch {
		title := strings.Split(m, "=====")[0]
		message := strings.Split(m, "=====")[1]

		if secret != "" {
			result, err = notify.SendDingTalkNew(title, message, url, secret, true, nil)
		} else {
			result, err = notify.SendDingTalk(title, message, url, true, nil)
		}

		if err != nil {
			log.Println(err)
		}
		log.Println(result)
		time.Sleep(3 * time.Second)
	}
}
