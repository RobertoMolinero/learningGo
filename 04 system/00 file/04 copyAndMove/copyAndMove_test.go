package main

import (
	"os"
	"testing"
)

type args struct {
	testObject string
	isDir      bool
	dstDir     string
}

type test struct {
	name string
	args args
}

func TestMoveFile(t *testing.T) {
	var tests = []test{
		{"Test case with a file as test object", args{"abc.txt", false, "dstDir"}},
		{"Test case with a directory as test object", args{"dir", true, "dirParent"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dstDir := tt.args.dstDir
			testObject := tt.args.testObject

			// Create the destination directory
			os.Mkdir(dstDir, 0777)
			defer os.RemoveAll(dstDir)

			// Create the test object
			if tt.args.isDir {
				os.Mkdir(testObject, 0777)
			} else {
				_, err := os.Create(testObject)
				if err != nil {
					t.Errorf("Test object could not be created")
				}
			}

			// Move the test object to dstDir
			MoveFile(testObject, dstDir)

			// Check the result (destination directory)
			if _, err := os.Stat(dstDir); os.IsNotExist(err) {
				t.Errorf("Destination directory does not exist")
			}

			// Check the result (test object)
			if _, err := os.Stat(dstDir + "/" + testObject); os.IsNotExist(err) {
				t.Errorf("Test object does not exist")
			}
		})
	}
}

func TestCopyFile(t *testing.T) {
	var tests = []test{
		{"Test case with a file as test object", args{"abc.txt", false, "dstDir"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dstDir := tt.args.dstDir
			testObject := tt.args.testObject

			// Create the destination directory
			os.Mkdir(dstDir, 0777)
			defer os.RemoveAll(dstDir)

			// Create the test object
			if tt.args.isDir {
				os.Mkdir(testObject, 0777)
			} else {
				_, err := os.Create(testObject)
				if err != nil {
					t.Errorf("Test object could not be created")
				}
				defer os.Remove(testObject)
			}

			// Copy the test object to dstDir
			CopyFile(testObject, dstDir+"/"+testObject)

			// Check the result (destination directory)
			if _, err := os.Stat(dstDir); os.IsNotExist(err) {
				t.Errorf("Destination directory does not exist")
			}

			// Check the test object (original)
			if _, err := os.Stat(testObject); os.IsNotExist(err) {
				t.Errorf("Test object (original) does not exist")
			}

			// Check the test object (copy)
			if _, err := os.Stat(dstDir + "/" + testObject); os.IsNotExist(err) {
				t.Errorf("Test object (copy) does not exist")
			}
		})
	}
}
