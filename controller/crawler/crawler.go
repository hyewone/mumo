package crawler

import (
	"context"
	"fmt"
	"log"
	"mumogo/model"
	"mumogo/service"
	"regexp"
	"strconv"
	"strings"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

type CrawlerController struct {
	Service *service.CrawlerService
}

func NewCrawlerController() *CrawlerController {
	return &CrawlerController{
		Service: service.NewCrawlerService(),
	}
}

func (con *CrawlerController) AddMovie(names []string) {
	var movies []*model.Movie
	for _, name := range names {
		movie := &model.Movie{
			Name: name,
		}
		movies = append(movies, movie)
	}

	err := con.Service.AddMovie(movies)
	if err != nil {
		// Handle error
	}
}

func (con *CrawlerController) AddStageGreeting(stageGreetingList []map[string]string) {

	var stageGreetingModels []*model.StageGreeting

	for _, data := range stageGreetingList {
		remainingSeats, err := strconv.Atoi(data["RemainingSeats"])
		if err != nil {
			// 변환 실패 시 에러 처리
		}

		movieId, err := strconv.Atoi(data["MovieID"])
		if err != nil {
			// 변환 실패 시 에러 처리
		}
		stageGreeting := &model.StageGreeting{
			ShowDate:       data["ShowDate"],
			ShowTime:       data["ShowTime"],
			RemainingSeats: remainingSeats,
			Theater:        data["Theater"],
			AttendeeName:   data["AttendeeName"],
			MovieID:        movieId,
			CinemaType:     data["CinemaType"],
		}
		stageGreetingModels = append(stageGreetingModels, stageGreeting)
	}
	// 서비스 호출
	err := con.Service.AddStageGreeting(stageGreetingModels)
	if err != nil {
		// 에러 처리
	}
}

func (con *CrawlerController) AddStageGreetingUrl(names []string, urls []string, titles []string, imgs []string, cinemaType string) {

	var urlModels []*model.StageGreetingUrl

	for i, name := range names {
		// 영화 이름으로 영화 ID 조회
		movie, err := con.Service.GetMovieByName(name)
		if err != nil {
			// Handle error
			continue
		}

		urlModel := model.NewStageGreetingUrl(int(movie.ID), cinemaType, titles[i], urls[i], imgs[i], "N")
		urlModels = append(urlModels, urlModel)
	}

	err := con.Service.AddStageGreetingUrl(urlModels)
	if err != nil {
		// Handle error
	}
}

func (con *CrawlerController) CrawlMegabox() {
	// https://www.megabox.co.kr/event/curtaincall
	// https://www.megabox.co.kr/event/detail?eventNo=13894
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	url := "https://www.megabox.co.kr/event/curtaincall"
	selector := "#event-list-wrap > div > div > ul > li > a"

	var dataNoList, movieList, imgSrcList, siteList, titleList []string
	var aNode []*cdp.Node
	var stageGreetingList []map[string]string

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector),
		chromedp.Nodes(selector, &aNode, chromedp.NodeVisible),
	)

	if err != nil {
		log.Fatal(err)
	}

	for _, node := range aNode {
		dataNo, _ := node.Attribute("data-no")
		title, _ := node.Attribute("title")
		re := regexp.MustCompile(`<([^>]+)>`)
		match := re.FindStringSubmatch(title)
		movieNm := match[1]

		dataNoList = append(dataNoList, dataNo)
		movieList = append(movieList, movieNm)
		siteList = append(siteList, "https://m.megabox.co.kr/event/detail?eventNo="+dataNo)
		titleList = append(titleList, title)

		var imgSrc string
		var ok bool
		chromedp.Run(ctx,
			chromedp.AttributeValue("a > p > img", "src", &imgSrc, &ok, chromedp.ByQuery, chromedp.FromNode(node)),
		)
		imgSrcList = append(imgSrcList, imgSrc)
	}

	con.AddMovie(movieList)
	con.AddStageGreetingUrl(movieList, siteList, titleList, imgSrcList, "MEGABOX")

	// // 각 페이지에 접속하여 추가 크롤링 진행
	for i, siteUrl := range siteList {
		movie, err := con.Service.GetMovieByName(movieList[i])
		if err != nil {
			// Handle error
			continue
		}

		dataList, err := fetchStagereetingDataMegabox(ctx, siteUrl, int(movie.ID), "MEGABOX")
		for _, data := range dataList {
			stageGreetingList = append(stageGreetingList, data)
			fmt.Println(data)
		}
	}
	con.AddStageGreeting(stageGreetingList)
}

