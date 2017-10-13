public class InvokeSpecialSuperTest {
  public static void main (String[] args) {
    GrandFather obj = new Son();
    obj.foo(); // invokevirtual
  }
}
  
  class GrandFather {
    public void foo(){
      System.out.println ("grandfather");
    }
  }
  
  class Father extends GrandFather {
  }
  
  class Son extends Father {
    public void foo(){
      super.foo(); // invokespecial?
    }
  }
  
