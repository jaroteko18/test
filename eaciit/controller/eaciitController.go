package controller

import (
	"encoding/json"
	"exercise/eaciit/model"
	"exercise/eaciit/variable"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"sync"

	"io"
	"os"
	"time"

	"github.com/eaciit/knot/knot.v1"
	"github.com/eaciit/toolkit"
)

type EaciitController struct {
}

var mx sync.RWMutex
var wg sync.WaitGroup
var arNoise [][]float32
var arMax []float32
var a1, a2 int

func (w *EaciitController) Exercise1(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputJson

	var knots model.ExerciseBasic
	knots.DTime, _ = time.Parse(time.RFC3339, "2017-07-22T01:50:34+00:00")
	return knots
}

// func (w *EaciitController) Ex1(r *knot.WebContext) interface{} {

// 	r.Config.OutputType = knot.OutputJson
// 	model := model.ExerciseBasic
// 	r.GetPayload(&model)
// 	t, _ := time.Parse(time.RFC3339, model.DTime)
// 	// RFC3339=2006-01-02T15:04:05Z07:00
// 	return t
// }

func (w *EaciitController) Exercise2(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputTemplate
	m := toolkit.M{}
	return m
}

func (w *EaciitController) Ex2(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputJson
	var knots model.ExerciseBasic
	r.GetPayload(&knots)
	// knots.DTime, _ = time.Parse(time.RFC3339, knots.GetDTime)
	// RFC3339=2006-01-02T15:04:05Z07:00
	return knots.DTime.Add(time.Hour * 12)
}

func (w *EaciitController) Exercise4(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputTemplate

	model := model.File
	mapHeader, _, err := r.GetPayloadMultipart(nil)
	if err != nil {
		model.IsError = true
		model.Message = err.Error()
	}
	fmt.Println("MAPHEADER", len(mapHeader))
	for key, _ := range mapHeader {
		fmt.Println(key)
	}

	file, err := mapHeader["import"][0].Open()
	// import >> Name dr input
	if err != nil {
		model.IsError = true
		model.Message = "OPEN ERROR" + err.Error()
	}

	defer file.Close()
	//xlsx, err := excelize.OpenReader(file)

	curTime := time.Now().Format("20060102150405")
	fName := curTime + "_" + mapHeader["import"][0].Filename
	fLink := variable.WebUploadDir + curTime + "_" + mapHeader["import"][0].Filename
	output, err := os.Create(fLink)
	if err != nil {
		model.IsError = true
		model.Message = "CREATE ERROR" + err.Error()
	}
	defer output.Close()
	_, err = io.Copy(output, file)
	if err != nil {
		model.IsError = true
		model.Message = "COPY ERROR" + err.Error()
	}

	m := toolkit.M{}
	m.Set("flink", fLink)
	m.Set("fname", fName)
	return m

}

// func (w *EaciitController) Download(r *knot.WebContext) interface{} {
// 	r.Config.OutputType = knot.OutputJson
// 	model := model.File
// 	r.GetPayload(&model)
// 	fmt.Println(model.Fname)
// 	// r.Writer.Header().Set("Content-Disposition", "attachment; filename='"+params["fname"]+"'")
// 	// r.Writer.Header().Set("Content-Type", r.Writer.Header().Get("Content-Type"))

// 	// f, err := os.Open(variable.WebUploadDir + model.Filename)
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// }

// 	// io.Copy(r.Writer, f)

// 	return model
// }

// getHost tries its best to return the request host.
// func getUrl(w http.ResponseWriter, r *http.Request) string {
// 	if r.URL.IsAbs() {
// 		host := r.Host
// 		// Slice off any port information.
// 		if i := strings.Index(host, ":"); i != -1 {
// 			host = host[:i]
// 		}
// 		return host
// 	}
// 	return r.URL.Host
// }

func (w *EaciitController) Exercise5(r *knot.WebContext) interface{} {
	a1 = 10 // INPUT
	a2 = 10 // INPUT

	arMax = arMax[:0]

	GridDefinition()

	wg.Add(a1 * a2)

	for i := 0; i < a1; i++ {
		for j := 0; j < a2; j++ {
			// go powMe(&wg, row, &sts)
			go DotProduct(i, j)
		}
	}

	wg.Wait()

	r.Config.OutputType = knot.OutputJson

	return arMax
}

// GridDefinition a
func GridDefinition() {
	// var noise = make([][]float32, x)
	arNoise = make([][]float32, a1)

	for i := 0; i < a1; i++ {
		var noiseY = make([]float32, a2)
		arNoise[i] = noiseY

		for j := 0; j < a2; j++ {
			arNoise[i][j] = rand.Float32()
		}
	}

}

// DotProduct a
func DotProduct(i, j int) {
	mx.Lock()

	// set cons
	var samplePeriod float64 = math.Pow(2, 1)
	var freq float32 = float32(1) / float32(samplePeriod)

	//calculate the horizontal sampling indices
	smple_i0 := (i / int(samplePeriod)) * int(samplePeriod)
	smple_i1 := (smple_i0 + int(samplePeriod)) % a1
	h_blend := (float32(i) - float32(smple_i0)) * freq

	// calculate the vertical sampling indices
	smple_j0 := (j / int(samplePeriod)) * int(samplePeriod)
	smple_j1 := (smple_j0 + int(samplePeriod)) % a2
	v_blend := (float32(j) - float32(smple_j0)) * freq

	//blend the top two corners
	top := Interpolate(arNoise[smple_i0][smple_j0], arNoise[smple_i1][smple_j0], h_blend)

	//blend the bottom two corners
	bot := Interpolate(arNoise[smple_i0][smple_j1], arNoise[smple_i1][smple_j1], h_blend)

	// return Interpolate(top, bot, v_blend)
	arMax = append(arMax, Interpolate(top, bot, v_blend))
	// arMax[row] = Interpolate(top, bot, v_blend)

	// row++
	defer mx.Unlock()
	defer wg.Done()

}

// Interpolate a
func Interpolate(x0, x1, alpha float32) float32 {
	return x0*(1-alpha) + alpha*x1
}

func (w *EaciitController) Exercise6(r *knot.WebContext) interface{} {

	name := "octocat"
	resp, err := http.Get("https://api.github.com/users/" + name)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	r.Config.OutputType = knot.OutputJson

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	user := model.User
	err = json.Unmarshal([]byte(string(body)), &user)
	// err = json.NewDecoder(string(body)).Decode(&user)
	if err != nil {
		panic(err)
	}
	// string(body)
	return user
}
