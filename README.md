# ci-with-skaffold

## requirements

- Docker Desktop
  - enable Kubernetes
- kubectl
- skaffold
- kustomize
- envsubst

## how to use

```console
make prepare
export DB_PORT=... # see the output the enf of `make prepare`.
skaffold dev -p local
```
