import io
import pytest

import calculate_part2


@pytest.mark.parametrize(
    "test_input,expected", [
        ("2x3x4\n", (58, 34)),
        ("1x1x10\n", (43, 14)),
    ]
)
def test_parser(test_input, expected):
    input_file = io.StringIO(test_input)
    result = calculate_part2.parse(input_file)

    assert result == expected

def test_calculate_volume():
    result = calculate_part2.volume(2, 3, 4)
    assert result == 52
