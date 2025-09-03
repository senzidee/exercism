class RnaTranscription {
  Map<String, String> transcription = {
    'G': 'C',
    'C': 'G',
    'T': 'A',
    'A': 'U'
  };
  String toRna(String nucleotide) {
    if (nucleotide.isEmpty) return '';
    return nucleotide.split('').fold('', (previousValue, element) => previousValue + transcription[element]!);
  }
}
