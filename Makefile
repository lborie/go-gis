launch-db:
	docker run --name=geo-gis -d \
	-e POSTGRES_DB=karnott \
	-e POSTGRES_USER=karnott \
	-e POSTGRES_PASSWORD=12345678 \
	-v $$(pwd)/resources:/workdir \
	-p 5432:5432 mdillon/postgis:latest

import-departements:
	docker exec -it geo-gis /bin/bash -c 'shp2pgsql -s 4326 -d -I /workdir/departements-20180101-shp/departements-20180101.shp | psql -U karnott -d karnott'

import-regions:
	docker exec -it geo-gis /bin/bash -c 'shp2pgsql -s 4326 -d -I /workdir/regions-20180101-shp/regions-20180101.shp | psql -U karnott -d karnott'
