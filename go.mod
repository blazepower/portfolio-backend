module github.com/blazepower/portfolio-backend

go 1.17

require (
	github.com/blazepower/portfolio-backend/Handlers v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
)

replace github.com/blazepower/portfolio-backend/Handlers => ./Handlers
