package day2

import (
	"aoc/internal/day1"
	"aoc/internal/filemanager"
	"fmt"
	"strconv"
)

func Run() error {
	lines, err := filemanager.ReadLines("./internal/day2/realInput.txt")

	num_safe := 0
	if err != nil {
		return err
	}
	for _, line := range lines {
		var num_list, err = getNumbersFromLine(line)
		if err != nil {
			return err
		}
		if isSafe(num_list) {
			num_safe += 1
		}
	}
	fmt.Println("Number Safe:", num_safe)
	return nil
}

const INCREASING = 0
const DECREASING = 1
const UNKNOWN = 2

func isSafe(num_list []int) bool {
	if len(num_list) <= 1 {
		return false
	}

	var change = UNKNOWN
	for i, value := range num_list {
		if i == len(num_list)-1 {
			return true
		}

		var next_value = num_list[i+1]
		var delta = next_value - value
		var current_change = UNKNOWN
		if delta > 0 {
			current_change = INCREASING
		} else {
			current_change = DECREASING
		}

		var distance = delta
		if delta < 0 {
			distance *= -1
		}

		if distance < 1 || distance > 3 {
			return false
		}

		if change == UNKNOWN {
			change = current_change
		} else if current_change != change {
			return false
		}
	}
	return false
}

//
//
//    }
//    unreachable;
//}
//
//fn isSafeWithDampen(num_list: []i32, dampened: bool) Allocator.Error!bool {
//    if (num_list.len <= 1) return false;
//
//    var change = CHANGE.UNKNOWN;
//    for (num_list, 0..) |value, i| {
//        if (i == num_list.len - 1) return true;
//
//        const next_value = num_list[i + 1];
//        const delta = next_value - value;
//        const current_change = if (delta > 0) CHANGE.INCREASING else CHANGE.DECREASING;
//
//        const distance = if (delta < 0) delta * -1 else delta;
//
//        if (distance < 1 or distance > 3) {
//            if (dampened) return false;
//
//            if (try dampen_list(num_list, i)) return true;
//            return dampen_list(num_list, i + 1);
//        }
//
//        if (change == CHANGE.UNKNOWN) {
//            change = current_change;
//        } else if (current_change != change) {
//            if (dampened) return false;
//
//            if (try dampen_list(num_list, i)) return true;
//            if (try dampen_list(num_list, i + 1)) return true;
//            if (i == 1) return dampen_list(num_list, i - 1);
//
//            return false;
//        }
//    }
//    unreachable;
//}
//
//fn dampen_list(num_list: []i32, i: usize) Allocator.Error!bool {
//    var buff: [4 * 32]u8 = undefined;
//    var fba = std.heap.FixedBufferAllocator.init(&buff);
//    const alloc = fba.allocator();
//    var dampened_list = ArrayList(i32).init(alloc);
//    try dampened_list.appendSlice(num_list);
//    _ = dampened_list.orderedRemove(i);
//    const dampened_slice = try dampened_list.toOwnedSlice();
//    const safe_with_dampen = isSafeWithDampen(dampened_slice, true);
//    alloc.free(dampened_slice);
//    return safe_with_dampen;
//}

func getNumbersFromLine(line string) (nums []int, err error) {
	var num_buff = []rune{}
	for i, char := range line {
		if day1.IsNum(char) {
			num_buff = append(num_buff, char)
			if len(line)-1 != i {
				continue
			}
		}

		if string(num_buff) == "" {
			continue
		}
		var val, err = strconv.ParseInt(string(num_buff), 10, 64)
		if err != nil {
			return []int{}, err
		}

		nums = append(nums, int(val))
        num_buff = []rune{}
	}

	return nums, nil

}
