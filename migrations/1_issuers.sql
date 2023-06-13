CREATE TABLE
    IF NOT EXISTS issuers(
        id SERIAL PRIMARY KEY,
        first_name VARCHAR NOT NULL,
        last_name VARCHAR,
        balance NUMERIC(10, 2)
    );

INSERT INTO
    issuers(first_name, last_name, balance)
VALUES ('John', 'Doe', 5000.0);

INSERT INTO
    issuers(first_name, last_name, balance)
VALUES ('Alice', 'Smith', 7000.0);

INSERT INTO
    issuers(first_name, last_name, balance)
VALUES ('Bob', 'Carlton', 1000.0);