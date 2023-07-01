-- Query
SELECT p.id,
       p.name,
       p.description,
       p.user_id,
       p.created_at,
       p.updated_at
  FROM products p
  WHERE p.id = 1


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
$1: 1



-- Result
ID: 1
Name: Rainbow Socks
Description: A rainbow colored socks
UserID: 27069b9d-dafb-45ce-9c48-5800b3b5823d
CreatedAt: "2023-07-01T13:13:59.223377Z"
UpdatedAt: "2023-07-01T13:13:59.223377Z"
User: null
