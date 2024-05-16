# TT双色球预测系统

## 简介
TT系统会预测下一次双色球开奖号码，并将预测的结果通过email形式发送给用户。
系统内置4个预测策略：完全随机预测、正向关联预测、负向关联预测、正负关联预测。

系统默认在每周2，周4，周7 12点将预测结果通过邮件形式发送给用户。


## 使用
1. 执行build.sh编译
2. 执行 cp ./tticket.toml /etc/tticket/tticket.toml并填写相关参数
    ```
   [mysql]
    host = ""
    port = ""
    username = ""
    password = ""
    log_location = "" # db日志位置
    
    [mail]
    address = "" # 邮件服务器host
    port = "" # 邮件服务器端口
    username = ""
    password = ""
    host = "" # 邮件服务器host
   ```
   
3. 在mysql中执行./pkg/dal/db.sql命令，初始化表与db数据
4. 在mysql user表中插入用户的name与邮箱
5. 启动编译的二进制文件


## 改进List
1. 添加HTTP接口，用于人工调用接口触发各种定时任务：爬虫任务、缓存任务、预测任务、邮件任务等
2. 优化配置项，各种可配置的信息加入tticket.toml中
3. 模块微服务划分：tuser、tspider、tmail、tstrategy等
4. 添加后台程序：如数据分析、用户订阅管理、爬虫数据人工录入或订、用户自定义预测策略正等。