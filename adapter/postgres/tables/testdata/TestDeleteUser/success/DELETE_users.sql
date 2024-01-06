-- Query
DELETE
  FROM users t
  WHERE t.id = 'd5717938-af86-4f5c-a826-774704583607'


-- Query Normalized
DELETE
  FROM users t
  WHERE t.id = $1


-- Args
$1: d5717938-af86-4f5c-a826-774704583607



-- Result
false
