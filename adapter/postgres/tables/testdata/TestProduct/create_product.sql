-- Query
INSERT INTO products (name, description, user_id)
  VALUES ('Rainbow Socks', 'A rainbow colored socks', '27069b9d-dafb-45ce-9c48-5800b3b5823d') RETURNING *


-- Query Normalized
INSERT INTO products (name, description, user_id)
  VALUES ($1, $2, $3) RETURNING *


-- Args
$1: Rainbow Socks
$2: A rainbow colored socks
$3: 27069b9d-dafb-45ce-9c48-5800b3b5823d



-- Result
ID: 1
Name: Rainbow Socks
Description: A rainbow colored socks
UserID: 27069b9d-dafb-45ce-9c48-5800b3b5823d
CreatedAt: "2023-07-01T13:13:59.223377Z"
UpdatedAt: "2023-07-01T13:13:59.223377Z"
User: null
