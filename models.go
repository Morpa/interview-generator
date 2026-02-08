package main

type Question struct {
	ID          int    `json:"id"`
	Question    string `json:"question"`
	Explanation string `json:"explanation"`
	Example     string `json:"example"`
	Category    string `json:"category"`
	Level       string `json:"level"`
	Eliminatory bool   `json:"eliminatory"`
}
