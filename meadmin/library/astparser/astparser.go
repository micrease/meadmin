package astparser

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"reflect"
	"runtime"
	"strings"
)

type StructInfo struct {
	PkgPath     string
	FilePath    string
	StructName  string
	Commit      *Commit
	ChildFields []*StructField
}

type FuncInfo struct {
	PkgPath     string
	FilePath    string
	ClassName   string
	ClassCommit *Commit
	FuncName    string
	Commit      *Commit
}

type Commit struct {
	TopList  []string
	BackList []string
}

type StructField struct {
	Path        string
	Index       int
	FiledName   string
	Type        string
	Kind        string
	Tag         string
	Anonymous   bool //是否匿名结构,就是继承的意思
	Value       interface{}
	Commits     *Commit
	ChildFields []*StructField
	T           *StructInfo
}

func ParseFunc(projectPath string, fn interface{}) (*FuncInfo, error) {
	if reflect.TypeOf(fn).Kind() != reflect.Func {
		//return
	}

	funcInfo := new(FuncInfo)
	//meadmin/application/controller/admin.(*ActivityBBDepositController).PageList-fm
	funcName := NameOfFunction(fn)
	arr := strings.Split(funcName, ".")

	if len(arr) < 3 {
		return funcInfo, errors.New("Get function error")
	}

	funcPath := arr[0]
	ctrl := arr[1]
	ctrl = strings.TrimLeft(ctrl, "(")
	ctrl = strings.TrimLeft(ctrl, "*")
	ctrl = strings.TrimRight(ctrl, ")")
	funcName = strings.TrimRight(arr[2], "-fm")

	path := ParentDir(projectPath) + "/" + funcPath
	funcInfo.PkgPath = path
	funcInfo.ClassName = ctrl
	funcInfo.FuncName = funcName

	fs := token.NewFileSet()
	pkgAst, err := parser.ParseDir(fs, path, filterFn, parser.ParseComments)

	for pkg, astPkg := range pkgAst {
		fmt.Println("Pckage:", pkg)
		for filePath, file := range astPkg.Files {
			for _, decl := range file.Decls {
				name := ""
				if genDecl, ok := decl.(*ast.GenDecl); ok {
					for _, spec := range genDecl.Specs {
						if typeSpec, ok := spec.(*ast.TypeSpec); ok {
							if typeSpec.Name != nil {
								name = typeSpec.Name.Name
							}
						}
					}

					if name != funcInfo.ClassName {
						continue
					}

					commits := Commit{}
					commits.TopList = make([]string, 0)
					if genDecl.Doc != nil {
						for _, v := range genDecl.Doc.List {
							if len(v.Text) > 0 {
								commits.TopList = append(commits.TopList, v.Text)
							}
						}
					}
					funcInfo.ClassCommit = &commits
				}

				if funcDecl, ok := decl.(*ast.FuncDecl); ok {
					if funcDecl.Name.Name != funcInfo.FuncName {
						continue
					}
					if funcDecl.Doc != nil {
						astType := funcDecl.Recv.List[0].Type
						if astStarType, ok := astType.(*ast.StarExpr); ok {
							if astIden, ok := astStarType.X.(*ast.Ident); ok {
								if astIden.Name == funcInfo.ClassName {
									//=====这里才是符合条件的=====
									funcInfo.Commit = new(Commit)
									funcInfo.Commit.TopList = make([]string, 0)
									for _, doc := range funcDecl.Doc.List {
										funcInfo.Commit.TopList = append(funcInfo.Commit.TopList, doc.Text)
									}
									funcInfo.FilePath = filePath
								}
							}
						}
					}
				}
			}
		}
	}
	return funcInfo, err
}

