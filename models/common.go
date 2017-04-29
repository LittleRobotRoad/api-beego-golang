package models

import "github.com/astaxie/beego/orm"

func PageAdd(qr orm.QuerySeter, pageSize int, pageNo int, dataFunc func() (interface{}, error)) (Page, error) {

	var err error
	var data interface{}
	var page Page = Page{PageNo: pageNo, PageSize: pageSize}

	countInt64, _ := qr.Count()
	count := int(countInt64)
	if count != 0 {
		page.TotalPage = count/pageSize + 1
	}
	page.TotalCount = count
	if (pageNo-1)*pageSize < count {

		if pageNo*pageSize < count {
			page.LastPage = true
		} else {
			page.LastPage = false
		}

		data, err = dataFunc()
		page.List = data
	} else {
		page.LastPage = false
	}
	return page, err
}
