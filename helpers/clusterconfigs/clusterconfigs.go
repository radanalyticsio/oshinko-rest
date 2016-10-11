package clusterconfigs

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"github.com/radanalyticsio/oshinko-rest/models"
)

var defaultConfig models.NewClusterConfig = models.NewClusterConfig{1, "", 1}

const defaultname = "default"
const configpath = "/etc/oshinko-cluster-configs/"
const globpath = configpath + "%s\\.*"
const failOnMissing = true
const allowMissing = false

func assignConfig(res *models.NewClusterConfig, src models.NewClusterConfig) {
	if src.MasterCount != 0 {
		res.MasterCount = src.MasterCount
	}
	if src.WorkerCount != 0 {
		res.WorkerCount = src.WorkerCount
	}
}

func checkConfiguration(config models.NewClusterConfig) error {
	var err error
	if config.MasterCount == 0 {
		err = errors.New("Cluster configuration may not have masterCount of 0")
	} else if config.WorkerCount == 0 {
		err = errors.New("Cluster configuration may not have workerCount of 0")
	}
	return err
}

func getInt(filename string) (res int64, err error) {
	fd, err := os.Open(filename)
	if err == nil {
		_, err = fmt.Fscanf(fd, "%d", &res)
		fd.Close()
	}
	return res, err
}

func process(config *models.NewClusterConfig, nameElements []string, filename string) error {

	var err error

	// At present we only have a single level of configs, but if/when we have
	// nested configs then we would descend through the levels beginning here with
	// the first element in the name
	switch nameElements[0] {
	case "mastercount":
		config.MasterCount, err = getInt(filename)
	case "workercount":
		config.WorkerCount, err = getInt(filename)
	}
	return err
}


func readConfig(name string, res *models.NewClusterConfig, failOnMissing bool) (err error) {

	filelist, err := filepath.Glob(fmt.Sprintf(globpath, name))
	if err == nil {
		if failOnMissing == true && len(filelist) == 0 {
			return errors.New(fmt.Sprintf("Named config '%s' does not exist", name))
		}
		for _, v := range (filelist) {

			// Break up each filename into elements by "."
			// The first element of every filename will be the config name, dump it
			elements := strings.Split(filepath.Base(v), ".")[1:]
			err = process(res, elements, v)
			if err != nil {
				break
			}
		}
	}
	return
}

func loadConfig(name string) (res models.NewClusterConfig, err error) {
	// If the default config has been modified use those mods.
	// This can probably be smarter, assuming file timestamps
	// work for ConfigMap volumes.
	res = defaultConfig
	err = readConfig(defaultname, &res, allowMissing)
	if err == nil && name != "" {
		err = readConfig(name, &res, failOnMissing)
	}
	return res, err
}

func GetClusterConfig(config *models.NewClusterConfig) (res models.NewClusterConfig, err error) {
	var name string = ""
	if config != nil {
		name, _ = config.Name.(string)
	}
	res, err = loadConfig(name)
	if err == nil && config != nil {
		assignConfig(&res, *config)
	}

	// Check that the final configuration is valid
	if err == nil {
		err = checkConfiguration(res)
	}
	return res, err
}
