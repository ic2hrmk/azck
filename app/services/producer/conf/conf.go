package conf

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-yaml/yaml"
	"github.com/pkg/errors"

	"github.com/ic2hrmk/azck/shared/checks"
	"github.com/ic2hrmk/azck/shared/conf"
)

type (
	ProducerConfiguration struct {
		JobPreferences          *JobPreferences
		ZookeeperConfigurations conf.ZookeeperPreferences
	}

	JobPreferences struct {
		GenerationFrequency float64 `yaml:"frequency" json:"frequency"`
	}
)

func (rcv *ProducerConfiguration) Validate() error {
	if rcv == nil {
		return errors.New("configuration is empty")
	}

	return validation.ValidateStruct(rcv,
		validation.Field(rcv.JobPreferences.GenerationFrequency, checks.IsZeroOrPositiveFloat64()),
	)
}

func ResolveProducerConfiguration() (*ProducerConfiguration, error) {
	//
	// We won't add any default configuration as the service fully depends on
	// Zookeeper settings
	//

	zookeeperConfigurations, err := resolveZookeeperConfigurations()
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve Zookeeper configurations")
	}

	/*
		TODO:
			- resolve Kafka's configurations
	*/

	jobPreferences, err := fetchJobPreferencesFromZookeeper(zookeeperConfigurations)
	if err != nil {
		return nil, errors.Wrap(err, "unable to fetch configurations from Zookeeper")
	}

	return &ProducerConfiguration{
		JobPreferences:          jobPreferences,
		ZookeeperConfigurations: zookeeperConfigurations,
	}, nil
}

//
// Zookeeper stored values
//

const (
	zookeeperProducerZNodePath = "/azck/config/producer"
	prod
)

func resolveZookeeperConfigurations() (conf.ZookeeperPreferences, error) {
	conf.GetStringVar("")
}

func fetchJobPreferencesFromZookeeper(zookeeper conf.ZookeeperPreferences) (*JobPreferences, error) {
	var (
		rawProducerConfiguration []byte
		jobPreferences           *JobPreferences
		err                      error
	)

	rawProducerConfiguration, err = conf.FetchZookeeperConfiguration(
		zookeeper, zookeeperProducerZNodePath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(rawProducerConfiguration, jobPreferences)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode configuration")
	}

	return jobPreferences, nil
}
