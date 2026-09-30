up:
	goose -dir sql/schema postgres "postgres://postgres:postgres@localhost:5432/gator" up

down:
	goose -dir sql/schema postgres "postgres://postgres:postgres@localhost:5432/gator" down

gator:
	sudo -u postgres psql gator




.ONESHELL:
ARGS := $(filter-out $@,$(MAKECMDGOALS))

.ONESHELL:
ARGS := $(filter-out $@,$(MAKECMDGOALS))

.ONESHELL:
ARGS := $(filter-out $@,$(MAKECMDGOALS))

.ONESHELL:
ARGS := $(filter-out $@,$(MAKECMDGOALS))

run:
	# Cleanly execute without if/else blocks by letting the shell expand $(ARGS) safely
	-@go run . $(ARGS) 2>/dev/null

git: 
	ENV_ARG="$(ARGS)"
	if [ -z "$$ENV_ARG" ]; then echo "Error: Please provide a commit message!"; exit 1; fi
	git add .
	-@git commit -m "$$ENV_ARG" 2>/dev/null
	-@git push -u origin main 2>/dev/null

%:
	@: