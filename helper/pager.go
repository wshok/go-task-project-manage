package helper

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Pager struct {
	Offset       int    // 起始行数
	Size         int    // 列表每页显示行数
	TotalPages   int    // 总页数
	TotalRows    int    // 总行数
	NowPage      int    // 当前页
	Method       string // 处理情况 Ajax分页 Html分页(静态化时) 普通get方式
	Parameter    string // 分页参数
	PageName     string //分页参数的名称
	AjaxFuncName string // ajax方法名
	Plus         int    // 分页偏移量
	Url          string
}

type showFunc func() string

const (
	MAX_ROWS     int = 10 // 每页最多显示
	MIN_ROWS     int = 1  // 每页最少显示
	DEFAULT_ROWS int = 5  // 每页默认显示

	PAGE_OFFSET int = 3 // 分布偏移量
	PAGE_SLICE  int = 5 // 分页条显示块
)

func NewPager(p *Pager, style string) template.HTML {
	p.init()
	return template.HTML(p.show(style))
}

func (p *Pager) init() {
	if (p.Size) >= MAX_ROWS || (p.Size < MIN_ROWS) {
		p.Size = DEFAULT_ROWS
	}
	p.TotalPages = int(math.Ceil(float64(p.TotalRows) / float64(p.Size)))
	if p.PageName == "" {
		p.PageName = "p"
	}
	if p.NowPage < 1 {
		p.NowPage = 1
	}
	if p.TotalPages > 0 && p.NowPage > p.TotalPages {
		p.NowPage = p.TotalPages
	}
	p.Offset = p.Size * (p.NowPage - 1)
	if p.Plus < 1 {
		p.Plus = PAGE_OFFSET
	}
}

func (p *Pager) show(num string) string {
	var link string
	// 少于2页时，不显示分页条
	if p.TotalPages < 2 {
		return link
	}

	var show showFunc
	switch num {
	case "1":
		show = p.showStyle1
	case "2":
		show = p.showStyle2
	case "3":
		show = p.showStyle3
	default:
		show = p.showStyle1
	}

	var html string = "<ol class=\"page-navigator\">%s</ol>"

	return fmt.Sprintf(html, show())
}

func (p *Pager) showStyle1() string {
	var (
		back  string
		plus  int = p.Plus
		begin int
		count int
	)

	// 根据当前页和翻页偏移量 计算分页栏输出的页码范围
	// 如 （1->7）,(5->11)
	if p.NowPage+plus > p.TotalPages {
		begin = p.TotalPages - plus*2
	} else {
		begin = p.NowPage - plus
	}
	if begin < 1 {
		begin = 1
	}

	back = p.FirstPage("« First")
	back += p.PrePage("« Prev")

	count = begin + plus*2
	for i := begin; i <= count; i++ {
		if i > p.TotalPages {
			break
		}
		if i == p.NowPage {
			back += fmt.Sprintf("<li class=\"current\"><a href=\"/page/%d/\">%d</a></li>\n", i, i)
		} else {
			back += p.getLink(i, strconv.Itoa(i))
		}
	}
	back += p.NextPage("Next »")
	back += p.LastPage("Last »")

	return back
}

func (p *Pager) showStyle2() string {
	var (
		back string
	)

	if p.TotalPages <= 1 {
		return back
	}

	back += p.PrePage("<")
	// 遍历分页,从第一页开始
	for i := 1; i <= p.TotalPages; i++ {
		if i == p.NowPage {
			back += fmt.Sprintf("<a class='current'>%d</a>\n", i)
		} else {
			if (p.NowPage-i >= PAGE_SLICE) && (i != 1) {
				back += "<a class='pageMore'>...</a>\n"
				i = p.NowPage - p.Plus
			} else {
				if (i >= p.NowPage+PAGE_SLICE) && (i != p.TotalPages) {
					back += "<span>...</span>\n"
					i = p.TotalPages
				}
				back += p.getLink(i, strconv.Itoa(i)) + "\n"
			}
		}
	}
	back += p.NextPage(">")

	return back
}

