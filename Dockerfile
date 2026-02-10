# SPDX-FileCopyrightText: Copyright (c) 2024-2026 Almaz Ilaletdinov <a.ilaletdinov@yandex.ru>
# SPDX-License-Identifier: MIT

FROM golang:1.26.0 AS base
WORKDIR /app
ENV EC_VERSION="v3.0.3"
ENV GOLANGCI_LINT_VERSION="v2.0.2"
ENV YAMLLINT_VERSION="1.35.1"
ENV PATH="/root/.local/bin:$PATH"
RUN apt-get update && apt-get install -y \
    python3 \
    python3-pip \
    python3-venv \
    curl \
    && rm -rf /var/lib/apt/lists/*
RUN pip install yamllint==$YAMLLINT_VERSION --break-system-packages
RUN curl -LSfs https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
    | sh -s -- -b /usr/local/bin $GOLANGCI_LINT_VERSION
RUN curl -LSfs https://astral.sh/uv/install.sh | sh
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /bin
RUN curl -O -L -C - https://github.com/editorconfig-checker/editorconfig-checker/releases/download/$EC_VERSION/ec-linux-amd64.tar.gz && \
    tar xzf ec-linux-amd64.tar.gz -C /tmp && \
    mv /tmp/bin/ec-linux-amd64 /root/.local/bin/ec
COPY . .
RUN task install-deps
