export default class PhoneNumber {
  private static readonly RE: RegExp = /^(?:\+?1)?\s*\(?(\d{3})\)?[\s\.-]*(\d{3})[\s\.-]*(\d{4})\s*$/;

  constructor(private phoneNumber: string) {}

  number(): string | undefined {
    const match = this.phoneNumber.match(PhoneNumber.RE);
    return match ? `${match[1]}${match[2]}${match[3]}` : undefined;
  }
}
