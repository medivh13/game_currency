CREATE TABLE game_currency.currencies (
    id bigint NOT NULL,
    name character varying(50) NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone
);

CREATE SEQUENCE game_currency.currencies_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE game_currency.currencies_id_seq OWNED BY game_currency.currencies.id;

ALTER TABLE ONLY game_currency.currencies ALTER COLUMN id SET DEFAULT nextval('game_currency.currencies_id_seq'::regclass);

ALTER TABLE ONLY game_currency.currencies
    ADD CONSTRAINT currencies_pkey PRIMARY KEY (id);

ALTER TABLE ONLY game_currency.currencies
    ADD CONSTRAINT currencies_account_number_key UNIQUE (name);

CREATE INDEX currenciesindex ON game_currency.currencies USING btree (deleted_at);

insert into game_currency.currencies (name) values ('knut');
insert into game_currency.currencies (name) values ('sickle');
insert into game_currency.currencies (name) values ('galleon');