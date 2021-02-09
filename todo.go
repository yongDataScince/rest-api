package rest

// Todo ... 
type Todo struct {
	ID 			int    `json:"id"`
	Title 		string `json:"title"`
	Description string `json:"description"`
}

// UserList ...
type UserList struct {
	ID 		int
	UserID 	int
	TodoID  int
}

// TodoItem ...
type TodoItem struct {
	ID 			int   	`json:"id"`
	Title 		string	`json"title"`
	Description string	`json:"description"`
	Done 		bool 	`json:"done"`
}

// ListsItem ... 
type ListsItem struct {
	ID 		int
	ListID 	int
	ItemID 	int
}