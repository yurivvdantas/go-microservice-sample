package test

import (
	"go-microservice-sample/internal/crypto-votes-service/repository"
	"testing"
)

//Test repository for get a crypto by id
func TestFindById(t *testing.T) {
	id := int64(1)
	want := "KLV"

	cry, err := repository.FindCryptoById(id)

	var code string
	if cry != nil {
		code = cry.Code
	}
	if code != want || err != nil {
		t.Fatalf(`FindById(%d) return = [value: %q - err: %v]. Expectation was = [value: %#q - err: nil]`, id, code, err, want)
	}
}

//Test repository for get all crypto
func TestFindAll(t *testing.T) {
	cryptos, err := repository.FindAllCrypto()
	if len(cryptos) < 1 || err != nil {
		t.Fatalf(`FindAll return = [value: %q - err: %v]. Expectation was err != nil and size of return was greater than 0`, cryptos, err)
	}
}
