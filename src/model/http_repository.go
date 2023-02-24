package model

type HttpRepository interface {
	Get(url string, response interface{}) error
}
