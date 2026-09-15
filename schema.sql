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

Pour chaque utilisateur, affiche ses 2 commandes les plus récentes (peu importe le produit), avec le nom du produit et la date. Un utilisateur avec une seule commande n'affiche qu'une ligne ; un utilisateur sans commande n'apparaît pas :

WITH RankedOrders AS (
    SELECT
        orders.id,
        orders.user_id,
        orders.product_id,
        orders.ordered_at,
        ROW_NUMBER() OVER (PARTITION BY orders.user_id ORDER BY orders.ordered_at DESC) as rn
    FROM orders
)

SELECT *
FROM RankedOrders
JOIN products ON products.id = RankedOrders.product_id
WHERE rn <= 2
ORDER BY user_id, rn;


UPDATE products SET created_at = '2026-01-01 00:00:00'
WHERE name IN ('Poste de travail portable', 'VM Cloud Standard');

Pour chaque catégorie de produit, classe les produits par date de création (created_at) du plus ancien au plus récent, en affichant leur rang. Si deux produits ont exactement la même date de création dans une catégorie, ils doivent recevoir le même rang, et le rang suivant doit sauter en conséquence (ex : 1, 2, 2, 4 - pas 1, 2, 2, 3).

SELECT name, category, created_at, RANK () OVER (
    PARTITION BY category
    ORDER BY created_at
    ) as rank_number,
    DENSE_RANK() OVER (PARTITION BY  category ORDER BY created_at) as dense_rank_number
FROM products;

 */