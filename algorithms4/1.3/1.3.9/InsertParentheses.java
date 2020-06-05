import edu.princeton.cs.algs4.Stack;
import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;

public class InsertParentheses {

  public static void main(String[] args) {
    Stack<String> ops = new Stack<String>();
    Stack<String> opers = new Stack<String>();

    while (!StdIn.isEmpty()) {
      String s = StdIn.readString();
      if (s.equals("+") || s.equals("-") || s.equals("*") || s.equals("/")) {
        opers.push(s);
      }
      else if (s.equals(")")) {
        String op2 = ops.pop(), op1 = ops.pop();
        String oper = opers.pop();
        ops.push("( " + op1 + " " + oper + " " + op2 + " )");
      } else {
        ops.push(s);
      }
    }

    String oper = "";
    String op1 = "", op2 = "";

    if (!ops.isEmpty()) op2 = ops.pop();
    if (!ops.isEmpty()) op1 = ops.pop();
    if (!opers.isEmpty()) oper = opers.pop();

    if (!op1.equals("")) StdOut.print(op1 + " ");
    if (!oper.equals("")) StdOut.print(oper + " ");
    if (!op2.equals("")) StdOut.print(op2);
    StdOut.println();
  }
}
