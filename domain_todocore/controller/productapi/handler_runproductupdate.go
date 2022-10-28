package productapi

import (
	"context"
	"demogogen1/domain_todocore/model/vo"
	"demogogen1/domain_todocore/usecase/runproductupdate"
	"demogogen1/shared/gogen"
	"demogogen1/shared/infrastructure/logger"
	"demogogen1/shared/model/payload"
	"demogogen1/shared/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *ginController) runProductUpdateHandler() gin.HandlerFunc {

	type InportRequest = runproductupdate.InportRequest
	type InportResponse = runproductupdate.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
		InportRequest
	}

	type response struct {
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
		req.ProductID = vo.ProductID(c.Param("product_id"))

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		_ = res

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
