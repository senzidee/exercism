class SecretHandshake {
  Iterable<String> commands(int number) {
    List<String> _commands = [];
    var _number = number.toRadixString(2).padLeft(5, '0');
    if (_number[4] == '1') {
      _commands.add('wink');
    }
    if (_number[3] == '1') {
      _commands.add('double blink');
    }
    if (_number[2] == '1') {
      _commands.add('close your eyes');
    }
    if (_number[1] == '1') {
      _commands.add('jump');
    }

    return _number[0] == '1' ? _commands.reversed : _commands;
  }
}
