package main

import "time"

// スライスに対して定義型を作成し、メソッドを定義
type Consumers []Consumer

func (c Consumers) expires(end time.Time) Consumers {
	// endの日時以降に契約が失効するユーザに絞り込む
}

func (c Consumers) sortByExpiredAt() Consumers {
	// 契約期限日で昇順ソートする
}

func (c Consumers) ActiveConsumer() Consumers {
	resp := make([]Consumer, 0, len(c))
	for _, v := range c {
		if v.ActiveFlg {
			resp = append(resp, v)
		}
	}
	return resp
}

func main() {

}
