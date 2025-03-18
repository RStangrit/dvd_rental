CREATE SEQUENCE public.store_store_id_seq1
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE public.store (
    store_id BIGINT PRIMARY KEY DEFAULT nextval('public.store_store_id_seq1'::regclass),
    manager_staff_id SMALLINT NOT NULL,
    address_id INTEGER NOT NULL,
    last_update TIMESTAMP DEFAULT now() NOT NULL,
    deleted_at TIMESTAMPTZ,
    CONSTRAINT fk_address_store FOREIGN KEY (address_id) REFERENCES public.address(address_id)
);

ALTER SEQUENCE public.store_store_id_seq1
    OWNED BY public.store.store_id;