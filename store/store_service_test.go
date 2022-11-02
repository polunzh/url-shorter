package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreSerice = &StorageService{}

func init() {
	testStoreSerice = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreSerice.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortURL := "Jsz4k57oAX"

	SaveUrlMapping(shortURL, initialLink, userUUId)

	retrievedUrl := RetriveInitialUrl(shortURL)

	assert.Equal(t, initialLink, retrievedUrl)
}
