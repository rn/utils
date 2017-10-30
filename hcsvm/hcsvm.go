package main

import (
	"github.com/Microsoft/hcsshim"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetLevel(logrus.DebugLevel)

	runtimeCfg := &hcsshim.HvRuntime{
		ImagePath:           "C:\\Program Files\\Linux Containers",
		LinuxInitrdFile:     "initrd.img",
		LinuxKernelFile:     "bootx64.efi",
		LinuxBootParameters: "console=ttyS0",
	}

	vmCfg := &hcsshim.ContainerConfig{
		SystemType:                  "container",
		Name:                        "FooBar",
		Owner:                       "Me",
		HvPartition:                 true,
		ContainerType:               "linux",
		TerminateOnLastHandleClosed: true,

		HvRuntime: runtimeCfg,
	}

	logrus.Info("Create VM")
	vm, err := hcsshim.CreateContainer(vmCfg.Name, vmCfg)
	if err != nil {
		logrus.Fatal("CreateContainer(): %v\n", err)
	}

	logrus.Info("Start VM")
	if err = vm.Start(); err != nil {
		logrus.Warnf("Start(): %v\n", err)
		vm.Terminate()
		return
	}

	logrus.Info("Waiting to terminate")
	vm.Wait()
}
