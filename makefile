command:
	echo frontend-dev
	echo backend-dev
	echo frontend-prod
	echo backend-prod
	echo backend-build
	echo create-db

create-db:
	echo > database.db && sqlite3 database.db < migration.sql

frontend-dev:
	cd frontend && pnpm run dev

frontend-build:
	cd frontend && time pnpm run build && pnpm run start

frontend-prod:
	cd frontend && pnpm run start

docker:
	docker compose up

backend-dev:
	cd backend && air

backend-prod:
	cd backend && go build && ./backend

.SILENT:
.PHONY: command