# Demo Script

## Environment setup

```sh
podman machine start
ibm_login_eu
iks cluster config --cluster af-cdevents
ibmcloud cr login
cp ~/.config/containers/auth.json ~/.docker/config.json
```

## Intro to the demo setup

- Fork of the podtato application, hosted on a gitea instance
- Tools I'm using
  - Tekton triggers as sync for webhooks and events
  - Knative eventing provides cloudevent broker and k8s api event source
  - The cdevents go SDK to produce CDEvents
  - IBM Cloud k8s cluster
  - Azure container registry

## Create a Deployment

```sh
ARTIFACT_URL=cdevents.azurecr.io/podtatoserver@sha256:3589637f88593d730fdb3cd32558588d846bc699cbdc04e8a97f200687c4b6cf
kubectl create deployment podtato-server --image=$ARTIFACT_URL -n production
open https://podtato-head.af-cdevents-dff43bc8701fcd5837d6de963718ad39-0000.eu-gb.containers.appdomain.cloud/
open https://tekton.af-cdevents-dff43bc8701fcd5837d6de963718ad39-0000.eu-gb.containers.appdomain.cloud/#/taskruns
```

## Make a PR

```sh
vim docker/Dockerfile
git add -u
git commit -s -m "Update server version"
open https://tekton.af-cdevents-dff43bc8701fcd5837d6de963718ad39-0000.eu-gb.containers.appdomain.cloud/#/taskruns
```

## Rebuild

```sh
git checkout main
git pull
./build/build_image.sh 0.1.1
ARTIFACT_URL=<output from build script>
open https://tekton.af-cdevents-dff43bc8701fcd5837d6de963718ad39-0000.eu-gb.containers.appdomain.cloud/#/taskruns
```

## Update the deployment

```sh
kubectl set image deployment/podtato-server -n production podtatoserver=$ARTIFACT_URL
```
