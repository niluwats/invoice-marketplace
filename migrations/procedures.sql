--SAVE INVESTOR

CREATE OR REPLACE PROCEDURE SAVE_INVESTOR(FIRST_NAME 
VARCHAR, LAST_NAME VARCHAR, BALANCE NUMERIC(10, 2)
, EMAIL VARCHAR, PASSWORD VARCHAR, STATUS BOOLEAN, 
IS_ISSUER BOOLEAN) LANGUAGE PLPGSQL AS 
	$$ begin
	insert into
	    investors (
	        first_name,
	        last_name,
	        balance,
	        email,
	        password,
	        status,
	        is_issuer
	    )
	values (
	        first_name,
	        last_name,
	        balance,
	        email,
	        password,
	        status,
	        is_issuer
	    );
	end 
$ 

$;

--GET INVESTOR BY ID

CREATE OR REPLACE FUNCTION GET_INVESTOR_BY_ID(INVESTOR_ID 
INTEGER) RETURNS TABLE(ID INTEGER, FIRST_NAME VARCHAR
, LAST_NAME VARCHAR, BALANCE NUMERIC(10, 2), STATUS 
BOOLEAN) LANGUAGE PLPGSQL AS 
	$$ begin
	return query
	select
	    investors.id,
	    investors.first_name,
	    investors.last_name,
	    investors.balance,
	    investors.status
	from investors
	where investors.id =
INVESTOR_ID; 

end $$;

-- GET ALL INVESTORS

CREATE OR REPLACE FUNCTION GET_INVESTORS() RETURNS 
TABLE(ID INTEGER, FIRST_NAME VARCHAR, LAST_NAME VARCHAR
, BALANCE NUMERIC(10, 2), STATUS BOOLEAN) LANGUAGE 
PLPGSQL AS 
	$$ begin
	return query
	select
	    investors.id,
	    investors.first_name,
	    investors.last_name,
	    investors.balance,
	    investors.status
	from investors
	where investors.status =
TRUE; 

end $$;

-- GET INVESTOR BY EMAIL

CREATE OR REPLACE FUNCTION GET_INVESTOR_BY_EMAIL(INVESTOR_EMAIL 
VARCHAR) RETURNS TABLE(ID INTEGER, PASSWORD VARCHAR
, IS_ISSUER BOOLEAN) LANGUAGE PLPGSQL AS 
	$$ begin
	return query
	select
	    investors.id,
	    investors.password,
	    investors.is_issuer
	from investors
	where investors.email =
INVESTOR_EMAIL; 

end $$;

-- GET ISSUER BY ID

CREATE OR REPLACE FUNCTION GET_ISSUER_BY_ID(ISSUER_ID 
INTEGER) RETURNS TABLE(ID INTEGER, COMPANY_NAME VARCHAR
, BALANCE NUMERIC(10, 2)) LANGUAGE PLPGSQL AS 
	$$ begin
	return query
	select
	    issuers.id,
	    issuers.company_name,
	    investors.balance
	from issuers
	    inner join investors on issuers.investor_id = investors.id
	where issuers.id =
ISSUER_ID; 

end $$;

-- GET ALL ISSUERS

CREATE OR REPLACE FUNCTION GET_ISSUERS() RETURNS TABLE
(ID INTEGER, COMPANY_NAME VARCHAR, BALANCE NUMERIC
(10, 2)) LANGUAGE PLPGSQL AS 
	$$ begin
	return query
	SELECT
	    issuers.id,
	    issuers.company_name,
	    investors.balance
	FROM issuers
	    INNER JOIN investors ON issuers.investor_id = investors.i
ID; 

end $$;

-- SAVE INVESTOR

CREATE OR REPLACE PROCEDURE SAVE_INVOICE(INVOICE_NUMBER 
VARCHAR, AMOUNT_DUE NUMERIC(10, 2), ASKING_PRICE 
	numeric(10, 2),
	duedate date,
	created_on timestamp,
	issuer_id integer
	) LANGUAGE PLPGSQL AS $$ declare row_id INTEGER;
	begin
	insert into
	    invoice (
	        invoice_number,
	        amount_due,
	        asking_price,
	        duedate,
	        created_on,
	        issuer_id
	    )
	values (
	        invoice_number,
	        amount_due,
	        asking_price,
	        duedate,
	        created_on,
	        issuer_id,
	        false,
	        false
	    ) RETURNING id INTO row_id;
	end 
$ 

$;

--SAVE INVOICE

