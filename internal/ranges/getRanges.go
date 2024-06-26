package ranges

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func getParameters() map[string][]int {
	return map[string][]int{
		"LFO RATE":       {},
		"LFO DELAY TIME": {},
		"LFO WAVE":       {},
		"LFO TRIG":       {},
		"OSC RANGE":      {},
		"OSC LFO MOD":    {},
		"PWM":            {},
		"PWM SOURCE":     {},
		"SQR SW":         {},
		"SAW SW":         {},
		"SUB LEVEL":      {},
		"NOISE LEVEL":    {},
		"SUB SW":         {},
		"HPF":            {},
		"CUTOFF":         {},
		"RESONANCE":      {},
		"ENV POLARITY":   {},
		"ENV MOD":        {},
		"FLT LFO MOD":    {},
		"FLT KEY FOLLOW": {},
		"AMP MODE":       {},
		"AMP LEVEL":      {},
		"ATTACK":         {},
		"DECAY":          {},
		"SUSTAIN":        {},
		"RELEASE":        {},
		"CHORUS SW":      {},
		"DELAY LEVEL":    {},
		"DELAY TIME":     {},
		"DELAY FEEDBACK": {},
		"DELAY SW":       {},
		"PORTA SW":       {},
		"PORTA TIME":     {},
		"ASSIGN MODE":    {},
		"BEND RANGE":     {},
		"TEMPO SYNC":     {},
	}
}

func createRangeMap(dataDirectory string) string {
	files, err := os.ReadDir(dataDirectory)
	if err != nil {
		panic("failed to read files")
	}
	freshMap := getParameters()
	for _, v := range files {
		if v.IsDir() {
			continue
		}
		if strings.Contains(v.Name(), "PRM") {
			res, err := parseFile(dataDirectory + "/" + v.Name())
			if err != nil {
				fmt.Println(err)
				continue
			}
			for k, _ := range res {
				freshMap[k] = append(freshMap[k], res[k]...)
			}
		}
	}
	m := "map[string][]int{\n"
	for k, v := range freshMap {
		m += fmt.Sprintf("\"%s\": {%d, %d},\n", k, slices.Min(v), slices.Max(v))
	}
	m += "}"
	return m
}

func parseFile(n string) (map[string][]int, error) {
	freshMap := getParameters()
	f, err := os.ReadFile(n)
	if err != nil {
		return freshMap, err
	}
	st := string(f)
	parts := strings.Split(st, "\n")
	var re = regexp.MustCompile(`(?m)(?P<param>^.*\()(?P<val>\d*)`)
	for _, v := range parts {
		if strings.Contains(v, "PATCH_NAME") {
			continue
		}
		parsed := re.FindString(v)
		working := strings.Split(parsed, "(")
		key := strings.TrimRight(working[0], " ")
		val, err := strconv.Atoi(working[1])
		if err != nil {
			return freshMap, err
		}
		freshMap[key] = append(freshMap[key], val)
	}
	return freshMap, nil
}

func GetRangeMap() map[string][]int {
	return map[string][]int{
		"LFO RATE":       {0, 254},
		"DELAY SW":       {0, 1},
		"PWM":            {0, 255},
		"AMP LEVEL":      {59, 255},
		"DELAY FEEDBACK": {0, 15},
		"RELEASE":        {0, 255},
		"ASSIGN MODE":    {0, 3},
		"OSC RANGE":      {0, 2},
		"SAW SW":         {0, 1},
		"NOISE LEVEL":    {0, 255},
		"HPF":            {0, 255},
		"FLT KEY FOLLOW": {0, 255},
		"ATTACK":         {0, 255},
		"TEMPO SYNC":     {0, 1},
		"LFO WAVE":       {0, 6},
		"SQR SW":         {0, 1},
		"ENV POLARITY":   {0, 1},
		"FLT LFO MOD":    {0, 255},
		"AMP MODE":       {0, 1},
		"CUTOFF":         {0, 255},
		"DELAY LEVEL":    {2, 15},
		"OSC LFO MOD":    {0, 255},
		"SUSTAIN":        {0, 255},
		"PORTA SW":       {0, 1},
		"PORTA TIME":     {0, 100},
		"BEND RANGE":     {0, 24},
		"LFO DELAY TIME": {0, 188},
		"LFO TRIG":       {0, 1},
		"RESONANCE":      {0, 255},
		"ENV MOD":        {0, 255},
		"CHORUS SW":      {0, 3},
		"PWM SOURCE":     {0, 2},
		"SUB LEVEL":      {0, 255},
		"SUB SW":         {0, 1},
		"DECAY":          {0, 255},
		"DELAY TIME":     {0, 14},
	}
}
