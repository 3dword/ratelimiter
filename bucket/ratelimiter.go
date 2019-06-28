package bucket

import (
	"time"
	"math"
)

func IsLimited (bucket *Bucket) bool {

	now := time.Now().UnixNano()

	if (now - bucket.lastDuration) >= int64(math.Pow(10,9)){
		bucket.SetLastDuration(now)
		bucket.Reset()
	}

	if bucket.capacity > 0 {
		bucket.Tick()
		return false
	}

	return true

}
