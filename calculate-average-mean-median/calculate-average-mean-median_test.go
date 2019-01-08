package calculate_average_mean_median

import "testing"

func TestMeanMedian_GetMinValue(t *testing.T) {
	tables := []struct {
		MeanMedian
		min float64
	}{
		{MeanMedian{
			numbers: []int{
				13, 7, 11, 27, 3,
			},
		}, 3},
	}

	for _, table := range tables {
		result := table.GetMinValue()
		if result != table.min {
			t.Errorf("GetMinValue() returned (%v) an incorrect response, supposed to be %v.", result, table.min)
		}

	}
}

func TestMeanMedian_GetMaxValue(t *testing.T) {
	tables := []struct {
		MeanMedian
		max float64
	}{
		{MeanMedian{
			numbers: []int{
				13, 7, 11, 27, 3,
			},
		}, 27},
	}

	for _, table := range tables {
		result := table.GetMaxValue()
		if result != table.max {
			t.Errorf("GetMaxValue() returned (%v) an incorrect response, supposed to be %v.", result, table.max)
		}

	}
}

func TestMeanMedian_CalcRangeValues(t *testing.T) {
	tables := []struct {
		MeanMedian
		r float64
	}{
		{MeanMedian{
			numbers: []int{
				3, 4, 2, 6,
			},
		}, 4},
	}

	for _, table := range tables {
		result := table.CalcRangeValues()
		if result != table.r {
			t.Errorf("CalcRangeValues() returned (%v) an incorrect response, supposed to be %v.", result, table.r)
		}

	}
}

func TestMeanMedian_CalcMean(t *testing.T) {
	tables := []struct {
		MeanMedian
		r float64
	}{
		{MeanMedian{
			numbers: []int{
				4, 4, 2, 2,
			},
		}, 3},
	}

	for _, table := range tables {
		result := table.CalcMean()
		if result != table.r {
			t.Errorf("CalcMean() returned (%v) an incorrect response, supposed to be %v.", result, table.r)
		}

	}
}

func TestMeanMedian_CalcMedianOdd(t *testing.T) {
	tables := []struct {
		MeanMedian
		r float64
	}{
		{MeanMedian{
			numbers: []int{
				3, 4, 2, 6, 5,
			},
		}, 4},
	}

	for _, table := range tables {
		result := table.CalcMedian()
		if result != table.r {
			t.Errorf("CalcMedian() (with odd) returned (%v) an incorrect response, supposed to be %v.", result, table.r)
		}

	}
}

func TestMeanMedian_CalcMedianEven(t *testing.T) {
	tables := []struct {
		MeanMedian
		r float64
	}{
		{MeanMedian{
			numbers: []int{
				3, 4, 2, 6,
			},
		}, 3.5},
	}

	for _, table := range tables {
		result := table.CalcMedian()
		if result != table.r {
			t.Errorf("CalcMedian() (with even) returned (%v) an incorrect response, supposed to be %v.", result, table.r)
		}

	}
}

func TestMeanMedian_IsOdd(t *testing.T) {
	tables := []struct {
		MeanMedian
		r bool
	}{
		{MeanMedian{
			numbers: []int{
				1, 2,
			},
		}, false},
		{MeanMedian{
			numbers: []int{
				1, 1, 1, 4,
			},
		}, false},
		{MeanMedian{
			numbers: []int{
				1, 1, 1, 1, 5,
			},
		}, true},
	}

	for _, table := range tables {
		result := table.IsOdd()
		if result != table.r {
			t.Errorf("isOdd() returned (%v) an incorrect response: %v, supposed to be %v.", result, table.MeanMedian.numbers, table.r)
		}
	}
}
