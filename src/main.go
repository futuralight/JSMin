package main

import (
	"C"
	"context"
	"os/exec"
	"strings"
	"time"
)

func main() {

}

//export MinifyJS
func MinifyJS(jsStringC *C.char, pathC *C.char) *C.char {
	path := C.GoString(pathC)
	jsString := C.GoString(jsStringC)
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond*1)
	defer cancel()
	cmd := exec.CommandContext(ctx, path+"/bin/JSMin")
	cmd.Stdin = strings.NewReader(jsString)
	out, _ := cmd.Output()

	return C.CString(string(out))
}
