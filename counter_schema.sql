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
-- Name: counter; Type: SCHEMA; Schema: -; Owner: COUNTER@2024
--

CREATE SCHEMA counter;


ALTER SCHEMA counter OWNER TO "COUNTER@2024";

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: item_scans; Type: TABLE; Schema: counter; Owner: COUNTER@2024
--

CREATE TABLE counter.item_scans (
    "time" timestamp with time zone NOT NULL,
    machine_id integer NOT NULL,
    qr_code_code character varying NOT NULL
);


ALTER TABLE counter.item_scans OWNER TO "COUNTER@2024";

--
-- Name: _hyper_1_1_chunk; Type: TABLE; Schema: _timescaledb_internal; Owner: COUNTER@2024
--

CREATE TABLE _timescaledb_internal._hyper_1_1_chunk (
    CONSTRAINT constraint_1 CHECK ((("time" >= '2024-10-24 00:00:00+00'::timestamp with time zone) AND ("time" < '2024-10-31 00:00:00+00'::timestamp with time zone)))
)
INHERITS (counter.item_scans);


ALTER TABLE _timescaledb_internal._hyper_1_1_chunk OWNER TO "COUNTER@2024";

--
-- Name: _hyper_1_2_chunk; Type: TABLE; Schema: _timescaledb_internal; Owner: COUNTER@2024
--

CREATE TABLE _timescaledb_internal._hyper_1_2_chunk (
    CONSTRAINT constraint_2 CHECK ((("time" >= '2024-10-31 00:00:00+00'::timestamp with time zone) AND ("time" < '2024-11-07 00:00:00+00'::timestamp with time zone)))
)
INHERITS (counter.item_scans);


ALTER TABLE _timescaledb_internal._hyper_1_2_chunk OWNER TO "COUNTER@2024";

--
-- Name: _hyper_1_3_chunk; Type: TABLE; Schema: _timescaledb_internal; Owner: COUNTER@2024
--

CREATE TABLE _timescaledb_internal._hyper_1_3_chunk (
    CONSTRAINT constraint_3 CHECK ((("time" >= '2024-11-07 00:00:00+00'::timestamp with time zone) AND ("time" < '2024-11-14 00:00:00+00'::timestamp with time zone)))
)
INHERITS (counter.item_scans);


ALTER TABLE _timescaledb_internal._hyper_1_3_chunk OWNER TO "COUNTER@2024";

--
-- Name: items; Type: TABLE; Schema: counter; Owner: COUNTER@2024
--

CREATE TABLE counter.items (
    code character varying NOT NULL,
    buyer_id integer NOT NULL,
    style_id integer NOT NULL,
    color_id integer NOT NULL,
    size_id integer NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);


ALTER TABLE counter.items OWNER TO "COUNTER@2024";

--
-- Name: setting_types; Type: TABLE; Schema: counter; Owner: COUNTER@2024
--

