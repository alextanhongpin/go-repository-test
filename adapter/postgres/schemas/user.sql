create table users (
	id uuid default gen_random_uuid(),
	name text not null,

	created_at timestamptz not null default current_timestamp,
	updated_at timestamptz not null default current_timestamp,
	primary key (id),
	unique (name)
);
