package day5

import (
	"aoc/internal/filemanager"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type node struct {
	value int
	after []int
}

func (node) new(value int, after_values ...int) node {
	if len(after_values) == 0 {
		return node{value, []int{}}
	}
	return node{value, after_values}
}

func Run() (err error) {
	lines, err := filemanager.ReadLines("./internal/day5/realInput.txt")
	if err != nil {
		return err
	}

	var rules = getRules(lines)
	nodes, err := getNodesFromRuleset(rules)
	if err != nil {
		return err
	}

	var updates = lines[len(rules)+1:]
	var total = 0
	var fixed_total = 0
	for _, update := range updates {
		middle, err := getMiddleOfUpdate(false, update, nodes)
		if err != nil {
			return err
		}
		total += middle
		fixed_middle, err := getMiddleOfUpdate(true, update, nodes)
		if err != nil {
			return err
		}
		fixed_total += fixed_middle
	}
	fmt.Printf("Output - %v\n", total)
	fmt.Printf("Fixed Output - %v\n", fixed_total)
	return nil
}

func sortUpdate(keys []int, nodes []*node) (sorted_keys []int) {
	var key_nodes = []*node{}
	for _, key := range keys {
		for _, current := range nodes {
			if current.value == key {
				key_nodes = append(key_nodes, current)
				break
			}
		}
	}

	var start = 0
	var sorted = false
	for !sorted {
		sorted = true
		var max_score_index = 0
		var max_score = 0
		for i, current := range key_nodes[start:] {
			var score = 0
			for _, next := range key_nodes {
				if slices.Contains(current.after, next.value) {
					score += 1
				}
			}

			if i == 0 {
				max_score = score
				continue
			}

			if score > max_score {
				max_score = score
				max_score_index = i
				sorted = false
			}

		}
		var temp = key_nodes[start]
		key_nodes[start] = key_nodes[start:][max_score_index]
		key_nodes[start:][max_score_index] = temp
		start += 1
		if start == len(key_nodes) {
			start = 0
		} else {
			sorted = false
		}
	}

	for _, n := range key_nodes {
		sorted_keys = append(sorted_keys, n.value)
	}
	return sorted_keys
}

func getMiddleOfUpdate(fix_values bool, update string, nodes []*node) (int, error) {
	var value_iterator = strings.Split(update, ",")
	var keys = []int{}

	for _, value := range value_iterator {
		var key, err = strconv.ParseInt(value, 10, 64)
		if err != nil {
			return 0, err
		}
		keys = append(keys, int(key))
	}

	var valid = true
	for i, key := range keys {
		var node_match *node = nil
		for _, current := range nodes {
			if current.value == key {
				node_match = current
				break
			}
		}
		if node_match == nil {
			return 0, nil
		}
		for _, after_key := range keys[i+1:] {
			var found = false
			for _, after_value := range node_match.after {
				if after_value == after_key {
					found = true
					break
				}
			}
			if !found {
				valid = false
			}
		}
	}

	if fix_values {
		if valid {
			return 0, nil
		}
		keys = sortUpdate(keys, nodes)
	} else {

		if valid {
			return keys[len(keys)/2], nil
		}
		return 0, nil
	}
	return keys[len(keys)/2], nil
}
func getRules(lines []string) []string {
	for i, line := range lines {
		if len(line) == 0 {
			return lines[0:i]
		}
	}
	return lines
}

func getNodesFromRuleset(rules []string) (nodes []*node, err error) {
	for _, rule := range rules {
		var split_iterator = strings.Split(rule, "|")

		firstNum64, err := strconv.ParseInt(split_iterator[0], 10, 64)
		if err != nil {
			return []*node{}, err
		}

		secondNum64, err := strconv.ParseInt(split_iterator[1], 10, 64)
		if err != nil {
			return []*node{}, err
		}
		firstNum := int(firstNum64)
		secondNum := int(secondNum64)

		var first_found = false
		var second_found = false
		for _, current := range nodes {
			if current.value == firstNum {
				current.after = append(current.after, secondNum)
				first_found = true
				break
			} else if current.value == secondNum {
				second_found = true
			}
		}
		if !first_found {
			var new_node = node{}.new(firstNum, secondNum)
			nodes = append(nodes, &new_node)
		}
		if !second_found {
			var new_node = node{}.new(secondNum, []int{}...)
			nodes = append(nodes, &new_node)
		}
	}
	return nodes, nil
}
