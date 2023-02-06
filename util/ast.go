package util

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

func Parse() {
	src, _ := os.Open("util/file.go")
	f := token.NewFileSet()
	p, err := parser.ParseFile(f, "file.go", src, parser.ParseComments)
	if err != nil {
		return
	}
	for _, v := range p.Decls {
		var (
			structComment string
			structName    string
			fieldNumber   int
		)
		if stc, ok := v.(*ast.GenDecl); ok && stc.Tok == token.TYPE {
			//fmt.Println(stc.Tok) //import、type、struct...
			if stc.Doc != nil {
				structComment = strings.TrimRight(stc.Doc.Text(), "\n")
			}
			for _, spec := range stc.Specs {
				if tp, ok := spec.(*ast.TypeSpec); ok {
					structName = tp.Name.Name
					fmt.Println("结构体名称:", structName)
					fmt.Println("结构体注释:", structComment)
					if stp, ok := tp.Type.(*ast.StructType); ok {
						if !stp.Struct.IsValid() {
							continue
						}
						fieldNumber = stp.Fields.NumFields()
						fmt.Println("字段数:", fieldNumber)
						for _, field := range stp.Fields.List {
							var (
								fieldName    string
								fieldType    string
								fieldTag     string
								fieldTagKind string
								fieldComment string
							)
							//获取字段名
							if len(field.Names) == 1 {
								fieldName = field.Names[0].Name //等同于field.Names[0].String()) //获取名称方法2
							} else if len(field.Names) > 1 {
								for _, name := range field.Names {
									fieldName = fieldName + name.String() + ","
								}
							}
							if field.Tag != nil {
								fieldTag = field.Tag.Value
								fieldTagKind = field.Tag.Kind.String()
							}
							if field.Comment != nil {
								fieldComment = strings.TrimRight(field.Comment.Text(), "\n")
							}
							if ft, ok := field.Type.(*ast.Ident); ok {
								fieldType = ft.Name
							}
							fmt.Println("\t字段名:", fieldName)
							fmt.Println("\t字段类型:", fieldType)
							fmt.Println("\t标签:", fieldTag, "标签类型", fieldTagKind)
							fmt.Println("\t字段注释:", fieldComment)
							fmt.Println("\t----------")
						}
					}
				}
			}
		}
	}
}
