package handler

import
	(
		"math"
		"github.com/SyncOrSink/HectoClash/backend/models"
	)

func ExpectedScore(ratingA, ratingB int64) float64 {
    return 1.0 / (1.0 + math.Pow(10, float64(ratingB-ratingA)/400))
}

func RecordMatchElo(winner, loser *models.User) {
    expectedWin := ExpectedScore(winner.Rating, loser.Rating)
    expectedLose := ExpectedScore(loser.Rating, winner.Rating)

    winner.Rating += int64(float64(1) * (1.0 - expectedWin))
    loser.Rating += int64(float64(1) * (0.0 - expectedLose))
}
