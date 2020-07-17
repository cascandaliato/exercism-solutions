import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;

public class Gigasecond {
  private static final long GIGASECOND = (long)Math.pow(10, 9);

  private final LocalDateTime endMoment;

  public Gigasecond(LocalDate moment) {
    this(LocalDateTime.of(moment, LocalTime.MIDNIGHT));
  }

  public Gigasecond(LocalDateTime moment) {
    this.endMoment = moment.plusSeconds(Gigasecond.GIGASECOND);
  }

  public LocalDateTime getDateTime() { return this.endMoment; }
}
