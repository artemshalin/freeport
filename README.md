# freeport
This is a command-line utility that gets an unoccupied port and prints it to stdout.

Утилита получает не занятый порт и пишет его в `stdout`. Используется для получения свободных портов, при создании окружений для тестов.

## Как пользоваться

Скачать последний релиз утилиты (например, для linux):

```sh
curl -L "https://gitlab.levelgroup.ru/api/v4/projects/104/releases/latest/downloads/binaries/freeport_linux_amd64" --output freeport
```

Сделать файл исполняемым:

```sh
chmod +x ./freeport
```

Вызвать утилиту из командной строки.

```sh
./freeport
```

В результате утилита покажет порт, которые никем не занят

```sh
39615
```