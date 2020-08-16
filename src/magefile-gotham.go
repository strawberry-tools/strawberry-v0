// +build mage

package main

func TestGotham() error {

	env := map[string]string{"GOFLAGS": testGoFlags()}

	return runCmd(env, "gotestsum", "./...", "-tags", buildTags())
}
