### Tests and linter status:
[![HexletCheck Status](https://github.com/yanasirina/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/yanasirina/go-project-242/actions)

[![GolangciLint Status](https://github.com/yanasirina/go-project-242/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/yanasirina/go-project-242/actions)

[![TidyVendor Status](https://github.com/yanasirina/go-project-242/actions/workflows/tidy-vendor.yml/badge.svg)](https://github.com/yanasirina/go-project-242/actions)

[![Test Status](https://github.com/yanasirina/go-project-242/actions/workflows/test.yml/badge.svg)](https://github.com/yanasirina/go-project-242/actions)

## Документация

hexlet-path-size - это консольная утилита для вычисления размера файлов и каталогов.

## Возможности

- Вычисление размера отдельных файлов
- Вычисление размера каталогов (рекурсивно всего содержимого или же только файлов первого уровня вложенности)
- Возможность форматировать вывод (КБ, МБ, ГБ и т.д.)
- Возможность учитывать или не учитывать скрытые файлы

## Установка

### Предварительные требования

- Go 1.25 или выше

### Сборка из исходников

1. Клонируйте репозиторий
2. Перейдите в директорию проекта
3. Соберите приложение одним из следующих способов:

**Используя Make:**
```bash
make build
```

Это создаст бинарный файл в директории `bin/`.

**Используя Go напрямую:**
```bash
go build -o bin/hexlet-path-size ./cmd/hexlet-path-size
```

### Использование собранного бинарного файла

После сборки можно запустить:
```bash
./bin/hexlet-path-size [опции] <путь>
```

## Опции

- `-H`, `--human`: Отображать размеры в удобочитаемом формате (КБ, МБ, ГБ и т.д.)
- `-a`, `--all`: Включать скрытые файлы и каталоги в вычисления
- `-r`, `--recursive`: Вычислять размер рекурсивно для каталогов

## Примеры

Получить размер файла:
```bash
./bin/hexlet-path-size /path/to/file.txt
```

Получить размер каталога без рекурсии (сумму размеров файлов на первом уровне вложенности):
```bash
./bin/hexlet-path-size /path/to/directory
```

Получить размер файла в удобочитаемом формате:
```bash
./bin/hexlet-path-size --human /path/to/file.txt
# или
./bin/hexlet-path-size -H /path/to/file.txt
```

Получить размер каталога включая скрытые файлы:
```bash
./bin/hexlet-path-size --all /path/to/directory
# или
./bin/hexlet-path-size -a /path/to/directory
```

Получить рекурсивный размер каталога в удобочитаемом формате:
```bash
./bin/hexlet-path-size --human --recursive /path/to/directory
# или
./bin/hexlet-path-size -H -r /path/to/directory
```

Получить рекурсивный размер каталога включая скрытые файлы:
```bash
./bin/hexlet-path-size --all --recursive /path/to/directory
# или
./bin/hexlet-path-size -a -r /path/to/directory
```

Получить полный размер каталога со всеми включёнными опциями:
```bash
./bin/hexlet-path-size --human --all --recursive /path/to/directory
# или
./bin/hexlet-path-size -H -a -r /path/to/directory
```

## Вывод

Инструмент выводит путь, за которым следует вычисленный размер:
```
/path/to/file.txt - 1024B
```

С флагом `--human`:
```
/path/to/file.txt - 1KB
```

## Разработка

### Запуск тестов

```bash
make test
```

### Линтинг

Для проверки качества кода:
```bash
make lint
```

Для исправления проблем линтинга:
```bash
make lint-fix
```

### Обновление зависимостей

```bash
make tidy-vendor
```
