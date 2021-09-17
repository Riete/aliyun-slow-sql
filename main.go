package main

import (
	"flag"
	"log"
	"strings"
	"time"

	"github.com/riete/aliyun-slow-sql/send"

	"github.com/riete/aliyun-slow-sql/rdsquery"
)

func main() {
	accessKeyId := flag.String("access.key.id", "", "aliyun access key id")
	accessKeySecret := flag.String("access.key.secret", "", "aliyun access key secret")
	regionId := flag.String("region.id", "", "aliyun region id")
	secret := flag.String("secret", "", "dingtalk callback url secret")
	url := flag.String("robot.url", "", "dingtalk callback url")
	instanceIds := flag.String("instance.ids", "", "rdsquery id, separated by ','")
	excludeDb := flag.String("exclude.db", "", "database not send alert, separated by ','")
	flag.Parse()

	if *accessKeyId == "" {
		log.Fatalln("access.key.id is required")
	}

	if *accessKeySecret == "" {
		log.Fatalln("access.key.secret is required")
	}

	if *regionId == "" {
		log.Fatalln("region.id is required")
	}

	if *url == "" {
		log.Fatalln("robot.url is required")
	}

	if *instanceIds == "" {
		log.Fatalln("instance.ids is required")
	}

	ids := strings.Split(*instanceIds, ",")
	client := rdsquery.NewClient(*regionId, *accessKeyId, *accessKeySecret)
	exclude := make(map[string]bool)
	for _, v := range strings.Split(*excludeDb, ",") {
		exclude[v] = true
	}

	message := make(chan string, 20)
	go send.DoSend(*url, *secret, message)

	for {
		for _, id := range ids {
			go func(instanceId string) {
				err := send.NewMessage(instanceId, client, exclude, message)
				if err != nil {
					log.Println(err)
				}
			}(id)
		}
		time.Sleep(5 * time.Minute)
	}
}
