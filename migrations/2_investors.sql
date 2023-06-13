CREATE TABLE
    IF NOT EXISTS investors(
        id SERIAL PRIMARY KEY,
        first_name VARCHAR NOT NULL,
        last_name VARCHAR,
        balance NUMERIC(10, 2)
    );

INSERT INTO
    investors(first_name, last_name, balance)
VALUES ('Jane', 'Daves', 10000.0);

INSERT INTO
    investors(first_name, last_name, balance)
VALUES ('Will', 'Johnson', 5000.0);

INSERT INTO
    investors(first_name, last_name, balance)
VALUES ('Robert', 'David', 2000.0);

INSERT INTO
    investors(first_name, last_name, balance)
VALUES ('Lisa', 'Nancy', 6000.0);

INSERT INTO
    investors(first_name, last_name, balance)
VALUES ('Sara', 'Lee', 1000.0);