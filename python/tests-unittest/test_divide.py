import unittest
from src.divide import divide


class TestDivide(unittest.TestCase):
    def test_divide_integer(self):
        self.assertEqual(divide(10, 2), 5)

    def test_divide_float(self):
        self.assertEqual(divide(10, 4), 2.5)

    def test_expect_exception(self):
        self.assertRaises(ZeroDivisionError, lambda: divide(10, 0))
