# 打包和工具链

- 如何组织Go代码
- 使用Go自带工具命令
- 与其他开发者合作

## 包

设计理念是使用包来封装不同的语义单元功能，更好的复用代码。

- 包名

    包名及目录命名使用，简洁，清晰，全小写的名字
- main包

    编译时会把包含main名字的包编译为二进制可执行文件。所有的Go语言编译的可执行程序必须有一个叫main的包
    main()函数是程序的入口,会使用main包代码所在目录目录名作为二进制可执行文件的文件名
- 命令 command 指可执行的程序。包值语义上可以导入的功能单元

例子 编写$GOPATH/src/test/hello.go
里面的package 写的是main 生成一个二进制文件名字叫test

    ```
    package main
    import "fmt"
    func main()
    fmt.Println("oooo")
    ```

如果$GOPATH/src/test/hello.go
里面的package 写的是hello并且包含了func main() 是无效的，认为它是一个包不是命令

## 导入

- import关键字

    告诉编译器到磁盘的哪里去找要到包

- 一旦找到一个满足import 语句的包就是停止查找，先找安装Go的目录然后才找GOPATH变量里面的目录
- 远程导入

    如果是url 工具链使用DVCS获取包。使用go get命令完成，go get将获取任意指定url的包 具有递归属性 扫描源码树
- 命名导入

    重名的包使用命名导入

    ```
        package main
        import (
            "fmt"
            myfmt "mylib/fmt"
        )
    ```
- 导入不需要引用这个包的标识符

    使用下划线 (_) 空白标识符

## 函数init

    每个包可以包含任意个init函数并且都在执行开始时被被调用。
都在main函数之前执行。
main.go 使用postgres.go init函数及空白符使用

## 工具命令

- go run 先构建程序 然后执行
- go build 先编译 然后执行
- go vet 检查错误
- go fmt 格式化代码
- go doc 查看文档 go doc tar …… 展示tar命令
- godoc -http=:6060启动web服务
- 注释风格 /* */ //
- 函数的文档

    写直接写在声明之前// 如果是大段文字 可以在工程里
    写一个叫做doc.go的文件 使用同样的包名
    并把包的介绍使用注释/**/加载报名之前

## 与其他Go合作开发

- 包应该在代码库的根目录中

    go get 指定了要导入的全部路径，意味着包名就是代码库的名字且包名应该位于代码库目录结构的根目录
- 包可以非常小
- 对代码指定 go fmt
- 写文档

## 依赖管理

- godep vender gopkg.in 工具

## 第三方依赖库

godep 和vender 第三方（verdoring）导入路径重写特性解决了依赖问题。
思想是把所有的依赖包复制到工程代码目录里，然后使用工程内部的依赖包在所在目录来重写所有的导入路径
- mod

- gb 第三方代码 vendored code 工具链 有些不兼容
