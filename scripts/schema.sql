--
-- PostgreSQL database dump
--

-- Dumped from database version 13.3 (Ubuntu 13.3-1.pgdg20.04+1)
-- Dumped by pg_dump version 13.3 (Ubuntu 13.3-1.pgdg20.04+1)

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
-- Name: wfg; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA wfg;

SET default_tablespace = '';

--SET default_table_access_method = heap;

--
-- Name: category; Type: TABLE; Schema: wfg; Owner: -
--

CREATE TABLE wfg.category (
    category_id integer NOT NULL,
    category_name character varying(255) NOT NULL,
    description character varying(255)
);


--
-- Name: city; Type: TABLE; Schema: wfg; Owner: -
--

CREATE TABLE wfg.city (
    city_id integer NOT NULL,
    city_name character varying(255) NOT NULL,
    country_id integer NOT NULL
);


--
-- Name: country; Type: TABLE; Schema: wfg; Owner: -
--

CREATE TABLE wfg.country (
    country_id integer NOT NULL,
    country_name character varying(255),
    country_name_es character varying(255),
    iso2 character varying(5),
    iso3 character varying(5),
    phone_code character varying(10)
);


--
-- Name: customer; Type: TABLE; Schema: wfg; Owner: -
--

CREATE TABLE wfg.customer (
    customer_id integer NOT NULL,
    name character varying(255) NOT NULL,
    address character varying(255) NOT NULL,
    website character varying(255) NOT NULL,
    credit_limit double precision NOT NULL
);


--
-- Name: customer_contact; Type: TABLE; Schema: wfg; Owner: -
--

CREATE TABLE wfg.customer_contact (
    contact_id integer NOT NULL,
    first_name character varying(255) NOT NULL,
    last_name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    phone character varying(15),
    customer_id integer NOT NULL
);


--
-- Name: district; Type: TABLE; Schema: wfg; Owner: -
--

CREATE TABLE wfg.district (
    district_id integer NOT NULL,
    district_name character varying(255),
    address character varying(255),
    postal_code text,
    warehouse_id integer
);


--
-- Name: employee; Type: TABLE; Schema: wfg; Owner: -
--

CREATE TABLE wfg.employee (
    employee_id integer NOT NULL,
    first_name character varying(255) NOT NULL,
    last_name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    phone character varying(20) NOT NULL,
    hire_date timestamp without time zone NOT NULL,
    manager_id integer,
    job_title character varying(255)
);


--
-- Name: history_district; Type: TABLE; Schema: wfg; Owner: -
--

CREATE TABLE wfg.history_district (
    district_id integer NOT NULL,
    product_id integer NOT NULL,
    change_inventory timestamp without time zone NOT NULL,
    sumed_quantity integer NOT NULL
);


--
-- Name: history_order; Type: TABLE; Schema: wfg; Owner: -
--

CREATE TABLE wfg.history_order (
    order_id integer NOT NULL,
    change_order timestamp without time zone,
    status_id integer
);


--
-- Name: inventory; Type: TABLE; Schema: wfg; Owner: -
--

CREATE TABLE wfg.inventory (
    district_id integer NOT NULL,
    product_id integer NOT NULL,
    quantity integer NOT NULL
);


--
-- Name: order_items; Type: TABLE; Schema: wfg; Owner: -
--

CREATE TABLE wfg.order_items (
    order_id UUID NOT NULL,
    item_id integer NOT NULL,
    product_id integer NOT NULL,
    quantity integer NOT NULL,
    unit_price double precision NOT NULL,
    district_id integer NOT NULL
);


--
-- Name: order_status; Type: TABLE; Schema: wfg; Owner: -
--

CREATE TABLE wfg.order_status (
    status_id integer NOT NULL,
    status_name character varying NOT NULL
);


--
-- Name: orders; Type: TABLE; Schema: wfg; Owner: -
--

CREATE TABLE wfg.orders (
    order_id UUID NOT NULL,
    customer_id integer,
    status_id integer,
    salesman_id integer,
    order_date timestamp without time zone,
    invoice_no character varying(255)
);


--
-- Name: product; Type: TABLE; Schema: wfg; Owner: -
--

CREATE TABLE wfg.product (
    product_id integer NOT NULL,
    barcode character varying NOT NULL,
    product_name character varying NOT NULL,
    description character varying NOT NULL,
    standard_cost double precision NOT NULL,
    list_price double precision NOT NULL,
    category_id integer
);


--
-- Name: region; Type: TABLE; Schema: wfg; Owner: -
--

CREATE TABLE wfg.region (
    region_id integer NOT NULL,
    region_name character varying(255) NOT NULL
);


--
-- Name: region_country; Type: TABLE; Schema: wfg; Owner: -
--

CREATE TABLE wfg.region_country (
    region_id integer NOT NULL,
    country_id integer NOT NULL
);


--
-- Name: warehouse; Type: TABLE; Schema: wfg; Owner: -
--

CREATE TABLE wfg.warehouse (
    warehouse_id integer NOT NULL,
    warehouse_name character varying(255) NOT NULL,
    city_id integer,
    address character varying(255),
    postal_code text
);


--
-- Name: category category_pkey; Type: CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.category
    ADD CONSTRAINT category_pkey PRIMARY KEY (category_id);


--
-- Name: country country_pkey; Type: CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.country
    ADD CONSTRAINT country_pkey PRIMARY KEY (country_id);


