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