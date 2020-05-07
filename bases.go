package sugusama

type Bases struct {
	Web  string
	Live string
}

var DefaultBases = &Bases{
	Web:  "https://instagram.com",
	Live: "https://i.instagram.com",
}
