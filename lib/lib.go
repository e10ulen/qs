package lib

import(
	"github.com/spiegel-im-spiegel/logf"
)

func Check(e error){
	if e != nil {
		logf.Errorf(e)
	}
}
