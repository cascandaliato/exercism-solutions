const letterValues = {
  a: 1,
  b: 3,
  c: 3,
  d: 2,
  e: 1,
  f: 4,
  g: 2,
  h: 4,
  i: 1,
  j: 8,
  k: 5,
  l: 1,
  m: 3,
  n: 1,
  o: 1,
  p: 3,
  q: 10,
  r: 1,
  s: 1,
  t: 1,
  u: 1,
  v: 4,
  w: 4,
  x: 8,
  y: 4,
  z: 10,
} as const;

type Letter = keyof typeof letterValues;

const hasScore = (ch: string): ch is Letter => {
  return ch in letterValues;
};

export default (word?: string): number => {
  if (word === undefined) {
    return 0;
  }

  return [...word.toLowerCase()].reduce(
    (score, ch) => score + (hasScore(ch) ? letterValues[ch] : 0),
    0,
  );
};
