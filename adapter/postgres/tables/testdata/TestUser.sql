-- Query
INSERT
INTO
  users (name)
VALUES
  ($1)
RETURNING
  *;


-- Args
{
 "$1": "john appleseed"
}


-- Rows
{
 "ID": "67d9669a-69e6-48b3-a7be-c05b85b5638b",
 "Name": "john appleseed",
 "CreatedAt": "2023-06-24T05:01:10.011289Z",
 "UpdatedAt": "2023-06-24T05:01:10.011289Z"
}