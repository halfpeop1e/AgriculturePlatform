<template>
  <div class="min-h-screen bg-[#f1f5f9] p-4 md:p-8 font-sans text-slate-900">
    <div class="max-w-7xl mx-auto">
      
      <!-- 头部：统计卡片 -->
      <div class="grid grid-cols-4 gap-4 mb-6">
        <div v-for="cat in apiData.categories" :key="cat.name" 
             class="bg-white p-4 rounded-2xl shadow-sm border-b-4" :style="{ borderColor: cat.color }">
          <div class="text-slate-400 text-xs font-bold uppercase tracking-wider">{{ cat.name }}类均价</div>
          <div class="flex items-baseline justify-between mt-1">
            <span class="text-xl font-black">¥{{ cat.data[cat.data.length - 1] }}</span>
            <span class="text-[10px] px-1.5 py-0.5 rounded bg-slate-100 font-bold" :style="{ color: cat.color }">成交价格</span>
          </div>
        </div>
      </div>

      <!-- 主看板 -->
      <div class="bg-white rounded-3xl shadow-xl shadow-slate-200/50 overflow-hidden border border-white">
        
        <!-- 工具栏 -->
        <div class="p-6 border-b border-slate-50 flex flex-wrap justify-between items-center gap-6">
          <div>
            <h1 class="text-2xl font-black italic tracking-tighter text-slate-800">木吉卡BUFF</h1>
            <p class="text-slate-400 text-sm">农产品每日价格实时追溯系统</p>
          </div>
          
          <div class="flex items-center bg-slate-100 p-1 rounded-xl">
            <button 
              v-for="mode in ['category', 'product']" 
              :key="mode"
              @click="()=>{viewMode = mode
                console.log(viewMode)
              }"
              :class="viewMode === mode ? 'bg-white shadow-sm text-blue-600' : 'text-slate-500'"
              class="px-6 py-2 rounded-lg text-sm font-bold transition-all"
            >
              {{ mode === 'category' ? '大类走势' : '单品详情' }}
            </button>
          </div>
        </div>

        <!-- 图表容器 -->
        <div class="relative p-4">
          <!-- 单品模式：多个独立图表 -->
          <div v-if="viewMode === 'product'" class="space-y-6">
            <div v-for="(product, index) in apiData.products" :key="product.name" 
                 class="bg-white rounded-2xl shadow-sm border border-slate-200 overflow-hidden">
              <!-- 产品标题栏 -->
              <div class="p-4 border-b border-slate-100 flex items-center justify-between">
                <div class="flex items-center gap-3">
                  <div class="w-3 h-3 rounded-full" :style="{ backgroundColor: getProductColor(index) }"></div>
                  <h3 class="font-bold text-slate-800">{{ product.name }}</h3>
                </div>
                <div class="flex items-center gap-3">
                  <span class="text-2xl font-black text-slate-800">¥ {{ product.price }}</span>
                  <span :class="product.change >= 0 ? 'text-emerald-500' : 'text-rose-500'" class="text-sm font-bold">
                    {{ product.change >= 0 ? '▲' : '▼' }} {{ Math.abs(product.change) }}%
                  </span>
                </div>
              </div>
              <!-- 单个产品图表 -->
              <div :ref="(el) => { if (el) productCharts[index] = el; }" class="h-[250px] w-full p-4"></div>
            </div>
          </div>
          
          <!-- 大类模式：单个图表 -->
          <div v-else ref="chartRef" class="h-[500px] w-full"></div>
        </div>
      </div>

      <!-- 底部状态 --> 
      <div class="mt-4 flex justify-between items-center px-4">
        <div class="flex items-center gap-2 text-slate-400 text-xs">
          <span class="relative flex h-2 w-2">
            <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-blue-400 opacity-75"></span>
            <span class="relative inline-flex rounded-full h-2 w-2 bg-blue-500"></span>
          </span>
          数据已根据后端更新: {{ lastUpdateTime }}
        </div>
        <el-button link type="primary" size="small" @click="fetchMarketData">手动刷新数据</el-button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, reactive, watch, onUnmounted, nextTick } from 'vue';
import * as echarts from 'echarts';
import dayjs from 'dayjs';

// --- 状态管理 ---
const chartRef = ref(null);
const viewMode = ref('category'); // 'category' 或 'product'
const lastUpdateTime = ref(dayjs().format('HH:mm:ss'));
let myChart = null;
const productCharts = ref([]);
let timer = null;
definePageMeta({ layout: 'home-page-layout' })
// 后端集中返回的数据结构
const apiData = reactive({
  dates: [],
  categories: [],
  products: []
});

