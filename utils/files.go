package utils

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"zestream-server/constants"
)

/*
GetDownloadFilePathName returns the absolute path of the file present in downloads folder
takes fileName as argument
*/
func GetDownloadFilePathName(fileName string) (string, error) {
	cwd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	pathName := path.Join(cwd, constants.DOWNLOAD_FILE_PATH_PREFIX)

	err = createDirPath(pathName)

	if err != nil {
		return "", err
	}

	newPath := path.Join(pathName, fileName)

	return newPath, nil
}

/*
GetOutputFilePathName returns the absolute path of the file present in downloads folder
takes fileName as argument
*/
func GetOutputFilePathName(fileName string, postfix string) (string, error) {
	cwd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	// TODO: replace filename with id
	fileName = strings.Replace(fileName, ".", "_", 1)

	pathName := path.Join(cwd, constants.OUTPUT_FILE_PATH_PREFIX, fileName)

	err = createDirPath(pathName)

	if err != nil {
		return "", err
	}

	newPath := path.Join(pathName, postfix)

	return newPath, nil
}

// createDirPath creates a directory at the given path
func createDirPath(pathName string) error {

	err := os.MkdirAll(pathName, os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}

// RemoveExtensionFromFile returns the fileName without the extension, if fileName doesn't end with ext, it returns the fileName
func RemoveExtensionFromFile(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

// IsFileValid checks if the file is present on that path, returns true, if file is there else false
func IsFileValid(filePath string) bool {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

// StringToArgsGenerator takes a map of arguments, and it returns the command line form of arguments
func StringToArgsGenerator(args map[string]string) string {
	var argsStr bytes.Buffer

	for k, v := range args {
		argsStr.WriteString("-")
		argsStr.WriteString(k + " ")
		argsStr.WriteString(v)

		if v != "" {
			argsStr.WriteString(" ")
		}
	}

	return argsStr.String()
}

// DeleteFiles deletes the file/s given in the filePaths
func DeleteFiles(filePaths string) error {
	_, err := exec.Command("rm", strings.Split(filePaths, " ")...).Output()

	return err
}
