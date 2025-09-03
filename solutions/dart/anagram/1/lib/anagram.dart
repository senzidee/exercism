class Anagram {
  List<String> findAnagrams(String target, List<String> candidates) {
    List<String> anagrams = [];
    var _target = target.toLowerCase().split('');
    _target.sort();
    candidates.forEach((candidate) {
      var _candidate = candidate.toLowerCase().split('');
      _candidate.sort();
      if (_candidate.join() == _target.join() && target.toLowerCase() != candidate.toLowerCase()) {
        anagrams.add(candidate);
      }
    });
    return anagrams;
  }
}
