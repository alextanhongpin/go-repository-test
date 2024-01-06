-- Query
DELETE
  FROM users t
  WHERE t.id = '25c67e2b-fa1f-4af2-9705-9e9bec5c5df3'


-- Query Normalized
DELETE
  FROM users t
  WHERE t.id = $1


-- Args
$1: 25c67e2b-fa1f-4af2-9705-9e9bec5c5df3



-- Result
false
