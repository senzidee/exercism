class Bob {
  String response(String input) {
    var regexAsk = RegExp(r'\?');
    if (regexAsk.hasMatch(input) && input == input.toUpperCase()) {
      return "Calm down, I know what I'm doing!";
    }
    if (regexAsk.hasMatch(input)) {
      return "Sure";
    }
    if (input == input.toUpperCase()) {
      return "Whoa, chill out!";
    }
    if (input.length > 0) {
      return "Whatever.";
    }
    return "Fine. Be that way!";
  }
}
