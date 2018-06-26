package context

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/rand"
	"sync"
	"time"
)

type Cxt struct {
	name string //用于表示这个元数据结构的名称，方便以后区分不同模型（TCC一致模型、补偿模型、不可靠消息模型、可靠消息模型）
	mu   *sync.Mutex
}
type metadata struct {
	s_id       int
	d_id       int
	s_funcname string
	d_funcname string
}

var ErrJson = errors.New("json fail")

func NewFactory() *Cxt {
	cxt := new(Cxt)
	cxt.mu = new(sync.Mutex)
	return cxt
}

func (this *Cxt) CreateActionID(funcname string) int {
	this.mu.Lock()
	defer this.mu.Unlock()
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(100)
	return id

}

//组装需要存储的元数据，以json格式存储
func (this *Cxt) ChangeToJson(s_id, d_id int, s_funcname, d_funcname string) (string, error) {
	m := new(metadata)
	m.s_id = s_id
	m.s_funcname = s_funcname
	m.d_id = d_id
	m.d_funcname = d_funcname

	bytes, err := json.Marshal(m)
	if err != nil {
		return "", ErrJson
	}
	return string(bytes), nil
}

//确定封装一组调用的元数据生成一个唯一id
func (this *Cxt) AssignUniqueID(meta string) string {
	h := md5.New()
	h.Write([]byte(meta))
	tmp := h.Sum(nil)
	id := hex.EncodeToString(tmp)
	return id
}

//确定每一步原子操作的正确性数据结构,enable : 1:成功，2:失败
func (this *Cxt) MakeSureAtomicOperation(metaId string, enable int) map[string]int {
	m := make(map[string]int)
	m[metaId] = enable
	return m
}
