import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;

public class Gigasecond {
  private static final long GIGASECOND = (long)Math.pow(10, 9);

  private final LocalDateTime initialMoment;

  public Gigasecond(LocalDate moment) {
    this.initialMoment = LocalDateTime.of(moment, LocalTime.MIDNIGHT);
  }

  public Gigasecond(LocalDateTime moment) { this.initialMoment = moment; }

  public LocalDateTime getDateTime() {
    return this.initialMoment.plusSeconds(Gigasecond.GIGASECOND);
  }
}
