-- Query
INSERT INTO users (name)
  VALUES ('John Appleseed') RETURNING *


-- Query Normalized
INSERT INTO users (name)
  VALUES ($1) RETURNING *


-- Args
$1: John Appleseed



-- Result
false
