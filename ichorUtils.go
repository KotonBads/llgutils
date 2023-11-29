package llgutils

import "fmt"

// Sorts Lunar's artifacts by TYPE.
//
// Specify the path to Lunar's artifacts with `path`.
//
// Returns []string: classpath, ichorClassPath, external, natives.
func (data LaunchMeta) SortFiles(path string) (classpath []string, ichorClassPath []string, external []string, natives []string) {
	for _, val := range data.LaunchTypeData.Artifacts {
		switch val.Type {
		case "CLASS_PATH":
			classpath = append(classpath, fmt.Sprintf("%s/%s", path, val.Name))
			ichorClassPath = append(ichorClassPath, val.Name)
		case "EXTERNAL_FILE":
			external = append(external, val.Name)
		case "NATIVES":
			natives = append(natives, fmt.Sprintf("%s/%s", path, val.Name))
		}
	}

	return classpath, ichorClassPath, external, natives
}
