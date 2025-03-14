CREATE TABLE public.film_category (
    film_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    last_update TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY (film_id, category_id),
    CONSTRAINT fk_film_films_categories FOREIGN KEY (film_id) REFERENCES public.film(film_id),
    CONSTRAINT fk_category_film_category FOREIGN KEY (category_id) REFERENCES public.category(category_id)
);

CREATE INDEX idx_film_category_film_id ON public.film_category USING btree (film_id);
CREATE INDEX idx_film_category_category_id ON public.film_category USING btree (category_id);
