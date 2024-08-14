CREATE TABLE public.balances (
    id bigserial NOT NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    user_id bigint,
    current_balance float(53) NOT NULL ,
    CONSTRAINT balances_pkey PRIMARY KEY (id),
    CONSTRAINT balances_user_ID_key UNIQUE (user_id)
);

CREATE INDEX idx_balances_user_ID ON public.balances USING btree (user_id);

CREATE TABLE public.operations (
     id bigserial NOT NULL,
     created_at timestamptz NULL,
     updated_at timestamptz NULL,
     user_id bigint,
     cost float(53) NOT NULL ,
     operation_type VARCHAR(64) NOT NULL,
     args JSONB,
     CONSTRAINT operations_pkey PRIMARY KEY (id)
);

CREATE INDEX idx_operations_user_id ON public.operations USING btree (user_id);

CREATE TABLE public.records (
    id bigserial NOT NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    operation_id bigint NOT NULL,
    user_id bigint NOT NULL,
    user_balance float(53) NOT NULL ,
    deleted bool NOT NULL ,
    operation_response text NOT NULL ,
    CONSTRAINT records_pkey PRIMARY KEY (id),
    CONSTRAINT fk_records_operations FOREIGN KEY (operation_id) REFERENCES public.operations(id)
);

CREATE INDEX idx_records_user_id ON public.records USING btree (user_id);
CREATE INDEX idx_records_operation_id ON public.records USING btree (operation_id);

GRANT ALL ON SCHEMA public TO public;
