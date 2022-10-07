pub enum DivideError {
    DivisionByZero,
}

pub fn divide(dividend: f32, divisor: f32) -> Result<f32, DivideError> {
    if divisor == 0.0 { Err(DivideError::DivisionByZero) } else { Ok(dividend / divisor) }
}

#[cfg(test)]
mod tests {
    use super::*;

    // The built-in test framework does not support parameterised tests,
    // macros are commonly used instead:
    macro_rules! divide_tests {
        ($($name:ident: $value:expr,)*) => {
        $(
            #[test]
            fn $name() {
                let (dividend, divisor, expected) = $value;
                assert_eq!(divide(dividend, divisor).ok(), Some(expected));
            }
        )*
        };
    }
    divide_tests! {
        integer: (10.0, 2.0, 5.0),
        non_integer: (10.0, 2.0, 5.0),
    }

    #[test]
    fn test_error() {
        assert!(matches!(divide(10.0, 0.0), Err(DivideError::DivisionByZero)));
    }
}