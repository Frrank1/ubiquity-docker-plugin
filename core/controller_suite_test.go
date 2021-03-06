package core_test

import (
	"fmt"
	"log"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var testLogger *log.Logger
var logFile *os.File

func TestController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}

var _ = BeforeEach(func() {
	var err error
	logFile, err = os.OpenFile("/tmp/test-spectrum-scale-plugin.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Failed to setup logger: %s\n", err.Error())
		return
	}
	testLogger = log.New(logFile, "spectrum: ", log.Lshortfile|log.LstdFlags)
})

var _ = AfterEach(func() {
	logFile.Sync()
	logFile.Close()
})
