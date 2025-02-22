# daggerverse

A collection of dagger modules for pipelines and local development.

#### golangci-lint

```sh {name="lint"}
dagger call -m github.com/vrunoa/daggerverse/go@main lint --src=$(pwd)
```


#### openapi validatation

```sh {name="openapi-validate"}
dagger call -m github.com/vrunoa/daggerverse/go@main openapi-validate --src=$(pwd) --spec=./petstore.yaml
```

#### oapi-codegen
```sh {name="oapi-codegen"}
dagger call -m github.com/vrunoa/daggerverse/go@main codegen --src=$(pwd) --config=./codegen/manager.yaml --spec=./spec/manager.yaml --target=gen/service -o ./gen/service
````
