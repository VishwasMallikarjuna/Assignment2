package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	s1 := "The fact of the matter is that although you may have numerous valid facts or descriptions related to your paragraph’s core idea, you may lose a reader’s attention if your paragraphs are too long. What’s more, if all of your paragraphs are long, you may lose opportunities to draw your reader in. Journalists, for example, know that their readers respond better to short paragraphs. News readers generally lose interest with long descriptions and even one-sentence paragraphs are considered both acceptable and impactful.The model regarding paragraph length that your teacher undoubtedly taught you involves a topic sentence, a number of facts that support that core idea, and a concluding sentence. The proviso about the number of sentences between the topic sentence and the conclusion was not given to you because it was the magic formula for creating paragraphs of the perfect length; rather, your educator was attempting to give you a good reason to do adequate research on your topic. Academic writing yields the best examples of the topic-support-conclusion paragraph structure.Recent research has provided a wealth of insight about how dogs came to be domesticated by humans and the roles they played in Native American culture. DNA studies on archaeological finds suggest that dogs may have been domesticated by humans as long as 40,000 years ago. When the first humans came to North America from Eurasia, at least 12,000 years ago, domesticated dogs came with them. They appear to have been highly prized by early North American hunter-gatherers and were their only animal companions for centuries, since there were no horses on the continent until the 16th century. You can see from this example how a topic is introduced, supported, and then brought to its natural conclusion. Yet, not all writing is academic, and once you have learned the concept behind good paragraph construction—which is really the art of focused writing in disguise—you should know that there are times when paragraph “rules” can, and should, be broken.READ: Splitting Paragraphs for Easier Reading How to write paragraphs people want to read The fact of the matter is that although you may have numerous valid facts or descriptions related to your paragraph’s core idea, you may lose a reader’s attention if your paragraphs are too long. What’s more, if all of your paragraphs are long, you may lose opportunities to draw your reader in. Journalists, for example, know that their readers respond better to short paragraphs. News readers generally lose interest with long descriptions and even one-sentence paragraphs are considered both acceptable and impactful."
	callServce1(s1)
}

func callServce1(s string) {

	responseBody := bytes.NewBuffer([]byte(s))
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post("http://localhost:8000/text", "application/text", responseBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	fmt.Println(sb)

}
