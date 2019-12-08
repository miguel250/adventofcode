from typing import IO, Tuple
import argparse

INCREASE_FLOOR = "("
DECREASE_FLOOR = ")"

def parse_input(input_file: IO[str], find_floor: int = -1) -> Tuple[int, int]:
    line = input_file.readline()
    floor: int = 0
    find_floor_position: int = -1

    for i, char in enumerate(line):
        if char is INCREASE_FLOOR:
            floor += 1
        elif char is DECREASE_FLOOR:
            floor -= 1

        if floor is find_floor and find_floor_position is -1:
            find_floor_position = i

    return floor, find_floor_position + 1

def main()-> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("-input-path", default="input", type=str, help="path to input file")
    args = parser.parse_args()

    input_file = open(args.input_path)
    output, floor_position = parse_input(input_file)

    print("result: %s" % output)
    print("result: %s" % floor_position)
    input_file.close()

if __name__ == '__main__':
    main()
