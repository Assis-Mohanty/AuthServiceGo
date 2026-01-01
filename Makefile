MIGRATIONS_PATH="db/migrations"

migrate-up:
	./goose -dir ${MIGRATIONS_PATH} mysql "root:qqqq@tcp(127.0.0.1:3306)/usersgo" up