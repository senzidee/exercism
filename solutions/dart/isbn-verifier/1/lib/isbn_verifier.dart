// Put your code here
bool isValid(String isbn) {
  isbn = isbn.replaceAll('-', '');
  if (
  isbn.length < 10 ||
  isbn.length > 10 ||
  isbn.contains(new RegExp(r'[^0-9xX]')) ||
      (isbn.contains(new RegExp(r'[xX]')) && isbn[9].toUpperCase() != 'X')
  ) {
    return false;
  }
  return _checkDigit(isbn);
}

bool _checkDigit(String isbn) {
  var sum = 0;
  for (var i = 0; i < 9; i++) {
    sum += int.parse(isbn[i]) * (10 - i);
  }
  if (isbn[9].toUpperCase() != 'X') {
    sum += int.parse(isbn[9]);
  } else {
    sum += 10;
  }

  return sum % 11 == 0;
}