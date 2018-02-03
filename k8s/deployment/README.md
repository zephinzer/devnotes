# Deployments

## Using `kubectl`
### Creating
> This assumes we want to start an `nginx` service using the `nginx` image, exposing port 80.

```bash
$ kubectl run nginx --image=nginx --port=80
```

### Exposing
> This exposes an existing deployment named `nginx` which has port 80 exposed.

```bash
$ kubectl expose deployment nginx --target-port=80 --type=NodePort
```

### Scaling
> This scales a deployment named `nginx` to use 2 replicas.

```bash
$ kubectl scale deployment nginx --replicas=2
```

### Updating
> This updates a deployment using a `specfile`

```bash
$ kubectl apply -f ./path/to/specfile.yaml
```

## Using Specfile
### Template
Save the following as a `.yaml` file:

```yaml
kind: Deployment
apiVersion: extensions/v1beta1
metadata:
  name: __DEPLOYMENT_NAME__
  labels:
    app: __DEPLOYMENT_NAME__
spec:
  replicas: 2
  strategy:
    type: RollingUpdate
    maxUnavailable: 1
    maxSurge: 4
  selector:
    matchLabels:
      app: __DEPLOYMENT_NAME__
  template:
    metadata:
      labels:
        app: __DEPLOYMENT_TEMPLATE_NAME__
    spec:
      containers:
      - name: __DEPLOYMENT_CONTAINER_NAME__
        image: asia.gcr.io/__PROJECT_NAME__/__IMAGE_NAME__:1.0.0
        imagePullPolicy: Always
        ports:
        - containerPort: __DEPLOYMENT_CONTAINER_PORT__
```

### Variables

`__DEPLOYMENT_NAME__` : The name of your deployment

`__DEPLOYMENT_TEMPLATE_NAME__` : The name of your template

`__DEPLOYMENT_CONTAINER_NAME__` : The name of your container

`__DEPLOYMENT_CONTAINER_PORT__` : The port which your container listens on

`__PROJECT_NAME__` : Your project name in Google Cloud Platform (GCP)

`__IMAGE_NAME__` : Your image's name in Google Container Registry (GCR)