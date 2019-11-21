package matchers
//包名
import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
//标准库 和search包
	"github.com/goinaction/code/chapter2/sample/search"
)

type (
	// item defines the fields associated with the item tag
	//item 根据item字段的标签，将定义的字段与rss文档字段关联起来
	// in the rss document.
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	// image defines the fields associated with the image tag
	// in the rss document.
	//根据image字段标签将定义的字段与rss文档字段关联起来
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// channel defines the fields associated with the channel tag
	// in the rss document.
	//根据channel 字段标签，将定义的字段与rss文档的字段关联起来
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	// rssDocument defines the fields associated with the rss document.
	//rssDocument定义了与rss文档关联的字段
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

// rssMatcher implements the Matcher interface.
// rssMatcher实现了Matcher 接口
type rssMatcher struct{}

// init registers the matcher with the program. 匹配器注册到程序里
func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

// Search looks at the document for the specified search term.
func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	var results []*search.Result//声明一个值为nil的切片
	//每个切片都指向Result类型值的指针

	log.Printf("Search Feed Type[%s] Site[%s] For URI[%s]\n", feed.Type, feed.Name, feed.URI)

	// Retrieve the data to search.
	document, err := m.retrieve(feed)
	if err != nil {
		return nil, err
	}

	for _, channelItem := range document.Channel.Item {
		//依次访问内部每一项
		// Check the title for the search term.
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		使用regexp包里的MatchString函数对channelItem里的title进行搜索 匹配
		if err != nil {
			return nil, err
		}

		// If we found a match save the result.
		if matched {//如果值为真
			//使用append 将结果加倒入results切片
			//append内置函数会根据切片需要决定是否增加切片的长度和容量
			results = append(results, &search.Result{
				Field:   "Title",
				Content: channelItem.Title,
			})//加入到切片里的是指向Results类型值的指针，取地址用& 最终将指针存入切片
		}

		// Check the description for the search term.
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}

		// If we found a match save the result.
		if matched {
			results = append(results, &search.Result{
				Field:   "Description",
				Content: channelItem.Description,
			})
		}
	}

	return results, nil
}

// retrieve performs a HTTP Get request for the rss feed and decodes the results.
//get请求rss数据源并解码
//小写名字 没有对外暴露
func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("No rss feed uri provided")
	}

	// Retrieve the rss feed document from the web.
	//从网络获取rss数据源文档
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}//http.Get返回一个指向Response 类型的指针

	// Close the response once we return from the function.
	//一旦函数返回 关闭返回的响应链接
	defer resp.Body.Close()

	// Check the status code for a 200 so we know we have received a
	// proper response. 检查状态码200 是不是收到了正确响应
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
	}

	// Decode the rss feed document into our struct type.
	//将rss数据源文档解码到我们定义的结构类型里 不需要检查错误调用者会做
	// We don't need to check for errors, the caller can do this.
	var document rssDocument //document 类型的rssDocument
	//传入局部变量document 地址并返回
	err = xml.NewDecoder(resp.Body).Decode(&document)
	//NewDecoder 函数返回一个指向Decoder值的指针，之后调用指针的Decode方法
	return &document, err
}
