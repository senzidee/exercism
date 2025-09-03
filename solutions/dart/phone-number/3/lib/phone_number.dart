class PhoneNumber {
  dynamic clean(String number) {
    final re = RegExp(r'(\D+)');
    var numberDigits = number.split(re).reduce((value, element) => value + element);
    if (numberDigits.length < 10) {
      if (number.contains(RegExp(r'([a-zA-Z]+)'))) {
        throw FormatException('letters not permitted');
      }
      if (number.contains(RegExp(r'[^\w\s-]+'))) {
        throw FormatException('punctuations not permitted');
      }
      throw FormatException('must not be fewer than 10 digits');
    }
    if (numberDigits.length > 11) {
      throw FormatException('must not be greater than 11 digits');
    }
    if (numberDigits.length == 11) {
      if (numberDigits[0] != '1') {
        throw FormatException('11 digits must start with 1');
      }
      numberDigits = numberDigits.replaceFirst('1', '');
    }
    if (numberDigits[0] == '0') {
      throw FormatException('area code cannot start with zero');
    }
    if (numberDigits[0] == '1') {
      throw FormatException('area code cannot start with one');
    }
    if (numberDigits[3] == '0') {
      throw FormatException('exchange code cannot start with zero');
    }
    if (numberDigits[3] == '1') {
      throw FormatException('exchange code cannot start with one');
    }
    return numberDigits;
  }
}
