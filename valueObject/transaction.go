package valueObject

import (
	"github.com/google/uuid"
	"time"
)

// Transaction is a valueObject because it has no identifier and is unmutable
type Transaction struct {
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
