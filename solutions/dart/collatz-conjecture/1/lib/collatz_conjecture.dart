class CollatzConjecture {
  int steps(int val) {
    if (val < 1) {
      throw ArgumentError('Only positive integers are allowed');
    }
    return _resolv(val, 0);
  }
  int _resolv(int val, int steps) {
    if (val == 1) {
      return steps;
    }
    if (val.isEven) {
      return _resolv(val ~/2, ++steps);
    }
    return _resolv((val * 3) + 1, ++steps);
  }
}
