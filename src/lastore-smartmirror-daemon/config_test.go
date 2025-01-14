package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	testDataPath := "./TemporaryTestDataDirectoryNeedDelete"
	err := os.Mkdir(testDataPath, 0777)
	require.NoError(t, err)
	defer func() {
		err := os.RemoveAll(testDataPath)
		require.NoError(t, err)
	}()
	tmpfile, err := ioutil.TempFile(testDataPath, "config.json")
	require.NoError(t, err)
	defer tmpfile.Close()

	data := []byte(`{"filePath":"/","Enable":true}`)
	err = ioutil.WriteFile(tmpfile.Name(), data, 0777)
	require.NoError(t, err)
	configBefore := newConfig(tmpfile.Name())
	require.NotNil(t, configBefore)
	config := newConfig(tmpfile.Name())
	require.NotNil(t, config)
	err = config.setEnable(!config.Enable)
	require.NoError(t, err)

	// 验证
	configAfter := newConfig(tmpfile.Name())
	require.NotNil(t, configAfter)
	assert.Equal(t, configAfter.Enable, !configBefore.Enable)
}
