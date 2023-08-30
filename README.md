# jsonnet-config-demo-go

A Go demo using https://jsonnet.org/ for configuration.
Package `config` provides a generic function `MustParse`, that evaluates the given config file using jsonnet and unmarshals the generated JSON into the given configuration struct type with exported fields. [github.com/go-playground/validator](https://github.com/go-playground/validator) struct tags can be used to constraint individual fields.
