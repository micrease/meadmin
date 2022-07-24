package yapi

import (
	"encoding/json"
)

// InterfaceService .
type InterfaceService struct {
	client *Client
}

type ReqKVItemSimple struct {
	Name    string `json:"name" structs:"name"`
	Value   string `json:"value" structs:"value"`
	Example string `json:"example" structs:"example"`
	Desc    string `json:"desc" structs:"desc"`
}

type ReqBodyForm struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Required string `json:"required"`
	Example  string `json:"example"`
	Desc     string `json:"desc"`
}

type ReqQuery struct {
	Name     string `json:"name"`
	Required string `json:"required"`
	Example  string `json:"example"`
	Desc     string `json:"desc"`
}

type ReqHeader struct {
	Name  string `json:"name" structs:"name"`
	Value string `json:"value" structs:"value"`
}

type ReqKVItemDetail struct {
	ReqKVItemSimple
	Type     string `json:"type" structs:"type"`
	Required bool   `json:"required" structs:"required"`
}

type interfaceBase struct {
	//ID        int `json:"_id" structs:"_id"`
	//UID       int `json:"uid" structs:"uid"`
	CatID     int `json:"catid" structs:"catid"`
	ProjectID int `json:"project_id" structs:"project_id"`
	//EditUID      int      `json:"edit_uid" structs:"edit_uid"`
	//AddTime      int      `json:"add_time" structs:"add_time"`
	//UpTime       int      `json:"up_time" structs:"up_time"`
	Status       string   `json:"status" structs:"status"`
	Title        string   `json:"title" structs:"title"`
	Path         string   `json:"path" structs:"path"`
	Method       string   `json:"method" structs:"method"`
	ApiOpened    bool     `json:"api_opened"  structs:"api_opened"`
	SwitchNotice bool     `json:"switch_notice" structs:"switch_notice"`
	Markdown     string   `json:"markdown"  structs:"markdown"`
	Desc         string   `json:"desc" structs:"desc"`
	Tag          []string `json:"tag" structs:"tag"`
}

type interfaceReq struct {
	ReqParams           []ReqKVItemSimple `json:"req_params" structs:"req_params"`
	ReqHeaders          []ReqHeader       `json:"req_headers" structs:"req_headers"`
	ReqQuery            []ReqQuery        `json:"req_query" structs:"req_query"`
	ReqBodyForm         []ReqBodyForm     `json:"req_body_form" structs:"req_body_form"`
	ReqBodyIsJsonSchema bool              `json:"req_body_is_json_schema" structs:"req_body_is_json_schema"`
	ReqBodyType         string            `json:"req_body_type" structs:"req_body_type"`
	ReqBodyOther        string            `json:"req_body_other" structs:"req_body_other"`
}

type interfaceRes struct {
	ResBodyIsJsonSchema bool   `json:"res_body_is_json_schema" structs:"res_body_is_json_schema"`
	ResBodyType         string `json:"res_body_type" structs:"res_body_type"`
	ResBody             string `json:"res_body" structs:"res_body"`
}

type InterfaceData struct {
	interfaceBase
	interfaceReq
	interfaceRes
}

type AddOrUpdateInterfaceData struct {
	Token string `json:"token" structs:"token"`
	InterfaceData
}

type ReqBodyOtherPropertiesValue struct {
	Type        string                                 `json:"type"`
	Description string                                 `json:"description"`
	Default     interface{}                            `json:"default"`
	MinLength   int                                    `json:"minLength"` //api就是这样,不是写错了
	MaxLength   int                                    `json:"maxLength"` //api就是这样,不是写错了
	Properties  map[string]ReqBodyOtherPropertiesValue `json:"properties"`
}

type Properties map[string]*Property
type Property struct {
	FieldName string `json:"-"`

	HasAnonymousObjectInGlobal bool `json:"-"` //整个结构树中是否有匿名结构体
	IsAnonymousObject          bool `json:"-"` //当前结点是否匿名结构体

	Type        string      `json:"type"`
	Description string      `json:"description"`
	Default     interface{} `json:"default"`
	Properties  Properties  `json:"properties"`
	Items       *Property   `json:"items"`
	Required    []string    `json:"required"`
}

type ReqBodyOther struct {
	Type       string                                 `json:"type"`
	Title      string                                 `json:"title"`
	Properties map[string]ReqBodyOtherPropertiesValue `json:"properties"`
	Required   []string                               `json:"required"`
}

func GetDefaultReqBodyOther(params map[string]ReqBodyOtherPropertiesValue, requiredFieldList []string) (string, error) {
	req := ReqBodyOther{
		Type:       "object",
		Title:      "empty objec",
		Required:   requiredFieldList,
		Properties: params,
	}
	by, err := json.Marshal(req)
	return string(by), err
}

func GetDefaultResBody(params map[string]ReqBodyOtherPropertiesValue, requiredFieldList []string) *ReqBodyOther {
	req := ReqBodyOther{
		Type:       "object",
		Title:      "empty objec",
		Required:   requiredFieldList,
		Properties: params,
	}
	//by, err := json.Marshal(req)
	//return string(by), err
	return &req
}

