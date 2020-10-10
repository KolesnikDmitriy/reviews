module github.com/KolesnikDmitriy/reviews

go 1.15

require (
	github.com/KolesnikDmitriy/reviews/pkg/api v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.33.0
)

replace github.com/KolesnikDmitriy/reviews/pkg/api => ./pkg/api
