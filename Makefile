up_db:
	docker-compose exec -T db psql "dbname=delivery user=user password=password" <backend/migrations/up.sql

down_db:
	docker-compose exec -T db psql "dbname=delivery user=user password=password" <backend/migrations/down.sql


