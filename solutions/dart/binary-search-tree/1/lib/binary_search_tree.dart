class Node {
  Node? left;
  Node? right;
  String data;

  Node(this.data);
}

class BinarySearchTree {
  late Node root;
  List<String> get sortedData => _sorted(this.root);

  BinarySearchTree(String val) {
    this.root = new Node(val);
  }

  void insert(String val) {
    _insert(this.root, val);
  }

  void _insert(Node? node, String val) {
    if (int.parse(val) > int.parse(node!.data)) {
      if (node.right == null) {
        node.right = new Node(val);
      } else {
        _insert(node.right, val);
      }
    } else {
      if (node.left == null) {
        node.left = new Node(val);
      } else {
        _insert(node.left, val);
      }
    }
  }

  List<String> _sorted(Node? node) {
    List<String> data = [];
    if (node != null) {
       data.addAll(_sorted(node.left));
       data.add(node.data);
       data.addAll(_sorted(node.right));
    }
    return data;
  }
}
