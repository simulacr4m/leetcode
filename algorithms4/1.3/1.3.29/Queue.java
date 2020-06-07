import edu.princeton.cs.algs4.StdOut;
import edu.princeton.cs.algs4.StdIn;

public class Queue<Item> {

  private int n;
  private Node last;

  private class Node {
    Item item;
    Node next;
  }

  public int size() { return n; }
  public boolean isEmpty() { return n == 0; }

  public void enqueue(Item item) {
    Node oldlast = last;
    last = new Node();
    last.item = item;
    if (oldlast == null) last.next = last;
    else {
      Node first = oldlast.next;
      oldlast.next = last;
      last.next = first;
    }
    ++n;
  }

  public Item dequeue() {
    Item item = last.next.item;
    if (last.next == last) last = null;
    else last.next = last.next.next;
    --n;
    return item;
  }

  public static void main(String[] args) {
    Queue<String> q = new Queue<String>();

    while (!StdIn.isEmpty()) {
      String item = StdIn.readString();
      if (!item.equals("-")) q.enqueue(item);
      else if (!q.isEmpty()) StdOut.print(q.dequeue() + " ");
    }
    StdOut.println("(" + q.size() + " left on queue)");
  }
}
