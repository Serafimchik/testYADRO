package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type person_in_table struct {
	client_name    string
	start_sit_time string
}
type table_info struct {
	total_income    int
	total_work_time time.Time
}

func output_time(hour int, minute int) string {
	out_hour := ""
	out_minute := ""
	if hour < 10 {
		out_hour = "0" + strconv.Itoa(hour)
	} else {
		out_hour = strconv.Itoa(hour)
	}
	if minute < 10 {
		out_minute = "0" + strconv.Itoa(minute)
	} else {
		out_minute = strconv.Itoa(minute)
	}
	out_time := out_hour + ":" + out_minute
	return out_time
}

func count_icome(answer *[]table_info, coast int, layout string, tables []person_in_table, client_is_at_the_table map[string]int, client_name string, time_end_sitting time.Time, table int) {
	time_start_sitting, _ := time.Parse(layout, tables[table].start_sit_time)
	time_spent_at_the_table := time_end_sitting.Sub(time_start_sitting)
	hours_for_pay := int(time_spent_at_the_table.Hours())
	if time_spent_at_the_table.Minutes() > 0 {
		hours_for_pay++
	}
	table_id := client_is_at_the_table[client_name]
	(*answer)[table_id].total_work_time = (*answer)[table_id].total_work_time.Add(time.Duration(time_spent_at_the_table.Minutes()) * time.Minute)
	(*answer)[table_id].total_income += hours_for_pay * coast
}

func client_from_queue_sit(client_is_at_the_table *map[string]int, queue *[]string, tables *[]person_in_table, table_is_free *map[int]bool, input_event []string) {
	table_id := (*client_is_at_the_table)[input_event[2]]
	if len(*queue) > 0 {
		(*tables)[table_id].client_name = (*queue)[0]
		(*tables)[table_id].start_sit_time = input_event[0]
		(*client_is_at_the_table)[(*queue)[0]] = table_id
	} else {
		(*table_is_free)[table_id] = true
		(*tables)[table_id].client_name = ""
		(*tables)[table_id].start_sit_time = ""
	}
	delete((*client_is_at_the_table), input_event[2])
	if len(*queue) > 1 {
		*queue = (*queue)[1:]
	} else {
		*queue = (*queue)[:0]
	}
}

func output_event_ID_11(out_time string, client_name string) {
	fmt.Printf("%s 11 %s", out_time, client_name)
	fmt.Println()
}

func output_event_ID_12(out_time string, client_name string, table_number int) {
	fmt.Printf("%s 12 %s %d", out_time, client_name, table_number)
	fmt.Println()
}

func output_event_ID_13(out_time string, type_of_error string) {
	fmt.Printf("%s 13 %s", out_time, type_of_error)
	fmt.Println()
}

func time_is_valid(time string) bool {
	if len(time) != 5 || time[2] != ':' {
		return false
	}
	parts := strings.Split(time, ":")
	if len(parts) != 2 {
		return false
	}
	hours, err1 := strconv.Atoi(parts[0])
	minutes, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil || hours < 0 || hours > 23 || minutes < 0 || minutes > 59 {
		return false
	}
	return true
}

func client_name_is_valid(name string) bool {
	for _, char := range name {
		if !(letter_is_valid(char) || digit_is_valid(char) || char == '_' || char == '-') {
			return false
		}
	}
	return true
}

func letter_is_valid(char rune) bool {
	return char >= 'a' && char <= 'z'
}

func digit_is_valid(char rune) bool {
	return char >= '0' && char <= '9'
}

func number_is_valid(number string) bool {
	x, err1 := strconv.Atoi(number)
	if err1 != nil {
		return false
	}
	if x < 0 {
		return false
	}
	return true
}

