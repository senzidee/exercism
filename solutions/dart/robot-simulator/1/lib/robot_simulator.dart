import 'package:robot_simulator/orientation.dart';
import 'package:robot_simulator/position.dart';

class Robot {
  Position position;
  Orientation orientation;

  Robot(this.position, this.orientation);

  void move(String s) {
    s.split('').forEach((element) {
      switch (element) {
        case 'A':
          this.advance();
          break;
        case 'L':
          this.turnLeft();
          break;
        case 'R':
          this.turnRight();
          break;
      }
    });
  }

  void turnLeft() {
    switch (this.orientation) {
      case Orientation.east:
        this.orientation = Orientation.north;
        break;
      case Orientation.north:
        this.orientation = Orientation.west;
        break;
      case Orientation.west:
        this.orientation = Orientation.south;
        break;
      case Orientation.south:
        this.orientation = Orientation.east;
        break;
    }
  }
  void turnRight() {
    switch (this.orientation) {
      case Orientation.east:
        this.orientation = Orientation.south;
        break;
      case Orientation.north:
        this.orientation = Orientation.east;
        break;
      case Orientation.west:
        this.orientation = Orientation.north;
        break;
      case Orientation.south:
        this.orientation = Orientation.west;
        break;
    }
  }
  void advance() {
    switch (this.orientation) {
      case Orientation.east:
        this.position = Position(this.position.x + 1, this.position.y);
        break;
      case Orientation.north:
        this.position = Position(this.position.x, this.position.y + 1);
        break;
      case Orientation.west:
        this.position = Position(this.position.x - 1, this.position.y);
        break;
      case Orientation.south:
        this.position = Position(this.position.x, this.position.y - 1);
        break;
    }
  }
}
