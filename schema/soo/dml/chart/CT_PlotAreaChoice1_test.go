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

func TestCT_PlotAreaChoice1Constructor(t *testing.T) {
	v := chart.NewCT_PlotAreaChoice1()
	if v == nil {
		t.Errorf("chart.NewCT_PlotAreaChoice1 must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_PlotAreaChoice1 should validate: %s", err)
	}
}

func TestCT_PlotAreaChoice1MarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_PlotAreaChoice1()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_PlotAreaChoice1()
	xml.Unmarshal(buf, v2)
}
