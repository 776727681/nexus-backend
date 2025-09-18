package controller

import (
	"github.com/gin-gonic/gin"
	"nexus-backend/api/common/statecode"
	"nexus-backend/api/model/request"
	"nexus-backend/api/model/response"
	"nexus-backend/api/service"
	"nexus-backend/api/validate"
	"nexus-backend/log"
)

type MultiSignPoolController struct {
}

// 处理设置多签信息的请求。验证请求参数，调用服务层设置多签信息
func (c *MultiSignPoolController) SetMultiSign(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.SetMultiSign{}
	log.Logger.Sugar().Info("SetMultiSign req ", req)

	errCode := validate.NewMutiSign().SetMultiSign(ctx, &req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil)
		return
	}

	errCode, err := services.NewMutiSign().SetMultiSign(&req)
	if errCode != statecode.CommonSuccess {
		log.Logger.Error(err.Error())
		res.Response(ctx, errCode, nil)
		return
	}

	res.Response(ctx, statecode.CommonSuccess, nil)
	return
}

// 处理获取多签信息的请求。验证参数（链ID），调用服务层获取多签信息，返回结果
func (c *MultiSignPoolController) GetMultiSign(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.GetMultiSign{}
	result := response.MultiSign{}
	log.Logger.Sugar().Info("GetMultiSign req ", nil)

	errCode := validate.NewMutiSign().GetMultiSign(ctx, &req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil)
		return
	}

	errCode, err := services.NewMutiSign().GetMultiSign(&result, req.ChainId)
	if errCode != statecode.CommonSuccess {
		log.Logger.Error(err.Error())
		res.Response(ctx, errCode, nil)
		return
	}

	res.Response(ctx, statecode.CommonSuccess, result)
	return
}