--
-- Name: customer_contact customer_contact_pkey; Type: CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.customer_contact
    ADD CONSTRAINT customer_contact_pkey PRIMARY KEY (contact_id, customer_id);


--
-- Name: customer customer_pkey; Type: CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.customer
    ADD CONSTRAINT customer_pkey PRIMARY KEY (customer_id);


--
-- Name: district district_pkey; Type: CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.district
    ADD CONSTRAINT district_pkey PRIMARY KEY (district_id);


--
-- Name: employee employee_pkey; Type: CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.employee
    ADD CONSTRAINT employee_pkey PRIMARY KEY (employee_id);


--
-- Name: history_district history_district_pkey; Type: CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.history_district
    ADD CONSTRAINT history_district_pkey PRIMARY KEY (district_id, product_id, change_inventory);


--
-- Name: history_order history_order_pkey; Type: CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.history_order
    ADD CONSTRAINT history_order_pkey PRIMARY KEY (order_id);


--
-- Name: inventory inventory_pkey; Type: CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.inventory
    ADD CONSTRAINT inventory_pkey PRIMARY KEY (district_id, product_id);


--
-- Name: city city_pkey; Type: CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.city
    ADD CONSTRAINT city_pkey PRIMARY KEY (city_id);


--
-- Name: order_items order_items_pkey; Type: CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.order_items
    ADD CONSTRAINT order_items_pkey PRIMARY KEY (order_id, item_id);


--
-- Name: orders order_pkey; Type: CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.orders
    ADD CONSTRAINT order_pkey PRIMARY KEY (order_id);


--
-- Name: product product_pkey; Type: CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.product
    ADD CONSTRAINT product_pkey PRIMARY KEY (product_id);


--
-- Name: region_country region_country_pkey; Type: CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.region_country
    ADD CONSTRAINT region_country_pkey PRIMARY KEY (region_id, country_id);


--
-- Name: region region_pkey; Type: CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.region
    ADD CONSTRAINT region_pkey PRIMARY KEY (region_id);


--
-- Name: order_status status_order_pkey; Type: CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.order_status
    ADD CONSTRAINT status_order_pkey PRIMARY KEY (status_id);


--
-- Name: warehouse warehouse_pkey; Type: CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.warehouse
    ADD CONSTRAINT warehouse_pkey PRIMARY KEY (warehouse_id);


--
-- Name: product category_fk; Type: FK CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.product
    ADD CONSTRAINT category_fk FOREIGN KEY (category_id) REFERENCES wfg.category(category_id) NOT VALID;


--
-- Name: warehouse city_fk; Type: FK CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.warehouse
    ADD CONSTRAINT city_fk FOREIGN KEY (city_id) REFERENCES wfg.city(city_id);


--
-- Name: customer_contact customer_fk; Type: FK CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.customer_contact
    ADD CONSTRAINT customer_fk FOREIGN KEY (customer_id) REFERENCES wfg.customer(customer_id);


--
-- Name: orders customer_fk; Type: FK CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.orders
    ADD CONSTRAINT customer_fk FOREIGN KEY (customer_id) REFERENCES wfg.customer(customer_id);


--
-- Name: inventory district_fk1; Type: FK CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.inventory
    ADD CONSTRAINT district_fk1 FOREIGN KEY (district_id) REFERENCES wfg.district(district_id) NOT VALID;


--
-- Name: order_items inventory_dist_fk; Type: FK CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.order_items
    ADD CONSTRAINT inventory_dist_fk FOREIGN KEY (product_id, district_id) REFERENCES wfg.inventory(product_id, district_id) NOT VALID;


--
-- Name: history_district inventory_history_district_fk; Type: FK CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.history_district
    ADD CONSTRAINT inventory_history_district_fk FOREIGN KEY (district_id, product_id) REFERENCES wfg.inventory(district_id, product_id);


--
-- Name: city location_country_id_fkey; Type: FK CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.city
    ADD CONSTRAINT location_country_id_fkey FOREIGN KEY (country_id) REFERENCES wfg.country(country_id);


--
-- Name: order_items order_fk; Type: FK CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.order_items
    ADD CONSTRAINT order_fk FOREIGN KEY (order_id) REFERENCES wfg.orders(order_id);


--
-- Name: inventory product_fk; Type: FK CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.inventory
    ADD CONSTRAINT product_fk FOREIGN KEY (product_id) REFERENCES wfg.product(product_id) NOT VALID;


--
-- Name: orders salesman_fk; Type: FK CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.orders
    ADD CONSTRAINT salesman_fk FOREIGN KEY (salesman_id) REFERENCES wfg.employee(employee_id);


--
-- Name: orders status_fk; Type: FK CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.orders
    ADD CONSTRAINT status_fk FOREIGN KEY (status_id) REFERENCES wfg.order_status(status_id) NOT VALID;


--
-- Name: history_order status_ho_fk; Type: FK CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.history_order
    ADD CONSTRAINT status_ho_fk FOREIGN KEY (status_id) REFERENCES wfg.order_status(status_id);


--
-- Name: district warehouse_fk1; Type: FK CONSTRAINT; Schema: wfg; Owner: -
--

ALTER TABLE ONLY wfg.district
    ADD CONSTRAINT warehouse_fk1 FOREIGN KEY (warehouse_id) REFERENCES wfg.warehouse(warehouse_id);


--
-- PostgreSQL database dump complete
--

