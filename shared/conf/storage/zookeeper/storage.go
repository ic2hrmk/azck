package zookeeper

import (
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/samuel/go-zookeeper/zk"

	"github.com/ic2hrmk/azck/shared/conf/storage"
)

const (
	zooConnectionTimeout = 60 * time.Second
)

type (
	ZooPreferences struct {
		Servers     []string
		Credentials []ZooCredentials
	}

	ZooCredentials struct {
		Scheme string
		Auth   []byte
	}

	zookeeperStorage struct {
		zkclient *zk.Conn
	}
)

func NewZookeeperConfigurationStorage(
	servers []string,
	credentials ...ZooCredentials,
) (
	storage.ConfigurationStorage, error,
) {
	connection, _, err := zk.Connect(servers, zooConnectionTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to Zookeeper")
	}

	for i := range credentials {
		if err = connection.AddAuth(credentials[i].Scheme, credentials[i].Auth); err != nil {
			connection.Close()
			return nil, errors.Wrap(err, "failed to add auth")
		}
	}

	return &zookeeperStorage{zkclient: connection}, nil
}

func (rcv *zookeeperStorage) Get(path *storage.ConfigurationPath) ([]byte, error) {
	if err := path.Validate(); err != nil {
		return nil, err
	}

	rawConfiguration, _, err := rcv.zkclient.Get(path.Build(zooPathBuilder))
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch configuration")
	}

	return rawConfiguration, nil
}

func (rcv *zookeeperStorage) Save(path *storage.ConfigurationPath, rawConfiguration []byte) error {
	if err := path.Validate(); err != nil {
		return err
	}

	zkPath := path.Build(zooPathBuilder)

	isExist, _, err := rcv.zkclient.Exists(zkPath)
	if err != nil {
		return errors.Wrap(err, "failed to verify path existence")
	}

	if !isExist {

		createSubNodeIfNotExists := func(subPath *storage.ConfigurationPath) error {
			builtSubPath := subPath.Build(zooPathBuilder)

			if isExist, _, _ = rcv.zkclient.Exists(builtSubPath); !isExist {
				if _, err = rcv.zkclient.Create(builtSubPath, []byte{}, 0, zk.WorldACL(zk.PermAll)); err != nil {
					return errors.Wrap(err, "failed to create a nested list")
				}
			}

			return nil
		}

		if err = path.IterateOverDirs(createSubNodeIfNotExists); err != nil {
			return err
		}

		_, err = rcv.zkclient.Create(zkPath, rawConfiguration, 0, zk.WorldACL(zk.PermAll))

	} else {

		_, err = rcv.zkclient.Set(zkPath, rawConfiguration, -1)

	}

	if err != nil {
		return errors.Wrap(err, "failed to persist setting")
	}

	return nil
}

func (rcv *zookeeperStorage) Delete(path *storage.ConfigurationPath) error {
	if err := path.Validate(); err != nil {
		return err
	}

	zkPath := path.Build(zooPathBuilder)

	isExist, _, err := rcv.zkclient.Exists(zkPath)
	if err != nil {
		return errors.Wrap(err, "failed to verify path existence")
	}

	if isExist {
		if err = rcv.zkclient.Delete(zkPath, -1); err != nil {
			return errors.Wrap(err, "failed to remove setting")
		}
	}

	return nil
}

func zooPathBuilder(setting string, dirs []string) string {
	path := "/" + strings.Join(dirs, "/")

	if setting != "" {
		path += "/" + setting
	}

	return path
}
