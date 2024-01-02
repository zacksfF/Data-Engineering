package wine

import (
	"fmt"
	"sort"

	"github.com/zacksfF/Data-Engineering/Data-pipeline/internal/api"
	"gonum.org/v1/gonum/stat"
	"google.golang.org/protobuf/proto"
)

// PrintStats ...
func PrintStats(in interface{}, data proto.Message) error {
	d := data.(*api.Interim)
	values := d.SepalLength
	sort.Float64s(values)
	mean := stat.Mean(values, nil)
	median := stat.Quantile(0.5, stat.Empirical, values, nil)
	variance := stat.Variance(values, nil)
	stddev := stat.StdDev(values, nil)
	fmt.Printf("SepalLength: %v\n", values)
	fmt.Printf("       Mean: %.2f\n", mean)
	fmt.Printf("     Median: %.2f\n", median)
	fmt.Printf("   Variance: %.2f\n", variance)
	fmt.Printf("   Std.Dev.: %.2f\n", stddev)
	return nil
}
