import io
import pytest

import calculate


@pytest.mark.parametrize(
    "test_input,expected", [
        ("2x3x4\n", 58),
        ("1x1x10\n", 43),
    ]
)
def test_parser(test_input, expected):
    input_file = io.StringIO(test_input)
    result = calculate.parse(input_file)

    assert result == expected

def test_calculate_volume():
    result = calculate.volume(2, 3, 4)
    assert result == 52