func chek_for_valid(scanner *bufio.Scanner) {
	line := 0
	N := 0
	layout := "15:04"
	previos_time, _ := time.Parse(layout, "00:00")
	for scanner.Scan() {
		if line == 0 {
			stri := strings.Split(scanner.Text(), " ")
			if len(stri) != 1 {
				fmt.Print(scanner.Text())
				os.Exit(0)
			}
			if number_is_valid(stri[0]) == false {
				fmt.Print(scanner.Text())
				os.Exit(0)
			}
			N, _ = strconv.Atoi(stri[0])
			line++
			continue
		}
		if line == 1 {
			stri := strings.Split(scanner.Text(), " ")
			if len(stri) != 2 {
				fmt.Print(scanner.Text())
				os.Exit(0)
			}
			timeStart := stri[0]
			timeEnd := stri[1]
			if time_is_valid(timeStart) == false || time_is_valid(timeEnd) == false {
				fmt.Print(scanner.Text())
				os.Exit(0)
			}
			line++
			continue
		}
		if line == 2 {
			stri := strings.Split(scanner.Text(), " ")
			if len(stri) != 1 {
				fmt.Print(scanner.Text())
				os.Exit(0)
			}
			if number_is_valid(stri[0]) == false {
				fmt.Print(scanner.Text())
				os.Exit(0)
			}
			line++
			continue

		}
		input_event := strings.Split(scanner.Text(), " ")
		if len(input_event) < 3 && len(input_event) > 4 {
			fmt.Print(scanner.Text())
			os.Exit(0)
		}
		if time_is_valid(input_event[0]) == false {
			fmt.Print(scanner.Text())
			os.Exit(0)
		}
		if number_is_valid(input_event[1]) == false {
			fmt.Print(scanner.Text())
			os.Exit(0)
		}
		input_id_event, _ := strconv.Atoi(input_event[1])
		if input_id_event == 2 {
			if len(input_event) != 4 {
				fmt.Print(scanner.Text())
				os.Exit(0)
			}
			if number_is_valid(input_event[3]) == false {
				fmt.Print(scanner.Text())
				os.Exit(0)
			}
			table, _ := strconv.Atoi(input_event[3])
			if table < 0 || table > N {
				fmt.Print(scanner.Text())
				os.Exit(0)
			}
		} else if input_id_event == 1 || input_id_event == 3 || input_id_event == 4 {
			if len(input_event) != 3 {
				fmt.Print(scanner.Text())
				os.Exit(0)
			}
		} else {
			fmt.Print(scanner.Text())
			os.Exit(0)
		}
		if client_name_is_valid(input_event[2]) == false {
			fmt.Print(scanner.Text())
			os.Exit(0)
		}
		new_time_event, _ := time.Parse(layout, input_event[0])
		if new_time_event.After(previos_time) == true || new_time_event == previos_time {
			previos_time = new_time_event
		} else {
			fmt.Print(scanner.Text())
			os.Exit(0)
		}
	}
}
func client_go_out(clients_in_clud *[]string, client_name string) {
	for i := 0; i < len(*clients_in_clud); i++ {
		if (*clients_in_clud)[i] == client_name {
			(*clients_in_clud) = append((*clients_in_clud)[:i], (*clients_in_clud)[i+1:]...)
			break
		}
	}
}
func main() {
	if !(len(os.Args) == 2) {
		panic("use task.exe test<N>.txt")
	}
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scannertest := scanner
	chek_for_valid(scannertest)
	file, err = os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner = bufio.NewScanner(file)
	line := 0 //для подсчета по какой строке файла идём
	layout := "15:04"
	var N, coast int                                   //кол-во столов и стоимость в час
	var opening_time_club, closing_time_club time.Time //время начала работы, время окончания работы клуба
	var timeStart, timeEnd string                      //время начала работы, время окончания работы клуба
	var queue []string                                 //очередь
	table_is_free := make(map[int]bool)                //проверка на свободен ли стол использую мапу для того чтоб получать ответ за O(1)
	client_in_club := make(map[string]bool)            //проверка клиент в клубе или нет использую мапу для того чтоб получать ответ за O(1)
	client_is_at_the_table := make(map[string]int)     //чтоб быстро узнать за каким столом сидит клиент использую мапу для того чтоб получать ответ за O(1)
	tables := make([]person_in_table, N+1, N+1)        //кто сидит за столом
	number_of_free_tables := 0                         // кол-во свободных столов
	answer := make([]table_info, N+1, N+1)             //ответ для каждого стола сколько он принёс денег и сколько времени был занят
	var clients_in_clud []string                       //хранит всех клиентов чтоб после закрытия клуба вывести что они ушли
	for scanner.Scan() {
		if line == 0 {
			N, _ = strconv.Atoi(scanner.Text())
			tables = make([]person_in_table, N+1, N+1)
			answer = make([]table_info, N+1, N+1)
			number_of_free_tables = N
			for i := 1; i <= N; i++ { //заполняю что все столы свободны
				table_is_free[i] = true
			}
			line++
			continue
		}
		if line == 1 {
			stri := strings.Split(scanner.Text(), " ")
			timeStart = stri[0]
			timeEnd = stri[1]
			fmt.Println(timeStart)
			opening_time_club, _ = time.Parse(layout, timeStart)
			closing_time_club, _ = time.Parse(layout, timeEnd)
			line++
			continue
		}
		if line == 2 {
			coast, _ = strconv.Atoi(scanner.Text())
			line++
			continue

		}
		input_event := strings.Split(scanner.Text(), " ")
		time_of_input_event, _ := time.Parse(layout, input_event[0])
		switch input_event[1] {
		case "1": //клиент пришел
			fmt.Println(scanner.Text())
			if client_in_club[input_event[2]] == true { //если клиент уже в клубе
				output_event_ID_13(output_time(int(time_of_input_event.Hour()), int(time_of_input_event.Minute())), "YouShallNotPass")
			} else if (time_of_input_event.After(opening_time_club) == true && time_of_input_event.Before(closing_time_club) == true) || time_of_input_event.Equal(opening_time_club) == true || time_of_input_event.Equal(closing_time_club) == true {
				client_in_club[input_event[2]] = true
				clients_in_clud = append(clients_in_clud, input_event[2])
			} else { //если клиент пришел рано
				output_event_ID_13(output_time(int(time_of_input_event.Hour()), int(time_of_input_event.Minute())), "NotOpenYet")
			}
		case "2": //клиент сел за стол
			fmt.Println(scanner.Text())
			table, _ := strconv.Atoi(input_event[3])    //номер стола за который садится клиент
			if client_in_club[input_event[2]] != true { //если клиент не в клубе
				output_event_ID_13(output_time(int(time_of_input_event.Hour()), int(time_of_input_event.Minute())), "ClientUnknown")
			} else if table_is_free[table] == true && client_is_at_the_table[input_event[2]] == 0 { //если место свободно и клиент не сидит за столом
				tables[table].client_name = input_event[2]     //имя клиента
				tables[table].start_sit_time = input_event[0]  //время во сколько сел за стол
				client_is_at_the_table[input_event[2]] = table //помечаю за каким столом сидит клиент
				table_is_free[table] = false                   //помечаю что стол занят
				number_of_free_tables--                        //кол-во свободных столов уменьшаю
			} else if table_is_free[table] == true && client_is_at_the_table[input_event[2]] != 0 { //если место свободно и клиент сидит за столом он пересаживается
				count_icome(&answer, coast, layout, tables, client_is_at_the_table, input_event[2], time_of_input_event, client_is_at_the_table[input_event[2]])
				if len(queue) > 0 {
					output_event_ID_12(input_event[0], queue[0], client_is_at_the_table[queue[0]])
					client_from_queue_sit(&client_is_at_the_table, &queue, &tables, &table_is_free, input_event)
				}
				tables[table].client_name = input_event[2]     //имя клиента
				tables[table].start_sit_time = input_event[0]  //время во сколько сел за стол
				client_is_at_the_table[input_event[2]] = table //помечаю за каким столом сидит клиент
				table_is_free[table] = false                   //помечаю что стол занят
			} else { //место занято
				output_event_ID_13(output_time(int(time_of_input_event.Hour()), int(time_of_input_event.Minute())), "PlaceIsBusy")
			}
		case "3": //клиент ожидает
			fmt.Println(scanner.Text())
			if number_of_free_tables > 0 { //Если в клубе есть свободные столы, то генерируется ошибка "ICanWaitNoLonger!"
				output_event_ID_13(output_time(int(time_of_input_event.Hour()), int(time_of_input_event.Minute())), "ICanWaitNoLonger!")
			} else if len(queue) > N { //Если в очереди ожидания клиентов больше, чем общее число столов, то клиент уходит и генерируется событие ID 11.
				output_event_ID_11(output_time(int(time_of_input_event.Hour()), int(time_of_input_event.Minute())), input_event[2])
			} else {
				queue = append(queue, input_event[2]) //добавляю клиента в очередь
			}
		case "4": //клиент ушел
			fmt.Println(scanner.Text())
			if client_in_club[input_event[2]] != true { //Если клиент не находится в компьютерном клубе, генерируется ошибка "ClientUnknown".
				output_event_ID_13(output_time(int(time_of_input_event.Hour()), int(time_of_input_event.Minute())), "ClientUnknown")
			} else {
				count_icome(&answer, coast, layout, tables, client_is_at_the_table, input_event[2], time_of_input_event, client_is_at_the_table[input_event[2]])
				table_id := client_is_at_the_table[input_event[2]]
				if len(queue) > 0 {
					output_event_ID_12(input_event[0], queue[0], table_id)
					client_from_queue_sit(&client_is_at_the_table, &queue, &tables, &table_is_free, input_event)
					client_go_out(&clients_in_clud, input_event[2])
				} else {
					number_of_free_tables++
					table_is_free[table_id] = true
					tables[table_id] = person_in_table{}
					delete(client_is_at_the_table, input_event[2])
					client_go_out(&clients_in_clud, input_event[2])
				}
			}
		}

	}
	for i := 1; i < N+1; i++ {
		if tables[i].client_name != "" {
			count_icome(&answer, coast, layout, tables, client_is_at_the_table, tables[i].client_name, closing_time_club, i)
		}
	}
	sort.Strings(clients_in_clud)
	for i := 0; i < len(clients_in_clud); i++ {
		output_event_ID_11(timeEnd, clients_in_clud[i])
	}
	fmt.Println(timeEnd)
	for i := 1; i < N+1; i++ {
		fmt.Printf("%d %d %s", i, answer[i].total_income, output_time(int(answer[i].total_work_time.Hour()), int(answer[i].total_work_time.Minute())))
		fmt.Println()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
