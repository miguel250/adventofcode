from typing import IO, Tuple
import argparse

def parse(input_file: IO[str]) -> Tuple[int, int]:
    lines = input_file.readlines()
    total_volume: int = 0
    total_ribbon_feet: int = 0
    shortest_side: int = 0
    min_value: int = 0

    for line in lines:
        length, wide, height = [int(i) for i in line.split("x")]
        min_value = wide *  length
        side_two = wide * height
        side_three = height * length

        shortest_side = wide + length

        if min_value > side_two:
            min_value = side_two
            shortest_side = wide + height

        if min_value > side_three:
            min_value = side_three
            shortest_side = height + length


        total_ribbon_feet += (length * height * wide) + (2 * shortest_side)
        total_volume += volume(length, wide, height) + min_value
    return total_volume, total_ribbon_feet

def volume(length: int, wide: int, height: int) -> int:
    return 2 * length * wide + 2 * wide * height + 2 * height * length

def main()-> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("-input-path", default="input", type=str, help="path to input file")
    args = parser.parse_args()

    input_file = open(args.input_path)
    output, ribbon = parse(input_file)

    print("result: %s" % output)
    print("ribbon: %s" % ribbon)
    input_file.close()

if __name__ == '__main__':
    main()
