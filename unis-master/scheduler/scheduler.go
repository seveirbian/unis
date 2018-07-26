package scheduler

import (
	"errors"
	"strconv"
)

func Schedule(policy string, nodesInfo []NodeInfo, imagetype, requestCPU, requestMem string) (string, error) {
	var nodeName string
	var err error
	switch policy {
	case FirstFit:
		nodeName, err = firstFitSchedule(nodesInfo, imagetype, requestCPU, requestMem)
	default:

	}

	return nodeName, err
}

func firstFitSchedule(nodesInfo []NodeInfo, imagetype, requestCPU, requestMem string) (string, error) {
	for _, node := range nodesInfo {
		if node.NodeEnv != imagetype {
			continue
		}
		if !node.NodeActive {
			continue
		}
		availableCPU := node.TotalCPU
		availableMem := node.TotalMem
		neededCPU, _ := strconv.Atoi(requestCPU)
		neededMem, _ := strconv.Atoi(requestMem)
		var usedCPU int64
		var usedMem int64
		// calculate used cpu at the node
		for _, instance := range node.Instances {
			usedCPU += instance.RequestCPU
			usedMem += instance.RequestMem
		}
		if (availableCPU-usedCPU) >= int64(neededCPU) && (availableMem-usedMem) >= int64(neededMem) {
			return node.NodeName, nil
		}
	}
	return "", errors.New("no nodes can be deployed")
}
