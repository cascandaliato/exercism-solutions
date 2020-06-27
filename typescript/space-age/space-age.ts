type Planet = 'Earth' | 'Jupiter' | 'Mars' | 'Mercury' | 'Neptune' | 'Saturn' | 'Uranus' | 'Venus';

class Space {
  private static readonly earthYearSeconds = 31557600;

  private static readonly earthOrbitalPeriods: Record<Planet, number> = {
    Earth: 1.0,
    Jupiter: 11.862615,
    Mars: 1.8808158,
    Mercury: 0.2408467,
    Neptune: 164.79132,
    Saturn: 29.447498,
    Uranus: 84.016846,
    Venus: 0.61519726,
  };

  constructor(public readonly seconds: number) {}

  onEarth = () => this.age('Earth');
  onJupiter = () => this.age('Jupiter');
  onMars = () => this.age('Mars');
  onMercury = () => this.age('Mercury');
  onNeptune = () => this.age('Neptune');
  onSaturn = () => this.age('Saturn');
  onUranus = () => this.age('Uranus');
  onVenus = () => this.age('Venus');

  private age(planet: Planet): number {
    const secondsInYear = Space.earthOrbitalPeriods[planet] * Space.earthYearSeconds;
    return Space.roundToTwoDecimals(this.seconds / secondsInYear);
  }

  private static roundToTwoDecimals(number: number): number {
    return Math.round(100 * number) / 100;
  }
}

export default Space;
