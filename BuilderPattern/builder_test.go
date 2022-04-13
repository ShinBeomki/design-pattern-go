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
	//patty의 계좌번호를 넣을수도 안넣을수도 있다. 또한 넣을데이터를 가공해서 넣을수도있다.
	//	ex) 오늘의 날짜와 시간을 조합해 계좌번호를 생성후 build

}
