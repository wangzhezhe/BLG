# BLG (Blocks in Golang)

平时写程序的时候，很多时候的时间都花在一些基本实现的细节问题的确认上，比如某个函数的参数是怎么样的，由于无法时刻记忆清除所有细节，对于具体语法是怎样实现的这种问题上总要花费很多精力。这个仓库的主要目的是把平时使用Golang的一些局部的功能（所谓的Blocks）整合在一起，这些代码大多来自于不同的项目（所谓的不生产代码仅仅是代码的搬运工），但都是自己认为的一些好的实践，随着不断的迭代整理，相信在项目中可以对于golang有更灵活的使用。

### BasicUsings

这一部分主要是基本数据结构在Golang中的使用 比如map使用时候的注意事项，struct使用时候的一些技巧，List使用时候的一些操作等等。更多的是一些语法层面的内容。

### BasicFunctions

这一部分更多的是一些常用到的局部小功能和技巧，比如json的解析，encode以及decode时候应该注意的技巧，定时器，goroutine的使用技巧等等。

### Components

这一部分是更完整的功能，有的是整理一些比较好的实践，比如日志部分的集中实践方式，参数传入的几种实践方式，webserver的一些特定功能，socket，rpc通信的实现方式等等。

### 参考资源

[awsomego](https://github.com/hackstoic/golang-open-source-projects)汇集了各种各样的golang开源项目
