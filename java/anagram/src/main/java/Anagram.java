import java.util.Arrays;
import java.util.Collection;
import java.util.stream.Collectors;

public class Anagram {
  private String word;
  private String normalized;

  public Anagram(String word) {
    this.word = word;
    this.normalized = normalize(word);
  }

  public Collection<String> match(Collection<String> candidates) {
    return candidates.stream().filter(c -> !c.equalsIgnoreCase(this.word)).filter(this::isAnagram)
        .collect(Collectors.toList());
  }

  private boolean isAnagram(String otherWord) {
    return normalize(otherWord).equals(this.normalized);
  }

  private String normalize(String word) {
    char[] arr = word.toLowerCase().toCharArray();
    Arrays.sort(arr);
    return new String(arr);
  }
}