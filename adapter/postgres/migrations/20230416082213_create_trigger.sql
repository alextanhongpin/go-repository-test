create or replace trigger moddatetime_user
	before update on users
	for each row
	execute procedure moddatetime(updated_at);
