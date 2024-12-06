package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/tony-montemuro/aoc-2024/aocutils"
)

func updateRulesMap(key, value int, ruleMap map[int][]int) {
	if rules, exists := ruleMap[key]; exists {
		rules = append(rules, value)
		ruleMap[key] = rules
	} else {
		values := []int{value}
		ruleMap[key] = values
	}
}

func extractRulesAndUpdates(input []string) (map[int][]int, map[int][]int, [][]int) {
	isAllRulesSeen := false
	beforeRulesMap, afterRulesMap := make(map[int][]int), make(map[int][]int)
	updates := [][]int{}

	for _, row := range input {
		if row == "" {
			isAllRulesSeen = true
		} else if !isAllRulesSeen {
			parts := strings.Split(row, "|")
			n1, n2 := aocutils.Stoi(parts[0]), aocutils.Stoi(parts[1])
			updateRulesMap(n2, n1, beforeRulesMap)
			updateRulesMap(n1, n2, afterRulesMap)
		} else {
			nums := []int{}
			for _, n := range strings.Split(row, ",") {
				nums = append(nums, aocutils.Stoi(n))
			}
			updates = append(updates, nums)
		}
	}

	return beforeRulesMap, afterRulesMap, updates
}

func isUpdateValid(update []int, beforeRulesMap, afterRulesMap map[int][]int) bool {
	for i, n := range update {
		beforeRules, afterRules := beforeRulesMap[n], afterRulesMap[n]

		for j := 0; j < i; j++ {
			if !slices.Contains(beforeRules, update[j]) {
				return false
			}
		}

		for j := i + 1; j < len(update); j++ {
			if !slices.Contains(afterRules, update[j]) {
				return false
			}
		}
	}
	return true
}

func fixUpdate(update []int, beforeRulesMap, afterRulesMap map[int][]int) int {
	for i, n := range update {
		beforeRules, afterRules := beforeRulesMap[n], afterRulesMap[n]

		for j := 0; j < i; j++ {
			if !slices.Contains(beforeRules, update[j]) {
				update[i], update[j] = update[j], update[i]
				if !isUpdateValid(update, beforeRulesMap, afterRulesMap) {
					return fixUpdate(update, beforeRulesMap, afterRulesMap)
				}
			}
		}

		for j := i + 1; j < len(update); j++ {
			if !slices.Contains(afterRules, update[j]) {
				update[i], update[j] = update[j], update[i]
				if !isUpdateValid(update, beforeRulesMap, afterRulesMap) {
					return fixUpdate(update, beforeRulesMap, afterRulesMap)
				}
			}
		}
	}

	return update[len(update)/2]
}

func main() {
	input := aocutils.GetInputRows()
	beforeRulesMap, afterRulesMap, updates := extractRulesAndUpdates(input)
	p1, p2 := 0, 0
	for _, update := range updates {
		if isUpdateValid(update, beforeRulesMap, afterRulesMap) {
			p1 += update[len(update)/2]
		} else {
			p2 += fixUpdate(update, beforeRulesMap, afterRulesMap)
		}
	}
	fmt.Println(p1, p2)
}
