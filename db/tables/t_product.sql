-- drop table t_product;
create table t_product (
	id integer default nextval('"s_product"'::regclass),
	code varchar not null,
	prod_name varchar not null,
	weight numeric not null,
	description varchar not null
);

alter table t_product add constraint t_product_pk primary key (id);
alter table t_product add constraint t_product_uq1 unique (code);
-- 
comment on table t_product is 'Table that contains products';