func (s *InterfaceData) SetDefaultJsonRequest(method string) {
	s.Method = method
	s.Tag = []string{}

	s.ReqQuery = []ReqQuery{}
	s.ReqParams = []ReqKVItemSimple{}
	s.ReqBodyForm = []ReqBodyForm{}
	s.ReqHeaders = []ReqHeader{}

	header := ReqHeader{}
	header.Name = "Content-Type"
	header.Value = "application/json"
	s.ReqHeaders = append(s.ReqHeaders, header)

	s.ApiOpened = false
	s.SwitchNotice = true
	s.ReqBodyType = "json"
	s.ResBodyType = "json"
	s.ResBodyIsJsonSchema = true
	s.ReqBodyIsJsonSchema = true
	s.Status = "undone"
}

func (s *InterfaceData) SetDefaultFormRequest(method string) {
	s.Method = method
	s.Tag = []string{}

	s.ReqQuery = []ReqQuery{}
	s.ReqParams = []ReqKVItemSimple{}
	s.ReqBodyForm = []ReqBodyForm{}
	s.ReqHeaders = []ReqHeader{}

	header := ReqHeader{}
	header.Name = "Content-Type"
	header.Value = "application/x-www-form-urlencoded"
	s.ReqHeaders = append(s.ReqHeaders, header)

	s.ApiOpened = false
	s.SwitchNotice = true
	s.ReqBodyType = "form"
	s.ResBodyType = "json"
	s.ResBodyIsJsonSchema = true
	s.ReqBodyIsJsonSchema = true
	s.Status = "undone"
}

type Interface struct {
	CommonResp
	Data   InterfaceData `json:"data" structs:"data"`
	string string        `json:"string"`
}

func (i *Interface) ToString() string {
	return i.string
}

type InterfaceParam struct {
	Token string `url:"token"`
	ID    int    `url:"id"`
}

type InterfaceListData struct {
	Count int             `json:"count" structs:"count"`
	Total int             `json:"total" structs:"total"`
	List  []InterfaceData `json:"list" structs:"list"`
}

type InterfaceList struct {
	ErrCode int               `json:"errcode" structs:"errcode"`
	ErrMsg  string            `json:"errmsg" structs:"errmsg"`
	Data    InterfaceListData `json:"data" structs:"data"`
	string  string            `json:"string"`
}

func (i *InterfaceList) ToString() string {
	return i.string
}

type InterfaceListParam struct {
	Token string `url:"token,omitempty"`
	CatID int    `url:"catid,omitempty"`
	Page  int    `url:"Page"`
	Limit int    `url:"limit"`
}

type UploadSwaggerReq struct {
	Type  string `json:"type" structs:"types"`
	Json  string `json:"json" structs:"jsons"`
	Merge string `json:"merge" structs:"string"`
	Token string `json:"token" structs:"token"`
	url   string `json:"url" structs:"url"`
}

func (s *InterfaceService) GetList(opt *InterfaceListParam) (*InterfaceList, error) {
	apiEndpoint := "api/interface/list_cat"
	opt.Token = s.client.Authentication.token
	url, err := addOptions(apiEndpoint, opt)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Get(url, nil)
	if err != nil {
		return nil, err
	}
	result := InterfaceList{}
	err = json.Unmarshal([]byte(resp), &result)
	return &result, err
}

func (s *InterfaceService) Get(id int) (*Interface, error) {
	apiEndpoint := "api/interface/get"
	interfaceParam := InterfaceParam{}
	interfaceParam.ID = id
	interfaceParam.Token = s.client.Authentication.token
	url, err := addOptions(apiEndpoint, &interfaceParam)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Get(url, nil)
	if err != nil {
		return nil, err
	}
	result := Interface{}
	result.string = resp
	err = json.Unmarshal([]byte(resp), &result)
	return &result, err
}

func (s *InterfaceService) AddOrUpdate(data *InterfaceData) (*ModifyResp, error) {
	apiEndpoint := "api/interface/save"
	addOrUpdateInterfaceData := AddOrUpdateInterfaceData{InterfaceData: *data}
	addOrUpdateInterfaceData.Token = s.client.Authentication.token
	resp, err := s.client.Post(apiEndpoint, addOrUpdateInterfaceData)
	if err != nil {
		return nil, err
	}
	result := ModifyResp{}
	result.string = resp
	err = json.Unmarshal([]byte(resp), &result)
	return &result, err
}

func (s *InterfaceService) UploadSwagger(data *string) (*ModifyResp, error) {
	apiEndpoint := "api/open/import_data"
	uploadSwaggerReq := new(UploadSwaggerReq)
	uploadSwaggerReq.Token = s.client.Authentication.token
	uploadSwaggerReq.Type = "swagger"
	uploadSwaggerReq.Merge = "merge"
	uploadSwaggerReq.Json = *data

	resp, err := s.client.Post(apiEndpoint, uploadSwaggerReq)
	if err != nil {
		return nil, err
	}
	result := ModifyResp{}
	err = json.Unmarshal([]byte(resp), &result)
	return &result, err
}
