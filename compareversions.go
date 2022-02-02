package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"strconv"

	"encoding/json"

	pdpb "github.com/holoplot/sw__protocols_generated/go/product"
	"google.golang.org/protobuf/proto"

	"github.com/hashicorp/go-version"
)

func loadVersionData(filenames []string) (*pdpb.CompatibilityMatrix, error) {

	ret := &pdpb.CompatibilityMatrix{}

	var err error
	var f *os.File

	for _, filename := range filenames {
		f, err = os.Open(filename)
		if err == nil {
			defer f.Close()
			break
		}
	}

	if err != nil {
		return nil, fmt.Errorf("cannot open any of %v", filenames)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("cannot read from %s: %v", f.Name(), err)
	}

	err = proto.Unmarshal(b, ret)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal %s: %v", f.Name(), err)
	}

	return ret, nil

}

func compareVersions(matrix *pdpb.CompatibilityMatrix, systemVersion, planVersion string) bool {
	for _, entry := range matrix.Entries {

		if minS, err := strconv.ParseInt(entry.MinimumSystemVersion, 10, 32); err == nil {
			if systemVersionNumeric, err := strconv.ParseInt(systemVersion, 10, 32); err == nil {
				if minS > systemVersionNumeric {
					//fmt.Printf("systemVersion: %v > %v\n", minS, systemVersionNumeric)

					//fmt.Printf("systemVersion match: %v <= %v\n", minS, systemVersionNumeric)

					v1, err := version.NewVersion(entry.MinimumPlanVersion)

					if err != nil {
						fmt.Printf("Error comparing planVersions: %v\n", err)
						v2, err := version.NewVersion(planVersion)
						if err != nil {
							fmt.Printf("Error comparing planVersions: %v\n", err)
						}
						if v1.GreaterThan(v2) {
							print("planVersion: %v > %v", v1, v2)
							print("planVersion match: %s <= %s", entry.MinimumPlanVersion, planVersion)
						}

					}
				}

			}
			return true
		}
	}
	return false
}

func main() {

	vd, err := loadVersionData([]string{"./versions.dat"})

	if err != nil {
		fmt.Printf("Error executing loadVersionData: %v\n", err)
		os.Exit(-1)
	}

	j, err := json.MarshalIndent(vd, "", "  ")
	if err != nil {
		fmt.Printf("Cannot marshall version data: %v\n", err)
		os.Exit(-1)
	}

	fmt.Printf("debug: %s\n", string(j))

	print(compareVersions(vd, "", ""))

}
