import java.lang.Object;
public class ClassFileTest extends Object{
  public static final boolean FLAG = true;
  public static final byte BYTE = 123;
  public static final char X = 'X';
  public static final short SHORT = 12345;
  public static final int INT = 123456789;
  public static final long LONG = 12345678901L;
  public static final float PI = 3.14f;
  public static final double E = 2.71828;
  public static final String Str = "This is a string";
  public static final Number num = 10086;
  public Object obj;
  public static void main(String[] args) throws RuntimeException {
    int res = 0;
    for (int i = 0; i < 10 ; i++) {
      res += i;
    }
  }
}
