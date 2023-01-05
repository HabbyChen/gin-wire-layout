# gin wire layout 
# 背景
> 解决以下几个问题
> 1. 我们是一个人负责多个系统的多个项目,每个项目都很小,我们的总体开发趋势是几个人负责一堆项目,每个项目内又有多个程序.人和项目之间必然出现高频穿插,交接成本巨大.
> 2. 我们海量的项目分散在不同的git库中,不同git库的技术细节可能又不一致,技术路线,代码质量无法保证
> 3. 肉眼可见的,我们手里会有越多的项目,基本上也不会出现需要多人维护一个的大型项目.
> 4. 我们必须人员复用,不可能永远一个人做用户中心,一个人做跨院调阅,不复用人员导致有的没事儿做,有的事儿堆着.
> 基于以上的背景,有如下考虑:
> 使用wire+类似kratos的layout来管理项目结构
> 继续使用gin框架
> 本项目结构和基础功能主要用于web api接口,或者http的api调用,不适合用作rpc的服务端,如果是用作rpc的接口定义,使用其他的框架或技术规范.

# 项目结构介绍
参考了
Standard Go Project Layout
https://github.com/golang-standards/project-layout/blob/master/README_zh.md


## 结构树状图
```
├── cmd
│   ├── app
│   │   ├── main.go  // 项目入口
│   │   └── wire.go // wire依赖注入
│   │   └── wire_gen.go // wire依赖注入自动生成的文件
│   │   └── boot.yml    // 项目启动配置文件,因为不同的项目可能启动参数不一样
│   ├── app2    // 项目2
│   │   ├── main.go  // 项目入口
│   │   └── wire.go // wire依赖注入
│   │   └── wire_gen.go // wire依赖注入自动生成的文件
│   │   └── boot.yml    // 项目启动配置文件,因为不同的项目可能启动参数不一样
│   ├── app...    // 项目...
├── internal
│   ├── app1
│   │   ├── configs
│   │   │   └── config.go // 配置文件解析,每个文件独有的解析结构
│   │   ├── server
│   │   │   └── server.go // server文件,一般都是一些服务启动相关的,比如gin的engin启动,grpc的server启动
│   │   ├── service //其中包含了DTO定义,以及DO的与dto的转换
│   │   │   └── service.go // service的SET文件
│   │   │   └── user_service.go // 用户相关的service如此命名,
│   │   ├── biz //其中包含了DO定义,以及repo接口的定义
│   │   │   └── biz.go // biz的SET文件
│   │   │   └── user_biz.go // 如果说有用户相关的biz,就如此命名,
│   │   ├── repo //在kratos的目录中是没有这一层的.新加的一层对biz的repo的实现.biz中repo接口的实现,以及PO和数据对象向的转换.
│   │   │   └── repo.go // repo的SET文件
│   │   │   └── user_repo.go // 如果说有用户相关的repo,就如此命名
│   │   ├── data   //repo依赖的数据层
│   │   │   └── user_data.go // 如果有user相关的data,就如此命名,一般都是表结构
│   ├── app2
├── pkg 公共类库
│   ├── user
│   │   ├── const  //常量定义
│   │   ├── repo  //业务公共库
│   │   ├── entity  //公共的模型库
│   ├── safety
│   ├── image
│   ├── frame 基础框架
│-- vendor 应用程序依赖项
├── docs 文档
├── scripts 脚本
├── third_party 第三方依赖
```
## internal内部结构定义
![框架层级设计.png](%E6%A1%86%E6%9E%B6%E5%B1%82%E7%BA%A7%E8%AE%BE%E8%AE%A1.png)

## 