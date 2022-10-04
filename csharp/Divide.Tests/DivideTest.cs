using Microsoft.VisualStudio.TestTools.UnitTesting;
using Divide.Service;

namespace Divide.Tests;

[TestClass]
public class DivideTest
{
    [TestMethod]
    [DataRow(10,2,5)]
    [DataRow(10,4,2.5)]
    public void TestDivision(double dividend, double divisor, double result)
    {
       Assert.AreEqual(Divisor.divide(dividend,divisor), result);
    }
    [ExpectedException(typeof(DivideByZeroException))]
    public void TestException() {
        Divisor.divide(10, 0);
    }
}