// --- ECharts 配置  ---
const updateChart = () => {
  const isProduct = viewMode.value === 'product';
  
  if (isProduct) {
    // 单品模式：为每个产品创建独立的图表
    updateProductCharts();
  } else {
    // 大类模式：使用原来的单个图表
    nextTick(() => {
      if (!chartRef.value) return;
      
      // 如果图表已存在，先销毁
      if (myChart) {
        myChart.dispose();
      }
      
      // 重新初始化图表
      myChart = echarts.init(chartRef.value);
    
      const series = apiData.categories.map(cat => ({
        name: cat.name,
        type: 'line',
        data: cat.data,
        smooth: 0.4,
        showSymbol: false,
        lineStyle: { width: 3, color: cat.color },
        emphasis: { lineStyle: { width: 5 } }
      }));

      const option = {
        animationDuration: 1000,
        legend: {
          data: apiData.categories.map(cat => ({
            name: cat.name,
            itemStyle: { color: cat.color }
          })),
          top: '2%',
          left: 'center',
          itemWidth: 15,
          itemHeight: 15,
          textStyle: {
            color: '#64748b',
            fontSize: 12,
            fontWeight: 'bold'
          }
        },
        tooltip: {
          trigger: 'axis',
          backgroundColor: 'rgba(15, 23, 42, 0.9)',
          borderWidth: 0,
          padding: 15,
          textStyle: { color: '#fff' },
          axisPointer: {
            type: 'cross',
            lineStyle: { color: '#cbd5e1', type: 'dashed' }
          },
          formatter: (params) => {
            let str = `<div class="mb-2 font-bold text-slate-400">${params[0].name} 报价</div>`;
            params.forEach(p => {
              // 获取线条的实际颜色
              const lineColor = apiData.categories.find(cat => cat.name === p.seriesName)?.color || p.color;
              str += `<div class="flex justify-between gap-8 items-center mt-1">
                        <span class="flex items-center gap-2">
                          <div class="w-2 h-2 rounded-full" style="background:${lineColor}"></div>
                          ${p.seriesName}
                        </span>
                        <span class="font-mono font-bold">¥${p.value}</span>
                      </div>`;
            });
            return str;
          }
        },
        grid: { left: '4%', right: '4%', top: '20%', bottom: '10%', containLabel: true },
        xAxis: {
          type: 'category',
          data: apiData.dates,
          boundaryGap: false,
          axisLine: { show: false },
          axisTick: { show: false },
          axisLabel: { color: '#94a3b8', margin: 20 }
        },
        yAxis: {
          type: 'value',
          position: 'left', 
          splitLine: { lineStyle: { color: '#f1f5f9', type: 'dashed' } },
          axisLabel: { color: '#94a3b8' }
        },
        series: series
      };

      myChart.setOption(option, true);
    });
  }
};

// 更新单品图表
const updateProductCharts = () => {
  console.log('Updating product charts. Products count:', apiData.products.length);
  console.log('Product charts refs:', productCharts.value.length);
  
  apiData.products.forEach((product, index) => {
    const chartElement = productCharts.value[index];
    if (!chartElement) {
      console.log(`Chart element not found for index ${index}`, productCharts.value);
      return;
    }
    
    let chart = echarts.getInstanceByDom(chartElement);
    if (!chart) {
      chart = echarts.init(chartElement);
    }

    const option = {
      animationDuration: 1000,
      tooltip: {
        trigger: 'axis',
        backgroundColor: 'rgba(15, 23, 42, 0.9)',
        borderWidth: 0,
        padding: 15,
        textStyle: { color: '#fff' },
        formatter: (params) => {
          return `<div class="font-mono font-bold">¥${params[0].value}</div>`;
        }
      },
      grid: { 
        left: '8%', 
        right: '4%', 
        top: '10%', 
        bottom: '15%', 
        containLabel: true 
      },
      xAxis: {
        type: 'category',
        data: apiData.dates,
        boundaryGap: false,
        axisLine: { show: false },
        axisTick: { show: false },
        axisLabel: { color: '#94a3b8', fontSize: 10 }
      },
      yAxis: {
        type: 'value',
        position: 'right', 
        splitLine: { lineStyle: { color: '#f1f5f9', type: 'dashed' } },
        axisLabel: { color: '#94a3b8', fontSize: 10 }
      },
      series: [{
        name: product.name,
        type: 'line',
        data: product.data,
        smooth: 0.4,
        showSymbol: false,
        lineStyle: { width: 3, color: getProductColor(index) },
        areaStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            { offset: 0, color: hexToRgba(getProductColor(index), 0.1) },
            { offset: 1, color: hexToRgba(getProductColor(index), 0) }
          ])
        }
      }]
    };

    chart.setOption(option, true);
  });
};

