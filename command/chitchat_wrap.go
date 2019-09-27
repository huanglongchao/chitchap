package command

import (
	"chitchat/data"
	"log"
	"math/rand"
	"strconv"
	"time"
)
type ThreadPost map[string][]string

var threadPost ThreadPost
/**
	注册机器人
 */

func init(){
	threadPost = make(ThreadPost)
	threadPost["曹操"] = []string{
		"奸雄：宁教我负天下人，休教天下人负我！/吾好梦中杀人！",
		"护驾：来人，护驾！/魏将何在？",
		"阵亡：霸业未成，未成啊...",
	}
	threadPost["司马懿"] = []string{
		"反馈：下次注意点！/出来混，早晚要还的！",
		"鬼才：吾乃天命之子！/天命？哈哈哈哈哈哈~！",
		"阵亡：难道真的是天命难违？",
	}
	threadPost["夏侯惇"] = []string{
		"刚烈：以彼之道，还施彼身！/鼠辈，竟敢伤我！",
		"阵亡：两，两边，都看不见了",
	}
	threadPost["邓艾"] = []string{
		"屯田：锄禾日当午，汗滴禾下土。/休养生息，备战待敌。",
		"凿险：开辟险路，奇袭敌军！/屯田日久，当建奇功。",
		"急袭：功其不备，出其不意！/偷渡阴平，直取蜀汉！",
		"阵亡：吾破蜀克敌，竟葬于奸贼之手....",
	}
	threadPost["甄姬"] = []string{
		"倾国：体迅飞凫，飘忽若神。/凌波微步，罗袜生尘。",
		"洛神：仿佛兮若轻云之蔽月。/飘摇兮若流风之回雪。",
		"阵亡：悼良会之永绝兮，哀一逝而异乡。",
	}
	log.Println("wrap: init...")
}

func RegisterWrap() data.User{

	randStr := strconv.Itoa(randNum(100000000,1000000))
	user := data.User{
		Name:     "Robot" + randStr,
		Email:    randStr + "@qq.com",
		Password: "944792",
	}
	if err := user.Create(); err != nil {
		log.Println("wrap: Cannot generate UUID", err)
	}
	return user
}

/**
	灌水发帖
 */
func CreateThreadWrap() string{
	users,err := data.Users()
	if err != nil {
		log.Println("wrap: query users fail", err)
	}
	rand.Seed(time.Now().Unix())
	user := users[rand.Intn(len(users))]
	for topic := range threadPost{
		if _, err := user.CreateThread(topic); err != nil {
			log.Println("wrap: Cannot create thread",err)
		}
		return topic
	}
}


func randNum(max,min int) int{
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min)
}
