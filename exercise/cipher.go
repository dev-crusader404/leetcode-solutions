package exercise

type Cipher interface {
	Encode(string) string
	Decode(string) string
}
