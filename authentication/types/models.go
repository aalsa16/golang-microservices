package types

type SignUpRequest struct {
	Email    string
	Password string
}

type User struct {
	ID        int64
	Email     string
	Password  string
	Uuid      string
	CreatedAt string
}

type SaveResponse struct {
	ID        int64
	Email     string
	Uuid      string
	CreatedAt int64
}
