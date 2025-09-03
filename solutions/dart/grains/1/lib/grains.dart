BigInt square(final int n) {
  if (n < 1 || n > 64) {
    throw ArgumentError('square must be between 1 and 64');
  }
  if (n == 64) {
    // It's because that's how the bitshift operation works on x86 assembly. For
    // example, in 32bit operants, the count is masked to 5 bits, for 31 bit
    // maximum shift (which clears the register, so there's no point in allowing
    // bigger shift).
    return new BigInt.from(1 << n -2) * new BigInt.from(2);
  }
  return new BigInt.from(1 << n -1);
}

BigInt total() {
  var sum = new BigInt.from(0);
  for (var i = 1; i <= 64; i++) {
    sum += square(i);
  }
  return sum;
}
