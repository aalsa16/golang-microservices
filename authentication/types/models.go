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
	Token     string
	CreatedAt string
}

type SaveResponse struct {
	ID        int64
	Email     string
	Uuid      string
	Token     string
	CreatedAt int64
}
