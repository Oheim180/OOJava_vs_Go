package beispiel1;

public class Kunde extends Person{
	int kundenNr;
	public Kunde(String name) {
		super(name);
	}
	@Override
	public String sprechen() {
		return "Ich bin Kunde " + this.name;
	}
}

