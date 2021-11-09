package main

// // No C code needed.
import "C"

import (
    "container/list"
    "common"
	"go/ast"
	"fmt"
)

// 
var GoFile common.FILE

//返回分析结果
var ResultList list.List

// 分析函数
func MyLintChecker()  {	 
    common.Walk(func(n ast.Node) bool {
		// 在这里实现分析逻辑，详情见if_check
		fmt.Println("demo")
		return true
    }, GoFile)
}
