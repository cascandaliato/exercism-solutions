class Bob {
  hey(sentence: string): string {
    if (this.yelling(sentence)) {
      return 'Whoa, chill out!';
    } else if (this.forcefulQuestion(sentence)) {
      return "Calm down, I know what I'm doing!";
    } else if (this.regularQuestion(sentence)) {
      return 'Sure.';
    } else if (this.isEmptySentence(sentence)) {
      return 'Fine. Be that way!';
    } else {
      return 'Whatever.';
    }
  }

  private yelling(sentence: string): boolean {
    return this.isAllUppercase(sentence) && !this.endsWithAQuestionMark(sentence);
  }

  private forcefulQuestion(sentence: string): boolean {
    return this.endsWithAQuestionMark(sentence) && this.isAllUppercase(sentence);
  }

  private regularQuestion(sentence: string): boolean {
    return this.endsWithAQuestionMark(sentence) && !this.isAllUppercase(sentence);
  }

  private isEmptySentence(sentence: string): boolean {
    return sentence.trim() === '';
  }

  private isAllUppercase(sentence: string): boolean {
    return this.containsLetter(sentence) && sentence.toUpperCase() === sentence;
  }

  private endsWithAQuestionMark(sentence: string): boolean {
    return sentence.trim().endsWith('?');
  }

  private containsLetter(sentence: string): boolean {
    return /[a-zA-Z]/.test(sentence);
  }
}

export default Bob;
