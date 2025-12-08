# freeport

[![Go Version](https://img.shields.io/github/go-mod/go-version/artemshalin/freeport?style=flat-square&logo=go&logoColor=white&color=00ADD8)](https://github.com/artemshalin/freeport) [![Releases](https://img.shields.io/github/v/release/artemshalin/freeport?style=flat-square&logo=github&logoColor=white&color=181717)](https://github.com/artemshalin/freeport/releases) [![License](https://img.shields.io/github/license/artemshalin/freeport?style=flat-square&logo=mit&logoColor=white&color=00D084)](https://github.com/artemshalin/freeport/blob/main/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/artemshalin/freeport?style=flat-square)](https://goreportcard.com/report/github.com/artemshalin/freeport) [![Actions](https://img.shields.io/github/actions/workflow/status/artemshalin/freeport/release.yml?style=flat-square&logo=github-actions&logoColor=white&color=2088FF)](https://github.com/artemshalin/freeport/actions) ![Docker Pulls](https://img.shields.io/docker/pulls/artemshalin/freeport)

Утилита командной строки для получения свободного порта. Выводит номер свободного для использования TCP-порта в stdout.

## Установка

### Скачать готовый релиз

#### Linux amd64

```sh
curl -L -o freeport https://github.com/artemshalin/freeport/releases/latest/download/freeport_Linux_x86_64.tar.gz
tar xzf freeport_Linux_x86_64.tar.gz && chmod +x freeport
```

#### macOS arm64  

```ah
curl -L -o freeport https://github.com/artemshalin/freeport/releases/latest/download/freeport_Darwin_arm64.tar.gz
tar xzf freeport_Darwin_arm64.tar.gz && chmod +x freeport
```

### Docker

```sh
docker pull ghcr.io/artemshalin/freeport:latest
```

## Использование

### Исполняемый файл

```sh
freeport
```

### Docker

```sh
docker run --network=host ghcr.io/artemshalin/freeport:latest
```
