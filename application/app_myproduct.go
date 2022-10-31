package application

import (
	"demogogen1/domain_todocore/controller/productapi"
	"demogogen1/domain_todocore/gateway/gatewayproduct"
	"demogogen1/domain_todocore/usecase/getallproduct"
	"demogogen1/domain_todocore/usecase/getoneproduct"
	"demogogen1/domain_todocore/usecase/runproductcreate"
	"demogogen1/domain_todocore/usecase/runproductdelete"
	"demogogen1/domain_todocore/usecase/runproductupdate"
	"demogogen1/shared/gogen"
	"demogogen1/shared/infrastructure/config"
	"demogogen1/shared/infrastructure/logger"
	"demogogen1/shared/infrastructure/server"
	"demogogen1/shared/infrastructure/token"
)

type myproduct struct{}

func NewMyproduct() gogen.Runner {
	return &myproduct{}
}

func (myproduct) Run() error {

	const appName = "myproduct"

	cfg := config.ReadConfig()

	appData := gogen.NewApplicationData(appName)

	//log := logger.NewSimpleJSONLogger(appData)

	log := logger.NewYourLog(appData)

	//log := logger.NewLogrusLog(appData)

	jwtToken := token.NewJWTToken(cfg.JWTSecretKey)

	datasource := gatewayproduct.NewGateway(log, appData, cfg)

	httpHandler := server.NewGinHTTPHandler(log, cfg.Servers[appName].Address, appData)

	x := productapi.NewGinController(log, cfg, jwtToken)
	x.AddUsecase(
		//
		runproductcreate.NewUsecase(datasource),
		getallproduct.NewUsecase(datasource),
		getoneproduct.NewUsecase(datasource),
		runproductupdate.NewUsecase(datasource),
		runproductdelete.NewUsecase(datasource),
	)
	x.RegisterRouter(httpHandler.Router)

	httpHandler.RunWithGracefullyShutdown()

	return nil
}
