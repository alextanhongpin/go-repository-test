-- Query
DELETE
  FROM products p
  WHERE p.id = -1


-- Query Normalized
DELETE
  FROM products p
  WHERE p.id = $1


-- Args
$1: -1



-- Result
false
