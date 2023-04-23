-- migrate:up
create table products (
	-- PK.
	id bigint generated always as identity,

	-- Attributes.
	name text not null check (length(name) > 3),
	description text not null check (length(description) > 3),
	user_id uuid not null,

	-- Timestamps.
	created_at timestamptz not null default current_timestamp,
	updated_at timestamptz not null default current_timestamp,

	-- Constraints.
	primary key (id),
	foreign key (user_id) references users(id),
	unique (name)
);

create trigger products_moddatetime
	before update on products
	for each row
	execute procedure moddatetime(updated_at);


-- migrate:down
drop table products;
