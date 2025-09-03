class SecretHandshake {
  Iterable<String> commands(int number) {
    List<String> _commands = [];
    if (number & 1 == 1) {
      _commands.add('wink');
    }
    if (number & 2 == 2) {
      _commands.add('double blink');
    }
    if (number & 4 == 4) {
      _commands.add('close your eyes');
    }
    if (number & 8 == 8) {
      _commands.add('jump');
    }

    return number & 16 == 16 ? _commands.reversed : _commands;
  }
}
