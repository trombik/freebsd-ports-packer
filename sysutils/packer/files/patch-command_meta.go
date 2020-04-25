--- command/meta.go.orig	2019-12-20 19:14:09 UTC
+++ command/meta.go
@@ -9,7 +9,7 @@ import (
 
 	kvflag "github.com/hashicorp/packer/helper/flag-kv"
 	sliceflag "github.com/hashicorp/packer/helper/flag-slice"
-	"github.com/hashicorp/packer/helper/wrappedreadline"
+	"github.com/hashicorp/packer/helper/wrappedstreams"
 	"github.com/hashicorp/packer/packer"
 	"github.com/hashicorp/packer/template"
 )
@@ -145,7 +145,7 @@ func (m *Meta) ValidateFlags() error {
 
 // StdinPiped returns true if the input is piped.
 func (m *Meta) StdinPiped() bool {
-	fi, err := wrappedreadline.Stdin().Stat()
+	fi, err := wrappedstreams.Stdin().Stat()
 	if err != nil {
 		// If there is an error, let's just say its not piped
 		return false
