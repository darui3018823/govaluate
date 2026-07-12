package govaluate

type tokenStream struct {
	tokens      []ExpressionToken
	index       int
	tokenLength int
}

func newTokenStream(tokens []ExpressionToken) *tokenStream {
	return &tokenStream{
		tokens:      tokens,
		tokenLength: len(tokens),
	}
}

func (stream *tokenStream) rewind() {
	stream.index--
}

func (stream *tokenStream) next() ExpressionToken {

	token := stream.tokens[stream.index]
	stream.index++
	return token
}

func (stream tokenStream) hasNext() bool {

	return stream.index < stream.tokenLength
}
