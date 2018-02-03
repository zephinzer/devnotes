# Secrets

## Using `kubectl`

### Creating a TLS Secret resource

```bash
$ kubectl create secret tls __SECRET_NAME__ --key ./path/to/key --cert ./path/to/crt
```

## Using Specfile

### Template for TLS Secret

```yaml
kind: Secret
apiVersion: v1
metadata:
  name: __SECRET_NAME__
type: kubernetes.io/tls
data:
  tls.key: __BASE64_ENCODED_KEY__
  tls.crt: __BASE64_ENCODED_CERT__
```

`__SECRET_NAME__` : Referenceable ID of this Secret resource

`__BASE64_ENCODED_KEY__` : Run `cat privateKey.key | base64` to get the base64 encoded private key

`__BASE64_ENCODED_CERT__` : Run `cat certificate.crt | base64` to get the base64 encoded certificate file