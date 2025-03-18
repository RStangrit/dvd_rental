CREATE SEQUENCE public.address_address_id_seq1
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE public.address (
    address_id BIGINT PRIMARY KEY DEFAULT nextval('public.address_address_id_seq1'::regclass),
    address VARCHAR(50) NOT NULL,
    address2 VARCHAR(50),
    district VARCHAR(20) NOT NULL,
    city_id INTEGER NOT NULL,
    postal_code VARCHAR(10),
    phone VARCHAR(20) NOT NULL,
    last_update TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ,
    CONSTRAINT fk_city_addresses FOREIGN KEY (city_id) REFERENCES public.city(city_id)
);

ALTER SEQUENCE public.address_address_id_seq1
    OWNED BY public.address.address_id;