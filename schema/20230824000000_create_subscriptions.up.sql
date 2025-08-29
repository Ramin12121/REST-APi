CREATE TABLE subscriptions (
    id VARCHAR(36) PRIMARY KEY,
    service_name VARCHAR(255) NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    start_date VARCHAR(7) NOT NULL,
    end_date VARCHAR(7) NULL
);