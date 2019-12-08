from typing import IO
import argparse

INCREASE_FLOOR = "("
DECREASE_FLOOR = ")"

def parse_input(input_file: IO[str]) -> int:
    line = input_file.readline()
    floor: int = 0

    for char in line:
        if char is INCREASE_FLOOR:
            floor += 1
        elif char is DECREASE_FLOOR:
            floor -= 1
    return floor

def main()-> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("-input-path", default="input", type=str, help="path to input file")
    args = parser.parse_args()

    input_file = open(args.input_path)
    output = parse_input(input_file)

    print("result: %s" % output)
    input_file.close()

if __name__ == '__main__':
    main()
