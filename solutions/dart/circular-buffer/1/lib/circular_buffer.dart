/* These implementation definitions enable zero-argument contructors:
 *    throw EmptyBufferException();
 */

class EmptyBufferException implements Exception {}

class FullBufferException implements Exception {}

class CircularBuffer {
  List<int> _buffer = [];
  int _readPointer = 0, _writePointer = 0, capacity;

  CircularBuffer(this.capacity) {
    for (var i = 0; i < capacity; i++) {
      _buffer.add(0);
    }
  }
  int read() {
    if (_buffer.every((element) => element == 0) || _buffer[_readPointer] == 0) {
      throw EmptyBufferException();
    }
    var i = _buffer[_readPointer];
    _buffer[_readPointer] = 0;
    _readPointer++;
    if (_readPointer >= capacity) {
      _readPointer = 0;
    }
    return i;
  }

  void write(int i, {bool force = false}) {
    if (_buffer[_writePointer] != 0) {
      if (!force) throw FullBufferException();
      _readPointer = _writePointer + 1;
      if (_readPointer >= capacity) {
        _readPointer = 0;
      }
    }
    _buffer[_writePointer] = i;

    _writePointer++;
    if (_writePointer >= capacity) {
      _writePointer = 0;
    }
  }

  void clear() => _buffer.fillRange(0, capacity, 0);
  // Put your code here
}