func (con *CrawlerController) CrawlLotteCinema() {
	// https://event.lottecinema.co.kr/NLCHS/Event/DetailList?code=40
	// https://event.lottecinema.co.kr/NLCHS/Event/EventTemplateStageGreeting?eventId=401070016923136

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	url := "https://event.lottecinema.co.kr/NLCHS/Event/DetailList?code=40"
	url2 := "https://www.lottecinema.co.kr/NLCMW/Event"
	selector := "#contents > ul > li"
	selector2 := "#eventContainer > div > div > div > div.tab_con.active > div > div.tab_inner_con_wrap > div.tab_inner_con.active > ul > li"

	var movieList, siteList, titleList, imgSrcList []string
	// var dataNoList, movieList, imgSrcList, siteList []string
	var aNode, aNode2 []*cdp.Node
	var stageGreetingList []map[string]string

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector),
		chromedp.Nodes(selector, &aNode, chromedp.NodeVisible),
	)

	if err != nil {
		log.Fatal(err)
	}

	err = chromedp.Run(ctx,
		chromedp.Navigate(url2),
		chromedp.WaitVisible(selector2),
		chromedp.Nodes(selector2, &aNode2, chromedp.NodeVisible),
	)

	if err != nil {
		log.Fatal(err)
	}

	for _, node := range aNode2 {
		var onClick string
		var href string
		var ok bool
		err := chromedp.Run(ctx,
			chromedp.AttributeValue("a", "href", &href, &ok, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.AttributeValue("a", "onClick", &onClick, &ok, chromedp.ByQuery, chromedp.FromNode(node)),
		)

		if err != nil {
			log.Println("Failed to get alt attribute:", err)
			continue
		}

		fmt.Println("onClick : ", onClick)
		fmt.Println("href : ", href)
	}

	for _, node := range aNode {
		var altData string
		var href string
		var ok bool
		err := chromedp.Run(ctx,
			chromedp.AttributeValue("a > img", "alt", &altData, &ok, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.AttributeValue("a", "href", &href, &ok, chromedp.ByQuery, chromedp.FromNode(node)),
		)

		if err != nil {
			log.Println("Failed to get alt attribute:", err)
			continue
		}

		if strings.Contains(altData, "무대인사") {
			fmt.Println(altData)

		}

		// altData: <한 남자>GV(이동진평론가)
		// altData: <볼코노고프대위 탈출하다>굿즈상영회
		// altData: <치악산>회원시사회(서울/경기)
		// altData: <치악산>회원시사회(서울/경기 외)
		// altData: <한 남자>무대인사
		// altData: <달짝지근해>무대인사(2주차)
		// altData: <한 남자>GV(8/28, 건대입구)
		// altData: <타겟> 개봉주 무대인사
		// altData: <잠>중계 GV

		// var title string

		// 각 노드 클릭
		// err = chromedp.Run(ctx,
		// 	chromedp.Click(node, chromedp.NodeVisible),
		// 	chromedp.Sleep(1*time.Second), // 적절한 대기 시간을 설정
		// 	chromedp.Title(&title),
		// )
		// if err != nil {
		// 	log.Println("Failed to click node:", err)
		// 	continue
		// }
		// fmt.Println("Navigated Title:", title)

		// // 이전 페이지로 돌아가기
		// err = chromedp.Run(ctx,
		// 	chromedp.NavigateBack(),
		// )
		// if err != nil {
		// 	log.Println("Failed to navigate back:", err)
		// }

		// altData, _ := node.Attribute("alt")

		// var nodeContent string
		// err := chromedp.Run(ctx,
		// 	chromedp.TextContent(node, &nodeContent),
		// )
		// if err != nil {
		// 	log.Println("Failed to get node content:", err)
		// 	continue
		// }

		// fmt.Println("Node Content:", nodeContent)

		// var altData string
		// var allData string
		// chromedp.Run(ctx,
		// 	chromedp.Text(" > img", "alt", &altData, nil, chromedp.FromNode(node)),
		// 	// chromedp.AttributesAll(node, &allData),
		// )

		// fmt.Println("altData: ", altData)
		// fmt.Println("allData: ", allData)
		// fmt.Println("node: ", node)

		// if strings.Contains(altData, "무대인사") {
		// 	fmt.Println("무대인사를 포함")

		// 	href, _ := node.Attribute("href")
		// 	fmt.Println("href: ", href)

		// 무대인사가 포함되면 제목도 따서 DB에 저장하고, 상세화면 들어가서 작업해줘야함

		// movie, err := con.Service.GetMovieByName(movieList[i])
		// if err != nil {
		// 	// Handle error
		// 	continue
		// }

		// dataCtx, dataCancel := chromedp.NewContext(ctx)

		// err := chromedp.Run(dataCtx,
		// 	chromedp.Click(node, chromedp.NodeVisible),
		// )
		// if err != nil {
		// 	log.Println("Failed to click component:", err)
		// 	continue
		// }

		// dataList, err := fetchStagereetingDataLotteCinema(dataCtx, "LOTTECINEMA")
		// for _, data := range dataList {
		// 	stageGreetingList = append(stageGreetingList, data)
		// 	fmt.Println(data)
		// }
		// defer dataCancel()
		// }

		// title, _ := node.Attribute("title")
		// re := regexp.MustCompile(`<([^>]+)>`)
		// match := re.FindStringSubmatch(title)
		// title = match[1]

		// dataNoList = append(dataNoList, dataNo)
		// movieList = append(movieList, title)
		// siteList = append(siteList, "https://m.megabox.co.kr/event/detail?eventNo="+dataNo)

		// var imgSrc string
		// chromedp.Run(ctx,
		// 	chromedp.AttributeValue("p.img > img", "src", &imgSrc, nil),
		// )
		// imgSrcList = append(imgSrcList, imgSrc)
	}

	con.AddMovie(movieList)
	con.AddStageGreetingUrl(movieList, siteList, titleList, imgSrcList, "LOTTECINEMA")
	con.AddStageGreeting(stageGreetingList)
}

func (con *CrawlerController) CrawlCgv() {
	// #contents > div.cols-content > div.col-detail.event > ul > li
	// http://www.cgv.co.kr/culture-event/event/defaultNew.aspx?mCode=004#1
	// http://www.cgv.co.kr/culture-event/event/detailViewUnited.aspx?seq=36631&menu=004
	// /<a id="tile_0" href="./detailViewUnited.aspx?seq=38137&amp;menu=001"><div class="evt-thumb"><img src="https://img.cgv.co.kr/WebApp/contents/eventV4/38137/16927756832350.jpg" alt="[해리포터와 혼혈왕자]
	// 4DX 포스터"></div><div class="evt-desc"><p class="txt1">[해리포터와 혼혈왕자]
	// 4DX 포스터</p><p class="txt2">2023.08.23~2023.09.12</p></div></a>

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	url := "http://www.cgv.co.kr/search/?query=%uBB34%uB300%uC778%uC0AC"
	selector := "#ctl00_PlaceHolderContent_divSearchEvent > ul > li"

	var movieList, siteList, titleList, imgSrcList []string
	// var dataNoList, movieList, imgSrcList, siteList []string
	var aNode []*cdp.Node
	// var stageGreetingList []map[string]string

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector),
		chromedp.Nodes(selector, &aNode, chromedp.NodeVisible),
	)

	if err != nil {
		log.Fatal(err)
	}

	for _, node := range aNode {
		var pData, href, imgSrc string
		var ok bool
		err := chromedp.Run(ctx,
			chromedp.Text("div > strong", &pData, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.AttributeValue("a", "href", &href, &ok, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.AttributeValue("a > img", "src", &imgSrc, &ok, chromedp.ByQuery, chromedp.FromNode(node)),
		)

		if err != nil {
			log.Println("Failed to get alt attribute:", err)
			continue
		}

		re := regexp.MustCompile(`\[(.*?)\]`)
		match := re.FindStringSubmatch(pData)
		movieNm := match[1]

		movieList = append(movieList, movieNm)
		siteList = append(siteList, "http://www.cgv.co.kr"+href)
		titleList = append(titleList, pData)
		imgSrcList = append(imgSrcList, imgSrc)
	}

	fmt.Println(movieList)
	fmt.Println(siteList)

	con.AddMovie(movieList)
	con.AddStageGreetingUrl(movieList, siteList, titleList, imgSrcList, "CGV")

	// Vision API
	// ctx := context.Background()

	// // 이미지 URL
	// file := "./images/target.jpeg"

	// client, err := vision.NewImageAnnotatorClient(ctx, option.WithCredentialsFile("./singular-rope-396203-c379014d95b9.json"))
	// if err != nil {
	// 	log.Fatalf("Failed to create Vision client: %v", err)
	// }

	// f, err := os.Open(file)
	// if err != nil {
	// 	log.Fatalf("Failed to open image: %v", err)
	// }
	// defer f.Close()

	// image, err := vision.NewImageFromReader(f)
	// if err != nil {
	// 	//    return err
	// }
	// annotations, err := client.DetectTexts(ctx, image, nil, 10)
	// if err != nil {

	// }

	// if len(annotations) == 0 {
	// 	fmt.Println("No text found.")
	// } else {
	// 	fmt.Println("Text :")
	// 	for _, annotation := range annotations {
	// 		fmt.Println(annotation.Description)
	// 	}
	// }
}

func fetchStagereetingDataMegabox(ctx context.Context, url string, id int, cinemaType string) ([]map[string]string, error) {
	// url := "https://www.megabox.co.kr/event/curtaincall"
	// body > div.container.event-detail-cont > div.event-sec.type0401 > div > div

	var ticketingBoxes []*cdp.Node
	selector := "body > div.container.event-detail-cont > div.event-sec.type0401 > div > div"

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector),
		chromedp.Nodes(selector, &ticketingBoxes),
	)
	if err != nil {
		log.Fatal(err)
	}

	var dataList []map[string]string
	var showDate, showTime, remainingSeats, theater, attendeeName string

	fmt.Println()

	for _, box := range ticketingBoxes {
		chromedp.Run(ctx,
			chromedp.Text("li:nth-child(1)", &showDate, chromedp.ByQuery, chromedp.FromNode(box)),
			chromedp.Text("li:nth-child(2)", &showTime, chromedp.ByQuery, chromedp.FromNode(box)),
			chromedp.Text("li:nth-child(3)", &remainingSeats, chromedp.ByQuery, chromedp.FromNode(box)),
			chromedp.Text("li:nth-child(4)", &theater, chromedp.ByQuery, chromedp.FromNode(box)),
			chromedp.Text("li:nth-child(5)", &attendeeName, chromedp.ByQuery, chromedp.FromNode(box)),
		)

		showDate = strings.TrimPrefix(showDate, "상영날짜 ")
		showTime = strings.TrimPrefix(showTime, "상영시간 ")
		remainingSeats = strings.TrimPrefix(remainingSeats, "잔여좌석 ")
		theater = strings.TrimPrefix(theater, "상영극장 ")
		attendeeName = strings.TrimPrefix(attendeeName, "참석자명 ")

		data := map[string]string{
			"ShowDate":       strings.TrimSpace(showDate),
			"ShowTime":       strings.TrimSpace(showTime),
			"RemainingSeats": strings.TrimSpace(remainingSeats),
			"Theater":        strings.TrimSpace(theater),
			"AttendeeName":   strings.TrimSpace(attendeeName),
			"MovieID":        strconv.Itoa(id),
			"CinemaType":     cinemaType,
		}
		dataList = append(dataList, data)
	}
	return dataList, nil
}

