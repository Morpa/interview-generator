package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func readMarkdownFiles(root string) ([]string, error) {
	files := []string{}

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".md" {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func main() {
	os.MkdirAll("repos", 0755)
	os.MkdirAll("output", 0755)

	fmt.Println("📥 Cloning repositories...")
	cloneRepo(
		"https://github.com/sudheerj/reactjs-interview-questions.git",
		"repos/reactjs",
	)
	cloneRepo(
		"https://github.com/sudheerj/javascript-interview-questions.git",
		"repos/javascript",
	)

	fmt.Println("⚙️ Parsing JavaScript...")
	jsFiles, _ := readMarkdownFiles("repos/javascript")
	jsQuestions := []Question{}

	for _, file := range jsFiles {
		content, _ := os.ReadFile(file)
		category := filepath.Base(file)
		jsQuestions = append(jsQuestions, parseMarkdown(string(content), category)...)
	}

	fmt.Println("⚙️ Parsing React...")
	reactFiles, _ := readMarkdownFiles("repos/reactjs")
	reactQuestions := []Question{}

	for _, file := range reactFiles {
		content, _ := os.ReadFile(file)
		category := filepath.Base(file)
		reactQuestions = append(reactQuestions, parseMarkdown(string(content), category)...)
	}

	writeTypes()
	writeTS("output/js.questions.ts", "jsQuestions", jsQuestions)
	writeTS("output/react.questions.ts", "reactQuestions", reactQuestions)

	fmt.Println("✅ Done!")
}
