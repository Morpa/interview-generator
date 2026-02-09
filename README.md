# Interview Generator

Gerador de questões de entrevista para JavaScript e React a partir de repositórios públicos do GitHub.

## 📋 Sobre

Este projeto clona repositórios com questões de entrevista (JavaScript e React) e gera arquivos TypeScript com as questões estruturadas, prontas para uso em aplicações.

## 🚀 Como usar

```bash
# Executar o gerador
go run .
```

Os arquivos serão gerados na pasta `output/`:

- `js.questions.ts` - Questões de JavaScript
- `react.questions.ts` - Questões de React

## 📦 Estrutura do Projeto

```
├── main.go          # Ponto de entrada
├── models.go        # Definição dos tipos (Question, Category)
├── parser.go        # Parser de Markdown
├── utils.go         # Funções auxiliares
├── writer.go        # Gerador de arquivos TypeScript
├── repos/           # Repositórios clonados (gerado)
└── output/          # Arquivos de saída (gerado)
```

## 📝 Estrutura da Question

```go
type Question struct {
    ID          uuid.UUID `json:"id"`
    Question    string    `json:"question"`
    Explanation string    `json:"explanation"`
    Example     string    `json:"example"`
    Category    Category  `json:"category"` // "JS" ou "React"
}
```

## 🔗 Fontes

- [JavaScript Interview Questions](https://github.com/sudheerj/javascript-interview-questions)
- [React Interview Questions](https://github.com/sudheerj/reactjs-interview-questions)
