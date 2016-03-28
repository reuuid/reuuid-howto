package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var ErrorCantReadBody = errors.New("Couldn't read response body")
var ErrorUUIDCount = errors.New("Tried to get < 1 UUID")

func GetOne() (string, error) {
	res, err := http.Get("http://reuuid.org/get/")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	uuid, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", ErrorCantReadBody
	}

	return string(uuid), nil
}

func GetSome(n int) ([]string, error) {
	if n < 1 {
		return nil, ErrorUUIDCount
	}

	res, err := http.Get("http://reuuid.org/get/" + strconv.Itoa(n))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	rawuuid, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	uuidbytes := bytes.Split(rawuuid, []byte{'\n'})
	uuids := make([]string, len(uuidbytes))
	for i, b := range uuidbytes {
		uuids[i] = string(b)
	}

	return uuids, nil
}

func Donate(uuids []string) error {
	rawuuid := bytes.NewBufferString(strings.Join(uuids, "\n"))
	_, err := http.Post("http://reuuid.org/give/", "text/plain", rawuuid)
	return err
}

func main() {
	uuid, err := GetOne()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println(uuid)
	}

	uuids, err := GetSome(5)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println(strings.Join(uuids, "\n"))
	}

	donations := []string{
		"3b54969c-d9fa-4ac9-aa38-4c69590ebaa5",
		"1237168c-35c4-437f-94fe-f48fe972eafa",
		"ac1c7d22-4903-4231-ae9a-c042c3a6211d",
		"b63f76ac-3d7c-43fd-b966-38ce938a126e",
		"03e65fe3-21d2-4ef4-bbf1-f14bb42f06e3",
		uuid, // Don't be greedy!
	}
	err = Donate(append(donations, uuids...))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
