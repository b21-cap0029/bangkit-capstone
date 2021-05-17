# Cloud Computing

Short README here


## Directories

```
.
|-- aaida-backend
|-- cloudbuild.yaml
`-- k8s
```

- aaida-backend directory contains the aaida-backend application source code and Dockerfile
- cloudbuild.yaml is located here and run on Google Cloud Build to make multiple images
- k8s directory holds kubernetes manifest files such as deployment and services, for now

## How to upload ML models to Google Cloud Storage

```bash
unzip model.zip
gsutil cp -r model/ gs://bangkit-aaida-model/
```

## TODO

- Implement CI/CD using Spinnaker
