up:
	docker-compose  --env-file .env up

down:
	docker-compose down

mocks:
	mockgen -source src/model/album_repository.go -destination src/service/mocks/mock_album_mysql_repository.go -package mocks
	mockgen -source src/model/http_repository.go -destination src/service/mocks/mock_http_repository.go -package mocks

test:
	go test ./...