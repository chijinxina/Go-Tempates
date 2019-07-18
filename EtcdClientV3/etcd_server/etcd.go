package etcd_server

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"

	"code.byted.org/tao/tao/unittests/minicluster/common"
)

var (
	etcdExec         = "etcd-%s"
	url              = "http://%s:%d"
	host             = "localhost"
	peerListenPort   = 2380
	clientListenPort = 2379
)

type Etcd struct {
	server   *exec.Cmd
	executor string
	dir      string
}

func New(execDir string) *Etcd {
	etcd := &Etcd{
		executor: path.Join(execDir, fmt.Sprintf(etcdExec, common.OS)),
		dir:      "/tmp/etcd_data",
	}

	etcd.server = exec.Command(etcd.executor,
		"--force-new-cluster="+"true",
		"--listen-peer-urls="+fmt.Sprintf(url, "0.0.0.0", peerListenPort),
		"--data-dir="+etcd.dir,
		"--initial-advertise-peer-urls="+fmt.Sprintf(url, "localhost", peerListenPort),
		"--advertise-client-urls="+fmt.Sprintf(url, "localhost", clientListenPort),
		"--listen-client-urls="+fmt.Sprintf(url, "0.0.0.0", clientListenPort),
	)
	return etcd
}

func (s *Etcd) StartServer() error {
	if s == nil {
		return fmt.Errorf("etcd is nil")
	}

	if err := s.cleanProcess(); err != nil {
		return fmt.Errorf("stop old etcd server failed: %s", err)
	}
	if err := s.server.Start(); err != nil {
		return fmt.Errorf("start etcd server failed: %s", err)
	}
	return nil
}

func (s *Etcd) Close() error {
	if s == nil {
		return nil
	}
	if s.server != nil {
		if err := s.server.Process.Kill(); err != nil {
			return err
		}
	}
	err := os.RemoveAll(s.dir)
	return err
}

func (s *Etcd) cleanProcess() error {
	cmd0 := exec.Command("ps", "-ef")
	cmd1 := exec.Command("grep", path.Base(s.executor))
	cmd2 := exec.Command("grep", "-v", "grep")
	var err error
	if cmd1.Stdin, err = cmd0.StdoutPipe(); err != nil {
		panic(err)
	}
	if cmd2.Stdin, err = cmd1.StdoutPipe(); err != nil {
		panic(err)
	}
	wg := &sync.WaitGroup{}
	run(wg, cmd0)
	run(wg, cmd1)
	wg.Add(1)
	if out, err := cmd2.CombinedOutput(); err != nil {
		if string(out) == "" {
			return nil
		}
		return fmt.Errorf("%s:%s", err, out)
	} else {
		return exec.Command("kill", "-9", strings.Fields(string(out))[1]).Run()
	}
}

func run(wg *sync.WaitGroup, cmd *exec.Cmd) {
	wg.Add(1)
	go func() {
		if err := cmd.Run(); err != nil {
			panic(err)
		}
		wg.Done()
	}()
	return
}
