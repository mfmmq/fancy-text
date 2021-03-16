package emoji

import (
	"fmt"
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().Unix())

func emojis() []string {
	return []string{"🎀", "💗", "🍭", "💎", "🍪"}
}

func getRandomEmoji() string {
	r := rand.New(randSource) // initialize local pseudorandom generator
	randomIndex := r.Intn(len(emojis()))
	item := emojis()[randomIndex]
	time.Sleep(time.Duration(randomIndex) * time.Second)
	fmt.Printf("Emoji service sent the '%v' emoji. Took %v seconds \n", item, randomIndex)
	return item
}


func GetEmojiFromSuperSlowService() string {
	return getRandomEmoji()
}

// func WriteEmojiToChannel(c chan string) {
// 	item := getRandomEmoji()
// 	c <- item	
// }

// func WriteEmojiToChannelAndWait(wg *sync.WaitGroup, c chan string) {
// 	defer wg.Done()
// 	item := getRandomEmoji()
// 	c <- item
// }