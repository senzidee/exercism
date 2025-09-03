class Hamming {
  int distance(String src, String dest) {
    if (src.length.compareTo(dest.length) != 0) {
      throw ArgumentError('strands must be of equal length');
    }
    if (src.compareTo(dest) == 0) {
      return 0;
    }
    var count = 0;
    for (var i = 0; i < src.length; i++) {
      if (src.codeUnitAt(i) != dest.codeUnitAt(i)) {
        count++;
      }
    }
    return count;
  }
}
