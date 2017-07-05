package coolq

import "testing"
import "fmt"

// SendPriviateMoreMsg
func Test_MoreUserMsg(t *testing.T) {
	cool := NewCoolQ("http://123.206.134.36:5700", "")
	QQnumber := []int{
		526133625,
		465534829,
	}
	haveError, errs := cool.SendPriviateMoreMsg(QQnumber, "消息Test", true)
	if haveError == true {
		// t.Fatalf(err)
		for _, err := range errs {
			if err != nil {
				fmt.Println(err)
			}
		}
		t.Fatalf("Error")
	}
}
