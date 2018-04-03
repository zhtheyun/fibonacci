package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/zhtheyun/fibonacci/lib/config"
	"net/http"
	"strconv"
)

type fibonacciResult struct {
	Data []string
}

// FibonacciHandler is the entry point to handle the fibonacci return logic.
func FibonacciHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	rawNumbers := vars["numbers"]
	//Validate parameter.
	//Note: In router configuration, this field has defined as mandatory and here we only check the scope.
	numbers, err := strconv.ParseUint(rawNumbers, 10, 64)
	if err != nil {
		logrus.Errorf("invalid numbers parameter. value: %v, err:%s", rawNumbers, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rawConfig := req.Context().Value(configKey)

	config, ok := (rawConfig).(config.Config)
	if !ok {
		err := fmt.Errorf("invalid configuration. value: %v", rawConfig)
		logrus.Errorf(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	if numbers >= config.MaximumNumbers {
		err := fmt.Errorf("numbers parameter exceeds maximum limit. numbers is %d, maximum limit is %d", numbers, config.MaximumNumbers)
		logrus.Errorf(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Check the range

	fib := new(fibonacciResult)

	if numbers <= fibCache.Numbers {
		logrus.Debugf("Cache hit!")
		fib.Data = fibCache.Data[:numbers]

	} else {
		logrus.Debugf("Cache missing!")

		result, _, _ := config.Generator.Generate(fibCache.Start, fibCache.Next, numbers-fibCache.Numbers)
		fib.Data = append(fibCache.Data, result...)
	}

	//FIXME: ffjson is a open source project which claim a faster json mashaling
	// we may consider to replace it if necessary.
	jsonResult, err := json.Marshal(fib)

	if err != nil {
		logrus.Errorf("failed to marshal fib result. value: %v, err:%s", fib, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResult)

}

// HomeHandler is a dummy handler for home page.
func HomeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to use fib system!")
}

// HeartBeatHandler is used for external tool to monitor the service status
func HeartBeatHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "I am alive!")
}
