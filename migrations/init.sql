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
        balance NUMERIC(10, 2),
        email VARCHAR,
        password VARCHAR,
        is_active BOOLEAN,
        is_issuer BOOLEAN
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
        asking_price NUMERIC(10, 2),
        created_on TIMESTAMP,
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
        timestamp TIMESTAMP,
        is_approved BOOLEAN,
        invoice_id INT NOT NULL,
        investor_id INT,
        status SMALLINT,
        FOREIGN KEY (invoice_id) REFERENCES invoice (id),
        FOREIGN KEY (investor_id) REFERENCES investors (id)
    );

INSERT INTO
    investors(
        first_name,
        last_name,
        balance,
        email,
        password,
        is_active,
        is_issuer
    )
VALUES (
        'Jane',
        'Daves',
        10000.0,
        'jane123@gmail.com',
        '$2a$10$p.d6hCOF16/jvAhmt5yBK.5S2piteKX0KPHt3VQqFGWoIn1cESdp.',
        true,
        true
    );

INSERT INTO
    investors(
        first_name,
        last_name,
        balance,
        email,
        password,
        is_active,
        is_issuer
    )
VALUES (
        'Will',
        'Johnson',
        5000.0,
        'will123@gmail.com',
        '$2a$10$p.d6hCOF16/jvAhmt5yBK.5S2piteKX0KPHt3VQqFGWoIn1cESdp.',
        true,
        true
    );

INSERT INTO
    investors(
        first_name,
        last_name,
        balance,
        email,
        password,
        is_active,
        is_issuer
    )
VALUES (
        'Robert',
        'David',
        2000.0,
        'robert123@gmail.com',
        '$2a$10$p.d6hCOF16/jvAhmt5yBK.5S2piteKX0KPHt3VQqFGWoIn1cESdp.',
        true,
        false
    );

INSERT INTO
    investors(
        first_name,
        last_name,
        balance,
        email,
        password,
        is_active,
        is_issuer
    )
VALUES (
        'Lisa',
        'Nancy',
        6000.0,
        'lisa123@gmail.com',
        '$2a$10$p.d6hCOF16/jvAhmt5yBK.5S2piteKX0KPHt3VQqFGWoIn1cESdp.',
        true,
        true
    );

INSERT INTO
    investors(
        first_name,
        last_name,
        balance,
        email,
        password,
        is_active,
        is_issuer
    )
VALUES (
        'Sara',
        'Lee',
        1000.0,
        'sara123@gmail.com',
        '$2a$10$p.d6hCOF16/jvAhmt5yBK.5S2piteKX0KPHt3VQqFGWoIn1cESdp.',
        true,
        false
    );

INSERT INTO issuers(company_name,investor_id ) VALUES ('test1', 1);

INSERT INTO issuers(company_name, investor_id) VALUES ('test2',2 );

INSERT INTO issuers(company_name, investor_id) VALUES ('test3',4 );