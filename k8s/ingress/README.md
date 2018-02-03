# Ingresses
 
## Using Specfile

### Template for `nginx-ingress-controller`

> Creates a basic Ingress resource using `nginx-ingress-controller`

```yaml
kind: Ingress
apiVersion: extensions/v1beta1
metadata:
  name: __INGRESS_NAME__
  labels:
    app: __INGRESS_NAME__
    url: "__DOMAIN_NAME__"
  annotations:
    kubernetes.io/ingress.class: "nginx"
    kubernetes.io/ingress.global-static-ip-name: __GLOBAL_STATIC_IP_ADDRESS__
spec:
  backend:
    serviceName: __BACKEND_DEFAULT_SERVICE_NAME__
    servicePort: __BACKEND_DEFAULT_SERVICE_PORT__
  rules:
    - host: __DOMAIN_NAME__
      http:
        paths:
        - path: /
          backend:
            serviceName: __BACKEND_LIVE_SERVICE_NAME__
            servicePort: __BACKEND_LIVE_SERVICE_PORT__
```

`__BACKEND_DEFAULT_SERVICE_NAME__` : A default application which should return 404 on `/` endpoint and 200 on `/healthz` endpoint

`__BACKEND_DEFAULT_SERVICE_PORT__` : Port which the default application listens on at the Service level

`__BACKEND_LIVE_SERVICE_NAME__` : Your application's Service's name

`__BACKEND_LIVE_SERVICE_NAME__` : Port which your application listens on at the Service level

`__DOMAIN_NAME__` : Domain name for accessing your application

`__GLOBAL_STATIC_IP_ADDRESS__` : Name for a created reserved global static IP address

`__INGRESS_NAME__` : Name for your Ingress resource
