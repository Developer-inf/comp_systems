create table t_product (
	id numeric,
	code varchar,
	prod_name varchar not null,
	weight numeric not null,
	description varchar not null
);

alter table t_product add constraint t_product_pk primary key (id);