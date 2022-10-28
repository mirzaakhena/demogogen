package application

import (
	"demogogen1/domain_todocore/controller/restapi"
	"demogogen1/domain_todocore/gateway/withmongodb"
	"demogogen1/domain_todocore/usecase/getalltodo"
	"demogogen1/domain_todocore/usecase/runtodocheck"
	"demogogen1/domain_todocore/usecase/runtodocreate"
	"demogogen1/shared/gogen"
	"demogogen1/shared/infrastructure/config"
	"demogogen1/shared/infrastructure/logger"
	"demogogen1/shared/infrastructure/server"
	"demogogen1/shared/infrastructure/token"
)

type mytodo struct{}

func NewMytodo() gogen.Runner {
	return &mytodo{}
}

func (mytodo) Run() error {

	const appName = "mytodo"

	cfg := config.ReadConfig()

	appData := gogen.NewApplicationData(appName)

	log := logger.NewSimpleJSONLogger(appData)

	jwtToken := token.NewJWTToken(cfg.JWTSecretKey)

	datasource := withmongodb.NewGateway(log, appData, cfg)

	httpHandler := server.NewGinHTTPHandler(log, cfg.Servers[appName].Address, appData)

	x := restapi.NewGinController(log, cfg, jwtToken)
	x.AddUsecase(
		//
		getalltodo.NewUsecase(datasource),
		runtodocheck.NewUsecase(datasource),
		runtodocreate.NewUsecase(datasource),
	)
	x.RegisterRouter(httpHandler.Router)

	httpHandler.RunWithGracefullyShutdown()

	return nil
}
