package integration_test

import (
	"encoding/json"
	"math/rand"
	"runtime"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gexec"
)

var fsMounterPath string

func TestIntegration(t *testing.T) {
	suiteName := "FS Mounter Suite"
	if runtime.GOOS != "linux" {
		suiteName = suiteName + " - skipping btrfs tests because non-linux"
	}

	rand.Seed(time.Now().Unix())

	RegisterFailHandler(Fail)
	RunSpecs(t, suiteName)
}

type suiteData struct {
	FSMounterPath string
}

var _ = SynchronizedBeforeSuite(func() []byte {
	fsmPath, err := gexec.Build("github.com/concourse/baggageclaim/cmd/fs_mounter")
	Ω(err).ShouldNot(HaveOccurred())

	data, err := json.Marshal(suiteData{
		FSMounterPath: fsmPath,
	})
	Ω(err).ShouldNot(HaveOccurred())

	return data
}, func(data []byte) {
	var suiteData suiteData
	err := json.Unmarshal(data, &suiteData)
	Ω(err).ShouldNot(HaveOccurred())

	fsMounterPath = suiteData.FSMounterPath
})

var _ = SynchronizedAfterSuite(func() {}, func() {
	gexec.CleanupBuildArtifacts()
})