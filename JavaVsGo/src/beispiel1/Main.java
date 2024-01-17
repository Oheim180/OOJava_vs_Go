package beispiel1;

public class Main {	
	public static void main(String[] args) {
		Person p = new Person("Max");
        Mitarbeiter m = new Mitarbeiter("Bernd");
        Kunde k = new Kunde("Herbert");
    
        methodeSprechen(p);
        // Folgegendde Aufrufe sind auch möglich, da Mitarbeiter und Kunde Untertypen von Person sind.
        methodeSprechen(m); 
        methodeSprechen(k);
    }
	
	/**
	 * Fuehrt die sprechne-Methode des Übergebenen Objekts aus.
	 * 
	 * @param p erwartet als parameter ein Objekt von Typ Person.
	 */
	public static void methodeSprechen(Person p) {
		System.out.println(p.sprechen());
	}
}
