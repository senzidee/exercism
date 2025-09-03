class Forth {
  List<int> stack = [];
  Map<String, String> custom = {};

  void evaluate(String input) {
    input = input.toUpperCase();
    if (input.startsWith(':')) {
      var _input = input.split(' ');
      _input.removeAt(0);
      _input.removeLast();
      var _command = _input[0];
      if (int.tryParse(_command) != null) {
        throw Exception('Invalid definition');
      }
      _input.removeAt(0);
      custom[_command] = _input
          .map((el) => _explainCustomCommand(el))
          .join(' ');
    } else {
      input
          .split(' ')
          .map((el) => _explainCustomCommand(el))
          .toList()
          .join(' ')
          .split(' ')
          .map((el) => _computeStack(el))
          .toList();
    }
  }

  void _canArithmeticAction() {
    if (stack.length < 2) {
      throw Exception('Stack empty');
    }
  }

  String _explainCustomCommand(String customCommand) {
    if (custom.containsKey(customCommand)) {
      return custom[customCommand]!;
    }
    return customCommand;
  }
  
  void _computeStack(String el) {
    if (int.tryParse(el) != null) {
      stack.add(int.parse(el));
    } else {
      switch (el) {
        case '+':
          _canArithmeticAction();
          var t = stack.last;
          stack.removeLast();
          t += stack.last;
          stack.removeLast();
          stack.add(t);
          break;
        case '-':
          _canArithmeticAction();
          var t = stack.last;
          stack.removeLast();
          t = stack.last - t;
          stack.removeLast();
          stack.add(t);
          break;
        case '*':
          _canArithmeticAction();
          var t = stack.last;
          stack.removeLast();
          t *= stack.last;
          stack.removeLast();
          stack.add(t);
          break;
        case '/':
          _canArithmeticAction();
          var t = stack.last;
          if (t == 0) {
            throw Exception('Division by zero');
          }
          stack.removeLast();
          t = stack.last ~/ t;
          stack.removeLast();
          stack.add(t);
          break;
        case 'DUP':
          if (stack.length == 0) {
            throw Exception('Stack empty');
          }
          stack.add(stack.last);
          break;
        case 'DROP':
          if (stack.length == 0) {
            throw Exception('Stack empty');
          }
          stack.removeLast();
          break;
        case 'SWAP':
          if (stack.length < 2) {
            throw Exception('Stack empty');
          }
          var t1 = stack.last;
          stack.removeLast();
          var t2 = stack.last;
          stack.removeLast();
          stack.add(t1);
          stack.add(t2);
          break;
        case 'OVER':
          if (stack.length < 2) {
            throw Exception('Stack empty');
          }
          var t1 = stack.last;
          stack.removeLast();
          var t2 = stack.last;
          stack.add(t1);
          stack.add(t2);
          break;
        default:
          throw Exception('Unknown command');
      }
    }
  }
}
