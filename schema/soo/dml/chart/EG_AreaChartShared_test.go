// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package chart_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/dml/chart"
)

func TestEG_AreaChartSharedConstructor(t *testing.T) {
	v := chart.NewEG_AreaChartShared()
	if v == nil {
		t.Errorf("chart.NewEG_AreaChartShared must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.EG_AreaChartShared should validate: %s", err)
	}
}

func TestEG_AreaChartSharedMarshalUnmarshal(t *testing.T) {
	v := chart.NewEG_AreaChartShared()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewEG_AreaChartShared()
	xml.Unmarshal(buf, v2)
}
