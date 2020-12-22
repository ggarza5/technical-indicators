
// Package indicators provides functions that can be used as
// indicators for finance products.
package indicators

import "fmt"
// Sma calculates simple moving average of a slice for a certain
// number of time periods.
func (slice mfloat) SMA(period int) []float64 {

	var smaSlice []float64

	for i := period; i <= len(slice); i++ {
		smaSlice = append(smaSlice, Sum(slice[i-period:i])/float64(period))
	}

	return smaSlice
}

func (slice mfloat) RollingMax(period int) []float64 {

	var maxSlice []float64
	// println("period is")
	// println(period)
	// println(len(slice))
	for i := 0; i < len(slice) - period; i++ {
		maxSlice = append(maxSlice, sliceMax(slice[i:i+period]))
		// println("recent value is")
		// println(slice[i])
		// println(slice[i+period])
	}
	// println("rolling past rollingmax")
	return maxSlice
}

func (slice mfloat) RollingMin(period int) []float64 {

	var minSlice []float64

	for i := 0; i < len(slice) - period; i++ {
		minSlice = append(minSlice, sliceMin(slice[i:i+period]))
	}
	// println("rolling past rollingMin")	
	return minSlice
}

// Ema calculates exponential moving average of a slice for a certain
// number of tiSmame periods.
func (slice mfloat) EMA(period int) []float64 {

	var emaSlice []float64

	ak := period + 1
	k := float64(2) / float64(ak)

	emaSlice = append(emaSlice, slice[0])

	for i := 1; i < len(slice); i++ {
		emaSlice = append(emaSlice, (slice[i]*float64(k)) + (emaSlice[i-1]*float64(1-k)))
	}

	return emaSlice
}


// BollingerBands returns upper band, lower band and simple moving
// average of a slice.
func BollingerBands(slice mfloat, period int, nStd float64) ([]float64, []float64, []float64) {

	var upperBand, lowerBand, middleBand mfloat

	middleBand = slice.SMA(period)
	std := Std(middleBand)
	upperBand = middleBand.AddToAll(std * nStd)
	lowerBand = middleBand.AddToAll(-1.0 * std * nStd)

	return middleBand, upperBand, lowerBand
}


// MACD stands for moving average convergence divergence.
func MACD(data mfloat, ema ...int) ([]float64, []float64) {

	var macd, ema1, ema2, ema3 mfloat

	if len(ema) < 3 {
		ema = []int{12, 26, 9}
	}

	ema1 = data.EMA(ema[0])
	ema2 = data.EMA(ema[1])
	macd = SubSlices(ema1, ema2)
	ema3 = macd.EMA(ema[2])

	return macd, ema3
}


// OBV means On Balance Volume.
func OBV(priceData, volumeData mfloat) []float64 {

	obv := []float64{volumeData[0]}

	for i, vol := range volumeData[1:] {
		if priceData[i] > priceData[i-1] {
			obv = append(obv, obv[i-1]+vol)
		} else if priceData[i] < priceData[i-1] {
			obv = append(obv, obv[i-1]-vol)
		} else {
			obv = append(obv, obv[i-1])
		}
	}

	return obv
}


// Ichimoku Cloud.
func IchimokuCloud(priceData, lowData, highData mfloat, configs []int) ([]float64, []float64, []float64,[]float64, []float64) {

	var conversionLine, baseLine, leadSpanA, leadSpanB, lagSpan []float64
	//SubSlices(highData.SMA(configs[0]), lowData.SMA(configs[0]))
	// println("WE GOT PAST CONVERSION LINE CA:C")
	// fmt.Println(highData.SMA(configs[0])[len(highData)-50:])
	// fmt.Println(lowData.SMA(configs[0])[len(lowData)-50:])
	// conversionLine = AddSlices(lowData.DivSlice(SubSlices(highData.SMA(configs[0]), lowData.SMA(configs[0])),2)
	// fmt.Println(highData[len(highData)-configs[0]*2:].RollingMax(configs[0]))
	conversionLine = DivSlice(AddSlices(highData.RollingMax(configs[0]), lowData.RollingMin(configs[0])),2)
	//subtract conversion value from high or add it to low to generate line point
	// println("WE GOT PAST CONVERSION LINE CA:C")
	baseLine	   = DivSlice(AddSlices(highData.RollingMax(configs[1]), lowData.RollingMin(configs[1])),2)
	// println("WE GOT PAST BASE LINE CA:C")
	// fmt.Println(baseLine)	
	// fmt.Println(conversionLine)
	leadSpanA	   = DivSlice(AddSlicesFromReverse(conversionLine, baseLine),2)
	// println("WE GOT PAST leadSpanA LINE CA:C")		
	leadSpanB	   = DivSlice(AddSlices(highData.RollingMax(configs[1]*2), lowData.RollingMin(configs[1]*2)),2)
	// println("WE GOT PAST leadSpanB LINE CA:C")		
	lagSpan		   = priceData[configs[3]:len(priceData)]
	fmt.Println("Done with cloud")
	return conversionLine, baseLine, leadSpanA, leadSpanB, lagSpan
}