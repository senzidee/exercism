class PascalsTriangle {
  List<List<int>> rows(int num) {
    List<List<int>> triangle = [];
    for (var i = 0; i < num; i++) {
      List<int> row = [];
      int c = 1;
      for (var j = 0; j <= i; j++) {
        row.add(c);
        c = c * (i - j) ~/ (j + 1);
      }
      triangle.add(row);
    }
    return triangle;
  }
}
