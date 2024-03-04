CREATE TABLE client (  
    client_id              SERIAL PRIMARY KEY,
    client_created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    client_updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    client_deleted_at      TIMESTAMP,
    client_account_limit   INTEGER,
    client_account_balance INTEGER
);

CREATE TABLE transaction (
    transaction_id                    SERIAL PRIMARY KEY,
    transaction_created_at            TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    transaction_updated_at            TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    transaction_deleted_at            TIMESTAMP,
    transaction_value                 INTEGER NOT NULL,
    transaction_type                  VARCHAR(255) NOT NULL,
    transaction_description           VARCHAR(255) NOT NULL,
    transaction_client_id               INTEGER,
    FOREIGN KEY (transaction_client_id) REFERENCES client(client_id)
);

INSERT INTO client (client_account_limit, client_account_balance) VALUES (100000, 0);
INSERT INTO client (client_account_limit, client_account_balance) VALUES (80000, 0);
INSERT INTO client (client_account_limit, client_account_balance) VALUES (1000000, 0);
INSERT INTO client (client_account_limit, client_account_balance) VALUES (10000000, 0);
INSERT INTO client (client_account_limit, client_account_balance) VALUES (500000, 0);