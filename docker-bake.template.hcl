group "default" {
  targets = ["multi"]
}

# Multi-platform build for pushing to registry
target "multi" {
  dockerfile = "Dockerfile"
  platforms = ["linux/amd64", "linux/arm64"]
  tags = ["astropulseinc/latency-app:{{TAG}}", "astropulseinc/latency-app:latest"]
}
