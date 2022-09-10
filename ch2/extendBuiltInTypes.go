// 2.2.1 メソッドを追加して組み込み型を拡張する
package main

import (
	"fmt"
)

type HTTPStatus int

const (
	StatusOK              HTTPStatus = 200
	StatusUnauthorized    HTTPStatus = 401
	StatusPaymentRequired HTTPStatus = 402
	StatusForbidden       HTTPStatus = 403
)

type NationalRoute int

const (
	NagasakiKaido   NationalRoute = 200
	AizuNumataKaido NationalRoute = 401
	HokurikuDo      NationalRoute = 402
	KurinokiBypass  NationalRoute = 403
)

func (n NationalRoute) String() string {
	switch n {
	case NagasakiKaido:
		return "長崎街道"
	case AizuNumataKaido:
		return "会津沼田街道"
	case HokurikuDo:
		return "北陸道"
	case KurinokiBypass:
		return "栗ノ木バイパス"
	default:
		return fmt.Sprintf("国道%d号線", n)
	}
}

func main() {
	var status HTTPStatus = 401
	fmt.Println(status.String()) // Unauthorized

	var route NationalRoute = 201
	fmt.Println(route.String())
}

func (s HTTPStatus) String() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusUnauthorized:
		return "Unauthorized"
	case StatusPaymentRequired:
		return "Payment Required"
	case StatusForbidden:
		return "Forbidden"
	default:
		return fmt.Sprintf("HTTPStatus(%d", s)
	}
}
