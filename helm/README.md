# hnrss (Helm)

Deploy from this directory (`helm/`):

```bash
helm upgrade --install hnrss ./ --namespace hnrss --set registryAuth.password="$REGISTRY_PASSWORD"
```

Set `REGISTRY_PASSWORD` to the registry’s HTTP basic-auth password so the chart can create the pull secret (see `values.yaml` → `registryAuth`).
