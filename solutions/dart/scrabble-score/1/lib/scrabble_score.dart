import 'dart:collection';

int score(String word) {
  final Map<String, int> points = HashMap();
  points.addAll({ 'A': 1, 'B': 3, 'C': 3, 'D': 2, 'E': 1, 'F': 4, 'G': 2,
    'H': 4, 'I': 1, 'J': 8, 'K': 5, 'M': 3, 'O': 1, 'P': 3, 'Q': 10, 'U': 1,
    'L': 1, 'N': 1, 'R': 1, 'S': 1, 'T': 1, 'V': 4, 'W': 4, 'Y': 4, 'X': 8,
    'Z': 10
  });
  int score = 0;
  if (word.length > 0) {
    word = word.toUpperCase();
    for (var i = 0; i < word.length; i++) {
      score += points[word[i]]!.toInt();
    }
  }
  return score;
}
