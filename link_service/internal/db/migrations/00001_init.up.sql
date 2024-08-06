begin;

create table if not exists links (
  slug character varying(5) not null,
  date timestamp without time zone, 
  metadata jsonb,
  link character varying(256)
);
alter table only links add primary key (slug);

commit;
