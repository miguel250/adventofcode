from typing import IO
import argparse

def parse(input_file: IO[str]) -> int:
    lines = input_file.readlines()
    total_volume: int = 0
    min_value: int = 0

    for line in lines:
        length, wide, height = [int(i) for i in line.split("x")]
        min_value = wide *  length
        side_two = wide * height
        side_three = height * length

        if min_value > side_two:
            min_value = (wide * height)

        if min_value > side_three:
            min_value = side_three

        total_volume += volume(length, wide, height) + min_value
    return total_volume

def volume(length: int, wide: int, height: int) -> int:
    return 2 * length * wide + 2 * wide * height + 2 * height * length

def main()-> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("-input-path", default="input", type=str, help="path to input file")
    args = parser.parse_args()

    input_file = open(args.input_path)
    output = parse(input_file)

    print("result: %s" % output)
    input_file.close()

if __name__ == '__main__':
    main()
