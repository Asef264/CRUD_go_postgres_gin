package constants

const (
	CREATE_USER_TABLE = `
	CREATE TABLE IF NOT EXISTS users(
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		name text not null,
		username text not null
	);
	`

	INSERT_USER = `insert into users (name, username)
	values($1,$2)
	`

	DELETE_USER = `delete from users where id=$1
	`
	UPDATE_USER = `UPDATE users SET name=$1, username=$2 WHERE id=$3;`

	GET_USER = `SELECT * from users where id=$1;`

	GET_USERS = `select * from users;`
)
