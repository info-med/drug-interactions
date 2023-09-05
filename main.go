package main

import (
	"encoding/csv"
	"fmt"
	"github.com/google/uuid"
	"github.com/meilisearch/meilisearch-go"
	"os"
)

type interaction struct {
	Id       string
	DrugOne  string
	DrugTwo  string
	Severity string
}

// Nema nikakov data clean up, primer tuka ima mn lekarstva sto ne se
// dostapni vo Makedonija, terribly inefficient ama odime jako
// do golive, posle refactor
func main() {
	meilisearchClient := meilisearch.NewClient(meilisearch.ClientConfig{
		Host: "http://127.0.0.1:7700",
	})
	// Manually changed the files here each run, no time to waste.
	// But if you want it's easy to just ReadDir and for loop all CSVs there.
	file, err := os.Open("./csvs/V.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV")
		panic("Error")
	}

	var interactions []interaction
	for _, record := range records {
		interactions = append(interactions, interaction{
			Id:       uuid.NewString(),
			DrugOne:  record[1],
			DrugTwo:  record[3],
			Severity: record[4],
		})
	}

	drugInteractions := meilisearchClient.Index("drug-interactions")

	results, err := drugInteractions.AddDocuments(interactions)
	if err != nil {
		fmt.Println(err)
		fmt.Println(results)
		panic("Error")
	}
	fmt.Println("Done")
}
