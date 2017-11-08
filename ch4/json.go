// Json test go json package
package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"reflect"
	"strings"
)

const PATH = "/var/lib/docker/image/overlay/repositories.json"

type Repositories struct {
	//repo string `json:"10.10.7.161/base/nginx"`
	Repo *Repository `json:"Repositories"`
}

type Repository struct {
	Nginx *Tag `json:"10.10.7.161/base/nginx"`
}

type Tag struct {
	Sha string `json:"10.10.7.161/base/nginx:1.10"`
}

type RepositoriesSecond struct {
	NginxSha string `json:"10.10.7.161/base/nginx:1.10"`
}

/*
type RepositoriesFunc struct {
	Repo *RepositoryFunc `json:"Repositories"`
}
*/

func main() {
	fmt.Println("--- test json unmarshal ---")
	fd, err := os.Open(PATH)
	if err != nil {
		panic(err)
	}
	var nginx Repositories
	if err := json.NewDecoder(fd).Decode(&nginx); err != nil {
		fmt.Printf("err != nil\n")
		return
	}
	fmt.Printf("embeded:\n%v\n", nginx.Repo.Nginx.Sha)
	fd.Close()

	fdSecond, err := os.Open(PATH)
	if err != nil {
		panic(err)
	}
	var nginxSecond RepositoriesSecond
	if err := json.NewDecoder(fdSecond).Decode(&nginxSecond); err != nil {
		fmt.Printf("err != nil\n")
		return
	}
	fmt.Printf("direct:\n%v\n", nginxSecond.NginxSha)
	fdSecond.Close()

	fmt.Printf("\n--- test reflect ---\n")
	t := reflect.TypeOf(nginx.Repo)
	field := t.Elem().Field(0)
	tag := field.Tag.Get("json")
	fmt.Printf("initial: %s\n", tag)
	newTag := `json:"alpine"`
	//reflect.TypeOf(nginx.Repo).Elem().Field(0).Tag = reflect.StructTag(newTag)
	field.Tag = reflect.StructTag(newTag)
	tag = field.Tag.Get("json")
	fmt.Printf("changed: %s\n", tag)
	t2 := reflect.TypeOf(nginx.Repo.Nginx)
	field2 := t2.Elem().Field(0)
	tag2 := field2.Tag.Get("json")
	fmt.Printf("initial: %s\n", tag2)
	newTag2 := `json:"alpine:latest"`
	field2.Tag = reflect.StructTag(newTag2)
	tag2 = field2.Tag.Get("json")
	fmt.Printf("changed: %s\n", tag2)

	fd, err = os.Open(PATH)
	if err != nil {
		panic(err)
	}
	if err := json.NewDecoder(fd).Decode(&nginx); err != nil {
		fmt.Printf("err != nil\n")
		return
	}
	fmt.Printf("embeded&changed:\n%v\n", nginx.Repo.Nginx.Sha)
	fd.Close()
	t = reflect.TypeOf(nginx.Repo)
	field = t.Elem().Field(0)
	tag = field.Tag.Get("json")
	fmt.Printf("tag: %s\n", tag)

	fmt.Printf("\n--- test url encode ---\n")
	terms := []string{"repo:golang/go", "is:open", "json", "decoder"}
	q := url.QueryEscape(strings.Join(terms, " "))
	fmt.Printf("%s\n", q)

	fmt.Printf("\n--- test struct defined in func ---\n")
	/*
		var repoTagFunc = `json:"alpine"`
		var tagTagFunc = `json:"alpine:latest"`
		type RepositoryFunc struct {
			Tag *TagFunc //repoTagFunc
		}
		type TagFunc struct {
			Sha string //tagTagFunc
		}
		var tagFunc RepositoriesFunc
		fmt.Printf("%T\n", tagFunc.Repo)
	*/

	fmt.Printf("\n--- test interface{} ---\n")
	var inter interface{}
	fd, err = os.Open(PATH)
	if err != nil {
		panic(err)
	}
	if err := json.NewDecoder(fd).Decode(&inter); err != nil {
		fmt.Printf("err != nil\n")
		return
	}
	fmt.Printf("interface{}: %v\n", inter)
	fd.Close()
	inter2 := inter.(map[string]interface{})
	//fmt.Printf("inter2: %v\n", inter2)
	for k, v := range inter2 {
		fmt.Printf("key: %v\nvalue: %v\n\n", k, v)
	}
	inter3 := inter2["Repositories"].(map[string]interface{})
	for k, v := range inter3 {
		//fmt.Printf("key: %v\nvalue: %v\n\n", k, v)
		fmt.Printf("\nkey: %v\n", k)
		inter4 := v.(map[string]interface{})
		for k, v := range inter4 {
			fmt.Printf("key: %v\nvalue: %v\n", k, v)
		}
	}
	fmt.Printf("\nrepository count: %d\n", len(inter3))
	repoName := "alpine"
	if repo, ok := inter3[repoName]; ok {
		fmt.Printf("%s > \n%v\n", repoName, repo)
	} else {
		fmt.Printf("%s not found\n", repoName)
	}
	repoName = "10.10.7.161/base/ubuntu"
	if repo, ok := inter3[repoName]; ok {
		fmt.Printf("%s > \n%v\n", repoName, repo)
	} else {
		fmt.Printf("%s not found\n", repoName)
	}
	repoName = "10.107.161/base/ubuntu"
	if repo, ok := inter3[repoName]; ok {
		fmt.Printf("%s > \n%v\n", repoName, repo)
	} else {
		fmt.Printf("%s not found\n", repoName)
	}
}
