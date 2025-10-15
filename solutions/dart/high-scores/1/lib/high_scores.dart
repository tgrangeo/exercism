class HighScores {
  final List<int> scores;

  HighScores(this.scores);

  int latest() {
    return scores.last;
  }

  int personalBest() {
    return scores.reduce((a, b) => a > b ? a : b);
  }

  List<int> personalTopThree() {
    final sorted = [...scores]..sort((b, a) => a.compareTo(b));
    return sorted.take(3).toList();
  }
}
