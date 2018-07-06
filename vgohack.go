package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"os/exec"

	"github.com/CJ-Jackson/vgohack/listjson"
	"github.com/CJ-Jackson/vgohack/modjson"
)

const replace = "}\n{"

const with = "},\n{"

func main() {
	if len(os.Args) <= 2 {
		other()
	}
	if os.Args[1] == "mod" && os.Args[2] == "-json" {
		useListInstead()
		return
	}
	other()
}

func useListInstead() {
	buf := &bytes.Buffer{}
	cmd := exec.Command("vgo", "list", "-m", "-json", "all")
	cmd.Stdin = cmd.Stdin
	cmd.Stderr = cmd.Stderr
	cmd.Stdout = buf
	if nil != cmd.Run() {
		os.Exit(1)
	}

	b := buf.Bytes()
	buf.Reset()

	b = bytes.Replace(b, []byte(replace), []byte(with), -1)
	buf.Write([]byte("["))
	buf.Write(b)
	buf.Write([]byte("]"))

	listJ := []listjson.Module{}
	if err := json.NewDecoder(buf).Decode(&listJ); nil != err {
		log.Fatal(err)
	}
	buf.Reset()

	goMod := &modjson.GoMod{}

	for _, j := range listJ {
		if j.Main {
			goMod.Module = modjson.Module{
				Path:    j.Path,
				Version: j.Version,
			}
			continue
		}

		goMod.Require = append(goMod.Require, modjson.Module{
			Path:    j.Path,
			Version: j.Version,
		})

		if nil != j.Replace {
			goMod.Replace = append(goMod.Replace, modjson.Replace{
				Old: modjson.Module{
					Path:    j.Path,
					Version: j.Version,
				},
				New: modjson.Module{
					Path:    j.Path,
					Version: j.Version,
				},
			})
		}
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("\t", "\t")
	if nil != enc.Encode(*goMod) {
		os.Exit(1)
	}
	os.Exit(0)
}

func other() {
	cmd := exec.Command("vgo", os.Args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if nil != cmd.Run() {
		os.Exit(1)
	}
	os.Exit(0)
}
