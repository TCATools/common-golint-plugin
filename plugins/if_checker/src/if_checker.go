package main

// // No C code needed.
import "C"

import (
    "container/list"
    "common"
	"go/ast"
)
// 
var GoFile common.FILE  //

// 返回结果
var ResultList list.List
// 
func checkIfStamt(n *ast.BinaryExpr) bool {
	op := n.Op.String()
	if op ==  "||" || op == "&&" || op == "^" || op == "&^" {
		x, ok := n.X.(*ast.BinaryExpr)
		if ok {
			checkIfStamt(x)
		}
		y, ok := n.Y.(*ast.BinaryExpr)
		if ok {
			checkIfStamt(y)
		}
		return true
	}
	x := n.X
	// y := binExpr.Y
	x_type, ok := x.(*ast.Ident)
	if ok && (x_type.Name == "nil" || x_type.Name == "true" || x_type.Name == "false"){
		// var err common.RESULT
		ResultList.PushBack(common.RESULT{CheckerName: "custom/if_checker",
		ErrPosition: x,
		ErrMsg: "if 对两个值进行判断时，约定如下顺序：变量在左，常量在右",
		Confidence: 0.8,
		HasErr: true})
		return true
	}
	_, ok = x.(*ast.BasicLit)
	if ok {
		// var err common.RESULT
		ResultList.PushBack(common.RESULT{CheckerName: "custom/if_checker",
		ErrPosition: x,
		ErrMsg: "if 对两个值进行判断时，约定如下顺序：变量在左，常量在右",
		Confidence: 0.8,
		HasErr: true})
		return true
	}
	// y_type, ok = y.(*ast.)
	// An error return parameter should be the last parameter.
	// Flag any error parameters found before the last.
	return true
}

// 分析函数
func MyLintChecker()  {
    common.Walk(func(n ast.Node) bool {
		fn, ok := n.(*ast.IfStmt)
		if !ok{
			return true
		}
		cond := fn.Cond
		if cond == nil{
			return true
		}
		binExpr, ok := cond.(*ast.BinaryExpr)
		if !ok{
			return true
		}
		checkIfStamt(binExpr)
		return true
    }, GoFile)
}
