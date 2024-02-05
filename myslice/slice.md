
# slice

对~不是很理解
~T 表示底层类型为T类型的值传入也不会引起错误，比如type intX int 将[]intX和[]int传入参数不会引起错误。
在slice的源码中的使用形式为S ~[] E代表底层类型为 []T的值传入也不会引起错误

Equal比较的是底层类型相同的切片之间是否完全相等。

EqualFunc比较的是底层类型不同两个切片是否完全相同。

slice的扩容机制：代码在`runtime.growslice`，目前是256机制会发生改变。