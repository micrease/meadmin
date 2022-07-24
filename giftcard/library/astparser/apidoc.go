package astparser

type ApiDoc struct {
	Enable        bool //是否开启
	ApiName       string
	Method        string
	TraceId       string
	ParamAst      *StructInfo
	ResultAst     *StructInfo
	FuncAst       *FuncInfo
	ParamExample  string
	ResultExample string
}
