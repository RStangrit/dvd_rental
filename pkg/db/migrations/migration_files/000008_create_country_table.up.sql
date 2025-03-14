CREATE SEQUENCE public.country_country_id_seq1
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE public.country (
    country_id BIGINT PRIMARY KEY DEFAULT nextval('public.country_country_id_seq1'::regclass),
    country VARCHAR(50) NOT NULL,
    last_update TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ
);

ALTER SEQUENCE public.country_country_id_seq1
    OWNED BY public.country.country_id;
