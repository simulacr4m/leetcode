import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;
import edu.princeton.cs.algs4.Stack;

public class InfixToPostfix {
  public static void main(String[] args) {
    Stack<String> values = new Stack<String>();
    Stack<String> ops = new Stack<String>();
    while (!StdIn.isEmpty()) {
      String s = StdIn.readString();
      if (s.equals("(")) continue;
      if (s.equals("+") || s.equals("-") || s.equals("*") || s.equals("/")) {
        ops.push(s);
      }
      else if (s.equals(")")) {
        String v = values.pop();
        String op = ops.pop();
        values.push(values.pop() + " " + v + " " + op);
      } else {
        values.push(s);
      }
    }
    StdOut.println(values.pop());
  }
}
