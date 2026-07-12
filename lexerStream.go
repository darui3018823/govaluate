package govaluate

type lexerStream struct {
	source   []rune
	position int
	length   int
}

func newLexerStream(source string) *lexerStream {

	runes := []rune(source)
	return &lexerStream{
		source: runes,
		length: len(runes),
	}
}

func (stream *lexerStream) readCharacter() rune {

	character := stream.source[stream.position]
	stream.position++
	return character
}

func (stream *lexerStream) rewind(amount int) {
	stream.position -= amount
}

func (stream lexerStream) canRead() bool {
	return stream.position < stream.length
}
