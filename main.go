package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"

	"net/http"

	"os"

	chart "github.com/wcharczuk/go-chart/v2"

	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gorilla/mux"
)

//////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////////
//
//
//
//
//
//
//
func createBot() TGtoken {
	var Config = TGtoken{
		Token:         "",
		LocalHostPort: "",
		Password:      "",
		PortWeb:       "",
		updateTimeSec: "",
	}
	fileT, _ := ioutil.ReadFile("data/Config.json")
	json.Unmarshal([]byte(fileT), &Config)
	return Config
}

//
//
//
//
//
//
//
// чтобы изменять сайт через бота
func changerDataTelegram(num string, name string, addres *string, u string, bot tgbotapi.BotAPI, update tgbotapi.Update) {
	if len(u) > 4 {
		if u[0:4] == num {
			*addres = u[4:len(u)]
			answ := "Changed " + name + ":\n" + *addres + "\nCheck it; \n" + dataSialum.DomainName + "/8ADF14ADC5E6CEBB2Bd6B9CF7FF70B5F9Bb56B1C59E8A1C0BEF69d0FCE36128DgEBC"
			file, _ := json.MarshalIndent(dataSialum, "", " ")
			_ = ioutil.WriteFile("data/dataSialum.json", file, 0644)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, answ)
			bot.Send(msg)
		}
	}
}

//bot
//фнкция оставляющая в списке строк только уникальные строки
func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

//      NOTIFICATIONS
////чтобы хранить уведомленияс сайта
var Notifiations = notifiations{
	Notification: []Notif{},
}

//bot
//форма записи уведомления
type Notif struct {
	Ip     string
	Nano   int
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
	Secund int
	Deal   string
}

type dialogSample struct {
	ParrentID string
	User      string
	Bot       string
}

//bot
//чтобы хранить верифицированные айди
type RouteVerifUsers struct {
	ChatID []string
}

//bot
//чтобы хранить скрипт диалога
type script struct {
	Scripts map[string]dialogSample
}

//bot
//чтобы хранить уведомления
type notifiations struct {
	Notification []Notif
}

//bot
//чтобы хранить аналитические данные
type TGtoken struct {
	Token           string
	LocalHostPort   string
	Password        string
	PortWeb         string
	TelegramBotLink string
	updateTimeSec   string
}

type Data struct {
	Admin  bool
	Anonim bool

	NoteStatus string

	DomainName string

	Warning string

	TitleIcon string
	Title     string
	Logo      string

	HeadPhoto1       string
	HomePhotoLink1   string
	HomeTextTitle1   string
	HomeText1        string
	HeadPhoto2       string
	HomePhotoLink2   string
	HomeTextTitle2   string
	HomeText2        string
	HeadPhoto3       string
	HomePhotoLink3   string
	HomeTextTitle3   string
	HomeText3        string
	HeadPhoto4       string
	HomePhotoLink4   string
	HomeTextTitle4   string
	HomeText4        string
	HeadPhoto5       string
	HomePhotoLink5   string
	HomeTextTitle5   string
	HomeText5        string
	HeadPhoto6       string
	HomePhotoLink6   string
	HomeTextTitle6   string
	HomeText6        string
	HeadPhoto7       string
	HomePhotoLink7   string
	HomeTextTitle7   string
	HomeText7        string
	HeadPhoto8       string
	HomePhotoLink8   string
	HomeTextTitle8   string
	HomeText8        string
	HeadPhoto9       string
	HomePhotoLink9   string
	HomeTextTitle9   string
	HomeText9        string
	HeadPhoto10      string
	HomePhotoLink10  string
	HomeTextTitle10  string
	HomeText10       string
	HeadPhoto11      string
	HomePhotoLink11  string
	HomeTextTitle11  string
	HomeText11       string
	HeadPhoto12      string
	HomePhotoLink12  string
	HomeTextTitle12  string
	HomeText12       string
	HeadPhoto13      string
	HomePhotoLink13  string
	HomeTextTitle13  string
	HomeText13       string
	HeadPhoto14      string
	HomePhotoLink14  string
	HomeTextTitle14  string
	HomeText14       string
	HeadPhoto15      string
	HomePhotoLink15  string
	HomeTextTitle15  string
	HomeText15       string
	Header1          string
	Header2          string
	Header3          string
	Head_botton      string
	Head_bottonPhoto string

	SpecHead       string
	SpecFirstPhoto string
	SpecFirstHead  string
	SpecFirstBody1 string
	SpecFirstBody2 string
	SpecFirstBody3 string
	SpecFirstLink  string
	SpecSecPhoto   string
	SpecSecHead    string
	SpecSecBody1   string
	SpecSecBody2   string
	SpecSecBody3   string
	SpecSecLink    string
	SpecThirdPhoto string
	SpecThirdHead  string
	SpecThirdBody1 string
	SpecThirdBody2 string
	SpecThirdBody3 string
	SpecThirdLink  string

	Company1 string
	Company2 string
	Company3 string
	Company4 string
	Company5 string
	Company6 string

	GallDisc string

	ChoiseHead string
	Choise1    string
	Choise2    string
	Choise3    string

	Worker1      string
	WorkerPhoto1 string
	WorkerDeal1  string
	WorkerLink1  string
	Worker2      string
	WorkerPhoto2 string
	WorkerDeal2  string
	WorkerLink2  string
	Worker3      string
	WorkerPhoto3 string
	WorkerDeal3  string
	WorkerLink3  string
	Worker4      string
	WorkerPhoto4 string
	WorkerDeal4  string
	WorkerLink4  string
	Worker5      string
	WorkerPhoto5 string
	WorkerDeal5  string
	WorkerLink5  string
	Worker6      string
	WorkerPhoto6 string
	WorkerDeal6  string
	WorkerLink6  string
	Worker7      string
	WorkerPhoto7 string
	WorkerDeal7  string
	WorkerLink7  string

	ContactTelegram  string
	ContactInstsgram string
	ContactWhatsApp  string
	ContactYoutube   string
	ContactPhone1    string
	ContactPhone2    string
	ContactPhone3    string
	AddresStreet     string
	AddresCountry    string
	AddresLink       string
	Email1           string
	Email2           string

	AdminPass string
}

