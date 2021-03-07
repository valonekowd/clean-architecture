CREATE TABLE IF NOT EXISTS accounts (
   id BIGSERIAL PRIMARY KEY,
   name VARCHAR(100) UNIQUE NOT NULL,
   bank VARCHAR(10) NOT NULL CHECK (bank IN ('VCB', 'ACB', 'VIB')),
   user_id BIGSERIAL NOT NULL,
   created_at TIMESTAMP NOT NULL,
   updated_at TIMESTAMP NOT NULL,
   deleted_at TIMESTAMP,
   CONSTRAINT fk_user
      FOREIGN KEY(user_id) 
	      REFERENCES users(id)
);