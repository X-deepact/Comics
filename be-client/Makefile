.PHONY: docs
docs:
	@echo "Generating Swagger Documentation..."
	swag init --parseDependency --parseInternal
	@echo "Swagger Documentation Generated!"

.PHONY: run
run:
	@echo "Starting the server..."
	sh run.sh start

.PHONY: all
all: docs run