var dataSialum = Data{

	Admin:  false,
	Anonim: true,

	NoteStatus: "on",

	DomainName: createBot().LocalHostPort,

	Warning: "",

	TitleIcon: "https://i.postimg.cc/KY25ybqM/logo-title.png",
	Title:     "Sialum",

	SpecFirstPhoto: "https://i.postimg.cc/jd5M270L/img-06.jpg",
	SpecFirstHead:  "Владельцам квартир, домов, кафе.",
	SpecFirstBody1: "· Посмотрите реализацию ваших идей.",
	SpecFirstBody2: "· Получите правильную и удобную планировку.",
	SpecFirstBody3: "· Получите необходимую документацию для строителей.",
	SpecFirstLink:  "https://t.me/",
	SpecSecPhoto:   "https://i.postimg.cc/qqmQQknJ/img-07.jpg",
	SpecSecHead:    "Строительным компаниям",
	SpecSecBody1:   "· Получите визуализацию проектируемого дома/микрорайона.",
	SpecSecBody2:   "· Как следствие, сократите время соглавования работ с клиентами..",
	SpecSecBody3:   "· Как следствие, повысьте конверсию продаж.",
	SpecSecLink:    "https://t.me/",
	SpecThirdPhoto: "https://i.postimg.cc/wxFW9krC/img-08.jpg",
	SpecThirdHead:  "Дизайнерам и архитекторам",
	SpecThirdBody1: "· С нами - быстрее согласовывайте и закрывать свои проекты.",
	SpecThirdBody2: "· С нами - наработайте большую клиентскую базу.",
	SpecThirdBody3: "· Выполняйте одновременно большее количество проектов и получайте больше прибыли.",
	SpecThirdLink:  "https://t.me/",

	GallDisc:         "Более 3700 материалов из 40 проектов",
	Header2:          "Студия визуализации",
	Header3:          "интерьеров и экстерьеров",
	Header1:          "Sialum",
	Head_botton:      createBot().TelegramBotLink,
	Head_bottonPhoto: "https://i.postimg.cc/7PQgBtdZ/premium-features-3.png",

	HeadPhoto1:     "https://i.postimg.cc/BQTsnQYZ/11-500.jpg",
	HomePhotoLink1: "127.0.0.1:8080",
	HomeTextTitle1: "Салон красоты",
	HomeText1:      "Нюансы подготовки помещения под парикмахерскую часто играют решающую роль при создании комфортного рабочегго процесса мастера",

	HeadPhoto2:     "https://i.postimg.cc/ZKSt2zTG/110-500.jpg",
	HomePhotoLink2: "127.0.0.1:8080",
	HomeTextTitle2: "Конференс-зал",
	HomeText2:      "Чтобы общение всегда проходило плодотворно в первую очедь неоюходимо позаботиться об освещении помещения и вентиляции.",

	HeadPhoto3:     "https://i.postimg.cc/PrQskVJP/111-500.jpg",
	HomePhotoLink3: "127.0.0.1:8080",
	HomeTextTitle3: "Кухня ресторана",
	HomeText3:      "Недавно франчайзи одной очень известной сети ресторанов раказали у нас дизайн рабочего пространства поваров.",

	HeadPhoto4:     "https://i.postimg.cc/QMyZyBtm/12-500.jpg",
	HomePhotoLink4: "127.0.0.1:8080",
	HomeTextTitle4: "Уголок ожидания",
	HomeText4:      "Сделали дизайн зеленого уголка для ожидания клиентов парикмахерской. Оптимизация трпфика в маленьком помещении с помощью планировки",

	HeadPhoto5:     "https://i.postimg.cc/DyknhbcY/13-5050.jpg",
	HomePhotoLink5: "127.0.0.1:8080",
	HomeTextTitle5: "Просторное кафе",
	HomeText5:      "Что создает чувство уюта и теплую атмосферу в помещении ? Конечно это теплые оттенки и природный материал.",

	HeadPhoto6:     "https://i.postimg.cc/wMJzJ69V/14-500.jpg",
	HomePhotoLink6: "127.0.0.1:8080",
	HomeTextTitle6: "Просторное кафе",
	HomeText6:      "Что создает чувство уюта и теплую атмосферу в помещении ? Конечно это теплые оттенки и природный материал.",

	HeadPhoto7:     "https://i.postimg.cc/dVJFkPp2/15-5050.jpg",
	HomePhotoLink7: "127.0.0.1:8080",
	HomeTextTitle7: "Просторное кафе",
	HomeText7:      "Что создает чувство уюта и теплую атмосферу в помещении ? Конечно это теплые оттенки и природный материал.",

	HeadPhoto8:     "https://i.postimg.cc/MHykHdY3/16-500.jpg",
	HomePhotoLink8: "127.0.0.1:8080",
	HomeTextTitle8: "Просторное кафе",
	HomeText8:      "Что создает чувство уюта и теплую атмосферу в помещении ? Конечно это теплые оттенки и природный материал.",

	HeadPhoto9:     "https://i.postimg.cc/0jgR9L7J/17-500.jpg",
	HomePhotoLink9: "127.0.0.1:8080",
	HomeTextTitle9: "Просторное кафе",
	HomeText9:      "Что создает чувство уюта и теплую атмосферу в помещении ? Конечно это теплые оттенки и природный материал.",

	HeadPhoto10:     "https://i.postimg.cc/wvRY86Rp/18-50-50.jpg",
	HomePhotoLink10: "127.0.0.1:8080",
	HomeTextTitle10: "Просторное кафе",
	HomeText10:      "Что создает чувство уюта и теплую атмосферу в помещении ? Конечно это теплые оттенки и природный материал.",

	HeadPhoto11:     "https://i.postimg.cc/RZGBwpxj/19-500.jpg",
	HomePhotoLink11: "127.0.0.1:8080",
	HomeTextTitle11: "Просторное кафе",
	HomeText11:      "Что создает чувство уюта и теплую атмосферу в помещении ? Конечно это теплые оттенки и природный материал.",

	HeadPhoto12:     "https://i.postimg.cc/ZKSt2zTG/110-500.jpg",
	HomePhotoLink12: "127.0.0.1:8080",
	HomeTextTitle12: "Просторное кафе",
	HomeText12:      "Что создает чувство уюта и теплую атмосферу в помещении ? Конечно это теплые оттенки и природный материал.",

	HeadPhoto13:     "https://i.postimg.cc/W43P8m2t/15-5050-2.jpg",
	HomePhotoLink13: "127.0.0.1:8080",
	HomeTextTitle13: "Просторное кафе",
	HomeText13:      "Что создает чувство уюта и теплую атмосферу в помещении ? Конечно это теплые оттенки и природный материал.",

	HeadPhoto14:     "https://i.postimg.cc/05JPhC7w/13-5050-2.jpg",
	HomePhotoLink14: "127.0.0.1:8080",
	HomeTextTitle14: "Просторное кафе",
	HomeText14:      "Что создает чувство уюта и теплую атмосферу в помещении ? Конечно это теплые оттенки и природный материал.",

	HeadPhoto15:     "https://i.postimg.cc/NjGhG82b/111-500-2.jpg",
	HomePhotoLink15: "127.0.0.1:8080",
	HomeTextTitle15: "Просторное кафе",
	HomeText15:      "Что создает чувство уюта и теплую атмосферу в помещении ? Конечно это теплые оттенки и природный материал.",

	Company1: "https://i.postimg.cc/QxhNnrMb/logo-1.png",
	Company2: "https://i.postimg.cc/qv7J4Xy6/logo-2.png",
	Company3: "https://i.postimg.cc/h49PrGKv/logo-3.png",
	Company4: "https://i.postimg.cc/zX7zsPb1/logo-4.png",
	Company5: "https://i.postimg.cc/qv7J4Xy6/logo-2.png",
	Company6: "https://i.postimg.cc/cCMx4FFw/logo-5.png",

	ChoiseHead:   "Почему нас выбирают?",
	Choise1:      "· Сроки выполнения от 2-х дней!",
	Choise2:      "· 7 специалистов и 6 лет в индустрии. ",
	Choise3:      "· Стоимость визуализации от 400 р/м2",
	Worker1:      "Егоров Андрей",
	WorkerPhoto1: "https://i.postimg.cc/pXfmSTCp/team-04.jpg",
	WorkerDeal1:  "Проект-менеджер",

	Worker2:      "Каримова Наталья",
	WorkerPhoto2: "https://i.postimg.cc/SNFXV8WB/team-05.jpg",
	WorkerDeal2:  "Старший дизайнер",

	Worker3:      "Свиридов Николай",
	WorkerPhoto3: "https://i.postimg.cc/0yZb8BLv/team-03.jpg",
	WorkerDeal3:  "Дизайнер-проектировщик",

	Worker4:      "Петров Антон",
	WorkerPhoto4: "https://i.postimg.cc/fLSLfTkW/team-07.jpg",
	WorkerDeal4:  "Художник по свету",

	Worker5:      "Наумова Елена",
	WorkerPhoto5: "https://i.postimg.cc/wBv33kpm/team-02.jpg",
	WorkerDeal5:  "Контент-визор",

	Worker6:      "Парфенова Евгения",
	WorkerPhoto6: "https://i.postimg.cc/DzM0GbvY/team-06.jpg",
	WorkerDeal6:  "Архитектор",

	Worker7:      "Скворцова Карина",
	WorkerPhoto7: "https://i.postimg.cc/mkQgKXQ0/team-08.jpg",
	WorkerDeal7:  "3D-художник",

	ContactTelegram:  "https://www.telegram.com",
	ContactInstsgram: "https://www.instagram.com",
	ContactWhatsApp:  "https://",
	ContactYoutube:   "https://",
	ContactPhone1:    "+7 999 767 76 85",
	ContactPhone2:    "+7 943 209 00 00",
	ContactPhone3:    "+7 077 655 44 55",
	Email1:           "sialuminfo@sialum.com",
	Email2:           "sialumsupport@sialum.com",

	AddresStreet:  "Пролетпрский пр., 9 к4",
	AddresCountry: "Москва - Россия",
	AddresLink:    "https://yandex.ru/maps/?rtext=~55.733836%2C37.588134",

	AdminPass: createBot().Password, // 29c9=a0a0a842-41231=862-377=f78-2e4bcf15-9-166=f9b60-0cc7ce14=d1ee6=128078df
}

//
//
//
//
//
//
//
//
//
//
//
//
//

