### 一：包介绍
1. go语言的包是多个go源码的集合，内置包：fmt、os、io，可以自定义包
2. 在使用其他包的对象或者函数时，首先要引入import

### 二：自定义包
1. 在项目目录下新建一个目录，如/tools
2. 在/tools目录下创建一个math.go
3. 编写math工具代码
4. 通过go mod初始化包
5. 在其他文件中使用自定义包

### 三：go mod
1. go modules是golang1.11 新加的特性
2. 模块是相关包的集合。modules是源代码交换和版本控制的单元。
3. go 命令直接支持使用modules，包括记录和解析对其他模块的依赖性
4. modules替换旧的基于GOPATH的方法指定在给定构建中使用那些源文件。
```
go mod init         生成go.mod文件
go mod download     下载go.mod文件中指明的所有的依赖
go mod tidy         整理现有的依赖
go mod graph        查看现有的依赖结构
go mod edit         编辑go.mod文件
go mod vendor       导出项目所有的依赖到vendor目录
go mod verify       校验一个模块是否被篡改过
go mod why          查看为什么需要依赖某模块
```

### 四：第三包的使用
1. https://pkg.go.dev/ 搜索包及查看使用说明
2. 搜索decimal
3. 点击搜索结果打开decimal包详情页面，找到安装命令
4. 安装完成使用decimal工具包