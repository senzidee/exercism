import 'dart:math';

class ArmstrongNumbers {
  bool isArmstrongNumber(int number) {
    String asWord = number.toString();
    int j = asWord.length;
    int score = 0;
    for (var i = 0; i < j; i++) {
      score += pow(int.parse(asWord[i]), j).toInt();
    }
    print(score);
    if (score == number) {
      return true;
    }
    return false;
  }
}
