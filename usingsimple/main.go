package main

import "unumeration/simple"
import "fmt"
import "runtime"
import "testing"

func enumerate(length int) {
	charSet := [...]string{
		"a","","b","c","d","e","f","g","h","i","j",
		"k","l","m","n","o","p","q","r","s","t",
		"u","v","w","x","y","z","0","1","2","3",
		"4","5","6","7","8","9","!","@","#","$",
		"%","^","&","*","(",")","-","}","{","_",
		"+","/","\\","<",">",".","`","'","\"","~",
		"A","B","C","D","E","F","G","H","I","J","K",
		"L","M","N","O","P","R","S","T","U","V","W",
		"X","Y","Z",
		}
	
	maxLength := length
	const numberThreads = 4
	
	var laborArray [numberThreads]simple.Combinator
	var deliveryArray [numberThreads]chan string
	runtime.GOMAXPROCS(numberThreads)
	
	for i := 0; i < len(laborArray); i++ {
		deliveryArray[i] = make(chan string, 10)
		laborArray[i] = simple.NewCombinator(charSet[:], maxLength)
		blockSize := laborArray[i].MaxVal()/int64(len(laborArray))
		laborArray[i].Skip(blockSize*int64(i))
	}
	complete := make(chan int) 
	for i := 0; i < len(laborArray); i++ {
		blockSize := laborArray[i].MaxVal()/int64(len(laborArray))
		if i == len(laborArray)-1 {
			blockSize += (laborArray[i].MaxVal()%int64(len(laborArray)))
		}
		go func(index int, blockSize int64, mailbox chan string) {
			for j := 0; int64(j) < blockSize; j++ {
				fmt.Printf("combination: \t%s\n", laborArray[index])
				laborArray[index].Next()
			}
			complete <- 1
		}(i, blockSize, deliveryArray[i])
	}
	
		
	for i := 0; i < len(deliveryArray); i++ { <-complete }
	//fmt.Printf("finished")
}

func benchmarkEnumerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		enumerate(3)
	}
}

func main() {
	//br := testing.Benchmark(benchmarkEnumerate)
	//fmt.Println(br)
	enumerate(5)
}

