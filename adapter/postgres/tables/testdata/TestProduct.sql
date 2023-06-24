-- Query
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
 "$3": "3135eb66-05d9-4017-886e-d32bfe0f545e"
}


-- Rows
{
 "ID": 1,
 "Name": "Rainbow Socks",
 "Description": "A rainbow colored socks",
 "UserID": "3135eb66-05d9-4017-886e-d32bfe0f545e",
 "CreatedAt": "2023-06-24T04:58:24.574653Z",
 "UpdatedAt": "2023-06-24T04:58:24.574653Z",
 "User": null
}