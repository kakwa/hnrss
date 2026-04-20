#!/bin/sh

cd "$(dirname $0)/.."

pwd
image_repo="hnrss"
IMAGE_TAG="latest"


if [ -z "${REGISTRY_USER}" ] || \
    [ -z "${REGISTRY_PASSWORD}" ] || \
    [ -z "${REGISTRY_HOST}" ];
then
    echo "Please set the following env vars:\n* REGISTRY_USER\n* REGISTRY_PASSWORD\n* REGISTRY_HOST"
fi

echo "${REGISTRY_PASSWORD}" | docker login "${REGISTRY_HOST}" -u "${REGISTRY_USER}" --password-stdin

ref="${REGISTRY_HOST}/${image_repo}:${IMAGE_TAG}"


docker build ${DOCKER_EXTRA_BUILD_ARGS:-} \
	--build-arg "VERSION=${IMAGE_TAG}" \
	--build-arg "CONTAINER_NAME=${image_repo}" \
	-t "${ref}" .

docker push "${ref}"
echo "Pushed ${ref}"
