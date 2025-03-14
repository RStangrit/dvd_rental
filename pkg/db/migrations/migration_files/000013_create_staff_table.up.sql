CREATE SEQUENCE public.staff_staff_id_seq1
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE public.staff (
    staff_id BIGINT PRIMARY KEY DEFAULT nextval('public.staff_staff_id_seq1'::regclass),
    first_name VARCHAR(45) NOT NULL,
    last_name VARCHAR(45) NOT NULL,
    address_id INTEGER NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    store_id INTEGER NOT NULL,
    active BOOLEAN DEFAULT TRUE NOT NULL,
    username VARCHAR(16) NOT NULL,
    password VARCHAR(40) NOT NULL,
    last_update TIMESTAMP DEFAULT now() NOT NULL,
    picture BYTEA NOT NULL,
    deleted_at TIMESTAMPTZ,
    CONSTRAINT fk_address_staff FOREIGN KEY (address_id) REFERENCES public.address(address_id),
    CONSTRAINT fk_store_staff FOREIGN KEY (store_id) REFERENCES public.store(store_id)
);

ALTER SEQUENCE public.staff_staff_id_seq1
    OWNED BY public.staff.staff_id;

CREATE UNIQUE INDEX idx_staff_email ON public.staff(email);