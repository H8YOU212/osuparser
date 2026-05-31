# osuparser

Go библиотека для парсинга osu! файлов (`.osu`, `.osz`, `.osr`)

## Быстрый старт

### 1. Импорт

```go
import "github.com/example/osuparser"
```

### 2. Создание парсера

```go
parser := osuparser.NewParser()
```

### 3. Парсинг файла

```go
result, err := parser.Parse("path/to/file.osz")
if err != nil {
    log.Fatal(err)
}
```

### 4. Работа с результатом

```go
// Для файлов .osz (бимапы)
if result.Osz != nil {
    fmt.Println("OSZ файл распарсен")
    fmt.Println("Содержимое:", result.Osz.Content)
    // Доступны распарсенные .osu файлы в result.Osz.osufiles
}

// Для файлов .osr (реплеи)
if result.Osr != nil {
    fmt.Println("OSR файл декодирован")
    fmt.Printf("Игрок: %s\n", result.Osr.Format.ReplayInfo.Username)
    fmt.Printf("Счет: %d\n", result.Osr.Format.Stats.TotalScore)
}
```

## Примеры

### Парсинг beatmap'а

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/example/osuparser"
)

func main() {
    parser := osuparser.NewParser()
    
    result, err := parser.Parse("maps/beatmap.osz")
    if err != nil {
        log.Fatal(err)
    }
    
    if result.Osz != nil {
        fmt.Printf("Карта распарсена\n")
        for _, osufile := range result.Osz.Osufiles {
            fmt.Printf("Файл: %s\n", osufile.Path)
            fmt.Printf("Название: %s\n", osufile.Osufile.Metadata.Title)
            fmt.Printf("Сложность: %.2f\n", osufile.Osufile.Difficulty.CS)
        }
    }
}
```

### Парсинг репея

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/example/osuparser"
)

func main() {
    parser := osuparser.NewParser()
    
    result, err := parser.Parse("replays/my_score.osr")
    if err != nil {
        log.Fatal(err)
    }
    
    if result.Osr != nil {
        fmt.Printf("Игрок: %s\n", result.Osr.Format.ReplayInfo.Username)
        fmt.Printf("Beatmap: %s\n", result.Osr.Format.ReplayInfo.BeatmapMD5)
        fmt.Printf("Счет: %d\n", result.Osr.Format.Stats.TotalScore)
        fmt.Printf("Combo: %d\n", result.Osr.Format.Stats.GreatestCombo)
    }
}
```

## API

### Package-level функции

```go
// NewParser создает новый парсер
func NewParser() *parser.Parser

// Parse парсит файл и возвращает результат
func Parse(path string) (*parser.ParseResult, error)
```

### Методы Parser

```go
// Parse парсит файл по пути
func (p *Parser) Parse(path string) (*ParseResult, error)
```

### ParseResult

```go
type ParseResult struct {
    Osz *osz.OszModel  // Для .osz файлов
    Osr *osr.OsrModel  // Для .osr файлов
}
```

## Поддерживаемые форматы

- `.osz` - архив с beatmap`s (ZIP)
- `.osr` - файл репроея
- `.osu` - файл beatmap (парсится как часть .osz)

## Обработка ошибок

```go
result, err := parser.Parse("file.osz")
if err != nil {
    switch err.Error() {
    case "path is empty":
        // Путь не указан
    case strings.Contains(err.Error(), "unsupported file extension"):
        // Неподдерживаемый формат файла
    default:
        log.Fatal(err)
    }
}
```

## Лицензия

MIT
