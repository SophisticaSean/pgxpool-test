# Quick Test of PGX Pool in Golang
  - Confirmed working on go 1.19

## To Run:
  - docker run --name pgxpooltest-db -p 5455:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=pgxpooltest_dev -d postgres
  - export DATABASE_URL="postgresql://postgres:postgres@0.0.0.0:5455"
  - go get
  - go run main.go
