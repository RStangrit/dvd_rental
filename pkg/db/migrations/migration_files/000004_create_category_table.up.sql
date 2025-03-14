CREATE TABLE public.category (
    category_id BIGSERIAL PRIMARY KEY,
    name VARCHAR(25) NOT NULL,
    last_update TIMESTAMP DEFAULT now() NOT NULL,
    deleted_at TIMESTAMPTZ
);
