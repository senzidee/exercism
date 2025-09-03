import 'dart:math';

class ArmstrongNumbers {
  bool isArmstrongNumber(int number) {
    String asWord = number.toString();
    return number == asWord
        .split('')
        .map((i) => pow(int.parse(i), asWord.length))
        .reduce((val, el) => val + el);
  }
}
