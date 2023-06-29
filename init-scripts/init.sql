CREATE TABLE public.Reports (
                                report_id serial PRIMARY KEY,
                                creation_date date default now(),
                                report_info text not null,
                                model_id int
);

CREATE TABLE public.requests_history (
                                         id serial PRIMARY KEY,
                                         date timestamp with time zone default now(),
                                         user_name text,
                                         endpoint text
);


COPY Reports(report_id, creation_date, report_info, model_id) FROM '/docker-entrypoint-initdb.d/data.csv' DELIMITER ',' CSV;

ALTER SEQUENCE public.reports_report_id_seq RESTART WITH 185;