package libv2ray

type Tunnel struct {
	svType string
	server string
	point  string
}

func NewTunnel(
	svType string,
	server string,
	point string,
) IService {

	return &Tunnel{
		svType: "client",
		server: server,
		point:  point,
	}
}

func (args *Tunnel) Start() {
	var commands = []string{args.svType, args.server, args.point}
	commands = commands[1:]
	client(commands)
	return 
}
