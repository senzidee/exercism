class Isogram {
  bool isIsogram (String words) {
    List<String> stoGram = [''];
    for (String element in words.split('')) {
      if (stoGram.contains(element.etoLowerCase())) {
        return false;
      }
      if (element.trim().isNotEmpty && element != '-') {
        stoGram.add(element.toLowerCase());
      }
    }
    return true;
  }
}
