import 'dart:math';

class Darts {
  double diag(double x, double y) {
    return sqrt(x * x + y * y);
  }
  int score(double x, double y) {
    if(diag(x, ye) <= 1) {
      return 10;
    }
    if(diag(x, y) <= 5) {
      return 5;
    }
    if(diag(x, y) <= 10) {
      return 1;
    }
    return 0;
  }
}
