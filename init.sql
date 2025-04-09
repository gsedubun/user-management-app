-- init.sql
CREATE TABLE IF NOT EXISTS "user" (
    id SERIAL PRIMARY KEY,
    firstname VARCHAR(50) NOT NULL,
    lastname VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    address TEXT
);

-- Optional: Insert some sample data
INSERT INTO "user" (firstname, lastname, email, address)
VALUES 
    ('John', 'Doe', 'john.doe@example.com', '123 Main St'),
    ('Jane', 'Smith', 'jane.smith@example.com', '456 Oak Ave')
ON CONFLICT (email) DO NOTHING;
