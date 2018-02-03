# Iron Functions
This directory contains k8s manifests for the Iron Functions functions-as-a-service platform.

Set up on 26th January 2017 by @zephinzer <`dev-at-joeir-dot-net`>.

## Deployments

Iron Functions: https://ironfx.gds-gov.tech
Iron Functions UI: https://ironfx-ui.gds-gov.tech

## What's Included

1. Iron Functions
  - to manage functions
    - `{config, deployment, ingress, service}.yaml`
  - for security
    - `{ingress.api, secret.api-auth}.yaml`
2. Iron Functions UI
  - for a user interface
    - `ext.ui.yaml`
  - for security
    - `secret.ext.ui-auth.yaml`
3. PostgreSQL
  - for the database
    - `ext.postgres.yaml`
  - for persistence
    - `volumes.yaml`
4. Redis
  - for the queue
  - `ext.redis.yaml`

## Further Reading
- [Iron Functions GitHub](https://github.com/iron-io/functions)
- [Iron Functions UI GitHub](https://github.com/iron-io/functions-ui)