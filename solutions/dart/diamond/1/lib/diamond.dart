class Diamond {
  List<String> rows(String s) {
    var _rows = <String>[];
    for (var i = 65; i <= s.codeUnitAt(0); i++) {
      var _pre = [for (var j = 0; j < s.codeUnitAt(0) - i; j++) ' '];
      var _post = [for (var j = i; j > 65; j--) ' '];
      var _row = _pre.join() + String.fromCharCode(i) + _post.join();
      if (_post.isNotEmpty) {
        _row = _row.substring(0, _row.length - 1) + _post.join() + String.fromCharCode(i);
      }
      _rows.add(_row + _pre.join());
    }
    for (var i = _rows.length - 1; i > 0; i--) {
      _rows.add(_rows[i - 1]);
    }
    return _rows;
  }
}
