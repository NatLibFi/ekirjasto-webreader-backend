package cache

import (
	"time"

	"github.com/readium/go-toolkit/pkg/pub"
)

// CachedPublication implements Evictable
type CachedPublication struct {
	*pub.Publication
	Remote   bool
	CachedAt time.Time
}

func EncapsulatePublication(pub *pub.Publication, remote bool) *CachedPublication {
	return &CachedPublication{pub, remote, time.Now()}
}

func (cp *CachedPublication) OnEvict() {
	// Cleanup
	if cp.Publication != nil {
		cp.Publication.Close()
	}
}
