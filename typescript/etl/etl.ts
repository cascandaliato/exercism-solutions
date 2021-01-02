type LettersByScore = { [key: number]: string[] };
type ScoreOfLetter = { [key: string]: number };

function transform(byScore: LettersByScore): ScoreOfLetter {
  const byLetter: ScoreOfLetter = {};

  for (const [score, letters] of Object.entries(byScore))
    for (const letter of letters)
      byLetter[letter.toLowerCase()] = Number(score);

  return byLetter;
}

export default transform;
