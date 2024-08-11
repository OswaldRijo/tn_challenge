CREATE TABLE public.users (
    id bigserial NOT NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    status varchar(16),
    username varchar(64) NOT NULL,
    password text NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_user_name_key UNIQUE (username)
);

CREATE INDEX idx_users_username ON public.users USING btree (username);

GRANT ALL ON SCHEMA public TO public;
