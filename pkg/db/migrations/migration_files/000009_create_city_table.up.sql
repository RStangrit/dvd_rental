CREATE SEQUENCE public.city_city_id_seq1
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE public.city (
    city_id BIGINT PRIMARY KEY DEFAULT nextval('public.city_city_id_seq1'::regclass),
    city VARCHAR(50) NOT NULL,
    country_id INTEGER NOT NULL,
    last_update TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ,
    CONSTRAINT fk_country_cities FOREIGN KEY (country_id) REFERENCES public.country(country_id)
);

ALTER SEQUENCE public.city_city_id_seq1
    OWNED BY public.city.city_id;