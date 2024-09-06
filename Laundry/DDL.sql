CREATE DATABASE enigma-laudry;

CREATE TABLE mst_customer(
    id bigserial primary key,
    nama varchar,
    phone_number varchar(13),
);

CREATE TABLE mst_employee(
    id bigserial primary key,
    nama varchar,
    phone_number varchar(13),
);

CREATE TABLE mst_service(
    id bigserial primary key,
    nama varchar,
    satuan char(6),
    price integer
);

CREATE TABLE mst_customer(
    id bigserial primary key,
    nama varchar,
    phone_number varchar(13),
);

CREATE TABLE tr_trasaction(
    id bigserial primary key,
    id_cust integer,
    id_emp integer,
    id_srv integer,
    jumlah integer,
    total integer,
    FOREIGN KEY (id_cust) REFERENCES mst_customer(id),
    FOREIGN KEY (id_emp) REFERENCES mst_employee(id),
    FOREIGN KEY (id_srv) REFERENCES mst_service(id),

);
