class SumOfMultiples {
  int sum(List<int> numbers, int limit) {
    List<int> _numbers = [];
    numbers.forEach((element) { _numbers += getMultiple(element, limit); });
    _numbers.sort();
    for (var i = 1; i < _numbers.length; i++) {
      if (_numbers[i] == _numbers[i -1]) {
        _numbers.removeAt(i);
      }
    }

    return _numbers.fold(0, (previousValue, element) => previousValue + element);
  }

  List<int> getMultiple(int number, int limit) {
    List<int> multiple = [0];
    if (number == 0) {
      return multiple;
    }
    var i = 2;
    var _number = number;
    while (_number < limit) {
      multiple.add(_number);
      _number = number * i++;
    }
    return multiple;
  }
}
