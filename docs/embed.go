package docs

import "embed"

//go:embed openapi.yaml
var spec embed.FS

func OpenApiSpec() (*[]byte, error) {
	f, err := spec.ReadFile("openapi.yaml")
	if err != nil {
		return nil, err
	}
	return &f, nil
}
