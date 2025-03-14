CREATE TABLE public.film_actor (
    actor_id INTEGER NOT NULL,
    film_id INTEGER NOT NULL,
    last_update TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY (actor_id, film_id),
    CONSTRAINT fk_actor_actor_films FOREIGN KEY (actor_id) REFERENCES public.actor(actor_id),
    CONSTRAINT fk_film_film_actors FOREIGN KEY (film_id) REFERENCES public.film(film_id)
);

CREATE INDEX idx_film_actor_actor_id ON public.film_actor (actor_id);
CREATE INDEX idx_film_actor_film_id ON public.film_actor (film_id);

