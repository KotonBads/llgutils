package llgutils

import "fmt"

func (data LaunchMeta) SortFiles(path string) (classpath []string, external []string) {
	for _, val := range data.LaunchTypeData.Artifacts {
		switch val.Type {
		case "CLASS_PATH":
			classpath = append(classpath, fmt.Sprintf("%s/%s", path, val.Name))
		case "EXTERNAL_FILE":
			external = append(external, fmt.Sprintf("%s/%s", path, val.Name))
		}
	}

	return classpath, external
}
