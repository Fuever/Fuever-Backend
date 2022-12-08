package authentication

import (
	"fmt"
	"testing"
	"time"
)

func TestAuthentication(t *testing.T) {
	year, month, day := time.Now().Date()
	fmt.Println(year, int(month), day)
}
