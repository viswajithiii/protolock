package protolock

import (
	"errors"
	"os"
)

// Status will report on any issues encountered when comparing the updated tree
// of parsed proto files and the current proto.lock file.
func Status(cfg Config) (*Report, error) {
	updated, err := getUpdatedLock(cfg)
	if err != nil {
		return nil, err
	}

	lockFile, err := openLockFile(cfg)
	if err != nil {
		if os.IsNotExist(err) {
			msg := `no "proto.lock" file found, first run "init"`
			return nil, errors.New(msg)
		}
		return nil, err
	}
	defer lockFile.Close()

	current, err := protolockFromReader(lockFile)
	if err != nil {
		return nil, err
	}

	return compare(current, *updated)
}
