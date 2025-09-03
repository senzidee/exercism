class MatchingBrackets {
  // Match brackets in a string
  bool isPaired(String s) {
    var _buffer = [];
    for (var _s in s.split('')) {
      if (_s == '{' || _s == '[' || _s == '(') {
        _buffer.add(_s);
      } else if (_s == '}' || _s == ']' || _s == ')') {
        if (_buffer.isEmpty) {
          return false;
        }
        if (_s == '}' && _buffer.last == '{') {
          _buffer.removeLast();
          continue;
        }
        if (_s == ']' && _buffer.last == '[') {
          _buffer.removeLast();
          continue;
        }
        if (_s == ')' && _buffer.last == '(') {
          _buffer.removeLast();
          continue;
        }
        return false;
      }
    }
    return _buffer.isEmpty;
  }
}
