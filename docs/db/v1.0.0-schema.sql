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
    product_id STRING NOT NULL,
    FOREIGN KEY(product_id) REFERENCES product(id)
);


CREATE TABLE subscription (
    id         STRING PRIMARY KEY,
    starts_at   DATE   NOT NULL,
    ends_at     DATE   NOT NULL,
    status     STRING NOT NULL,
    voucher_id STRING,
    account_id STRING NOT NULL,
    product_id STRING NOT NULL,
    FOREIGN KEY(voucher_id) REFERENCES voucher(id),
    FOREIGN KEY(account_id) REFERENCES account(id),
    FOREIGN KEY(product_id) REFERENCES product(id)
);

CREATE TABLE payment (
    id              STRING  PRIMARY KEY,
    value           DECIMAL NOT NULL,
    subscription_id STRING  NOT NULL,
    FOREIGN KEY(subscription_id) REFERENCES subscription(id)
);