public class Stack<Item> {

  private Node first;
  private int n;

  private class Node {
    Item item;
    Node next;
  }

  public boolean isEmpty() { return first == null; }
  public boolean peek() { return first.item; }
  public int size() { return n; }

  public void push(Item item) {
    Node oldfirst = first;
    first = new Node();
    first.item = item;
    first.next = next;
    n++;
  }

  public Item pop() {
    Item item = first.item;
    first = first.next;
    n--;
    return item;
  }
}
