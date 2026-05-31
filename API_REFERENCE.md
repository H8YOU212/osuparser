# API Библиотеки osuparser

## Твои требования ✅

```
1. go get "github/example/myparser"
2. import this
3. osuparser := osuparser.NewParser()
4. osuparser.Parse(path/to/osr?osz(любой))
5. давать возвращаемое значение
```

## Реализация

### 1. Установка

```bash
go get github.com/example/osuparser
```

### 2. Импорт

```go
import "github.com/example/osuparser"
```

### 3. Создание парсера

```go
parser := osuparser.NewParser()
```

### 4. Парсинг файла (auto-detect формата)

```go
// Автоматически определяет .osu, .osz или .osr
result, err := parser.Parse("path/to/file.osz")
if err != nil {
    log.Fatal(err)
}
```

### 5. Возвращаемое значение

```go
type ParseResult struct {
    Osz *osz.OszModel  // nil если не .osz файл
    Osr *osr.OsrModel  // nil если не .osr файл
}
```

## Полный пример

```go
package main

import (
    "fmt"
    "log"
    "github.com/example/osuparser"
)

func main() {
    // Создание парсера
    parser := osuparser.NewParser()

    // Парсинг любого типа файла (определяется по расширению)
    result, err := parser.Parse("beatmap.osz")
    if err != nil {
        log.Fatal(err)
    }

    // Проверка результата
    if result.Osz != nil {
        fmt.Printf("OSZ распарсена, файлов: %d\n", len(result.Osz.Osufiles))
        for _, file := range result.Osz.Osufiles {
            fmt.Printf("  - %s: %s\n", file.Path, file.Osufile.Metadata.Title)
        }
    }

    if result.Osr != nil {
        fmt.Println("OSR успешно декодирована")
    }
}
```

## Альтернативный способ (Package-level)

Если не хочешь создавать Parser напрямую:

```go
result, err := osuparser.Parse("path/to/file.osz")
// То же самое, что:
// parser := osuparser.NewParser()
// result, err := parser.Parse("path/to/file.osz")
```

## Поддерживаемые форматы

| Расширение | Тип | Возвращается в |
|-----------|------|-------------------|
| `.osz` | Beatmap архив (ZIP) | `result.Osz` |
| `.osr` | Файл репроея | `result.Osr` |
| `.osu` | Beatmap файл | Часть `.osz` |

## Обработка ошибок

```go
result, err := parser.Parse("file.osz")
if err != nil {
    switch {
    case err.Error() == "path is empty":
        fmt.Println("Путь не указан")
    case strings.Contains(err.Error(), "unsupported file extension"):
        fmt.Println("Неподдерживаемый формат файла")
    default:
        log.Fatal(err)
    }
}
```

## API функции Parse()

```go
func (p *Parser) Parse(path string) (*ParseResult, error)
```

**Параметры:**
- `path string` - путь к файлу (.osz, .osr или .osu)

**Возвращает:**
- `*ParseResult` - структура с результатом парсинга
- `error` - ошибка если что-то пошло не так

**Поддерживаемые форматы:**
- Определяет тип файла по расширению
- Автоматически распаковывает `.osz`
- Декодирует данные `.osr`
- Парсит `.osu` файлы

## Статус структур

Только публичные поля (`Osufiles`, `Content` и т.д.):

```go
type OszModel struct {
    Osufiles []osu.Commonosu  // Публичное
    Content  []string         // Публичное
    // ... приватные поля
}
```
