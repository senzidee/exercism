extension ListOps on List {
  void append<T>(List<T> other) {
    for (var elem in other) {
      this.add(elem);
    }
  }

  List<T> concat<T>() {
    List<T> _concat = [];
    for (var elem in this) {
      if (elem is List) {
        for (var elem2 in elem) {
          _concat.add(elem2);
        }
      }
    }
    return _concat;
  }

  List<T> filter<T>(bool Function(T elem) predicate) {
    List<T> _filtered = [];
    for(var elem in this) {
      if(predicate(elem)) {
        _filtered.add(elem);
      }
    }
    return _filtered;
  }

  int count() {
    int _count = 0;
    for(var elem in this) {
      _count++;
    }
    return _count;
  }

  List<T> myMap<T>(T Function(T elem) fn) {
    List<T> _mapped = [];
    for(var elem in this) {
      if(fn(elem) != null) {
        _mapped.add(fn(elem));
      }
    }
    return _mapped;
  }

  U foldl<T, U>(U initial, U Function(U acc, T elem) fn) {
    var _acc = initial;
    for (var i = 0; i < this.length; i++) {
      _acc = fn(_acc, this[i]);
    }
    return _acc;
  }

  U foldr<T, U>(U initial, U Function(T elem, U acc) fn) {
    var _acc = initial;
    for (var i = this.length; i > 0; i--) {
      _acc = fn( this[i - 1], _acc);
    }
    return _acc;
  }

  List<T> reverse<T>() {
    List<T> _reversed = [];
    for (var i = this.length; i > 0; i--) {
      _reversed.add(this[i - 1]);
    }
    return _reversed;
  }
}
