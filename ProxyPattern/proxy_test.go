package ProxyPattern

import "testing"

func Test(t *testing.T) {
	t.Run("ProxyPattern", Proxy)
}

func Proxy(t *testing.T) {
	car := NewCarProxy(&Driver{12})
	car.Drive()
}
