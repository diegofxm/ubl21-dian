package security

import (
	"time"
)

// TimestampData contiene los datos para generar un timestamp WS-Security
type TimestampData struct {
	ID      string
	Created string
	Expires string
}

// GenerateTimestamp crea un timestamp con TTL de 60000 segundos (como PHP)
func GenerateTimestamp(id string) TimestampData {
	now := time.Now().UTC()
	return TimestampData{
		ID:      id,
		Created: now.Format("2006-01-02T15:04:05Z"),
		Expires: now.Add(60000 * time.Second).Format("2006-01-02T15:04:05Z"),
	}
}
