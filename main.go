package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	target := numRand(1, 100)
	fmt.Println("1 ile 100 arasindaki sayiyi bulmaya calisin")
	reader := bufio.NewReader(os.Stdin)
	for attempts := 0; attempts < 10; attempts++ {

		fmt.Println(10-attempts, "hakkiniz kaldi")
		fmt.Print("tahmininizi yapiniz : ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
		}
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}
		if num > target {
			fmt.Println("tahminizi kucultun ")
		} else if num < target {
			fmt.Println("tahminizi buyutun")
		} else {
			fmt.Println("tahmin dogru", attempts, ".deneyisinizde buldunuz")
			break
		}

	}

}
func numRand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min

}
