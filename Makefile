.PHONY: pgTestSetup \
		pgTestStart \
		pgTestStop \
		test \
        install

pgTestSetup:
	@ docker help > /dev/null 2>&1 || (echo "Please install docker." && exit 1)
	docker run --name myapp-postgres-test -p 5432:5432 \
    -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=pass -e POSTGRES_DB=testdb \
    -d postgres:12.6 || echo "Postgres container for tests already set up"

pgTestStart:
	docker start myapp-postgres-test || (make pgTestSetup)

pgTestStop:
	docker stop myapp-postgres-test

install: pgTestSetup
	go mod download

test: pgTestStart
	go test ./...