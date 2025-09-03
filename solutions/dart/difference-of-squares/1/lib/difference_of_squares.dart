import 'dart:math';

class DifferenceOfSquares {

  int squareOfSum(int num) {
    int sum = 0;
    for (var i = num; i > 0; i--) {
      sum += i;
    }
    return pow(sum, 2).toInt();
  }

  int sumOfSquares(int num) {
    int sum = 0;
    for (var i = num; i > 0; i--) {
      sum += pow(i, 2).toInt();
    }
    return sum;
  }

  int differenceOfSquares(int num) {
    return squareOfSum(num) - sumOfSquares(num);
  }
}
