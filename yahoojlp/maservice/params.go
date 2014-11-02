package maservice

const (
	ResultType_形態素解析  = "ma"
	ResultType_出現頻度情報 = "uniq"

	ResponseType_表記    = "surface"
	ResponseType_読みがな  = "reading"
	ResponseType_品詞    = "pos"
	ResponseType_基本形表記 = "baseform"
	ResponseType_全情報   = "feature"

	FilterType_形容詞  = 1
	FilterType_形容動詞 = 2
	FilterType_感動詞  = 3
	FilterType_副詞   = 4
	FilterType_連体詞  = 5
	FilterType_接続詞  = 6
	FilterType_接頭辞  = 7
	FilterType_接尾辞  = 8
	FilterType_名詞   = 9
	FilterType_動詞   = 10
	FilterType_助詞   = 11
	FilterType_助動詞  = 12
	FilterType_特殊   = 13
)

var ResultTypes = []string{ResultType_形態素解析, ResultType_出現頻度情報}

var ResponseTypes = []string{ResponseType_表記, ResponseType_読みがな, ResponseType_品詞, ResponseType_基本形表記, ResponseType_全情報}

var FilterTypes = []int{FilterType_形容詞, FilterType_形容動詞, FilterType_感動詞, FilterType_副詞, FilterType_連体詞, FilterType_接続詞, FilterType_接頭辞, FilterType_接尾辞, FilterType_名詞, FilterType_動詞, FilterType_助詞, FilterType_助動詞, FilterType_特殊}
