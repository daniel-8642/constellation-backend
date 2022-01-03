# 星座运势数据后端

#### 介绍

这个项目是一个星座运势查询项目,分为 前台 后台 后端 三个项目

#### 软件架构

本项目前后端分离,

有简单的前台星座运势查询页面,

后端给前台提供星座运势查询接口,给后端提供账户管理,星座运势查询数据统计接口,带有鉴权,时间戳,与登录状态session,权限管理功能,

后台主要使用 Ant Design Vue 组件库与 ECharts 搭建

后台使用Vue CLI 仿照Ant Design Vue PRO构建

登录页面风格仿照Stisla 登录界面



#### 项目文件

Command 存放接口请求处理方法

​	data.go 数据分析接口

​	star.go 星座运势接口

​	user.go 用户接口

Global 类似Utils 存放工具

​	DataBase.go 获取数据库连接 

​	getConfig.go 获取配置项

Proxy

​	auth.go 鉴权中间件

​	cors.go 跨域中间件

​	log.go 日志中间件

​	timestamp 时间戳校验中间件

Routers 存放路由配置文件

static 存放静态页面文件

config.yaml 软件配置文件

initTable.sql  数据库 表定义

main.go

#### 启动项目

1. 配置数据库

   1. 安装mysql或MariaDB
   2. 确保数据库可以连接,将连接数据填入config.yaml
   3. 在数据库新建database(项目中使用的数据库名是starWeb),将数据库名复制到config.yaml的Mysql.database (注意yaml空格与tab敏感, :后需要跟一个空格)
   4. use 你新建的database, 执行initTable.sql建立表

2. 启动服务

   1. (如果没有修改前台代码,可忽略本步骤)npm build 前台项目,将生成的dist文件夹中文件复制到static/web文件夹(路径与config.yaml 的Web.static_web对应)覆盖本项目中的文件
   2. (如果没有修改后台代码,可忽略本步骤)npm build 后台项目,将生成的dist文件夹中文件复制到static/backend文件夹(路径与config.yaml 的Web.static_backend对应)覆盖本项目中的文件
   3. (如果没有修改本项目中代码,可直接使用编译好的可执行文件),安装go环境,使用go build编译本项目
   4. 运行可执行文件, 运行时需要文件同级存放有config.yaml配置文件,static文件夹,配置文件指向的数据库可连接.

3. 到这里,你的项目应该可以正常运行了,

   星座运势的接口是使用免费的聚合数据的接口.

   配合缓存节约请求次数,你也可以到聚合数据的官网([聚合数据](https://www.juhe.cn/))申请自己的key


最近我在学习分布式,和微服务相关的东西

熬了一晚的夜完成了Dockerfile

并且在我的云服务器上成功的用Docker运行本项目

现在下面的测试页就是使用docker镜像部署的



项目效果

前台:http://star.86428642.xyz/

后台:http://back.86428642.xyz/

项目地址
https://gitee.com/daniel8642/constellation-fortune-backend

前台项目:
https://gitee.com/daniel8642/vue-constellation-chart

后台项目:
https://gitee.com/daniel8642/constellation-chart-backend

