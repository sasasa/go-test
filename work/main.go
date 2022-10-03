// docker-compose exec app go run main.go -n 100 -m hgoehoge -x

package main

import(
	"fmt"
	"os"
	"log"
	"time"
	"flag"
	"encoding/json"
	"local.packages/user"
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
	user := user.NewUser()
	/* JSON エン コード */
	bs, err := json.Marshal(user)
	if err != nil {
		log. Fatal( err)
	} else {
		fmt.Println(string(bs)) 
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
	go appleGoroutine("Apple", 10, quit)
	fmt.Println("Waiting for the goroutine to complete")
	<-quit
	fmt.Println("Test compleated")

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

func appleGoroutine(fruit string, a int, quit chan bool) {
	fmt.Printf("Start goroutine - %s %d\n", fruit, a)
	time.Sleep(3 * time.Second)
	fmt.Println("End goroutine")
	quit <- true
}