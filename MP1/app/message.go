package app

type Message struct {
	Source Process
	Destination Process
	Content string
}
