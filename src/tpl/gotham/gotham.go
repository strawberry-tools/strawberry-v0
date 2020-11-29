// Copyright 2020 Gotham Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package gotham

import (
	"encoding/base64"
	"html/template"

	qrcode "github.com/skip2/go-qrcode"
)

func New() *Namespace {
	return &Namespace{}
}

type Namespace struct{}

func (ns *Namespace) QRCoder(url interface{}) (template.HTML, error) {

	var png []byte

	png, err := qrcode.Encode(url.(string), qrcode.Medium, 256)
	if err != nil {
		return "", err
	}

	encodedImage := base64.StdEncoding.EncodeToString(png)

	fullImage := "data:image/png;base64," + encodedImage

	return template.HTML(fullImage), nil
}
