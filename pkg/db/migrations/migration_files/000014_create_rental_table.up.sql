CREATE SEQUENCE public.rental_rental_id_seq1
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE public.rental (
    rental_id BIGINT PRIMARY KEY DEFAULT nextval('public.rental_rental_id_seq1'),
    rental_date TIMESTAMP NOT NULL,
    inventory_id INTEGER NOT NULL,
    customer_id INTEGER NOT NULL,
    return_date TIMESTAMP NOT NULL,
    staff_id INTEGER NOT NULL,
    last_update TIMESTAMP DEFAULT now() NOT NULL,
    deleted_at TIMESTAMPTZ,
    CONSTRAINT fk_customer_rentals FOREIGN KEY (customer_id) REFERENCES public.customer(customer_id),
    CONSTRAINT fk_inventory_rental FOREIGN KEY (inventory_id) REFERENCES public.inventory(inventory_id),
    CONSTRAINT fk_staff_rentals FOREIGN KEY (staff_id) REFERENCES public.staff(staff_id)
);

ALTER SEQUENCE public.rental_rental_id_seq1
    OWNED BY public.rental.rental_id;