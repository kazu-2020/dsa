# syntax=docker/dockerfile:1

# see: https://img.atcoder.jp/file/language-update/language-list.html
ARG GO_VERSION=1.20.6
FROM golang:${GO_VERSION}-bullseye

SHELL ["/bin/bash", "-o", "pipefail", "-c"]
WORKDIR /workspace

ENV PATH=/root/.local/bin:/root/.volta/bin:$PATH
RUN <<EOF
  curl https://get.volta.sh | bash
  volta install node
EOF
RUN --mount=type=cache,target=/root/.npm <<EOF
  npm install -g atcoder-cli
EOF

RUN --mount=type=cache,target=/var/lib/apt,sharing=locked \
   --mount=type=cache,target=/var/cache/apt,sharing=locked <<EOF
  apt-get update
  apt-get --no-install-recommends install -y python3 python3-pip
EOF
RUN --mount=type=cache,target=/root/.cache/pip <<EOF
  pip3 install --user online-judge-tools
EOF

# see: https://taskfile.dev/installation/#install-script
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b ~/.local/bin
