public class TestMethod extends B{
  public static void main(String[] args) {
    B b = new B();
    int a = 10;
    a = b.foo(a);
  }
}

class B  {
  public int foo(int a) {
    return a * 2;
  }
}
