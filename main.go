package main

import (
	"flag"
	"fmt"
	"goStudy/ginutils"
	"goStudy/gormodel"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

var (
	//\d代表数字
	reQQEmail = `(\d+)@qq.com`
	//匹配邮箱
	reEmail = `\w+@\w+\.\w+(\.\w+)?`
	//链接
	reLink  = `href="(https?://[\s\S]+?)"`
	rePhone = `1[3456789]\d\s?\d{4}\s?\d{4}`
	//410222 1987 06 13 4038
	reIdcard = `[12345678]\d{5}((19\d{2})|(20[01]))((0[1-9]|[1[012]]))((0[1-9])|[12]\d|[3[01]])\d{3}[\dXx]`
	reImg    = `"(https?://[^"]+?(\.((jpg)|(jpeg)|(png)|(gif)|(ico))))"`
	dgraph   = flag.String("d", "127.0.0.1:9080", "Dgraph server address")
)

func main() {
	//fmt.Printf("hahah-gomod",congigor.Config{})
	//name := route.GetName()
	//fmt.Print(name)

	//创建无缓冲通道
	/*c := make(chan int, 3)
	//长度和容量
	fmt.Printf("len(c)=%d, cap(c)=%d\n", len(c), cap(c))
	//子协程存数据
	go func() {
		defer fmt.Println("子协程结束")
		//向通道添加数据
		for i := 0; i < 3; i++ {
			c <- i
			fmt.Printf("子协程正在运行[%d]:len(c)=%d,cap(c)=%d\n", i, len(c), cap(c))
		}
	}()

	time.Sleep(2 * time.Second)
	//主协程取数据
	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("num=", num)
	}
	fmt.Println("主协程结束")*/
	/*for {
		timer1 := time.NewTimer(2 * time.Second)
		//打印当前时间
		//t1 := time.Now()
		//fmt.Printf("t1:%v\n", t1)

		//从管道中取出C打印

		t2 := <-timer1.C
		fmt.Printf("t2:%v\n", t2)
	}*/

	//getEmail()

	//------dgraph
	/*dgraph := schema.NewDgrapClient()

	p1 := structV2.Person{
		Name:     "wanghaha",
		Age:      190,
		From:     "China",
		NameOFen: "wanghaha",
		NameOFcn: "王哈哈",
		NameOFjp: "王ハ",
	}
	p2 := structV2.Person{
		Name:     "chenchao",
		Age:      22,
		From:     "China",
		NameOFen: "ChaoChen",
		NameOFcn: "陈超",
	}
	p3 := structV2.Person{
		Name:     "xhe",
		Age:      18,
		From:     "Japan",
		NameOFen: "wanghe",
		NameOFcn: "x鹤",
	}
	p4 := structV2.Person{
		Name:     "changyang",
		Age:      19,
		From:     "England",
		NameOFcn: "常飏",
	}
	p5 := structV2.Person{
		Name:     "yetao",
		Age:      18,
		From:     "Russian",
		NameOFen: "TaoYe",
		NameOFcn: "叶掏",
	}
	op := &api.Operation{}
	op.Schema = `
		name: string .
		age: int .
		from: string .
		nameOFcn: string @index(term) .
		nameOFjp: string @index(term) .
		nameOFen: string @index(term) .
	`
	ctx := context.Background()

	if err := dgraph.Alter(ctx, op); err != nil {
		log.Fatal(err)
	}

	mu := &api.Mutation{
		CommitNow: true,
	}

	var p = [5]structV2.Person{p1, p2, p3, p4, p5}

	for _, x := range p {
		pb, err := json.Marshal(x)
		if err != nil {
			log.Println(err)
		}
		mu.SetJson = pb
		_, err = dgraph.NewTxn().Mutate(ctx, mu)
		if err != nil {
			log.Println(err)
		}
	}*/

	//-----gin
	err := gormodel.Init()
	if err != nil {
		log.Println(err)
		return
	}

	r := ginutils.InitGin()
	r.Run()
}

func getEmail() {
	resp, err := http.Get("https://www.cnblogs.com/konghui/p/10742153.html")
	HandleError(err, "http.getUrl")
	defer resp.Body.Close()
	//接收页面
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	//打印页面内容
	pageStr := string(pageBytes)
	fmt.Println(pageStr)

	//2.捕获邮箱，先搞定qq邮箱
	//传入正则
	re := regexp.MustCompile(reQQEmail)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		//fmt.Println(result)
		fmt.Printf("email=%s qq=%s\n", result[0], result[1])
	}
}

//处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}
