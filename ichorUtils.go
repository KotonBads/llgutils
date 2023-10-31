package llgutils

func (data LaunchMeta) SortFiles() (classpath []string, external []string) {
	for _, val := range data.LaunchTypeData.Artifacts {
		switch val.Type {
		case "CLASS_PATH":
			classpath = append(classpath, val.Name)
		case "EXTERNAL_FILE":
			external = append(external, val.Name)
		}
	}

	return classpath, external
}
