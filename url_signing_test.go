package urlsign

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_timedSecret(t *testing.T) {
	Convey("timedSecret()", t, func() {
		secret := []byte("king-under-the-mountain")

		Convey("generates a different secret for each baseTime", func() {
			secret1 := timedSecret(secret, time.Now().UTC().UnixNano())
			secret2 := timedSecret(secret, time.Now().UTC().UnixNano()+1)

			So(secret1, ShouldNotBeEmpty)
			So(secret2, ShouldNotBeEmpty)
			So(secret1, ShouldNotEqual, secret2)
		})

		Convey("generates a different secret for each key", func() {
			baseTime := time.Now().UTC().UnixNano()
			secret1 := timedSecret(secret, baseTime)
			secret2 := timedSecret([]byte("abc123"), baseTime)

			So(secret1, ShouldNotBeEmpty)
			So(secret2, ShouldNotBeEmpty)
			So(secret1, ShouldNotEqual, secret2)
		})
	})
}

func Test_GenerateToken(t *testing.T) {
	Convey("GenerateToken()", t, func() {
		Convey("generates some expected values", func() {
			secret := "asdfasdf"
			bucketSize := 8 * time.Hour
			baseTime := time.Date(2017, 9, 26, 13, 47, 0, 0, time.UTC)

			uris := map[string]string{
				"/hobbiton1?page=2&width=1024": "5cba73b7cfc04a86dbbeeef094f219f1883a08d8",
				"/hobbiton2?page=2&width=1024": "dc6a803f73b49d7421f432039ea42e1b5e8a276e",
				"/hobbiton3?page=2&width=1024": "1651f219ff5998d0791dabe3bbf78ab973281da0",
				"/hobbiton4?page=2&width=1024": "5974369c4c8acd0bf2ed4300e2738494d0e1e486",
				"/hobbiton5?page=2&width=1024": "ed224dfe0c8366abc0b864014c058e7224fcb754",
			}

			for uri, token := range uris {
				So(GenerateToken(secret, bucketSize, baseTime, uri), ShouldEqual, token)
			}
		})
	})
}

func Test_IsValidSignature(t *testing.T) {
	Convey("IsValidSignature()", t, func() {
		secret := "asdfasdf"
		bucketSize := 8 * time.Hour
		baseTime := time.Date(2017, 9, 26, 13, 47, 0, 0, time.UTC)

		urls := []string{
			"http://example.com/hobbiton1?page=2&width=1024&token=5cba73b7cfc04a86dbbeeef094f219f1883a08d8",
			"http://example.com/hobbiton2?page=2&width=1024&token=dc6a803f73b49d7421f432039ea42e1b5e8a276e",
			"http://example.com/hobbiton3?page=2&width=1024&token=1651f219ff5998d0791dabe3bbf78ab973281da0",
			"http://example.com/hobbiton4?page=2&width=1024&token=5974369c4c8acd0bf2ed4300e2738494d0e1e486",
			"http://example.com/hobbiton5?page=2&width=1024&token=ed224dfe0c8366abc0b864014c058e7224fcb754",
		}

		Convey("accepts valid signatures from this bucket", func() {
			for _, url := range urls {
				So(IsValidSignature(secret, bucketSize, baseTime, url), ShouldBeTrue)
			}
		})

		Convey("accepts signatures from the previous time bucket", func() {
			baseTime = baseTime.Add(0-bucketSize)
			for _, url := range urls {
				So(IsValidSignature(secret, bucketSize, baseTime, url), ShouldBeTrue)
			}
		})

		Convey("accepts signatures from the next time bucket", func() {
			baseTime = baseTime.Add(bucketSize)
			for _, url := range urls {
				So(IsValidSignature(secret, bucketSize, baseTime, url), ShouldBeTrue)
			}
		})

		Convey("does not accept signatures from outside the bucket", func() {
			baseTime = baseTime.Add(2*bucketSize)
			for _, url := range urls {
				So(IsValidSignature(secret, bucketSize, baseTime, url), ShouldBeFalse)
			}
		})
	})
}
