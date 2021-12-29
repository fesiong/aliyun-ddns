package main

import (
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"log"
	"net"
	"net/http"

)

var LastIp string

func main() {
	log.Println("自动更新阿里云dns程序正在执行：")
	Crond()

	select {
	//保持业务不停止
	}
}

func Crond() {
	//开机的时候执行一次
	CheckUpdateDns()
	//每20分钟读取一次
	crontab := cron.New(cron.WithSeconds())
	crontab.AddFunc("1 */20 * * * *", CheckUpdateDns)
}

func GetInternetIp() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://4.ipw.cn/", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	ip := net.ParseIP(string(body))
	if ip == nil {
		return ""
	}

	return ip.String()
}

func CheckUpdateDns() {
	newIp := GetInternetIp()
	if newIp == "" {
		//伪获取到ip
		log.Println("未获取到本地网络的公网IP")
		DebugLog("ddns", "未获取到本地网络的公网IP")
		return
	}

	if newIp != LastIp {
		//需要更新
		log.Println("公网IP更新为：", newIp)
		DebugLog("ddns", fmt.Sprintf("公网IP更新为：%s", newIp))
		LastIp = newIp

		err := UpdateAliDns(newIp)
		if err != nil {
			log.Println("更新本地公网IP失败：", err.Error())
			DebugLog("ddns", fmt.Sprintf("更新本地公网IP失败：%s", err.Error()))
		}
	}
}

func UpdateAliDns(ipValue string) error {
	client, err := alidns.NewClientWithAccessKey("cn-hangzhou", JsonData.AccessKey, JsonData.AccessSecret)
	if err != nil {
		return err
	}
	req := alidns.CreateDescribeSubDomainRecordsRequest()
	req.DomainName = JsonData.Domain
	req.SubDomain = fmt.Sprintf("%s.%s", JsonData.RR, JsonData.Domain)
	response, err := client.DescribeSubDomainRecords(req)
	if err != nil {
		return err
	}

	if response.TotalCount == 0 {
		return errors.New("相关记录值为空")
	}
	//拿到解析记录
	record := response.DomainRecords.Record[0]

	//如果相同，则不需要重新提交解析
	if record.Value == ipValue {
		return nil
	}

	log.Println(record.RR, record)
	//开始尝试更新
	reqU := alidns.CreateUpdateDomainRecordRequest()
	reqU.RR = record.RR
	reqU.RecordId = record.RecordId
	reqU.Value = ipValue
	reqU.Type = record.Type
	respU, err := client.UpdateDomainRecord(reqU)

	if err != nil {
		return err
	}

	log.Println("更改解析成功")

	DebugLog("ddns", fmt.Sprintf("response is %#v\n", respU))

	return nil
}