//
//
//
//
//
//
//                 HOME PAGE
//
func Home(w http.ResponseWriter, r *http.Request) {

	ipuser := ReadUserIP(r)

	// запись времени события
	now := time.Now()
	nanos := now.UnixNano()
	yy, _, dd := time.Now().UTC().Date()
	hh, mm, secc := time.Now().UTC().Clock()
	mon := now.Format("20060102150405")
	mon = mon[4:6]
	monInt, _ := strconv.Atoi(mon)
	id := now.Format("20060102150405") + strconv.FormatInt(nanos, 10)[0:6]
	nanoInt, _ := strconv.Atoi(strconv.FormatInt(nanos, 10)[0:6])

	var Nota = Notif{
		Ip:     ipuser,
		Year:   yy,
		Month:  monInt,
		Day:    dd,
		Hour:   hh,
		Minute: mm,
		Secund: secc,
		Nano:   nanoInt,
		Deal:   "Using HomePage",
	}
	//читаем БД с уведомлениями
	fileN, _ := ioutil.ReadFile("data/NotificationData.json")
	json.Unmarshal([]byte(fileN), &Notifiations)
	//Загружаем уведомления в БД онлайн (не в джсон)
	Notifiations.Notification = append(Notifiations.Notification, Nota)
	fmt.Println("NOTIFICATION\n№ " + id)
	//загрузка уведомления обратно в файл уведомления  (в джсон)
	fileN, _ = json.MarshalIndent(Notifiations, "", " ")
	_ = ioutil.WriteFile("data/NotificationData.json", fileN, 0644)
	//
	//
	//
	//
	if dataSialum.NoteStatus == "on" {
		x := createBot().Token
		bot, err := tgbotapi.NewBotAPI(x)
		if err != nil {
			log.Panic(err)
		}
		var VerifData = RouteVerifUsers{
			ChatID: []string{},
		}
		fileV, _ := ioutil.ReadFile("data/VrifUsersID.json")
		json.Unmarshal([]byte(fileV), &VerifData)
		for _, aidee := range VerifData.ChatID {
			r, _ := strconv.ParseInt(aidee, 10, 64)
			if len(VerifData.ChatID) > 0 {
				for _, checkID := range VerifData.ChatID {
					if checkID == strconv.FormatInt(r, 10) {

						answ := "NOTIFICATION\n№ " + id
						answ = answ + "\n" + Notifiations.Notification[len(Notifiations.Notification)-1].Deal + "\n" + ipuser
						msg := tgbotapi.NewMessage(r, answ)
						bot.Send(msg)

					}
				}
			}
		}

	}

	//очищаем уведомления администратора
	dataSialum.Warning = ""

	///READ JSON для получения данных о наполнеии сайта
	file, _ := ioutil.ReadFile("data/dataSialum.json")
	json.Unmarshal([]byte(file), &dataSialum)

	///SHOW показываем сайт
	var tmpls = template.Must(template.ParseFiles("templates/home.html"))
	if err := tmpls.ExecuteTemplate(w, "home.html", dataSialum); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//
//
//
//
//
//
//   MIRROR HOME PAGE MIRROR FOR CHECK A CHANGED
//    ELEMENT WIYHOUT REGISTR NOTIF OR ANALITS
//
func Mirror(w http.ResponseWriter, r *http.Request) {

	//очищаем уведомления администратора
	dataSialum.Warning = ""

	///READ JSON для получения данных о наполнеии сайта
	file, _ := ioutil.ReadFile("data/dataSialum.json")
	json.Unmarshal([]byte(file), &dataSialum)

	///SHOW показываем сайт
	var tmpls = template.Must(template.ParseFiles("templates/home.html"))
	if err := tmpls.ExecuteTemplate(w, "home.html", dataSialum); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//
//
//
///

//

//
//
//
//
//
///
//
//
//

//                 SPECIAL PAGE
//
func special(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/special.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tmpl.ExecuteTemplate(w, "special", dataSialum)

}

//
//
//
//
//
//
//
//
//
//
//
//
func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

//
//
//
//
//
///
//
//
//
///
//
//
//
//
//
//

func main() {

	//  SSSSIIIITTTEEEEE SSEERRRRVVEEERRR  //
	go func() {

		port := createBot().PortWeb //os.Getenv("PORT")

		r := mux.NewRouter()

		r.HandleFunc("/", Home)

		r.HandleFunc("/8ADF14ADC5E6CEBB2Bd6B9CF7FF70B5F9Bb56B1C59E8A1C0BEF69d0FCE36128DgEBC", Mirror)

		r.PathPrefix("").Handler(http.StripPrefix("",
			http.FileServer(http.Dir("templates/"))))

		http.Handle("css/template-style.css", http.StripPrefix("", http.FileServer(http.Dir("./template-style.css"))))
		http.Handle("/", r)
		log.Fatalln(http.ListenAndServe(":"+port, nil))
	}() /////////////////////////////////////
	//
	//
	//
	// СОЗДВЕМ ДЕФОЛТНУЮ БД
	fileC, _ := json.MarshalIndent(dataSialum, "", " ")
	_ = ioutil.WriteFile("data/defaultDataSialum.json", fileC, 0644)
	//
	//
	//
	//
	//

	//
	//

	////////////////////////////////////////////////////////////////////////////
	//bot///////////////////////////////////////////////////////////////////////
	bot, err := tgbotapi.NewBotAPI(createBot().Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	updStr, err := strconv.Atoi(createBot().updateTimeSec)
	if err != nil {
		log.Panic(err)
	}
	u.Timeout = updStr

	updates, err := bot.GetUpdatesChan(u)

	var VerifData = RouteVerifUsers{
		ChatID: []string{},
	}
	//
	//
	//
	//
	//
	//

	//
	//
	//
	//
	//
	//
	//
	//
	//
	//читаем БД с уведомлениями
	fileN, _ := ioutil.ReadFile("data/NotificationData.json")
	json.Unmarshal([]byte(fileN), &Notifiations)
	//
	//
	//
	//
	//
	//
	//
	//         SCRIPTS
	// чтобы хранить скрипт диалога
	// подготовка к циклу
	var dialogS = dialogSample{
		ParrentID: "",
		User:      "",
		Bot:       "",
	}

	var ScriptData = script{
		Scripts: make(map[string]dialogSample),
	}

	StepDialog := map[string]string{}
	dialogIDbuffer := ""

	//
	//

	//
	//
	//
	//
	//
	//
	//
	//
	// пароль и статус верфикации

	verification := false
	Exited := false
	//
	//
	//
	//
	//
	//
	//
	//

	answ := "Бот может:\n\n\n1. Показать актуальные данные на сайте.\nДля этого есть два варианта команды:"
	answ = answ + "\n   1. 'conf all' - запросить всю\n   информацию.\n   2. Запросить один раздел\n   по номеру или\n   имени из списка:\n"
	answ = answ + "      'conf 0' или 'title'\n      'conf 1' или 'head'"
	answ = answ + "\n      'conf 2' или 'special'\n      'conf 3' или 'gallery'"
	answ = answ + "\n      'conf 4' или 'company'\n      'conf 5' или 'choice'\n      'conf 6' или 'workers'\n      'conf 7' или 'contacts'"
	answ = answ + "\n\n\n2. Заменить данные на сайте.\nДля этого используйте кодовую команду:"
	answ = answ + "\n   'XXX=' - код элемента раздела.\n   'Текст' - данные для замены.\n   Примеры:\n      '307=Это новые данные'\n      '102=https://i.site.com/T/1.jpg'"
	answ = answ + "\n\n\n3. Сбросить все изменения сайта.\n   'set default conf' - команда сброса"
	answ = answ + "\n\n\n4. Уведомлять о посещении сайта.\nДля управления используйте команды:"
	answ = answ + "\n   'shut up' - отключить\n   'talk all' - включить"
	answ = answ + "\n\n\n5. Записать пару вопрос/ответ. Используйте команду:"
	answ = answ + "\n   'nscr' - команда для начала\nЭтапы записи пары:"
	answ = answ + "\n   1. Введите идентификационный номер пары, от которой вы хотите продолжать диалог с пользователем.\nЕсли вы хотите, чтобы эта пара срабатывала как начало новой линии диалога, то введите '0'."
	answ = answ + "\n   2. Введите сообщение, которое ожидаете от пользователя.\n   3. Введите сообщение, которое ответит бот."
	answ = answ + "\n   4. Всё! Пара создана. Бот создаст для неё уникальный номер, к которому вы сможете позже привязать развитие диалога."
	answ = answ + "\n\n\n6. Удалить пару вопрос/ответ.\n    'dscrX' - замените X на уникальный номер."
	answ = answ + "\n\n\n7. Просмотреть все пары вопрос/ответ. Используйте команду:"
	answ = answ + "\n   'conf scr' - команда для просмотра."
	answ = answ + "\n\n\n8. Установить новый пароль. Используйте команду:"
	answ = answ + "\n   'npasX' - замените X на новый пароль."
	answ = answ + "\n\n\n9. Запросить статистику за день. Используйте команду:"
	answ = answ + "\n   'uday' - команда для посмотра посещаемости."
	answ = answ + "\n\n\n10. Выйти из режима администратора:\n   'close' - команда выхода."

	botmenu := answ
	//
	//
	//
	//
	//
	//
	//    ЦИКЛ ОБНОВЛЕНИЙ
	for update := range updates {
		fmt.Printf("\n\n\n\n\n\n%T", update)

		//проверка статуса данного чата и того
		//был ли в нем введен пароль
		file, _ := ioutil.ReadFile("data/VrifUsersID.json")
		json.Unmarshal([]byte(file), &VerifData)
		if len(VerifData.ChatID) > 0 {
			for _, checkID := range VerifData.ChatID {
				if checkID == strconv.FormatInt(update.Message.Chat.ID, 10) {
					if Exited != true {
						verification = true
					}

				}
			}
		}

		//сообщение пользователя
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		//когда пользователь ввел пароль- создание пользователя по айди
		if msg.Text == dataSialum.AdminPass {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, botmenu)
			bot.Send(msg)

			verification = true
			Exited = false
			var VerifData = RouteVerifUsers{
				ChatID: append(VerifData.ChatID, strconv.FormatInt(update.Message.Chat.ID, 10)),
			}
			//удаление повторяющихся айди
			if len(VerifData.ChatID) > 0 {
				VerifData.ChatID = unique(VerifData.ChatID)
			}
			fmt.Println("Повторяющиеся айди удалены:\n", VerifData.ChatID)
			//запись айди пользователя в джсон
			fileY, _ := json.MarshalIndent(VerifData, "", " ")
			_ = ioutil.WriteFile("data/VrifUsersID.json", fileY, 0644)

		}

		//выйти из настроек администратора
		if msg.Text == "close" && verification == true {
			verification = false
			newVerifList := []string{}
			for _, idd := range VerifData.ChatID {
				if idd != strconv.FormatInt(update.Message.Chat.ID, 10) {
					newVerifList = append(newVerifList, idd)
				}
			}
			VerifData.ChatID = newVerifList
			fileY, _ := json.MarshalIndent(VerifData, "", " ")
			_ = ioutil.WriteFile("data/VrifUsersID.json", fileY, 0644)
			answ := "see you soon"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, answ)
			bot.Send(msg)
			Exited = true
		}
		//
		//
		//
		//
		//
		//
		//
		//
		//
		//
		//
		//
		//
		//
		//

		//ответ по скрипту из файла
		file, _ = ioutil.ReadFile("data/ScriptData.json")
		json.Unmarshal([]byte(file), &ScriptData)

		for nameLine := range ScriptData.Scripts {
			if ScriptData.Scripts[nameLine].ParrentID == "0" {
				if msg.Text == ScriptData.Scripts[nameLine].User {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, ScriptData.Scripts[nameLine].Bot)
					bot.Send(msg)
					dialogIDbuffer = nameLine
				}
			} else {
				if ScriptData.Scripts[nameLine].ParrentID == dialogIDbuffer {
					if msg.Text == ScriptData.Scripts[nameLine].User {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, ScriptData.Scripts[nameLine].Bot)
						bot.Send(msg)
						dialogIDbuffer = nameLine
					}
				}
			}
		}
		//
		//
		//
		//
		//
		//

		//
		//
		//
		//
		//
		//
		//
		//

		//разговор с верифицированным пользователем
		if verification == true {

			//рудимент
			userText := msg.Text

			//удалить всё
			if userText == "set default conf" {
				if _, err := os.Stat("data/dataSialum.json"); err == nil {
					e := os.Remove("data/dataSialum.json")
					if e != nil {
						log.Fatal(e)
					}
				}
				if _, err := os.Stat("data/NotificationData.json"); err == nil {
					e := os.Remove("data/NotificationData.json")
					if e != nil {
						log.Fatal(e)
					}
				}

			}

			//CREATE DIALOG//////////////////////////
			if userText == "nscr" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Parent ID: ")
				bot.Send(msg)
				StepDialog["Parent ID"] = "name"

			} //CREATE DIALOG//////////////////////////
			if len(StepDialog) == 2 {
				if userText == "not nscr" {
					StepDialog = map[string]string{}
				} else {
					StepDialog["Parent ID"] = userText
					StepDialog["User"] = "userText"
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "User: ")
					bot.Send(msg)

				}
			} //CREATE DIALOG//////////////////////////
			if len(StepDialog) == 4 {
				if userText == "not nscr" {
					StepDialog = map[string]string{}
				} else {
					StepDialog["User"] = userText
					StepDialog["Bot"] = "userText"
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Bot: ")
					bot.Send(msg)

				}
			} //CREATE DIALOG//////////////////////////
			if len(StepDialog) == 6 {
				if userText == "not nscr" {
					StepDialog = map[string]string{}
				} else {
					StepDialog["Bot"] = userText
					file, _ := ioutil.ReadFile("data/ScriptData.json")
					json.Unmarshal([]byte(file), &ScriptData)
					// установка лимита на сообщения
					if len(ScriptData.Scripts) < 500 {
						dialogS.User = StepDialog["User"]
						dialogS.Bot = StepDialog["Bot"]
						dialogS.ParrentID = StepDialog["Parent ID"]
						StepDialog["Name"] = time.Now().Format("20060102150405")
						ScriptData.Scripts[StepDialog["Name"]] = dialogS
						fmt.Println(ScriptData.Scripts)
						file, _ = json.MarshalIndent(ScriptData, "", " ")
						_ = ioutil.WriteFile("data/ScriptData.json", file, 0644)
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Done!\n\nID: ")
						bot.Send(msg)
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, StepDialog["Name"])
						bot.Send(msg)
						StepDialog = map[string]string{}
					} else {
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Превышен лимит количества сообщений")
						bot.Send(msg)
					}
				}
			} //CREATE DIALOG//////////////////////////
			if len(StepDialog) == 1 {
				StepDialog["1"] = "1"
			} //CREATE DIALOG//////////////////////////
			if len(StepDialog) == 3 {
				StepDialog["3"] = "3"
			} //CREATE DIALOG//////////////////////////
			if len(StepDialog) == 5 {
				StepDialog["5"] = "5"
			}
			//
			//
			//
			//
			//
			//
			//
			//
			//  просмотр сообщений
			if userText == "conf scr" {
				file, _ = ioutil.ReadFile("data/ScriptData.json")
				json.Unmarshal([]byte(file), &ScriptData)
				for nameLine := range ScriptData.Scripts {
					answ := "Parent ID: " + ScriptData.Scripts[nameLine].ParrentID
					answ = answ + "\n\n   User:\n   " + ScriptData.Scripts[nameLine].User
					answ = answ + "\n\n   Bot:\n   " + ScriptData.Scripts[nameLine].Bot
					answ = answ + "\n\n   ID:\n   " + nameLine
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, answ)
					bot.Send(msg)
				}
			}
			//  удалить сообщение
			if len(userText) > 3 {
				if userText[0:4] == "dscr" {
					if len(userText[4:len(userText)]) == len("20060102150405") {
						file, _ = ioutil.ReadFile("data/ScriptData.json")
						json.Unmarshal([]byte(file), ScriptData)
						delete(ScriptData.Scripts, userText[4:len(userText)])
						file, _ = json.MarshalIndent(ScriptData, "", " ")
						_ = ioutil.WriteFile("data/ScriptData.json", file, 0644)
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Done!")
						bot.Send(msg)
					}
				}
			}
			//
			// заменить пароль администратора
			if len(userText) > 3 {
				if userText[0:4] == "npas" {
					dataSialum.AdminPass = userText[4:len(userText)]
					file, _ = json.MarshalIndent(dataSialum, "", " ")
					_ = ioutil.WriteFile("data/dataSialum.json", file, 0644)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Done!")
					bot.Send(msg)
				}
			}
			//
			//
			//
			//
			//
			//
			//
			//

			// Игнорирование пустых обновлений
			if update.Message == nil {
				continue
			}
			//
			//
			//
			//
			//
			//      Аналитика
			if userText == "uday" {
				if len(Notifiations.Notification) > 0 {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "loading")
					bot.Send(msg)
					analiticList := []Notif{}
					now := time.Now()
					yy, _, dd := time.Now().UTC().Date()
					mon := now.Format("20060102150405")
					mon = mon[4:6]
					monInt, _ := strconv.Atoi(mon)
					//читаем БД с уведомлениями
					fileN, _ := ioutil.ReadFile("data/NotificationData.json")
					json.Unmarshal([]byte(fileN), &Notifiations)
					//Отбераем уведомления за последний день (не 24 часа)
					for _, dayy := range Notifiations.Notification {
						if dayy.Year == yy && dayy.Month == monInt && dayy.Day == dd && dayy.Deal == "Using HomePage" {
							analiticList = append(analiticList, dayy)
						}
					}

					XValues := []float64{0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0,
						9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0,
						17.0, 18.0, 19.0, 20.0, 21.0, 22.0, 23.0} //Time
					preYValues := map[float64][]Notif{}
					for _, ana := range analiticList {
						for _, XV := range XValues {
							if float64(ana.Hour) == XV {
								preYValues[XV] = append(preYValues[XV], ana)
							}
						}
					}

					YValues := map[float64]float64{}
					for key, _ := range preYValues {

						YValues[key] = float64(len(preYValues[key]))
					}

					ggX := []float64{0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0,
						8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0,
						16.0, 17.0, 18.0, 19.0, 20.0, 21.0, 22.0, 23.0}
					ggY := []float64{YValues[0.0], YValues[1.0], YValues[2.0], YValues[3.0], YValues[4.0], YValues[5.0],
						YValues[6.0], YValues[7.0], YValues[8.0], YValues[9.0], YValues[10.0], YValues[11.0],
						YValues[12.0], YValues[13.0], YValues[14.0], YValues[15.0], YValues[16.0], YValues[17.0],
						YValues[18.0], YValues[19.0], YValues[20.0], YValues[21.0], YValues[22.0], YValues[23.0],
					}
					//ggX = ggX[0:now.Hour()]
					//ggY = ggY[0:now.Hour()]

					graph := chart.Chart{
						XAxis: chart.XAxis{
							Name: "Time UTC at " + now.Format("2006.01.02"),
						},
						YAxis: chart.YAxis{
							Name: "Using Home Page",
						},

						Series: []chart.Series{
							chart.ContinuousSeries{
								Style: chart.Style{
									StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
									FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
								},
								XValues: ggX,
								YValues: ggY}},
					}

					f, _ := os.Create("uday.png")
					defer f.Close()
					graph.Render(chart.PNG, f)
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "photo created")
					bot.Send(msg)

					//
					//
					f, _ = os.Open("uday.png")
					reader := tgbotapi.FileReader{Name: "uday.png", Reader: f, Size: -1}

					msgv := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, reader)
					msgv.Caption = "Using Home Page at " + now.Format("2006.01.02")
					bot.Send(msgv)

				}
			}
			//
			//
			//
			//
			//
			//
			//
			//

			//отчет о диалоге в лог сервера
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			//ответ на сообщение с пересылкой эхо
			//msg.ReplyToMessageID = update.Message.MessageID
			//bot.Send(msg)

			//  МЕНЮ БОТА
			if userText == "help" || userText == "commands" || userText == "bot" || userText == "hi" || userText == "menu" {

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, botmenu)
				bot.Send(msg)
			}

			// ВКЛЮЧИТЬ УВЕДОМЛЕНИЯ
			if userText == "talk all" {
				dataSialum.NoteStatus = "on"
				file, _ := json.MarshalIndent(dataSialum, "", " ")
				_ = ioutil.WriteFile("data/dataSialum.json", file, 0644)
				answ := "Notifications on"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, answ)
				bot.Send(msg)
			}
			// ОТКЛЮЧИТЬ УВЕДОМЛЕНИЯ
			if userText == "shut up" {
				dataSialum.NoteStatus = "off"
				answ := "Notifications off"
				file, _ := json.MarshalIndent(dataSialum, "", " ")
				_ = ioutil.WriteFile("data/dataSialum.json", file, 0644)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, answ)
				bot.Send(msg)
			}

			// ВЕРНУТЬСЯ ИЗНАЧАЛЬНОЙ БД
			if userText == "set default conf" {
				file, _ := ioutil.ReadFile("data/defaultDataSialum.json")
				json.Unmarshal([]byte(file), &dataSialum)
				file, _ = json.MarshalIndent(dataSialum, "", " ")
				_ = ioutil.WriteFile("data/dataSialum.json", file, 0644)
				answ := "Data Erised"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, answ)
				bot.Send(msg)
			}

			//демонстрация контента на сайте
			if userText == "conf all" || userText == "conf 0" || userText == "title" {
				answ := "000\n   Title Icon:  \n   ..." + dataSialum.TitleIcon
				answ = answ + "\n\n001\n   Title Head Name:  \n   " + dataSialum.Title
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, answ)
				bot.Send(msg)
			}

			if userText == "conf all" || userText == "conf 1" || userText == "head" {
				answ := "100\n   Head Text 1:  \n   " + dataSialum.Header1
				answ = answ + "\n\n101\n   Head Text 2:  \n   " + dataSialum.Header2
				answ = answ + "\n\n102\n   Head Text 3:  \n   " + dataSialum.Header3
				answ = answ + "\n\n103\n   Head Button Link:  \n   " + dataSialum.Head_botton
				answ = answ + "\n\n104\n   Head Button Photo:  \n   " + dataSialum.Head_bottonPhoto
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, answ)
				bot.Send(msg)
			}

			//демонстрация контента на сайте
			if userText == "conf all" || userText == "conf 2" || userText == "special" {
				answ := "200\n   Special Offer 1 Photo :  \n   ..." + dataSialum.SpecFirstPhoto
				answ = answ + "\n\n201\n   Special Offer 1 Name :  \n   " + dataSialum.SpecFirstHead
				answ = answ + "\n\n202\n   Special Offer 1 Text 1:  \n   " + dataSialum.SpecFirstBody1
				answ = answ + "\n\n203\n   Special Offer 1 Text 2:  \n   " + dataSialum.SpecFirstBody2
				answ = answ + "\n\n204\n   Special Offer 1 Text 3:  \n   " + dataSialum.SpecFirstBody3
				answ = answ + "\n\n205\n   Special Offer 1 Link:  \n   ..." + dataSialum.SpecFirstLink
				answ = answ + "\n\n206\n   Special Offer 2 Photo :  \n   ..." + dataSialum.SpecSecPhoto
				answ = answ + "\n\n\n207\n   Special Offer 2 Name :  \n   " + dataSialum.SpecSecHead
				answ = answ + "\n\n208\n   Special Offer 2 Text 1:  \n   " + dataSialum.SpecSecBody1
				answ = answ + "\n\n209\n   Special Offer 2 Text 2:  \n   " + dataSialum.SpecSecBody2
				answ = answ + "\n\n210\n   Special Offer 2 Text 3:  \n   " + dataSialum.SpecSecBody3
				answ = answ + "\n\n211\n   Special Offer 2 Link:  \n   ..." + dataSialum.SpecSecLink
				answ = answ + "\n\n212\n   Special Offer 3 Photo :  \n   ..." + dataSialum.SpecThirdPhoto
				answ = answ + "\n\n\n213\n   Special Offer 3 Name :  \n   " + dataSialum.SpecThirdHead
				answ = answ + "\n\n214\n   Special Offer 3 Text 1:  \n   " + dataSialum.SpecThirdBody1
				answ = answ + "\n\n215\n   Special Offer 3 Text 2:  \n   " + dataSialum.SpecThirdBody2
				answ = answ + "\n\n216\n   Special Offer 3 Text 3:  \n   " + dataSialum.SpecThirdBody3
				answ = answ + "\n\n217\n   Special Offer 3 Link:  \n   ..." + dataSialum.SpecThirdLink
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, answ)
				bot.Send(msg)
			}

			//демонстрация контента на сайте
			if userText == "conf all" || userText == "conf 3" || userText == "gallery" {
				answ := "300\n   Gallery Description:  \n   ..." + dataSialum.GallDisc
				answ = answ + "\n\n301\n   Gallery Photo image 1:  \n   ..." + dataSialum.HeadPhoto1
				answ = answ + "\n\n302\n   Gallery Photo link 1:  \n   ..." + dataSialum.HomePhotoLink1
				answ = answ + "\n\n303\n   Gallery Photo Text Title 1:  \n   " + dataSialum.HomeTextTitle1
				answ = answ + "\n\n304\n   Gallery Photo Text 1:  \n   " + dataSialum.HomeText1
				answ = answ + "\n\n305\n   Gallery Photo image 2:  \n   ..." + dataSialum.HeadPhoto2
				answ = answ + "\n\n306\n   Gallery Photo link 2:  \n   ..." + dataSialum.HomePhotoLink2
				answ = answ + "\n\n307\n   Gallery Photo Text Title 2:  \n   " + dataSialum.HomeTextTitle2
				answ = answ + "\n\n308\n   Gallery Photo Text 2:  \n   " + dataSialum.HomeText2
				answ = answ + "\n\n309\n   Gallery Photo image 3:  \n   ..." + dataSialum.HeadPhoto3
				answ = answ + "\n\n310\n   Gallery Photo link 3:  \n   ..." + dataSialum.HomePhotoLink3
				answ = answ + "\n\n311\n   Gallery Photo Text Title 3:  \n   " + dataSialum.HomeTextTitle3
				answ = answ + "\n\n312\n   Gallery Photo Text 3:  \n   " + dataSialum.HomeText3
				answ = answ + "\n\n313\n   Gallery Photo image 4:  \n   ..." + dataSialum.HeadPhoto4
				answ = answ + "\n\n314\n   Gallery Photo link 4:  \n   ..." + dataSialum.HomePhotoLink4
				answ = answ + "\n\n315\n   Gallery Photo Text Title 4:  \n   " + dataSialum.HomeTextTitle4
				answ = answ + "\n\n316\n   Gallery Photo Text 4:  \n   " + dataSialum.HomeText4
				answ = answ + "\n\n317\n   Gallery Photo image 5:  \n   ..." + dataSialum.HeadPhoto5
				answ = answ + "\n\n318\n   Gallery Photo link 5:  \n   ..." + dataSialum.HomePhotoLink5
				answ = answ + "\n\n319\n   Gallery Photo Text Title 5:  \n   " + dataSialum.HomeTextTitle5
				answ = answ + "\n\n320\n   Gallery Photo Text 5:  \n   " + dataSialum.HomeText5
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, answ)
				bot.Send(msg)
			}

			if userText == "conf all" || userText == "conf 3" || userText == "gallery" {
				answ := "\n\n321\n   Gallery Photo image 6:  \n   ..." + dataSialum.HeadPhoto6
				answ = answ + "\n\n322\n   Gallery Photo link 6:  \n   ..." + dataSialum.HomePhotoLink6
				answ = answ + "\n\n323\n   Gallery Photo Text Title 6:  \n   " + dataSialum.HomeTextTitle6
				answ = answ + "\n\n324\n   Gallery Photo Text 6:  \n   " + dataSialum.HomeText6
				answ = answ + "\n\n325\n   Gallery Photo image 7:  \n   ..." + dataSialum.HeadPhoto7
				answ = answ + "\n\n326\n   Gallery Photo link 7:  \n   ..." + dataSialum.HomePhotoLink7
				answ = answ + "\n\n327\n   Gallery Photo Text Title 7:  \n   " + dataSialum.HomeTextTitle7
				answ = answ + "\n\n328\n   Gallery Photo Text 7:  \n   " + dataSialum.HomeText7
				answ = answ + "\n\n329\n   Gallery Photo image 8:  \n   ..." + dataSialum.HeadPhoto8
				answ = answ + "\n\n330\n   Gallery Photo link 8:  \n   ..." + dataSialum.HomePhotoLink8
				answ = answ + "\n\n331\n   Gallery Photo Text Title 8:  \n   " + dataSialum.HomeTextTitle8
				answ = answ + "\n\n332\n   Gallery Photo Text 8:  \n   " + dataSialum.HomeText8
				answ = answ + "\n\n333\n   Gallery Photo image 9:  \n   ..." + dataSialum.HeadPhoto9
				answ = answ + "\n\n334\n   Gallery Photo link 9:  \n   ..." + dataSialum.HomePhotoLink9
				answ = answ + "\n\n335\n   Gallery Photo Text Title 9:  \n   " + dataSialum.HomeTextTitle9
				answ = answ + "\n\n336\n   Gallery Photo Text 9:  \n   " + dataSialum.HomeText9
				answ = answ + "\n\n337\n   Gallery Photo image 10:  \n   ..." + dataSialum.HeadPhoto10
				answ = answ + "\n\n338\n   Gallery Photo link 10:  \n   ..." + dataSialum.HomePhotoLink10
				answ = answ + "\n\n339\n   Gallery Photo Text Title 10:  \n   " + dataSialum.HomeTextTitle10
				answ = answ + "\n\n340\n   Gallery Photo Text 10:  \n   " + dataSialum.HomeText10
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, answ)
				bot.Send(msg)
			}

			if userText == "conf all" || userText == "conf 3" || userText == "gallery" {
				answ := "\n\n341\n   Gallery Photo image 11:  \n   ..." + dataSialum.HeadPhoto11
				answ = answ + "\n\n342\n   Gallery Photo link 11:  \n   ..." + dataSialum.HomePhotoLink11
				answ = answ + "\n\n343\n   Gallery Photo Text Title 11:  \n   " + dataSialum.HomeTextTitle11
				answ = answ + "\n\n344\n   Gallery Photo Text 11:  \n   " + dataSialum.HomeText11
				answ = answ + "\n\n345\n   Gallery Photo image 12:  \n   ..." + dataSialum.HeadPhoto12
				answ = answ + "\n\n346\n   Gallery Photo link 12:  \n   ..." + dataSialum.HomePhotoLink12
				answ = answ + "\n\n347\n   Gallery Photo Text Title 12:  \n   " + dataSialum.HomeTextTitle12
				answ = answ + "\n\n348\n   Gallery Photo Text 12:  \n   " + dataSialum.HomeText12
				answ = answ + "\n\n349\n   Gallery Photo image 13:  \n   ..." + dataSialum.HeadPhoto13
				answ = answ + "\n\n350\n   Gallery Photo link 13:  \n   ..." + dataSialum.HomePhotoLink13
				answ = answ + "\n\n351\n   Gallery Photo Text Title 13:  \n   " + dataSialum.HomeTextTitle13
				answ = answ + "\n\n352\n   Gallery Photo Text 13:  \n   " + dataSialum.HomeText13
				answ = answ + "\n\n353\n   Gallery Photo image 14:  \n   ..." + dataSialum.HeadPhoto14
				answ = answ + "\n\n354\n   Gallery Photo link 14:  \n   ..." + dataSialum.HomePhotoLink14
				answ = answ + "\n\n355\n   Gallery Photo Text Title 14:  \n   " + dataSialum.HomeTextTitle14
				answ = answ + "\n\n356\n   Gallery Photo Text 14:  \n   " + dataSialum.HomeText14
				answ = answ + "\n\n357\n   Gallery Photo image 15:  \n   ..." + dataSialum.HeadPhoto15
				answ = answ + "\n\n358\n   Gallery Photo link 15:  \n   ..." + dataSialum.HomePhotoLink15
				answ = answ + "\n\n359\n   Gallery Photo Text Title 15:  \n   " + dataSialum.HomeTextTitle15
				answ = answ + "\n\n360\n   Gallery Photo Text 15:  \n   " + dataSialum.HomeText15
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, answ)
				bot.Send(msg)
			}

			//демонстрация контента на сайте
			if userText == "conf all" || userText == "conf 4" || userText == "company" {
				answ := "400\n   Company 1:  \n   ..." + dataSialum.Company1
				answ = answ + "\n\n401\n   Company 2:  \n   ..." + dataSialum.Company2
				answ = answ + "\n\n402\n   Company 3:  \n   ..." + dataSialum.Company3
				answ = answ + "\n\n403\n   Company 4:  \n   ..." + dataSialum.Company4
				answ = answ + "\n\n404\n   Company 5:  \n   ..." + dataSialum.Company3
				answ = answ + "\n\n405\n   Company 6:  \n   ..." + dataSialum.Company4
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, answ)
				bot.Send(msg)
			}

			//демонстрация контента на сайте
			if userText == "conf all" || userText == "conf 5" || userText == "choice" {
				answ := "500\n   Choice Head Name:  \n   " + dataSialum.ChoiseHead
				answ = answ + "\n\n501\n   Choice Text 1:  \n   " + dataSialum.Choise1
				answ = answ + "\n\n502\n   Choice Text 2:  \n   " + dataSialum.Choise2
				answ = answ + "\n\n503\n   Choice Text 3:  \n   " + dataSialum.Choise3
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, answ)
				bot.Send(msg)
			}

			if userText == "conf all" || userText == "conf 6" || userText == "workers" {
				answ := "600\n   Worker 1 Name:  \n   " + dataSialum.Worker1
				answ = answ + "\n\n601\n   Worker 1 deal:  \n   " + dataSialum.WorkerDeal1
				answ = answ + "\n\n602\n   Worker 1 photo:  \n   " + dataSialum.WorkerPhoto1
				answ = answ + "\n\n603\n   Worker 1 link:  \n   " + dataSialum.WorkerLink1
				answ = answ + "\n\n604\n   Worker 2 Name:  \n   " + dataSialum.Worker2
				answ = answ + "\n\n605\n   Worker 2 deal:  \n   " + dataSialum.WorkerDeal2
				answ = answ + "\n\n606\n   Worker 2 photo:  \n   " + dataSialum.WorkerPhoto2
				answ = answ + "\n\n607\n   Worker 2 link:  \n   " + dataSialum.WorkerLink2
				answ = answ + "\n\n608\n   Worker 3 Name:  \n   " + dataSialum.Worker3
				answ = answ + "\n\n609\n   Worker 3 deal:  \n   " + dataSialum.WorkerDeal3
				answ = answ + "\n\n610\n   Worker 3 photo:  \n   " + dataSialum.WorkerPhoto3
				answ = answ + "\n\n611\n   Worker 3 link:  \n   " + dataSialum.WorkerLink3
				answ = answ + "\n\n612\n   Worker 4 Name:  \n   " + dataSialum.Worker4
				answ = answ + "\n\n613\n   Worker 4 deal:  \n   " + dataSialum.WorkerDeal4
				answ = answ + "\n\n614\n   Worker 4 photo:  \n   " + dataSialum.WorkerPhoto4
				answ = answ + "\n\n615\n   Worker 4 link:  \n   " + dataSialum.WorkerLink4
				answ = answ + "\n\n616\n   Worker 5 Name:  \n   " + dataSialum.Worker5
				answ = answ + "\n\n617\n   Worker 5 deal:  \n   " + dataSialum.WorkerDeal5
				answ = answ + "\n\n618\n   Worker 5 photo:  \n   " + dataSialum.WorkerPhoto5
				answ = answ + "\n\n619\n   Worker 5 link:  \n   " + dataSialum.WorkerLink5
				answ = answ + "\n\n620\n   Worker 6 Name:  \n   " + dataSialum.Worker6
				answ = answ + "\n\n621\n   Worker 6 deal:  \n   " + dataSialum.WorkerDeal6
				answ = answ + "\n\n622\n   Worker 6 photo:  \n   " + dataSialum.WorkerPhoto6
				answ = answ + "\n\n623\n   Worker 6 link:  \n   " + dataSialum.WorkerLink6
				answ = answ + "\n\n624\n   Worker 7 Name:  \n   " + dataSialum.Worker7
				answ = answ + "\n\n625\n   Worker 7 deal:  \n   " + dataSialum.WorkerDeal7
				answ = answ + "\n\n626\n   Worker 7 photo:  \n   " + dataSialum.WorkerPhoto7
				answ = answ + "\n\n627\n   Worker 7 link:  \n   " + dataSialum.WorkerLink7
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, answ)
				bot.Send(msg)
			}

			//демонстрация контента на сайте
			if userText == "conf all" || userText == "conf 7" || userText == "contacts" {
				answ := "700\n   Contact Call 1:  \n   " + dataSialum.ContactPhone1
				answ = answ + "\n\n701\n   Contact Call 2:  \n   " + dataSialum.ContactPhone2
				answ = answ + "\n\n702\n   Contact Call 3:  \n   " + dataSialum.ContactPhone3
				answ = answ + "\n\n703\n   Contact Addres Street:  \n   " + dataSialum.AddresStreet
				answ = answ + "\n\n704\n   Contact Addres Country:  \n   " + dataSialum.AddresCountry
				answ = answ + "\n\n705\n   Contact Addres Link:  \n   ..." + dataSialum.AddresLink
				answ = answ + "\n\n706\n   Contact Addres e-mail 1:  \n   " + dataSialum.Email1
				answ = answ + "\n\n707\n   Contact Addres e-mail 2:  \n   " + dataSialum.Email2
				answ = answ + "\n\n\n\n708\n   Contact Telegram:  \n   ..." + dataSialum.ContactTelegram
				answ = answ + "\n\n709\n   Contact WhatsApp:  \n   ..." + dataSialum.ContactWhatsApp
				answ = answ + "\n\n710\n   Contact Instagram:  \n   ..." + dataSialum.ContactInstsgram
				answ = answ + "\n\n711\n   Contact YouTube:  \n   ..." + dataSialum.ContactYoutube

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, answ)
				bot.Send(msg)
			}

			//демонстрация пароля
			if userText == "show pas" || userText == "pas" || userText == "password" {
				answ := "Password: \n"
				for i, e := range dataSialum.AdminPass {
					if i < len(dataSialum.AdminPass)/4 {
						answ = answ + string(e)
					} else {
						answ = answ + "*"
					}
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, answ)
				bot.Send(msg)
			}

			//
			//
			//
			//
			//
			//
			//
			//
			//
			//
			//
			//
			//
			//
			//
			//
			//
			//
			//
			//
			//
			////////////////////   CHANGING SITE CONFIG BY BOT   ////////////////////
			changerDataTelegram("pas=", "Password", &dataSialum.AdminPass, userText, *bot, update)

			changerDataTelegram("000=", "Title Icon", &dataSialum.TitleIcon, userText, *bot, update)
			changerDataTelegram("001=", "Title Head Name", &dataSialum.Title, userText, *bot, update)

			changerDataTelegram("100=", "Head Text 1", &dataSialum.Header1, userText, *bot, update)
			changerDataTelegram("101=", "Head Text 2", &dataSialum.Header2, userText, *bot, update)
			changerDataTelegram("102=", "Head Head Name", &dataSialum.Header3, userText, *bot, update)
			changerDataTelegram("103=", "Head Button Link", &dataSialum.Head_botton, userText, *bot, update)
			changerDataTelegram("104=", "Head Button Photo", &dataSialum.Head_bottonPhoto, userText, *bot, update)

			changerDataTelegram("200=", "Special Offer 1 Photo", &dataSialum.SpecFirstPhoto, userText, *bot, update)
			changerDataTelegram("201=", "Special Offer 1 Name", &dataSialum.SpecFirstHead, userText, *bot, update)
			changerDataTelegram("202=", "Special Offer 1 Text 1", &dataSialum.SpecFirstBody1, userText, *bot, update)
			changerDataTelegram("203=", "Special Offer 1 Text 2", &dataSialum.SpecFirstBody2, userText, *bot, update)
			changerDataTelegram("204=", "Special Offer 1 Text 3", &dataSialum.SpecFirstBody3, userText, *bot, update)
			changerDataTelegram("205=", "Special Offer 1 Link", &dataSialum.SpecFirstLink, userText, *bot, update)
			changerDataTelegram("206=", "Special Offer 2 Photo", &dataSialum.SpecSecPhoto, userText, *bot, update)
			changerDataTelegram("207=", "Special Offer 2 Name", &dataSialum.SpecSecHead, userText, *bot, update)
			changerDataTelegram("208=", "Special Offer 2 Text 1", &dataSialum.SpecSecBody1, userText, *bot, update)
			changerDataTelegram("209=", "Special Offer 2 Text 2", &dataSialum.SpecSecBody2, userText, *bot, update)
			changerDataTelegram("210=", "Special Offer 2 Text 3", &dataSialum.SpecSecBody3, userText, *bot, update)
			changerDataTelegram("211=", "Special Offer 2 Link", &dataSialum.SpecSecLink, userText, *bot, update)
			changerDataTelegram("212=", "Special Offer 3 Photo", &dataSialum.SpecThirdPhoto, userText, *bot, update)
			changerDataTelegram("213=", "Special Offer 3 Name", &dataSialum.SpecThirdHead, userText, *bot, update)
			changerDataTelegram("214=", "Special Offer 3 Text 1", &dataSialum.SpecThirdBody1, userText, *bot, update)
			changerDataTelegram("215=", "Special Offer 3 Text 2", &dataSialum.SpecThirdBody2, userText, *bot, update)
			changerDataTelegram("216=", "Special Offer 3 Text 3", &dataSialum.SpecThirdBody3, userText, *bot, update)
			changerDataTelegram("217=", "Special Offer 3 Link", &dataSialum.SpecThirdLink, userText, *bot, update)

			changerDataTelegram("300=", "Gallery Description", &dataSialum.GallDisc, userText, *bot, update)

			changerDataTelegram("301=", "Gallery Photo image 1", &dataSialum.HeadPhoto1, userText, *bot, update)
			changerDataTelegram("302=", "Gallery Photo link 1", &dataSialum.HomePhotoLink1, userText, *bot, update)
			changerDataTelegram("303=", "Gallery Photo Text Title 1", &dataSialum.HomeTextTitle1, userText, *bot, update)
			changerDataTelegram("304=", "Gallery Photo Text 1", &dataSialum.HomeText1, userText, *bot, update)

			changerDataTelegram("305=", "Gallery Photo image 2", &dataSialum.HeadPhoto2, userText, *bot, update)
			changerDataTelegram("306=", "Gallery Photo link 2", &dataSialum.HomePhotoLink2, userText, *bot, update)
			changerDataTelegram("307=", "Gallery Photo Text Title 2", &dataSialum.HomeTextTitle2, userText, *bot, update)
			changerDataTelegram("308=", "Gallery Photo Text 2", &dataSialum.HomeText2, userText, *bot, update)

			changerDataTelegram("309=", "Gallery Photo image 3", &dataSialum.HeadPhoto3, userText, *bot, update)
			changerDataTelegram("310=", "Gallery Photo link 3", &dataSialum.HomePhotoLink3, userText, *bot, update)
			changerDataTelegram("311=", "Gallery Photo Text Title 3", &dataSialum.HomeTextTitle3, userText, *bot, update)
			changerDataTelegram("312=", "Gallery Photo Text 3", &dataSialum.HomeText3, userText, *bot, update)

			changerDataTelegram("313=", "Gallery Photo image 4", &dataSialum.HeadPhoto4, userText, *bot, update)
			changerDataTelegram("314=", "Gallery Photo link 4", &dataSialum.HomePhotoLink4, userText, *bot, update)
			changerDataTelegram("315=", "Gallery Photo Text Title 4", &dataSialum.HomeTextTitle4, userText, *bot, update)
			changerDataTelegram("316=", "Gallery Photo Text 4", &dataSialum.HomeText4, userText, *bot, update)

			changerDataTelegram("317=", "Gallery Photo image 5", &dataSialum.HeadPhoto5, userText, *bot, update)
			changerDataTelegram("318=", "Gallery Photo link 5", &dataSialum.HomePhotoLink5, userText, *bot, update)
			changerDataTelegram("319=", "Gallery Photo Text Title 5", &dataSialum.HomeTextTitle5, userText, *bot, update)
			changerDataTelegram("320=", "Gallery Photo Text 5", &dataSialum.HomeText5, userText, *bot, update)

			changerDataTelegram("321=", "Gallery Photo image 6", &dataSialum.HeadPhoto6, userText, *bot, update)
			changerDataTelegram("322=", "Gallery Photo link 6", &dataSialum.HomePhotoLink6, userText, *bot, update)
			changerDataTelegram("323=", "Gallery Photo Text Title 6", &dataSialum.HomeTextTitle6, userText, *bot, update)
			changerDataTelegram("324=", "Gallery Photo Text 6", &dataSialum.HomeText6, userText, *bot, update)

			changerDataTelegram("325=", "Gallery Photo image 7", &dataSialum.HeadPhoto7, userText, *bot, update)
			changerDataTelegram("326=", "Gallery Photo link 7", &dataSialum.HomePhotoLink7, userText, *bot, update)
			changerDataTelegram("327=", "Gallery Photo Text Title 7", &dataSialum.HomeTextTitle7, userText, *bot, update)
			changerDataTelegram("328=", "Gallery Photo Text 7", &dataSialum.HomeText7, userText, *bot, update)

			changerDataTelegram("329=", "Gallery Photo image 8", &dataSialum.HeadPhoto8, userText, *bot, update)
			changerDataTelegram("330=", "Gallery Photo link 8", &dataSialum.HomePhotoLink8, userText, *bot, update)
			changerDataTelegram("331=", "Gallery Photo Text Title 8", &dataSialum.HomeTextTitle8, userText, *bot, update)
			changerDataTelegram("332=", "Gallery Photo Text 8", &dataSialum.HomeText8, userText, *bot, update)

			changerDataTelegram("333=", "Gallery Photo image 9", &dataSialum.HeadPhoto9, userText, *bot, update)
			changerDataTelegram("334=", "Gallery Photo link 9", &dataSialum.HomePhotoLink9, userText, *bot, update)
			changerDataTelegram("335=", "Gallery Photo Text Title 9", &dataSialum.HomeTextTitle9, userText, *bot, update)
			changerDataTelegram("336=", "Gallery Photo Text 9", &dataSialum.HomeText9, userText, *bot, update)

			changerDataTelegram("337=", "Gallery Photo image 10", &dataSialum.HeadPhoto10, userText, *bot, update)
			changerDataTelegram("338=", "Gallery Photo link 10", &dataSialum.HomePhotoLink10, userText, *bot, update)
			changerDataTelegram("339=", "Gallery Photo Text Title 10", &dataSialum.HomeTextTitle10, userText, *bot, update)
			changerDataTelegram("340=", "Gallery Photo Text 10", &dataSialum.HomeText10, userText, *bot, update)

			changerDataTelegram("341=", "Gallery Photo image 11", &dataSialum.HeadPhoto11, userText, *bot, update)
			changerDataTelegram("342=", "Gallery Photo link 11", &dataSialum.HomePhotoLink11, userText, *bot, update)
			changerDataTelegram("343=", "Gallery Photo Text Title 11", &dataSialum.HomeTextTitle11, userText, *bot, update)
			changerDataTelegram("344=", "Gallery Photo Text 11", &dataSialum.HomeText11, userText, *bot, update)

			changerDataTelegram("345=", "Gallery Photo image 12", &dataSialum.HeadPhoto12, userText, *bot, update)
			changerDataTelegram("346=", "Gallery Photo link 12", &dataSialum.HomePhotoLink12, userText, *bot, update)
			changerDataTelegram("347=", "Gallery Photo Text Title 12", &dataSialum.HomeTextTitle12, userText, *bot, update)
			changerDataTelegram("348=", "Gallery Photo Text 12", &dataSialum.HomeText12, userText, *bot, update)

			changerDataTelegram("349=", "Gallery Photo image 13", &dataSialum.HeadPhoto13, userText, *bot, update)
			changerDataTelegram("350=", "Gallery Photo link 13", &dataSialum.HomePhotoLink13, userText, *bot, update)
			changerDataTelegram("351=", "Gallery Photo Text Title 13", &dataSialum.HomeTextTitle13, userText, *bot, update)
			changerDataTelegram("352=", "Gallery Photo Text 13", &dataSialum.HomeText13, userText, *bot, update)

			changerDataTelegram("353=", "Gallery Photo image 14", &dataSialum.HeadPhoto14, userText, *bot, update)
			changerDataTelegram("354=", "Gallery Photo link 14", &dataSialum.HomePhotoLink14, userText, *bot, update)
			changerDataTelegram("355=", "Gallery Photo Text Title 14", &dataSialum.HomeTextTitle14, userText, *bot, update)
			changerDataTelegram("356=", "Gallery Photo Text 14", &dataSialum.HomeText14, userText, *bot, update)

			changerDataTelegram("357=", "Gallery Photo image 15", &dataSialum.HeadPhoto15, userText, *bot, update)
			changerDataTelegram("358=", "Gallery Photo link 15", &dataSialum.HomePhotoLink15, userText, *bot, update)
			changerDataTelegram("359=", "Gallery Photo Text Title 15", &dataSialum.HomeTextTitle15, userText, *bot, update)
			changerDataTelegram("360=", "Gallery Photo Text 15", &dataSialum.HomeText15, userText, *bot, update)

			changerDataTelegram("400=", "Company 1", &dataSialum.Company1, userText, *bot, update)
			changerDataTelegram("401=", "Company 2", &dataSialum.Company2, userText, *bot, update)
			changerDataTelegram("402=", "Company 3", &dataSialum.Company3, userText, *bot, update)
			changerDataTelegram("403=", "Company 4", &dataSialum.Company4, userText, *bot, update)
			changerDataTelegram("404=", "Company 5", &dataSialum.Company5, userText, *bot, update)
			changerDataTelegram("405=", "Company 6", &dataSialum.Company6, userText, *bot, update)

			changerDataTelegram("500=", "Choice Head Name", &dataSialum.ChoiseHead, userText, *bot, update)
			changerDataTelegram("501=", "Choice Text 1", &dataSialum.Choise1, userText, *bot, update)
			changerDataTelegram("502=", "Choice Text 2", &dataSialum.Choise2, userText, *bot, update)
			changerDataTelegram("503=", "Choice Text 3", &dataSialum.Choise3, userText, *bot, update)
			changerDataTelegram("600=", "Worker 1 Name", &dataSialum.Worker1, userText, *bot, update)
			changerDataTelegram("601=", "Worker 1 Deal", &dataSialum.WorkerDeal1, userText, *bot, update)
			changerDataTelegram("602=", "Worker 1 Photo", &dataSialum.WorkerPhoto1, userText, *bot, update)
			changerDataTelegram("603=", "Worker 1 Link", &dataSialum.WorkerLink1, userText, *bot, update)
			changerDataTelegram("600=", "Worker 2 Name", &dataSialum.Worker2, userText, *bot, update)
			changerDataTelegram("601=", "Worker 2 Deal", &dataSialum.WorkerDeal2, userText, *bot, update)
			changerDataTelegram("602=", "Worker 2 Photo", &dataSialum.WorkerPhoto2, userText, *bot, update)
			changerDataTelegram("603=", "Worker 2 Link", &dataSialum.WorkerLink2, userText, *bot, update)
			changerDataTelegram("600=", "Worker 3 Name", &dataSialum.Worker3, userText, *bot, update)
			changerDataTelegram("601=", "Worker 3 Deal", &dataSialum.WorkerDeal3, userText, *bot, update)
			changerDataTelegram("602=", "Worker 3 Photo", &dataSialum.WorkerPhoto3, userText, *bot, update)
			changerDataTelegram("603=", "Worker 3 Link", &dataSialum.WorkerLink3, userText, *bot, update)
			changerDataTelegram("600=", "Worker 4 Name", &dataSialum.Worker4, userText, *bot, update)
			changerDataTelegram("601=", "Worker 4 Deal", &dataSialum.WorkerDeal4, userText, *bot, update)
			changerDataTelegram("602=", "Worker 4 Photo", &dataSialum.WorkerPhoto4, userText, *bot, update)
			changerDataTelegram("603=", "Worker 4 Link", &dataSialum.WorkerLink4, userText, *bot, update)
			changerDataTelegram("600=", "Worker 5 Name", &dataSialum.Worker5, userText, *bot, update)
			changerDataTelegram("601=", "Worker 5 Deal", &dataSialum.WorkerDeal5, userText, *bot, update)
			changerDataTelegram("602=", "Worker 5 Photo", &dataSialum.WorkerPhoto5, userText, *bot, update)
			changerDataTelegram("603=", "Worker 5 Link", &dataSialum.WorkerLink5, userText, *bot, update)
			changerDataTelegram("600=", "Worker 6 Name", &dataSialum.Worker6, userText, *bot, update)
			changerDataTelegram("601=", "Worker 6 Deal", &dataSialum.WorkerDeal6, userText, *bot, update)
			changerDataTelegram("602=", "Worker 6 Photo", &dataSialum.WorkerPhoto6, userText, *bot, update)
			changerDataTelegram("603=", "Worker 6 Link", &dataSialum.WorkerLink6, userText, *bot, update)
			changerDataTelegram("600=", "Worker 7 Name", &dataSialum.Worker7, userText, *bot, update)
			changerDataTelegram("601=", "Worker 7 Deal", &dataSialum.WorkerDeal7, userText, *bot, update)
			changerDataTelegram("602=", "Worker 7 Photo", &dataSialum.WorkerPhoto7, userText, *bot, update)
			changerDataTelegram("603=", "Worker 7 Link", &dataSialum.WorkerLink7, userText, *bot, update)
			changerDataTelegram("700=", "Contact Call 1", &dataSialum.ContactPhone1, userText, *bot, update)
			changerDataTelegram("701=", "Contact Call 2", &dataSialum.ContactPhone2, userText, *bot, update)
			changerDataTelegram("702=", "Contact Call 3", &dataSialum.ContactPhone3, userText, *bot, update)
			changerDataTelegram("703=", "Addres Street", &dataSialum.AddresStreet, userText, *bot, update)
			changerDataTelegram("704=", "Adress Country", &dataSialum.AddresCountry, userText, *bot, update)
			changerDataTelegram("705=", "Adress Link", &dataSialum.AddresLink, userText, *bot, update)
			changerDataTelegram("706=", "E-mail 1", &dataSialum.Email1, userText, *bot, update)
			changerDataTelegram("707=", "E-mail 2", &dataSialum.Email2, userText, *bot, update)
			changerDataTelegram("703=", "Contact Telegram", &dataSialum.ContactTelegram, userText, *bot, update)
			changerDataTelegram("704=", "Contact WhatsApp", &dataSialum.ContactWhatsApp, userText, *bot, update)
			changerDataTelegram("705=", "Contact Instagram", &dataSialum.ContactInstsgram, userText, *bot, update)
			changerDataTelegram("706=", "Contact YouTube", &dataSialum.ContactWhatsApp, userText, *bot, update)

		}
	}

}
