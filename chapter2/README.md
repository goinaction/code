# README

## 代码目录

    ```
        /-sample 外部目录
        /-data
            data.json --数据源
        /-matachers
            rss.go --搜索rss源的匹配器
        /-search
            default.go --搜索默认匹配起
            feed.go --用于读取json数据的文件
            match.go --用于支持不同匹配器的接口
            search.go --执行搜索的主控制逻辑
        main.go    -- 程序的入口
        rss.xml -- 示例rss文档 RSS匹配器会下载这些RSS文档。 并将这些返回给results通道
    ```

第二章介绍的的是一个rss 匹配程序内容，介绍从每个代码片段介绍语法规则及现象，介绍众多的使用方法和具体实例。

## 小结

- 每个代码文件都属于一个包，报名应该与代码文件所在文件夹同名
    每个文件下面的文件头部都必须明确写 package + 文件夹名字
- 多种初始化方式和声明的变量的方式 如果值没有显式初始化，会默认为零/nil
    var 初始化那些是空的变量
    := 初始化赋值 也会有类型推断
- 指针可以在函数或者goroutine间共享数据
    例子在search.go 的Run函数中 go func
- 通过启动goroutine和使用通道完成并发和同步
    search.go 使用sync.WaitGroup方法
- 内置函数支持Go语言内部数据结构
    make方法的使用初始化 引用类型的数据比如
    make(map[string]Matcher)
    make(chan *Result)
- 标准库包含很多包
- Go接口可以编写通用的代码和框架
    接口里面定义方法，单一方法使用er结尾，多个使用相关名字
