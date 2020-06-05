import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;
import edu.princeton.cs.algs4.Stack;

public class Parentheses {

  public static boolean isValid(String text) {
    Stack<Character> s = new Stack<Character>();
    for (int i = 0; i < text.length(); i++) {
      char c = text.charAt(i);
      if (c == ')' || c == '}' || c == ']') {
        if (s.isEmpty()) return false;
        char value = s.pop();
        if (value == '(' && c != ')') return false;
        if (value == '{' && c != '}') return false;
        if (value == '[' && c != ']') return false;
      } else {
        s.push(c);
      }
    }
    return s.isEmpty();
  }

  public static void main(String[] args) {
    String parentheses = StdIn.readLine();
    StdOut.println(isValid(parentheses));
  }
}
