class Allergies {
  static const Map<String, int> alle = {
    'eggs': 1,
    'peanuts': 2,
    'shellfish': 4,
    'strawberries': 8,
    'tomatoes': 16,
    'chocolate': 32,
    'pollen': 64,
    'cats': 128
  };

  bool allergicTo(String item, int score) => score & alle[item]! != 0;

  List<String> list(int score) {
    List<String> list = [];
    alle.forEach((key, value) {
      if (score & value != 0) {
        list.add(key);
      }
    });
    return list;
  }
}
