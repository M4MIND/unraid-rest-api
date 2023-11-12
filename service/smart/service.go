package smart

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"unraid-rest-api/service/smart/types"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetDiskSmartInfo(disk string) types.SmartCtl {
	output, _ := exec.Command("smartctl", "-j", "-a", "/dev/"+disk).Output()

	fmt.Println()

	smart := types.SmartCtl{}

	err := json.Unmarshal(output, &smart)

	if err != nil {
		return types.SmartCtl{}
	}

	return smart
}
