CREATE TABLE IF NOT EXISTS transactions (
   id BIGSERIAL PRIMARY KEY,
   type VARCHAR(8) NOT NULL CHECK (type IN ('withdraw', 'deposit')),
   amount DOUBLE PRECISION NOT NULL,
   account_id BIGSERIAL NOT NULL,
   created_at TIMESTAMP NOT NULL,
   CONSTRAINT fk_account
      FOREIGN KEY(account_id) 
	      REFERENCES accounts(id)
);