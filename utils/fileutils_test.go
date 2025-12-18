package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDirIfExists(t *testing.T) {
	t.Run("directory exists and is removed successfully", func(t *testing.T) {
		// Create a temporary directory
		tmpDir, err := os.MkdirTemp("", "test_remove_dir_*")
		assert.NoError(t, err)
		defer func() {
			// Clean up in case test fails
			_ = os.RemoveAll(tmpDir)
		}()

		// Verify directory exists
		_, err = os.Stat(tmpDir)
		assert.NoError(t, err)

		// Remove the directory
		err = RemoveDirIfExists(tmpDir)
		assert.NoError(t, err)

		// Verify directory no longer exists
		_, err = os.Stat(tmpDir)
		assert.True(t, os.IsNotExist(err))
	})

	t.Run("directory does not exist", func(t *testing.T) {
		// Use a non-existent directory path
		nonExistentDir := filepath.Join(os.TempDir(), "non_existent_dir_12345")

		// Verify directory doesn't exist
		_, err := os.Stat(nonExistentDir)
		assert.True(t, os.IsNotExist(err))

		// Should succeed without error when directory doesn't exist
		err = RemoveDirIfExists(nonExistentDir)
		assert.NoError(t, err)
	})

	t.Run("directory with nested files and subdirectories is removed recursively", func(t *testing.T) {
		// Create a temporary directory
		tmpDir, err := os.MkdirTemp("", "test_nested_dir_*")
		assert.NoError(t, err)
		defer func() {
			// Clean up in case test fails
			_ = os.RemoveAll(tmpDir)
		}()

		// Create nested structure
		subDir := filepath.Join(tmpDir, "subdir")
		err = os.Mkdir(subDir, 0o750)
		assert.NoError(t, err)

		file1 := filepath.Join(tmpDir, "file1.txt")
		err = os.WriteFile(file1, []byte("content1"), 0o600)
		assert.NoError(t, err)

		file2 := filepath.Join(subDir, "file2.txt")
		err = os.WriteFile(file2, []byte("content2"), 0o600)
		assert.NoError(t, err)

		// Verify structure exists
		_, err = os.Stat(tmpDir)
		assert.NoError(t, err)
		_, err = os.Stat(subDir)
		assert.NoError(t, err)
		_, err = os.Stat(file1)
		assert.NoError(t, err)
		_, err = os.Stat(file2)
		assert.NoError(t, err)

		// Remove the directory
		err = RemoveDirIfExists(tmpDir)
		assert.NoError(t, err)

		// Verify entire structure is removed
		_, err = os.Stat(tmpDir)
		assert.True(t, os.IsNotExist(err))
		_, err = os.Stat(subDir)
		assert.True(t, os.IsNotExist(err))
		_, err = os.Stat(file1)
		assert.True(t, os.IsNotExist(err))
		_, err = os.Stat(file2)
		assert.True(t, os.IsNotExist(err))
	})

	t.Run("empty directory is removed successfully", func(t *testing.T) {
		// Create an empty temporary directory
		tmpDir, err := os.MkdirTemp("", "test_empty_dir_*")
		assert.NoError(t, err)
		defer func() {
			// Clean up in case test fails
			_ = os.RemoveAll(tmpDir)
		}()

		// Verify directory exists
		_, err = os.Stat(tmpDir)
		assert.NoError(t, err)

		// Remove the directory
		err = RemoveDirIfExists(tmpDir)
		assert.NoError(t, err)

		// Verify directory no longer exists
		_, err = os.Stat(tmpDir)
		assert.True(t, os.IsNotExist(err))
	})

	t.Run("calling multiple times on same path", func(t *testing.T) {
		// Create a temporary directory
		tmpDir, err := os.MkdirTemp("", "test_multiple_calls_*")
		assert.NoError(t, err)
		defer func() {
			// Clean up in case test fails
			_ = os.RemoveAll(tmpDir)
		}()

		// First call should succeed
		err = RemoveDirIfExists(tmpDir)
		assert.NoError(t, err)

		// Second call should also succeed (directory doesn't exist anymore)
		err = RemoveDirIfExists(tmpDir)
		assert.NoError(t, err)

		// Third call should also succeed
		err = RemoveDirIfExists(tmpDir)
		assert.NoError(t, err)
	})
}
