package payloads

/*ResendPacket - пакет перенаправления*/
type ResendPacket struct {
	ResolutionMethod  string      `json:"method"`
	ResolutionAddress string      `json:"service"`
	ResolutionRoute   string      `json:"route"`
	Body              interface{} `json:"body"`
}
