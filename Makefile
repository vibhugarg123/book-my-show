APP=book-my-show
APP_EXECUTABLE="./out/$(APP)"
GO=GO111MODULE=on go
# BUILD #############
clean: ##@build remove executable
	rm -f $(APP_EXECUTABLE)

compile: ##@build build the executable
	mkdir -p out/
	GO111MODULE=on go build -o $(APP_EXECUTABLE)

build: compile ##@build a fresh build

# TESTS #############
clean:
	$(GO) mod tidy -v

setup:
	$(GO) get -u golang.org/x/lint/golint

test: clean setup
	$(GO) test ./... -v -coverprofile=coverage.out.tmp -p 1
	cat coverage.out.tmp | grep -v "_mock.go" > coverage.out

# DB ################
mysql:
	docker stop full_db_mysql
	docker rm full_db_mysql
	docker run --name full_db_mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=ons_vg -d mysql

db-create:
	mysql -u root -h 127.0.0.1 -pons_vg -e "CREATE DATABASE IF NOT EXISTS BOOK_MY_SHOW;"

db-drop:
	mysql -u root -h 127.0.0.1 -pons_vg -e "DROP DATABASE IF EXISTS BOOK_MY_SHOW;"

db-migrate-up:
	migrate -path db/migration -database "mysql://root:ons_vg@tcp(127.0.0.1:3306)/BOOK_MY_SHOW" -verbose up

db-migrate-down:
	migrate -path db/migration -database "mysql://root:ons_vg@tcp(127.0.0.1:3306)/BOOK_MY_SHOW" -verbose down

db-up: db-create db-migrate-up
# Swagger ###########
check-swagger-install:
	which swagger || brew tap go-swagger/go-swagger || brew install go-swagger

generate-swagger: check-swagger-install
	swagger generate spec -o swagger-ui/swagger.yaml --scan-models

# Build ###########
docker-up:
	docker-compose down
	docker rmi book_my_show
	docker-compose up --build
	db-migrate-up

.PHONY: mysql db-create db-drop