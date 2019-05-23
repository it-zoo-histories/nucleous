package payloads

type ResendPacket struct {
	ResolutionAddress string `json:"service"`
	ResolutionRoute   string `json:"route"`
}
