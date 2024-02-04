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

## 源码阅读

### slice

对~不是很理解
~T 表示底层类型为T类型的值传入也不会引起错误，比如type intX int 将[]intX和[]int传入参数不会引起错误。
在slice的源码中的使用形式为S ~[] E代表底层类型为 []T的值传入也不会引起错误

Equal比较的是底层类型相同的切片之间是否完全相等。

EqualFunc比较的是底层类型不同两个切片是否完全相同。