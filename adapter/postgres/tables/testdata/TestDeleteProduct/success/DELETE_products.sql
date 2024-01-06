-- Query
DELETE
  FROM products p
  WHERE p.id = 3


-- Query Normalized
DELETE
  FROM products p
  WHERE p.id = $1


-- Args
$1: 3



-- Result
false
