-- Query
SELECT p.id,
       p.name,
       p.description,
       p.user_id,
       p.created_at,
       p.updated_at
  FROM products p
  WHERE p.id = -1


-- Query Normalized
SELECT p.id,
       p.name,
       p.description,
       p.user_id,
       p.created_at,
       p.updated_at
  FROM products p
  WHERE p.id = $1


-- Args
$1: -1



-- Result
false
