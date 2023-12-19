package response

type Pokedex struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Total     string `json:"total"`
	HP        string `json:"hp"`
	Attack    string `json:"attack"`
	Defense   string `json:"defense"`
	Speed     string `json:"speed"`
	SpAttack  string `json:"spattack"`
	SpDefense string `json:"spdefense"`
}

type PokedexList struct {
	List []Pokedex `json:"list"`
}
