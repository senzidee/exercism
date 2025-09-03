
  String translate(String input) {
    if (input.isEmpty) {
      return '';
    }
    return input.split(' ').map((element) => _translateWord(element)).join(' ');
  }
  String _translateWord(String word) {
    if (word.startsWith('a') || word.startsWith('e') || word.startsWith('i') ||
        word.startsWith('o') || word.startsWith('u') || word.startsWith('yt') || word.startsWith('xr')) {
      return word + 'ay';
    } else {
      if (word.startsWith('thr') || word.startsWith('sch') || (word[1] == 'q' && word[2] == 'u')) {
        return word.substring(3) + word.substring(0, 3) + 'ay';
      }
      if (word.startsWith('ch') || word.startsWith('st') || word.startsWith('qu') || word.startsWith('th')) {
        return word.substring(2) + word.substring(0, 2) + 'ay';
      }
      if (word.length == 2 && word[1] == 'y') {
        return word.substring(1) + word.substring(0, 1) + 'ay';
      }
      if (word[1] == 'y' || word[2] == 'y') {
        return 'y' + word
            .split('y')
            .last + word
            .split('y')
            .first + 'ay';
      }
      return word.substring(1) + word.substring(0, 1) + 'ay';
    }
  }