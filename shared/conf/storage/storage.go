package storage

import "github.com/pkg/errors"

type ConfigurationStorage interface {
	Get(path *ConfigurationPath) ([]byte, error)
	Save(path *ConfigurationPath, rawConfiguration []byte) error
	Delete(path *ConfigurationPath) error
}

type ConfigurationPath struct {
	dirs    []string
	setting string
}

func NewConfigurationPath() *ConfigurationPath {
	return &ConfigurationPath{}
}

func (rcv *ConfigurationPath) Validate() error {
	if rcv == nil {
		return errors.New("path is empty")
	}

	if rcv.setting == "" {
		return errors.New("setting is empty")
	}

	return nil
}

func (rcv *ConfigurationPath) Dir(dir string) *ConfigurationPath {
	rcv.dirs = append(rcv.dirs, dir)
	return rcv
}

func (rcv *ConfigurationPath) IterateOverDirs(w func(subPath *ConfigurationPath) error) error {
	var (
		walkPath = NewConfigurationPath()
		err      error
	)

	for i := range rcv.dirs {
		if err = w(walkPath.Dir(rcv.dirs[i])); err != nil {
			return err
		}
	}

	return nil
}

func (rcv *ConfigurationPath) GetSettingName() string {
	return rcv.setting
}

func (rcv *ConfigurationPath) Setting(name string) *ConfigurationPath {
	rcv.setting = name
	return rcv
}

func (rcv *ConfigurationPath) Build(rule func(setting string, dirs []string) string) string {
	return rule(rcv.setting, rcv.dirs)
}

func (rcv *ConfigurationPath) Flush() {
	rcv.setting = ""
	rcv.dirs = []string{}
}
