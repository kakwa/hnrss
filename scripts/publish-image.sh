#!/usr/bin/env bash
# Build and push the container image to a registry.
#
# Required:
#   REGISTRY_HOST   Registry hostname for docker login and image prefix
#                   (e.g. ghcr.io, docker.io, registry.example.com)
#
# Optional:
#   REGISTRY_USER / REGISTRY_PASSWORD
#                   When both are set, runs docker login with basic auth
#                   (password-stdin). Omit both to skip login.
#   CONTAINER_NAME  Repository path under the registry (default: hnrss/hnrss-ai-filtering)
#   IMAGE_NAME      Alias for CONTAINER_NAME when CONTAINER_NAME is unset
#   IMAGE_TAG       Image tag (default: git describe --tags --always, else latest)
#   BUILD_CONTEXT   Docker build context directory (default: repo root)
#   DOCKER_EXTRA_BUILD_ARGS
#                   Extra args passed to docker build (quoted string)

set -euo pipefail

root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$root"

: "${REGISTRY_HOST:?REGISTRY_HOST is required (registry hostname, e.g. ghcr.io)}"

image_repo="${CONTAINER_NAME:-${IMAGE_NAME:-hnrss/hnrss-ai-filtering}}"
IMAGE_TAG="${IMAGE_TAG:-}"
if [[ -z "${IMAGE_TAG}" ]]; then
	IMAGE_TAG="$(git describe --tags --always 2>/dev/null || true)"
fi
if [[ -z "${IMAGE_TAG}" ]]; then
	IMAGE_TAG="latest"
fi

BUILD_CONTEXT="${BUILD_CONTEXT:-$root}"

if [[ -n "${REGISTRY_USER:-}" || -n "${REGISTRY_PASSWORD:-}" ]]; then
	: "${REGISTRY_USER:?REGISTRY_USER is required when using registry credentials}"
	: "${REGISTRY_PASSWORD:?REGISTRY_PASSWORD is required when using registry credentials}"
	echo "${REGISTRY_PASSWORD}" | docker login "${REGISTRY_HOST}" -u "${REGISTRY_USER}" --password-stdin
fi

ref="${REGISTRY_HOST}/${image_repo}:${IMAGE_TAG}"
# shellcheck disable=SC2086
docker build ${DOCKER_EXTRA_BUILD_ARGS:-} \
	--build-arg "VERSION=${IMAGE_TAG}" \
	--build-arg "CONTAINER_NAME=${image_repo}" \
	-t "${ref}" \
	-f "${root}/Dockerfile" \
	"${BUILD_CONTEXT}"

docker push "${ref}"
echo "Pushed ${ref}"
