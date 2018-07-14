package apiserver

type Server struct {
	Version string
}

type ServerFilePath struct {
	UnisPath         string
	RootPath         string
	ImagesPath       string
	NodesPath        string
	ImagesPublicPath string
	NodesPublicPath  string
	UsersJSONPath    string
}

type userInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ImageInfo struct {
	Repository string
	Tag        string
	ImageID    string
	Created    string
	Size       string
	Type       string
	Owner      string
}

type Instance struct {
	ImageRepository string
	ImageTag        string
	ImageID         string
	InstanceID      string

	RequestCPU int64
	RequestMem int64
	LimitCPU   int64
	LimitMem   int64
}

type NodeInfo struct {
	NodeName       string
	NodeAddr       string
	NodeType       string // public or private
	NodeEnv        string // Docker or Unikernel
	DockerInfo     string
	HypervisorInfo string

	TotalCPU int64
	TotalMem int64

	images    []ImageInfo
	Instances []Instance
}
