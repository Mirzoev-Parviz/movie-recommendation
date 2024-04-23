package utils

import (
	"encoding/csv"
	"github.com/Mirzoev-Parviz/movie-recommendation/internal/dto"
	"github.com/Mirzoev-Parviz/movie-recommendation/models"
	"math/rand"
	"os"
	"strings"
	"time"
)

func HasCommonGenres(itemGenres []string, userGenres map[string]bool) bool {
	for _, genre := range itemGenres {
		if userGenres[genre] {
			return true
		}
	}

	return false
}

func SortByPopularity(items []models.Item) []models.Item {
	rand.Seed(time.Now().UnixNano())
	sortedItems := make([]models.Item, len(items))
	copy(sortedItems, items)
	rand.Shuffle(len(sortedItems), func(i, j int) {
		sortedItems[i], sortedItems[j] = sortedItems[j], sortedItems[i]
	})
	return sortedItems
}

func CalculateSimilarity(itemID int, userInteractions map[int]float64, items []models.Item) float64 {
	var targetItem models.Item

	for _, item := range items {
		if item.ID == itemID {
			targetItem = item
			break
		}
	}

	targetActors := make(map[string]bool)
	for _, actor := range targetItem.Actors {
		targetActors[actor] = true
	}
	targetDirectors := make(map[string]bool)
	for _, director := range targetItem.Directors {
		targetDirectors[director] = true
	}
	var totalSimilarity float64
	var numInteractions float64
	for itemID, pct := range userInteractions {
		var userItem models.Item
		for _, item := range items {
			if item.ID == itemID {
				userItem = item
				break
			}
		}
		actorSimilarity := calculateSetSimilarity(targetActors, userItem.Actors)
		directorSimilarity := calculateSetSimilarity(targetDirectors, userItem.Directors)
		similarity := (actorSimilarity + directorSimilarity) / 2 // Среднее значение
		totalSimilarity += similarity * pct
		numInteractions += pct
	}
	if numInteractions > 0 {
		return totalSimilarity / numInteractions
	}
	return 0
}

func calculateSetSimilarity(set1 map[string]bool, set2 []string) float64 {
	var commonItems int
	for _, item := range set2 {
		if set1[item] {
			commonItems++
		}
	}
	similarity := float64(commonItems) / float64(len(set1)+len(set2)-commonItems)
	return similarity
}

func ReadCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func SplitData(data string) []string {
	return strings.Split(data, ",")
}

func ConvertRCM(ides []int) []models.Item {
	var items []models.Item

	for _, id := range ides {
		for _, item := range dto.Items {
			if item.ID == id {
				items = append(items, item)
			}
		}
	}

	return items
}
