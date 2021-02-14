package dot

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func There_is_error(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func CrearArchivo(commands string) {

	var file, err = os.Create("./DotFile.txt")
	if There_is_error(err) {
		return
	}

	defer file.Close()

	var OpenFile, err1 = os.OpenFile("./DotFile.txt", os.O_RDWR, 0644)
	if There_is_error(err1) {
		return
	}

	defer OpenFile.Close()

	_, err = file.WriteString(commands)
	if There_is_error(err1) {
		return
	}

	err = file.Sync()
	if There_is_error(err1) {
		return
	}
}

func Graph() {
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tsvg", "DotFile.txt").Output()
	mode := int(0777)
	ioutil.WriteFile("Diagram.svg", cmd, os.FileMode(mode))
}
