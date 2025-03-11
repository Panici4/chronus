package chronus_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/panici4/chronus"
)

func TestOffset(t *testing.T) {
	c := chronus.NewChronus("aa", time.Hour*24)
	now := c.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))
}
