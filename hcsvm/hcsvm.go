package main

import (
	"flag"

	"github.com/Microsoft/hcsshim"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	imgPath := flag.String("dir", "C:\\Program Files\\Linux Containers", "Directory with initrd.img and bootx64.efi")
	name := flag.String("name", "", "Name of the VM (default a UUID v4)")
	extraCmdLine := flag.String("cmdline", "", "Additional kernel command line arguments")

	flag.Parse()

	if *name == "" {
		*name = uuid.NewV4().String()
	}

	cmdLine := "console=ttyS0"
	if *extraCmdLine != "" {
		cmdLine = cmdLine + " " + *extraCmdLine
	}

	logrus.SetLevel(logrus.DebugLevel)

	runtimeCfg := &hcsshim.HvRuntime{
		ImagePath:           *imgPath,
		LinuxInitrdFile:     "initrd.img",
		LinuxKernelFile:     "bootx64.efi",
		LinuxBootParameters: cmdLine,
	}

	vmCfg := &hcsshim.ContainerConfig{
		SystemType:                  "container",
		Name:                        *name,
		Owner:                       "Me",
		HvPartition:                 true,
		ContainerType:               "linux",
		TerminateOnLastHandleClosed: true,

		HvRuntime: runtimeCfg,
	}

	logrus.Info("Create VM: %s", *name)
	vm, err := hcsshim.CreateContainer(vmCfg.Name, vmCfg)
	if err != nil {
		logrus.Fatalf("CreateContainer(): %v\n", err)
	}
	// XXX The above hangs and times out after 240s

	// logrus.Info("Start VM")
	// if err = vm.Start(); err != nil {
	// 	logrus.Warnf("Start(): %v\n", err)
	// 	vm.Terminate()
	// 	return
	// }

	logrus.Info("Waiting to terminate")
	vm.Wait()
}
