package lsblk

import (
	"encoding/json"
	"os/exec"
	"unraid-rest-api/service/disks/types"
)

type Lsblk struct {
}

func NewService() *Lsblk {
	return &Lsblk{}
}

func (s *Lsblk) GetInfo() types.Lsblk {

	output, _ := exec.Command("lsblk", "--json", "--output-all", "--bytes").Output()

	blockDevices := types.Lsblk{}

	json.Unmarshal(output, &blockDevices)

	return blockDevices
}
