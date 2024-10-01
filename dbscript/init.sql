   CREATE TYPE gateway AS ENUM ('A', 'B');
   CREATE TYPE transactiontype AS ENUM ('DEPOSIT', 'WITHDRAWAL', 'DEPOSIT_REVERSAL', 'WITHDRAWAL_REVERSAL');
   
   CREATE TABLE IF NOT EXISTS client (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    gateway gateway NOT NULL,
    balance DECIMAL(10, 2) NOT NULL DEFAULT 0   
   );

   CREATE TABLE IF NOT EXISTS transaction (
    id SERIAL PRIMARY KEY,
    client_id INT NOT NULL,
    type transactiontype NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (client_id) REFERENCES client(id)
   );

   INSERT INTO client (name, gateway, balance) VALUES ('Alice', 'A', 100.00);
   INSERT INTO client (name, gateway, balance) VALUES ('Bob', 'B', 200.00);