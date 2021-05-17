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

## ML Model

> WARNING: Never commit models, model.zip, or model binary to `tensorflow-serving/model` dir!

### Build tensorflow-serving locally

```bash
gsutil cp -r gs://bangkit-aaida-model/* tensorflow-serving/model/ 
docker build -t tensorflow-serving tensorflow-serving/
docker run -it --rm -p 8500:8500 -p 8501:8501 tensorflow-serving
```

### How to upload ML models to Google Cloud Storage

```bash
unzip model.zip
gsutil cp -r model/ gs://bangkit-aaida-model/
```

## TODO

- Implement CI/CD using Spinnaker
