package runner

import (
	"github.com/peco/peco"
	// "github.com/yamayo/ec2ssh/internal/util"
	"io/ioutil"
	"os"
	"golang.org/x/net/context"
	"bytes"
	"fmt"
)

type Runner struct {
	tmpFile    *os.File
}

func NewRunner() *Runner {
	return &Runner{
	}
}

func (r *Runner) Transform(data string) (string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r.tmpFile, _ = ioutil.TempFile(os.TempDir(), "ec2ssh")
	defer os.Remove(r.tmpFile.Name())

	ioutil.WriteFile(r.tmpFile.Name(), []byte(data), 0644)

	cli := peco.New()
	cli.Argv = []string{"ec2ssh", r.tmpFile.Name()}
	// cli.Stdin = os.Stdin //bytes.NewBufferString(data)
	buf := bytes.Buffer{}
	cli.Stdout = &buf

	blocker := make(chan struct{})
	errCh := make(chan error, 1)

	go func() {
		defer close(blocker)
		if err := cli.Run(ctx); err != nil {
			// switch {
			// case util.IsCollectResultsError(err):
			// 	cli.PrintResults()
			// case util.IsIgnorableError(err):
			// 	if st, ok := util.GetExitStatus(err); ok {
			// 		fmt.Println(st)
			// 	}
			// default:
			// 	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			// }

		}
		// select {
		// case <-ctx.Done():
		// 	fmt.Println("done:", ctx.Err()) // done: context canceled
		// }
	}()

	select {
	case <-blocker:
		fmt.Println("blocker")
	case <-errCh:
	}

	cli.PrintResults()
	return buf.String(), nil
}