func fetchStagereetingDataLotteCinema(ctx context.Context, cinemaType string) ([]map[string]string, error) {

	fmt.Println("fetchStagereetingDataLotteCinema")
	var ticketingBoxes []*cdp.Node
	selector := "#eventContainer > div > div:nth-child(2) > div.preview_list > div:nth-child(1)"

	err := chromedp.Run(ctx,
		chromedp.WaitVisible(selector),
		chromedp.Nodes(selector, &ticketingBoxes),
	)
	if err != nil {
		log.Fatal(err)
	}

	var dataList []map[string]string
	var showDate, showTime, theater, attendeeName string

	for _, box := range ticketingBoxes {
		chromedp.Run(ctx,
			chromedp.Text("dl:nth-child(1) > dd", &showDate, chromedp.ByQuery, chromedp.FromNode(box)),
			chromedp.Text("dl:nth-child(2) > dd", &showTime, chromedp.ByQuery, chromedp.FromNode(box)),
			// chromedp.Text("dl:nth-child(3) > dd", &remainingSeats, chromedp.ByQuery, chromedp.FromNode(box)),
			chromedp.Text("dl:nth-child(3) > dd", &theater, chromedp.ByQuery, chromedp.FromNode(box)),
			chromedp.Text("dl:nth-child(4) > dd", &attendeeName, chromedp.ByQuery, chromedp.FromNode(box)),
		)

		data := map[string]string{
			"ShowDate": strings.TrimSpace(showDate),
			"ShowTime": strings.TrimSpace(showTime),
			// "RemainingSeats": strings.TrimSpace(remainingSeats),
			"Theater":      strings.TrimSpace(theater),
			"AttendeeName": strings.TrimSpace(attendeeName),
			"MovieID":      strconv.Itoa(1), //strconv.Itoa(id),
			"CinemaType":   cinemaType,
		}
		dataList = append(dataList, data)
	}
	return dataList, nil
}

