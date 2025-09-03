class Pangram {
  bool isPangram(String input) {
    if (input.length < 26) {
      return false;
    }
    input = input.toLowerCase();
    for (var i = 97; i < 123; i++) {
      if (!input.contains(String.fromCharCode(i))) {
        return false;
      }
    }
    return true;
  }
}
