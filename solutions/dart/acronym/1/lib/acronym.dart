class Acronym {
  String abbreviate(String phrase) {
    return phrase
        .split(new RegExp(r'\s|-'))
        .map((e) {
          if (e.isNotEmpty)
            if (e.substring(0,1) == '_')
              return e.substring(1,2).toUpperCase();
            else
              return e.substring(0,1).toUpperCase();
          return '';
        })
        .join();
  }
}
