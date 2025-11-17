// Package: main
// File: config.go
// Author: Kavi
// Created: 2025-11-17
// Updated: 2025-11-17
// Description: 加载配置文件
//
// -----------------------------------------------------------------------------
// Copyright (c) 2025 Kavi. All rights reserved.
// This file is licensed under the Hsiyue Software License v1.0.
// See the LICENSE file in the project root for the full license text.
// For commercial licensing, contact: kavi@hsiyue.com
// -----------------------------------------------------------------------------

package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	COS struct {
		Bucket    string `yaml:"bucket"`
		Region    string `yaml:"region"`
		SecretID  string `yaml:"secret_id"`
		SecretKey string `yaml:"secret_key"`
		Prefix    string `yaml:"prefix"`
		CertPath  string `yaml:"cert_path"`
		KeyPath   string `yaml:"key_path"`
	} `yaml:"cos"`

	Local struct {
		CertPath  string `yaml:"cert_path"`
		KeyPath   string `yaml:"key_path"`
		ReloadCmd string `yaml:"reload_cmd"`
	} `yaml:"local"`

	Domain   string `yaml:"domain"`
	Schedule string `yaml:"schedule"`
}

// 读取YAML配置文件
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var c Config
	if err := yaml.Unmarshal(data, &c); err != nil {
		return nil, err
	}

	return &c, nil
}
