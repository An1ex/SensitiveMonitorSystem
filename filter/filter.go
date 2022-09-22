package filter

import (
	"log"

	"bili-monitor-system/alarm"
	"bili-monitor-system/db"

	filter "github.com/antlinker/go-dirtyfilter"
	"github.com/antlinker/go-dirtyfilter/store"
)

var (
	memStore *store.MemoryStore
)

func Init() error {
	var err error
	//sf, _ := os.Open("sensitive/guns.txt")
	//defer sf.Close()
	//s := bufio.NewReader(sf)
	memStore, err = store.NewMemoryStore(store.MemoryConfig{
		//Reader:     s,
		DataSource: []string{"苦难"},
	})
	if err != nil {
		return err
	}
	return nil
}

func Filter(bid string, comments db.Comments) int {
	count := 0
	for _, comment := range comments {
		originText := comment.Content
		filterManage := filter.NewDirtyManager(memStore)
		result, _ := filterManage.Filter().Filter(originText, '@', '，', '。', '[', ']', '！')
		if len(result) != 0 {
			count += 1
			err := alarm.Alarm(bid, comment)
			if err != nil {
				log.Printf("[alarm] %v", err)
			}
		}
	}
	return count
}
