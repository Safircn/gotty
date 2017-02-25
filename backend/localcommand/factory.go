package localcommand

import (
	"syscall"

	"github.com/yudai/gotty/webtty"
)

type Options struct {
	CloseSignal int `hcl:"close_signal" flagName:"close-signal" flagSName:"" flagDescribe:"Signal sent to the command process when gotty close it (default: SIGHUP)" default:"1"`
}

type Factory struct {
	command string
	argv    []string
	options *Options
}

func NewFactory(command string, argv []string, options *Options) (*Factory, error) {
	return &Factory{
		command: command,
		options: options,
	}, nil
}

func (factory *Factory) Name() string {
	return "local command"
}

func (factory *Factory) New(params map[string][]string) (webtty.Slave, error) {
	argv := factory.argv
	// todo conststant?
	if params["args"] != nil && len(params["args"]) > 0 {
		argv = append(argv, params["args"]...)
	}
	return New(
		factory.command,
		argv,
		WithCloseSignal(syscall.Signal(factory.options.CloseSignal)),
	)
}
