class Bob {
  private readonly answers: {
    pattern: RegExp;
    answer: string;
  }[] = [
    // forceful question
    { pattern: /^[^a-z]*[A-Z][^a-z]*\?\s*$/, answer: "Calm down, I know what I'm doing!" },
    // regular question
    { pattern: /^.*[a-z0-9].*\?\s*$/, answer: 'Sure.' },
    // yelling
    { pattern: /^[^a-z]*[A-Z][^a-z\?]*$/, answer: 'Whoa, chill out!' },
    // nothing
    { pattern: /^\s*$/, answer: 'Fine. Be that way!' },
  ];

  hey(sentence: string): string {
    for (const { pattern, answer } of this.answers) {
      if (pattern.test(sentence)) return answer;
    }
    return 'Whatever.';
  }
}

export default Bob;
