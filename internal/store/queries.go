package store

const (
	getReportByIDSQL      = "select report_info from public.reports where report_id = $1"
	addReportSQL          = "insert into public.reports(report_info, model_id) VALUES ($1, $2)"
	maxObservationTimeSQL = "with observation_time as (\n    select creation_date - lag(creation_date) over (ORDER BY creation_date) as dif from reports where model_id = 1089\n)  select max(dif) from observation_time;"
)
