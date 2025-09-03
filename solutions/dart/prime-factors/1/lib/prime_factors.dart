class PrimeFactors {
  List factors(int number) {
    var _factors = [];
    var _i = 2;
    while (number > 1) {
      if (number % _i == 0) {
        _factors.add(_i);
        number = number ~/ _i;
      } else {
        _i++;
      }
    }
    return _factors;
  }
}
