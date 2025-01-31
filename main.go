package main
//adding comment for  check
//created new branch
//new main branch
import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Brotchu/tasks/cmd"
	"github.com/Brotchu/tasks/db"

	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println("[ERR]", err)
		os.Exit(1)
	}
}
