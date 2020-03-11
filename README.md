# go-gis
Sample with Google Maps / Golang / Postgis 

## Data
The Shape data come from [data.gouv.fr](https://www.data.gouv.fr/fr/datasets/contours-des-regions-francaises-sur-openstreetmap/)
Same for the [SNCF Data](https://www.data.gouv.fr/fr/datasets/fichier-de-formes-des-lignes-du-reseau-ferre-national/)

## Init project
### Create the database
You need docker :
`make launch-db`

### Import datasets
```
make import-departements
make import-regions
make import-sncf
```

### Launch project
Environment Variables : 
- DB_CONNECTION_URI (should be `DB_CONNECTION_URI=postgresql://karnott:12345678@localhost:5432/karnott`)
- GOOGLE_MAPS_KEY
- APPSETTING_PORT (optional, 80 by default)

Then : 
```
go mod vendor
go run main.go
```
