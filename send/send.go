package send

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/riete/dingtalk"
)

type Record struct {
	HostAddress        string
	DBName             string
	SQLText            string
	QueryTimes         int64
	LockTimes          int64
	ParseRowCounts     int64
	ReturnRowCounts    int64
	ExecutionStartTime string
}

type Records []Record

func (r Records) newMessage(title string, record Record, exclude map[string]bool, ch chan<- string) {
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

func (r Records) NewMessage(instanceName string, dbType string, exclude map[string]bool, ch chan<- string) {
	title := fmt.Sprintf("%s数据库[%s]新增慢SQL信息", dbType, instanceName)
	for _, record := range r {
		r.newMessage(title, record, exclude, ch)
	}
}

func DoSend(d dingtalk.DingTalk, ch <-chan string) {
	for m := range ch {
		title := strings.Split(m, "=====")[0]
		message := strings.Split(m, "=====")[1]

		result := d.SendMarkdown(title, message, false)
		log.Println(result)
		time.Sleep(3 * time.Second)
	}
}
