package scheduler

var FirstFit = "FirstFit"

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
	DockerID        string
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

	Images    []ImageInfo
	Instances []Instance
}
