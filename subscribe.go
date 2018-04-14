package main

type subscribe struct {
	MsgType    string   `json:"type"`
	ProductIds []string `json:"product_ids"`

	// fields used for signature
	signature  string
	passphrase string
	timestamp  string
	apiKey     string
}
