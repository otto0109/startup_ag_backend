# Rekursion

```Um Rekursion zu verstehen, muss man Rekursion verstanden haben``` (Daniel Alpers, 2017)

Dieser obere Satz fasst eigentlich ganz gut zusammen, was Rekursion ist.

Aber mal in anderen Worten, Rekursion beschreibt die Möglichkeit ein Problem zu lösen,
indem man die Größe des Problems verringert.

In der Informatik nennt man dieses Prinzip, bzw. das Vorgehen auch: 

```divide and conquer / teile und heersche```

Kleinere Probleme sind oft einfacher zu lösen.

---

Rekursive Funktionen versuchen also das Problem so weit zu verkleiner, bis wir die Lösung kennen.

Rekursive Funktionen rufen sich immer selber wieder auf und brauchen immer
eine Abbruchbedingung.

Allgemein:

```
func recursive(parameter) rückgabe {
    if (<Abbruchbedingung>) {
        return irgendwas
    }
    
    return recursive(<kleineres problem>)
}
```
Beispiel (fibonacci):
```
func fib(n int) int {
    if n == 1 || n == 2 {
        return 0
    }
    return fib(n-1) + fib(n-2)
}
```
