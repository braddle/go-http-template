test: docker_build_test
	docker-compose up -d
	docker-compose exec http go test ./...
	docker-compose down

unit_test:
	go test `go list ./... | grep -v e2e_test`

docker_build:
	docker build . -t template

docker_build_test:
	docker build . -t template_test --target=test

docker_run:
	docker run template