class Triangle {
  bool equilateral(double side1, double side2, double side3) {
    return isTriangle(side1, side2, side3) && side1 == side2 && side2 == side3;
  }
  bool isosceles(double side1, double side2, double side3) {
    return isTriangle(side1, side2, side3) && (side1 == side2 || side2 == side3 || side1 == side3);
  }
  bool scalene(double side1, double side2, double side3) {
    return isTriangle(side1, side2, side3) && side1 != side2 && side2 != side3 && side1 != side3;
  }
  bool isTriangle(double side1, double side2, double side3) {
    return side1 + side2 > side3 && side1 + side3 > side2 && side2 + side3 > side1 && side1 > 0 && side2 > 0 && side3 > 0;
  }
}
