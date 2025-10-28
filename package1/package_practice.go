package package1

import (
	"fmt"
	"math"
	"math/big"
	"regexp"
	"strconv"
)

const LINUX_REBOOT_MAGIC1 uintptr = 0xfee1dead
const LINUX_REBOOT_MAGIC2 uintptr = 672274793
const LINUX_REBOOT_CMD_RESTART uintptr = 0x1234567

func reboot_go() {
	//syscall.Syscall(syscall.SYS_REBOOT, LINUX_REBOOT_MAGIC1, LINUX_REBOOT_MAGIC2, LINUX_REBOOT_CMD_RESTART)
}

func regexp1() {
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+"
	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 64)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}
	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match Found!")
	}
	re, _ := regexp.Compile(pat)
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)

}
func regjs() {
	jscontent := `password:123`

	// 修正后的模式：使用反引号 `...` 避免转义问题
	// 1. 去掉双引号字符串中的所有额外 \
	// 2. 模式依然要求值被引号包裹 ( ' 或 " ) 且长度至少为 8
	pat := `password`

	re, err := regexp.Compile(pat)
	if err != nil {
		fmt.Println("匹配失败，请检查正则表达式:", err)
		return
	}

	fmt.Println("--- 查找结果 ---")
	// 使用 FindAllStringSubmatch 来查找并确认所有捕获组
	matches := re.FindAllStringSubmatch(jscontent, -1)

	if len(matches) == 0 {
		fmt.Println("未找到匹配内容。")
	}

	for i, match := range matches {
		fmt.Printf("%d. 完整匹配: %s\n", i+1, match[0])
	}

	// ----------------------------------------------------
	// 添加一个能匹配的测试案例，以便看到正确的输出
	fmt.Println("\n--- 修正后的测试案例 ---")
	jscontent_fixed := `var myPassword = "securepassword123";`

	matches_fixed := re.FindAllStringSubmatch(jscontent_fixed, -1)
	fmt.Println(matches_fixed)
	if len(matches_fixed) > 0 {
		// 假设 "PasswordValue" 在索引 5

		fmt.Printf("找到匹配！完整匹配: %s\n", matches_fixed[0][0])
		fmt.Printf("捕获的值: %s\n", matches_fixed[0])
	} else {
		fmt.Println("测试案例未匹配。")
	}
}
func bigInt() {
	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(1)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Printf("Big Int: %v\n", ip)
	fmt.Println("===============")
	rm := big.NewRat(math.MaxInt64, 1956)
	rn := big.NewRat(-1956, math.MaxInt64)
	ro := big.NewRat(19, 56)
	rp := big.NewRat(1111, 2222)
	rq := big.NewRat(1, 1)
	rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
	fmt.Printf("Big Rat:%v\n", rq)
}