func ParseStruct(projectPath string, obj interface{}) (*StructInfo, error) {
	if obj == nil {
		return nil, nil
	}

	structInfo := new(StructInfo)
	structInfo.ChildFields = []*StructField{}
	m := make(map[string]interface{})
	path := ParentDir(projectPath)

	kind := reflect.TypeOf(obj).Kind()
	var t reflect.Type
	var v reflect.Value
	if kind == reflect.Struct {
		t = reflect.TypeOf(obj)
		v = reflect.ValueOf(obj)
	} else if kind == reflect.Pointer {
		t = reflect.TypeOf(obj).Elem()
		v = reflect.ValueOf(obj).Elem()
	} else if kind == reflect.Slice {
		t = reflect.TypeOf(obj).Elem()
		v = reflect.New(t).Elem()
	}

	if t == nil {
		return nil, errors.New("t is nil")
	}
	pkgPath := path + "/" + t.PkgPath()
	structInfo.PkgPath = pkgPath
	structInfo.StructName = t.Name()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		value, _ := f.Tag.Lookup("json")
		field := v.Field(i)

		filedInfo := StructField{
			Path:      f.PkgPath,
			Index:     i,
			FiledName: t.Field(i).Name,
			Value:     field.Interface(),
			Type:      field.Type().Name(),
			Kind:      field.Kind().String(),
			Tag:       value,
			Anonymous: f.Anonymous,
		}

		if filedInfo.Type != "DBTime" {

			/*
				if filedInfo.Tag == "-" {
					structInfo.ChildFields = append(structInfo.ChildFields, &filedInfo)
					continue
				}
			*/

			if len(filedInfo.Type) == 0 || filedInfo.Kind == "struct" {
				var err error
				filedInfo.T, err = ParseStruct(projectPath, filedInfo.Value)
				if err != nil {
					return structInfo, err
				}

				err = ParseAST(filedInfo.T)
				if err != nil {
					return structInfo, err
				}
			}
		}

		structInfo.ChildFields = append(structInfo.ChildFields, &filedInfo)
		m[value] = v.Field(i).Interface()
	}
	ParseAST(structInfo)
	return structInfo, nil
}

func ParseAstIdent(ident *ast.Ident, parent *StructField) {
	//structField := StructField{}
	if ident.Obj == nil {
		return
	}

	fieldList := []*StructField{}
	if typeSpec, ok := ident.Obj.Decl.(*ast.TypeSpec); ok {
		if structType, ok := typeSpec.Type.(*ast.StructType); ok {
			for fieldIndex, field := range structType.Fields.List {

				fieldType := ""
				if astIdent, ok := field.Type.(*ast.Ident); ok {
					fieldType = astIdent.Name
				}

				tag := ""
				if field.Tag != nil {
					tagStr := strings.Trim(field.Tag.Value, "`")
					if len(tagStr) > 0 {
						structTag := reflect.StructTag(tagStr)
						tag, _ = structTag.Lookup("json")
					}
				}

				topList := []string{}
				if field.Doc != nil {
					for _, doc := range field.Doc.List {
						topList = append(topList, doc.Text)
					}
				}

				commitList := []string{}
				if field.Comment != nil {
					for _, commit := range field.Comment.List {
						commitList = append(commitList, commit.Text)
					}
				}

				structField := new(StructField)

				//匿名结构:field.Names为空
				if field.Names != nil {
					structField = &StructField{
						Index:     fieldIndex,
						FiledName: field.Names[0].Name,
						Type:      fieldType,
						Kind:      fieldType,
						Tag:       tag,
						Commits: &Commit{
							TopList:  topList,
							BackList: commitList,
						},
					}
					fieldList = append(fieldList, structField)
				}

				if childAstIdent, ok := field.Type.(*ast.Ident); ok {
					//struct
					if childAstIdent.Obj != nil {
						ParseAstIdent(childAstIdent, structField)
					}
				} else if astFieldType, ok := field.Type.(*ast.StarExpr); ok {
					structField.Kind = "struct"
					//struct pointer, *a
					if astIdent, ok := astFieldType.X.(*ast.Ident); ok {
						structField.Type = astIdent.Name

						ParseAstIdent(astIdent, structField)
					} else if astFieldType, ok := astFieldType.X.(*ast.ArrayType); ok {
						//struct pointer array, *a[]
						structField.Kind = "slice"
						if astIdent, ok := astFieldType.Elt.(*ast.Ident); ok {
							structField.Type = astIdent.Name
							ParseAstIdent(astIdent, structField)
						}
					}
				} else if astFieldType, ok := field.Type.(*ast.ArrayType); ok {
					//array, []a
					structField.Kind = "slice"
					if astIdent, ok := astFieldType.Elt.(*ast.Ident); ok {
						structField.Type = astIdent.Name
						ParseAstIdent(astIdent, structField)
					} else if astStarExpr, ok := astFieldType.Elt.(*ast.StarExpr); ok {
						//array pointer, []*a
						if astStarIdent, ok := astStarExpr.X.(*ast.Ident); ok {
							structField.Type = astStarIdent.Name
							ParseAstIdent(astStarIdent, structField)
						}
					}
				} else {
					fmt.Println("不是Ident跳过解析")
				}
			}
		}
	}
	parent.ChildFields = fieldList
}

