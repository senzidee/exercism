class ArmstrongNumbers {
  bool isArmstrongNumber(int number) {
    String asWord = number.toString();
    int j = asWord.length;
    var score = BigInt.from(0);
    for (var i = 0; i < j; i++) {
      score += BigInt.from(int.parse(asWord[i])).pow(j);
    }
    print(score);
    if (score == BigInt.from(number)) {
      return true;
    }
    return false;
  }
}
