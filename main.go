package main

import "github.com/Zzocker/bookolab/pkg/blog"

func main() {
	l := blog.New(blog.DebugLevel)
	a := 10
	l.Debugf("debug called %+v", a)
	l.Errorf("error called %+v", a)
	l.Infof("info called %+v", a)
	blog.NewWithFields(l, map[string]interface{}{
		"username": "pritam",
		"age":      "56",
	}).Infof("calling info with fields")
}
