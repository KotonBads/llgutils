package llgutils

type LaunchBody struct {
	OS      string
	Version string
	Module  string
	Arch    string
}

type Artifacts struct {
	Name string `json:"name"`
	Sha1 string `json:"sha1"`
	Url  string `json:"url"`
	Type string `json:"type"`
}

type LaunchMeta struct {
	Success        bool `json:"success"`
	LaunchTypeData struct {
		Artifacts []Artifacts `json:"artifacts"`
		MainClass string      `json:"mainClass"`
	} `json:"launchTypeData"`
	Licenses []struct {
		File string `json:"file"`
		URL  string `json:"url"`
		Sha1 string `json:"sha1"`
	} `json:"licenses"`
	Textures struct {
		IndexURL  string `json:"indexUrl"`
		IndexSha1 string `json:"indexSha1"`
		BaseURL   string `json:"baseUrl"`
	} `json:"textures"`
	Jre struct {
		Download struct {
			URL       string `json:"url"`
			Extension string `json:"extension"`
		} `json:"download"`
		ExecutablePathInArchive []string    `json:"executablePathInArchive"`
		CheckFiles              [][]string  `json:"checkFiles"`
		ExtraArguments          []string    `json:"extraArguments"`
		JavawDownload           interface{} `json:"javawDownload"`
		JavawExeChecksum        interface{} `json:"javawExeChecksum"`
		JavaExeChecksum         string      `json:"javaExeChecksum"`
	} `json:"jre"`
}