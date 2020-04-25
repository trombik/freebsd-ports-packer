--- helper/wrappedreadline/wrappedreadline_unix.go.orig	2019-12-20 19:14:09 UTC
+++ helper/wrappedreadline/wrappedreadline_unix.go
@@ -3,15 +3,16 @@
 package wrappedreadline
 
 import (
-	"os"
 	"syscall"
 	"unsafe"
+
+	"github.com/hashicorp/packer/helper/wrappedstreams"
 )
 
 // getWidth impl for Unix
 func getWidth() int {
-	stdoutFd := int(Stdout().Fd())
-	stderrFd := int(Stderr().Fd())
+	stdoutFd := int(wrappedstreams.Stdout().Fd())
+	stderrFd := int(wrappedstreams.Stderr().Fd())
 
 	w := getWidthFd(stdoutFd)
 	if w < 0 {
@@ -42,11 +43,4 @@ func getWidthFd(stdoutFd int) int {
 	}
 
 	return int(ws.Col)
-}
-
-func init() {
-	// The standard streams are passed in via extra file descriptors.
-	wrappedStdin = os.NewFile(uintptr(3), "stdin")
-	wrappedStdout = os.NewFile(uintptr(4), "stdout")
-	wrappedStderr = os.NewFile(uintptr(5), "stderr")
 }
