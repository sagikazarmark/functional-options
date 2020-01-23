package advanced

import "time"

func ExampleNewServer() {
	NewServer(":8080", Timeout(10*time.Second), MaxConnections(100))
}
