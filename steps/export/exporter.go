package export

import (
	"patrol_install/utils/print"
)

type Exporter interface {
	FindAndExportBuilds() error
}

func Run(exporter Exporter) error {
	print.StepIniciated("--- Getting Patrol builds ---")

	if err := exporter.FindAndExportBuilds(); err != nil {
		// Handle error
	}

	return nil
}
