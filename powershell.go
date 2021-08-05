// powershell.go
// Copyright (c) 2021 Tobotobo
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

// +build windows

package powershell

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"syscall"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func Execute(cmd string) (string, error) {
	command := exec.Command("powershell.exe")
	cmdLine := fmt.Sprintf(`-Command "chcp 65001 | Out-Null
	$ErrorActionPreference = 'Stop'
	%s`, cmd)
	command.SysProcAttr = &syscall.SysProcAttr{CmdLine: cmdLine}
	var stderr bytes.Buffer
	command.Stderr = &stderr

	out, err := command.Output()
	if err != nil {
		// TODO: powershellのstderrをUTF8で受け取る方法
		errText, _ := sjis_to_utf8(stderr.String())
		return "", errors.New(errText)
	}
	return strings.Trim(string(out), "\r\n"), nil
}

// ShiftJIS から UTF-8
func sjis_to_utf8(str string) (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(str), japanese.ShiftJIS.NewDecoder()))
	if err != nil {
		return "", err
	}
	return string(ret), err
}
