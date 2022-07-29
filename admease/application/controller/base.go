package controller

type BaseController struct {
}

//panic不建议在代码里大量使用,只限在controller层使用。因此写在这里
func (ctrl *BaseController) PanicIf(err error, message string) {
	if err != nil {
		panic(message)
	}
}
