class Bob {
  hey(sentence: string): string {
    if (this.forcefulQuestion(sentence)) {
      return "Calm down, I know what I'm doing!";
    } else if (this.question(sentence)) {
      return 'Sure.';
    } else if (this.yelling(sentence)) {
      return 'Whoa, chill out!';
    } else if (this.silence(sentence)) {
      return 'Fine. Be that way!';
    } else {
      return 'Whatever.';
    }
  }

  private forcefulQuestion(sentence: string): boolean {
    return this.yelling(sentence) && this.question(sentence);
  }

  private yelling(sentence: string): boolean {
    return sentence.toUpperCase() === sentence && this.containsWord(sentence);
  }

  private question(sentence: string): boolean {
    return sentence.trim().endsWith('?');
  }

  private silence(sentence: string): boolean {
    return sentence.trim() === '';
  }

  private containsWord(sentence: string): boolean {
    return /[a-zA-Z]+/.test(sentence);
  }
}

export default Bob;
