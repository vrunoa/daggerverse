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

