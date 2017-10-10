//Test sum for loop if control
public class SumTest {
  public static void main(String[] args) {
    int res = 0;
    for( int i = 0; i < 10; i++) {
      if (i%2 ==0) {
        res += i;
      }
    }
  }
}
