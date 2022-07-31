package validate

import (
	"admease/library/astparser"
	"admease/library/context/api"
	"admease/library/logger"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"reflect"
)

func BindWithPanic(ctx *api.Context, obj interface{}) error {
	return bind(ctx, obj, true)
}

func Bind(ctx *api.Context, obj interface{}) error {
	return bind(ctx, obj, false)
}

func bind(ctx *api.Context, obj interface{}, withPanic bool) error {
	//绑定参数
	err := ctx.GinCtx.ShouldBind(obj)
	if err != nil {
		//解析错误信息
		value := reflect.TypeOf(obj)
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range errs {
				if f, exist := value.Elem().FieldByName(e.Field()); exist {
					if value, ok := f.Tag.Lookup("tips"); ok {
						err = fmt.Errorf("%s", value)
						break
					} else {
						err = fmt.Errorf("%s", e.Value())
						break
					}
				} else {
					err = fmt.Errorf("%s参数错误", e.Field())
					break
				}
			}
		}

		if withPanic {
			//这里的panic将会被中间件捕获
			panic(errors.WithMessage(err, "参数验证失败"))
		}
		return err
	}

	bindTraceId(ctx, obj)
	if ctx.ApiDoc.Enable {
		ctx.ApiDoc.ParamAst, err = astparser.ParseStruct(ctx.ProjectPath, obj)
		paramJson, err := json.MarshalIndent(obj, "", "\t")
		if err != nil {
			ctx.Log.Error("参数错误", zap.Any("json error", obj))
		}
		ctx.ApiDoc.ParamExample = string(paramJson)
	}
	return nil
}

func bindTraceId(ctx *api.Context, obj interface{}) {
	traceIdField := reflect.ValueOf(obj).Elem().FieldByName(logger.TraceId)

	if !traceIdField.IsValid() {
		return
	}

	traceId := traceIdField.String()
	if len(traceId) == 0 {
		//取URL上的platform_id
		traceId = ctx.TraceId
	}

	if len(traceId) > 0 {
		traceIdField.SetString(traceId)
	}

	//logger traceId
	loggerField := reflect.ValueOf(obj).Elem().FieldByName("Logger")
	if !loggerField.IsValid() {
		return
	}

	logTraceIdField := loggerField.FieldByName(logger.TraceId)
	if !logTraceIdField.IsValid() {
		return
	}
	logTraceIdField.SetString(traceId)
}

func RegisterExtension() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("datetime", datetime)
		if err != nil {
			panic("ValidatorInstall Error" + err.Error())
		}

		err = v.RegisterValidation("date", date)
		if err != nil {
			panic("ValidatorInstall Error" + err.Error())
		}

		err = v.RegisterValidation("time", time)
		if err != nil {
			panic("ValidatorInstall Error" + err.Error())
		}
	}
}
