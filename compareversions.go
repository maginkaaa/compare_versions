package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"strconv"

	pdpb "github.com/holoplot/sw__protocols_generated/go/product"
	"google.golang.org/protobuf/proto"

	"github.com/hashicorp/go-version"
)

func loadMatrix(filenames []string) (*pdpb.CompatibilityMatrix, error) {

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

func compatible(matrix *pdpb.CompatibilityMatrix, systemVersion, planVersion string) bool {

	var minimalSystemVersion, availableSystemVersion int64
	var err error

	if availableSystemVersion, err = strconv.ParseInt(systemVersion, 10, 32); err != nil {
		return false
	}

	for _, entry := range matrix.Entries {
		if minimalSystemVersion, err = strconv.ParseInt(entry.MinimumSystemVersion, 10, 32); err != nil {
			return false
		}

		if availableSystemVersion < minimalSystemVersion {
			continue
		}

		minPlanSemanticVersion, err := version.NewVersion(entry.MinimumPlanVersion)
		if err != nil {
			continue
		}

		semanticPlanVersion, err := version.NewVersion(planVersion)
		if err != nil {
			continue
		}

		if semanticPlanVersion.GreaterThanOrEqual(minPlanSemanticVersion) {
			return true
		}
		return false
	}
	return false
}
