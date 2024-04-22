package services

import (
	"github.com/Mirzoev-Parviz/movie-recommendation/models"
	"github.com/Mirzoev-Parviz/movie-recommendation/utils"
	"sort"
	"sync"
)

type RCM interface {
	Recommend(interactions []models.Interactions, items []models.Item,
		users []models.User, userID int) []int
	RecommendRandomMovies(userGenres map[string]bool, items []models.Item) []int
}

type RCM_Service struct {
}

func NewRCM_Service() *RCM_Service {
	return &RCM_Service{}
}

func (r *RCM_Service) Recommend(interactions []models.Interactions, items []models.Item,
	users []models.User, userID int) []int {
	userInteractions := make(map[int]float64)
	watchedMovies := make(map[int]bool)

	for _, interaction := range interactions {
		if interaction.UserID == userID {
			userInteractions[interaction.ItemID] = interaction.WatchedPCT
			watchedMovies[interaction.ItemID] = true
		}
	}

	var userHasKids bool
	for _, user := range users {
		if user.ID == userID {
			userHasKids = user.HasKids
			break
		}
	}

	userGenres := make(map[string]bool)
	for itemID, _ := range userInteractions {
		for _, item := range items {
			if item.ID == itemID {
				for _, genre := range item.Genres {
					userGenres[genre] = true
				}
			}
		}
	}

	var recommendationItems []int
	for _, item := range items {
		if _, watched := userInteractions[item.ID]; !watched {
			if (!userHasKids && !item.ForKids) || (userHasKids && item.ForKids) {
				if utils.HasCommonGenres(item.Genres, userGenres) {
					recommendationItems = append(recommendationItems, item.ID)
				}
			}
		}
	}

	if len(recommendationItems) == 0 {
		recommendationItems = r.RecommendRandomMovies(userGenres, items)
	}

	similarityChan := make(chan struct {
		index int
		value float64
	})

	var wg sync.WaitGroup
	for i, itemID := range recommendationItems {
		wg.Add(1)
		go func(index int, itemID int) {
			defer wg.Done()
			similarity := utils.CalculateSimilarity(itemID, userInteractions, items)
			similarityChan <- struct {
				index int
				value float64
			}{index, similarity}
		}(i, itemID)
	}

	go func() {
		wg.Wait()
		close(similarityChan)
	}()

	similarities := make([]struct {
		index int
		value float64
	}, len(recommendationItems))

	for result := range similarityChan {
		similarities[result.index] = result
	}

	sort.SliceStable(recommendationItems, func(i, j int) bool {
		return similarities[i].value > similarities[j].value
	})
	if len(recommendationItems) > 10 {
		recommendationItems = recommendationItems[:10]
	}
	return recommendationItems
}

func (r *RCM_Service) RecommendRandomMovies(userGenres map[string]bool, items []models.Item) []int {
	var recommendedItems []int
	for _, item := range items {
		for _, genre := range item.Genres {
			if userGenres[genre] {
				recommendedItems = append(recommendedItems, item.ID)
				break
			}
		}
	}
	if len(recommendedItems) == 0 || len(recommendedItems) < 10 {
		popularItems := utils.SortByPopularity(items)
		for _, item := range popularItems {
			recommendedItems = append(recommendedItems, item.ID)
			if len(recommendedItems) >= 10 {
				break
			}
		}
	}
	return recommendedItems
}
