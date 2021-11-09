package common 

import (
    "go/ast"
    "go/token"
)

// 用于分析的文件
type FILE struct {
    F        *ast.File
    Fset     *token.FileSet
    Src      []byte
    Filename string
}

//  type RESULT 用于返回结果的格式
type RESULT struct {    
    CheckerName            string
    ErrPosition            ast.Node
    ErrMsg                string
    Confidence            float64
    HasErr                bool  // 设为true
}

// 遍历函数
type walker func(ast.Node) bool

// 遍历语法树
func (w walker) Visit(node ast.Node) ast.Visitor {
    if w(node) {
        return w
    }
    return nil
}

// 
func Walk(fn func(ast.Node) bool, file FILE) {
    ast.Walk(walker(fn), file.F)
}