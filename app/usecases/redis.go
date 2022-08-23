package usecases

import (
	"errors"

	"github.com/nutchanonc/try-k8s/app/repositories"
)

type RedisUsecase struct {
	repository repositories.RedisRepository
}

func NewRedisUsecase(repository repositories.RedisRepository) RedisUsecase {
	return RedisUsecase{
		repository: repository,
	}
}

func (u RedisUsecase) Get(key string) (string, error) {
	if (key == "") {
		return "", errors.New("Key is an empty string.")
	}
	return u.repository.Get(key)
}

func (u RedisUsecase) Set(key string, value string) (error) {
	if (key == "") {
		return errors.New("Key is an empty string.")
	}
	return u.repository.Set(key, value, 0)
}