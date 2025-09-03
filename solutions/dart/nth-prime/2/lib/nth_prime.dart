extension _ on List<int> {
  void next_prime_number() {
    var next = last + 1;
    while(true) {
      if (sublist(2).firstWhere((element) => next % element == 0, orElse: () => -1) == -1) {
        add(next);
        break;
      }
      next++;
    }
  }
}
class NthPrime {
  int prime(int position) {
    switch(position) {
      case 0:
        throw new ArgumentError("There is no zeroth prime");
      case 1:
        return 2;
      case 2:
        return 3;
      case 3:
        return 5;
      case 4:
        return 7;
      case 5:
        return 11;
      case 6:
        return 13;
      default:
        var prime_numbers = [2, 3, 5, 7, 11, 13];
        var iterator = 2;
        while(position > iterator) {
          prime_numbers.next_prime_number();
          iterator++;
        }
        return prime_numbers.last;
    }
  }
}
