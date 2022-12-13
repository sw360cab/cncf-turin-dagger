package main

import (
	"dagger.io/dagger"
	"dagger.io/dagger/core"
	"universe.dagger.io/go"
)

dagger.#Plan & {
	actions: {
		getCode: core.#Source & {
			path: "."
		}

		build: {
			goBuild: go.#Build & {
				source: getCode.output
			}
			
			goTest: go.#Test & {
				source:  getCode.output
				package: "./..."
			}
		}
	}
}
