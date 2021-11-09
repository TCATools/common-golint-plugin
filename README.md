# 编写规范
## 结构及编译
目前暂不支持windows使用自定义规则，但工具不涉及编译所以windows用户可以使用codedog——linux机器扫描
首先把 common 文件夹放到 $GOROOT/src 目录下
在plugins 目录下创建相应的checker.go 文件目录结构为
```
-plugins
    |---mychecker
            |--- linux // 存放linux编译的so文件
            |       |--- mychecker.so
            |--- mac  // 存放mac编译的so文件
            |       |--- mychecker.so
            |--- src  // 存放源代码
            |       |--- mychecker.go
```
编译命令：
`go build -buildmode=plugin mychecker.go`
## 编写格式
报错信息格式为
```
type RESULT struct {
    CheckerName            string  
    // 规则命名为 custom/XXXXX  开头为custom/  后续只支持字母，数字以及下划线， 名称和目录以及.so文件名字保持一致
    ErrPosition            ast.Node
    // ast节点
    ErrMsg                string
    // 错误信息
    Confidence            float64
    // 暂时未使用设置为0.8, 规则可信度
    HasErr                bool
    // 设置为true
}
可以参考if_checker
```
提交代码到master后，在codedog->工具->golint 添加规则，规则实际名称请和 `CheckerName` 保持一致, 展示名称可自定义。规则其他信息可自由选择
# 参考文档
go 语言ast文档https://golang.org/pkg/go/ast/
go ast语法书可视化工具https://yuroyoro.github.io/goast-viewer/index.html
实例文档 https://docs.qq.com/doc/DQ2JyakZQT1NZeU1a
# 联系人
anakinliu