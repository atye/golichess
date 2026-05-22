module github.com/atye/golichess/examples/oapi-codegen

go 1.24.13

replace github.com/atye/golichess/oapi-codegen => ../../oapi-codegen

require github.com/atye/golichess/oapi-codegen v0.0.0-00010101000000-000000000000

require (
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/oapi-codegen/runtime v1.4.1 // indirect
)
