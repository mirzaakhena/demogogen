package restapi

import (
	"demogogen1/shared/gogen"
	"demogogen1/shared/infrastructure/config"
	"demogogen1/shared/infrastructure/logger"
	"demogogen1/shared/infrastructure/token"

	"github.com/gin-gonic/gin"
)

type selectedRouter = gin.IRouter

type ginController struct {
	*gogen.BaseController
	log      logger.Logger
	cfg      *config.Config
	jwtToken token.JWTToken
}

func NewGinController(log logger.Logger, cfg *config.Config, tk token.JWTToken) gogen.RegisterRouterHandler[selectedRouter] {
	return &ginController{
		BaseController: gogen.NewBaseController(),
		log:            log,
		cfg:            cfg,
		jwtToken:       tk,
	}
}

func (r *ginController) RegisterRouter(router selectedRouter) {

	resource := router.Group("/api/v1", r.authentication())
	resource.GET("/todo", r.authorization(), r.getAllTodoHandler())
	resource.POST("/todo/", r.authorization(), r.runTodoCheckHandler())
	resource.POST("/runtodocreate", r.authorization(), r.runTodoCreateHandler())

}