func (p *Pager) showStyle3() string {
	var (
		back  string
		plus  int = p.Plus
		begin int
		count int
	)
	if plus+p.NowPage > p.TotalPages {
		begin = p.TotalPages - plus*2
	} else {
		begin = p.NowPage - plus
	}
	if begin < 1 {
		begin = 1
	}
	back = fmt.Sprintf("Total %d Records Split %d Page, Current %d Page, EachPage",
		p.TotalRows, p.TotalPages, p.NowPage)
	back += fmt.Sprintf("<input type=\"text\" value=\"%d\" id=\"pageSize\" size=\"3\" />", p.Size)
	back += p.FirstPage() + "\n"
	back += p.PrePage() + "\n"
	back += p.NextPage() + "\n"
	back += p.LastPage() + "\n"

	back += fmt.Sprintf("<select onchange=\"%s(this.value)\" id=\"gotoPage\">", p.AjaxFuncName)
	count = begin + PAGE_SLICE*2
	for i := begin; i < count; i++ {
		if i > p.TotalPages {
			break
		}
		if i == p.NowPage {
			back += fmt.Sprintf("<option selected=\"true\" value=\"%d\">%d</option>", i, i)
		} else {
			back += fmt.Sprintf("<option value=\"%d\">%d</option>", i, i)
		}
	}
	back += "</select>"

	return back
}

func (p *Pager) FirstPage(arg ...string) string {
	var link, text string
	if len(arg) > 0 {
		text = arg[0]
	} else {
		text = "First Page"
	}

	if p.NowPage > PAGE_SLICE {
		link = p.getLink(1, text)
	}

	return link
}

func (p *Pager) LastPage(arg ...string) string {
	var link, text string
	if len(arg) > 0 {
		text = arg[0]
	} else {
		text = "Last Page"
	}
	if p.NowPage < p.TotalPages-PAGE_SLICE {
		link = p.getLink(p.TotalPages, text)
	}

	return link
}

func (p *Pager) PrePage(arg ...string) string {
	var link, text string
	if len(arg) > 0 {
		text = arg[0]
	} else {
		text = "Prefix Page"
	}
	if p.NowPage > 1 {
		link = p.getLink(p.NowPage-1, text)
	}

	return link
}

func (p *Pager) NextPage(arg ...string) string {
	var link, text string
	if len(arg) > 0 {
		text = arg[0]
	} else {
		text = "Next Page"
	}
	if p.NowPage < p.TotalPages {
		link = p.getLink(p.NowPage+1, text)
	}

	return link
}

func (p *Pager) getLink(pageNum int, text string) (link string) {
	link = "<li><a href=\"%s\" onlick=\"%s\">%s</a></li>\n"
	switch p.Method {
	case "ajax":
		var parameter string
		if p.Parameter != "" {
			parameter = "," + p.Parameter
		}
		var jsFunc string = "%s('%d'%v)"
		jsFunc = fmt.Sprintf(jsFunc, p.AjaxFuncName, pageNum, parameter)
		link = fmt.Sprintf(link, "javascript:;", jsFunc, text)
	case "html":
		var url string = strings.Replace(p.Parameter, "?", strconv.Itoa(pageNum), 1)
		link = fmt.Sprintf(link, url, "javascript:;", text)
	default:
		link = fmt.Sprintf(link, p.getCurUrl(pageNum), "javascript:;", text)
	}

	return
}

func (p *Pager) getCurUrl(page int) string {
	if p.Url == "" {
		var r *http.Request
		p.setUrl(r)
	}
	var link string = p.Url + p.PageName + "=" + strconv.Itoa(page)

	return link
}

func (p *Pager) setUrl(r *http.Request) {
	var link string = r.RequestURI
	if !strings.Contains(r.RequestURI, "?") {
		link += "?"
	}
	link += p.Parameter
	var queryMap url.Values
	if r.URL.RawQuery != "" {
		queryMap = r.URL.Query()
		if _, ok := queryMap[p.PageName]; ok {
			delete(queryMap, p.PageName)
		}
		link = r.URL.RawQuery + "?" + queryMap.Encode()
	}
	if len(queryMap) > 0 {
		link += "&"
	}
	p.Url = link
}
