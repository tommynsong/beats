// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package azure

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "azure", asset.ModuleFieldsPri, AssetAzure); err != nil {
		panic(err)
	}
}

// AssetAzure returns asset data.
// This is the base64 encoded gzipped contents of module/azure.
func AssetAzure() string {
	return "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"
}
