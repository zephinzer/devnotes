# Google Container Registry

## Format
```
[HOSTNAME]/[PROJECT-ID]/[IMAGE][:TAG|@DIGEST]
```
### Hostname
One of:
- `us.gcr.io`
- `eu.gcr.io`
- `asia.gcr.io`

### Project ID
This is your project's ID

### Image
Your application's container name

### Tag
Tag for your container

## Pushing Images
### Build It
```
docker build \
  -t [] \
  -f %__dockerfile_path__% \
  .
```

### Tag It
```
docker tag %__tag_name__% asia.gcr.io/
```

### Push It
```
gcloud docker -- push [HOSTNAME]/[PROJECT-ID]/[IMAGE][:TAG]
```

## Relevant Links
[Pushing and Pulling Images](https://cloud.google.com/container-registry/docs/pushing-and-pulling?hl=en_US&_ga=2.233967006.-1261252008.1518271864)