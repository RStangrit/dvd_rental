CREATE SEQUENCE public.inventory_inventory_id_seq1
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE public.inventory (
    inventory_id BIGINT PRIMARY KEY DEFAULT nextval('public.inventory_inventory_id_seq1'::regclass),
    film_id INTEGER NOT NULL,
    store_id SMALLINT NOT NULL,
    last_update TIMESTAMP DEFAULT now() NOT NULL,
    deleted_at TIMESTAMPTZ,
    CONSTRAINT fk_film_films_inventory FOREIGN KEY (film_id) REFERENCES public.film(film_id)
);

ALTER SEQUENCE public.inventory_inventory_id_seq1
    OWNED BY public.inventory.inventory_id;
