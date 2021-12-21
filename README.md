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

config.yaml

initTable.sql

main.go





项目效果

前台:http://star.86428642.xyz/

后台:http://back.86428642.xyz/

项目地址
https://gitee.com/daniel8642/constellation-fortune-backend

前台项目:
https://gitee.com/daniel8642/vue-constellation-chart

后台项目:
https://gitee.com/daniel8642/constellation-chart-backend