CREATE TABLE counter.setting_types (
    alias character varying NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE counter.setting_types OWNER TO "COUNTER@2024";

--
-- Name: settings; Type: TABLE; Schema: counter; Owner: COUNTER@2024
--

CREATE TABLE counter.settings (
    id integer NOT NULL,
    setting_type_alias character varying NOT NULL,
    value character varying,
    parent_id integer
);


ALTER TABLE counter.settings OWNER TO "COUNTER@2024";

--
-- Name: settings_id_seq; Type: SEQUENCE; Schema: counter; Owner: COUNTER@2024
--

CREATE SEQUENCE counter.settings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE counter.settings_id_seq OWNER TO "COUNTER@2024";

--
-- Name: settings_id_seq; Type: SEQUENCE OWNED BY; Schema: counter; Owner: COUNTER@2024
--

ALTER SEQUENCE counter.settings_id_seq OWNED BY counter.settings.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: COUNTER@2024
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO "COUNTER@2024";

--
-- Name: settings id; Type: DEFAULT; Schema: counter; Owner: COUNTER@2024
--

ALTER TABLE ONLY counter.settings ALTER COLUMN id SET DEFAULT nextval('counter.settings_id_seq'::regclass);


--
-- Name: items items_pk; Type: CONSTRAINT; Schema: counter; Owner: COUNTER@2024
--

ALTER TABLE ONLY counter.items
    ADD CONSTRAINT items_pk PRIMARY KEY (code);


--
-- Name: setting_types setting_types_pk; Type: CONSTRAINT; Schema: counter; Owner: COUNTER@2024
--

ALTER TABLE ONLY counter.setting_types
    ADD CONSTRAINT setting_types_pk PRIMARY KEY (alias);


--
-- Name: settings settings_pk; Type: CONSTRAINT; Schema: counter; Owner: COUNTER@2024
--

ALTER TABLE ONLY counter.settings
    ADD CONSTRAINT settings_pk PRIMARY KEY (id);


--
-- Name: settings settings_pk_2; Type: CONSTRAINT; Schema: counter; Owner: COUNTER@2024
--

ALTER TABLE ONLY counter.settings
    ADD CONSTRAINT settings_pk_2 UNIQUE (setting_type_alias, value, parent_id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: COUNTER@2024
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: _hyper_1_1_chunk_idx_machine_id_time; Type: INDEX; Schema: _timescaledb_internal; Owner: COUNTER@2024
--

CREATE UNIQUE INDEX _hyper_1_1_chunk_idx_machine_id_time ON _timescaledb_internal._hyper_1_1_chunk USING btree (machine_id, "time");


--
-- Name: _hyper_1_1_chunk_item_scans_time_idx; Type: INDEX; Schema: _timescaledb_internal; Owner: COUNTER@2024
--

CREATE INDEX _hyper_1_1_chunk_item_scans_time_idx ON _timescaledb_internal._hyper_1_1_chunk USING btree ("time" DESC);


--
-- Name: _hyper_1_2_chunk_idx_machine_id_time; Type: INDEX; Schema: _timescaledb_internal; Owner: COUNTER@2024
--

CREATE UNIQUE INDEX _hyper_1_2_chunk_idx_machine_id_time ON _timescaledb_internal._hyper_1_2_chunk USING btree (machine_id, "time");


--
-- Name: _hyper_1_2_chunk_item_scans_time_idx; Type: INDEX; Schema: _timescaledb_internal; Owner: COUNTER@2024
--

CREATE INDEX _hyper_1_2_chunk_item_scans_time_idx ON _timescaledb_internal._hyper_1_2_chunk USING btree ("time" DESC);


--
-- Name: _hyper_1_3_chunk_idx_machine_id_time; Type: INDEX; Schema: _timescaledb_internal; Owner: COUNTER@2024
--

CREATE UNIQUE INDEX _hyper_1_3_chunk_idx_machine_id_time ON _timescaledb_internal._hyper_1_3_chunk USING btree (machine_id, "time");


--
-- Name: _hyper_1_3_chunk_item_scans_time_idx; Type: INDEX; Schema: _timescaledb_internal; Owner: COUNTER@2024
--

CREATE INDEX _hyper_1_3_chunk_item_scans_time_idx ON _timescaledb_internal._hyper_1_3_chunk USING btree ("time" DESC);


--
-- Name: idx_machine_id_time; Type: INDEX; Schema: counter; Owner: COUNTER@2024
--

CREATE UNIQUE INDEX idx_machine_id_time ON counter.item_scans USING btree (machine_id, "time");


--
-- Name: item_scans_time_idx; Type: INDEX; Schema: counter; Owner: COUNTER@2024
--

CREATE INDEX item_scans_time_idx ON counter.item_scans USING btree ("time" DESC);


--
-- Name: item_scans ts_insert_blocker; Type: TRIGGER; Schema: counter; Owner: COUNTER@2024
--

CREATE TRIGGER ts_insert_blocker BEFORE INSERT ON counter.item_scans FOR EACH ROW EXECUTE FUNCTION _timescaledb_functions.insert_blocker();


--
-- Name: items items_settings_id_fk; Type: FK CONSTRAINT; Schema: counter; Owner: COUNTER@2024
--

ALTER TABLE ONLY counter.items
    ADD CONSTRAINT items_settings_id_fk FOREIGN KEY (buyer_id) REFERENCES counter.settings(id);


--
-- Name: items items_settings_id_fk_2; Type: FK CONSTRAINT; Schema: counter; Owner: COUNTER@2024
--

ALTER TABLE ONLY counter.items
    ADD CONSTRAINT items_settings_id_fk_2 FOREIGN KEY (style_id) REFERENCES counter.settings(id);


--
-- Name: items items_settings_id_fk_3; Type: FK CONSTRAINT; Schema: counter; Owner: COUNTER@2024
--

ALTER TABLE ONLY counter.items
    ADD CONSTRAINT items_settings_id_fk_3 FOREIGN KEY (color_id) REFERENCES counter.settings(id);


--
-- Name: items items_settings_id_fk_4; Type: FK CONSTRAINT; Schema: counter; Owner: COUNTER@2024
--

ALTER TABLE ONLY counter.items
    ADD CONSTRAINT items_settings_id_fk_4 FOREIGN KEY (size_id) REFERENCES counter.settings(id);


--
-- Name: settings settings_setting_types_alias_fk; Type: FK CONSTRAINT; Schema: counter; Owner: COUNTER@2024
--

ALTER TABLE ONLY counter.settings
    ADD CONSTRAINT settings_setting_types_alias_fk FOREIGN KEY (setting_type_alias) REFERENCES counter.setting_types(alias);


--
-- PostgreSQL database dump complete
--

