CREATE TABLE public.balances (
    id bigserial NOT NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    user_id int8,
    current_balance int NOT NULL ,
    CONSTRAINT balances_pkey PRIMARY KEY (id),
    CONSTRAINT balances_user_ID_key UNIQUE (user_id)
);

CREATE INDEX idx_balances_user_ID ON public.balances USING btree (user_id);

CREATE TABLE public.operations (
     id bigserial NOT NULL,
     created_at timestamptz NULL,
     updated_at timestamptz NULL,
     user_id int8,
     cost int NOT NULL ,
     operation_type VARCHAR(64) NOT NULL,
     args JSONB,
     CONSTRAINT operations_pkey PRIMARY KEY (id),
     CONSTRAINT operations_user_ID_key UNIQUE (user_id)
);

CREATE INDEX idx_operations_user_id ON public.operations USING btree (user_id);

CREATE TABLE public.records (
    id bigserial NOT NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    operation_id in8 NOT NULL,
    user_id in8 NOT NULL,
    amount int NOT NULL ,
    user_balance int NOT NULL ,
    deleted bool NOT NULL ,
    operation_response float(53) NOT NULL ,
    CONSTRAINT records_pkey PRIMARY KEY (id),
    CONSTRAINT fk_records_operations FOREIGN KEY (operation_id) REFERENCES public.operations(id),
    CONSTRAINT records_user_id_key UNIQUE (user_id)
);

CREATE INDEX idx_records_user_id ON public.records USING btree (user_id);
CREATE INDEX idx_records_operation_id ON public.records USING btree (operation_id);

GRANT ALL ON SCHEMA public TO public;
