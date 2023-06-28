package controllers

import (
	"SberDS/internal/models"
	"SberDS/internal/store"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ReportController struct {
	store *store.Repository
}

func NewReportController(store *store.Repository) *ReportController {
	return &ReportController{store: store}
}

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

func (c *ReportController) SetReport(ctx *gin.Context) {
	var dto models.SetReportReq
	err := ctx.BindJSON(&dto)

	if err != nil {
		dto.ErrorMsg = err.Error()
		ctx.JSON(http.StatusBadRequest, dto.ErrorMessage)
		return
	}

	err = c.store.SetReport(ctx, dto)
	if err != nil {
		dto.ErrorMsg = err.Error()
		ctx.JSON(http.StatusInternalServerError, dto.ErrorMessage)
		return
	}

	ctx.JSON(http.StatusOK, dto.ErrorMessage)
}

func (c *ReportController) GetObservationTime(ctx *gin.Context) {

}
