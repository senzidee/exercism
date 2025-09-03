extension Ex on double {
  double toPrecision(int n) => double.parse(toStringAsFixed(n));
}

class SpaceAge {
  Map<String, double> planetToEarth = {
    "Mercury": 0.2408467,
    "Venus": 0.61519726,
    "Mars": 1.8808158,
    "Jupiter": 11.862615,
    "Saturn": 29.447498,
    "Uranus": 84.016846,
    "Neptune": 164.79132,
  };
  Map<String, double> earth = {
    "days": 365.25,
    "seconds": 31557600,
  };
  double age({String? planet, int? seconds}) {
    if (planetToEarth.containsKey(planet)) {
      return planetToEarth[planet]! * earth['seconds']!;
    }
    if (planet == "Earth") {
      return (seconds! / earth['seconds']!).toPrecision(2);
    }
    return 0.0;
  }
}
