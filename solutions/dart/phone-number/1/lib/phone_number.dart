class PhoneNumber {
  dynamic clean(String number) {
    final re = RegExp(r'(\D+)');
    number = number.split(re).reduce((value, element) => value + element);
    if (number.length < 10) {
      throw FormatException('incorrect number of digits');
    }
    if (number.length > 11) {
      throw FormatException('more than 11 digits');
    }
    if (number.length == 11) {
      if (number[0] != '1') {
        throw FormatException('11 digits must start with 1');
      }
      number = number.replaceFirst('1', '');
    }
    if (number[0] == '0') {
      throw FormatException('area code cannot start with zero');
    }
    if (number[0] == '1') {
      throw FormatException('area code cannot start with one');
    }
    if (number[3] == '0') {
      throw FormatException('exchange code cannot start with zero');
    }
    if (number[3] == '1') {
      throw FormatException('exchange code cannot start with one');
    }
    return number;
  }
}
