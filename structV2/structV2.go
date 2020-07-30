package structV2

type Person struct {
	Uid      string   `json:"uid,omitempty"`
	Name     string   `json:"name,omitempty"`
	From     string   `json:"from,omitempty"`
	NameOFcn string   `json:"nameOFcn,omitempty"`
	NameOFjp string   `json:"nameOFjp,omitempty"`
	NameOFen string   `json:"nameOFen,omitempty"`
	Age      int      `json:"age,omitempty"`
	Friend   []Person `json:"friend,omitempty"`
}
