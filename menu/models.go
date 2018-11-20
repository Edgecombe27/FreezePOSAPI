package menu

type MenuItem struct {
	ID int `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Price float32 `json:"price,omitempty"`
	Ingredients []int `json:"ingredients,omitempty"`
	Preparations []int `json:"preparations,omitempty"`
	Groups []int `json:"groups,omitempty"`
}

type Ingredient struct {
	ID int `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Price float32 `json:"price,omitempty"`
}

type Preparation struct {
	ID int `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

type Group struct {
	ID int `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Ingredients []int `json:"ingredients,omitempty"`
	Preparations []int `json:"preparations,omitempty"`
}

type Menu struct {
	Ingredients []Ingredient `json:"ingredients,omitempty"`
	MenuItems []MenuItem `json:"menu_items,omitempty"`
	Preparations []Preparation `json:"preparations,omitempty"`
	Groups []Group `json:"groups,omitempty"`
}



