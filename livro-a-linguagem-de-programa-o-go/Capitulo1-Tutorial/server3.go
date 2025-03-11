// Server3 é um servidor mínimo de "eco" e contador
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var pallete = []color.Color{color.Black, color.RGBA{0, 255, 0, 1}, color.White}
var mu sync.Mutex
var count int

const (
	blackIndex  = 0 // próxima cor da paleta
	greeenIndex = 1 // próxima cor da paleta
	whiteIndex  = 2 // primeira cor da paleta
)

func main() {
	//http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", executeLissajous)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler ecoa o componente path do URL requisitado
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// counter ecoa o número de chamadas até agora.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

// executeLissajous executa a animação de Lissajous e envia o resultado para o cliente.
func executeLissajous(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UTC().UnixNano())
	cyclesURL := r.URL.Query()["cycles"]
	cycles, err := strconv.Atoi(cyclesURL[0])
	if err != nil {
		cycles = 5
	}
	lissajous(w, float64(cycles))
}

func lissajous(out io.Writer, cycles float64) {
	const ( // número de revoluções completas do oscilador x
		res     = 0.001 // resolução angular
		size    = 100   // canvas da imagem cobre [-size..+size]
		nframes = 64    // número de quadros da animação
		delay   = 8     // tempo entre quadros em unidades de 10ms
	)

	freq := rand.Float64() * 3.0 // frequência relativa do oscilador y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // diferença de fase
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greeenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTA: ignorando erros de codificação
}
