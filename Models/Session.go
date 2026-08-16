package Models

type Session struct {
	ID      string
	UserID  string
	Token   string
	Expires int64
}
