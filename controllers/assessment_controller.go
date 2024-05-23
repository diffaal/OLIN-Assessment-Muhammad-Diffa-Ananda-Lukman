package controllers

import (
	"fmt"
	"net/http"
	"olin-assessment-muhammad-diffa/package/enum"
	"olin-assessment-muhammad-diffa/package/logger"
	"olin-assessment-muhammad-diffa/package/response"
	"olin-assessment-muhammad-diffa/schemas"
	"olin-assessment-muhammad-diffa/services"

	"github.com/gin-gonic/gin"
)

type AssessmentController struct {
	assessmentService services.IAssessmentService
}

func NewAssessmentController(assessmentService services.IAssessmentService) *AssessmentController {
	assessmentController := AssessmentController{
		assessmentService: assessmentService,
	}
	return &assessmentController
}

func (ac *AssessmentController) AssessmentCodeOne(c *gin.Context) {
	var cor schemas.CodeOneRequest

	err := c.BindJSON(&cor)
	if err != nil {
		logger.Error(fmt.Sprintf("[FAILED][AssessmentCodeOne] %v: %+v", enum.INVALID_JSON, err))
		response.BaseResponseWriter(
			c,
			http.StatusBadRequest,
			nil,
			enum.BAD_REQUEST,
			enum.INVALID_JSON,
		)
		return
	}

	err = cor.ValidatePayload()
	if err != nil {
		logger.Error(fmt.Sprintf("[FAILED][AssessmentCodeOne] %v: %+v", enum.INVALID_REQUEST, err))
		response.BaseResponseWriter(
			c,
			http.StatusUnprocessableEntity,
			nil,
			enum.INVALID_REQUEST,
			err.Error(),
		)
		return
	}

	result := ac.assessmentService.TwoSumArrayTarget(cor.Nums, cor.Target)

	logger.Info(fmt.Sprintf("[SUCCESS][AssessmentCodeOne] Result: %v", result))
	response.BaseResponseWriter(
		c,
		http.StatusOK,
		result,
		enum.SUCCESS,
		enum.SUCCESS_DETAIL,
	)
}

func (ac *AssessmentController) AssessmentCodeTwo(c *gin.Context) {
	var ctr schemas.CodeTwoRequest

	err := c.BindJSON(&ctr)
	if err != nil {
		logger.Error(fmt.Sprintf("[FAILED][AssessmentCodeTwo] %v: %+v", enum.INVALID_JSON, err))
		response.BaseResponseWriter(
			c,
			http.StatusBadRequest,
			nil,
			enum.BAD_REQUEST,
			enum.INVALID_JSON,
		)
		return
	}

	err = ctr.ValidatePayload()
	if err != nil {
		logger.Error(fmt.Sprintf("[FAILED][AssessmentCodeTwo] %v: %+v", enum.INVALID_REQUEST, err))
		response.BaseResponseWriter(
			c,
			http.StatusUnprocessableEntity,
			nil,
			enum.INVALID_REQUEST,
			err.Error(),
		)
		return
	}

	result := ac.assessmentService.ThreeSumArrayZero(ctr.Nums)

	logger.Info(fmt.Sprintf("[SUCCESS][AssessmentCodeTwo] Result: %v", result))
	response.BaseResponseWriter(
		c,
		http.StatusOK,
		result,
		enum.SUCCESS,
		enum.SUCCESS_DETAIL,
	)
}
