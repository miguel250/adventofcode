import io
import pytest

import floors_part2

@pytest.mark.parametrize(
    "test_input,expected", [
        ("(())", (0, 0)),
        ("(()(()(", (3, 0)),
        ("(((", (3, 0)),
        ("(()(()(", (3, 0)),
        ("))(((((", (3, 1)),
        ("())", (-1, 3)),
        ("))(", (-1, 1)),
        (")))", (-3, 1)),
        (")())())", (-3, 1)),
    ],
)
def test_parse_input(test_input: str, expected: int) -> None:
    input_file = io.StringIO(test_input)
    output = floors_part2.parse_input(input_file, -1)
    assert output == expected
