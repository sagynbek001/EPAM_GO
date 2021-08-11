package saver

//creating a map of tokens (database)
var Tokens map[int64]Token = make(map[int64]Token)

type Token struct {
	Token     string
	CreatedAt string
	ExpiredAt string
}
