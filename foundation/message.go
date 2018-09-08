package foundation

// A Message is a unit of data that can be disseminated throughout the network.
// An outdated Message can be overwritten by disseminating a newer Message with
// the same `Key` but an incremented `Nonce`. Nodes in the network will discard
// the lower `Nonce` Message in favour of the higher `Nonce` Message. A
// `Signature` is used to verify the authenticity of the Message.
type Message struct {
	Nonce     uint64 `json:"nonce"`
	Key       []byte `json:"key"`
	Value     []byte `json:"value"`
	Signature []byte `json:"signature"`
}
