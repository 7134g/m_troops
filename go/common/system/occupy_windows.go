package system

/*
#include <io.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <sys/locking.h>
#include <share.h>
#include <fcntl.h>

int IsFileUsed(char* filePath)
{
    int ret = 0;
    int fh = _sopen(filePath, _O_RDWR, _SH_DENYRW,
        _S_IREAD | _S_IWRITE );
    if(-1 == fh)
        ret = 1;
    else
        _close(fh);
    return ret;
}
*/
import "C"
import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"unsafe"
)

func Occupy(path string) bool {

	gbkPath, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(path)), simplifiedchinese.GBK.NewEncoder()))

	in := (*C.char)(unsafe.Pointer(&gbkPath[0]))

	ok := int(C.IsFileUsed(in))
	// 1 占用 0 没占用
	if ok == 1 {
		return true
	} else {
		return false
	}
}

func Permit() {

}
