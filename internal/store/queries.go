package store

const (
	getReportByIDSQL      = `select report_info from public.reports where report_id = $1`
	addReportSQL          = `insert into public.reports(report_info, model_id) VALUES ($1, $2)`
	maxObservationTimeSQL = `with max_observation_time as (select r.creation_date - lag(r.creation_date) over (ORDER BY r.creation_date) as dif
                         from reports r
                         where model_id = $1),
    from_current_date as (
        select current_date - r.creation_date from reports r where model_id = $1
        )
select json_build_object(
    'max_observation_period', case
    when count(*) = 1 then (select * from from_current_date)
    else max(abs(dif)) end
           )
from max_observation_time`
	addSnippetRequestHistorySQL = `insert into public.requests_history (user_name, endpoint)
select user_name, endpoint
from json_to_record($1) as x( user_id int, user_name text, endpoint text)`
	getRequestsHistorySQL = ` select json_build_object(
                   'list', (select json_agg(row_to_json(public.requests_history.*)) from public.requests_history offset $1 limit $2 )
               );`
)
