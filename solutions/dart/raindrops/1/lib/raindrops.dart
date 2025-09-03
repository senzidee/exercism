class Raindrops {
  Map<int, String> divisibleBy = {
    3: 'Pling',
    5: 'Plang',
    7: 'Plong'
  };

  String convert(int input) {
    var result = '';
    divisibleBy.forEach((key, value) {
      if (input % key == 0) {
        result += value;
      }
    });
    if (result.isEmpty) {
      result = input.toString();
    }
    return result;
  }
}
