-- Query
INSERT
INTO
  products (name, description, user_id)
VALUES
  (
    'Rainbow Socks',
    'A rainbow colored socks',
    '534941da-06ca-4d9a-88cf-5fc82d4943b5'
  )
RETURNING
  *;


-- Query Normalized
INSERT
INTO
  products (name, description, user_id)
VALUES
  ($1, $2, $3)
RETURNING
  *;


-- Args
{
 "$1": "Rainbow Socks",
 "$2": "A rainbow colored socks",
 "$3": "534941da-06ca-4d9a-88cf-5fc82d4943b5"
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