package application

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const sampleAppDescription = `[
	{
		"id": "some-app",
		"name": "Some App",
		"description": "A simple app",
		"redirectURIs": ["some-url"],
		"grantTypes": ["implicit", "refresh_token"],
		"responseTypes": ["token", "code"],
		"scopes": ["permissions"],
		"public": true
	}
]`

func Test_Storage_LoadFromFile_GetAll_success(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "app.*.txt")
	require.NoError(t, err)
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.Write([]byte(sampleAppDescription))
	require.NoError(t, err)

	storage := NewStorage()
	storage.LoadFromFile(tmpfile.Name())

	res := storage.GetAll(context.Background())

	assert.EqualValues(t, []Application{ValidApp}, res)
}

func Test_Storage_LoadFromFile_with_an_invalid_file(t *testing.T) {
	assert.PanicsWithValue(t,
		"failed to read the file content: open some-invalid-file-path: no such file or directory",
		func() {
			storage := NewStorage()
			storage.LoadFromFile("some-invalid-file-path")
		})
}

func Test_Storage_LoadFromFile_with_an_unmarshable_content(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "app.*.txt")
	require.NoError(t, err)
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.Write([]byte("some invalid json"))
	require.NoError(t, err)

	assert.PanicsWithValue(t,
		"failed to decode the application description file: invalid character 's' looking for beginning of value",
		func() {

			storage := NewStorage()
			storage.LoadFromFile(tmpfile.Name())
		})
}
