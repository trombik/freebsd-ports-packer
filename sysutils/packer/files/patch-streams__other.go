--- streams_other.go.orig	2020-04-09 22:22:46 UTC
+++ streams_other.go
@@ -0,0 +1,21 @@
+// +build !windows
+
+package wrappedstreams
+
+import (
+	"os"
+	"sync"
+)
+
+var initOnce sync.Once
+
+func initPlatform() {
+	// These must be initialized lazily, once it's been determined that this is
+	// a wrapped process.
+	initOnce.Do(func() {
+		// The standard streams are passed in via extra file descriptors.
+		wrappedStdin = os.NewFile(uintptr(3), "stdin")
+		wrappedStdout = os.NewFile(uintptr(4), "stdout")
+		wrappedStderr = os.NewFile(uintptr(5), "stderr")
+	})
+}
