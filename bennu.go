package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type User struct {
	Name string
}

type Message struct {
	MsgContent string
	user       User
	Date       string
	Hour       string
}

func checkIfDate(date string) bool {
	count := strings.Count(date, "/")
	if count == 2 {
		return true
	}
	return false
}

func checkIfHour(hour string) bool {
	if len(hour) == 5 && strings.Count(hour, ":") == 1 {
		return true
	}
	return false
}

func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func trimUser(user string) string {
	return strings.Trim(user, ":")
}

func msglm(txt string) map[string]int {
	const emojiPattern = "[\\x{2712}\\x{2714}\\x{2716}\\x{271d}\\x{2721}\\x{2728}\\x{2733}\\x{2734}\\x{2744}\\x{2747}\\x{274c}\\x{274e}\\x{2753}-\\x{2755}\\x{2757}\\x{2763}\\x{2764}\\x{2795}-\\x{2797}\\x{27a1}\\x{27b0}\\x{27bf}\\x{2934}\\x{2935}\\x{2b05}-\\x{2b07}\\x{2b1b}\\x{2b1c}\\x{2b50}\\x{2b55}\\x{3030}\\x{303d}\\x{1f004}\\x{1f0cf}\\x{1f170}\\x{1f171}\\x{1f17e}\\x{1f17f}\\x{1f18e}\\x{1f191}-\\x{1f19a}\\x{1f201}\\x{1f202}\\x{1f21a}\\x{1f22f}\\x{1f232}-\\x{1f23a}\\x{1f250}\\x{1f251}\\x{1f300}-\\x{1f321}\\x{1f324}-\\x{1f393}\\x{1f396}\\x{1f397}\\x{1f399}-\\x{1f39b}\\x{1f39e}-\\x{1f3f0}\\x{1f3f3}-\\x{1f3f5}\\x{1f3f7}-\\x{1f4fd}\\x{1f4ff}-\\x{1f53d}\\x{1f549}-\\x{1f54e}\\x{1f550}-\\x{1f567}\\x{1f56f}\\x{1f570}\\x{1f573}-\\x{1f579}\\x{1f587}\\x{1f58a}-\\x{1f58d}\\x{1f590}\\x{1f595}\\x{1f596}\\x{1f5a5}\\x{1f5a8}\\x{1f5b1}\\x{1f5b2}\\x{1f5bc}\\x{1f5c2}-\\x{1f5c4}\\x{1f5d1}-\\x{1f5d3}\\x{1f5dc}-\\x{1f5de}\\x{1f5e1}\\x{1f5e3}\\x{1f5ef}\\x{1f5f3}\\x{1f5fa}-\\x{1f64f}\\x{1f680}-\\x{1f6c5}\\x{1f6cb}-\\x{1f6d0}\\x{1f6e0}-\\x{1f6e5}\\x{1f6e9}\\x{1f6eb}\\x{1f6ec}\\x{1f6f0}\\x{1f6f3}\\x{1f910}-\\x{1f918}\\x{1f980}-\\x{1f984}\\x{1f9c0}\\x{3297}\\x{3299}\\x{a9}\\x{ae}\\x{203c}\\x{2049}\\x{2122}\\x{2139}\\x{2194}-\\x{2199}\\x{21a9}\\x{21aa}\\x{231a}\\x{231b}\\x{2328}\\x{2388}\\x{23cf}\\x{23e9}-\\x{23f3}\\x{23f8}-\\x{23fa}\\x{24c2}\\x{25aa}\\x{25ab}\\x{25b6}\\x{25c0}\\x{25fb}-\\x{25fe}\\x{2600}-\\x{2604}\\x{260e}\\x{2611}\\x{2614}\\x{2615}\\x{2618}\\x{261d}\\x{2620}\\x{2622}\\x{2623}\\x{2626}\\x{262a}\\x{262e}\\x{262f}\\x{2638}-\\x{263a}\\x{2648}-\\x{2653}\\x{2660}\\x{2663}\\x{2665}\\x{2666}\\x{2668}\\x{267b}\\x{267f}\\x{2692}-\\x{2694}\\x{2696}\\x{2697}\\x{2699}\\x{269b}\\x{269c}\\x{26a0}\\x{26a1}\\x{26aa}\\x{26ab}\\x{26b0}\\x{26b1}\\x{26bd}\\x{26be}\\x{26c4}\\x{26c5}\\x{26c8}\\x{26ce}\\x{26cf}\\x{26d1}\\x{26d3}\\x{26d4}\\x{26e9}\\x{26ea}\\x{26f0}-\\x{26f5}\\x{26f7}-\\x{26fa}\\x{26fd}\\x{2702}\\x{2705}\\x{2708}-\\x{270d}\\x{270f}]|\\x{23}\\x{20e3}|\\x{2a}\\x{20e3}|\\x{30}\\x{20e3}|\\x{31}\\x{20e3}|\\x{32}\\x{20e3}|\\x{33}\\x{20e3}|\\x{34}\\x{20e3}|\\x{35}\\x{20e3}|\\x{36}\\x{20e3}|\\x{37}\\x{20e3}|\\x{38}\\x{20e3}|\\x{39}\\x{20e3}|\\x{1f1e6}[\\x{1f1e8}-\\x{1f1ec}\\x{1f1ee}\\x{1f1f1}\\x{1f1f2}\\x{1f1f4}\\x{1f1f6}-\\x{1f1fa}\\x{1f1fc}\\x{1f1fd}\\x{1f1ff}]|\\x{1f1e7}[\\x{1f1e6}\\x{1f1e7}\\x{1f1e9}-\\x{1f1ef}\\x{1f1f1}-\\x{1f1f4}\\x{1f1f6}-\\x{1f1f9}\\x{1f1fb}\\x{1f1fc}\\x{1f1fe}\\x{1f1ff}]|\\x{1f1e8}[\\x{1f1e6}\\x{1f1e8}\\x{1f1e9}\\x{1f1eb}-\\x{1f1ee}\\x{1f1f0}-\\x{1f1f5}\\x{1f1f7}\\x{1f1fa}-\\x{1f1ff}]|\\x{1f1e9}[\\x{1f1ea}\\x{1f1ec}\\x{1f1ef}\\x{1f1f0}\\x{1f1f2}\\x{1f1f4}\\x{1f1ff}]|\\x{1f1ea}[\\x{1f1e6}\\x{1f1e8}\\x{1f1ea}\\x{1f1ec}\\x{1f1ed}\\x{1f1f7}-\\x{1f1fa}]|\\x{1f1eb}[\\x{1f1ee}-\\x{1f1f0}\\x{1f1f2}\\x{1f1f4}\\x{1f1f7}]|\\x{1f1ec}[\\x{1f1e6}\\x{1f1e7}\\x{1f1e9}-\\x{1f1ee}\\x{1f1f1}-\\x{1f1f3}\\x{1f1f5}-\\x{1f1fa}\\x{1f1fc}\\x{1f1fe}]|\\x{1f1ed}[\\x{1f1f0}\\x{1f1f2}\\x{1f1f3}\\x{1f1f7}\\x{1f1f9}\\x{1f1fa}]|\\x{1f1ee}[\\x{1f1e8}-\\x{1f1ea}\\x{1f1f1}-\\x{1f1f4}\\x{1f1f6}-\\x{1f1f9}]|\\x{1f1ef}[\\x{1f1ea}\\x{1f1f2}\\x{1f1f4}\\x{1f1f5}]|\\x{1f1f0}[\\x{1f1ea}\\x{1f1ec}-\\x{1f1ee}\\x{1f1f2}\\x{1f1f3}\\x{1f1f5}\\x{1f1f7}\\x{1f1fc}\\x{1f1fe}\\x{1f1ff}]|\\x{1f1f1}[\\x{1f1e6}-\\x{1f1e8}\\x{1f1ee}\\x{1f1f0}\\x{1f1f7}-\\x{1f1fb}\\x{1f1fe}]|\\x{1f1f2}[\\x{1f1e6}\\x{1f1e8}-\\x{1f1ed}\\x{1f1f0}-\\x{1f1ff}]|\\x{1f1f3}[\\x{1f1e6}\\x{1f1e8}\\x{1f1ea}-\\x{1f1ec}\\x{1f1ee}\\x{1f1f1}\\x{1f1f4}\\x{1f1f5}\\x{1f1f7}\\x{1f1fa}\\x{1f1ff}]|\\x{1f1f4}\\x{1f1f2}|\\x{1f1f5}[\\x{1f1e6}\\x{1f1ea}-\\x{1f1ed}\\x{1f1f0}-\\x{1f1f3}\\x{1f1f7}-\\x{1f1f9}\\x{1f1fc}\\x{1f1fe}]|\\x{1f1f6}\\x{1f1e6}|\\x{1f1f7}[\\x{1f1ea}\\x{1f1f4}\\x{1f1f8}\\x{1f1fa}\\x{1f1fc}]|\\x{1f1f8}[\\x{1f1e6}-\\x{1f1ea}\\x{1f1ec}-\\x{1f1f4}\\x{1f1f7}-\\x{1f1f9}\\x{1f1fb}\\x{1f1fd}-\\x{1f1ff}]|\\x{1f1f9}[\\x{1f1e6}\\x{1f1e8}\\x{1f1e9}\\x{1f1eb}-\\x{1f1ed}\\x{1f1ef}-\\x{1f1f4}\\x{1f1f7}\\x{1f1f9}\\x{1f1fb}\\x{1f1fc}\\x{1f1ff}]|\\x{1f1fa}[\\x{1f1e6}\\x{1f1ec}\\x{1f1f2}\\x{1f1f8}\\x{1f1fe}\\x{1f1ff}]|\\x{1f1fb}[\\x{1f1e6}\\x{1f1e8}\\x{1f1ea}\\x{1f1ec}\\x{1f1ee}\\x{1f1f3}\\x{1f1fa}]|\\x{1f1fc}[\\x{1f1eb}\\x{1f1f8}]|\\x{1f1fd}\\x{1f1f0}|\\x{1f1fe}[\\x{1f1ea}\\x{1f1f9}]|\\x{1f1ff}[\\x{1f1e6}\\x{1f1f2}\\x{1f1fc}]"
	re, _ := regexp.Compile(emojiPattern)
	var lol = 0
	var lmao = 0
	lol += strings.Count(strings.ToLower(txt), "lol")
	lmao += strings.Count(strings.ToLower(txt), "lmao")
	emoji := re.FindAllStringSubmatch(txt, -1)

	return map[string]int{
		"lol":   lol,
		"lmao":  lmao,
		"emoji": len(emoji),
	}
}

