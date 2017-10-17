public class ObjTest {
  public static int staticV = 5;
  public int instanceV;
  public static void main(String[]args){
    ObjTest o = new ObjTest();
    o.instanceV = 100;
    ObjTest.staticV += 10086;
    System.out.println(o.instanceV);
    System.out.println(ObjTest.staticV);
  }
}
