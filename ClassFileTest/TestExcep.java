public class TestExcep {
  public void test() throws Throwable{
    try {
      int i = 0;
    } catch (Throwable e){
      throw e;
    }
  }
}
