.PHONY: docs
docs:
	@echo "Generating Swagger Documentation..."
	swag init -g main.go
	@echo "Swagger Documentation Generated!"

.PHONY: run
run:
	@echo "Starting the server..."
	sh run.sh start

.PHONY: all
all: docs run
