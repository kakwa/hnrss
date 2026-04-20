# hnrss (Helm)

Deploy from this directory (`helm/`):

```bash
helm upgrade --install hnrss ./ --namespace hnrss --set registryAuth.password="$REGISTRY_PASSWORD"
