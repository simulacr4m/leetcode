import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;
import edu.princeton.cs.algs4.Stack;

public class EvaluatePostfix {
  public static void main(String[] args) {
    Stack<Double> s = new Stack<Double>();
    while (!StdIn.isEmpty()) {
      String a = StdIn.readString();
      if (a.equals("+") || a.equals("-") || a.equals("*") || a.equals("/")) {
        double op2 = s.pop(), op1 = s.pop();
        if (a.equals("+")) s.push(op1 + op2);
        if (a.equals("-")) s.push(op1 - op2);
        if (a.equals("*")) s.push(op1 * op2);
        if (a.equals("/")) s.push(op1 / op2);
      }
      else s.push(Double.parseDouble(a));
    }
    StdOut.println(s.pop());
  }
}
