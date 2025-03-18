CREATE SEQUENCE public.payment_payment_id_seq1
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE public.payment (
    payment_id BIGINT PRIMARY KEY DEFAULT nextval('public.payment_payment_id_seq1'),
    customer_id INTEGER NOT NULL,
    staff_id INTEGER NOT NULL,
    rental_id INTEGER NOT NULL,
    amount NUMERIC(5,2) NOT NULL,
    payment_date TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ,
    CONSTRAINT fk_customer_payments FOREIGN KEY (customer_id) REFERENCES public.customer(customer_id),
    CONSTRAINT fk_rental_payment FOREIGN KEY (rental_id) REFERENCES public.rental(rental_id),
    CONSTRAINT fk_staff_payments FOREIGN KEY (staff_id) REFERENCES public.staff(staff_id)
);

ALTER SEQUENCE public.payment_payment_id_seq1
    OWNED BY public.payment.payment_id;