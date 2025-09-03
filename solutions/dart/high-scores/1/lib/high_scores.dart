class HighScores {
  List<int> scores;

  HighScores(this.scores);

  int latest() {
    return this.scores.last;
  }

  int personalBest() {
    return this.scores.reduce((value, element) => (element > value) ? element : value);
  }

  List<int> personalTopThree() {
    this.scores.sort((a, b) => b.compareTo(a));
    return this.scores.getRange(0, (this.scores.length < 3) ? this.scores.length : 3).toList();
  }
}
