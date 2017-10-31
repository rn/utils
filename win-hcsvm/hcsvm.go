package main

import (
	"flag"
	"time"

	"github.com/Microsoft/hcsshim"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	imgPath := flag.String("dir", "C:\\Program Files\\Linux Containers", "Directory with initrd.img and bootx64.efi")
	name := flag.String("name", "", "Name of the VM (default a UUID v4)")
	cmdLine := flag.String("cmdline", "console=ttyS0", "Kernel command line arguments")
	createTimeout := flag.Duration("timeout", 10*time.Second, "Timeout for VM creation")

	flag.Parse()

	if *name == "" {
		*name = uuid.NewV4().String()
	}

	logrus.SetLevel(logrus.DebugLevel)

	runtimeCfg := &hcsshim.HvRuntime{
		ImagePath:           *imgPath,
		LinuxInitrdFile:     "initrd.img",
		LinuxKernelFile:     "bootx64.efi",
		LinuxBootParameters: *cmdLine,
	}

	vmCfg := &hcsshim.ContainerConfig{
		SystemType:                  "container",
		Name:                        *name,
		HvPartition:                 true,
		ContainerType:               "linux",
		TerminateOnLastHandleClosed: true,

		HvRuntime: runtimeCfg,
	}

	logrus.Infof("Create VM: %s", *name)

	// XXX For some reason, the CreateContainer just hangs (and
	// then times out after 240s). The VM boots up, but then there
	// seems to be a hand shake missing between the VM and HCS. So
	// we implement our own, shorter timeout.
	var vm hcsshim.Container
	c := make(chan error, 1)
	go func() {
		var err error
		vm, err = hcsshim.CreateContainer(vmCfg.Name, vmCfg)
		c <- err
	}()
	select {
	case err := <-c:
		if err != nil {
			logrus.Fatalf("CreateContainer(): %v\n", err)
		}
	case <-time.After(*createTimeout):
		logrus.Fatalf("CreateContainer(): timed out after %s", createTimeout)
	}

	logrus.Info("Start VM")
	if err := vm.Start(); err != nil {
		logrus.Warnf("Start(): %v\n", err)
		vm.Terminate()
		return
	}

	logrus.Info("Waiting to finish")
	vm.Wait()
}
