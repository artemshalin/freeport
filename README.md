# freeport

Утилита командной строки для получения свободного порта. Выводит номер свободного для использования TCP-порта в stdout. 

[![Go Version](https://img.shields.io/github/go-mod/go-version)]

[![Releases](https://img.shields.io/github/v/release/arthttps://github.com/artemshalin/freeport/releaseshttps://img.shields.io/github/license/art)]

[![Licence](https://img.shields.io/github/license/Ileriayo/markdown-badges?style=for-the-badge)](./LICENSE)

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
