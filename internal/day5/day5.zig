const std = @import("std");
const ArrayList = std.ArrayList;
const Allocator = std.mem.Allocator;
const FileManager = @import("file_manager.zig");

const node = struct {
    value: u32,
    after: ArrayList(u32),
    //    previous: ArrayList(u8)
    pub fn init(value: u32, after_value: ?u32, alloc: Allocator) !*node {
        var newNode = try alloc.create(node);
        var afterList = ArrayList(u32).init(alloc);
        if (after_value != null) try afterList.append(after_value.?);
        newNode.after = afterList;
        newNode.value = value;
        return newNode;
    }

    pub fn deinit(self: *node) void {
        self.after.deinit();
    }
};

pub fn run(alloc: Allocator) !void {
    const data = try FileManager.readFullFile("realInput.txt", alloc);
    defer alloc.free(data);
    const lines = try FileManager.splitByLine(data, alloc);
    defer alloc.free(lines);

    const rules = getRules(lines);
    const nodes = try getNodesFromRuleset(rules, alloc);
    defer nodes.deinit();

    const updates = lines[rules.len + 1 ..];
    var total: u32 = 0;
    for (updates) |update| {
        total += try getMiddleOfUpdate(true, update, nodes, alloc);
    }
    std.debug.print("Output - {}\n", .{total});
}

fn sortUpdate(keys: ArrayList(u32), nodes: ArrayList(*node), alloc: Allocator) !ArrayList(u32) {
    var key_nodes = ArrayList(*node).init(alloc);
    var sorted_keys = ArrayList(u32).init(alloc);
    defer keys.deinit();

    for (keys.items) |key| {
        for (nodes.items) |current| {
            if (current.value == key) {
                try key_nodes.append(current);
                break;
            }
        }
    }

    var start: usize = 0;
    var sorted = false;
    while (!sorted) {
        sorted = true;
        var max_score_index: usize = 0;
        var max_score: u32 = 0;

        for (key_nodes.items[start..], 0..) |current, i| {
            var score: u32 = 0;
            for (key_nodes.items) |next| {
                if (std.mem.containsAtLeast(u32, current.after.items, 1, &.{next.value})) score += 1;
            }

            if (i == 0) {
                max_score = score;
                continue;
            }

            if (score > max_score) {
                max_score = score;
                max_score_index = i;
                sorted = false;
            }
        }
        const temp = key_nodes.items[start];
        key_nodes.items[start] = key_nodes.items[start..][max_score_index];
        key_nodes.items[start..][max_score_index] = temp;
        start += 1;
        if (start == key_nodes.items.len) {
            start = 0;
        } else {
            sorted = false;
        }
    }

    for (key_nodes.items) |n| {
        try sorted_keys.append(n.value);
    }

    return sorted_keys;
}

fn getMiddleOfUpdate(fix_values: bool, update: []const u8, nodes: ArrayList(*node), alloc: Allocator) !u32 {
    var value_iterator = std.mem.splitSequence(u8, update, ",");
    var keys = ArrayList(u32).init(alloc);

    while (value_iterator.next()) |value| {
        const key = try std.fmt.parseUnsigned(u32, value, 10);
        try keys.append(key);
    }

    var valid = true;
    for (keys.items, 0..) |key, i| {
        var node_match: ?*node = null;
        for (nodes.items) |current| {
            if (current.value == key) {
                node_match = current;
                break;
            }
        }
        if (node_match == null) {
            return 0;
        }

        for (keys.items[i + 1 ..]) |after_key| {
            var found = false;
            for (node_match.?.after.items) |after_value| {
                if (after_value == after_key) {
                    found = true;
                    break;
                }
            }
            if (!found) {
                valid = false;
            }
        }
    }

    if (fix_values) {
        if (valid) {
            return 0;
        }
        keys = try sortUpdate(keys, nodes, alloc);
    }
    return keys.items[keys.items.len / 2];
}

pub fn getRules(lines: [][]const u8) [][]const u8 {
    for (lines, 0..) |line, i| {
        if (line.len == 0) {
            return lines[0..i];
        }
    }
    return lines;
}

pub fn getNodesFromRuleset(rules: [][]const u8, alloc: Allocator) !ArrayList(*node) {
    var nodes = ArrayList(*node).init(alloc);
    for (rules) |rule| {
        var split_iterator = std.mem.split(u8, rule, "|");
        const firstNum = try std.fmt.parseUnsigned(u32, split_iterator.next().?, 10);
        const secondNum = try std.fmt.parseUnsigned(u32, split_iterator.next().?, 10);

        var first_found = false;
        var second_found = false;
        for (nodes.items) |current| {
            if (current.value == firstNum) {
                try current.after.append(secondNum);
                first_found = true;
                break;
            } else if (current.value == secondNum) {
                second_found = true;
            }
        }

        if (!first_found) {
            const new_node = try node.init(firstNum, secondNum, alloc);
            try nodes.append(new_node);
        }
        if (!second_found) {
            const new_node = try node.init(secondNum, null, alloc);
            try nodes.append(new_node);
        }
    }
    return nodes;
}
