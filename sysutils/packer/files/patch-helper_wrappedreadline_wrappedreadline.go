--- helper/wrappedreadline/wrappedreadline.go.orig	2019-12-20 19:14:09 UTC
+++ helper/wrappedreadline/wrappedreadline.go
@@ -1,6 +1,8 @@
-// Shamelessly copied from the Terraform repo because it wasn't worth vendoring
-// out two hundred lines of code so Packer could use it too.
+// STOLEN SHAMELESSLY FROM THE TERRAFORM REPO BECAUSE VENDORING OUT
+// WRAPPEDREADLINE AND WRAPPEDSTREAMS FELT LIKE TOO MUCH WORK.
 //
+// "a little copying is better than a lot of dependency"
+//
 // wrappedreadline is a package that has helpers for interacting with
 // readline from a panicwrap executable.
 //
@@ -13,24 +15,24 @@
 package wrappedreadline
 
 import (
-	"os"
 	"runtime"
 
 	"github.com/chzyer/readline"
-	"github.com/mitchellh/panicwrap"
+
+	"github.com/hashicorp/packer/helper/wrappedstreams"
 )
 
 // Override overrides the values in readline.Config that need to be
 // set with wrapped values.
 func Override(cfg *readline.Config) *readline.Config {
-	cfg.Stdin = Stdin()
-	cfg.Stdout = Stdout()
-	cfg.Stderr = Stderr()
+	cfg.Stdin = wrappedstreams.Stdin()
+	cfg.Stdout = wrappedstreams.Stdout()
+	cfg.Stderr = wrappedstreams.Stderr()
 
 	cfg.FuncGetWidth = TerminalWidth
 	cfg.FuncIsTerminal = IsTerminal
 
-	rm := RawMode{StdinFd: int(Stdin().Fd())}
+	rm := RawMode{StdinFd: int(wrappedstreams.Stdin().Fd())}
 	cfg.FuncMakeRaw = rm.Enter
 	cfg.FuncExitRaw = rm.Exit
 
@@ -45,9 +47,9 @@ func IsTerminal() bool {
 	}
 
 	// Same implementation as readline but with our custom fds
-	return readline.IsTerminal(int(Stdin().Fd())) &&
-		(readline.IsTerminal(int(Stdout().Fd())) ||
-			readline.IsTerminal(int(Stderr().Fd())))
+	return readline.IsTerminal(int(wrappedstreams.Stdin().Fd())) &&
+		(readline.IsTerminal(int(wrappedstreams.Stdout().Fd())) ||
+			readline.IsTerminal(int(wrappedstreams.Stderr().Fd())))
 }
 
 // TerminalWidth gets the terminal width in characters.
@@ -78,43 +80,3 @@ func (r *RawMode) Exit() error {
 
 	return readline.Restore(r.StdinFd, r.state)
 }
-
-// Package provides access to the standard OS streams
-// (stdin, stdout, stderr) even if wrapped under panicwrap.
-// Stdin returns the true stdin of the process.
-func Stdin() *os.File {
-	stdin := os.Stdin
-	if panicwrap.Wrapped(nil) {
-		stdin = wrappedStdin
-	}
-
-	return stdin
-}
-
-// Stdout returns the true stdout of the process.
-func Stdout() *os.File {
-	stdout := os.Stdout
-	if panicwrap.Wrapped(nil) {
-		stdout = wrappedStdout
-	}
-
-	return stdout
-}
-
-// Stderr returns the true stderr of the process.
-func Stderr() *os.File {
-	stderr := os.Stderr
-	if panicwrap.Wrapped(nil) {
-		stderr = wrappedStderr
-	}
-
-	return stderr
-}
-
-// These are the wrapped standard streams. These are setup by the
-// platform specific code in initPlatform.
-var (
-	wrappedStdin  *os.File
-	wrappedStdout *os.File
-	wrappedStderr *os.File
-)
