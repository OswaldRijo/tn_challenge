CREATE SEQUENCE public.users_id_seq
    INCREMENT BY 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
	CACHE 1
	NO CYCLE;

-- public.users definition

CREATE TABLE public.users (
    id bigserial NOT NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    username varchar(64) NOT NULL,
    password text NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_user_name_key UNIQUE (username),
);

CREATE INDEX idx_users_username ON public.users USING btree (username);

GRANT ALL ON SCHEMA public TO public;
