module github.com/KolesnikDmitriy/reviews

go 1.15

require (
	github.com/KolesnikDmitriy/reviews/pkg/api v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.6.1 // indirect
	google.golang.org/grpc v1.33.0
	google.golang.org/protobuf v1.25.0 // indirect
)

replace github.com/KolesnikDmitriy/reviews/pkg/api => ./pkg/api
