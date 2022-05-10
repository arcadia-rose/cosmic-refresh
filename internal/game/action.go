package game

type Action struct {
	Do Event  `json:"do"`
	It string `json:"it"`
	To Id     `json:"to"`
	Is Entity `json:"is"`
}
