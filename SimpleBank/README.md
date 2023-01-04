migrate create -ext sql -dir db/migration -seq init_schema

docker exec -it postgres12 createdb --username=root --owner=root simple_bank
docker exec -it postgres12 psql -U root simple_bank
