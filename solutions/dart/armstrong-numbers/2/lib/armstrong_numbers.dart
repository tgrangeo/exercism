class ArmstrongNumbers {
  bool isArmstrongNumber(String str) {
    BigInt res = BigInt.zero;
    int length = str.length;

    for (int i = 0; i < length; i++) {
      BigInt digit = BigInt.parse(str[i]);
      res += _bigIntPow(digit, length);
    }

    return res == BigInt.parse(str);
  }

  BigInt _bigIntPow(BigInt base, int exponent) {
    BigInt result = BigInt.one;
    for (int i = 0; i < exponent; i++) {
      result *= base;
    }
    return result;
  }
}
