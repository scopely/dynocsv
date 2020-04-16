package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
	"strings"
)

type runtime int

const (
	none runtime = iota
	snap
)

const (
	snapEnvName         = "SNAP"
	snapNameEnvName     = "SNAP_NAME"
	snapRevisionEnvName = "SNAP_REVISION"
	homeEnvName         = "HOME"
)

// GetSession returns the AWS session for the corresponding profile.
func GetSession(profile string) *session.Session {
	p := profile
	if profile == "" {
		return setupSessionWithoutProfile()
	}
	return setupSession(p)
}

func setupSessionWithoutProfile() *session.Session {
	return session.Must(session.NewSession())
}

func setupSession(profile string) *session.Session {
	r := detectRuntime()
	if r == snap {
		setActualUserHome()
	}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           profile,
		SharedConfigState: session.SharedConfigEnable,
	}))
	return sess
}

func detectRuntime() runtime {
	if isEnvSet(snapEnvName) || isEnvSet(snapNameEnvName) || isEnvSet(snapRevisionEnvName) {
		return snap
	}
	return none
}

func isEnvSet(name string) bool {
	_, found := os.LookupEnv(name)
	return found
}

func setActualUserHome() {
	home := os.ExpandEnv("$" + homeEnvName)
	_ = os.Setenv(
		homeEnvName,
		strings.TrimSuffix(home, os.ExpandEnv("/snap/$"+snapNameEnvName+"/$"+snapRevisionEnvName)))
}
