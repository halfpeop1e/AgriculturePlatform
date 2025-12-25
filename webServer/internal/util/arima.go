package util

import (
	"errors"
	"math"
	"math/rand"
)

// ARIMAModel 实现简化的ARIMA时间序列预测模型
type ARIMAModel struct {
	P int // AR项的阶数
	D int // 差分阶数
	Q int // MA项的阶数
}

// NewARIMAModel 创建ARIMA模型
func NewARIMAModel(p, d, q int) *ARIMAModel {
	return &ARIMAModel{
		P: p,
		D: d,
		Q: q,
	}
}

// Difference 计算时间序列的差分
func (m *ARIMAModel) Difference(data []float64, d int) []float64 {
	if d <= 0 {
		return data
	}

	result := make([]float64, len(data))
	copy(result, data)

	for i := 0; i < d; i++ {
		temp := make([]float64, len(result)-1)
		for j := 1; j < len(result); j++ {
			temp[j-1] = result[j] - result[j-1]
		}
		result = temp
	}

	return result
}

// Autoregression 计算自回归系数
func (m *ARIMAModel) Autoregression(data []float64, p int) []float64 {
	if p <= 0 || len(data) <= p {
		return make([]float64, p)
	}

	n := len(data)
	coefficients := make([]float64, p)

	// 使用Yule-Walker方程的简化版本
	for i := 0; i < p; i++ {
		var numerator, denominator float64
		count := 0

		for t := p; t < n; t++ {
			numerator += data[t] * data[t-i-1]
			denominator += data[t-i-1] * data[t-i-1]
			count++
		}

		if denominator > 0 {
			coefficients[i] = numerator / denominator
		}
	}

	return coefficients
}

// MovingAverage 计算移动平均系数
func (m *ARIMAModel) MovingAverage(data []float64, q int) []float64 {
	if q <= 0 || len(data) <= q {
		return make([]float64, q)
	}

	coefficients := make([]float64, q)

	// 简化的MA系数计算
	for i := 0; i < q; i++ {
		var sum float64
		count := 0

		for t := i; t < len(data); t++ {
			sum += data[t]
			count++
		}

		if count > 0 {
			coefficients[i] = sum / float64(count)
		}
	}

	return coefficients
}

// Fit 拟合ARIMA模型
func (m *ARIMAModel) Fit(data []float64) error {
	if len(data) < 3 {
		return errors.New("数据点太少，无法拟合ARIMA模型")
	}

	// 数据预处理：确保数据平稳
	differenced := m.Difference(data, m.D)

	if len(differenced) < m.P+1 {
		return errors.New("差分后数据不足以拟合模型")
	}

	return nil
}

// Forecast 进行预测
func (m *ARIMAModel) Forecast(data []float64, steps int) ([]float64, error) {
	if len(data) == 0 {
		return nil, errors.New("输入数据为空")
	}

	if len(data) < 3 {
		// 数据不足时，使用简单平均作为预测
		forecasts := make([]float64, steps)
		avg := m.calculateAverage(data)
		for i := range forecasts {
			forecasts[i] = avg
		}
		return forecasts, nil
	}

	// 如果数据点很少，使用加权移动平均
	if len(data) < 7 {
		return m.weightedMovingAverageForecast(data, steps), nil
	}

	// 对于足够多的数据，使用简化的ARIMA预测
	return m.simplifiedARIMAForecast(data, steps), nil
}

// calculateAverage 计算平均值
func (m *ARIMAModel) calculateAverage(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}

	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum / float64(len(data))
}

// weightedMovingAverageForecast 加权移动平均预测
func (m *ARIMAModel) weightedMovingAverageForecast(data []float64, steps int) []float64 {
	forecasts := make([]float64, steps)
	n := len(data)

	for i := 0; i < steps; i++ {
		// 使用指数加权移动平均
		var weightedSum, weightSum float64
		alpha := 0.3 // 平滑因子

		for j := 0; j < n; j++ {
			weight := math.Pow(1-alpha, float64(n-j-1))
			weightedSum += data[j] * weight
			weightSum += weight
		}

		if weightSum > 0 {
			forecasts[i] = weightedSum / weightSum
		} else {
			forecasts[i] = data[n-1]
		}

		// 将预测值添加到数据序列中，用于下一步预测
		data = append(data, forecasts[i])
		n++
	}

	return forecasts
}

// simplifiedARIMAForecast 简化的ARIMA预测
func (m *ARIMAModel) simplifiedARIMAForecast(data []float64, steps int) []float64 {
	forecasts := make([]float64, steps)

	// 计算趋势
	trend := m.calculateTrend(data)

	// 计算季节性因子（7天周期）
	seasonal := m.calculateSeasonality(data, 7)

	// 计算最近7天的平均值
	recentAvg := m.calculateRecentAverage(data, 7)

	for i := 0; i < steps; i++ {
		// 基础预测值 = 趋势 + 季节性调整
		baseValue := recentAvg + trend*float64(i+1)

		// 添加季节性调整
		seasonalIndex := (len(data) + i) % 7
		if seasonalIndex < len(seasonal) {
			baseValue *= seasonal[seasonalIndex]
		}

		// 添加随机波动（相比原始方法更保守）
		randomFactor := 1 + (rand.Float64()-0.5)*0.06 // ±3%波动
		forecasts[i] = baseValue * randomFactor

		// 确保预测值合理
		if forecasts[i] <= 0 {
			forecasts[i] = recentAvg
		}
	}

	return forecasts
}

// calculateTrend 计算趋势
func (m *ARIMAModel) calculateTrend(data []float64) float64 {
	if len(data) < 2 {
		return 0
	}

	// 使用线性回归计算趋势
	n := float64(len(data))
	var sumX, sumY, sumXY, sumX2 float64

	for i, y := range data {
		x := float64(i)
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
	}

	denominator := n*sumX2 - sumX*sumX
	if denominator == 0 {
		return 0
	}

	slope := (n*sumXY - sumX*sumY) / denominator
	return slope
}

// calculateSeasonality 计算季节性因子
func (m *ARIMAModel) calculateSeasonality(data []float64, period int) []float64 {
	if len(data) < period*2 {
		// 数据不足以计算季节性，返回中性因子
		return make([]float64, period)
	}

	seasonal := make([]float64, period)
	counts := make([]int, period)

	// 计算每个周期的平均值
	for i, value := range data {
		index := i % period
		seasonal[index] += value
		counts[index]++
	}

	// 计算总平均值
	totalAvg := m.calculateAverage(data)

	// 计算季节性因子
	for i := 0; i < period; i++ {
		if counts[i] > 0 && totalAvg > 0 {
			seasonal[i] = (seasonal[i] / float64(counts[i])) / totalAvg
		} else {
			seasonal[i] = 1.0
		}
	}

	return seasonal
}

// calculateRecentAverage 计算最近n天的平均值
func (m *ARIMAModel) calculateRecentAverage(data []float64, n int) float64 {
	if len(data) == 0 {
		return 0
	}

	if n > len(data) {
		n = len(data)
	}

	sum := 0.0
	start := len(data) - n
	for i := start; i < len(data); i++ {
		sum += data[i]
	}

	return sum / float64(n)
}
