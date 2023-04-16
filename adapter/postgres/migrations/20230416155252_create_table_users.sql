-- migrate:up
create table users (
	id uuid default gen_random_uuid(),
	name text not null,
	created_at timestamptz not null default current_timestamp,
	updated_at timestamptz not null default current_timestamp,
	primary key(id)
);

create trigger users_moddatetime
	before update on users
	for each row
	execute procedure moddatetime(updated_at);

-- migrate:down
drop table users;
