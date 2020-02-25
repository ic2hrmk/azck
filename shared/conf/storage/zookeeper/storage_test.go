// +build integration !unit

package zookeeper

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/ic2hrmk/azck/shared/conf/storage"
)

var (
	TestZookeeperServers = []string{"127.0.0.1"}
)

func TestNewZookeeperConfigurationStorage(t *testing.T) {
	var (
		testServers  = TestZookeeperServers
		settingsPath = storage.NewConfigurationPath().
			Dir("zk").
			Dir("cycle").
			Dir("test").
			Setting("value-" + strconv.FormatInt(time.Now().Unix(), 10))
		settingValue = []byte("KEY=VALUE")
	)

	zkClient, err := NewZookeeperConfigurationStorage(testServers)
	if err != nil {
		t.Errorf("failed to create Zookeeper client, %s", err)
		return
	}

	defer zkClient.Delete(settingsPath) // Clear after tests

	err = zkClient.Save(settingsPath, settingValue)
	if err != nil {
		t.Errorf("failed to persist value, %s", err)
		return
	}

	snapshotSettingValue, err := zkClient.Get(settingsPath)
	if err != nil {
		t.Errorf("failed to retrieve value, %s", err)
		return
	}

	if !reflect.DeepEqual(settingValue, snapshotSettingValue) {
		t.Errorf("persisted and retrieved data is different,"+
			"expcted %v, got %v", settingValue, snapshotSettingValue)
		return
	}

	err = zkClient.Delete(settingsPath)
	if err != nil {
		t.Errorf("failed to retrieve value, %s", err)
		return
	}
}
