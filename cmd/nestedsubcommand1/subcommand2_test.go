package nestedsubcommand1

import (
	"github.com/stretchr/testify/assert"
	"testing"
        "bytes"
        "bufio"
)

func TestRunSubCommand2(t *testing.T) {
  var outputStream bytes.Buffer
  writer := bufio.NewWriter(&outputStream)

  runSubCommand2([]string{"I am a strange loop"}, writer)

  writer.Flush()
  assert.Regexp(t, "I am a strange loop I am a strange loop I am a strange loop", outputStream.String(), "sds")
}
