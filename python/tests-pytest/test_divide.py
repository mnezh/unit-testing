import pytest
from src.divide import divide


@pytest.mark.parametrize("a, b, result", [[10, 2, 5], [10, 4, 2.5]])
def test_division(a, b, result):
    assert divide(a, b) == result


def test_expect_exception():
    with pytest.raises(ZeroDivisionError):
        divide(10, 0)
