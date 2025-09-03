class BeerSong {
  List<String> recite(int bottles, int stanzas) {
    List<String> result = [];
    for (var i = 0; i < stanzas; i++) {
      if (result.isNotEmpty) {
        result.add('');
      }
      result.add('${bottles == 0 ? 'No more' : bottles} bottle${bottles != 1 ? 's' : ''} of beer on the wall, ${bottles == 0 ? 'no more' : bottles} bottle${bottles != 1 ? 's' : ''} of beer.');
      if (bottles == 0) {
        result.add('Go to the store and buy some more, 99 bottles of beer on the wall.');
      } else {
        result.add('Take ${bottles != 1 ? 'one' : 'it'} down and pass it around, ${bottles == 1 ? 'no more' : bottles -1} bottle${bottles != 2 ? 's' : ''} of beer on the wall.');
      }
      bottles--;
    }
    return result;
  }
}

