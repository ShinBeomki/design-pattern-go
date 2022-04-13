package BuilderPattern

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("BuilderPattern", BuilderPatternTest)
}

func BuilderPatternTest(t *testing.T) {
	KBbank := New()
	account := KBbank.Owner("Beomki").AccountNumber("1234-123-22").Build()
	fmt.Println(account.Deposit(10000))

	BSbank := New()
	account2 := BSbank.Owner("Patty").Build()
	fmt.Println(account2.ShowInfo())

}
