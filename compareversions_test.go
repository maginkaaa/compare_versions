package main

import (
	"testing"

	pdpb "github.com/holoplot/sw__protocols_generated/go/product"
)

func Test_compareVersions(t *testing.T) {
	type args struct {
		matrix        *pdpb.CompatibilityMatrix
		systemVersion string
		planVersion   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareVersions(tt.args.matrix, tt.args.systemVersion, tt.args.planVersion); got != tt.want {
				t.Errorf("compareVersions() = %v, want %v", got, tt.want)

				test1 := compareVersions(tt.args.matrix, "1000", "0.7.1") // should return False
				if test1 != false {
					t.Errorf("compareVersions() = %v, want %v", test1, false)
				}

				test2 := compareVersions(tt.args.matrix, "1001", "0.7.1") // should return True
				if test2 != true {
					t.Errorf("compareVersions() = %v, want %v", test2, true)
				}

				test3 := compareVersions(tt.args.matrix, "1001", "0.7.0") // should return False
				if test3 != false {
					t.Errorf("compareVersions() = %v, want %v", test3, false)
				}

				test4 := compareVersions(tt.args.matrix, "1128", "0.7.3") // should return True
				if test4 != true {
					t.Errorf("compareVersions() = %v, want %v", test4, true)
				}

			}
		})
	}

}
