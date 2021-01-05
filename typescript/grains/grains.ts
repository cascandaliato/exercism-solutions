function square(n: number): number {
  if (n < 1 || n > 64) throw new Error("Input must be in the range [1, 64]");

  return 2 ** (n - 1);
}

function total(): number {
  let t = 0;
  for (let i = 1; i <= 64; i++) t += square(i);
  return t;
}

export default { square, total };
