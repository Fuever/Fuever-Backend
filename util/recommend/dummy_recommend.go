package recommend

import (
	"math/rand"
	"time"
)

type RandomRecommendSystem struct {
}

func (r RandomRecommendSystem) InitSystem() error {
	rand.Seed(time.Now().Unix())
	return nil
}

func (r RandomRecommendSystem) Recommend(userID int) []int {
	return []int{
		rand.Int(),
		rand.Int(),
		rand.Int(),
		rand.Int(),
		rand.Int(),
	}
}
