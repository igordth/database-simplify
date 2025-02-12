# List off commands
list:
	@grep '^[^[:space:]].*:' Makefile | grep -Eo '^[^:]*'

# Run application with
#	args <app> - directory of application
#	example > make run app=connection
run: docker-recreate
	go run ./$(app)

# Short command of docker-recreate
docker: docker-recreate

# Recreate docker containers and recreate db scheme
docker-recreate: docker-down docker-up

# Run docker compose
docker-up:
	@docker compose -p ds-pggorm --project-directory ./build up -d

# Down docker containers & remove used volume
docker-down:
	@docker compose -p ds-pggorm down -v
