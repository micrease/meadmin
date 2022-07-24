package user_action_cache_test

import (
	"fmt"
	"giftcard/application/entity/user_action"
	"giftcard/application/service"
	"giftcard/system"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestBuildIndexCache(t *testing.T) {
	s := service.NewUserActionCacheService()
	s.Run()

	ret := service.ActionIgnore(user_action.UserActionReq{PlatformId: 10001, ActionId: 104, RequestId: "test"})
	fmt.Println(ret)
}

func TestMain(m *testing.M) {
	fmt.Println("run main")
	system.ServerInit("../../resources/config-dev.ini")
	m.Run()
}

func setupRouter() *gin.Engine {
	return system.SetupTestRouter()
}
