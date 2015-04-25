## Deployment on Google Compute Engine

### Register Image

The docker image must first be registered with the [Google Container Registry](https://cloud.google.com/tools/container-registry/).

The docker image must be tagged.

```
$ docker build -t gcr.io/{project-id}/{image-name} .
```

The docker image must be pushed to the registry.

```
$ gcloud preview docker push gcr.io/{project-id}/{image-name}
```

### Create Instance

Specify the image in the `containers.yaml` manifest file.

```
version: v1beta2
containers:
  - name: {image-name}
    image: gcr.io/{project-id}/{image-name}
```

Creating the instance.
    
```
$ gcloud compute instances create {instance-name} \
    --image container-vm \
    --metadata-from-file google-container-manifest=containers.yaml \
    --tags http-server \
    --zone us-central1-a \
    --machine-type f1-micro
```