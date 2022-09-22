# Übungsaufgaben
## Aufgabe 1 (Schwierigkeit 1)
Die Multiplikation kann man auch als Addition von Zahlen darstellen:

5 * 5 = 5 + 5 + 5 + 5 + 5

Implementiere eine Funktion, welche 2 Zahlen multipliziert, ohne
Verwendung des Multiplikationsoperators.

Du kannst mit folgendem code starten:
```
    func main() {
        fmt.Println(mult(5,5))
    }
    
    func mult(x int, y int) int {
        ...
    } 
```

## Aufgabe 2 (Schwierigkeit 1)
Die Potenzfunction kann man auch als Multiplikation darstellen:

5^2 = 5*5

Implementiere die function hoch. Der erste parameter gibt die Zahl an und der zweite den Exponenten.

Du kannst mit folgendem code starten:
```
    func main() {
        fmt.Println(hoch(5,2))
    }
    
    func hoch(x int, y int) int {
        ...
    } 
```

### Aufgabe 3 (Schwierigkeit 1)
Die Fakultät ist die Multiplikation von allen Zahlen davor.

fac(0) = 1

fac(x) = x * fac(x-1)

fac(5) = 5 * 4 * 3 * 2 * 1 * 1

Du kannst mit folgendem code anfangen:
```
    func main() {
        fmt.Println(fac(5))
    }
    
    func fac(x int) int {
        ...
    } 
```

## Aufgabe 4 (Schwierigkeit 2)
Die Reihenfolge von Buchstaben umdrehen.

Tipp: 

Zwei Strings können mit dem + Operator zusammengefügt werden.

Mit len(string) wird die länge eines string zurück gegeben.

Mit [x:y] können gewisse zeichen von ein String zurück gegeben werden,.

Bsp.:

```
    x := "Test1234"
    x[0:1] == "T" oder einfach x[:1] === "T"
    x[1:len(x)] == "est1234" oder einfach x[1:] == "est1234"
```

Du kannst mit folgendem code starten:
```
    func main() {
        fmt.Println(reverse("test"))
    }
    
    func reverse(word string) string {
        ...
    } 
```

## Aufgabe 5 (Schwierigkeit 3)
Palindrom erkennen
Ein Palindrom ist ein Wort, welches rückwärts und vorwärts gleich ist.

Bsp.: "abba", "anna", "lagerregal"

Tipps:
 Siehe Aufgabe 2

Du kannst mit folgendem code starten:
```
    func main() {
        fmt.Println(checkPalindrom("anna"))
    }
    
    func checkPalindrom(word string) bool {
        ...
    } 
    
    func firstLetterIsLastLetter(word string) bool {
	 return word[:1] == word[len(word)-1:]
    }
```

## Aufgabe 6 (Knobel)
Türme von Hanoi.
Lies dir die Spielregeln von den Türmen von Hanoi durch: https://www.bernhard-gaul.de/spiele/tower/tower.php

Die Türme von Hanoi lassen sich ziemlich einfach mit rekursion lösen.
Male dir doch hierzu einmal die Türme von Hanoi mit 1 Scheibe auf und versuche sie zu lösen.
Wiederhole dies für die Türme von Hanoi mit 2 und 3 Scheibe. Fällt dir etwas auf?

Um 3 Scheiben auf den rechten Platz zu bewegen, muss man nur 2 Scheiben in die Mitte bewegen, danach die 3. Scheibe auf den letzten Platz
und 2 Scheiben von der Mitte auf den rechten Platz.

Vorgehen:
Die func hanoi soll die Schritte zu Lösen in die Konsole schreiben.

Bsp.:
"Schiebe die Scheibe von 0 nach 2"

Tipp:
Nutze dafür folgendes Statement:

```fmt.Println(fmt.Sprintf("Schiebe die Scheibe von %v nach %v\n", turmposition, turmziel))```

Wir sprechen später darüber, was das macht.

Überlege dir, wann die function abbrechen könnte. Welche Lösung kennst du? 

Überlege dir, was du machen würdest. 
* Wenn du einen Turm von Hanoi 3 von Platz 0 auf Platz 2 Schieben möchtest,
* dann baust du einen Turm von Hanoi 2 von Platz 0 auf Platz 1
* danach schiebst du die letzte Scheibe auf Platz 2
* Und baust du einen Turm von Hanoi 2 von Platz 1 auf Platz 2

Die 3 Positionen bekommst du bei der Funktion mit.

Du kannst mit folgendem Code starten:
```
 func main() {
     hanoi(3, 0, 1, 2)
 }
 
 func hanoi(turmhöhe int, turmposition int, hilfsposition int, turmziel int) {
 }
```