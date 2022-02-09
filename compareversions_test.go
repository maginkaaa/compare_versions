package main

import (
	"testing"

	pdpb "github.com/holoplot/sw__protocols_generated/go/product"
)

func Test_compareVersions(t *testing.T) {
	mm := &pdpb.CompatibilityMatrix{
		Entries: []*pdpb.CompatibilityMatrix_CompatibilityMatrixEntry{
			{
				MinimumSystemVersion: "1001",
				MinimumPlanVersion:   "0.7.1",
			},
			{
				MinimumSystemVersion: "1125",
				MinimumPlanVersion:   "0.7.2",
			},
			{
				MinimumSystemVersion: "1179",
				MinimumPlanVersion:   "0.8.0",
			},
			{
				MinimumSystemVersion: "1289",
				MinimumPlanVersion:   "0.9.0",
			},
			{
				MinimumSystemVersion: "1375",
				MinimumPlanVersion:   "0.10.0",
			},
		},
	}

	type args struct {
		matrix        *pdpb.CompatibilityMatrix
		systemVersion string
		planVersion   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test1 - Not compatible",
			args: args{
				matrix:        mm,
				systemVersion: "1000",
				planVersion:   "0.7.1",
			},
			want: false,
		},
		{
			name: "Test2 - Compatible",
			args: args{
				matrix:        mm,
				systemVersion: "1001",
				planVersion:   "0.7.1",
			},
			want: true,
		},
		{
			name: "Test3 - Not Compatible",
			args: args{
				matrix:        mm,
				systemVersion: "1001",
				planVersion:   "0.7.0",
			},
			want: false,
		},
		{
			name: "Test4 - Compatible",
			args: args{
				matrix:        mm,
				systemVersion: "1128",
				planVersion:   "0.7.3",
			},
			want: true,
		},
		{
			name: "Test5 - Broken Plan Version",
			args: args{
				matrix:        mm,
				systemVersion: "1128",
				planVersion:   "0.7,3",
			},
			want: false,
		},
		{
			name: "Test6 - Broken System Version",
			args: args{
				matrix:        mm,
				systemVersion: "1128Z",
				planVersion:   "0.7,3",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compatible(tt.args.matrix, tt.args.systemVersion, tt.args.planVersion); got != tt.want {
				t.Errorf("compareVersions() = %v, want %v", got, tt.want)
			}
		})
	}

}
