package rest

// User ...
type User struct {
	ID 		 int     `json:"-"`
	Name 	 string  `json:"name"`
	Usename  string  `json:"username"`
	Password string  `json:"password"`
}