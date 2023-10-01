# NewMallApi
 Golang基于gin的商城项目

## 参考

1. [后端参考](https://github.com/newbee-ltd/newbee-mall-api-go)
2. [gin框架]()

## 涉及

1. gin框架
2. Zap日志文件系统
3. Viper 应用程序配置
4. Mysql及redis
5. Md5加密以及token、jwt验证

## 结构

|文件夹  |说明        |功能                      |
|------|---------|---------------|
|model  | 模型层      | 用来对数据库进行操作|
|router  | 路由层      | 获取请求并分配到service进行处理|
|service | 逻辑处理层 | 进行业务逻辑处理                     |
|utils     | 工具层       | 存放编写的辅助工具代码            |



## 界面

1. register 注册界面  
2. login 登陆界面
3. index 首页
4. 





## 运行方法如下


## 项目步骤及部分代码如下

1. 初始化文件、导入gin框架及创建目录

```golang
# 创建文件
mkdir NewMallApi
# 初始化文件
go mod init
# 导入gin框架
go get -u github.com/gin-gonic/gin
# 创建基础目录
model、router、service、utils
```




