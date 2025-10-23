BigInt square(final int n) {
  BigInt res = BigInt.from(1);
  if (n > 64 || n < 1 ){
    throw ArgumentError('square must be between 1 and 64');
    return res;
  }
  for (int i = 1; i < n; i+= 1){
    res = res * BigInt.from(2);
  }
  return res;
}

BigInt total() {
  return square(64) + square(64) - BigInt.from(1);
}
