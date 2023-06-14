SELECT
    'CREATE DATABASE invoice_marketplace'
WHERE NOT EXISTS (
        SELECT
        FROM pg_database
        WHERE
            datname = 'invoice_marketplace'
    );

CREATE TABLE
    IF NOT EXISTS investors(
        id SERIAL PRIMARY KEY,
        first_name VARCHAR NOT NULL,
        last_name VARCHAR,
        balance NUMERIC(10, 2)
    );

CREATE TABLE
    IF NOT EXISTS issuers(
        id SERIAL PRIMARY KEY,
        company_name VARCHAR NOT NULL,
        investor_id INT,
        FOREIGN KEY (investor_id) REFERENCES investors (id)
    );

CREATE TABLE
    IF NOT EXISTS invoice(
        id SERIAL PRIMARY KEY,
        invoice_number VARCHAR NOT NULL,
        amount_due NUMERIC(10, 2),
        amount_enclosed NUMERIC(10, 2),
        created_on DATE,
        duedate DATE,
        is_locked BOOLEAN,
        is_traded BOOLEAN,
        issuer_id INT,
        FOREIGN KEY (issuer_id) REFERENCES issuers (id)
    );

CREATE TABLE
    IF NOT EXISTS bids(
        id SERIAL PRIMARY KEY,
        bid_amount NUMERIC(10, 2) NOT NULL,
        timestamp DATE,
        is_approved BOOLEAN,
        invoice_id INT NOT NULL,
        investor_id INT,
        FOREIGN KEY (invoice_id) REFERENCES invoice (id),
        FOREIGN KEY (investor_id) REFERENCES investors (id)
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

INSERT INTO issuers(company_name,investor_id ) VALUES ('test1', 1);

INSERT INTO issuers(company_name, investor_id) VALUES ('test2',2 );

INSERT INTO issuers(company_name, investor_id) VALUES ('test3',4 );