include .env

up:
	@echo "Starting containers..."
	docker-compose up --build -d --remove-orphans

down:
	@echo "Stopping containers..."
	docker-compose down

build:
	go build -o ${BINARY} ./cmd/api/

start:
	@env MONGO_DB_USERNAME=${MONGO_DB_USERNAME} MONGO_DB_PASSWORD=${MONGO_DB_PASSWORD} ./${BINARY}

restart: build start


#connection string
# mongodb://admin:password@localhost:27017/mongo_api?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false