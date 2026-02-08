package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func writeTypes() error {
	content := `
export interface Question {
  id: number;
  question: string;
  explanation: string;
  example: string;
  category: string;
  level: 'senior';
  eliminatory: boolean;
}
`
	return os.WriteFile("output/types.ts", []byte(content), 0644)
}

func writeTS(filename, varName string, questions []Question) error {
	data, _ := json.MarshalIndent(questions, "", "  ")

	content := fmt.Sprintf(`
import { Question } from './types';

export const %s: Question[] = %s;
`, varName, string(data))

	return os.WriteFile(filename, []byte(content), 0644)
}
