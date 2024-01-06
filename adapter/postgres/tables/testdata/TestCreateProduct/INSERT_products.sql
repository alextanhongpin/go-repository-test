-- Query
INSERT INTO products (name, description, user_id)
  VALUES ('table', 'a wooden table', '80b44f42-22d0-4980-b204-bcd600324374') RETURNING *


-- Query Normalized
INSERT INTO products (name, description, user_id)
  VALUES ($1, $2, $3) RETURNING *


-- Args
$1: table
$2: a wooden table
$3: 80b44f42-22d0-4980-b204-bcd600324374



-- Result
false
