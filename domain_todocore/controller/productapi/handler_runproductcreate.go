package productapi

import (
	"context"
	"demogogen1/domain_todocore/usecase/runproductcreate"
	"demogogen1/shared/gogen"
	"demogogen1/shared/infrastructure/logger"
	"demogogen1/shared/model/payload"
	"demogogen1/shared/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (r *ginController) runProductCreateHandler() gin.HandlerFunc {

	type InportRequest = runproductcreate.InportRequest
	type InportResponse = runproductcreate.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
		InportRequest
	}

	type response struct {
		*InportResponse
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		if err := c.BindJSON(&jsonReq); err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req InportRequest
		req = jsonReq.InportRequest
		req.Now = time.Now()
		req.RandomString = util.GenerateID(16)

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.InportResponse = res

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
