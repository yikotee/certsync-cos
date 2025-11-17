// Package: certsync-cos
// File: main.go
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
	"log"
	"os"
	"path/filepath"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime)

	cfgFile := "config.yaml"
	cfg, err := LoadConfig(cfgFile)
	if err != nil {
		log.Println("加载配置文件失败:", err)
		os.Exit(1)
	}

	// 创建COS客户端
	client := NewClient(
		cfg.COS.Bucket,
		cfg.COS.Region,
		cfg.COS.SecretID,
		cfg.COS.SecretKey,
	)

	// 构建下载路径
	certKey := filepath.Join(cfg.COS.Prefix, cfg.COS.CertPath)
	secrtKey := filepath.Join(cfg.COS.Prefix, cfg.COS.KeyPath)

	// 下载证书
	certBytes, err := client.Download(certKey)
	if err != nil {
		log.Println("下载证书失败:", err)
		os.Exit(1)
	}

	// 下载私钥
	secrtBytes, err := client.Download(secrtKey)
	if err != nil {
		log.Println("下载私钥失败:", err)
		os.Exit(1)
	}

	// 写入本地目录
	if err := os.MkdirAll(cfg.Local.CertPath, 0755); err != nil {
		log.Println("创建证书目录失败:", err)
		os.Exit(1)
	}

	if err := os.MkdirAll(cfg.Local.KeyPath, 0755); err != nil {
		log.Println("创建私钥目录失败:", err)
		os.Exit(1)
	}

	certPath := filepath.Join(cfg.Local.CertPath, cfg.Domain+".crt")
	secrtPath := filepath.Join(cfg.Local.KeyPath, cfg.Domain+".key")

	if err := os.WriteFile(certPath, certBytes, 0600); err != nil {
		log.Println("写入证书失败:", err)
		os.Exit(1)
	}
	if err := os.WriteFile(secrtPath, secrtBytes, 0600); err != nil {
		log.Println("写入私钥失败:", err)
		os.Exit(1)
	}

	log.Println("证书同步完成:", cfg.Domain)
	log.Println("Cert:", certPath)
	log.Println("Key:", secrtPath)

	// reload
	if cfg.Local.ReloadCmd != "" {
		log.Println("执行reload:", cfg.Local.ReloadCmd)
		if err := RunCmd(cfg.Local.ReloadCmd); err != nil {
			log.Println("reload执行失败", err)
		} else {
			log.Println("reload成功")
		}
	}
}