CREATE OR REPLACE FUNCTION SAVE_INVOICE(INVOICE_NUMBER 
VARCHAR, AMOUNT_DUE NUMERIC(10, 2), ASKING_PRICE 
	numeric(10, 2),
	created_on timestamp,
	duedate date,
	is_locked boolean,
	is_traded boolean,
	issuer_id integer
	) RETURNS integer LANGUAGE PLPGSQL AS $$ declare new_id INTEGER;
	begin
	insert into
	    invoice (
	        invoice_number,
	        amount_due,
	        asking_price,
	        duedate,
	        created_on,
	        issuer_id,
	        is_locked,
	        is_traded
	    )
	values (
	        invoice_number,
	        amount_due,
	        asking_price,
	        duedate,
	        created_on,
	        issuer_id,
	        is_locked,
	        is_traded
	    ) RETURNING id INTO new_id;
	RETURN 
NEW_ID; 

end $$;

-- GET INVOICE BY ID

CREATE OR REPLACE FUNCTION GET_INVOICE_BY_ID(INV_ID 
INTEGER) RETURNS TABLE(ID INTEGER, INVOICE_NUMBER 
VARCHAR, AMOUNT_DUE NUMERIC(10, 2), ASKING_PRICE 
	NUMERIC(10, 2),
	created_on TIMESTAMP,
	duedate DATE,
	is_locked BOOLEAN,
	is_traded BOOLEAN,
	issuer_id INTEGER,
	investors INTEGER []
	) LANGUAGE PLPGSQL AS $$ begin
	return query
	select
	    invoice.id,
	    invoice.invoice_number,
	    invoice.amount_due,
	    invoice.asking_price,
	    invoice.created_on,
	    invoice.duedate,
	    invoice.is_locked,
	    invoice.is_traded,
	    invoice.issuer_id,
	    CASE
	        WHEN invoice.is_traded = true THEN ARRAY_AGG (bids.investor_id)
	        ELSE NULL
	    END AS investors
	FROM invoice
	    LEFT JOIN bids ON invoice.id = bids.invoice_id
	WHERE invoice.id = INV_ID
	GROUP BY invoice.i
ID; 

end $$;

--GET ALL INVOICES

CREATE OR REPLACE FUNCTION GET_INVOICES() RETURNS TABLE
(ID INTEGER, INVOICE_NUMBER VARCHAR, ASKING_PRICE 
	NUMERIC(10, 2),
	created_on timestamp,
	is_locked boolean,
	is_traded boolean,
	issuer_id integer
	) LANGUAGE PLPGSQL AS $$ begin
	return query
	SELECT
	    invoice.id,
	    invoice.invoice_number,
	    invoice.asking_price,
	    invoice.created_on,
	    invoice.is_locked,
	    invoice.is_traded,
	    invoice.issuer_id
	FROM invoice
	order by invoice.i
ID; 

end $$;

--GET TOTAL INVESTMENT OF INVOICE

CREATE OR REPLACE FUNCTION GET_TOTAL_INVESTMENT(INV_ID 
INTEGER) RETURNS TABLE(AMOUNT NUMERIC(10, 2)) LANGUAGE 
PLPGSQL AS 
	$$ begin
	return query
	select
	    COALESCE(SUM(bid_amount), 0)
	FROM bids
	WHERE
	    invoice_id = INV_ID
	    AND status =
1; 

end $$;

--GET INVOICE BY INVOICE NUMBER

CREATE OR REPLACE FUNCTION GET_INVOICE_BY_INVOICENUMBER
(INV_NUMBER VARCHAR) RETURNS TABLE(INVOICE_NUMBER 
VARCHAR) LANGUAGE PLPGSQL AS 
	$$ begin
	return query
	select
	    invoice.invoice_number
	FROM invoice
	WHERE invoice.invoice_number =
INV_NUMBER; 

end $$;

--SAVE BID

CREATE OR REPLACE FUNCTION SAVE_BID(INVOICE_ID INTEGER
, BID_AMOUNT NUMERIC(10, 2), CREATED_AT TIMESTAMP, 
INVESTOR_ID INTEGER) RETURNS INTEGER LANGUAGE PLPGSQL 
AS 
	$$ DECLARE row_id INTEGER;
	BEGIN
	INSERT INTO
	    bids (
	        bid_amount,
	        created_at,
	        is_approved,
	        invoice_id,
	        investor_id,
	        status
	    )
	values (
	        bid_amount,
	        created_at,
	        false,
	        invoice_id,
	        investor_id,
	        1
	    ) returning id into row_id;
	RETURN 
