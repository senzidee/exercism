extension ConvertToRomanNumerals on int {
  String toRoman() {
    return _convert(this ~/ 1000, 'M', '', '') +
        _convert((this ~/ 100) % 10, 'C', 'D', 'M') +
        _convert((this ~/ 10 % 10), 'X', 'L', 'C') +
        _convert(this % 10, 'I', 'V', 'X');

  }
  String _convert(int number, String ones, String fives, String tens) {
    switch (number) {
      case 1:
        return ones;
      case 2:
        return ones + ones;
      case 3:
        return ones + ones + ones;
      case 4:
        return ones + fives;
      case 5:
        return fives;
      case 6:
        return fives + ones;
      case 7:
        return fives + ones + ones;
      case 8:
        return fives + ones + ones + ones;
      case 9:
        return ones + tens;
      default:
        return '';
    }
  }
}
