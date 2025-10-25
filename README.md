# Latency App

A real-time network latency measurement application built with Go, showcasing deployment on the [AstroPulse](https://astropulse.io) platform.

![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)
[![Docker Hub](https://img.shields.io/docker/v/astropulseinc/latency-app?label=docker)](https://hub.docker.com/r/astropulseinc/latency-app)
[![Docker Pulls](https://img.shields.io/docker/pulls/astropulseinc/latency-app)](https://hub.docker.com/r/astropulseinc/latency-app)

## Features

- 🎯 Real-time latency measurement with interactive web UI
- 📊 Statistics dashboard with min/max/average metrics
- 📈 Visual latency history (last 50 measurements)
- 🔄 Continuous monitoring mode
- 🏥 Health check endpoints for Kubernetes
- 🚀 Production-ready deployment with kpack + Kubernetes

## Quick Start

### Run with Docker

```bash
# Pull and run from Docker Hub
docker run -p 8080:8080 astropulseinc/latency-app:latest

# Open browser
open http://localhost:8080
```

### Local Development

```bash
# Build
make build

# Run
./bin/latency

# Open browser
open http://localhost:8080
```

### Build Docker Image

```bash
# For local testing (single platform, loads into Docker)
make docker-build-local

# For publishing (multi-platform: amd64 + arm64)
# Requires: git tag + docker login
git tag v1.0.0
docker login
make docker-build  # This builds AND pushes to Docker Hub
```

### Deploy to AstroPulse

Follow the complete deployment guide at **[astropulse.io/get-started](https://astropulse.io/get-started)**

**Quick deployment:**

```bash
# 1. Install astroctl CLI (see link above)
# 2. Configure your API key
export ASTROPULSE_API_KEY=your-api-key

# 3. Customize configuration files
# - Update .astropulse/*.yaml with your AWS account, domain, and repo URL
# - See "Configuration" section below

# 4. Deploy
astroctl app profile apply -f .astropulse/profile.yaml
astroctl app apply -f .astropulse/build-app.yaml
astroctl app apply -f .astropulse/deploy-app.yaml
```

## Configuration

Before deploying, update these files:

| File | Line | What to Change |
|------|------|----------------|
| `.astropulse/profile.yaml` | 5 | `clusterName` - Your cluster name |
| `.astropulse/build-app.yaml` | 11 | `repoURL` - Your forked repository |
| `.astropulse/deploy-app.yaml` | 13-14 | `registry` and `repository` - Your AWS ECR |
| `.astropulse/deploy-app.yaml` | 20 | `dnsName` - Your domain |
| `.astropulse/resources/kpack-image.yaml` | 13 | `url` - Your forked repository |
| `.astropulse/resources/kpack-image.yaml` | 19 | `tag` - Your AWS ECR |

## Architecture

```
┌─────────────┐      ┌──────────────┐      ┌─────────────────┐
│   Browser   │─────▶│  Go Server   │─────▶│   Kubernetes    │
│   (UI)      │◀─────│  (Port 8080) │◀─────│   (AstroPulse)  │
└─────────────┘      └──────────────┘      └─────────────────┘
                            │
                            ▼
                     ┌──────────────┐
                     │  kpack Build │
                     │   (ECR)      │
                     └──────────────┘
```

**How it works:**
1. **Build**: kpack watches Git repo, builds container image, pushes to ECR
2. **Deploy**: AstroPulse creates Kubernetes Deployment, Service, and Ingress
3. **Access**: Automatic HTTPS via cert-manager + DNS via external-dns

## API Endpoints

- `GET /` - Web UI
- `GET /ping-latency` - Latency measurement endpoint (returns JSON)
- `GET /health` - Health check for Kubernetes probes

## GitHub Actions (Optional)

This repository includes an **example workflow** (`.github/workflows/deploy.yml`) for automated deployments to AstroPulse.

**This workflow is disabled by default** and serves as a reference for users who fork the repository.

### To use automated deployment:

1. **Follow the complete setup guide:** [astropulse.io/get-started](https://astropulse.io/get-started)
2. **Fork this repository**
3. **Update configuration files** (see Configuration section above)
4. **Add GitHub Secret:** `ASTROPULSE_API_KEY`
5. **Enable the workflow** by editing `.github/workflows/deploy.yml`:
   ```yaml
   # Change from:
   on:
     workflow_dispatch:

   # To:
   on:
     push:
       branches: [main]
   ```

The workflow will then automatically deploy your app to AstroPulse on every push to main.

## Tech Stack

- **Backend**: Go 1.21
- **Frontend**: Vanilla JavaScript + CSS
- **Build**: kpack (Cloud Native Buildpacks)
- **Deploy**: AstroPulse platform
- **Container Registry**: AWS ECR
- **Infrastructure**: Kubernetes

## Development

```bash
# Run locally
go run main.go

# Build binary
make build

# Build Docker image (if needed)
docker build -t latency-app .
```

## Learn More

- [AstroPulse Documentation](https://docs.astropulse.io)
- [Get Started Guide](https://astropulse.io/get-started)
- [kpack Documentation](https://github.com/buildpacks-community/kpack)

## License

Apache 2.0
