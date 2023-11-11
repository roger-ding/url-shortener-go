package lib

import (
  "encoding/base64"
  "crypto/sha256"
)

func sha256Encrypt(input string) []byte {
  encryption := sha256.New()
  encryption.Write([]byte(input))
  return encryption.Sum(nil)
}

func base64Encode(data []byte) string {
  encodedString := base64.StdEncoding.EncodeToString([]byte(data))
  return string(encodedString)
}

func GenerateShortenedUrl(sourceUrl string) string {
  urlHashBytes := sha256Encrypt(sourceUrl)
  encodedUrl := base64Encode(urlHashBytes)
  return string(encodedUrl)[:10]
}
