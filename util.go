// Package: certsync-cos
// File: util.go
// Author: Kavi
// Created: 2025-11-17
// Updated: 2025-11-17
// Description: 文件描述
//
// -----------------------------------------------------------------------------
// Copyright (c) 2025 Kavi. All rights reserved.
// This file is licensed under the Hsiyue Software License v1.0.
// See the LICENSE file in the project root for the full license text.
// For commercial licensing, contact: kavi@hsiyue.com
// -----------------------------------------------------------------------------

package main

import (
	"fmt"
	"os/exec"
)

func RunCmd(cmd string) error {
	c := exec.Command("/bin/sh", "-c", cmd)

	output, err := c.CombinedOutput()
	if err != nil {
		return fmt.Errorf("执行失败: %v,输出:%s", err, output)
	}
	return nil
}
