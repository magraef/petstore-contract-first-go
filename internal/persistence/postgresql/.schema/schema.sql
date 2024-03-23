-- name: create-categories-table
-- Create the categories table
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Add ON CONFLICT clause for 'name' column
CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_category_name ON categories(name);

-- name: create-pets-table
-- Create the pets table
CREATE TABLE IF NOT EXISTS pets (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category_id BIGINT NOT NULL REFERENCES categories(id)
);

-- Add ON CONFLICT clause for 'name' column
CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_pet_name ON pets(name);


