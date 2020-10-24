// Package profile collects following information
// - The number of requests
// - The fastest time
// - The slowest time
// - The mean & median times
// - The percentage of requests that succeeded
// - Any error codes returned that weren't a success
// - The size in bytes of the smallest response
// - The size in bytes of the largest response
package profile

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/magiciiboy/gurl/pkg/common"
	"github.com/magiciiboy/gurl/pkg/http"
)

// SessionProfile stores all information of profiling
type SessionProfile struct {
	TotalRequest int

	MinResponseTime    int     //us
	MaxResponseTime    int     //us
	MeanResponseTime   float32 //us
	MedianResponseTime float32 //us
	TotalTime          int     //us
	SpentTimes         []int   // us

	MinResponseSize int64 //byte
	MaxResponseSize int64 //byte

	TotalStatusError int
	StatusErrorCodes []int16

	TotalError           int
	Errors               []error
	SuccessfulPercentage int8
}

// Profiler helps to profile a session
type Profiler struct {
	Profile SessionProfile
}

// DoStat computes statistics
func (p *SessionProfile) DoStat() {
	if p.TotalRequest > 0 && len(p.SpentTimes) > 0 {
		p.MinResponseTime, _ = common.GetMin(p.SpentTimes)
		p.MaxResponseTime, _ = common.GetMax(p.SpentTimes)
		p.MeanResponseTime = float32(p.TotalTime / p.TotalRequest)
		p.MedianResponseTime, _ = common.GetMedian(p.SpentTimes)
		p.SuccessfulPercentage = int8(100 * (p.TotalRequest - p.TotalError) / p.TotalRequest)
	}
}

// PrintSummary Print out the profiling result
func (p *SessionProfile) PrintSummary() {
	fmt.Printf("\nSummary:\r\n"+
		"\t- Requests:\t %v \r\n"+
		"\t- Fastest:\t %v ms \r\n"+
		"\t- Slowest:\t %v ms \r\n"+
		"\t- Mean time:\t %v ms \r\n"+
		"\t- Median time:\t %v ms \r\n"+
		"\t- Success Rate:\t %v %% \r\n"+
		"\t- Max Size:\t %v B \r\n"+
		"\t- Min Size:\t %v B \r\n"+
		"\t- Error Status Codes:\t %v \r\n",
		p.TotalRequest,
		float32(p.MinResponseTime)/1000,
		float32(p.MaxResponseTime)/1000,
		float32(p.MeanResponseTime)/1000,
		float32(p.MedianResponseTime)/1000,
		p.SuccessfulPercentage,
		p.MinResponseSize,
		p.MaxResponseSize,
		p.StatusErrorCodes,
	)
}

// DefaultProfiler is the default profiler
var DefaultProfiler = &Profiler{
	Profile: SessionProfile{},
}

// DoRequest sends a request and captures profiling information
func (p *Profiler) DoRequest(client *http.Client, req *http.Request) {

	res, err := p.trackRequestTime(client, req)

	if err != nil {
		p.Profile.TotalError++
		p.Profile.Errors = append(p.Profile.Errors, err)
	}

	if statusCodeRange := res.StatusCode / 100; statusCodeRange == 4 || statusCodeRange == 5 {
		p.Profile.StatusErrorCodes = append(p.Profile.StatusErrorCodes, res.StatusCode)
	}

	size := int64(unsafe.Sizeof(res.Raw))
	p.Profile.MinResponseSize = common.MinInt64NonZero(p.Profile.MinResponseSize, size)
	p.Profile.MaxResponseSize = common.MaxInt64(p.Profile.MaxResponseSize, size)
	p.Profile.TotalRequest++
}

func (p *Profiler) trackRequestTime(client *http.Client, req *http.Request) (*http.Response, error) {
	defer p.timeTrack(time.Now())
	res, err := client.SendRequest(req)
	return res, err
}

func (p *Profiler) timeTrack(start time.Time) {
	dur := time.Since(start)
	us := int(dur.Microseconds())
	p.Profile.TotalTime += us
	p.Profile.SpentTimes = append(p.Profile.SpentTimes, us)
}
