-- Create accounts table
CREATE TABLE IF NOT EXISTS accounts (
    Id VARCHAR (100) UNIQUE NOT NULL, 
    username VARCHAR (50) UNIQUE NOT NULL, 
    currency VARCHAR (20) NOT NULL, 
    balance VARCHAR (100) NOT NULL
);

-- Create payments table
CREATE TABLE IF NOT EXISTS payments (
    id VARCHAR(100) UNIQUE NOT NULL, 
    sender VARCHAR(100) NOT NULL, 
    reciever VARCHAR(100) NOT NULL, 
    amount NUMERIC NOT NULL, 
    initiated TIME NOT NULL, 
    completed TIME NOT NULL
);