# skaffold example

## local dev-cycle

The profile must set to `local`.

```bash
# build
skaffold dev -p local
```

## profiles

### kaniko

```bash
# pullsecret for dockerhub
DOCKERHUB_USER=alitari
DOCKERHUB_EMAIL=alitari67@gmail.com
DOCKERHUB_PASSWORD=....
kubectl create secret docker-registry regcred --docker-server=https://index.docker.io/v1/ --docker-username=$DOCKERHUB_USER --docker-password=$DOCKERHUB_PASSWORD --docker-email=alitari67@gmail.com


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



```
