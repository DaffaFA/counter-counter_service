--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4
-- Dumped by pg_dump version 16.1 (Ubuntu 16.1-1.pgdg22.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: timescaledb; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS timescaledb WITH SCHEMA public;


--
-- Name: EXTENSION timescaledb; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION timescaledb IS 'Enables scalable inserts and complex queries for time-series data (Community Edition)';


--
-- Name: user_access_control; Type: SCHEMA; Schema: -; Owner: COUNTER@2024
--

CREATE SCHEMA user_access_control;


ALTER SCHEMA user_access_control OWNER TO "COUNTER@2024";

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: COUNTER@2024
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO "COUNTER@2024";

--
-- Name: departments; Type: TABLE; Schema: user_access_control; Owner: COUNTER@2024
--

CREATE TABLE user_access_control.departments (
    id integer NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE user_access_control.departments OWNER TO "COUNTER@2024";

--
-- Name: departments_id_seq; Type: SEQUENCE; Schema: user_access_control; Owner: COUNTER@2024
--

CREATE SEQUENCE user_access_control.departments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE user_access_control.departments_id_seq OWNER TO "COUNTER@2024";

--
-- Name: departments_id_seq; Type: SEQUENCE OWNED BY; Schema: user_access_control; Owner: COUNTER@2024
--

ALTER SEQUENCE user_access_control.departments_id_seq OWNED BY user_access_control.departments.id;


--
-- Name: permissions; Type: TABLE; Schema: user_access_control; Owner: COUNTER@2024
--

CREATE TABLE user_access_control.permissions (
    id integer NOT NULL,
    name character varying NOT NULL,
    alias character varying NOT NULL,
    parent_id integer
);


ALTER TABLE user_access_control.permissions OWNER TO "COUNTER@2024";

--
-- Name: permissions_id_seq; Type: SEQUENCE; Schema: user_access_control; Owner: COUNTER@2024
--

CREATE SEQUENCE user_access_control.permissions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE user_access_control.permissions_id_seq OWNER TO "COUNTER@2024";

--
-- Name: permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: user_access_control; Owner: COUNTER@2024
--

ALTER SEQUENCE user_access_control.permissions_id_seq OWNED BY user_access_control.permissions.id;


--
-- Name: user_department_permissions; Type: TABLE; Schema: user_access_control; Owner: COUNTER@2024
--

CREATE TABLE user_access_control.user_department_permissions (
    id integer NOT NULL,
    user_id bigint,
    department_id integer,
    permission_id integer NOT NULL,
    read boolean NOT NULL,
    write boolean NOT NULL
);


ALTER TABLE user_access_control.user_department_permissions OWNER TO "COUNTER@2024";

--
-- Name: user_department_permissions_id_seq; Type: SEQUENCE; Schema: user_access_control; Owner: COUNTER@2024
--

CREATE SEQUENCE user_access_control.user_department_permissions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE user_access_control.user_department_permissions_id_seq OWNER TO "COUNTER@2024";

--
-- Name: user_department_permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: user_access_control; Owner: COUNTER@2024
--

ALTER SEQUENCE user_access_control.user_department_permissions_id_seq OWNED BY user_access_control.user_department_permissions.id;


--
-- Name: users; Type: TABLE; Schema: user_access_control; Owner: COUNTER@2024
--

CREATE TABLE user_access_control.users (
    id bigint NOT NULL,
    department_id integer NOT NULL,
    full_name character varying,
    username character varying NOT NULL,
    password character varying NOT NULL,
    expired_at timestamp with time zone,
    activated_at timestamp with time zone DEFAULT now() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE user_access_control.users OWNER TO "COUNTER@2024";

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: user_access_control; Owner: COUNTER@2024
--

CREATE SEQUENCE user_access_control.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE user_access_control.users_id_seq OWNER TO "COUNTER@2024";

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: user_access_control; Owner: COUNTER@2024
--

ALTER SEQUENCE user_access_control.users_id_seq OWNED BY user_access_control.users.id;


--
-- Name: departments id; Type: DEFAULT; Schema: user_access_control; Owner: COUNTER@2024
--

ALTER TABLE ONLY user_access_control.departments ALTER COLUMN id SET DEFAULT nextval('user_access_control.departments_id_seq'::regclass);


--
-- Name: permissions id; Type: DEFAULT; Schema: user_access_control; Owner: COUNTER@2024
--

ALTER TABLE ONLY user_access_control.permissions ALTER COLUMN id SET DEFAULT nextval('user_access_control.permissions_id_seq'::regclass);


--
-- Name: user_department_permissions id; Type: DEFAULT; Schema: user_access_control; Owner: COUNTER@2024
--

ALTER TABLE ONLY user_access_control.user_department_permissions ALTER COLUMN id SET DEFAULT nextval('user_access_control.user_department_permissions_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: user_access_control; Owner: COUNTER@2024
--

ALTER TABLE ONLY user_access_control.users ALTER COLUMN id SET DEFAULT nextval('user_access_control.users_id_seq'::regclass);


--
-- Data for Name: hypertable; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

COPY _timescaledb_catalog.hypertable (id, schema_name, table_name, associated_schema_name, associated_table_prefix, num_dimensions, chunk_sizing_func_schema, chunk_sizing_func_name, chunk_target_size, compression_state, compressed_hypertable_id, status) FROM stdin;
\.


--
-- Data for Name: chunk; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

COPY _timescaledb_catalog.chunk (id, hypertable_id, schema_name, table_name, compressed_chunk_id, dropped, status, osm_chunk, creation_time) FROM stdin;
\.


--
-- Data for Name: chunk_column_stats; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

COPY _timescaledb_catalog.chunk_column_stats (id, hypertable_id, chunk_id, column_name, range_start, range_end, valid) FROM stdin;
\.


--
-- Data for Name: dimension; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

COPY _timescaledb_catalog.dimension (id, hypertable_id, column_name, column_type, aligned, num_slices, partitioning_func_schema, partitioning_func, interval_length, compress_interval_length, integer_now_func_schema, integer_now_func) FROM stdin;
\.


--
-- Data for Name: dimension_slice; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

COPY _timescaledb_catalog.dimension_slice (id, dimension_id, range_start, range_end) FROM stdin;
\.


--
-- Data for Name: chunk_constraint; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

COPY _timescaledb_catalog.chunk_constraint (chunk_id, dimension_slice_id, constraint_name, hypertable_constraint_name) FROM stdin;
\.


--
-- Data for Name: chunk_index; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

COPY _timescaledb_catalog.chunk_index (chunk_id, index_name, hypertable_id, hypertable_index_name) FROM stdin;
\.


--
-- Data for Name: compression_chunk_size; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

COPY _timescaledb_catalog.compression_chunk_size (chunk_id, compressed_chunk_id, uncompressed_heap_size, uncompressed_toast_size, uncompressed_index_size, compressed_heap_size, compressed_toast_size, compressed_index_size, numrows_pre_compression, numrows_post_compression, numrows_frozen_immediately) FROM stdin;
\.


--
-- Data for Name: compression_settings; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

COPY _timescaledb_catalog.compression_settings (relid, segmentby, orderby, orderby_desc, orderby_nullsfirst) FROM stdin;
\.


--
-- Data for Name: continuous_agg; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

COPY _timescaledb_catalog.continuous_agg (mat_hypertable_id, raw_hypertable_id, parent_mat_hypertable_id, user_view_schema, user_view_name, partial_view_schema, partial_view_name, direct_view_schema, direct_view_name, materialized_only, finalized) FROM stdin;
\.


--
-- Data for Name: continuous_agg_migrate_plan; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

COPY _timescaledb_catalog.continuous_agg_migrate_plan (mat_hypertable_id, start_ts, end_ts, user_view_definition) FROM stdin;
\.


--
-- Data for Name: continuous_agg_migrate_plan_step; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

COPY _timescaledb_catalog.continuous_agg_migrate_plan_step (mat_hypertable_id, step_id, status, start_ts, end_ts, type, config) FROM stdin;
\.


--
-- Data for Name: continuous_aggs_bucket_function; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

COPY _timescaledb_catalog.continuous_aggs_bucket_function (mat_hypertable_id, bucket_func, bucket_width, bucket_origin, bucket_offset, bucket_timezone, bucket_fixed_width) FROM stdin;
\.


--
-- Data for Name: continuous_aggs_hypertable_invalidation_log; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

COPY _timescaledb_catalog.continuous_aggs_hypertable_invalidation_log (hypertable_id, lowest_modified_value, greatest_modified_value) FROM stdin;
\.


--
-- Data for Name: continuous_aggs_invalidation_threshold; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

COPY _timescaledb_catalog.continuous_aggs_invalidation_threshold (hypertable_id, watermark) FROM stdin;
\.


--
-- Data for Name: continuous_aggs_materialization_invalidation_log; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

COPY _timescaledb_catalog.continuous_aggs_materialization_invalidation_log (materialization_id, lowest_modified_value, greatest_modified_value) FROM stdin;
\.


--
-- Data for Name: continuous_aggs_watermark; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

COPY _timescaledb_catalog.continuous_aggs_watermark (mat_hypertable_id, watermark) FROM stdin;
\.


--
-- Data for Name: metadata; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

COPY _timescaledb_catalog.metadata (key, value, include_in_telemetry) FROM stdin;
install_timestamp	2024-10-29 20:05:08.564735+00	t
timescaledb_version	2.17.0	f
exported_uuid	0e4f788c-0f0b-4099-86ec-3532333911ce	t
\.


--
-- Data for Name: tablespace; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

COPY _timescaledb_catalog.tablespace (id, hypertable_id, tablespace_name) FROM stdin;
\.


--
-- Data for Name: bgw_job; Type: TABLE DATA; Schema: _timescaledb_config; Owner: COUNTER@2024
--

COPY _timescaledb_config.bgw_job (id, application_name, schedule_interval, max_runtime, max_retries, retry_period, proc_schema, proc_name, owner, scheduled, fixed_schedule, initial_start, hypertable_id, config, check_schema, check_name, timezone) FROM stdin;
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: COUNTER@2024
--

COPY public.schema_migrations (version, dirty) FROM stdin;
2	f
\.


--
-- Data for Name: departments; Type: TABLE DATA; Schema: user_access_control; Owner: COUNTER@2024
--

COPY user_access_control.departments (id, name) FROM stdin;
1	Admin
\.


--
-- Data for Name: permissions; Type: TABLE DATA; Schema: user_access_control; Owner: COUNTER@2024
--

COPY user_access_control.permissions (id, name, alias, parent_id) FROM stdin;
\.


--
-- Data for Name: user_department_permissions; Type: TABLE DATA; Schema: user_access_control; Owner: COUNTER@2024
--

COPY user_access_control.user_department_permissions (id, user_id, department_id, permission_id, read, write) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: user_access_control; Owner: COUNTER@2024
--

COPY user_access_control.users (id, department_id, full_name, username, password, expired_at, activated_at, created_at, updated_at) FROM stdin;
1	1	Admin	admin	$2a$12$pg01Rl0QX3VXUTTuQTAas.27qdm/RtXvFlFDPvsFwQFzyBl8YU4bC	\N	2024-10-29 20:06:48.656025+00	2024-10-29 20:06:48.656025+00	2024-10-29 20:06:48.656025+00
\.


--
-- Name: chunk_column_stats_id_seq; Type: SEQUENCE SET; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

SELECT pg_catalog.setval('_timescaledb_catalog.chunk_column_stats_id_seq', 1, false);


--
-- Name: chunk_constraint_name; Type: SEQUENCE SET; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

SELECT pg_catalog.setval('_timescaledb_catalog.chunk_constraint_name', 1, false);


--
-- Name: chunk_id_seq; Type: SEQUENCE SET; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

SELECT pg_catalog.setval('_timescaledb_catalog.chunk_id_seq', 1, false);


--
-- Name: continuous_agg_migrate_plan_step_step_id_seq; Type: SEQUENCE SET; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

SELECT pg_catalog.setval('_timescaledb_catalog.continuous_agg_migrate_plan_step_step_id_seq', 1, false);


--
-- Name: dimension_id_seq; Type: SEQUENCE SET; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

SELECT pg_catalog.setval('_timescaledb_catalog.dimension_id_seq', 1, false);


--
-- Name: dimension_slice_id_seq; Type: SEQUENCE SET; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

SELECT pg_catalog.setval('_timescaledb_catalog.dimension_slice_id_seq', 1, false);


--
-- Name: hypertable_id_seq; Type: SEQUENCE SET; Schema: _timescaledb_catalog; Owner: COUNTER@2024
--

SELECT pg_catalog.setval('_timescaledb_catalog.hypertable_id_seq', 1, false);


--
-- Name: bgw_job_id_seq; Type: SEQUENCE SET; Schema: _timescaledb_config; Owner: COUNTER@2024
--

SELECT pg_catalog.setval('_timescaledb_config.bgw_job_id_seq', 1000, false);


--
-- Name: departments_id_seq; Type: SEQUENCE SET; Schema: user_access_control; Owner: COUNTER@2024
--

SELECT pg_catalog.setval('user_access_control.departments_id_seq', 1, true);


--
-- Name: permissions_id_seq; Type: SEQUENCE SET; Schema: user_access_control; Owner: COUNTER@2024
--

SELECT pg_catalog.setval('user_access_control.permissions_id_seq', 1, false);


--
-- Name: user_department_permissions_id_seq; Type: SEQUENCE SET; Schema: user_access_control; Owner: COUNTER@2024
--

SELECT pg_catalog.setval('user_access_control.user_department_permissions_id_seq', 1, false);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: user_access_control; Owner: COUNTER@2024
--

SELECT pg_catalog.setval('user_access_control.users_id_seq', 1, true);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: COUNTER@2024
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: departments departments_pk; Type: CONSTRAINT; Schema: user_access_control; Owner: COUNTER@2024
--

ALTER TABLE ONLY user_access_control.departments
    ADD CONSTRAINT departments_pk PRIMARY KEY (id);


--
-- Name: permissions permissions_alias_uk; Type: CONSTRAINT; Schema: user_access_control; Owner: COUNTER@2024
--

ALTER TABLE ONLY user_access_control.permissions
    ADD CONSTRAINT permissions_alias_uk UNIQUE (alias);


--
-- Name: permissions permissions_pk; Type: CONSTRAINT; Schema: user_access_control; Owner: COUNTER@2024
--

ALTER TABLE ONLY user_access_control.permissions
    ADD CONSTRAINT permissions_pk PRIMARY KEY (id);


--
-- Name: user_department_permissions user_department_permissions_pk; Type: CONSTRAINT; Schema: user_access_control; Owner: COUNTER@2024
--

ALTER TABLE ONLY user_access_control.user_department_permissions
    ADD CONSTRAINT user_department_permissions_pk UNIQUE (user_id, permission_id);


--
-- Name: user_department_permissions user_department_permissions_pk_2; Type: CONSTRAINT; Schema: user_access_control; Owner: COUNTER@2024
--

ALTER TABLE ONLY user_access_control.user_department_permissions
    ADD CONSTRAINT user_department_permissions_pk_2 PRIMARY KEY (id);


--
-- Name: user_department_permissions user_department_permissions_pk_3; Type: CONSTRAINT; Schema: user_access_control; Owner: COUNTER@2024
--

ALTER TABLE ONLY user_access_control.user_department_permissions
    ADD CONSTRAINT user_department_permissions_pk_3 UNIQUE (department_id, permission_id);


--
-- Name: users users_pk; Type: CONSTRAINT; Schema: user_access_control; Owner: COUNTER@2024
--

ALTER TABLE ONLY user_access_control.users
    ADD CONSTRAINT users_pk PRIMARY KEY (id);


--
-- Name: users users_pk_2; Type: CONSTRAINT; Schema: user_access_control; Owner: COUNTER@2024
--

ALTER TABLE ONLY user_access_control.users
    ADD CONSTRAINT users_pk_2 UNIQUE (username);


--
-- Name: permissions permissions_permissions_id_fk; Type: FK CONSTRAINT; Schema: user_access_control; Owner: COUNTER@2024
--

ALTER TABLE ONLY user_access_control.permissions
    ADD CONSTRAINT permissions_permissions_id_fk FOREIGN KEY (parent_id) REFERENCES user_access_control.permissions(id);


--
-- Name: user_department_permissions user_department_permissions_departments_id_fk; Type: FK CONSTRAINT; Schema: user_access_control; Owner: COUNTER@2024
--

ALTER TABLE ONLY user_access_control.user_department_permissions
    ADD CONSTRAINT user_department_permissions_departments_id_fk FOREIGN KEY (department_id) REFERENCES user_access_control.departments(id);


--
-- Name: user_department_permissions user_department_permissions_permissions_id_fk; Type: FK CONSTRAINT; Schema: user_access_control; Owner: COUNTER@2024
--

ALTER TABLE ONLY user_access_control.user_department_permissions
    ADD CONSTRAINT user_department_permissions_permissions_id_fk FOREIGN KEY (permission_id) REFERENCES user_access_control.permissions(id);


--
-- Name: user_department_permissions user_permissions___fk; Type: FK CONSTRAINT; Schema: user_access_control; Owner: COUNTER@2024
--

ALTER TABLE ONLY user_access_control.user_department_permissions
    ADD CONSTRAINT user_permissions___fk FOREIGN KEY (user_id) REFERENCES user_access_control.users(id);


--
-- Name: users users_departments_id_fk; Type: FK CONSTRAINT; Schema: user_access_control; Owner: COUNTER@2024
--

ALTER TABLE ONLY user_access_control.users
    ADD CONSTRAINT users_departments_id_fk FOREIGN KEY (department_id) REFERENCES user_access_control.departments(id);


--
-- PostgreSQL database dump complete
--

