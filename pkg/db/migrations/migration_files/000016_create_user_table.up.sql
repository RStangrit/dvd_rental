CREATE SEQUENCE public.user_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE public."user" (
    user_id BIGINT PRIMARY KEY DEFAULT nextval('public.user_user_id_seq'),
    email VARCHAR(45) NOT NULL UNIQUE,
    password VARCHAR(60) NOT NULL,
    last_update TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

ALTER SEQUENCE public.user_user_id_seq
    OWNED BY public."user".user_id;