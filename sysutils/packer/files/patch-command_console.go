--- command/console.go.orig	2019-12-20 19:14:09 UTC
+++ command/console.go
@@ -9,6 +9,7 @@ import (
 
 	"github.com/chzyer/readline"
 	"github.com/hashicorp/packer/helper/wrappedreadline"
+	"github.com/hashicorp/packer/helper/wrappedstreams"
 	"github.com/hashicorp/packer/packer"
 	"github.com/hashicorp/packer/template"
 	"github.com/hashicorp/packer/template/interpolate"
@@ -115,7 +116,7 @@ func (*ConsoleCommand) AutocompleteFlags() complete.Fl
 
 func (c *ConsoleCommand) modePiped(session *REPLSession) int {
 	var lastResult string
-	scanner := bufio.NewScanner(wrappedreadline.Stdin())
+	scanner := bufio.NewScanner(wrappedstreams.Stdin())
 	for scanner.Scan() {
 		result, err := session.Handle(strings.TrimSpace(scanner.Text()))
 		if err != nil {
