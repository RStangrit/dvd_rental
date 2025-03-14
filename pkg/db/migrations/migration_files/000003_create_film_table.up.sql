CREATE TABLE public.film (
    film_id BIGSERIAL PRIMARY KEY,
    title character varying(255) NOT NULL,
    description text,
    release_year integer,
    language_id integer NOT NULL,
    rental_duration smallint DEFAULT 3 NOT NULL,
    rental_rate numeric(4,2) DEFAULT 4.99 NOT NULL,
    length smallint,
    replacement_cost numeric(5,2) DEFAULT 19.99 NOT NULL,
    rating text,
    last_update timestamp without time zone NOT NULL,
    special_features text[],
    fulltext tsvector,
    deleted_at timestamp with time zone
);

CREATE INDEX idx_film_release_year ON public.film USING btree (release_year);
CREATE INDEX idx_film_title ON public.film USING btree (title);

