export DBURL="postgres://postgres:password@localhost:5432/postgres?sslmode=disable" 
migrate -database ${DBURL} -path ./database/migration up