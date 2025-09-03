class Sieve {
  Map<int, bool> _primeNumbers = {};
  Sieve(int number) {
    for (var i = 2; i <= number; i++) {
      if (!_primeNumbers.containsKey(i)) _primeNumbers[i] = true;
      if (_primeNumbers[i] != true) continue;
      for (var j = i + i; j <= number; j += i) {
        _primeNumbers[j] = false;
      }
    }
  }
  get primes => _primeNumbers.isEmpty ? [] : _primeNumbers.keys.where((i) => _primeNumbers[i] == true).toList();
}