func ParseAST(structInfo *StructInfo) error {
	//AST解析
	if structInfo == nil {
		return nil
	}

	fs := token.NewFileSet()
	pkgAst, err := parser.ParseDir(fs, structInfo.PkgPath, filterFn, parser.ParseComments)
	if err != nil {
		return err
	}
	fieldLen := len(structInfo.ChildFields)
	for _, astItem := range pkgAst {
		//fmt.Println("包名:", astItem.Name)
		for filePath, file := range astItem.Files {
			//	fmt.Println("文件名:", filePath)
			structInfo.FilePath = filePath
			//	fmt.Println("Doc:", file.Doc.Text())
			for _, decl := range file.Decls {
				if genDecl, ok := decl.(*ast.GenDecl); ok {
					for _, spec := range genDecl.Specs {
						if typeSpec, ok := spec.(*ast.TypeSpec); ok {

							//有参数的结构体
							//TODO::structInfo.StructName 分页
							if typeSpec.TypeParams != nil && len(structInfo.StructName) == 0 {
								//fmt.Println("这是范型")
								arr := strings.Split(typeSpec.Name.Name, "[")
								structInfo.StructName = arr[0]
							}
							//	fmt.Println("typeSpec字段:", typeSpec.Name.Name, "Struct字段:", structInfo.Name)
							if typeSpec.Name.Name == structInfo.StructName {
								//fmt.Println("字段:", structInfo.StructName)
								topList := []string{}
								if genDecl.Doc != nil {
									for _, doc := range genDecl.Doc.List {
										topList = append(topList, doc.Text)
									}
									structInfo.Commit = &Commit{TopList: topList}
								}

								if structType, ok := typeSpec.Type.(*ast.StructType); ok {
									for fieldIndex, astField := range structType.Fields.List {
										if fieldIndex >= fieldLen {
											continue
										}

										tag := ""
										if astField.Tag != nil {
											tagStr := strings.Trim(astField.Tag.Value, "`")
											if len(tagStr) > 0 {
												structTag := reflect.StructTag(tagStr)
												tag, _ = structTag.Lookup("json")
											}
										}

										if tag == "-" {
											continue
										}

										//var commits = []string{}
										topList := []string{}
										if astField.Doc != nil {
											for _, doc := range astField.Doc.List {
												topList = append(topList, doc.Text)
											}
										}
										commitList := []string{}
										if astField.Comment != nil {
											for _, commit := range astField.Comment.List {
												commitList = append(commitList, commit.Text)
											}
										}
										commits := Commit{
											TopList:  topList,
											BackList: commitList,
										}
										structInfo.ChildFields[fieldIndex].Commits = &commits
										if astFieldType, ok := astField.Type.(*ast.StarExpr); ok {
											if astIdent, ok := astFieldType.X.(*ast.Ident); ok {
												structInfo.ChildFields[fieldIndex].Kind = "struct"
												structInfo.ChildFields[fieldIndex].Type = astIdent.Name
												//fmt.Println("StarExpr->astIdent", fieldIndex, astIdent.Name, astIdent.Obj.Kind.String(), astIdent.Obj.Decl)

												if len(structInfo.ChildFields[fieldIndex].Type) <= 0 {
													structInfo.ChildFields[fieldIndex].Type = astIdent.Name
												}

												ParseAstIdent(astIdent, structInfo.ChildFields[fieldIndex])
											} else if astFieldType, ok := astFieldType.X.(*ast.ArrayType); ok {
												structInfo.ChildFields[fieldIndex].Kind = "slice"
												if astIdent, ok := astFieldType.Elt.(*ast.Ident); ok {

													if len(structInfo.ChildFields[fieldIndex].Type) <= 0 {
														structInfo.ChildFields[fieldIndex].Type = astIdent.Name
													}

													//fmt.Println("StarExpr->astFieldType->astIdent", fieldIndex, astIdent.Name, astIdent.Obj.Kind.String(), astIdent.Obj.Decl)
													ParseAstIdent(astIdent, structInfo.ChildFields[fieldIndex])
												}
											}
										}
										//array
										if astFieldType, ok := astField.Type.(*ast.ArrayType); ok {
											if astIdent, ok := astFieldType.Elt.(*ast.Ident); ok {
												//范性数组
												if len(structInfo.ChildFields[fieldIndex].Type) <= 0 {
													structInfo.ChildFields[fieldIndex].Type = astIdent.Name
												}

												ParseAstIdent(astIdent, structInfo.ChildFields[fieldIndex])
											} else if astStarExpr, ok := astFieldType.Elt.(*ast.StarExpr); ok {
												if astStarIdent, ok := astStarExpr.X.(*ast.Ident); ok {

													if len(structInfo.ChildFields[fieldIndex].Type) <= 0 {
														structInfo.ChildFields[fieldIndex].Type = astStarIdent.Name
													}

													ParseAstIdent(astStarIdent, structInfo.ChildFields[fieldIndex])
												}
											}
										}
										//object,function
										if astIdent, ok := astField.Type.(*ast.Ident); ok {
											if astIdent.Obj != nil {
												if isFunc(astIdent) {
													fmt.Println("is function")
												} else {
													structInfo.ChildFields[fieldIndex].Type = astIdent.Name
													structInfo.ChildFields[fieldIndex].Kind = "struct"
													if len(structInfo.ChildFields[fieldIndex].Type) <= 0 {
														structInfo.ChildFields[fieldIndex].Type = astIdent.Name
													}
													ParseAstIdent(astIdent, structInfo.ChildFields[fieldIndex])
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return nil
}

func isFunc(ident *ast.Ident) bool {
	if ident.Obj == nil {
		return false
	}
	if ident.Obj.Decl == nil {
		return false
	}
	if typeSpec, ok := ident.Obj.Decl.(*ast.TypeSpec); ok {
		if typeSpec.Type == nil {
			return false
		}
		if funcType, ok := typeSpec.Type.(*ast.FuncType); ok {
			fmt.Println("function", funcType)
			return true
		}
	} else {
		return false
	}
	return false
}

func filterFn(info fs.FileInfo) bool {
	return true
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func ParentDir(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func NameOfFunction(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func (c Commit) ToString() string {
	arr := make([]string, 0)
	for _, top := range c.TopList {
		//@这些不标准的注释写法将不会同步到文档中
		//@也就是说如果你不想同步到文档中的注释就不要写的太标准
		if strings.Contains(top, "///") {
			continue
		}

		if strings.Contains(top, "//@") {
			continue
		}

		if strings.Contains(top, "//#") {
			continue
		}

		if strings.Contains(top, "/**") {
			continue
		}

		if strings.Contains(top, "**/") {
			continue
		}

		if strings.Contains(top, "*//") {
			continue
		}

		top := strings.ReplaceAll(top, "//", "")
		top = strings.ReplaceAll(top, "/*", "")
		top = strings.ReplaceAll(top, "*/", "")
		top = strings.ReplaceAll(top, "*", "")
		arr = append(arr, top)
	}

	for _, back := range c.BackList {
		if strings.Contains(back, "///") {
			continue
		}

		if strings.Contains(back, "//@") {
			continue
		}

		if strings.Contains(back, "//#") {
			continue
		}

		back := strings.ReplaceAll(back, "//", "")
		arr = append(arr, back)
	}
	return strings.Join(arr, ",")
}
