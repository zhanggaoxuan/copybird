package lz4

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/pierrec/lz4"
	"github.com/stretchr/testify/assert"
)

func TestDecompress(t *testing.T) {
	wr := new(bytes.Buffer)
	var decompressor RestoreDecompressLz4

	s := `I bomb atomically, Socrates' philosophies
	And hypotheses can't define how I be droppin' these
	Mockeries, lyrically perform armed robbery
	Flee with the lottery, possibly they spotted me
	Battle-scarred shogun, explosion when my pen hits
	Tremendous, ultra-violet shine blind forensics
	I inspect view through the future see millennium
	Killa Beez sold fifty gold sixty platinum
	Shackling the masses with drastic rap tactics
	Graphic displays melt the steel like blacksmiths
	Black Wu jackets Queen Beez ease the guns in
	Rumblin' patrolmen tear gas laced the function
	Heads by the score take flight incite a war
	Chicks hit the floor, die hard fans demand more
	Behold the bold soldier, control the globe slowly
	Proceeds to blow swingin' swords like Shinobi
	Stomp grounds I pound footprints in solid rock
	Wu got it locked, performin' live on your hottest block`
	hw, err := compress(s)
	assert.NoError(t, err)
	rd := bytes.NewReader(hw)
	assert.NoError(t, decompressor.InitModule(nil))
	assert.NoError(t, decompressor.InitPipe(wr, rd))
	assert.NoError(t, decompressor.Run(context.TODO()))

	assert.Equal(t, s, wr.String())
}

func compress(s string) ([]byte, error) {
	r := bytes.NewReader([]byte(s))
	w := &bytes.Buffer{}
	zw := lz4.NewWriter(w)
	_, err := io.Copy(zw, r)
	if err != nil {
		return nil, err
	}
	if err := zw.Close(); err != nil {
		return nil, err
	}
	return w.Bytes(), nil

}
