migrate create -ext sql -dir db/migration -seq init_schema


docker exec -it postgres12 psql -U root simple_bank