package models

type CPU struct {
	Usage float64 `json:"usage,omitempty"`
	Cores int     `json:"cores,omitempty"`
}

type RAM struct {
	Usage float64 `json:"usage,omitempty"`
	Total uint64  `json:"total,omitempty"`
	Used  uint64  `json:"used,omitempty"`
	Free  uint64  `json:"free,omitempty"`
}

type HDD struct {
	Partition string  `json:"partition,omitempty"`
	Usage     float64 `json:"usage,omitempty"`
	Total     uint64  `json:"total,omitempty"`
	Used      uint64  `json:"used,omitempty"`
	Free      uint64  `json:"free,omitempty"`
}

type Stats struct {
	CPU CPU `json:"cpu,omitempty"`
	RAM RAM `json:"ram,omitempty"`
	HDD HDD `json:"hdd,omitempty"`
}
