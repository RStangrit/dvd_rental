CREATE TABLE public.language (
    language_id BIGSERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    last_update TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_language_name ON public.language USING btree (name);
