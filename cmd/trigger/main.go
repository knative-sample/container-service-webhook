package main

import (
	"flag"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/golang/glog"
	"github.com/knative-sample/container-service-webhook/cmd/trigger/app"
	"github.com/knative-sample/container-service-webhook/cmd/trigger/app/signals"
	"github.com/knative-sample/container-service-webhook/pkg/utils/logs"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	logs.InitLogs()
	defer logs.FlushLogs()

	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	stopCh := signals.SetupSignalHandler()

	glog.Flush()
	// Start runner
	cmd := app.NewCommandStartServer(stopCh)
	cmd.Flags().AddGoFlagSet(flag.CommandLine)
	flag.CommandLine.Parse([]string{})

	if err := cmd.Execute(); err != nil {
		glog.Fatal(err)
	}
}
