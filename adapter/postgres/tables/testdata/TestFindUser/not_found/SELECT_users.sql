-- Query
SELECT t.id,
       t.name,
       t.created_at,
       t.updated_at
  FROM users t
  WHERE t.id = 'b64607df-f393-4300-bbf4-6c7d4e83f0d5'


-- Query Normalized
SELECT t.id,
       t.name,
       t.created_at,
       t.updated_at
  FROM users t
  WHERE t.id = $1


-- Args
$1: b64607df-f393-4300-bbf4-6c7d4e83f0d5



-- Result
false
