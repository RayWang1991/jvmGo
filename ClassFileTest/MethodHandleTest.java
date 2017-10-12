import java.lang.invoke.*;

public class MethodHandleTest {
  private static void hello() {
    System.out.println("Hello world!");
  }
  
  public static void main(String[] args) throws Throwable {
    MethodType type = MethodType.methodType(void.class);
    MethodHandle method = MethodHandles.lookup().
      findStatic(MethodHandleTest.class, "hello", type);
    method.invoke();
    }
}
