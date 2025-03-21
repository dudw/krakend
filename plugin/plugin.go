// SPDX-License-Identifier: Apache-2.0

/*
Package plugin provides tools for loading and registering plugins
*/
package plugin

import (
	"os"
	"path/filepath"
	"strings"
)

// Scan returns all the files contained in the received folder with a filename matching the given pattern
func Scan(folder, pattern string) ([]string, error) {
	files, err := os.ReadDir(folder)
	if err != nil {
		return []string{}, err
	}

	var plugins []string

	for _, file := range files {
		if !file.IsDir() && strings.Contains(file.Name(), pattern) {
			plugins = append(plugins, filepath.Join(folder, file.Name()))
		}
	}

	return plugins, nil
}
