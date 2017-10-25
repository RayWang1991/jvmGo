public abstract class TestMethod {
  public abstract void foo(int a);
}

class B extends TestMethod {
  public void test(){
    this.foo(1);
  }
}
