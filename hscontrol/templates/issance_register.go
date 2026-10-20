package templates

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

// IssanceRegister is the page a phone or computer opens when it joins a Tower network from the Tailscale
// app. The customer pastes the issance key of their account, and the device joins their network with no
// administrator involved. message is shown above the form (an error, or empty); done shows the success
// page instead of the form.
func IssanceRegister(formAction, message string, done bool) *elem.Element {
	if done {
		return page(
			"Device added - issance",
			H1(elem.Text("Device added")),
			P(elem.Text("This device is now in your issance network. Go back to the Tailscale app: it connects in a few seconds.")),
		)
	}
	nodes := []elem.Node{
		H1(elem.Text("Add this device to issance")),
		P(elem.Text("Paste the issance key of your account. You find it in the issance installer's last screen, or with `issance key` on your computer.")),
	}
	if message != "" {
		nodes = append(nodes, P(elem.Strong(nil, elem.Text(message))))
	}
	nodes = append(nodes, elem.Form(
		attrs.Props{attrs.Method: "POST", attrs.Action: formAction},
		elem.Input(attrs.Props{
			attrs.Type:         "password",
			attrs.Name:         "key",
			attrs.Placeholder:  "issance key",
			attrs.Autocomplete: "off",
			attrs.Required:     "true",
			attrs.Style:        "width:100%;padding:12px;font-size:16px;margin:12px 0;box-sizing:border-box",
		}),
		elem.Button(attrs.Props{attrs.Type: "submit", attrs.Style: "padding:12px 18px;font-size:16px"}, elem.Text("Add device")),
	))
	return page("Add device - issance", nodes...)
}
