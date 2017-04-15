package peco

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"

	"github.com/peco/peco"
	"github.com/yamayo/ec2ssh/peco/internal/util"
)

type Runner struct {
}

func NewRunner() *Runner {
	return &Runner{}
}

func (r *Runner) Select(data string) (string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tmpFile, _ := ioutil.TempFile(os.TempDir(), "ec2ssh")
	defer os.Remove(tmpFile.Name())

	ioutil.WriteFile(tmpFile.Name(), []byte(data), 0644)

	cli := peco.New()
	cli.Argv = []string{"ec2ssh", tmpFile.Name()}
	buf := bytes.Buffer{}
	cli.Stdout = &buf

	resultCh := make(chan error)

	go func() {
		defer close(resultCh)
		resultCh <- cli.Run(ctx)
	}()

	err := <-resultCh

	switch {
	case util.IsCollectResultsError(err):
		cli.PrintResults()
		return buf.String(), nil
	case util.IsIgnorableError(err):
		if _, ok := util.GetExitStatus(err); ok {
			return "", err
		}
		return "", nil
	default:
		return "", err
	}
}
