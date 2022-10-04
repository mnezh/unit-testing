namespace Divide.Service;

public class Divisor
{
    public static double divide(double a, double b) {
        double result = a/b;
        if (Double.IsInfinity(result)) {
            throw new DivideByZeroException($"{a}/{b}");
        }
        return result;
    }
}
