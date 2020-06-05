import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;
import edu.princeton.cs.algs4.Stack;

public class StackClient {

  public static Stack<String> copy(Stack<String> s) {
    Stack<String> tmp = new Stack<String>();
    Stack<String> copy = new Stack<String>();
    for (String str : s)
      tmp.push(str);
    while (!tmp.isEmpty())
      copy.push(tmp.pop());
    return copy;
  }

  public static void main(String[] args) {
    Stack<String> s = new Stack<String>();
    while (!StdIn.isEmpty()) {
      String str = StdIn.readString();
      s.push(str);
    }
    StdOut.println(s + "\n");
    StdOut.println(copy(s));
  }
}
