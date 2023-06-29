package controllers

import (
	"SberDS/internal/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ReportController struct {
	*BaseController
}

func NewReportController(base *BaseController) *ReportController {
	return &ReportController{base}
}

// GetReport
// @Tags Reports
// @Summary Получить информацию о отчете по report_id
// @Description Получить информацию о отчете по report_id
// @Param report_id path int true "report_id"
// @Produce json
// @Success 200 {object} models.GetReportResp
// @Failure 400 {object} models.ErrorMessage
// @Failure 500 {object} models.ErrorMessage
// @Router /api/v1/get_report/{report_id} [get]
func (c *ReportController) GetReport(ctx *gin.Context) {
	var dto models.GetReportResp

	reportId := ctx.Param("report_id")
	if reportId == "" {
		dto.ErrorMsg = fmt.Sprintf("invalid report_id: %s", reportId)
		ctx.JSON(http.StatusBadRequest, dto)
		return
	}

	reportIdInt, err := strconv.Atoi(reportId)
	if err != nil {
		dto.ErrorMsg = err.Error()
		ctx.JSON(http.StatusInternalServerError, dto)
		return
	}

	data, err := c.store.GetReportById(ctx, reportIdInt)
	if err != nil {
		dto.ErrorMsg = err.Error()
		ctx.JSON(http.StatusInternalServerError, dto)
		return
	}

	dto.ReportInfo = data

	ctx.JSON(http.StatusOK, dto)
}

// SetReport
// @Tags Reports
// @Summary Отправить отчет
// @Description Отправить отчет
// @Accept json
// @Produce json
// @Param report body models.SetReportReq true "Отчет"
// @Success 200 {object} models.ErrorMessage
// @Failure 400 {object} models.ErrorMessage
// @Failure 500 {object} models.ErrorMessage
// @Router /api/v1/set_report [post]
func (c *ReportController) SetReport(ctx *gin.Context) {
	var dto models.SetReportReq
	var dtoErr models.ErrorMessage
	err := ctx.BindJSON(&dto)

	if err != nil {
		dtoErr.ErrorMsg = err.Error()
		ctx.JSON(http.StatusBadRequest, dtoErr)
		return
	}

	err = c.store.SetReport(ctx, dto)
	if err != nil {
		dtoErr.ErrorMsg = err.Error()
		ctx.JSON(http.StatusInternalServerError, dtoErr)
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

// GetMaxObservationTime
// @Tags Reports
// @Summary Получить максимальное промежуток времени в течении которых не было поступлений отчетов по модели
// @Description Получить максимальное промежуток времени в течении которых не было поступлений отчетов по модели
// @Accept json
// @Produce json
// @Param model_id path int true "model_id"
// @Success 200 {object} models.ErrorMessage
// @Failure 400 {object} models.ErrorMessage
// @Failure 500 {object} models.ErrorMessage
// @Router /api/v1/get_observation_time/{model_id} [get]
func (c *ReportController) GetMaxObservationTime(ctx *gin.Context) {
	var dto models.MaxObservationTime

	modelId := ctx.Param("model_id")
	if modelId == "" {
		dto.ErrorMsg = fmt.Sprintf("invalid model_id: %s", modelId)
		ctx.JSON(http.StatusBadRequest, dto)
		return
	}

	modelIdInt, err := strconv.Atoi(modelId)
	if err != nil {
		dto.ErrorMsg = fmt.Sprintf("invalid model_id: %v", err)
		ctx.JSON(http.StatusBadRequest, dto)
		return
	}

	data, err := c.store.GetMaxObservationTime(ctx, modelIdInt)
	if err != nil {
		dto.ErrorMsg = err.Error()
		ctx.JSON(http.StatusInternalServerError, dto)
		return
	}

	err = json.Unmarshal(data, &dto)
	if err != nil {
		dto.ErrorMsg = err.Error()
		ctx.JSON(http.StatusInternalServerError, dto)
		return
	}

	if dto.MaxObservationPeriod == nil {
		dto.ErrorMsg = fmt.Sprintf("models id %s does not exists", modelId)
		ctx.JSON(http.StatusInternalServerError, dto)
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

// GetRequestsHistory
// @Tags Reports
// @Summary Получить историю запросов
// @Description Получить историю запросов
// @Produce json
// @Param offset path int false "offset"
// @Param limit path int false "limit"
// @Success 200 {object} models.RequestsHistoryResp
// @Failure 500 {object} models.ErrorMessage
// @Router /api/v1/requests_history [get]
func (c *ReportController) GetRequestsHistory(ctx *gin.Context) {
	var errMsg models.ErrorMessage
	offset, err := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	if err != nil {
		errMsg.ErrorMsg = err.Error()
		ctx.JSON(http.StatusInternalServerError, errMsg)
		return
	}
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil {
		errMsg.ErrorMsg = err.Error()
		ctx.JSON(http.StatusInternalServerError, errMsg)
		return
	}

	data, err := c.store.RequestsHistory(ctx, offset, limit)
	if err != nil {
		errMsg.ErrorMsg = err.Error()
		ctx.JSON(http.StatusInternalServerError, errMsg)
		return
	}
	ctx.JSON(http.StatusOK, data)
}
