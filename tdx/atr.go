package tdx

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const historyDailyQuotationPath = `D:\new_tdx\vipdoc`

type dayData struct {
	Date                   int32
	Open, High, Low, Close int32
	Amount, Qty            int32
	Other                  int32
}

// To 转换为日线行情数据
func (d *dayData) To() *DayQuotation {
	return &DayQuotation{
		Date:   d.Date,
		Open:   float32(d.Open) / 100,
		High:   float32(d.High) / 100,
		Low:    float32(d.Low) / 100,
		Close:  float32(d.Close) / 100,
		Amount: float32(d.Amount),
		Qty:    d.Qty,
	}
}

// DayQuotation 日线行情
type DayQuotation struct {
	Date   int32   //日期
	Open   float32 //开盘价
	High   float32 //最高价
	Low    float32 //最低价
	Close  float32 //收盘价
	Amount float32 //总成交金额
	Qty    int32   //  总成交量
}

func (day DayQuotation) toString() string {
	return "日期" + strconv.Itoa(int(day.Date)) +
		"开盘价(元):" + strconv.FormatFloat(float64(day.Open), 'f', 2, 32) +
		"最高价(元):" + strconv.FormatFloat(float64(day.High), 'f', 2, 32) +
		"最低价(元):" + strconv.FormatFloat(float64(day.Low), 'f', 2, 32) +
		"收盘价(元):" + strconv.FormatFloat(float64(day.Close), 'f', 2, 32) +
		"成交金额(元):" + strconv.FormatFloat(float64(day.Amount), 'f', 2, 32) +
		"成交量:" + strconv.FormatFloat(float64(day.Qty), 'f', 2, 32)

}

//GetStockQuoation 获取股票历史行情
func getStockQuoation(marker, stockCode string, n int) ([]*DayQuotation, error) {
	marker = strings.ToLower(strings.TrimSpace(marker))
	stockCode = strings.ToLower(strings.TrimSpace(stockCode))
	if marker == "" || stockCode == "" {
		return nil, errors.New("marker和stockCode不能为空")
	}
	//文件路径，e.g. D:\new_tdx\vipdoc\sz\lday\sz000002.day
	name := filepath.Join(historyDailyQuotationPath, marker, "lday", fmt.Sprintf("%s%s.day", marker, stockCode))
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	quos := []*DayQuotation{}
	oneDay := make([]byte, 32)
	data := &dayData{}
	for {
		l, err := f.Read(oneDay)
		if err == io.EOF {
			break
		} else if err != nil {
			return quos, err
		} else if l != 32 {
			return quos, errors.New("数据不完整")
		}
		buf := bytes.NewBuffer(oneDay)
		err = binary.Read(buf, binary.LittleEndian, data)
		if err != nil {
			return quos, err
		}
		quos = append(quos, data.To())
	}

	return getTdxArr(quos, n), nil
}

//数组倒序并且获取指点条数
func getTdxArr(arr []*DayQuotation, n int) []*DayQuotation {
	var temp *DayQuotation
	length := len(arr)
	for i := 0; i < length/2; i++ {
		temp = (arr)[i]
		(arr)[i] = (arr)[length-1-i]
		(arr)[length-1-i] = temp
	}

	resArr := []*DayQuotation{}
	for i := 0; i < n; i++ {
		resArr = append(resArr, arr[i])
	}
	return resArr
}

//计算atr
func CalculationAtr(code string, n int) (float32, []string, error) {
	quos, err := getStockQuoation("sz", code, n)
	if err != nil {
		return 0, nil, err
	}
	//def get_atr(df):
	//df['yesterday_colse'] = df['close'].shift(1)
	//df['1'] = df['high'] - df['low']
	//df['2'] = df['high'] - df['yesterday_colse']
	//df['3'] = df['yesterday_colse'] - df['low']
	//
	//temp_df = df[['1', '2', '3']]
	//temp_df = temp_df.iloc[1:, :]
	//TR = temp_df.max(axis=1)
	//
	//ATR = np.around(np.mean(TR), 0)
	//return ATR
	//
	//Open   float32 //开盘价
	//High   float32 //最高价
	//Low    float32 //最低价

	var totalTr float32
	dataArr := []string{}
	for i := 0; i < len(quos)-1; i++ {
		temp := quos[i]            //当天的数据指标
		yesterdayTemp := quos[i+1] //昨天的数据指标
		//组装计算的数据，打印到页面
		dataArr = append(dataArr, quos[i].toString())
		totalTr += func(vals ...float32) float32 {
			var max float32
			for _, val := range vals {
				if val > max {
					max = val
				}
			}
			return max
		}(temp.High-temp.Low, temp.High-yesterdayTemp.Close, yesterdayTemp.Close-yesterdayTemp.Low)
	}
	return totalTr / float32(n), dataArr, err
}
