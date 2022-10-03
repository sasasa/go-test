// docker-compose exec app go run main.go -n 100 -m hgoehoge -x

package main

import(
	"fmt"
	"os"
	"log"
	// "time"
	"flag"
	"local.packages/user"
	"github.com/tlorens/go-ibgetkey"
	. "local.packages/utility"
)

type MyError struct {
	errMessage string
}
// implements error interface
func (e *MyError) Error() string {
    return e.errMessage
}

func fn() (int, error) { // *MyErrorではなくerror
    return 42, &MyError{ errMessage: "errorが起こりました！" }
    // return 42, nil
}

func main() {
	u := user.NewUser(Dict{
		"Id": 1,
		"Name": "山田 太郎",
		"Email": "yamada@example.com",
	})
	encodedUser, err := u.Encode()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(encodedUser)
	}

	p, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)
	fmt.Println("Hello, World!")

	result, err := fn()
	if err != nil {
			fmt.Println(err.Error())
	} else {
		// エラーなし
		fmt.Println(result)
	}

	fmt.Println("Test start")
	fmt.Println("Create goroutine")
	quit := make(chan bool)
	counter := make(chan int)
	go appleGoroutine("Apple", 10, quit, counter)
	fmt.Println("Waiting for the goroutine to complete")
	
	var c int
	var isEnd bool
	loop:
	for {
			select {
				case isEnd = <-quit:
						break loop
				case c = <-counter:
					fmt.Println(c)
			}
	}
	fmt.Println(fmt.Sprintf("Test compleated  --- %t ---", isEnd))

	var ( 
		max int 
		msg string 
		x bool
	)
	
	flag.IntVar(&max, "n", 32, "処理数の最大値")
	flag.StringVar(&msg, "m", "---", "処理メッセージ")
	flag.BoolVar(&x, "x", false, "拡張オプション")
	/*コマンド ライン を パース*/
	flag.Parse()
	fmt.Println("処理数の最大値 = ", max)
	fmt.Println("処理メッセージ = ", msg)
	fmt.Println("拡張オプション = ", x)

}

func appleGoroutine(fruit string, a int, quit chan bool, counter chan int) {
	fmt.Printf("Start goroutine - %s %d\n", fruit, a)
	for i:=1; i<=a; i++ {
		counter <- i
	}
	targetkey := "."
	t := int(targetkey[0])
	loop:
	for {
		input := keyboard.ReadKey()
		fmt.Println(input)
		if input == t {
			quit <- true
			break loop
		}
	}
	// time.Sleep(1 * time.Second)
	fmt.Println("End goroutine")

}