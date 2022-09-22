# Lösungen
### Aufgabe 1
```
    func main() {
        fmt.Println(mult(5, 5))
    }

    func mult(x int, y int) int {

	    if x == 0 {
		    return 0
	    }

	    return mult(x-1, y) + y
    }
```

### Aufgabe 2
```
    func main() {
        fmt.Println(hoch(5, 2))
    }
    
    func hoch(x int, y int) int {
        if y == 0 {
            return 1
        }
    
        return x * hoch(x, y-1)
    }
```

### Aufgabe 3
```
    func main() {
        fmt.Println(fac(5))
    }
    
    func fac(x int) int {
        if x == 0 {
            return 1
        }
    
        return fac(x-1) * x
    }
```

### Aufgabe 4
```
    func main() {
        fmt.Println(reverse("test"))
    }
    
    func reverse(word string) string {
        if len(word) == 0 {
            return ""
        }
    
        return reverse(word[1:]) + word[:1]
    }
```

### Aufgabe 5
```
    func main() {
        fmt.Println(checkPalindrom("test"))
    }
    
    func checkPalindrom(word string) bool {
        if len(word) == 0 {
            return true
        }
    
        return checkPalindrom(word[1:len(word)-1]) && firstLetterIsLastLetter(word)
    }
    
    func firstLetterIsLastLetter(word string) bool {
        return word[:1] == word[len(word)-1:]
    }
```

### Aufgabe 6
```
    func main() {
        hanoi(3, 0, 1, 2)
    }
    
    func hanoi(turmhöhe int, turmposition int, hilfsposition int, turmziel int) {
    
        if turmhöhe == 0 {
            return
        }
    
        hanoi(turmhöhe-1, turmposition, turmziel, hilfsposition)
        fmt.Println(fmt.Sprintf("Schiebe die Scheibe von %v nach %v\n", turmposition, turmziel))
        hanoi(turmhöhe-1, hilfsposition, turmposition, turmziel)
    }
```