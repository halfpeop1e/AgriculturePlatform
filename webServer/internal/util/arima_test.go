package util

import (
	"testing"
)

func TestARIMAForecast(t *testing.T) {
	// 测试数据：模拟7天的价格数据
	data := []float64{2.5, 2.8, 2.6, 2.7, 2.9, 2.4, 2.6}

	// 创建ARIMA模型
	arima := NewARIMAModel(2, 1, 1)

	// 拟合模型
	err := arima.Fit(data)
	if err != nil {
		t.Logf("模型拟合警告: %v", err)
	}

	// 进行预测
	forecasts, err := arima.Forecast(data, 2)
	if err != nil {
		t.Errorf("预测失败: %v", err)
		return
	}

	if len(forecasts) != 2 {
		t.Errorf("预测结果数量错误，期望2个，实际%d个", len(forecasts))
		return
	}

	// 检查预测值是否合理
	for i, forecast := range forecasts {
		if forecast <= 0 {
			t.Errorf("第%d个预测值不合理: %.2f", i+1, forecast)
		}
	}

	t.Logf("测试数据: %v", data)
	t.Logf("预测结果: %.2f, %.2f", forecasts[0], forecasts[1])
}

func TestARIMAWithInsufficientData(t *testing.T) {
	// 测试数据不足的情况
	data := []float64{2.5}

	arima := NewARIMAModel(2, 1, 1)
	forecasts, err := arima.Forecast(data, 2)

	if err != nil {
		t.Errorf("数据不足时预测失败: %v", err)
		return
	}

	if len(forecasts) != 2 {
		t.Errorf("数据不足时预测结果数量错误，期望2个，实际%d个", len(forecasts))
		return
	}

	t.Logf("数据不足时预测: %.2f, %.2f", forecasts[0], forecasts[1])
}

func TestARIMAWithEmptyData(t *testing.T) {
	// 测试空数据的情况
	var data []float64

	arima := NewARIMAModel(2, 1, 1)
	_, err := arima.Forecast(data, 2)

	if err == nil {
		t.Error("空数据应该返回错误")
	}

	t.Logf("空数据测试通过，返回错误: %v", err)
}

// 性能测试
func BenchmarkARIMAForecast(b *testing.B) {
	data := make([]float64, 30)
	for i := range data {
		data[i] = 2.0 + float64(i%5)*0.2
	}

	arima := NewARIMAModel(2, 1, 1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = arima.Forecast(data, 2)
	}
}
