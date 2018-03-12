# Google Kubernetes Engine

## Setting Up A Cluster
1. Create a new project (top navigation bar next to the **Google Cloud Platform** logo).
2. Click on the bar icon (top navigation bar left most) and click on **Kubernetes Engine**.
3. In the left navigation panel, click on **Kubernetes clusters**.
4. Click on the **Create cluster**.
5. Set options as required.

## Accessing a Cluster
1. Install the Gcloud API - instructions at https://cloud.google.com/sdk/docs/quickstarts
2. Run `gcloud init` and setup a configuration
3. Run `gcloud config set compute/zone <zone>`
4. Run `gcloud container clusters get-credentials <cluster_name>`

You should now be able to use your `kubectl` to access your cluster.