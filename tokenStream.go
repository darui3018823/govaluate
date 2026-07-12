package govaluate

type tokenStream struct {
	tokens      []ExpressionToken
	index       int
	tokenLength int
}

func newTokenStream(tokens []ExpressionToken) *tokenStream {

	var ret *tokenStream

	ret = new(tokenStream)
	ret.tokens = tokens
	ret.tokenLength = len(tokens)
	return ret
}

func (stream *tokenStream) rewind() {
	stream.index -= 1
}

func (stream *tokenStream) next() ExpressionToken {

	var token ExpressionToken

	token = stream.tokens[stream.index]

	stream.index += 1
	return token
}

func (stream tokenStream) hasNext() bool {

	return stream.index < stream.tokenLength
}
