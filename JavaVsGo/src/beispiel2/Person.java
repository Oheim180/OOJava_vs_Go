package beispiel2;

public class Person {
	String name;
	public Person(String name) {
		this.name = name;
	}
	
	public String sprechen() {
		return "Ich bin Person " + this.name;
	}
}

