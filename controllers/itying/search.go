package itying

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"xiaomi/models"

	"github.com/astaxie/beego"
	"github.com/olivere/elastic/v7"
)

type SearchController struct {
	BaseController
}

//初始化的时候判断goods是否存在 创建索引配置映射
func init() {
	exists, err := models.EsClient.IndexExists("goods").Do(context.Background())
	if err != nil {
		beego.Error(err)
	}
	if !exists {
		// Create a new index.
		mapping := `
		{
			"settings": {
			  "number_of_shards": 1,
			  "number_of_replicas": 0
			},
			"mappings": {
			  "properties": {
				"content": {
				  "type": "text",
				  "analyzer": "ik_max_word",
				  "search_analyzer": "ik_max_word"
				},
				"title": {
				  "type": "text",
				  "analyzer": "ik_max_word",
				  "search_analyzer": "ik_max_word"
				}
			  }
			}
		  }
		`
		_, err := models.EsClient.CreateIndex("goods").Body(mapping).Do(context.Background())
		if err != nil {
			// Handle error
			beego.Error(err)
		}

	}
}

//增加商品数据
func (c *SearchController) AddGoods() {
	goods := []models.Goods{}
	models.DB.Find(&goods)

	for i := 0; i < len(goods); i++ {
		_, err := models.EsClient.Index().
			Index("goods").
			Id(strconv.Itoa(goods[i].Id)).
			BodyJson(goods[i]).
			Do(context.Background())
		if err != nil {
			// Handle error
			beego.Error(err)
		}
	}

	c.Ctx.WriteString("AddGoods success")

}

//更新数据
func (c *SearchController) Update() {

	// res, err := models.EsClient.Update().
	// 	Index("goods").
	// 	Type("_doc").
	// 	Id("19").
	// 	Doc(map[string]interface{}{
	// 		"Title": "哈哈哈",
	// 	}).
	// 	Do(context.Background())

	// if err != nil {
	// 	beego.Error(err)
	// }
	// fmt.Printf("update %s\n", res.Result)

	//从数据库获取修改

	goods := models.Goods{}
	models.DB.Where("id=20").Find(&goods)
	goods.Title = "苹果电脑"
	goods.SubTitle = "苹果电脑"
	res, err := models.EsClient.Update().
		Index("goods").
		Type("_doc").
		Id("20").
		Doc(goods).
		Do(context.Background())
	if err != nil {
		beego.Error(err)
	}
	fmt.Printf("update %s\n", res.Result)

	c.Ctx.WriteString("修改数据")
}

//删除
func (c *SearchController) Delete() {
	res, err := models.EsClient.Delete().
		Index("goods").
		Type("_doc").
		Id("20").
		Do(context.Background())

	if err != nil {
		beego.Error(err)
	}
	fmt.Printf("Delete %s\n", res.Result)

	c.Ctx.WriteString("删除成功")
}

//查询一条数据
func (c *SearchController) GetOne() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			c.Ctx.WriteString("GetOne")
		}
	}()

	result, _ := models.EsClient.Get().
		Index("goods").
		Id("19").
		Do(context.Background())

	fmt.Println(result.Source)

	goods := models.Goods{}
	json.Unmarshal(result.Source, &goods)
	c.Data["json"] = goods
	c.ServeJSON()

}

//查询多条数据
func (c *SearchController) Query() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			c.Ctx.WriteString("Query")
		}
	}()

	query := elastic.NewMatchQuery("Title", "旗舰")
	searchResult, err := models.EsClient.Search().
		Index("goods").          // search in index "twitter"
		Query(query).            // specify the query
		Do(context.Background()) // execute
	if err != nil {
		// Handle error
		panic(err)
	}
	goodsList := []models.Goods{}
	var goods models.Goods
	for _, item := range searchResult.Each(reflect.TypeOf(goods)) {
		g := item.(models.Goods)
		fmt.Printf("标题： %v\n", g.Title)
		goodsList = append(goodsList, g)
	}

	c.Data["json"] = goodsList
	c.ServeJSON()

}

//条件筛选查询
func (c *SearchController) FilterQuery() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			c.Ctx.WriteString("Query")
		}
	}()

	//筛选
	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewMatchQuery("Title", "小米"))
	boolQ.Filter(elastic.NewRangeQuery("Id").Gt(19))
	boolQ.Filter(elastic.NewRangeQuery("Id").Lt(31))
	searchResult, err := models.EsClient.Search().Index("goods").Type("_doc").Query(boolQ).Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	var goods models.Goods
	for _, item := range searchResult.Each(reflect.TypeOf(goods)) {
		t := item.(models.Goods)
		fmt.Printf("Id:%v 标题：%v\n", t.Id, t.Title)
	}

	c.Ctx.WriteString("filter Query")
}

//分页查询
func (c *SearchController) GoodsList() {
	c.SuperInit()
	keyword := c.GetString("keyword")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			c.Ctx.WriteString("GoodsList")
		}
	}()

	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}
	pageSize := 5

	query := elastic.NewMatchQuery("Title", keyword)
	searchResult, err := models.EsClient.Search().
		Index("goods").
		Query(query).
		Sort("Price", true). //true 升序
		Sort("Id", false).   //false 降序
		From((page - 1) * pageSize).Size(pageSize).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}

	//查询符合条件的商品的总数
	searchResult2, _ := models.EsClient.Search().
		Index("goods").          // search in index "twitter"
		Query(query).            // specify the query
		Do(context.Background()) // execute

	goodsList := []models.Goods{}
	var goods models.Goods
	for _, item := range searchResult.Each(reflect.TypeOf(goods)) {
		g := item.(models.Goods)
		fmt.Printf("标题： %v\n", g.Title)
		goodsList = append(goodsList, g)
	}
	c.Data["goodsList"] = goodsList
	c.Data["totalPages"] = math.Ceil(float64(len(searchResult2.Each(reflect.TypeOf(goods)))) / float64(pageSize))
	c.Data["page"] = page
	c.Data["keyword"] = keyword
	// c.ServeJSON()
	c.TplName = "itying/elasticsearch/list.html"
}
