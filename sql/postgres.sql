--
-- PostgreSQL database dump
--

-- Dumped from database version 11.18 (Ubuntu 11.18-1.pgdg18.04+1)
-- Dumped by pg_dump version 14.6

-- Started on 2023-02-16 16:20:26 WIB

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
-- TOC entry 24 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO postgres;

--
-- TOC entry 4007 (class 0 OID 0)
-- Dependencies: 24
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

--
-- TOC entry 222 (class 1259 OID 10491051)
-- Name: users; Type: TABLE; Schema: public; Owner: zixjtwbp
--

CREATE TABLE public.users (
    user_id character varying NOT NULL,
    name character varying,
    email character varying,
    password character varying,
    date_created timestamp without time zone,
    date_updated timestamp without time zone
);


ALTER TABLE public.users OWNER TO zixjtwbp;

--
-- TOC entry 4001 (class 0 OID 10491051)
-- Dependencies: 222
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: zixjtwbp
--

COPY public.users (user_id, name, email, password, date_created, date_updated) FROM stdin;
a5f7ae2a-aaf9-11ed-8350-60599c074d6b	Habib Irfan Mahaasin	admin@mahaasin.com	$2a$10$BHZ3hC8RnNXRwmZOT83E4.enMiY9ARlmOjxd/C0sR2Uu5HNA18QQq	2023-02-16 15:21:08.755693	2023-02-16 15:21:08.755693
5015b561-5ceb-49a7-b673-78cebd134c50	Super Admin	superadmin@mahaasin.com	$2a$10$14EWUaUllmqiQcoeCmjnTuBcs9Z2by.oyoIbd3h7loHu.BUyZ63kW	2023-02-16 16:15:22.310404	2023-02-16 16:15:22.310404
\.


--
-- TOC entry 3878 (class 2606 OID 10491058)
-- Name: users user_pk; Type: CONSTRAINT; Schema: public; Owner: zixjtwbp
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT user_pk PRIMARY KEY (user_id);


-- Completed on 2023-02-16 16:20:31 WIB

--
-- PostgreSQL database dump complete
--

