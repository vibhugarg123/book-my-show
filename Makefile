APP=book-my-show
APP_EXECUTABLE="./out/$(APP)"

# BUILD #############
clean: ##@build remove executable
	rm -f $(APP_EXECUTABLE)

compile: ##@build build the executable
	mkdir -p out/
	GO111MODULE=on go build -o $(APP_EXECUTABLE)

build: compile ##@build a fresh build

# DB ################
mysql:
#	x=$(docker ps --filter status=running --filter "name=mysql" | wc -l)
#	echo $x
	#if [ ${x} -eq 1 ] ; then
	docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=ons_vg -d mysql
#; fi
db-create:
	mysql -u root -h 127.0.0.1 -pons_vg -e "CREATE DATABASE IF NOT EXISTS BOOK_MY_SHOW;"

db-drop:
	mysql -u root -h 127.0.0.1 -pons_vg -e "DROP DATABASE IF EXISTS BOOK_MY_SHOW;"

db-migrate-up:
	migrate -path db/migration -database "mysql://root:ons_vg@tcp(127.0.0.1:3306)/BOOK_MY_SHOW" -verbose up

db-migrate-down:
	migrate -path db/migration -database "mysql://root:ons_vg@tcp(127.0.0.1:3306)/BOOK_MY_SHOW" -verbose down

# Swagger ###########
check-swagger-install:
	which swagger || brew tap go-swagger/go-swagger || brew install go-swagger

generate-swagger: check-swagger-install
	swagger generate spec -o swagger-ui/swagger.yaml --scan-models

.PHONY: mysql db-create db-drop