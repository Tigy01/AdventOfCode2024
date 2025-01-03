package day3

import (
	"aoc/internal/day1"
	"aoc/internal/filemanager"
	"fmt"
	"strconv"
	"strings"
)

func Run() (err error) {
	data, err := filemanager.ReadFullFile("./internal/day3/realInput.txt")
	if err != nil {
		return err
	}

	var total = 0
	var do_total = 0

	sections, err := getValidSections(data)
	if err != nil {
		return err
	}

	operations, err := getOps(sections)
	if err != nil {
		return err
	}

	do_characters, err := getDoSections(data)
	if err != nil {
		return err
	}

	do_sections, err := getValidSections(do_characters)
	if err != nil {
		return err
	}

	do_operations, err := getOps(do_sections)
	if err != nil {
		return err
	}

	for _, op := range operations {
		total += op[0] * op[1]
	}

	for _, op := range do_operations {
		do_total += op[0] * op[1]
	}

	fmt.Println("Total:", total)
	fmt.Println("Do Total:", do_total)
	return nil
}

func getDoSections(data string) (string, error) {
	var do_chunks = []string{}
	var char_count = 0
	var dont_sections = strings.Split(data, "don't()")
	var first = dont_sections[0]
	do_chunks = append(do_chunks, first)
	char_count += len(first)
	for _, dont_section := range dont_sections {
		for i := range dont_section {
			if i+4 >= len(dont_section) {
				continue
			}
			if dont_section[i:i+4] == "do()" {
				do_chunks = append(do_chunks, dont_section[i+4:])
				char_count += len(dont_section[i+4:])
				break
			}
		}
	}
	var index = 0
	var do_chars = make([]rune, char_count)
	for _, value := range do_chunks {
		for _, char := range value {
			if index >= len(do_chars) {
				return string(do_chars), nil
			}
			do_chars[index] = char
			index += 1
		}
	}
	return string(do_chars), nil
}

// fn getDoSections(data: []const u8, alloc: Allocator) ![]const u8 {
//
//	   const first = dont_sections.first();
//	   try do_chunks.append(first);
//	   char_count += first.len;
//
//	   while (dont_sections.next()) |dont_section| {
//	       for (dont_section, 0..) |_, i| {
//	           if (i + 4 >= dont_section.len) continue;
//	           if (std.mem.eql(u8, dont_section[i .. i + 4], "do()")) {
//	               try do_chunks.append(dont_section[i + 4 ..]);
//	               char_count += dont_section[i + 4 ..].len;
//	               break;
//	           }
//	       }
//	   }
//	   var index: usize = 0;
//	   var do_chars = try alloc.alloc(u8, char_count);
//	   for (do_chunks.items) |value| {
//	       for (value) |char| {
//	           if (index >= do_chars.len) return do_chars;
//	           do_chars[index] = char;
//	           index += 1;
//	       }
//	   }
//	   return do_chars;
//	}
func getValidSections(chars string) (sections []string, err error) {
	var mul_sections = strings.Split(chars, "mul(")
	for _, section := range mul_sections {
		var num_section = strings.Split(section, ")")
		for _, nums := range num_section {
			if !validateSection(nums) {
				continue
			}
			sections = append(sections, nums)
		}
	}
	return sections, nil
}

func getOps(sections []string) (operations [][2]int, err error) {
	for _, section := range sections {
		var current_op = [2]int{}
		var current_num = [3]rune{}
		var num_size = 0
		for i, char := range section {
			if day1.IsNum(char) {
				current_num[num_size] = char
				num_size += 1
				if i == len(section)-1 {
					result, err := strconv.ParseInt(string(current_num[0:num_size]), 10, 64)
					if err != nil {
						return nil, err
					}
					current_op[1] = int(result)
					operations = append(operations, current_op)
				}
			} else if char == ',' {
				if num_size == 0 {
					break
				}
				result, err := strconv.ParseInt(string(current_num[0:num_size]), 10, 64)
				if err != nil {
					return nil, err
				}
				current_op[0] = int(result)
				num_size = 0
			}
		}

	}
	return operations, nil
}

func validateSection(line string) bool {
	for i, char := range line {
		if day1.IsNum(char) {
			continue
		}
		if char == ',' && i != 0 {
			continue
		}
		return false
	}
	return true
}
