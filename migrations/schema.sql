--
-- PostgreSQL database dump
--

-- Dumped from database version 14.3
-- Dumped by pg_dump version 14.3

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
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: erald
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO erald;

--
-- Name: stats; Type: TABLE; Schema: public; Owner: erald
--

CREATE TABLE public.stats (
    id integer NOT NULL,
    date timestamp with time zone NOT NULL,
    breakfast integer NOT NULL,
    lunch integer NOT NULL,
    dinner integer NOT NULL,
    snacks integer NOT NULL,
    protein integer NOT NULL,
    carbs integer NOT NULL,
    fats integer NOT NULL,
    user_id integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.stats OWNER TO erald;

--
-- Name: stats_id_seq; Type: SEQUENCE; Schema: public; Owner: erald
--

CREATE SEQUENCE public.stats_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.stats_id_seq OWNER TO erald;

--
-- Name: stats_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: erald
--

ALTER SEQUENCE public.stats_id_seq OWNED BY public.stats.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: erald
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(60) NOT NULL,
    access_level integer DEFAULT 1 NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users OWNER TO erald;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: erald
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO erald;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: erald
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: stats id; Type: DEFAULT; Schema: public; Owner: erald
--

ALTER TABLE ONLY public.stats ALTER COLUMN id SET DEFAULT nextval('public.stats_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: erald
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: schema_migration schema_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: erald
--

ALTER TABLE ONLY public.schema_migration
    ADD CONSTRAINT schema_migration_pkey PRIMARY KEY (version);


--
-- Name: stats stats_pkey; Type: CONSTRAINT; Schema: public; Owner: erald
--

ALTER TABLE ONLY public.stats
    ADD CONSTRAINT stats_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: erald
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: erald
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: users_email_idx; Type: INDEX; Schema: public; Owner: erald
--

CREATE UNIQUE INDEX users_email_idx ON public.users USING btree (email);


--
-- Name: stats stats_users_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: erald
--

ALTER TABLE ONLY public.stats
    ADD CONSTRAINT stats_users_id_fk FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

