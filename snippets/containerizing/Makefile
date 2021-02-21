folder := containerizing

.DEFAULT_GOAL := help
.PHONY: help
help: ## Affiche cette aide
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Construit l'image de l'application
	podman build -t $(folder) .

.PHONY: run
run: ## Lance le container de l'application
	podman run -p 8081:8081 -it $(folder)

.PHONY: ps
ps: ## Les containers lancés
	podman ps
