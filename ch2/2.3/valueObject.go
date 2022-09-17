package main

func before() {
	// 桁数が異なっていたり、想定しない文字が入っていた場合が考慮されていない
	skuCD, _ := r.URL.Query()["sku_code"]
	itemCD, sizeCD, colorCD := skuCD[0:5], skuCD[5:7], skuCD[7:9]
}

func after() {
	param, _ := r.URL.Query()["sku_code"]
	skuCD := SKUCode(param)

	if skuCD.Invalid() {
		// 異常系のハンドリング
	}
	itemCD, sizeCD, colorCD := skuCD
}

type SKUCode string

func (c SKUCode) Invalid() bool {
	// 桁数や利用可能文字のチェックを行う
}

func (c SKUCode) ItemCD() string {
	return c[0:5]
}

func (c SKUCode) SizeCD() string {
	return c[5:7]
}

func (c SKUCode) ColorCD() string {
	return c[7:9]
}
