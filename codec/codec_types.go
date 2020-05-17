package codec

const (
	SeparatorBetweenContentAndID = ":"

	ID_ASCI = "asci" // Ascii (Raw) codec - limitations: No ';' in any key's value
)

type Codec interface {
	Encode(keyValues map[string]string) (opaque string, err error)

	Decode(opaque string) (keyValues map[string]string, err error)
}
