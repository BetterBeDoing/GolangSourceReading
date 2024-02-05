# GolangSourceReading

1.18 引入泛型

静态编译型语言

## 基本语法

### 变量

#### 声明以及初始化

Gopher在变量声明形式的选择上应尽量保持项目范围内一致。

```go
var a int32
//声明并显式初始化
var s string = "hello"
var i = 13
n := 17
var (
    crlf       = []byte("\r\n")
    colonSpace = []byte(": ")
)
```

#### 局部变量

短变量声明

```
a  := 17
```

#### 无类型常量/有类型变量

**不支持隐式转换**

无类型常量的默认值对应关系

```
布尔型常量  整数常量   字符常量       浮点数常量 复数常量    字符串常量
bool        int     int32(rune)  float64  complex128  string
```

**const**

隐式重复前一个非空表达式

**iota**

## source code learning

### package usage

AUTHORS：Golang官方作者清单
CONTRIBUTING.md：加入贡献者队列的指导文件
CONTRIBUTORS：第三方贡献者清单
LICENSE：授权协议
PATENTS：专利
README.boringcrypto.md：因为Golang是Google发布的，这是针对Google内部研究分支的说明
README.md：说明文件，大家都明白，每个开源库都有
SECURITY.md：安全政策
api：Golang每个版本的功能列表归档文件，下面有具体介绍
doc：Golang文档说明，和官方文档相同，可以离线查看
lib：看起来像是库文档模板，里面列举了time包的说明
misc：汇集了Go语言相关的IDE、插件、cgo测试程序、示例等乱七八糟的东西
robots.txt：主要用来控制各大搜索引擎爬虫的爬取规则
src：Golang核心实现都在这里，下面详细讲述
test：Golang单元测试程序，通过查看测试程序可以学习到golang的用法和特性

### api

该目录中的每个文件都是Go语言API列表，每行一个，方便IDE使用。

### src

#### archive&compress

压缩包处理，tar、zip等

#### bufio

文本读写

#### builtin

定义了内置类型、函数和接口，比如make、new、len、error

#### bytes

操作字节

#### cmd

提供了Go的基本工具，比如gofmt、静态检查工具vet

#### container

list，heap，ring 

#### context

goroutine's context

#### crypto

manage the algorithm of encryption

#### database

manage the api of db

#### debug

Go program debug

#### encoding

encoding , eg. base64,json,xml,hex and so on

#### errors

define the error

#### expvar

use the HTTP to print the variable of service with the format of JSON

#### flag

Tools for parsing and processing command line parameters

#### fmt

format output

### go

#### hash

crc32,crc64 hash algorithm

#### html

HTML template engine

#### image

handle the image

#### index

string search

#### internal

used to control the import authority of package

#### io

provide some basic api of file I/O, bufio is implemented based on the package

#### log

some functions about the recording of log

#### math

some functions about the math

#### mime

package the decoding of media

#### net

network I/O, such as tcp, udp, Socket

#### os

operating of os

#### path

handle the /

#### plugin

dynamically load the .so file

#### reflect

package the function of reading the reflect

#### regexp

regex

#### runtime

Go runtime

#### strconv

some basic type convertion to string such as int to sting

#### sync

mutex

#### unsafe

some unsafe operation such as pointer to pointer