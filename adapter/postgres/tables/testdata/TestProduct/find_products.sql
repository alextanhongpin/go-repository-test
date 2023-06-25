-- Query
SELECT
  p.id,
  p.name,
  p.description,
  p.user_id,
  p.created_at,
  p.updated_at
FROM
  products AS p
WHERE
  p.id = 1;


-- Query Normalized
SELECT
  p.id,
  p.name,
  p.description,
  p.user_id,
  p.created_at,
  p.updated_at
FROM
  products AS p
WHERE
  p.id = $1;


-- Args
{
 "$1": 1
}


-- Rows
{
 "ID": 1,
 "Name": "Rainbow Socks",
 "Description": "A rainbow colored socks",
 "UserID": "534941da-06ca-4d9a-88cf-5fc82d4943b5",
 "CreatedAt": "2023-06-25T15:26:47.87562Z",
 "UpdatedAt": "2023-06-25T15:26:47.87562Z",
 "User": null
}