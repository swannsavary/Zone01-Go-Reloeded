# Go-Reloaded

A simple Go tool that edits and corrects text automatically.

It does **not** fix spelling mistakes. It only handles formatting and the rules below.

## Features

- Fixes punctuation spacing (`.`, `,`, `!`, `?`, `:`, `;`)
- Removes extra spaces
- Fixes quotation marks: removes spaces inside `' '`
- Changes "a" to "an" before a word starting with a vowel or an "h"
- Applies modifiers to the preceding word(s):

| Modifier | Effect |
|----------|--------|
| `(cap)`  | Capitalizes the word |
| `(up)`   | Converts the word to uppercase |
| `(low)`  | Converts the word to lowercase |
| `(hex)`  | Converts a hexadecimal number to decimal |
| `(bin)`  | Converts a binary number to decimal |

Modifiers `(up)`, `(low)` and `(cap)` accept an optional number to apply the change to the previous N words, e.g. `(cap, 6)`.

## Requirements

- Go 1.22.2 or later

## Usage

```bash
go run . <input.txt> <output.txt>
```

## Examples

(garde tes deux exemples tels quels)

## Project structure

```
.
├── formatage
│   ├── An.go
│   ├── apostrophe.go
│   ├── formatagecomplet.go
│   ├── formatage.go
│   └── modif_mot_with_para.go
├── go.mod
├── main.go
├── main_test.go
├── README.md
├── result.txt
├── sample.txt
└── tools
    ├── parametres.go
    └── utils.go
```

## Tests

```bash
go test ./...
```

## Author

swannsavary