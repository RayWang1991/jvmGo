public class TestPrivate extends A {
  public static void main(String[] args) {
    System.out.printf("Before %d\n", A.sn);
    TestPrivate.sn = 1024;
    System.out.printf("After %d\n", A.sn);
  }
}

class A {
  public static int sn = 2048;
  protected int n;
  public void setN(int a) {
    n = a;
  }
  public int N() {
    return n;
  }
}

