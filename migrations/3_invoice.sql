CREATE TABLE
    IF NOT EXISTS invoice(
        id SERIAL PRIMARY KEY,
        invoice_number VARCHAR NOT NULL,
        customer_first_name VARCHAR NOT NULL,
        customer_last_name VARCHAR NOT NULL,
        amount_due NUMERIC(10, 2),
        amount_enclosed NUMERIC(10, 2),
        duedate DATE,
        is_locked BOOLEAN,
        is_traded BOOLEAN,
        investor_id INT,
        FOREIGN KEY (investor_id) REFERENCES investors (id)
    );