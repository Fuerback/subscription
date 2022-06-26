CREATE TABLE account (
    id   STRING PRIMARY KEY,
    name STRING NOT NULL
);

CREATE TABLE product (
    id     STRING  PRIMARY KEY,
    name   STRING  NOT NULL
                   UNIQUE,
    period STRING  NOT NULL,
    price  DECIMAL NOT NULL,
    active BOOLEAN NOT NULL
                   DEFAULT (1) 
);


CREATE TABLE voucher (
    id         STRING PRIMARY KEY,
    code       STRING NOT NULL,
    type       STRING NOT NULL,
    status     STRING NOT NULL,
    product_id STRING REFERENCES product (id) 
                      NOT NULL
);


CREATE TABLE subscription (
    id         STRING PRIMARY KEY,
    startsAt   DATE   NOT NULL,
    endsAt     DATE   NOT NULL,
    status     STRING NOT NULL,
    voucher_id STRING REFERENCES voucher (id),
    account_id STRING REFERENCES account (id) 
                      NOT NULL,
    product_id STRING REFERENCES product (id) 
                      NOT NULL
);

CREATE TABLE payment (
    id              STRING  PRIMARY KEY,
    value           DECIMAL NOT NULL,
    subscription_id STRING  REFERENCES subscription (id) 
                            NOT NULL
);