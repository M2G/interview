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

/*

Liste les produits qui n'ont jamais été commandés (aucune ligne dans orders) :

SELECT * FROM products LEFT JOIN orders ON products.id = orders.product_id WHERE orders.id IS NULL
SELECT products.* FROM products LEFT JOIN orders ON products.id = orders.product_id WHERE orders.id IS NULL;

Pour chaque entité, calcule le nombre total de commandes passées par ses utilisateurs (une ligne par entité, avec son nom et le total) :

SELECT
    entities.id,
    entities.name,
    COUNT(orders.id) AS total_orders
FROM orders
JOIN users ON orders.user_id = users.id
JOIN entities ON users.entity_id = entities.id
GROUP BY entities.id, entities.name

Pour chaque catégorie de produit, trouve le produit le plus récemment commandé (celui dont la commande la plus récente est la plus tardive), avec le nom du produit et la date de cette commande :

SELECT DISTINCT ON (products.category)
    products.category,
    products.name,
    orders.ordered_at
FROM products
LEFT JOIN orders ON products.id = orders.product_id
ORDER BY products.category, orders.ordered_at DESC NULLS LAST;



 */