package imgproc

import server "github.com/jekabolt/imgproc/image-server"

type Config struct {
	Router server.ProcRouter `mapstructure:"Router"`
}
