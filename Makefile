MONGO_NAME ?= your_mongo
MONGO_PORT ?= 27017
SERVER_NAME ?= your_server

MONGODB_VERSION=6.0-ubi8
MONGO_USER ?= user
MONGO_PASSWORD ?= pwd

default_setup:
	echo '{"username": "$(MONG_USER)","password": "$(MONGO_PASSWORD)"}' > /tmp/mongo_creds.json


run_db:
	docker run -e MONGODB_INITDB_ROOT_USERNAME=$(MONGO_USER) -e MONGODB_INITDB_ROOT_PASSWORD=$(MONGO_PASSWORD) \
--name $(MONGO_NAME) -d -p $(MONGO_PORT):27017 mongodb/mongodb-community-server:$(MONGODB_VERSION)


run_server:
	make build
	# the server should run from "back-end/src" due to `configs/` directory
	cd back-end/src && ../../bin/$(SERVER_NAME)


run_fe:
	cd front-end && npm run-script dev


run_all:
	make run_db
	make run_fe &> /tmp/fe.log &
	make run_server &> /tmp/server.log &


stop_all:
	docker stop $(MONGO_NAME) && docker rm $(MONGO_NAME) || true
	lsof -i :3000 | awk '{print $2}' | xargs kill -9 || true
	pkill -f "npm run-script dev" || true


force_run_all:
	stop_all || true
	default_setup
	docker stop $(MONGO_NAME) && docker rm $(MONGO_NAME) || true
	make run_all
