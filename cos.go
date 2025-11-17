// Package: main
// File: cos.go
// Author: Kavi
// Created: 2025-11-17
// Updated: 2025-11-17
// Description: COS下载封装
//
// -----------------------------------------------------------------------------
// Copyright (c) 2025 Kavi. All rights reserved.
// This file is licensed under the Hsiyue Software License v1.0.
// See the LICENSE file in the project root for the full license text.
// For commercial licensing, contact: kavi@hsiyue.com
// -----------------------------------------------------------------------------

package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"

	cosSDK "github.com/tencentyun/cos-go-sdk-v5"
)

type Client struct {
	c *cosSDK.Client
}

// 实例化cos客户端
func NewClient(bucket, region, secretID, secretKey string) *Client {
	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", bucket, region))
	client := cosSDK.NewClient(&cosSDK.BaseURL{BucketURL: u}, &http.Client{
		Transport: &cosSDK.AuthorizationTransport{
			SecretID:  secretID,
			SecretKey: secretKey,
		},
	})
	return &Client{c: client}
}

// 下载指定Key文件内容
func (cl *Client) Download(key string) ([]byte, error) {
	resp, err := cl.c.Object.Get(context.Background(), key, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
