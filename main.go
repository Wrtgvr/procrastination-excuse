package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type WordGroups struct {
	Guilty     []string `json:"guilty"`
	Action     []string `json:"action"`
	ActionTool []string `json:"action_tool"`
}

func main() {
	wordGroupsJSON, err := os.Open("./word_groups.json")
	if err != nil {
		panic(err)
	}

	var wordG WordGroups
	err = json.NewDecoder(wordGroupsJSON).Decode(&wordG)
	if err != nil {
		panic(err)
	}

	s := fmt.Sprintf(
		"I procrastinated because %s %s %s.",
		wordG.Guilty[rand.Intn(len(wordG.Guilty))],
		wordG.Action[rand.Intn(len(wordG.Action))],
		wordG.ActionTool[rand.Intn(len(wordG.ActionTool))],
	)

	fmt.Println(s)
}
