# skaffold example

## local
The profile must set to `local`.

### build

```bash
skaffold build
```

###  dev-cycle


```bash
skaffold dev -p local
```

## profiles

### kaniko

```bash
# pullsecret for dockerhub
DOCKERHUB_USER=...
DOCKERHUB_EMAIL=...
DOCKERHUB_PASSWORD=...
kubectl create secret docker-registry regcred --docker-server=https://index.docker.io/v1/ --docker-username=$DOCKERHUB_USER --docker-password=$DOCKERHUB_PASSWORD --docker-email=$DOCKERHUB_EMAIL

DOCKER_HUB_AUTH=$(echo -n $DOCKERHUB_USER:$DOCKERHUB_PASSWORD | base64)

cat <<EOF >config.json
{
    "auths": {
        "https://index.docker.io/v1/": {
            "auth": "${DOCKER_HUB_AUTH}"
        }
    }
}
EOF

kubectl create secret generic kaniko-secret --from-file=config.json

# build and push from k8s to dockerhub
skaffold build -p cloud-dev
```
