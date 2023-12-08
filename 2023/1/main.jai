
#import "Basic";
#import "File";
#import "String";

main :: () {
    f, ok := read_entire_file("input.txt");
    assert(ok, "Couldn't read input.txt file");
    if ok {
        part_one(f);
        part_two(f);
    }
}

digit_strings: []string = string.["zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"];
part_two :: (f: string) {
    stream := f;
    sum: int;
    ascii: [2]u8;
    first_digit: bool;
    while stream.count > 0 {
        for digit_strings if begins_with(stream, it) {
            digit: u8 = cast(u8)it_index + #char "0";
            add_digit(*ascii, digit, first_digit);
            first_digit = true;
            break;
        }
        if stream[0] >= #char "0" && stream[0] <= #char "9" { // Char is a digit
            add_digit(*ascii, stream[0], first_digit);
            first_digit = true; // just do it everytime
        }
        if stream[0] == #char "\n" { // End of Line
            val: int = 10*(ascii[0] - #char"0");
            val += ascii[1] - #char"0";
            sum += val;
            first_digit = false;
        }
        stream.data += 1; stream.count -= 1; // advance
    }
    print("part two sum: %\n", sum);
}

part_one :: (f: string) {
    stream := f;
    sum: int;
    ascii: [2]u8;
    first_digit: bool;
    while stream.count > 0 {
        if stream[0] >= #char "0" && stream[0] <= #char "9" { // Char is a digit
            add_digit(*ascii, stream[0], first_digit);
            first_digit = true;
        }
        if stream[0] == #char "\n" { // End of Line
            val: int = 10*(ascii[0] - #char"0");
            val += ascii[1] - #char"0";
            sum += val;
            first_digit = false;
        }
        stream.data += 1; stream.count -= 1; // advance
    }
    print("part one sum: %\n", sum);
}

add_digit :: (arr: *[2]u8, digit: u8, first: bool) {
    if !first
        arr.*[0] = digit;
    arr.*[1] = digit;
}
