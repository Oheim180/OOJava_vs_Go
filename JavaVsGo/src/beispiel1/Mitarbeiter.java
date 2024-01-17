package beispiel1;

public class Mitarbeiter extends Person {
	int mitarbeiterNr;
	public Mitarbeiter(String name) {
		super(name);
	}
	@Override
	public String sprechen() {
		return "Ich bin Mitarbeiter " + this.name;
	}
}
