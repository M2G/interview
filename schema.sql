CREATE TABLE entities (
                          id SERIAL PRIMARY KEY,
                          name TEXT NOT NULL
);

CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       entity_id INTEGER REFERENCES entities(id),
                       email TEXT UNIQUE NOT NULL
);

CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          name TEXT NOT NULL,
                          category TEXT NOT NULL,
                          metadata JSONB,
                          created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        user_id INTEGER REFERENCES users(id),
                        product_id INTEGER REFERENCES products(id),
                        status TEXT NOT NULL DEFAULT 'pending',
                        ordered_at TIMESTAMPTZ DEFAULT now()
);