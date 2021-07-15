M = $(shell printf "\033[34;1m▶\033[0m")
COMPOSE := docker-compose -f docker/docker-compose.yaml
.PHONY: run
run: 
	$(info $(M) Starting development server with reload...)
	realize start
.PHONY: dev
dev: ##@Dev Run Dev server with Reload
	$(info $(M) Starting development...)
	$(COMPOSE) up $(ARGS)