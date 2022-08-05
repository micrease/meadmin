package apidoc

import (
	"encoding/json"
	"errors"
	"fmt"
	yapi "meadmin/library/apidoc/yapiclient"
	"meadmin/library/astparser"
	"strings"
)

func NewYapi() (*yapi.Client, error) {
	yapiUrl := "https://yapi.kgtindok.com"
	yapiToken := "3a0ccf8ec49f404bff302b622069e1d42f11b2bc49c77e5a90ea7b38c55b17b1"
	yapiClient, err := yapi.NewClient(yapiUrl, yapiToken)
	yapiClient.CateId = 3761 //3706
	if err != nil {
		panic(err)
	}
	return yapiClient, err
}

func SaveApi(doc *astparser.ApiDoc) error {
	yapiClient, err := NewYapi()
	if err != nil {
		return err
	}

	project, err := yapiClient.Project.Get()
	if err != nil {
		return err
	}

	if yapiClient == nil {
		return errors.New("doc is nil")
	}

	if doc.Method == "POST" || doc.Method == "PUT" {
		PostJsonApi(doc.Method, project.Data.ID, yapiClient, doc)
	} else if doc.Method == "GET" || doc.Method == "DELETE" {
		GetFormApi(doc.Method, project.Data.ID, yapiClient, doc)
	}
	return nil
}

func PostJsonApi(method string, projectId int, yapiClient *yapi.Client, apiDoc *astparser.ApiDoc) error {
	data := new(yapi.InterfaceData)
	data.SetDefaultJsonRequest(method)
	data.CatID = yapiClient.CateId
	if apiDoc.FuncAst.Commit != nil {
		data.Title = apiDoc.FuncAst.Commit.ToString()
	} else {
		data.Title = apiDoc.ApiName
	}

	if apiDoc.FuncAst.ClassCommit != nil {
		classDoc := apiDoc.FuncAst.ClassCommit.ToString()
		data.Title = "[" + classDoc + "]" + data.Title
	}

	data.Path = apiDoc.ApiName
	data.ProjectID = projectId

	//request
	reqProperty := &yapi.Property{}
	reqProperty.Type = "object"
	reqCommitArr := []string{}

	reqBody := createJsonProperty(reqProperty, apiDoc.ParamAst.ChildFields, &reqCommitArr)
	reqBody = mergeAnonymous(reqBody)

	reqJsonBody, _ := json.Marshal(reqBody)
	data.ReqBodyOther = string(reqJsonBody)
	reqCommit := strings.Join(reqCommitArr, "")

	//response
	respProperty := &yapi.Property{}
	respProperty.Type = "object"
	commitArr := []string{}

	resBody := createJsonProperty(respProperty, apiDoc.ResultAst.ChildFields, &commitArr)
	resBody = mergeAnonymous(resBody)

	resJsonBody, _ := json.Marshal(resBody)
	data.ResBody = string(resJsonBody)
	respCommit := strings.Join(commitArr, "")

	//commit
	data.Desc = FormatCommitDesc(reqCommit, respCommit, apiDoc)

	//save
	resp, err := yapiClient.Interface.AddOrUpdate(data)
	fmt.Println(resp, err)
	return nil
}

func GetFormApi(method string, projectId int, yapiClient *yapi.Client, apiDoc *astparser.ApiDoc) error {
	data := new(yapi.InterfaceData)
	data.SetDefaultFormRequest(method)
	data.CatID = yapiClient.CateId

	if apiDoc.FuncAst == nil || apiDoc.FuncAst.Commit == nil {
		return nil
	}

	data.Title = apiDoc.FuncAst.Commit.ToString()
	if apiDoc.FuncAst.ClassCommit != nil {
		classDoc := apiDoc.FuncAst.ClassCommit.ToString()
		data.Title = "[" + classDoc + "]" + data.Title
	}

	data.Path = apiDoc.ApiName
	data.ProjectID = projectId

	reqBodyForm, paramsCommits := CreateFormParam(apiDoc.ParamAst.ChildFields)
	data.ReqQuery = reqBodyForm

	respProperty := &yapi.Property{}
	respProperty.Type = "object"
	commitArr := []string{}
	resBody := createJsonProperty(respProperty, apiDoc.ResultAst.ChildFields, &commitArr)
	resBody = mergeAnonymous(resBody)

	resJsonBody, _ := json.Marshal(resBody)
	data.ResBody = string(resJsonBody)
	respCommit := strings.Join(commitArr, "")

	data.Desc = FormatCommitDesc(paramsCommits, respCommit, apiDoc)
	resp, err := yapiClient.Interface.AddOrUpdate(data)
	fmt.Println(resp, err, paramsCommits)
	return nil
}

//合并匿名类,就是继承的意思
func mergeAnonymous(property *yapi.Property) *yapi.Property {
	if property.HasAnonymousObjectInGlobal {
		for childKey, child := range property.Properties {
			if child.IsAnonymousObject {
				for k, v := range child.Properties {
					property.Properties[k] = v
				}
				property.Required = append(property.Required, child.Required...)
				delete(property.Properties, childKey)
			}
			if child.HasAnonymousObjectInGlobal {
				mergeAnonymous(child)
			}
		}
	}
	return property
}

