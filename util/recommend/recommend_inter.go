package recommend

type RecommenderSystem interface {
	InitSystem() error
	Recommend(userID int) []int
}
