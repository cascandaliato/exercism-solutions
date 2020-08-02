class ComplexNumber {
  constructor(public real: number, public imag: number) {}

  get conj(): ComplexNumber {
    return new ComplexNumber(this.real, this.imag === 0 ? 0 : -this.imag);
  }

  get abs(): number {
    return Math.sqrt(this.real ** 2 + this.imag ** 2);
  }

  get exp(): ComplexNumber {
    return new ComplexNumber(Math.exp(this.real), 0).mul(
      new ComplexNumber(Math.cos(this.imag), Math.sin(this.imag)),
    );
  }

  add(other: ComplexNumber): ComplexNumber {
    return new ComplexNumber(this.real + other.real, this.imag + other.imag);
  }

  sub(other: ComplexNumber): ComplexNumber {
    return new ComplexNumber(this.real - other.real, this.imag - other.imag);
  }

  mul(other: ComplexNumber): ComplexNumber {
    return new ComplexNumber(
      this.real * other.real - this.imag * other.imag,
      this.imag * other.real + this.real * other.imag,
    );
  }

  div(other: ComplexNumber): ComplexNumber {
    return this.mul(other.conj).mul(new ComplexNumber(1 / other.abs ** 2, 0));
  }
}

export default ComplexNumber;
