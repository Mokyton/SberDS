CREATE TABLE public.Reports (
                                report_id serial PRIMARY KEY,
                                creation_date date default now(),
                                report_info text not null,
                                model_id int
);

CREATE TABLE public.requests_history (
                                         id serial PRIMARY KEY,
                                         creation_date timestamp with time zone default now(),
                                         type text,
                                         req_info jsonb
);

CREATE OR REPLACE FUNCTION public.max_observation_time(_model_id int) RETURNS integer
    LANGUAGE plpgsql
AS $$
declare _res int;
begin
    with max_interval as (
        select a.creation_date - b.creation_date as dif from (select creation_date from reports where model_id = _model_id)
                                                                 a cross join (select creation_date from reports where model_id = _model_id) b
    )
    select max(abs(dif)) from max_interval into _res;
    return _res;
end;
$$;