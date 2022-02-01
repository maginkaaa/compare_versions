package main

import (
	"fmt"
	"io/ioutil"
	"os"

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

		//fmt.Printf("MinimumPlanVersion: %s -- MinimumSystemVersion: %s \n", entry.MinimumPlanVersion, entry.MinimumSystemVersion)
		if entry.MinimumSystemVersion > systemVersion {
			fmt.Printf("systemVersion: %s > %s\n", entry.MinimumSystemVersion, systemVersion)
		}
		fmt.Printf("systemVersion match: %s <= %s\n", entry.MinimumSystemVersion, systemVersion)

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
		return true
	}
	return false
}

/*
		for minSymstemVersion, minPlanVersion in loadVersionData():

	        if minSystemVersion > systemVersion:
	            continue

	        if semver.compare(minPlanVersion, planVersion) > 0:
	            continue


	        return True

	    return False
*/

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

	compareVersions(vd, "", "")

}
