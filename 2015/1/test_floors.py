import io
import pytest

import floors

@pytest.mark.parametrize(
    "test_input,expected", [
        ("(())", 0),
        ("(()(()(", 3),
        ("(((", 3),
        ("(()(()(", 3),
        ("))(((((", 3),
        ("())", -1),
        ("))(", -1),
        (")))", -3),
        (")())())", -3),
    ],
)
def test_parse_input(test_input: str, expected: int) -> None:
    input_file = io.StringIO(test_input)
    output = floors.parse_input(input_file)
    assert output == expected
