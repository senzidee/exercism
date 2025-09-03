class Luhn {
  bool valid(String number) {
    number = number.replaceAll(' ', '');
    if (number.length < 2 || number.contains(new RegExp(r'[^0-9]'))) {
      return false;
    }
    var _number = number.split('').map((e) => int.parse(e)).toList();
    return _checkValidity(_number.reversed.toList());
  }

  int _replaceDigit(int number) {
    number *= 2;
    if (number > 9) number -= 9;
    return number;
  }

  bool _checkValidity(List<int> digits) {
    var _sum = 0;
    for (var i = 0; i < digits.length; i++) {
      _sum += (i + 1) % 2 == 0 ? _replaceDigit(digits[i]) : digits[i];
    }
    return _sum % 10 == 0;
  }
}
