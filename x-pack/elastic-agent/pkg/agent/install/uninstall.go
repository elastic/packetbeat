package install

import (
	"fmt"
	"os"

	"github.com/kardianos/service"

	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/agent/errors"
)

// Uninstall uninstalls persistently Elastic Agent on the system.
func Uninstall() error {
	// uninstall the current service
	svc, err := newService()
	if err != nil {
		return err
	}
	status, _ := svc.Status()
	if status == service.StatusRunning {
		err := svc.Stop()
		if err != nil {
			return errors.New(
				err,
				fmt.Sprintf("failed to stop service (%s)", ServiceName),
				errors.M("service", ServiceName))
		}
		status = service.StatusStopped
	}
	if status == service.StatusStopped {
		err := svc.Uninstall()
		if err != nil {
			return errors.New(
				err,
				fmt.Sprintf("failed to uninstall service (%s)", ServiceName),
				errors.M("service", ServiceName))
		}
	}

	// remove existing directory
	err = os.RemoveAll(InstallPath)
	if err != nil {
		return errors.New(
			err,
			fmt.Sprintf("failed to remove installation directory (%s)", InstallPath),
			errors.M("directory", InstallPath))
	}

	// remove, if present on platform
	if ShellWrapperPath != "" {
		err = os.Remove(ShellWrapperPath)
		if !os.IsNotExist(err) && err != nil {
			return errors.New(
				err,
				fmt.Sprintf("failed to remove shell wrapper (%s)", ShellWrapperPath),
				errors.M("destination", ShellWrapperPath))
		}
	}

	return nil
}
