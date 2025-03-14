CREATE SEQUENCE public.customer_customer_id_seq1
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE public.customer (
    customer_id BIGINT PRIMARY KEY DEFAULT nextval('public.customer_customer_id_seq1'::regclass),
    store_id SMALLINT NOT NULL,
    first_name VARCHAR(45) NOT NULL,
    last_name VARCHAR(45) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    address_id INTEGER NOT NULL,
    activebool BOOLEAN DEFAULT TRUE NOT NULL,
    create_date DATE DEFAULT CURRENT_DATE NOT NULL,
    last_update TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,
    active INTEGER NOT NULL,
    CONSTRAINT fk_address_customer FOREIGN KEY (address_id) REFERENCES public.address(address_id)
);

ALTER SEQUENCE public.customer_customer_id_seq1
    OWNED BY public.customer.customer_id;

CREATE UNIQUE INDEX idx_customer_email ON public.customer(email);