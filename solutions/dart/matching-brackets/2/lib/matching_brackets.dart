class MatchingBrackets {
  // Match brackets in a string
  bool isPaired(String s) {
    var stack = [];
    var pairs = {
      '}': '{',
      ']': '[',
      ')': '(',
    };
    for (var char in s.split('')) {
      if (pairs.containsValue(char)) {
        stack.add(char);
      } else if (pairs.containsKey(char)) {
        if (stack.isEmpty || stack.last != pairs[char]) {
          return false;
        }
        stack.removeLast();
      }
    }
    return stack.isEmpty;
  }
}
