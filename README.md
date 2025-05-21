# go-http-client-benchmark

Инструмент для сравнения производительности HTTP-клиентов `net/http` и `fasthttp` в последовательном и параллельном режимах.

## Возможности

- Последовательный (`--mode seq`) и параллельный (`--mode par`) режимы
- Подсчёт количества успешных и неуспешных запросов
- Расчёт общего времени выполнения и среднего времени на запрос

## Установка и запуск

### 1. Клонирование репозитория и установка зависимости

```bash
git clone https://github.com/Deymos01/go-http-client-benchmark.git
cd go-http-client-benchmark
go mod tidy
```

### 2. Запуск тестового сервера

```bash
go run server/server.go
```
Сервер будет доступен по адресу: http://localhost:8080/ping

### 3. Запуск бенчмарка

```bash
go run main.go [flags]
```

#### Доступные флани

| Флаг      | Значение по умолчанию        | Описание                              |
| --------- | ---------------------------- | ------------------------------------- |
| `-n`      | `1000`                       | Количество HTTP-запросов              |
| `-url`    | `http://localhost:8080/ping` | Целевой URL                           |
| `-mode`   | `seq`                        | Режим тестирования: `seq` или `par`   |
| `-client` | `nethttp`                    | HTTP-клиент: `nethttp` или `fasthttp` |


### Пример запуска

Последовательный запуск с использованием net/http:
```bash
go run main.go -n 1000 -client nethttp -mode seq
```

Параллельный запуск с использованием fasthttp:
```bash
go run main.go -n 1000 -client fasthttp -mode par
```

Пример вывода:
```cmd
[Bench client fasthttp (Parallel mode)] Completed 910 requests in 67.329ms (Failed: 90, Avg Time: 67.329µs)
Success Rate: 91.00%
```