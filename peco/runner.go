package peco

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/peco/peco"
	"github.com/yamayo/ec2ssh/internal/util"
	"golang.org/x/net/context"
)

type Runner struct {
	tmpFile *os.File
}

func NewRunner() *Runner {
	return &Runner{}
}

func (r *Runner) Transform(data string) (string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r.tmpFile, _ = ioutil.TempFile(os.TempDir(), "ec2ssh")
	defer os.Remove(r.tmpFile.Name())

	ioutil.WriteFile(r.tmpFile.Name(), []byte(data), 0644)

	cli := peco.New()
	cli.Argv = []string{"ec2ssh", r.tmpFile.Name()}
	buf := bytes.Buffer{}
	cli.Stdout = &buf

	// blocker := make(chan struct{})
	blocker := make(chan error, 1)

	go func() {
		// defer close(blocker)
		if err := cli.Run(ctx); err != nil {
			blocker <- err
		}
	}()

	err := <-blocker

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
