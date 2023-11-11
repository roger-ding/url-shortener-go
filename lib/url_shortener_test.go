package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortenedUrlGeneration(t *testing.T) {
  sourceUrl1:= "https://en.wikipedia.org/wiki/Kobe_Bryant"
  shortenedUrl1 := GenerateShortenedUrl(sourceUrl1)

  sourceUrl2:= "https://www.cnn.com/"
  shortenedUrl2 := GenerateShortenedUrl(sourceUrl2)

  assert.Equal(t, shortenedUrl1, "l6QuSjqijC")
  assert.Equal(t, shortenedUrl2, "H2tnGNyizt")
}
