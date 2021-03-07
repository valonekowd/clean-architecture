CREATE TABLE IF NOT EXISTS users (
   id BIGSERIAL	 PRIMARY KEY,
   email VARCHAR(300) UNIQUE NOT NULL,
   password BYTEA NOT NULL,
   first_name VARCHAR(50) NOT NULL,
   last_name VARCHAR(50) NOT NULL,
   gender CHAR(1) NOT NULL CHECK (gender IN ('M', 'F', 'O')),
   date_of_birth TIMESTAMP NOT NULL,
   created_at TIMESTAMP NOT NULL,
   updated_at TIMESTAMP NOT NULL,
   deleted_at TIMESTAMP
);