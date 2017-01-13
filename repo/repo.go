package repo

//import "github.com/learning/short-url-go/model"

var currentId int

//var longUrls model.Urls

//
// Mock the database table <LONG_URL>
//+-------------+----------+---------+
//| Field       | Type     |
//+-------------+----------+---------+
//| ID          | int      |   PK    |
//| LongUrl     | char(256)|   Index |
//+-------------+----------+---------+
//
var mLongUrlToId map[string]int
var mIdToLongUrl map[int]string


func InitRepo()  {
	mLongUrlToId = make(map[string]int)
	mIdToLongUrl = make(map[int]string)
}


func RepoFindUrl(longUrl string) int {
	i, ok := mLongUrlToId[longUrl]
	if !ok {
		return -1
	}
	return i
}

func RepoFindUrlById(id int) string {
	url, ok := mIdToLongUrl[id]
	if !ok {
		return ""
	}
	return url
}

func RepoCreateUrl(longUrl string) int {
	//Note: capacity of the ids is MaxInt, 1<<64 - 1 (64-bit system) 1<<31 -1 (32-bit system)
	currentId++
	mLongUrlToId[longUrl] = currentId
	mIdToLongUrl[currentId] = longUrl
	return currentId
}
