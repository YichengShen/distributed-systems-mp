package messages

// AckType = 0
// EmailType = 1
const (
	AckType = iota
	EmailType
)

type Msg struct {
	Type int
	Content Email
}
