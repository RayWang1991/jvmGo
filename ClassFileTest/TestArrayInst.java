public class TestArrayInst {
  public static void main (String[] args) {
    boolean[] bArr = {true, false, true}; // new array b
    bArr[1] = bArr[2]; //bastore; baload
    long[] lArr = {1024, 1000000000000L, 0} ;// new array J
    lArr[1] = lArr[2] ;//lastore; laload
    double[] dArr = {102.4, 100.0000000000,  0.1} ;// new array D
    dArr[1] = dArr[2] ;//dastore; daload
    
//    Object[][][] oArr = new Object[3][4][5] ;// multianewarray 
//    oArr[1][2][3] = new Object();
  }
}
