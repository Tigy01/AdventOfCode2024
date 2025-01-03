const std = @import("std");
const io = std.io;
const ArrayList = std.ArrayList;
const Allocator = std.mem.Allocator;
const day1 = @import("day1.zig");
const FileManager = @import("file_manager.zig");

pub fn run(allocator: Allocator) !void {
    var arena = std.heap.ArenaAllocator.init(allocator);
    var alloc = arena.allocator();
    defer arena.deinit();

    const data = try FileManager.readFullFile("realInput.txt", alloc);
    defer alloc.free(data);
    var total: i32 = 0;
    const do_characters = try getDoSections(data, alloc);
    const sections = try getValidSections(do_characters, alloc);
    defer alloc.free(sections);
    const operations = try getOps(sections, alloc);
    defer alloc.free(operations);
    for (operations) |operation| {
        total += operation[0] * operation[1];
    }
    std.debug.print("{}\n", .{total});
}


fn getDoSections(data: []const u8, alloc: Allocator) ![]const u8 {
    var do_chunks = ArrayList([]const u8).init(alloc);
    defer do_chunks.deinit();
    var char_count: usize = 0;

    var dont_sections = std.mem.split(u8, data, "don't()");

    const first = dont_sections.first();
    try do_chunks.append(first);
    char_count += first.len;

    while (dont_sections.next()) |dont_section| {
        for (dont_section, 0..) |_, i| {
            if (i + 4 >= dont_section.len) continue;
            if (std.mem.eql(u8, dont_section[i .. i + 4], "do()")) {
                try do_chunks.append(dont_section[i + 4 ..]);
                char_count += dont_section[i + 4 ..].len;
                break;
            }
        }
    }
    var index: usize = 0;
    var do_chars = try alloc.alloc(u8, char_count);
    for (do_chunks.items) |value| {
        for (value) |char| {
            if (index >= do_chars.len) return do_chars;
            do_chars[index] = char;
            index += 1;
        }
    }
    return do_chars;
}

fn getValidSections(chars: []const u8, alloc: Allocator) ![][]const u8 {
    var sections = ArrayList([]const u8).init(alloc);
    var mul_sections = std.mem.split(u8, chars, "mul(");
    while (mul_sections.next()) |section| {
        var num_section = std.mem.split(u8, section, ")");
        while (num_section.next()) |nums| {
            if (!validateSection(nums)) {
                continue;
            }
            try sections.append(nums);
        }
    }

    return sections.toOwnedSlice();
}

fn getOps(sections: [][]const u8, alloc: Allocator) ![][2]i32 {
    var operations = ArrayList([2]i32).init(alloc);
    for (sections) |section| {
        var current_op: [2]i32 = undefined;
        var current_num: [3]u8 = undefined;
        var num_size: usize = 0;
        for (section, 0..) |char, i| {
            if (day1.isNum(char)) {
                current_num[num_size] = char;
                num_size += 1;
                if (i == section.len - 1) {
                    current_op[1] = try std.fmt.parseInt(i32, current_num[0..num_size], 10);
                    try operations.append(current_op);
                }
            } else if (char == ',') {
                if (num_size == 0) {
                    break;
                }
                current_op[0] = std.fmt.parseInt(i32, current_num[0..num_size], 10) catch |err| {
                    std.debug.panic("{s}", .{current_num[0..num_size]});
                    return err;
                };
                num_size = 0;
            }
        }
    }
    return operations.toOwnedSlice();
}

fn validateSection(line: []const u8) bool {
    for (line, 0..) |char, i| {
        if (day1.isNum(char)) continue;
        if (char == ',' and i != 0) {
            continue;
        }
        return false;
    }
    return true;
}

//test "safety" {
//    const alloc = std.testing.allocator;
//    try run(alloc);
//}
