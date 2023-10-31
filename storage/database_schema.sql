create table transactions
(
    id             bigserial
        primary key,
    created_at     timestamp with time zone,
    updated_at     timestamp with time zone,
    deleted_at     timestamp with time zone,
    transaction_id bigint                   not null,
    account_id     bigint                   not null,
    date           timestamp with time zone not null,
    amount         numeric(20, 8)           not null
);

alter table transactions
    owner to postgres;

create index idx_transactions_deleted_at
    on transactions (deleted_at);

create table accounts
(
    id         bigserial
        primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name       text not null,
    surname    text not null,
    email      text not null
);

alter table accounts
    owner to postgres;

create index idx_accounts_deleted_at
    on accounts (deleted_at);
