CREATE TABLE game_currency.conversion_rates (
    id bigint NOT NULL,
    currency_id_from bigint NOT NULL,
    currency_id_to bigint NOT NULL,
    rate double precision NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone
);

CREATE SEQUENCE game_currency.conversion_rates_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE game_currency.conversion_rates_id_seq OWNED BY game_currency.conversion_rates.id;

ALTER TABLE ONLY game_currency.conversion_rates ALTER COLUMN id SET DEFAULT nextval('game_currency.conversion_rates_id_seq'::regclass);

ALTER TABLE ONLY game_currency.conversion_rates
    ADD CONSTRAINT conversion_rates_pkey PRIMARY KEY (id);

CREATE INDEX conversion_ratesindex ON game_currency.conversion_rates USING btree (deleted_at);

insert into game_currency.conversion_rates (currency_id_from, currency_id_to, rate) 
values (2,1,29);