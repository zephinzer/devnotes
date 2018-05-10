# Goldfish UI for Vault
This set of instructions is for a Goldfish + Vault setup on Kubernetes.

## Run a Kube Execute into Vault
```sh
kubectl get pods -n common-services | grep vault;
kubectl exec -it -n common-services ${VAULT_POD_ID} -c vault /bin/sh;
```
## Unseal Vault
```sh
vault unseal [-tls-skip-verify];
```

## Get Access Token
```sh
vault auth [-tls-skip-verify];
```

## Create Wrapping Token
```sh
vault write [-tls-skip-verify] -f -wrap-ttl=5m auth/approle/role/goldfish/secret-id;
```

Paste the generated token into Goldfish.