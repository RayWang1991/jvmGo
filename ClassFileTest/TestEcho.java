public class TestEcho {
  public static void main (String[] args) {
 //   System.out.printf("%d",1);
    int[] array = new int[args.length];
    for(int i = 0; i < args.length; i++ ) {
      System.out.println(args[i]);
      array[i] = args[i].length();
    }
    int res = 0;
    for(int i : array) {
      res += i;
    }
    System.out.print("Num of chars: ");
    System.out.println(res);
  }
}