ROW_ID; 

end $$;

--UPDATE INVESTOR BALANCE

CREATE OR REPLACE PROCEDURE UPDATE_INVESTOR_BALANCE
(BID_AMOUNT NUMERIC(10, 2), INVESTOR_ID INTEGER) LANGUAGE 
PLPGSQL AS 
	$$ BEGIN
	UPDATE investors
	SET
	    balance = balance - bid_amount
	where id = investor_id;
	end 
$ 

$;

--UPDATE INVOICE LOCK STATUS

CREATE OR REPLACE PROCEDURE UPDATE_INVOICE_STATUS(INV_ID 
INTEGER, STATUS BOOLEAN) LANGUAGE PLPGSQL AS 
	$$ BEGIN
	UPDATE invoice
	SET is_locked = STATUS
	where id = INV_ID;
	end 
$ 

$;

--GET ALL BIDS BY INVOICE

CREATE OR REPLACE FUNCTION GET_ALL_BIDS_BY_INVOICE(
INV_ID INTEGER) RETURNS TABLE(ID INTEGER, BID_AMOUNT 
NUMERIC(10, 2), CREATED_AT TIMESTAMP, IS_APPROVED 
BOOLEAN, INVOICE_ID INTEGER, INVESTOR_ID INTEGER, 
STATUS SMALLINT) LANGUAGE PLPGSQL AS 
	$$ begin
	return query
	select
	    bids.id,
	    bids.bid_amount,
	    bids.created_at,
	    bids.is_approved,
	    bids.invoice_id,
	    bids.investor_id,
	    bids.status
	from bids
	where
	    bids.invoice_id = INV_ID
	    and bids.status = 1
	order by bids.i
ID; 

end $$;

--GET BID BY ID

CREATE OR REPLACE FUNCTION GET_BID_BY_ID(BID_ID INTEGER
) RETURNS TABLE(ID INTEGER, BID_AMOUNT NUMERIC(10, 
2), CREATED_AT TIMESTAMP, IS_APPROVED BOOLEAN, INVOICE_ID 
INTEGER, INVESTOR_ID INTEGER, STATUS SMALLINT) LANGUAGE 
PLPGSQL AS 
	$$ begin
	return query
	select
	    bids.id,
	    bids.bid_amount,
	    bids.created_at,
	    bids.is_approved,
	    bids.invoice_id,
	    bids.investor_id,
	    bids.status
	from bids
	where bids.id =
BID_ID; 

end $$;

--UPDATE BID APPROVAL STATUS

CREATE OR REPLACE PROCEDURE UPDATE_BID_APPROVAL_STATUS
(INV_ID INTEGER) LANGUAGE PLPGSQL AS 
	$$ BEGIN
	UPDATE bids
	SET is_approved = true
	WHERE invoice_id = INV_ID;
	end 
$ 

$;

--UPDATE ISSUER BALANCE

CREATE OR REPLACE PROCEDURE UPDATE_ISSUER_BALANCE(AMOUNT 
NUMERIC(10, 2), ISSUER_ID INTEGER) LANGUAGE PLPGSQL 
AS 
	$$ BEGIN
	UPDATE investors
	SET balance = balance + amount
	FROM issuers
	WHERE
	    issuers.investor_id = investors.id
	    and issuers.id = ISSUER_ID;
	end 
$ 

$;

--UPDATE INVOICE TRADE STATUS

CREATE OR REPLACE PROCEDURE UPDATE_INVOICE_TRADED_STATUS
(INV_ID INTEGER) LANGUAGE PLPGSQL AS 
	$$ BEGIN UPDATE invoice SET is_traded = true WHERE id = INV_ID;
	end 
$ 

$;

--UPDATE BID STATUS

CREATE OR REPLACE PROCEDURE UPDATE_BID_STATUS(INV_ID 
INTEGER) LANGUAGE PLPGSQL AS 
	$$ BEGIN UPDATE bids SET status = 0 where invoice_id = INV_ID;
	end 
$ 

$;

--UPDATE ALL INVESTORS BALANCES BY INVOICE ID

CREATE OR REPLACE PROCEDURE UPDATE_ALL_INVESTORS_BALANCES
(INV_ID INTEGER) LANGUAGE PLPGSQL AS 
	$$ BEGIN
	UPDATE investors
	SET
	    balance = investors.balance + bids.bid_amount
	FROM bids
	WHERE
	    bids.investor_id = investors.id
	    AND bids.invoice_id = INV_ID;
	end 
$ 

$;