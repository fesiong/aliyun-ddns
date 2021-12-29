# aliyun-ddns
阿里云域名动态解析处理golang版


如果会GO语言，则建议使用源码自己编译，按照自己的喜好修改。   
不会的同学，可以直接使用，请从后边的Release 中下载
> Windows:   
> aliyun_ddns.exe
>
> Linux:   
> aliyun_ddns

将以上提到的文件添加进系统的自动执行策略里，就可以实现动态更新自己的域名到阿里云的解析上了。

系统额外增加了定时检查ip机制，每半小时校对一次ip如果ip变化了，则执行更新

## 配置文件说明
需要同步的配置内容在config.json文件内，此文件需要同执行文件放在同一目录下


    access_key:你的阿里云访问key，在你的阿里云控制台里面的accesskeys里面去找
    access_secret:同上
    domain:你申请的域名，例如xxxx.com
    rr:对应解析配置中的主机记录一项


> 注意：不要使用子用户accessskey，目前阿里的Go语言SDK暂不支持子用户key