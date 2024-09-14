--
-- PostgreSQL database dump
--

-- Dumped from database version 12.12
-- Dumped by pg_dump version 14.10 (Homebrew)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: selector_score_daily; Type: TABLE; Schema: chs; Owner: centos
--

CREATE TABLE chs.selector_score_daily (
    client_id uuid NOT NULL,
    identity_id uuid NOT NULL,
    selector_type identity.selector_type NOT NULL,
    selector_id uuid NOT NULL,
    date date NOT NULL,
    score double precision,
    breakdown jsonb,
    breakdown_details jsonb,
    date_added timestamp without time zone DEFAULT timezone('utc'::text, CURRENT_TIMESTAMP),
    date_updated timestamp without time zone DEFAULT timezone('utc'::text, CURRENT_TIMESTAMP)
);


ALTER TABLE chs.selector_score_daily OWNER TO centos;

--
-- Name: selector_score_daily selector_score_daily_pkey; Type: CONSTRAINT; Schema: chs; Owner: centos
--

ALTER TABLE ONLY chs.selector_score_daily
    ADD CONSTRAINT selector_score_daily_pkey PRIMARY KEY (client_id, identity_id, selector_type, selector_id, date);


--
-- Name: selector_score_daily selector_score_daily_identity_id_fkey; Type: FK CONSTRAINT; Schema: chs; Owner: centos
--

ALTER TABLE ONLY chs.selector_score_daily
    ADD CONSTRAINT selector_score_daily_identity_id_fkey FOREIGN KEY (identity_id) REFERENCES identity.identity(id);


--
-- PostgreSQL database dump complete
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 12.12
-- Dumped by pg_dump version 14.10 (Homebrew)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: alembic_version; Type: TABLE; Schema: chs; Owner: centos
--

CREATE TABLE chs.alembic_version (
    version_num character varying(32) NOT NULL
);


ALTER TABLE chs.alembic_version OWNER TO centos;

--
-- Name: alembic_version alembic_version_pkc; Type: CONSTRAINT; Schema: chs; Owner: centos
--

ALTER TABLE ONLY chs.alembic_version
    ADD CONSTRAINT alembic_version_pkc PRIMARY KEY (version_num);


--
-- PostgreSQL database dump complete
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 12.12
-- Dumped by pg_dump version 14.10 (Homebrew)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: identity_score_daily; Type: TABLE; Schema: chs; Owner: centos
--

CREATE TABLE chs.identity_score_daily (
    client_id uuid NOT NULL,
    identity_id uuid NOT NULL,
    date date NOT NULL,
    score double precision,
    breakdown jsonb,
    breakdown_details jsonb,
    date_added timestamp without time zone DEFAULT timezone('utc'::text, CURRENT_TIMESTAMP),
    date_updated timestamp without time zone DEFAULT timezone('utc'::text, CURRENT_TIMESTAMP)
);


ALTER TABLE chs.identity_score_daily OWNER TO centos;

--
-- Name: identity_score_daily identity_score_daily_pkey; Type: CONSTRAINT; Schema: chs; Owner: centos
--

ALTER TABLE ONLY chs.identity_score_daily
    ADD CONSTRAINT identity_score_daily_pkey PRIMARY KEY (client_id, identity_id, date);


--
-- Name: identity_score_daily identity_score_daily_identity_id_fkey; Type: FK CONSTRAINT; Schema: chs; Owner: centos
--

ALTER TABLE ONLY chs.identity_score_daily
    ADD CONSTRAINT identity_score_daily_identity_id_fkey FOREIGN KEY (identity_id) REFERENCES identity.identity(id);


--
-- PostgreSQL database dump complete
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 12.12
-- Dumped by pg_dump version 14.10 (Homebrew)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: identity_score_monthly; Type: TABLE; Schema: chs; Owner: centos
--

CREATE TABLE chs.identity_score_monthly (
    client_id uuid NOT NULL,
    identity_id uuid NOT NULL,
    date date NOT NULL,
    score double precision,
    breakdown jsonb,
    breakdown_details jsonb,
    date_added timestamp without time zone DEFAULT timezone('utc'::text, CURRENT_TIMESTAMP),
    date_updated timestamp without time zone DEFAULT timezone('utc'::text, CURRENT_TIMESTAMP)
);


ALTER TABLE chs.identity_score_monthly OWNER TO centos;

--
-- Name: identity_score_monthly identity_score_monthly_pkey; Type: CONSTRAINT; Schema: chs; Owner: centos
--

ALTER TABLE ONLY chs.identity_score_monthly
    ADD CONSTRAINT identity_score_monthly_pkey PRIMARY KEY (client_id, identity_id, date);


--
-- Name: identity_score_monthly identity_score_monthly_identity_id_fkey; Type: FK CONSTRAINT; Schema: chs; Owner: centos
--

ALTER TABLE ONLY chs.identity_score_monthly
    ADD CONSTRAINT identity_score_monthly_identity_id_fkey FOREIGN KEY (identity_id) REFERENCES identity.identity(id);


--
-- PostgreSQL database dump complete
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 12.12
-- Dumped by pg_dump version 14.10 (Homebrew)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: selector_score_monthly; Type: TABLE; Schema: chs; Owner: centos
--

CREATE TABLE chs.selector_score_monthly (
    client_id uuid NOT NULL,
    identity_id uuid NOT NULL,
    selector_type identity.selector_type NOT NULL,
    selector_id uuid NOT NULL,
    date date NOT NULL,
    score double precision,
    breakdown jsonb,
    breakdown_details jsonb,
    date_added timestamp without time zone DEFAULT timezone('utc'::text, CURRENT_TIMESTAMP),
    date_updated timestamp without time zone DEFAULT timezone('utc'::text, CURRENT_TIMESTAMP)
);


ALTER TABLE chs.selector_score_monthly OWNER TO centos;

--
-- Name: selector_score_monthly selector_score_monthly_pkey; Type: CONSTRAINT; Schema: chs; Owner: centos
--

ALTER TABLE ONLY chs.selector_score_monthly
    ADD CONSTRAINT selector_score_monthly_pkey PRIMARY KEY (client_id, identity_id, selector_type, selector_id, date);


--
-- Name: selector_score_monthly selector_score_monthly_identity_id_fkey; Type: FK CONSTRAINT; Schema: chs; Owner: centos
--

ALTER TABLE ONLY chs.selector_score_monthly
    ADD CONSTRAINT selector_score_monthly_identity_id_fkey FOREIGN KEY (identity_id) REFERENCES identity.identity(id);


--
-- PostgreSQL database dump complete
--

