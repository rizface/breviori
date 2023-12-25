package urlshortener

type Shortener struct{}

func New() *Shortener {
	return &Shortener{}
}

func (s *Shortener) Short(url string) (string, error) {
	return "shortened", nil
}
