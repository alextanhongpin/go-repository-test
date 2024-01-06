-- Query
SELECT t.id,
       t.name,
       t.created_at,
       t.updated_at
  FROM users t
  WHERE t.id = 'd1874e6c-ee79-4ac3-8588-d17bd4521f74'


-- Query Normalized
SELECT t.id,
       t.name,
       t.created_at,
       t.updated_at
  FROM users t
  WHERE t.id = $1


-- Args
$1: d1874e6c-ee79-4ac3-8588-d17bd4521f74



-- Result
false
