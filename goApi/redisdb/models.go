package redisdb

type Match struct {
	Team1 string `json:"team1"`
	Team2 string `json:"team2"`
	Score string `json:"score"`
	Phase int    `json:"phase"`
}

type Key struct {
	Key string `json:"key"`
}

type Fase struct {
	Fase string `json:"fase"`
}

type Fases struct {
	Fase1 []string
	Fase2 []string
	Fase3 []string
	Fase4 []string
}

type Conteo struct {
	Partido string
	Total   float64
}

type Grafica struct {
	Data   []float64
	Labels []string
}

type Resultados []Conteo
