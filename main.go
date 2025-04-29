package main

import (
	install "patrol_install/steps/install"
	print "patrol_install/utils/print"
)

func main() {
	err := install.Run(&install.PatrolInstaller{})
	if err != nil {
		print.Error("❌ Setup failed")
		print.Error(err.Error())
		print.Error("Please check the logs for more details.")
	}
	print.Success("✅ Setup completed successfully")

}