func (con *CrawlerController) CrawlTest() {

	// chromeDriver := webdriver.NewChromeDriver("./chromedriver_mac_arm64/chromedriver.exe")
	// err := chromeDriver.Start()
	// if err != nil {
	// 	log.Println(err)
	// }
	// desired := webdriver.Capabilities{"Platform": "Windows"}
	// required := webdriver.Capabilities{}
	// session, err := chromeDriver.NewSession(desired, required)
	// if err != nil {
	// 	log.Println(err)
	// }
	// err = session.Url("http://golang.org")
	// if err != nil {
	// 	log.Println(err)
	// }

	// // 웹 페이지 열기
	// if err := session.GetUrl("https://event.lottecinema.co.kr/NLCHS/Event/DetailList?code=40"); err != nil {
	// 	log.Fatal(err)
	// }

	// // 개발자 도구 패킷 캡처 시작
	// script := `
	// 	window.performance.clearResourceTimings();
	// 	return "Capture started";
	// `
	// _, err = wd.ExecuteScript(script, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // 작업 수행 후 일정 시간 대기 (예: 5초)
	// time.Sleep(5 * time.Second)

	// // 네트워크 패킷 정보 가져오기
	// script = `
	// 	return window.performance.getEntriesByType("resource");
	// `
	// packets, err := wd.ExecuteScript(script, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // 패킷 정보 출력
	// for _, packet := range packets.([]interface{}) {
	// 	fmt.Println(packet)
	// }
}

// func scrapIt(url string, str *string) chromedp.Tasks {
// 	return chromedp.Tasks{
// 		chromedp.Navigate(url),
// 		chromedp.ActionFunc(func(ctx context.Context) error {
// 			node, err := dom.GetDocument().Do(ctx)
// 			if err != nil {
// 				return err
// 			}
// 			*str, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
// 			return err
// 		}),
// 	}
// }
