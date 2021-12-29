package main

type aliConfig struct {
	AccessKey    string `json:"access_key"`
	AccessSecret string `json:"access_secret"`
	Domain       string `json:"domain"` //域名
	RR           string `json:"rr"`     //解析的主机记录
}
