package opentonapi

//go:generate go run github.com/ogen-go/ogen/cmd/ogen -clean -config ogen.yaml -package oas -target pkg/oas api/openapi.yml
//go:generate go run api/jsonify.go api/openapi.yml api/openapi.json
