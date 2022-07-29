package woker

import (
	"fmt"
	"giftcard/library/context/api"
	"math/rand"
)

//定时任务或者kafka队列中
func NewContext(traceIdArg ...string) *api.Context {
	ctx := new(api.Context)
	var traceId string
	if len(traceIdArg) > 0 {
		traceId = traceIdArg[0]
	} else {
		traceId = fmt.Sprintf("R%d", rand.Int63())
	}

	ctx.TraceId = traceId
	ctx.Log.TraceId = ctx.TraceId
	return ctx
}
