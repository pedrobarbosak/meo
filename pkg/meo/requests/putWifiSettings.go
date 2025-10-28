package requests

type PutWifiSettings struct {
	PrivateNetwork *PrivateNetwork `json:"privateNetwork,omitempty"`
	Band2_4GHz     *BandSettings   `json:"2_4ghz,omitempty"`
	Band5GHz       *BandSettings   `json:"5ghz,omitempty"`
	Version        string          `json:"version"`
}

type PrivateNetwork struct {
	SSID     string `json:"ssid,omitempty"`
	Password string `json:"password,omitempty"`
}

type BandSettings struct {
	Bandwidth     int `json:"bandwidth,omitempty"`
	Channel       int `json:"channel,omitempty"`
	TransmitPower int `json:"transmitPower,omitempty"`
}
