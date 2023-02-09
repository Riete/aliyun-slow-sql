package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/riete/aliyun-slow-sql/query"
	"github.com/riete/aliyun-slow-sql/send"
	"github.com/riete/dingtalk"
)

type Instance struct {
	InstanceId   string       `json:"instance_id"`
	ExcludedDB   []string     `json:"excluded_db"`
	InstanceType query.DBType `json:"instance_type"`
}

func (i Instance) Exclude() map[string]bool {
	ex := make(map[string]bool)
	for _, exclude := range i.ExcludedDB {
		ex[exclude] = true
	}
	return ex
}

type Task struct {
	AccessKeyId     string     `json:"access_key_id"`
	AccessKeySecret string     `json:"access_key_secret"`
	RegionId        string     `json:"region_id"`
	Instances       []Instance `json:"instances"`
	DingTalk        struct {
		Webhook string `json:"webhook"`
		Secret  string `json:"secret"`
	} `json:"ding_talk"`
}

type Config struct {
	Tasks []Task `json:"tasks"`
}

func (c Config) Validate() error {
	for _, task := range c.Tasks {
		for _, instance := range task.Instances {
			if !(instance.InstanceType == query.Rds || instance.InstanceType == query.PolarDB) {
				return errors.New(fmt.Sprintf("instance_type only support [%s] and [%s] for now", query.Rds, query.PolarDB))
			}
		}
	}
	return nil
}

func parseConfig() (Config, error) {
	config := flag.String("config", "config.json", "config file path")
	flag.Parse()

	var c Config
	d, err := os.ReadFile(*config)
	if err != nil {
		return c, errors.New("read config file error: " + err.Error())
	}
	if err := json.Unmarshal(d, &c); err != nil {
		return c, errors.New("parse config file error: " + err.Error())
	}
	return c, nil
}

func main() {
	config, err := parseConfig()
	if err != nil {
		log.Fatalln(err)
	}
	if err := config.Validate(); err != nil {
		log.Fatalln(err)
	}
	for _, task := range config.Tasks {
		ch := make(chan string)
		go send.DoSend(dingtalk.NewDingTalk(task.DingTalk.Webhook, task.DingTalk.Secret), ch)
		for _, instance := range task.Instances {
			go func(t Task, i Instance) {
				client := query.NewClient(i.InstanceType)
				if err := client.NewClient(t.RegionId, t.AccessKeyId, t.AccessKeySecret); err != nil {
					log.Println(fmt.Sprintf("[%s] init client error", t.AccessKeyId), err)
					return
				}
				instanceName := client.InstanceName(i.InstanceId)
				exclude := i.Exclude()

				for {
					if records, err := client.SlowSql(i.InstanceId); err != nil {
						log.Println(fmt.Sprintf("query [%s] slow sql records error", instanceName), err)
					} else {
						records.NewMessage(instanceName, string(i.InstanceType), exclude, ch)
					}
					time.Sleep(5 * time.Minute)
				}
			}(task, instance)
		}
	}
	infinite := make(chan struct{})
	<-infinite
}
