
class InvalidNucleotideException implements Exception {
  final String invalidChar;

  InvalidNucleotideException(this.invalidChar);

  @override
  String toString() => 'InvalidNucleotideException: invalid nucleotide "$invalidChar"';
}

class NucleotideCount {
  Map<String, int> count(String str) {
    var map = {"A": 0, "C": 0, "G": 0, "T": 0};
    for (var c in str.split('')) {
      if (map.containsKey(c)) {
        map[c] = map[c]! + 1;
      } else {
        throw InvalidNucleotideException(c);
      }
    }
    return map;
  }
}
