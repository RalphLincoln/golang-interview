package models

type Movie struct {
	ID                 string  `json:"id" gorm:"primaryKey"`
	Title              string  `json:"title"`
	USGross            int     `json:"us_gross"`
	WorldwideGross     int     `json:"worldwide_gross"`
	ReleaseDate        string  `json:"release_date"`
	IMDBVotes          int     `json:"imdb_votes"`
	Director           string  `json:"director"`
	USDVDSales         string  `json:"usdvd_sales"`
	ProductionBudget   int     `json:"production_budget"`
	MPAARating         string  `json:"mpaa_rating"`
	RunningTimeMin     string  `json:"running_time_min"`
	Distributor        string  `json:"distributor"`
	Source             string  `json:"source"`
	MajorGenre         string  `json:"major_genre"`
	CreativeType       string  `json:"creative_type"`
	RottenTomatoRating int     `json:"rotten_tomato_rating"`
	IMDBRating         float64 `json:"imdb_rating"`
}
