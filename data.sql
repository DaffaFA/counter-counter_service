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

