package log

import (
	"gopkg.in/pp.v2"
)

func init() {
	pp.ColoringEnabled = false
}

func ActionGRP(v ...interface{}) {
	args := []interface{}{"    [ACTIONGRP] "}
	args = append(args, v...)
	pp.Println(args)
}

func Action(v ...interface{}) {
	args := []interface{}{"    [ACTION] "}
	args = append(args, v...)
	pp.Println(args)
}

func Info(v ...interface{}) {
	args := []interface{}{" [INFO] "}
	args = append(args, v...)
	pp.Println(args)
}
