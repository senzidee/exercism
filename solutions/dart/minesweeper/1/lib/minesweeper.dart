extension on List<String> {
  int neighbors(int x, int y) {
    var _c = [];
    if (x > 0) {
      _c.add(this[x - 1][y]);
    }
    if (y > 0) {
      _c.add(this[x][y - 1]);
    }
    if (x < this.length - 1) {
      _c.add(this[x + 1][y]);
    }
    if (y < this[0].length - 1) {
      _c.add(this[x][y + 1]);
    }
    if (x > 0 && y > 0) {
      _c.add(this[x - 1][y - 1]);
    }
    if (x < this.length - 1 && y < this[0].length - 1) {
      _c.add(this[x + 1][y + 1]);
    }
    if (x > 0 && y < this[0].length - 1) {
      _c.add(this[x - 1][y + 1]);
    }
    if (x < this.length - 1 && y > 0) {
      _c.add(this[x + 1][y - 1]);
    }
    return _c
        .where((value) => value == '*')
        .length;
  }
}
class Minesweeper {
  List<String> annotated = [];
  late List<List<String>> board;

  Minesweeper(List <String> board) {
    this.annotated = board;
    this.board = board.map((element) => element.split('').toList()).toList();

    List<List<String>> _annotated = this.board;

    for (int i = 0; i < this.board.length; i++) {
      for (int j = 0; j < this.board[i].length; j++) {
        if (board[i][j] == '*') {
          continue;
        }
        _annotated[i][j] = board.neighbors(i, j).toString();
        _annotated[i][j] = _annotated[i][j] == '0' ? ' ' : _annotated[i][j];
      }
      this.annotated[i] = _annotated[i].join();
    }

  }
}
