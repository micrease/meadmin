package astparser

import (
	"fmt"
	"giftcard/library/files"
	"testing"
)

type Page struct {
	PageNo   int `json:"page_no"`   //pageno
	PageSize int `json:"page_size"` //pagesize
}

type A struct {
	Page
	Name string `json:"a_name"` //a name commits
	Age  int    `json:"a_age"`  //a age commits
}

//再来一个B
type B struct {
	Page Page
	Name string `json:"a_name"` //b name commits
	Age  int    `json:"a_age"`  //b age commits
}

func TestExtendStruct(t *testing.T) {
	//这种继承结构应该有4个字段,解析出来后我们可以看到structInfo中Anonymous为true,而structInfob中为false
	a := &A{}
	a.Page.PageNo = 1 //可以显示赋值
	a.PageSize = 10
	a.Name = "张三"
	a.Age = 34

	path := files.ProjectDir()
	structInfo, err := ParseStruct(path, a) //Anonymous为true
	fmt.Println(a, structInfo, err)

	b := &B{}
	b.Page.PageNo = 2 //与继承方式一模一样:a.Page.PageNo = 1
	b.Page.PageSize = 20
	b.Name = "张三B"
	b.Age = 23

	structInfob, err := ParseStruct(path, b) //Anonymous为false
	fmt.Println(a, structInfob, err)
}
