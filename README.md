LogLint

Простой линтер для Go, который проверяет ваши лог-сообщения:

Начинаются с маленькой буквы

Только на английском

Без спецсимволов и эмодзи

Без чувствительных данных (password, token, apiKey, secret)

Установка

Собираем бинарник линтера:

go build -o loglint ./cmd/loglint

Проверяем, что работает:

./loglint ./example

Использование с golangci-lint

Добавляем в .golangci.yml:

linters:
enable:
- custom-loglint

linters-settings:
custom:
custom-loglint:
path: ./loglint
description: "Checks log messages"

Запускаем проверку на проекте:

golangci-lint run