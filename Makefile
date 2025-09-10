# Carrega variáveis do .env se existir (não reclama se não existir)
-include .env

# Exporta só o que você precisa para os recipes (opcional, mas útil)
export DB_CONN_URL

# Permite sobrescrever via ambiente, senão usa do .env
POSTGRESQL_URL ?= $(DB_CONN_URL)
server:
	@go run ./cmd/http/.

exp:
	@go run ./cmd/exp/exp.go

db:
	@docker compose up -d

migrate-up:
	@migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrate-down:
	@migrate -database ${POSTGRESQL_URL} -path db/migrations down

.PHONY: server exp migrate-up migrate-down	