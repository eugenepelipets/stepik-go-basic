// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin) // создаем экземпляр структуры bufio.Scanner для чтения из консоли
// 	fmt.Println("Введите строку с пробелами:")
// 	_ = scanner.Scan()                                           // приложение остановится, пока пользователь не введет строку и не нажмет Enter
// 	input := scanner.Text()                                      // сохраняем введенную строку целиком в переменную name
// 	fmt.Println("Вы ввели:", input, "длина строки:", len(input)) // выводим строку и ее длину

// scanner := bufio.NewScanner(os.Stdin)
// fmt.Println("Enter a string:")

// if !scanner.Scan() {
// 	fmt.Println("Error reading input:", scanner.Err())
// 	return
// }

// input := scanner.Text()                                                 // сохраняем введенную строку целиком в переменную name
// fmt.Println("You entered:", input, "length of the string:", len(input)) // выводим строку и ее длину

// var num int
// fmt.Scan(&num)
// fmt.Println(num)
// var name string
// fmt.Scan(&name)
// fmt.Println("Привет,", name)

// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {

// 	scanner := bufio.NewScanner(os.Stdin)

// 	_ = scanner.Scan()
// 	stringSeparator := scanner.Text()

// 	_ = scanner.Scan()
// 	input1 := scanner.Text()

// 	_ = scanner.Scan()
// 	input2 := scanner.Text()

// 	_ = scanner.Scan()
// 	input3 := scanner.Text()

// 	fmt.Println(input1 + stringSeparator + input2 + stringSeparator + input3)
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	var num int = 52
// 	str := strconv.Itoa(num)
// 	fmt.Println(str + str) // вывод 55
// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	_ = scanner.Scan()
// 	a, _ := strconv.Atoi(scanner.Text())
// 	b := a * a
// 	fmt.Println(b)
// }

package main

import (
	"fmt"
)

func main() {
	var n int64
	fmt.Scan(&n)
	n2 := n * n
	n4 := n2 * n2
	n6 := n4 * n2
	fmt.Println(n6)
}
