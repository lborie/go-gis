launch-db:
	docker run --name=geo-gis -d -e POSTGRES_DB=karnott -e POSTGRES_USER=karnott -e POSTGRES_PASSWORD=12345678 -p 5432:5432 mdillon/postgis:latest