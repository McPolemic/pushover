#!/bin/bash
docker run --rm -it -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang ./inside_docker_build.sh
