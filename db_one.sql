SELECT u.name, SUM(o.amount) AS total_amount 
FROM users u 
INNER JOIN orders o 
ON u.id = o.user_id 
WHERE o.created_at >= '2022-01-01' 
GROUP BY u.name 
HAVING SUM(o.amount) > 1000;