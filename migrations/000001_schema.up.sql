
DO
$do$
BEGIN
   IF EXISTS (SELECT FROM pg_database WHERE datname = 'counter_db') THEN
      RAISE NOTICE 'Database already exists';  -- optional
   ELSE
      PERFORM dblink_exec('dbname=' || current_database()  -- current db
                        , 'CREATE DATABASE counter_db');
   END IF;
END
$do$;

create schema if not exists counter;

create table counter.setting_types
(
    alias varchar not null
        constraint setting_types_pk
            primary key,
    name  varchar not null
);

alter table counter.setting_types
    owner to "COUNTER@2024";

create table counter.settings
(
    id                 serial
        constraint settings_pk
            primary key,
    setting_type_alias varchar not null
        constraint settings_setting_types_alias_fk
            references counter.setting_types,
    value              varchar,
    parent_id          integer,
    constraint settings_pk_2
        unique (setting_type_alias, value, parent_id)
);

alter table counter.settings
    owner to "COUNTER@2024";

create table counter.items
(
    code     varchar not null
        constraint items_pk
            primary key,
    buyer_id integer not null
        constraint items_settings_id_fk
            references counter.settings,
    style_id integer not null
        constraint items_settings_id_fk_2
            references counter.settings,
    color_id integer not null
        constraint items_settings_id_fk_3
            references counter.settings,
    size_id  integer not null
        constraint items_settings_id_fk_4
            references counter.settings,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

alter table counter.items
    owner to "COUNTER@2024";

create table counter.item_scans
(
    time         timestamp with time zone,
    machine_id   integer not null,
    qr_code_code varchar not null
);

alter table counter.item_scans
    owner to "COUNTER@2024";

create unique index idx_machine_id_time
    on counter.item_scans (machine_id, time);

select * from public.create_hypertable('counter.item_scans', public.by_range('time'));
