.PHONEY: build up down logs sqlc db migrate-up migrate-reset migrate-status

build:
	docker compose build --no-cache

up:
	docker compose up -d

down:
	docker compose down

logs:
	docker compose logs -f

sqlc:
	docker compose exec -it app sqlc -f infra/db/sqlc.yaml generate

db:
	docker compose exec -it db mysql -ubspliter -pbspliter bspliter

migrate-up:
	docker compose exec -it app goose -dir infra/db/ddl mysql "bspliter:bspliter@tcp(db:3306)/bspliter" up

migrate-reset:
	docker compose exec -it app goose -dir infra/db/ddl mysql "bspliter:bspliter@tcp(db:3306)/bspliter" reset

migrate-status:
	docker compose exec -it app goose -dir infra/db/ddl mysql "bspliter:bspliter@tcp(db:3306)/bspliter" status
