CREATE TABLE "user" (  
    user_id              SERIAL PRIMARY KEY,
    user_created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_deleted_at      TIMESTAMP,
    user_account_limit   INTEGER,
    user_account_balance INTEGER
);

CREATE TABLE transaction (
    transaction_id                    SERIAL PRIMARY KEY,
    transaction_created_at            TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    transaction_updated_at            TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    transaction_deleted_at            TIMESTAMP,
    transaction_value                 INTEGER NOT NULL,
    transaction_type                  VARCHAR(255) NOT NULL,
    transaction_description           VARCHAR(255) NOT NULL,
    transaction_user_id               INTEGER,
    FOREIGN KEY (transaction_user_id) REFERENCES "user"(user_id)
);

INSERT INTO "user" (user_account_limit, user_account_balance) VALUES (100000, 0);
INSERT INTO "user" (user_account_limit, user_account_balance) VALUES (80000, 0);
INSERT INTO "user" (user_account_limit, user_account_balance) VALUES (1000000, 0);
INSERT INTO "user" (user_account_limit, user_account_balance) VALUES (10000000, 0);
INSERT INTO "user" (user_account_limit, user_account_balance) VALUES (500000, 0);