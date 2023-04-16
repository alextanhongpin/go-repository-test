-- migrate:up
create extension citext;
create extension moddatetime;

-- migrate:down
drop extension citext;
drop extension moddatetime;
