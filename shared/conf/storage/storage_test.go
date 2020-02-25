package storage

import (
	"reflect"
	"testing"
)

func TestStoragePathBuilder(t *testing.T) {
	type testCase struct {
		name string
		args struct {
			subDirs []string
			setting string
		}
		expected *ConfigurationPath
		wantErr  bool
	}
	cases := []testCase{
		{
			name: "Regular path",
			args: struct {
				subDirs []string
				setting string
			}{
				subDirs: []string{
					"producer",
					"job",
				},
				setting: "frequency",
			},
			expected: NewConfigurationPath().
				Dir("producer").
				Dir("job").
				Setting("frequency"),
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.name)

			newPath := NewConfigurationPath()

			for i := range tt.args.subDirs {
				newPath.Dir(tt.args.subDirs[i])
			}
			newPath.Setting(tt.args.setting)

			if !reflect.DeepEqual(tt.expected, newPath) && !tt.wantErr {
				t.Errorf("different pathes, expected %v, got %v",
					tt.expected, newPath)
				return
			}
		})
	}
}
