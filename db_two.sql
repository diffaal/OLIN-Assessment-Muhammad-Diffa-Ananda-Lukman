CREATE EXTENSION IF NOT EXISTS dblink;

SELECT u.name, SUM(o.amount) AS amount , MAX(o.created_at) AS created_at
FROM users u 
INNER JOIN dblink('dbname=second user=postgres password=pass', 'SELECT id, user_id, amount, created_at FROM orders')
AS o(id INT, user_id INT, amount NUMERIC(10,2), created_at TIMESTAMP)
ON u.id = o.user_id
GROUP BY u.id;
