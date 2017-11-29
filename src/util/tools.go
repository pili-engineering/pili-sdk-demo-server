package util

import (
	"cli"
	"encoding/base64"
	"errors"
	"github.com/qiniu/log"
	"strings"
)

func Authority(token string) (string, error) {
	if token == "" {
		return "", errors.New("authorization is null")
	}
	log.Infof("authorization:%s\n", token)
	decodeBytes, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", err
	}
	res := strings.Split(string(decodeBytes), ":")
	if len(res) != 2 {
		return "", errors.New("the authorization error, " + token)
	}
	existErr := cli.QueryUser(res[0], res[1])
	if existErr != nil {
		return "", existErr
	}
	return res[0], nil
}

func AuthorityOfAdmin(token string) (string, error) {
	if token == "" {
		return "", errors.New("authorization is null")
	}
	log.Infof("admin's authorization : %s\n", token)
	decodeBytes, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", err
	}
	res := strings.Split(string(decodeBytes), ":")
	if len(res) != 2 {
		return "", errors.New("the authorization error, " + token)
	}
	existErr := cli.QuerySaller(res[0], res[1])
	if existErr != nil {
		return "", existErr
	}
	return res[0], nil
}
