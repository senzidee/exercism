import 'dart:math';

class DifferenceOfSquares {

  int squareOfSum(int num) {
    return pow(summa(0, num), 2).toInt();
  }

  int sumOfSquares(int num) {
    return summaq(0, num);
  }

  int differenceOfSquares(int num) {
    return squareOfSum(num) - sumOfSquares(num);
  }

  int summa(int sum, int num) {
    if (num == 0) {
      return sum;
    }
    sum += num--;
    return summa(sum, num);
  }

  int summaq(int sum, int num) {
    if (num == 0) {
      return sum;
    }
    sum += pow(num--, 2).toInt();
    return summaq(sum, num);
  }
}