func msgProf(txt string) int {
	items := []string{"fuck", "merde", "putain", "ass"}
	matches := 0
	for _, item := range items {
		matches += strings.Count(strings.ToLower(txt), item)
	}
	return matches
}

func msgGrt(txt string) int {
	items := []string{"amen", "akpe", "merci", "nagode", "imela", "thanks", "thank you", "alhamdulillah", "shukran"}
	matches := 0
	for _, item := range items {
		matches += strings.Count(strings.ToLower(txt), item)
	}
	return matches
}

func angryEmoj(txt string) int {
	nbr := strings.Count(txt, "😡")
	return nbr
}

func main()  {
	file, err := os.Open("doit")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	messages := make(map[int]*Message)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	var users []string
	var msgNbr = 0
	for _, line := range lines {
		words := strings.Fields(line)
		if len(line) != 0 {
			if checkIfDate(words[0]) && checkIfHour(words[1]) {
				msgNbr++
				user := trimUser(words[3])
				messages[msgNbr] = &Message{
					MsgContent: strings.Join(words[4:], " "),
					user: User{
						Name: user,
					},
					Date: words[0],
					Hour: words[1],
				}
				if len(users) == 0  && user != "Messages" && user != "You"{
					users = append(users, user)
				}
				_, isFind := find(users, user)
				if !isFind && user != "Messages" && user != "You"{
					users = append(users, user)
				}
			} else {
				newLine := messages[msgNbr]
				newLine.MsgContent = newLine.MsgContent + " " + line
			}
		}
	}

	userStr := ""
	for _, user := range users {
		userStr += user
	}
	userStr = strings.Join(users, ", ")
	fmt.Println("------------------------------------")
	fmt.Println("\nHi! This application will try to analyse \nyour whatsapp history file.")
	fmt.Println("\nYour file contain ", len(users), " user(s) : ", userStr)
	fmt.Print("\nPlease enter the user on witch to run the analyse\nUser: ")

	var input string
	fmt.Scanf("%s", &input)
	txt1 := ""
	txt2 := ""
	nbrMsgUser := 0
	isValid := strings.Contains(userStr, input)
	if isValid {
		for _, message := range messages {
			if message.user.Name == input {
				txt1 += message.MsgContent
				nbrMsgUser++
			} else {
				txt2 += message.MsgContent
			}
		}
		lme := msglm(txt1)
		profan := msgProf(txt1)
		emo := msglm(txt2)
		angry := angryEmoj(txt2)
		u := msgGrt(txt1)
		o := msgGrt(txt2)
		//fmt.Printf("Users are %v", users)
		fmt.Println("Total number of messages sent by", input, " : ", nbrMsgUser)
		fmt.Println("Total number of times", input, "sent lol : ", lme["lol"])
		fmt.Println("Total number of times", input, "sent lmao : ", lme["lmao"])
		fmt.Println("Total number of times", input, "sent emojis : ", lme["emoji"])
		fmt.Println("Total number of profanities", input, "sent : ", profan)
		fmt.Println("Total number of times", input, "recieved emojis : ", emo["emoji"])
		fmt.Println("Total number of times", input, "recieved the angry 😡 emoji : ", angry)
		fmt.Println("Total number of times ", input, " sent : ", u, "and recieved : ", o, "the words amen, akpe, merci, nagode, imela, thanks, thank you, alhamdulillah, shukran")
	} else {
		println("This entry is not valid. Please try again")
	}
}
