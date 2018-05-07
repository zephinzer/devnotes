# SSL Certificates

## Generating a Self-Signed SSL Certificate

```sh
openssl req \
  -x509 \
  -newkey rsa:4096 \
  -keyout ./secrets/certs/key.pem \
  -passout pass:${SERVER_CERT_PASSWORD} \
  -out ./secrets/certs/cert.pem \
  -days 365 \
  -subj "/CN=${USER_CERT_DOMAIN}/O=${USER_CERT_ID}/C=US";
```

## Generating a User/Certificate Authority Key

```sh
openssl genrsa \
  -des3 \
  -out ./secrets/ca.key \
  -passout pass:${SERVER_CERT_PASSWORD} \
  4096;
```

## Generating a Certificate Authority

```sh
openssl req \
  -new \
  -x509 \
  -key ./secrets/ca.key \
  -passin pass:${SERVER_CERT_PASSWORD} \
  -out ./secrets/ca.crt \
  -days 365 \
  -subj "/CN=${USER_CERT_DOMAIN}/O=${USER_CERT_ID}/C=US";
```

## Generating a Certificate Signing Request for a User

```sh
openssl req \
  -new \
  -key ./secrets/user.key \
  -passin pass:${SERVER_CERT_PASSWORD} \
  -out ./secrets/user.csr \
  -subj "/CN=${USER_CERT_DOMAIN}/O=${USER_CERT_ID}/C=US";
```

## Signing a Certifcate Signing Request for a User using a Certificate Authority

```sh
openssl x509 \
  -req \
  -days 365 \
  -in ./secrets/user.csr \
  -CA ./secrets/ca.crt \
  -CAkey ./secrets/ca.key \
  -passin pass:${SERVER_CERT_PASSWORD} \
  -set_serial 01 \
  -out ./secrets/auth/user.crt;
```

## Exporting a User Certificate to PFX Format

```sh
openssl pkcs12 \
  -export \
  -out ./secrets/browser/user.pfx \
  -inkey ./secrets/user.key \
  -passin pass:${SERVER_CERT_PASSWORD} \
  -in ./secrets/auth/user.crt \
  -passout pass:${USER_CERT_PASSWORD} \
  -certfile ./secrets/ca.crt;
```