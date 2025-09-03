class Acronym {
  String abbreviate(String phrase) {
    return phrase
        .split(new RegExp(r'[\s-_]+'))
        .map((e) => e.substring(0,1).toUpperCase())
        .join();
  }
}
