package log

import (
	"fmt"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/gopkg.in/pp.v2"
)

func init() {
	pp.ColoringEnabled = false
}

func ActionGRP(v ...interface{}) {
	args := []interface{}{"[ACTIONGRP] "}
	args = append(args, v...)
	fmt.Println(pp.Sprint(args...))
}

func Action(v ...interface{}) {
	args := []interface{}{"    [ACTION] "}
	args = append(args, v...)
	fmt.Println(pp.Sprint(args...))
}

func Info(v ...interface{}) {
	args := []interface{}{"[INFO] "}
	args = append(args, v...)
	fmt.Println(pp.Sprint(args...))
}

func Error(v ...interface{}) {
	args := []interface{}{"[ERROR] "}
	args = append(args, v...)
	fmt.Println(pp.Sprint(args...))
}
