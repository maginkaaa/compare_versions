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

				test5 := compareVersions(tt.args.matrix, "25567", "0.7.3.2") // should return False
				if test5 != false {
					t.Errorf("compareVersions() = %v, want %v", test5, false)
				}

				test6 := compareVersions(tt.args.matrix, "1375", "0.10.0") // should return True
				if test6 != true {
					t.Errorf("compareVersions() = %v, want %v", test6, true)
				}

				test7 := compareVersions(tt.args.matrix, "1375", "0.9.0") // should return False
				if test7 != false {
					t.Errorf("compareVersions() = %v, want %v", test7, false)
				}

				test8 := compareVersions(tt.args.matrix, "this is test", "0.9.0") // should return False
				if test8 != false {
					t.Errorf("compareVersions() = %v, want %v", test8, false)
				}

				test9 := compareVersions(tt.args.matrix, "this is test", "wrong") // should return False
				if test9 != false {
					t.Errorf("compareVersions() = %v, want %v", test9, false)
				}

				test10 := compareVersions(tt.args.matrix, "1289", "0.8.0") // should return False
				if test10 != false {
					t.Errorf("compareVersions() = %v, want %v", test10, false)
				}

				test11 := compareVersions(tt.args.matrix, "1289", "0.9.1") // should return True
				if test11 != true {
					t.Errorf("compareVersions() = %v, want %v", test11, true)
				}

				test12 := compareVersions(tt.args.matrix, "1375", "0.10.1") // should return True
				if test12 != true {
					t.Errorf("compareVersions() = %v, want %v", test12, true)
				}

				test13 := compareVersions(tt.args.matrix, "1110", "0.7.2") // should return False
				if test13 != false {
					t.Errorf("compareVersions() = %v, want %v", test13, false)
				}

				test14 := compareVersions(tt.args.matrix, "1125", "0.7.2") // should return True
				if test14 != true {
					t.Errorf("compareVersions() = %v, want %v", test14, true)
				}

				test15 := compareVersions(tt.args.matrix, "1001", "0.7.0") // should return False
				if test15 != false {
					t.Errorf("compareVersions() = %v, want %v", test15, false)
				}

				test16 := compareVersions(tt.args.matrix, "543872", "0.7.8") // should return False
				if test16 != false {
					t.Errorf("compareVersions() = %v, want %v", test16, false)
				}

				test17 := compareVersions(tt.args.matrix, "1125", "0.7.1") // should return False
				if test17 != false {
					t.Errorf("compareVersions() = %v, want %v", test17, false)
				}

				test18 := compareVersions(tt.args.matrix, "1289", "0.9.0") // should return True
				if test18 != true {
					t.Errorf("compareVersions() = %v, want %v", test18, true)
				}

				test19 := compareVersions(tt.args.matrix, "1290", "0.9.0") // should return True
				if test19 != true {
					t.Errorf("compareVersions() = %v, want %v", test16, true)
				}

				test20 := compareVersions(tt.args.matrix, "1289", "0.8.9") // should return False
				if test20 != false {
					t.Errorf("compareVersions() = %v, want %v", test20, false)
				}

			}
		})
	}

}
