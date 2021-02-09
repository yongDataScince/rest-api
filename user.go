package rest

// User ...
type User struct {
	ID 		 int     `json:"-"`
	Name 	 string  `json:"name" binding:"required"`
	Usename  string  `json:"username" binding:"required"`
	Password string  `json:"password" binding:"required"`
}