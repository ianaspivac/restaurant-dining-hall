package components

import "fmt"

var Ratings = [6]float32{0, 0, 0, 0, 0, 0}

func GetRating(maxWait float32, cookingTime int64) {
	fCookingTime := float32(cookingTime)
	fMaxWait := float32(maxWait)
	if fCookingTime < fMaxWait {
		Ratings[5]++
	} else if fCookingTime < fMaxWait*1.1 {
		Ratings[4]++
	} else if fCookingTime < fMaxWait*1.2 {
		Ratings[3]++
	} else if fCookingTime < fMaxWait*1.3 {
		Ratings[2]++
	} else if fCookingTime < fMaxWait*1.4 {
		Ratings[1]++
	} else {
		Ratings[0]++
	}
	rating := (Ratings[1]*1 + Ratings[2]*2 + Ratings[3]*3 + Ratings[4]*4 + Ratings[5]*5) / (Ratings[0] + Ratings[1] + Ratings[2] + Ratings[3] + Ratings[4] + Ratings[5])
	fmt.Printf("Current rating: %f\n", rating)
}
