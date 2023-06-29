package models

type SetReportReq struct {
	ReportInfo string `json:"report_info"`
	ModelID    int    `json:"model_id"`
}

type ErrorMessage struct {
	ErrorMsg string `json:"error_msg"`
}

type GetReportResp struct {
	ReportInfo string `json:"report_info"`
	ErrorMessage
}

type MaxObservationTime struct {
	MaxObservationPeriod *int `json:"max_observation_period"`
	ErrorMessage
}

type RequestSnippet struct {
	UserName string `json:"user_name"`
	Endpoint string `json:"endpoint"`
	ErrorMessage
}

type RequestsHistoryResp struct {
	List []struct {
		Id       int    `json:"id"`
		Endpoint string `json:"endpoint"`
		Date     string `json:"date"`
		UserName string `json:"user_name"`
	} `json:"list"`
}
