-- Query
INSERT INTO users (name)
  VALUES ('john appleseed') RETURNING *


-- Query Normalized
INSERT INTO users (name)
  VALUES ($1) RETURNING *


-- Args
$1: john appleseed



-- Result
ID: 385e26fe-0ac2-4738-b3d8-9614ec1f76e6
Name: john appleseed
CreatedAt: "2023-07-01T13:13:59.764471Z"
UpdatedAt: "2023-07-01T13:13:59.764471Z"
