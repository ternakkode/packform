CREATE TABLE IF NOT EXISTS companies (
    id BIGSERIAL,
    company_name varchar UNIQUE,

    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS customers (
    id BIGSERIAL,
    user_id varchar UNIQUE,
    login varchar,
    password BYTEA,
    name varchar,
    company_id bigint,
    credit_cards varchar[] DEFAULT array[]::varchar[],
    
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS deliveries (
    id BIGSERIAL,
    order_item_id bigint,
    delivered_quantity int,

    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS order_items (
    id BIGSERIAL,
    order_id bigint,
    price_per_unit numeric(24, 3),
    quantity bigint,
    product varchar,

    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS orders (
    id BIGSERIAL,
    created_at timestamptz,
    order_name varchar UNIQUE,
    customer_user_id varchar,

    PRIMARY KEY(id)
);