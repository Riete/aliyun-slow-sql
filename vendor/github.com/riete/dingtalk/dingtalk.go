package dingtalk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/riete/requests"
)

type DingTalk interface {
	SendMarkdown(title, message string, isAtAll bool, atMobiles ...string) string
	SendText(title, message string, isAtAll bool, atMobiles ...string) string
}

func NewDingTalk(url, secret string) DingTalk {
	return &dingtalk{url: url, secret: secret}
}

type dingtalk struct {
	url       string
	secret    string
	signedUrl string
}

func (dt *dingtalk) sign() {
	timestamp := fmt.Sprintf("%d000", time.Now().Unix())
	sign := fmt.Sprintf("%s\n%s", timestamp, dt.secret)
	h := hmac.New(sha256.New, []byte(dt.secret))
	h.Write([]byte(sign))
	signB64 := base64.StdEncoding.EncodeToString([]byte(h.Sum(nil)))
	v := url.Values{}
	v.Add("sign", signB64)
	signUrlEncode := v.Encode()
	dt.signedUrl = fmt.Sprintf("%s&timestamp=%s&%s", dt.url, timestamp, signUrlEncode)
}

func (dt dingtalk) formatMarkdown(title, message string, isAtAll bool, atMobiles ...string) map[string]interface{} {
	var mobiles []string
	for _, m := range atMobiles {
		mobiles = append(mobiles, fmt.Sprintf("@%s", m))
	}
	return map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"title": title,
			"text":  fmt.Sprintf("### %s\n\n%s\n\n%s", title, message, strings.Join(mobiles, " ")),
		},
		"at": map[string]interface{}{
			"atMobiles": atMobiles,
			"isAtAll":   isAtAll,
		},
	}
}

func (dt dingtalk) formatText(title, message string, isAtAll bool, atMobiles ...string) map[string]interface{} {
	return map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": fmt.Sprintf("%s\n\n%s", title, message),
		},
		"at": map[string]interface{}{
			"atMobiles": atMobiles,
			"isAtAll":   isAtAll,
		},
	}
}

func (dt *dingtalk) send(body map[string]interface{}) string {
	dt.sign()
	if r, err := requests.Post(dt.signedUrl, body); err != nil {
		return err.Error()
	} else {
		return r.ContentToString()
	}
}

func (dt *dingtalk) SendMarkdown(title, message string, isAtAll bool, atMobiles ...string) string {
	return dt.send(dt.formatMarkdown(title, message, isAtAll, atMobiles...))
}

func (dt *dingtalk) SendText(title, message string, isAtAll bool, atMobiles ...string) string {
	return dt.send(dt.formatText(title, message, isAtAll, atMobiles...))
}
