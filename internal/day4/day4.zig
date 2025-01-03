const std = @import("std");
const io = std.io;
const ArrayList = std.ArrayList;
const Allocator = std.mem.Allocator;
const FileManager = @import("file_manager.zig");

pub fn run(alloc: Allocator) !void {
    const data = try FileManager.readFullFile("realInput.txt", alloc);
    defer alloc.free(data);
    const lines = try FileManager.splitByLine(data, alloc);
    defer alloc.free(lines);

    var matches: u32 = 0;
    var crosses: u32 = 0;
    for (lines, 0..) |line, y| {
        for (0..line.len) |x| {
            matches += checkDirections(lines, x, y, "XMAS");
            if (checkForX(lines, x, y, "MAS")) crosses += 1;
        }
    }
    std.debug.print("XMAS count: {}\n", .{matches});
    std.debug.print("X-MAS count: {}\n", .{crosses});
}

fn checkForX(lines: [][]const u8, x: usize, y: usize, token: []const u8) bool {
    var numFound: u8 = 0;
    var buff: [6]u8 = undefined;
    const south_east = getOffsetFromCenter(lines, buff[0..3], x, y, 1, 1);
    const south_west = getOffsetFromCenter(lines, buff[3..6], x, y, -1, 1);
    if (std.mem.eql(u8, south_east, token) or isReversed(south_east, token))
        numFound += 1;
    if (std.mem.eql(u8, south_west, token) or isReversed(south_west, token))
        numFound += 1;
    return numFound == 2;
}

fn checkDirections(lines: [][]const u8, x: usize, y: usize, token: []const u8) u8 {
    var numFound: u8 = 0;
    var buff: [16]u8 = undefined;
    var directions: [4][]const u8 = undefined;
    directions[0] = getOffset(lines, buff[0..4], x, y, 1, 1); //South East
    directions[1] = getOffset(lines, buff[4..8], x, y, 1, -1); //North West
    directions[2] = getOffset(lines, buff[8..12], x, y, 1, 0); //East
    directions[3] = getOffset(lines, buff[12..16], x, y, 0, 1); //South

    for (directions) |dir| {
        if (std.mem.eql(u8, dir, token) or isReversed(dir, token))
            numFound += 1;
    }
    return numFound;
}

fn isReversed(str: []const u8, token: []const u8) bool {
    var i: usize = token.len - 1;
    var j: usize = 0;
    if (str.len != token.len) return false;
    while (true) {
        if (str[j] != token[i]) return false;
        if (i == 0) break;
        i -= 1;
        j += 1;
    }
    return true;
}

fn getOffsetFromCenter(lines: [][]const u8, buffer: []u8, start_x: usize, start_y: usize, x_offset: i8, y_offset: i8) []const u8 {
    const x: i32 = @intCast(start_x);
    const y: i32 = @intCast(start_y);
    if (y + y_offset < 0 or y + y_offset >= lines.len) return "";
    if (x + x_offset < 0 or x + x_offset >= lines.len) return "";
    if (y - y_offset < 0 or y - y_offset >= lines.len) return "";
    if (x - x_offset < 0 or x - x_offset >= lines.len) return "";

    buffer[0] = lines[@intCast(y + y_offset)][@intCast(x + x_offset)];
    buffer[1] = lines[@intCast(y)][@intCast(x)];
    buffer[2] = lines[@intCast(y - y_offset)][@intCast(x - x_offset)];

    return buffer[0..3];
}

fn getOffset(lines: [][]const u8, buffer: []u8, start_x: usize, start_y: usize, x_offset: i8, y_offset: i8) []const u8 {
    var len: usize = 0;
    var x: i32 = @intCast(start_x);
    const y_len: i32 = @intCast(lines.len);
    var x_len: i32 = 0;
    var y: i32 = @intCast(start_y);

    for (0..buffer.len) |i| {
        buffer[i] = lines[@intCast(y)][@intCast(x)];
        len += 1;

        y += if ((y > 0 and y_offset < 0) or (y >= 0 and y_offset >= 0)) y_offset else break;
        if (y == y_len) break;

        x_len = @intCast(lines[@intCast(y)].len);
        x += if ((x > 0 and x_offset < 0) or (x >= 0 and x_offset >= 0)) x_offset else break;
        if (x == x_len) break;
    }

    return buffer[0..len];
}

test "safety" {
    const alloc = std.testing.allocator;
    try run(alloc);
}