//创建属性
func createJsonProperty(property *yapi.Property, structFieldList []*astparser.StructField, commitArr *[]string) *yapi.Property {
	property.Properties = map[string]*yapi.Property{} //字段名->描述信息
	property.Required = []string{}                    //当前对象的必选参数
	//property.Type = "object"

	for _, structField := range structFieldList {
		if structField.FiledName == "ReqHeader" {
			continue
		}

		//字段名称
		filedName := structField.Tag
		if len(filedName) == 0 {
			filedName = structField.FiledName
		}

		if len(filedName) == 0 {
			continue
		}
		typeVal, defaultVal := GetKind(structField.Kind)
		//如果是内联结构体
		property.Required = append(property.Required, filedName)
		fieldProperty := &yapi.Property{}
		fieldProperty.Type = typeVal
		fieldProperty.Default = defaultVal
		fieldProperty.IsAnonymousObject = structField.Anonymous

		if len(property.FieldName) > 0 && !fieldProperty.IsAnonymousObject {
			fieldProperty.FieldName = property.FieldName + "." + filedName
		} else {
			fieldProperty.FieldName = filedName
		}

		commit := ""
		if structField.Commits != nil {
			commit = structField.Commits.ToString()
			fieldProperty.Description = commit
		}

		//值是一个结构体时
		if typeVal == "object" {
			if structField.Type == "DBTime" {
				fieldProperty.Type = "string"
				fieldProperty.Default = ""
			} else {
				if structField.T != nil {
					fieldProperty = createJsonProperty(fieldProperty, structField.T.ChildFields, commitArr)
				}
			}
		}

		//值是一个列表时
		if typeVal == "array" {
			itemObj := &yapi.Property{Type: "object"}
			itemObj.FieldName = filedName
			if len(property.FieldName) > 0 {
				itemObj.FieldName = property.FieldName + "." + filedName
			}
			//这是一个范性列表,从 structField.T 中解析
			if structField.T != nil {
				fieldProperty.Items = createJsonProperty(itemObj, structField.T.ChildFields, commitArr)
			} else if structField.ChildFields != nil {
				fieldProperty.Items = createJsonProperty(itemObj, structField.ChildFields, commitArr)
			} else {
				//如果没有值,忽略
				continue
			}
		}

		if len(commit) > 0 {
			commitHtml := "<p><strong>" + fieldProperty.FieldName + "</strong>[" + typeVal + "]:" + commit + "</p>"
			*commitArr = append(*commitArr, commitHtml)
		}

		if fieldProperty.IsAnonymousObject || fieldProperty.HasAnonymousObjectInGlobal {
			property.HasAnonymousObjectInGlobal = true
		}

		property.Properties[filedName] = fieldProperty
	}
	return property
}

func FormatCommitDesc(paramDesc, resDesc string, apiDoc *astparser.ApiDoc) string {
	desc := ""
	desc += "<h5>请求字段概要:</h5>详情见后文:请求参数"
	desc += paramDesc
	desc += "<h5>返回字段概要:</h5>详情见后文:返回数据"
	desc += resDesc
	desc += "<h5>传参示例:</h5>"
	desc += "<pre><code>" + apiDoc.ParamExample + "</code></pre>"
	desc += "<h5>返回示例:</h5>"
	desc += "<pre><code>" + apiDoc.ResultExample + "</code></pre>"
	desc += "<p>*此文档由系统自动生成,如有错误请联系开发者</p>"
	return desc
}

type JsonParams struct {
	ReqBodyOther *yapi.ReqBodyOther
	Desc         []string
}

func astFieldToReqQuery(childField *astparser.StructField, reqQueryList *[]yapi.ReqQuery, paramsCommitArr *[]string) {
	if childField.FiledName == "ReqHeader" {
		return
	}

	filed := childField.Tag
	if len(filed) == 0 {
		filed = childField.FiledName
	}

	if len(filed) == 0 {
		return
	}

	typeVal, defaultVal := GetKind(childField.Kind)
	//匿名对象要递归调用
	if childField.Anonymous && typeVal == "object" {
		if childField.T != nil && childField.T.ChildFields != nil {
			for _, tfield := range childField.T.ChildFields {
				astFieldToReqQuery(tfield, reqQueryList, paramsCommitArr)
			}
		}
	} else {
		desc := childField.Commits.ToString()
		if len(desc) > 0 {
			p := "<p><strong>" + filed + "</strong>[" + typeVal + "]:" + desc + "</p>"
			*paramsCommitArr = append(*paramsCommitArr, p)
		}
		param := yapi.ReqQuery{
			Name:     filed,
			Required: "1",
			Example:  defaultVal,
			Desc:     desc,
		}
		*reqQueryList = append(*reqQueryList, param)
	}
}

//CreateFormParam
func CreateFormParam(childFields []*astparser.StructField) ([]yapi.ReqQuery, string) {
	params := []yapi.ReqQuery{}
	paramsCommitArr := []string{}

	//params
	for _, childField := range childFields {
		astFieldToReqQuery(childField, &params, &paramsCommitArr)
	}

	paramDesc := strings.Join(paramsCommitArr, "")
	return params, paramDesc
}

func GetKind(docType string) (string, string) {
	if IsInt(docType) {
		return "integer", "0"
	} else if IsFloat(docType) {
		return "number", "0.00"
	} else if IsString(docType) {
		return "string", "\"\""
	} else if docType == "bool" {
		return "boolean", "false"
	} else if docType == "slice" {
		return "array", "[]"
	} else if docType == "struct" || docType == "T" { //T是范性
		return "object", "null"
	} else {
		return "object", "null"
	}
}

func IsInt(t string) bool {
	if t == "int" || t == "int8" || t == "int16" || t == "int32" || t == "int64" || t == "uintptr" {
		return true
	}
	if t == "uint" || t == "uint8" || t == "uint16" || t == "uint32" || t == "uint64" {
		return true
	}
	return false
}

func IsFloat(t string) bool {
	if t == "float32" || t == "float64" {
		return true
	}
	return false
}

func IsString(t string) bool {
	if t == "string" {
		return true
	}
	return false
}
