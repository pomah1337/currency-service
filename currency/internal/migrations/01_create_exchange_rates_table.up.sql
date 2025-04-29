create table currency_exchange_rates
(
    id              serial primary key,
    base_currency   varchar        not null,
    target_currency varchar        not null,
    rate            numeric(10, 6) not null,
    update_date      date default now()
);