# TeleBaiduPan
基于Telegram Bot的百度网盘管理工具

## 安装使用  
目前支持Linux与Windows部署
在

# TeleBaiduPan
基于Telegram Bot的百度网盘管理工具

## 安装使用  
目前支持Linux与Windows部署
在config.yaml文件下填写User部分信息与tginfo部分信息

```yaml
user:
  #百度账号BDUSS,建议使用网页版Cookie
  bduss:
  #对下载无明显影响,可以不改
  is_vip: 0
  #奇妙小链接,不清楚就留空
  acclink:
```
```yaml
tginfo:
    #telegram bot token
    bottoken:
    #你的telegramID
    userid:
    #开启仅本人访问模式,关闭后会向全部用户开放获取文件列表和下载的权限
    privacy: true
```

## 功能
### V1.0.0版本
    1. 支持文件列表查看
    2. 支持文件下载
    3. 支持分享链接转存
    4. 支持文件管理
    5. 支持公共仓库模式(仅查看和下载)

## FaQ
如果有任何使用问题请在群组[@KinhWeb](https://t.me/kinhweb)提出,或者在本仓库下提交Issue.  
关注频道[@KinhWebPD](https://t.me/kinhwebpd)获取最新更新信息.  

## License
GPL-2.0 license
