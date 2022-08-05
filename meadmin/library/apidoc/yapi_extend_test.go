package apidoc

import (
	"fmt"
	yapi "meadmin/library/apidoc/yapiclient"
	"meadmin/library/astparser"
	"meadmin/library/files"
	"reflect"
	"testing"
)

type Page struct {
	PageNo   int `json:"page_no"`   //pageno
	PageSize int `json:"page_size"` //pagesize
}

type A struct {
	Page2 Page
	Name  string `json:"a_name"` //a name commits
	Age   int    `json:"a_age"`  //a age commits
}

type AA struct {
	Page
	Name string `json:"a_name"` //a name commits
	Age  int    `json:"a_age"`  //a age commits
}

//B不是继承于Page
type B struct {
	Page  Page
	NameB string `json:"name_b"` //name b commit
}

//B是继承于Page
type BB struct {
	B
	NameBB string `json:"name_bb"` //name bb commit
}

func TestNoExtendYapiProperty(t *testing.T) {
	a := &A{}
	a.Page2.PageNo = 1 //可以显示赋值
	a.Page2.PageSize = 10
	a.Name = "张三"
	a.Age = 34

	path := files.ProjectDir()
	structInfo, err := astparser.ParseStruct(path, a) //Anonymous为true
	fmt.Println(a, structInfo, err)

	respProperty := &yapi.Property{}
	respProperty.Type = "object"
	commitArr := []string{}

	resBody := createJsonProperty(respProperty, structInfo.ChildFields, &commitArr)
	fmt.Println(resBody)
}

func TestExtendYapiProperty(t *testing.T) {
	a := &AA{}
	a.PageNo = 1 //可以显示赋值
	a.PageSize = 10
	a.Name = "张三"
	a.Age = 34

	path := files.ProjectDir()
	structInfo, err := astparser.ParseStruct(path, a) //Anonymous为true
	fmt.Println(a, structInfo, err)

	respProperty := &yapi.Property{}
	respProperty.Type = "object"
	commitArr := []string{}

	resBody := createJsonProperty(respProperty, structInfo.ChildFields, &commitArr)
	fmt.Println(resBody)
}

//连续继承
type AAA struct {
	AA
	NameAAA string `json:"aaa_name"` //a name commits
	AgeAAA  int    `json:"aaaa_age"` //a age commits
}

func TestKeepExtendYapiProperty(t *testing.T) {
	a := &AAA{}
	a.PageNo = 1 //可以显示赋值
	a.PageSize = 10

	a.Name = "张三"
	a.Age = 34

	a.NameAAA = "aaa"
	a.AgeAAA = 12

	path := files.ProjectDir()
	structInfo, err := astparser.ParseStruct(path, a) //Anonymous为true
	fmt.Println(a, structInfo, err)

	respProperty := &yapi.Property{}
	respProperty.Type = "object"
	commitArr := []string{}

	resBody := createJsonProperty(respProperty, structInfo.ChildFields, &commitArr)
	fmt.Println(resBody)
}

//不连惯的继承:最顶层不是继承
func TestTopBreakExtendYapiProperty(t *testing.T) {
	b := &BB{}
	b.Page.PageNo = 1 //可以显示赋值
	b.Page.PageSize = 10

	b.NameB = "张三"
	b.NameBB = "namec"

	path := files.ProjectDir()
	structInfo, err := astparser.ParseStruct(path, b) //Anonymous为true
	fmt.Println(b, structInfo, err)

	respProperty := &yapi.Property{}
	respProperty.Type = "object"
	commitArr := []string{}

	resBody := createJsonProperty(respProperty, structInfo.ChildFields, &commitArr)
	fmt.Println(resBody)

	resBody = mergeAnonymous(resBody)
	fmt.Println(resBody)
}

//ABA不继承于AA,但AA继承于Page
type ABA struct {
	AA      AA     `json:"aa"`       //aa
	ABAName string `json:"aba_name"` // aba_name commits
}

type ABABA struct {
	ABA
	ABABAName string `json:"ababa_name"` // ababa_name commits
}

//不连惯的继承:中件层不是继承
func TestMidBreakExtendYapiProperty(t *testing.T) {
	b := &ABABA{}
	b.ABA.AA.Page.PageNo = 1 //可以显示赋值
	b.ABA.AA.Page.PageSize = 10

	b.ABA.ABAName = "b.ABA.ABAName"

	b.ABA.AA.Age = 11
	b.ABA.AA.Name = "b.ABA.AA.Name"

	path := files.ProjectDir()
	structInfo, err := astparser.ParseStruct(path, b) //Anonymous为true
	fmt.Println(b, structInfo, err)

	respProperty := &yapi.Property{}
	respProperty.Type = "object"
	commitArr := []string{}

	resBody := createJsonProperty(respProperty, structInfo.ChildFields, &commitArr)
	fmt.Println(resBody)

	resBody = mergeAnonymous(resBody)
	fmt.Println(resBody)
}

type T struct{}

func (t *T) Add(a, b int) {
	fmt.Printf("a + b is %+v\n", a+b)
}
func TestCall(t *testing.T) {
	funcName := "Add"
	typeT := &T{}
	a := reflect.ValueOf(1)
	b := reflect.ValueOf(2)
	in := []reflect.Value{a, b}
	reflect.ValueOf(typeT).MethodByName(funcName).Call(in)
}
