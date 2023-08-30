# jsonnet-config-demo-go

A Go demo using https://jsonnet.org/ for configuration.
Package `config` provides a generic function `MustParse`, that evaluates the given config file using jsonnet and unmarshals the generated JSON into the given configuration struct type with exported fields. [github.com/go-playground/validator](https://github.com/go-playground/validator) struct tags can be used to constraint individual fields.

In the provided [example](https://github.com/romshark/jsonnet-config-demo-go/tree/main/cmd/example1), the following jsonnet configuration:

```jsonnet
local AdminName(name='<untitled>') = 'user_admin_' + name;

{
  host: 'localhost:8080',
  admins: [AdminName(name='Bob'), AdminName(name='Alice')],
}
```

Is parsed into:

```go
type Config struct {
	Host   string   `json:"host" validate:"required,hostname_port"`
	Admins []string `json:"admins" validate:"required"`
}
```
