class WordCount {
  Map<String, int> countWords(String input) {
    Map<String, int> counts = {};
    var matcher = new RegExp(r"(\w+(?:\'\w+)*)");
    matcher.allMatches(input).forEach((item) {
      if (!counts.containsKey(item.group(0)!)) {
        counts[item.group(0)!] = 1;
      } else {
        counts[item.group(0)!] = counts[item.group(0)!]! + 1;
      }
    });

    return counts;
  }
}
