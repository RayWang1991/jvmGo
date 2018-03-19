public class TestInvokeVirtual {
  class GrandFa {
    protected int id = 1;
    public void say(){
      System.out.println("Grand Father id:");
      System.out.println(id);
    }
  }

  class Father extends GrandFa{
    protected int id = 2;
  }
  
  class Son extends Father {
    public void say (){
      System.out.println("Son id");
      System.out.println(id + 100);
    }
  }

  public void main(String[] args) {
    System.out.println("test1");
    GrandFa a = new Father();
    a.say(); // should be "Grand Father id:\n 1"

    System.out.println("test2");
    GrandFa b = new Son();
    b.say(); // should be "Son id:\n102";
  }
}
