package api

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func FormatContainerName(containerName string) string {
	// container names are supposed to be unique
	// since we are using the docker-compose service as the container name
	// make it unique by adding a time.
	// TODO: we should skip creating the container again if already exists
	// instead of creating a uniquely named container name
	now := time.Now()
	f := func(c rune) bool {
		if c == 58 {
			// 58 is the ':' character
			return true
		}
		return false
	}
	return strings.FieldsFunc(containerName, f)[0] + now.Format("2006-02-15-04-05") + strconv.Itoa(rand.Int())
}

func fomatLabels(label string) []string {
	f := func(c rune) bool {
		if c == 58 {
			// 58 is the ':' character
			return true
		} else if c == 61 {
			//61 is '=' char
			return true
		}
		return false
	}
	// TODO: we should trim any whitespace before returning.
	// this will prevent labels like type= web
	return strings.FieldsFunc(label, f)
}

func fomatPorts(port string) []string {
	f := func(c rune) bool {
		if c == 58 {
			// 58 is the ':' character
			return true
		} else if c == 61 {
			//61 is '=' char
			return true
		}
		return false
	}
	// TODO: we should trim any whitespace before returning.
	// this will prevent labels like type= web
	return strings.FieldsFunc(port, f)
}

func fomatServiceVolumes(volume string) []string {
	f := func(c rune) bool {
		if c == 58 {
			// 58 is the ':' character
			return true
		}
		return false
	}
	// TODO: we should trim any whitespace before returning.
	// this will prevent labels like type= web
	return strings.FieldsFunc(volume, f)
}

func fomatRegistryAuth(auth string) []string {
	f := func(c rune) bool {
		if c == 58 {
			// 58 is the ':' character
			return true
		}
		return false
	}
	// TODO: we should trim any whitespace before returning.
	// this will prevent labels like type= web
	return strings.FieldsFunc(auth, f)
}

func formatComposePath(path string) []string {
	f := func(c rune) bool {
		// TODO; check if this is cross platform
		if c == 47 {
			// 47 is the '/' character
			return true
		}
		return false
	}
	// TODO: we should trim any whitespace before returning.
	return strings.FieldsFunc(path, f)
}

type popagateError struct {
	originalErr error
	newErr      error
}

func (p *popagateError) Error() string {
	return fmt.Sprintf("originalErr:: %s \nThisErr:: %s", p.originalErr, p.newErr)
}