// --- 数据获取 (对接后端) ---
const fetchMarketData = async () => {
  try {
    // 这里模拟后端返回的数据更新
    const mockDates = ['12-20', '12-21', '12-22', '12-23', '12-24', '12-25'];

    apiData.dates = mockDates;
    apiData.categories = [
      { name: "粮食", data: randomData(2.5, 6), color: "#f59e0b" },
      { name: "蔬菜", data: randomData(3.5, 6), color: "#10b981" },
      { name: "水果", data: randomData(8.5, 6), color: "#ef4444" },
      { name: "种苗", data: randomData(1.5, 6), color: "#8b5cf6" }
    ];
    apiData.products = [
      {
        name: "红富士苹果 (特级)",
        price: (14 + Math.random()).toFixed(2),
        change: (Math.random() * 4 - 2).toFixed(1),
        data: randomData(14, 6)
      },
      {
        name: "有机西红柿",
        price: (8 + Math.random()).toFixed(2),
        change: (Math.random() * 4 - 2).toFixed(1),
        data: randomData(8, 6)
      },
      {
        name: "有机胡萝卜",
        price: (5 + Math.random()).toFixed(2),
        change: (Math.random() * 4 - 2).toFixed(1),
        data: randomData(5, 6)
      },
      {
        name: "优质黄瓜",
        price: (6 + Math.random()).toFixed(2),
        change: (Math.random() * 4 - 2).toFixed(1),
        data: randomData(6, 6)
      },
      {
        name: "优质黄瓜1",
        price: (6 + Math.random()).toFixed(2),
        change: (Math.random() * 4 - 2).toFixed(1),
        data: randomData(6, 6)
      },
      {
        name: "有机胡萝卜2",
        price: (5 + Math.random()).toFixed(2),
        change: (Math.random() * 4 - 2).toFixed(1),
        data: randomData(5, 6)
      },
      {
        name: "有机胡萝卜3",
        price: (5 + Math.random()).toFixed(2),
        change: (Math.random() * 4 - 2).toFixed(1),
        data: randomData(5, 6)
      },
    ];

    const newData = await getOrderDate();
    console.log('Fetched data from backend:', newData);
    if(newData){
      apiData.dates = newData.dates;
      apiData.categories = newData.categories;
      apiData.products = newData.products;
    }
    
    lastUpdateTime.value = dayjs().format('HH:mm:ss');
    updateChart();
  } catch (e) {
    console.error("数据加载失败", e);
  }
};

// 辅助：生成随机数模拟更新效果
const randomData = (base, len) => Array.from({length: len}, () => (base + Math.random()).toFixed(2));

// 辅助函数：为每个产品生成固定的随机颜色
const getProductColor = (index) => {
  // 使用种子随机数，确保每个产品的颜色在刷新后保持一致
  const colors = [
    '#3b82f6', '#ef4444', '#10b981', '#f59e0b', '#8b5cf6', 
    '#ec4899', '#14b8a6', '#f97316', '#84cc16', '#06b6d4',
    '#6366f1', '#a855f7', '#f43f5e', '#22c55e', '#eab308'
  ];
  
  // 如果索引在颜色数组范围内，使用预设颜色
  if (index < colors.length) {
    return colors[index];
  }
  
  // 如果超出范围，生成随机颜色但基于索引保持一致性
  let seed = index * 9301 + 49297;
  seed = (seed * 9301 + 49297) % 233280;
  const hue = (seed / 233280) * 360;
  return `hsl(${hue}, 70%, 50%)`;
};

// 辅助函数：将十六进制或HSL颜色转换为RGBA
const hexToRgba = (color, alpha) => {
  if (color.startsWith('hsl')) {
    // 处理 HSL 颜色
    const matches = color.match(/hsl\((\d+),\s*(\d+)%,\s*(\d+)%\)/);
    if (matches) {
      const h = parseInt(matches[1]) / 360;
      const s = parseInt(matches[2]) / 100;
      const l = parseInt(matches[3]) / 100;
      
      // HSL 转 RGB
      const hue2rgb = (p, q, t) => {
        if (t < 0) t += 1;
        if (t > 1) t -= 1;
        if (t < 1/6) return p + (q - p) * 6 * t;
        if (t < 1/2) return q;
        if (t < 2/3) return p + (q - p) * (2/3 - t) * 6;
        return p;
      };
      
      const q = l < 0.5 ? l * (1 + s) : l + s - l * s;
      const p = 2 * l - q;
      const r = Math.round(hue2rgb(p, q, h + 1/3) * 255);
      const g = Math.round(hue2rgb(p, q, h) * 255);
      const b = Math.round(hue2rgb(p, q, h - 1/3) * 255);
      
      return `rgba(${r}, ${g}, ${b}, ${alpha})`;
    }
  }
  
  // 处理十六进制颜色
  const hex = color.replace('#', '');
  const r = parseInt(hex.substr(0, 2), 16);
  const g = parseInt(hex.substr(2, 2), 16);
  const b = parseInt(hex.substr(4, 2), 16);
  return `rgba(${r}, ${g}, ${b}, ${alpha})`;
};

// --- 生命周期 ---
onMounted(() => {
  fetchMarketData();
});


// 监听模式切换
watch(viewMode, (newMode) => {
    nextTick(() => {
      updateChart();
    });
});

// 监听产品数据变化
watch(() => apiData.products, () => {
  nextTick(() => {
      updateChart();
    });
}, { deep: true });
</script>

<style scoped>
/* 隐藏 Element Plus 默认的一些边框 */
:deep(.el-button) {
  font-weight: bold;
}
</style>