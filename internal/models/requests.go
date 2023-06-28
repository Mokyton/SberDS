package models

type SetReportReq struct {
	ReportInfo string `json:"report_info"`
	ModelID    int    `json:"model_id"`
	ErrorMessage
}

type ErrorMessage struct {
	ErrorMsg string `json:"error_msg"`
}
