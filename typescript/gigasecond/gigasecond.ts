export default class Gigasecond {
  constructor(private start: Date) {}

  date(): Date {
    const end: Date = new Date(this.start);
    end.setSeconds(end.getSeconds() + 10 ** 9);
    return end;
  }
}
