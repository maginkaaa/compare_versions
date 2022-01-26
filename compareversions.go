package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"golang.org/x/mod/semver"
	"google.golang.org/protobuf/proto"
)


func loadVersionData(filenames []string) ([]*pdpb.Version, string, error) {
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
		return nil, "", fmt.Errorf("cannot open any of %v", filenames)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, "", fmt.Errorf("cannot read from %s: %v", f.Name(), err)
	}

	pd := pdpb.AllVersions{}
	err = proto.Unmarshal(b, &pd)
	if err != nil {
		return nil, "", fmt.Errorf("cannot unmarshal %s: %v", f.Name(), err)
	}

	return pd.Version, f.Name(), nil
}


type Version struct {
	SystemVersion  string
	PlanVersion    string
	VersionData *pdpb.Version
}


func compareVersions() {
	for minSymstemVersion, minPlanVersion in loadVersionData():
        
        if minSystemVersion > systemVersion:
            continue

        if semver.compare(minPlanVersion, planVersion) > 0:
            continue


        return True

    return False
}

func main() {

	fmt.Printf("Hello World")

}