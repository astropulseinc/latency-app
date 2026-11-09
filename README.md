# latency-app

A small Go web app that measures the round trip from your browser to the server. It is the sample service the [AstroPulse](https://astropulse.io) examples deploy.

Image: `docker.io/astropulseinc/latency-app` (linux/amd64 and linux/arm64)

## Run it

```bash
docker run -p 8080:8080 docker.io/astropulseinc/latency-app:latest
```

Open http://localhost:8080. From source: `make run`.

## Deploy it with AstroPulse

In the AstroPulse console, go to **Workloads > Applications > Deploy application**. Choose **Container image** with `docker.io/astropulseinc/latency-app` on port `8080`. The full walkthrough is [Build your own developer platform](https://astropulse.io/docs/examples/internal-developer-platform/path-to-production).

## Endpoints

| Path | Returns |
|------|---------|
| `/` | The page |
| `/ping-latency` | The server time as JSON; the page measures the round trip |
| `/healthz`, `/health` | Health, for readiness and liveness probes |

## What's here

| Path | What it is |
|------|------------|
| `main.go`, `index.html` | The app |
| `Dockerfile` | A static build on a distroless, non-root base |
| `k8s/` | Plain Kubernetes manifests: Deployment and Service |
| `charts/latency/` | The same app as a Helm chart, with `values-prod.yaml` for two replicas |
| `.astropulse/` | AstroPulse application files for building the image in your cluster with kpack and deploying it |

## Develop

```bash
make test    # vet and test
make build   # bin/latency
make docker  # local image
```

A merge to `main` creates the next version tag. A merge whose commit message contains `[DOCKER_GENERATE]` also publishes the image under that tag.

## License

Apache 2.0
