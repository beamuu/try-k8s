package controllers

type Query struct {
	Key string `form:"key"`
	Value string `form:"value"`
}