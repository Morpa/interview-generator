package main

import (
	"bytes"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/google/uuid"
)

func parseMarkdown(content string, category string) []Question {
	md := []byte(content)
	root := markdown.Parse(md, nil)

	var questions []Question

	// função recursiva para percorrer o AST
	var walk func(node ast.Node)
	walk = func(node ast.Node) {
		switch n := node.(type) {

		case *ast.List:
			for _, item := range n.GetChildren() {
				if li, ok := item.(*ast.ListItem); ok {
					text := getText(li)
					if text != "" {
						example := getCodeBlock(li)
						q := normalizeQuestion(text)
						exp := buildSeniorExplanation(q)
						questions = append(questions, Question{
							ID:          uuid.New(),
							Question:    q,
							Explanation: normalizeExplanation(exp),
							Example:     example,
							Category:    Category(category),
						})
					}
				}
			}

		case *ast.Heading:
			if n.Level >= 2 {
				text := string(n.Literal)
				if text != "" {
					q := normalizeQuestion(text)
					exp := buildSeniorExplanation(q)
					questions = append(questions, Question{
						ID:          uuid.New(),
						Question:    q,
						Explanation: normalizeExplanation(exp),
						Example:     "", // exemplos podem vir nos próximos nós
						Category:    Category(category),
					})
				}
			}
		}

		for _, c := range node.GetChildren() {
			walk(c)
		}
	}

	walk(root)
	return questions
}

// pega o texto completo de um nó de lista
func getText(node ast.Node) string {
	var buf bytes.Buffer
	ast.WalkFunc(node, func(n ast.Node, entering bool) ast.WalkStatus {
		if lit, ok := n.(*ast.Text); ok && entering {
			buf.Write(lit.Literal)
		}
		return ast.GoToNext
	})
	return strings.TrimSpace(buf.String())
}

// pega o primeiro bloco de código dentro do nó
func getCodeBlock(node ast.Node) string {
	var buf bytes.Buffer
	ast.WalkFunc(node, func(n ast.Node, entering bool) ast.WalkStatus {
		if code, ok := n.(*ast.CodeBlock); ok && entering {
			buf.WriteString("```" + string(code.Info) + "\n")
			buf.Write(code.Literal)
			buf.WriteString("\n```")
			return ast.Terminate
		}
		return ast.GoToNext
	})
	return buf.String()
}

// import (
// 	"regexp"
// 	"strings"
// )

// func parseMarkdown(content string, category string) []Question {
// 	blocks := strings.Split(content, "\n## ")
// 	questions := []Question{}

// 	codeRegex := regexp.MustCompile("```[\\s\\S]*?```")

// 	for i, block := range blocks[1:] {
// 		lines := strings.Split(block, "\n")
// 		rawQuestion := strings.TrimSpace(lines[0])

// 		example := ""
// 		if match := codeRegex.FindString(block); match != "" {
// 			example = match
// 		}

// 		q := normalizeQuestion(rawQuestion)
// 		exp := buildSeniorExplanation(q)

// 		questions = append(questions, Question{
// 			ID:          i + 1,
// 			Question:    q,
// 			Explanation: normalizeExplanation(exp),
// 			Example:     example,
// 			Category:    category,
// 			Level:       "senior",
// 			Eliminatory: isEliminatory(q),
// 		})
// 	}

// 	return questions
// }
