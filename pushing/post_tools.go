package pushing

import (
	"athena/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// 根据 news_id 查找一个 news
func Select(conn *mgo.Collection, news_id string, p *models.Post) {
	conn.Find(&bson.M{"_id": news_id}).One(&p)
}

// 新增 / 删除 一条评论
func CommentAddOrReduce(conn *mgo.Collection, p models.Post, action string, comment models.Comment) error {
	bsonid := bson.M{"_id": p.Id}

	switch action {
	case "+":
		comments := append(p.Comments, comment)
		bsonobj := bson.M{"$set": &bson.M{"comments": comments}}
		if err := conn.Update(&bsonid, &bsonobj); err != nil {
			return err
		}
		return nil
	case "-":
		arr := BinarySearch(p.Comments, comment.Id)
		bsonobj := bson.M{"$set": &bson.M{"comments": arr}}
		if err := conn.Update(&bsonid, &bsonobj); err != nil {
			return err
		}
		return nil
	}
	return nil
}

// 点赞 / 取消点赞 一条评论
func FabulousAddOrReduce(conn *mgo.Collection, p models.Post, action string, comment models.Comment) error {
	bsonid := bson.M{"_id": p.Id}

	switch action {
	case "+":
		arr := BinarySearchFabulous(p.Comments, comment.Id, "+")
		bsonobj := bson.M{"$set": &bson.M{"comments": arr}}
		if err := conn.Update(&bsonid, &bsonobj); err != nil {
			return err
		}
		return nil
	case "-":
		arr := BinarySearchFabulous(p.Comments, comment.Id, "-")
		bsonobj := bson.M{"$set": &bson.M{"comments": arr}}
		if err := conn.Update(&bsonid, &bsonobj); err != nil {
			return err
		}
	}
	return nil
}

// 踩 / 取消点赞 一条评论
func StepOnAddOrReduce(conn *mgo.Collection, p models.Post, action string, comment models.Comment) error {
	bsonid := bson.M{"_id": p.Id}

	switch action {
	case "+":
		arr := BinarySearchFabulous(p.Comments, comment.Id, "+")
		bsonobj := bson.M{"$set": &bson.M{"comments": arr}}
		if err := conn.Update(&bsonid, &bsonobj); err != nil {
			return err
		}
		return nil
	case "-":
		arr := BinarySearchFabulous(p.Comments, comment.Id, "-")
		bsonobj := bson.M{"$set": &bson.M{"comments": arr}}
		if err := conn.Update(&bsonid, &bsonobj); err != nil {
			return err
		}
	}
	return nil
}

// 二分查找 跟新一个评论的踩数量
func BinarySearchStepOn(arr []models.Comment, target bson.ObjectId, action string) []models.Comment {
	//变量l,r的意义是：从[l......r]区间内查找target的数的下标
	l, r := 0, len(arr)-1
	arr_index := 0
	//l == r时，说明数组只有1个数，若l>r时，则数组没有元素条件结束
	for l <= r {
		mid := (l + r) / 2
		//正好是中位数
		if target == arr[mid].Id {
			arr_index = mid
			break
		}
		//在右区间查找
		if target > arr[mid].Id {
			l = mid + 1
		} else { //在左区间查找
			r = mid - 1
		}
	}

	switch action {
	case "+":
		arr[arr_index].Eggs += 1
	case "-":
		if num := arr[arr_index].Eggs - int64(1); num < 0 {
			arr[arr_index].Eggs = 0
		}
	}

	return arr
}

// 二分查找 跟新一个评论的点赞数量
func BinarySearchFabulous(arr []models.Comment, target bson.ObjectId, action string) []models.Comment {
	//变量l,r的意义是：从[l......r]区间内查找target的数的下标
	l, r := 0, len(arr)-1
	arr_index := 0
	//l == r时，说明数组只有1个数，若l>r时，则数组没有元素条件结束
	for l <= r {
		mid := (l + r) / 2
		//正好是中位数
		if target == arr[mid].Id {
			arr_index = mid
			break
		}
		//在右区间查找
		if target > arr[mid].Id {
			l = mid + 1
		} else { //在左区间查找
			r = mid - 1
		}
	}

	switch action {
	case "+":
		arr[arr_index].Flowers += 1
	case "-":
		if num := arr[arr_index].Flowers - int64(1); num < 0 {
			arr[arr_index].Flowers = 0
		}
	}

	return arr
}

// 二分查找 删除一个 评论
func BinarySearch(arr []models.Comment, target bson.ObjectId) []models.Comment {
	//变量l,r的意义是：从[l......r]区间内查找target的数的下标
	l, r := 0, len(arr)-1
	arr_index := 0
	//l == r时，说明数组只有1个数，若l>r时，则数组没有元素条件结束
	for l <= r {
		mid := (l + r) / 2
		//正好是中位数
		if target == arr[mid].Id {
			arr_index = mid
			break
		}
		//在右区间查找
		if target > arr[mid].Id {
			l = mid + 1
		} else { //在左区间查找
			r = mid - 1
		}
	}

	arr1 := arr[:arr_index]
	for _, v := range arr[arr_index+1:] {
		arr = append(arr1, v)
	}

	return arr1
}
