package beispiel2;

public class Main {	
	public static void main(String[] args) {
		Person p = new Person("Max");
        Mitarbeiter m = new Mitarbeiter("Bernd");
        Kunde k = new Kunde("Herbert");

        System.out.println(p.sprechen());
        System.out.println(m.sprechen());
        System.out.println(k.sprechen());
    }
}
