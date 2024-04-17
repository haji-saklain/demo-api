CREATE TABLE output_table (
    id SERIAL PRIMARY KEY,
    command VARCHAR(255),
    output TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
