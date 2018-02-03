# Automatic SSL Certificate Provisioning with Let's Encrypt
The tool to use is `kube-lego`. Check out:

- [The GitHub Repository](https://github.com/jetstack/kube-lego)
- [This Useful Guide](https://www.matt-j.co.uk/2017/03/03/automatic-dns-and-ssl-on-kubernetes-with-letsencrypt-part-2/)

Instructions follow:

## Set up annotations on your Ingress
Set up your target Ingresses so that they contain the following annotation in the `metadata.annotations[]` field:

`kubernetes.io/tls-acme: "true"`

## Add the TLS specifications
Next, add the `tls` field to your `spec` field as follows:

```yaml
...
spec:
  ...
  tls:
  - secretName: secret1-tls
    hosts:
    - example.com
    - api.example.com
  - secretName: secret2-tls
    hosts:
    - example2.com
```

`kube-lego` will detect the annotation above, and read from the `tls` section. For the above example, two certificates will be registered, one for `example.com` and `api.example.com`, and the second certificate will be for `example2.com`. The two certificates will be automatically created as Kubernetes Secrets with the names `secret1-tls` and `secret2-tls` in the namespace which your Ingress exists in.

> It doesn't matter that `secret1-tls` and `secret2-tls` do not exist, they will be auto-created by `kube-lego`. Good stuff.

## Deploy `kube-lego`

Create a specfile with the following contents:

```yaml
kind: ConfigMap
apiVersion: v1
metadata:
  name: kube-lego-cm
  namespace: kube-system
data:
  lego.email: "email@example.com"
  lego.url: "https://acme-v01.api.letsencrypt.org/directory"
  # staging: "https://acme-staging.api.letsencrypt.org/directory"
  # production: "https://acme-v01.api.letsencrypt.org/directory"

---

kind: Deployment
apiVersion: extensions/v1beta1
metadata:
  name: kube-lego
  namespace: kube-system
spec:
  replicas: 1
  template:
    metadata:
      labels:
        app: kube-lego
    spec:
      containers:
      - name: kube-lego
        image: jetstack/kube-lego:0.1.5
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 8080
        env:
        - name: LEGO_EMAIL
          valueFrom:
            configMapKeyRef:
              name: kube-lego-cm
              key: lego.email
        - name: LEGO_URL
          valueFrom:
            configMapKeyRef:
              name: kube-lego-cm
              key: lego.url
        - name: LEGO_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: LEGO_POD_IP
          valueFrom:
            fieldRef:
              fieldPath: status.podIP
        readinessProbe:
          httpGet:
            path: /healthz
            port: 8080
          initialDelaySeconds: 10
          timeoutSeconds: 3
```

## Verifying Success/Debugging
To see what is happening, identify the Pod running the specified Deployment:

```bash
$ kubectl get pods -n=kube-system
```

Search for a pod named `kube-lego-XXXXXXXXXX-yyyyy` and check the logs:

```bash
$ kubectl log -n=kube-system -f kube-lego-XXXXXXXXXX-yyyyy
```

On success, you should see something resembling the following where:

- `example.com` is a specified domain in an entry in the `tls` field
- `secret-name-tls` is the `metadata.name` of the Secret as specified in the `secretName` field in an entry of the `tls` field.
- `your-ingress` is the `metadata.name` of your Ingress resource

```
time="2017-10-24T17:24:08Z" level=info msg="Attempting to create new secret" context=secret name=secret-name-tls namespace=default
time="2017-10-24T17:24:08Z" level=info msg="no cert associated with ingress" context="ingress_tls" name=your-ingress namespace=default
time="2017-10-24T17:24:08Z" level=info msg="requesting certificate for example.com" context="ingress_tls" name=your-ingress namespace=default
time="2017-10-24T17:24:08Z" level=info msg="Attempting to create new secret" context=secret name=kube-lego-account namespace=kube-system
time="2017-10-24T17:24:09Z" level=info msg="if you don't accept the TOS (https://letsencrypt.org/documents/LE-SA-v1.1.1-August-1-2016.pdf) please exit the program now" context=acme
time="2017-10-24T17:24:09Z" level=info msg="created an ACME account (registration url: https://acme-v01.api.letsencrypt.org/acme/reg/23208775)" context=acme
time="2017-10-24T17:24:09Z" level=info msg="Attempting to create new secret" context=secret name=kube-lego-account namespace=kube-system
time="2017-10-24T17:24:09Z" level=info msg="Secret successfully stored" context=secret name=kube-lego-account namespace=kube-system
time="2017-10-24T17:24:46Z" level=info msg="authorization successful" context=acme domain=example.com
time="2017-10-24T17:24:46Z" level=info msg="successfully got certificate: domains=[example.com] url=https://acme-v01.api.letsencrypt.org/acme/cert/034ee9178e8435897ea6a9069e2e8b9d5347" context=acme
time="2017-10-24T17:24:46Z" level=info msg="Attempting to create new secret" context=secret name=secret-name-tls namespace=default
time="2017-10-24T17:24:46Z" level=info msg="Secret successfully stored" context=secret name=secret-name-tls namespace=default
```

## Troubleshooting/Common Errors

### 403 Unauthorized

`authorization failed after 1m0s: getting authorization failed: 403 urn:acme:error:unauthorized: No registration exists matching provided key`

This error happens when you first run `kube-lego-nginx` with the staging Let's Encrypt URL and then again with the production URL. To resolve this, delete the `kube-lego-account` secret. Search for it:

```bash
$ kubectl get secret -n=kube-system | grep kube-lego
```

Down the `kube-lego` deployment:

```bash
$ kubectl delete deploy -n=kube-system kube-lego
$ kubectl delete cm -n=kube-system kube-lego-cm
```

Then delete the `kube-lego-account`:

```bash
$ kubectl delelte secret -n=kube-system kube-lego-account
```

Confirm you are using the **production** Let's Encrypt URL instead of the staging, and deploy the above specfile again. 
