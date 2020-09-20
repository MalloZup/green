package main

import (
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

const hanadbExporter = "prometheus-hanadb_exporter@PRD_HDB00.service"

func configAndInstallHanaMonitoring() {

	// if hana software is detected.

	// TODO: installing pip add. we should research if is more convenient to install the hdbcli like
	// python3 setup.py install
	log.Info("installing and configuring hana monitoring")
	err := zypperInstall("python3-pip")
	if err != nil {
		log.Errorf("package  installation failed")
		os.Exit(1)
	}

	// path where the tarbal is located
	// TODO this should be in a kind of CONF variable, if only the default is not set
	// this is the tarball for hana sql client
	hdbcli := "/tmp/pydbapi/hdbcli-*.tar.gz"
	out, err := exec.Command("/usr/bin/python3", "-m", "pip", "install", hdbcli).CombinedOutput()
	if err != nil {
		log.Errorf("pkg %s is not installed correctly. %s error %s", pkg, out, err)
		// be resilient, try to install the package 3 times
		os.Exit(1)
	}

}

func main() {

	configAndInstallHanaMonitoring()

	// netweaver and  ha

	// log.Info("Starting ha-control..")
}

// SystemctlStatus call systemctl status on service
func systemctlStatus(service string) error {
	out, err := exec.Command("/usr/bin/systemctl", "status", service).CombinedOutput()
	if err != nil {
		log.Errorf("service %s is not running correctly. %s		 error %s", service, out, err)
		return err
	}
	log.Infof("service %s is up and running", service)
	return nil
}

// SystemctlStatus call systemctl status on service
func zypperInstall(pkg string) error {
	out, err := exec.Command("/usr/bin/zypper", "-i", "--non-interactive", "install", pkg).CombinedOutput()
	if err != nil {
		log.Warnf("pkg %s is not installed correctly. %s error %s", pkg, out, err)
		// be resilient, try to install the package 3 times
		exec.Command("/usr/bin/zypper", "-i", "--non-interactive", "install", pkg).CombinedOutput()
		_, err = exec.Command("/usr/bin/zypper", "-i", "--non-interactive", "install", pkg).CombinedOutput()
		return err
	}
	log.Infof("pkg %s installed", pkg)
	return nil
}
