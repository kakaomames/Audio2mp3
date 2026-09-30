//go:build js && wasm

package main

import (
	"syscall/js"

	"github.com/formeo/go-audio-converter/pkg/converter"
)

func wasmMain() {
	js.Global().Set(
		"convertAudio",
		js.FuncOf(func(this js.Value, args []js.Value) any {

			if len(args) < 3 {
				return "convertAudio(data, inputFormat, outputFormat)"
			}

			input := make([]byte, args[0].Length())
			js.CopyBytesToGo(input, args[0])

			inFmt := converter.Format(args[1].String())
			outFmt := converter.Format(args[2].String())

			output, err := converter.New().ConvertBytes(
				input,
				inFmt,
				outFmt,
			)

			if err != nil {
				return err.Error()
			}

			result := js.Global().
				Get("Uint8Array").
				New(len(output))

			js.CopyBytesToJS(result, output)

			return result
		}),
	)

	println("convertAudio() ready")

	select {}
}
