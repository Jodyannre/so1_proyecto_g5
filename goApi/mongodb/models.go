package mongodb

type LogMongo struct {
	//ID    string `json:"id" bson:"-"`
	Team1 string `json:"team1" bson:"team1,omitempty"`
	Team2 string `json:"team2" bson:"team2,omitempty"`
	Score string `json:"score" bson:"score,omitempty"`
	Phase int    `json:"phase" bson:"phase,omitempty"`
}

type LogsMongo []LogMongo

type TotalMongo struct {
	Total int64
}
