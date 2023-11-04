CREATE TABLE IF NOT EXISTS accounts (
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
	garage_number TEXT UNIQUE,
	first_name    TEXT,
	last_name     TEXT,
	phone_number  TEXT,
	address_line  TEXT
);

CREATE TABLE IF NOT EXISTS bills (
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
	account_id INTEGER,
	year_id    INTEGER ,
	Payed      REAL 
);


CREATE TABLE IF NOT EXISTS years (
	id    INTEGER  PRIMARY KEY AUTOINCREMENT,
	price INTEGER
);

CREATE UNIQUE INDEX idx_accounts_garage_number ON accounts (garage_number);
CREATE  INDEX idx_accounts_last_name ON accounts (last_name);


