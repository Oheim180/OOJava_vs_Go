package beispiel1;

public class Main {	
	public static void main(String[] args) {
		Person p = new Person("Max");
        Mitarbeiter m = new Mitarbeiter("Bernd");
        Kunde k = new Kunde("Herbert");

        if (p instanceof Mitarbeiter) {//Hier: False
            System.out.println("Die Person " + p.name + " ist ein Mitarbeiter.");
        } else {
            System.out.println("Die Person " + p.name + " ist kein Mitarbeiter.");
        }
        if (m instanceof Person) {//Hier: True
            System.out.println("Die Mitarbeiter " + m.name + " ist eine Person.");
        } else {
            System.out.println("Der Mitarbeiter " + m.name + " ist keine Person.");
        }

        if (k instanceof Person) {//Hier: True
            System.out.println("Der Kunde " + k.name + " ist eine Person.");
        } else {
            System.out.println("Der Kunde " + k.name + " ist keine Person.");
        }
    }
}
