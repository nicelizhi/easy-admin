package router

import (
	"os"

	common "github.com/nicelizhi/easy-admin/common/middleware"

	log "github.com/nicelizhi/easy-admin-core/logger"
	"github.com/nicelizhi/easy-admin-core/sdk"

	"github.com/gin-gonic/gin"
)

// InitRouter 路由初始化，不要怀疑，这里用到了
func InitRouter() {
	var r *gin.Engine
	h := sdk.Runtime.GetEngine()
	if h == nil {
		log.Fatal("not found engine...")
		os.Exit(-1)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		os.Exit(-1)
	}

	authMiddleware, err := common.AuthInit()
	if err != nil {
		log.Fatalf("JWT Init Error, %s", err.Error())
	}

	initRouter(r, authMiddleware)
}
