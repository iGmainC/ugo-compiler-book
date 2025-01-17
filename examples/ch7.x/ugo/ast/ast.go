package ast

import (
	"github.com/chai2010/ugo/token"
)

// AST中全部结点
type Node interface {
	Pos() token.Pos
	End() token.Pos
	node_type()
}

// Package 表示一个包中全部的文件AST
type Package struct {
	Name  string           // 包名
	Files map[string]*File // 文件名(不包含目录路径)
}

// File 表示 µGo 文件对应的语法树.
type File struct {
	Filename string // 文件名
	Source   string // 源代码

	Pkg     *PackageSpec  // 包信息
	Imports []*ImportSpec // 导入包信息
	Consts  []*ConstSpec  // 常量信息
	Globals []*VarSpec    // 全局变量
	Funcs   []*Func       // 函数列表
}

// 包信息
type PackageSpec struct {
	PkgPos  token.Pos // package 关键字位置
	NamePos token.Pos // 包名位置
	Name    string    // 包名
}

// ImportSpec 表示一个导入包
type ImportSpec struct {
	ImportPos token.Pos
	Name      *Ident
	Path      *BasicLit
}

// 基础类型面值
type BasicLit struct {
	ValuePos  token.Pos
	ValueType token.TokenType
	ValueLit  string
	Value     interface{}
}

// 常量信息
type ConstSpec struct {
	ConstPos token.Pos // const 关键字位置
	Name     *Ident    // 常量名字
	Type     *Ident    // 常量类型, 可省略
	Value    Expr      // 常量表达式
}

// 变量信息
type VarSpec struct {
	VarPos token.Pos // var 关键字位置
	Name   *Ident    // 变量名字
	Type   *Ident    // 变量类型, 可省略
	Value  Expr      // 变量表达式
}

// 函数信息
type Func struct {
	FuncPos token.Pos
	NamePos token.Pos
	Name    string
	Body    *BlockStmt
}

// 块语句
type BlockStmt struct {
	Lbrace token.Pos // '{'
	List   []Stmt
	Rbrace token.Pos // '}'
}

// Stmt 表示一个语句节点.
type Stmt interface {
	Node
	stmt_type()
}

// 表达式语句
type ExprStmt struct {
	X Expr
}

// AssignStmt 表示一个赋值语句节点.
type AssignStmt struct {
	Target Expr            // 要赋值的目标
	OpPos  token.Pos       // ':=' 的位置
	Op     token.TokenType // '=' or ':='
	Value  Expr            // 值
}

// Expr 表示一个表达式节点。
type Expr interface {
	Node
	expr_type()
}

// Ident 表示一个标识符
type Ident struct {
	NamePos token.Pos
	Name    string
}

// Number 表示一个数值.
type Number struct {
	ValuePos token.Pos
	ValueEnd token.Pos
	Value    int
}

// BinaryExpr 表示一个二元表达式.
type BinaryExpr struct {
	OpPos token.Pos       // 运算符位置
	Op    token.TokenType // 运算符类型
	X     Expr            // 左边的运算对象
	Y     Expr            // 右边的运算对象
}

// UnaryExpr 表示一个一元表达式.
type UnaryExpr struct {
	OpPos token.Pos       // 运算符位置
	Op    token.TokenType // 运算符类型
	X     Expr            // 运算对象
}

// ParenExpr 表示一个圆括弧表达式.
type ParenExpr struct {
	X Expr // 圆括弧内的表达式对象
}

// CallExpr 表示一个函数调用
type CallExpr struct {
	FuncName *Ident    // 函数名字
	Lparen   token.Pos // '(' 位置
	Args     []Expr    // 调用参数列表
	Rparen   token.Pos // ')' 位置
}

// SelectorExpr 表示 x.Name 属性选择表达式
type SelectorExpr struct {
	X   Expr
	Sel *Ident
}
