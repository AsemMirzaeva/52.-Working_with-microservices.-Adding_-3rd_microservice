
CREATE TABLE if not exists sales_transactions (
    transaction_id VARCHAR PRIMARY KEY,
    product_id VARCHAR,
    quantity INT,
    price DECIMAL,
    timestamp BIGINT
);

create table if not exists sales_summary (
    total_amount decimal,
    total_transactions DECIMAL
);