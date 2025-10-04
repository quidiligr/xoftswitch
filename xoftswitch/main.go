
package main

import (
	"flag"
	"fmt"
	"path"

	"io"
	"crypto/rand"
	"math/big"

	"log"
	
	"net/url"

	"os"

	"os/exec"

	"os/user"

	"runtime"
	"os/signal"
	"path/filepath"


	"xoftswitch/agi"
	"xoftswitch/pkg/addexts"
	"xoftswitch/pkg/delexts"
	
	

	//"time"
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	"context"
	
	"net/http"
	"html/template"

	"time"
	//"strings"
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"database/sql"
	"encoding/base64"
	"errors"
	"regexp"
	"strconv"
	"strings"

	"net" // standard library net for LookupIP, ParseIP, etc.

    gnet "github.com/shirou/gopsutil/v3/net"
     // alias gopsutil/net as gnet
	"github.com/shirou/gopsutil/v3/process"
	"github.com/shirou/gopsutil/v3/disk"
	
	//"io"
	//"crypto/rand"
	

	_ "github.com/go-sql-driver/mysql"

	"github.com/googolgl/gami"
	"github.com/googolgl/gami/event"
	//"xoftswitch/pjsip" 
	"crypto/tls"
	"net/smtp"
	
)



type Map map[string]interface{}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Apitoken string `json:"apitoken"`
	Apitokencreated string `json:"apitokencreated"`
	
	//Id string `json:"id"`
	Name        string `json:"name"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Phonenumber string `json:"phonenumber"`

	Photo string `json:"photo"`

	Status    int    `json:"status"`
	Sessionid string `json:"sessionid"`
	Roleid    string `json:"roleid"` //admin,employee,customer
	//Extensions []UserExtension `json:"extensions"`
	Missed         int    `json:"missed"`
	Misseduniqueid string `json:"misseduniqueid"`
	Unreadmessages int    `json:"unreadmessages"`
	Homepageurl    string `json:"homepageurl"`
	Ver            string `json:"ver"`
}

// with this struct we added Blockedby properties to orig User
type UserMoreInfo struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	//Id string `json:"id"`
	Name        string `json:"name"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Phonenumber string `json:"phonenumber"`

	Photo string `json:"photo"`

	Status    int    `json:"status"`
	Sessionid string `json:"sessionid"`
	Roleid    string `json:"roleid"` //admin,employee,customer
	//Extensions []UserExtension `json:"extensions"`
	Missed         int    `json:"missed"`
	Misseduniqueid string `json:"misseduniqueid"`
	Unreadmessages int    `json:"unreadmessages"`
	Homepageurl    string `json:"homepageurl"`
	Ver            string `json:"ver"`
	Blockedby      string `json:"blockedby"`
}


type ActivateExtensionRequest struct {
	//id       int

	//password string
	//loggedAt time.Time
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Created  string `json:"created"`

	//Status int `json:"status"`
	//Roleid	string	`json:"roleid"`

}

type Role struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type NameValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}


type UserExtension struct {
	Username  string `json:"username"`
	Extension string `json:"extension"`
	Extenpassword string `json:"extenpassword"`
	
	Groupid   string `json:"groupid"`

	
}

type BlockedUser struct {
	Username        string `json:"username"`
	Blockedusername string `json:"blockedusername"`
	//Password string `json:"password"` //this is dynamic populated from auth_extensions

}

/*
type Extension struct {
    Number      string
    DisplayName string
    Status      string
    Contacts    string
    Source      string
    Secret      string
    DeviceType  string
    Public      string
    CanBlock    string
    CanUnlist   string
    CanAddList  string
    CanDial     string
    Ver         string
}
*/
type Extension struct {
	Number      string `json:"number"`
	Displayname string `json:"displayname"`
	Email string `json:"email"`
	Status      string `json:"status"`
	Contacts    string `json:"contacts"`
	Source      string `json:"source"` //comma separated
	Secret      string `json:"secret"`
	Devicetype  string `json:"devicetype"`
	Public      string `json:"public"`
	Webrtc      string `json:"webrtc"`
	Mediaencryption      string `json:"media_encryption"`
	Directmedia      string `json:"direct_media"`
	
	Canblock    string `json:"canblock"`
	Canunlist   string `json:"canunlist"`
	Canaddlist  string `json:"canaddlist"`
	Candial     string `json:"candial"`
	
	Ver         string `json:"ver"`

	//Index	int `json:"index"`
}
/*
type ExtensionConfig struct {
	Extension      string `json:"number"`
	Displayname string `json:"displayname"`
	Email string `json:"email"`
	
	
	Secret      string `json:"secret"`
	Devicetype  string `json:"devicetype"`
	Public      string `json:"public"`
	Webrtc      string `json:"webrtc"`
	Canblock    string `json:"canblock"`
	Canunlist   string `json:"canunlist"`
	Canaddlist  string `json:"canaddlist"`
	Candial     string `json:"candial"`
	//Ver         string `json:"ver"`

	//Index	int `json:"index"`
}
*/
type MoreExtensionInfo struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Photo    string `json:"photo"`

	Displayname string `json:"displayname"`
	Number      string `json:"number"`
	Status      string `json:"status"`
	Contacts    string `json:"contacts"`
	Source      string `json:"source"` //comma separated
	Secret      string `json:"secret"`
	Devicetype  string `json:"devicetype"`
	Public      string `json:"public"`
	Canblock    string `json:"canblock"`
	Canunlist   string `json:"canunlist"`
	Canaddlist  string `json:"canaddist"`
	Candial     string `json:"candial"`
	
	//Groupid string `json:"groupid"`

	//Missed int `json:"missed"`
	//Misseduniqueid string `json:"misseduniqueid"`
	Ver string `json:"ver"`

	//Index	int `json:"index"`
}

type MyContacts struct {
	Username string `json:"username"`

	Number      string `json:"number"`
	Displayname string `json:"displayname"`
}

//--- WS SERVER

type WSClient struct {
	conn         *websocket.Conn
	device_token string
	//exten string
	//username string

	//device_type string

}
type WSRegister struct {
	//conn *websocket.Conn
	Exten             string `json:"exten"`
	Username          string `json:"username"`
	Name              string `json:"name"`
	Photo             string `json:"photo"`
	Device_Token      string `json:"device_token"`
	Device_Type       string `json:"device_type"`
	Remote_Addr       string `json:"remote_addr"`
	Contact_Name      string `json:"contact_name"`
	Contact_Host      string `json:"contact_host"`
	Contact_Transport string `json:"contact_transport"`

	Register_Status int    `json:"register_status"`
	Awake_Status    string `json:"awake_status"`
	Date            string `json:"date"`
	Ver             string `json:"ver"`

	//Source string `json:"source"`

}

type WSIncomingCall struct {
	//conn *websocket.Conn
	Id   string `json:"id"`
	Uuid string `json:"uuid"`

	//Exten string `json:"exten"`
	//Username string `json:"username"`
	Device_Token   string `json:"device_token"`
	Date           string `json:"date"`
	CKAnswerState  int    `json:"ck_answer_state"`
	CKHangupState  int    `json:"ck_hangup_state"`
	SIPHangupState int    `json:"sip_hangup_state"`

	//FromExten string `json:"from_exten"`
	//Device_Type string `json:"device_type"`
	//Remote_Addr string `json:"remote_addr"`
	//Register_Status int `json:"register_status"`
	//Awake_Status string `json:"awake_status"`
	//Date string `json:"date"`
}

type WSContact struct {
	contact *WSRegister
	client  *WSClient
}

type IP struct {
	Query string
}

/*
*************************** 511. row ***************************

	calldate: 2023-06-02 15:34:43
		clid: "400 App" <400>
		src: 400
		dst: 405
	dcontext: ext-local
	channel: PJSIP/400-00000379

dstchannel: PJSIP/405-0000037a

	lastapp: Dial
	lastdata: PJSIP/405/sip:405@71.222.1.109:5060,,HhTtrb(func-apply-sipheaders^s^1)
	duration: 11
	billsec: 6

disposition: ANSWERED

	amaflags: 3

accountcode:

	uniqueid: 1685720083.945
	userfield:
		did:

recordingfile:

	cnum: 400
	cnam: 400 App

outbound_cnum:
outbound_cnam:

	dst_cnam:
	linkedid: 1685720083.945

peeraccount:

	sequence: 1054
*/
type CDR struct {
	Calldate    string `json:"calldate"`    // 2023-06-02 15:34:43
	Clid        string `json:"clid"`        //"400 App" <400>
	Src         string `json:"src"`         //400
	Dst         string `json:"dst"`         //405
	Duration    int    `json:"duration"`    //11
	Billsec     int    `json:"billsec"`     //6
	Uniqueid    string `json:"uniqueid"`    //1685720083.945
	Disposition string `json:"disposition"` //ANSWERED
	Srcusername string `json:"scrusername"` //ROM custom field
	Dstusername string `json:"dstusername"` //ROM custom field
	Srcname     string `json:"srcname"`     //ROM custom field
	Dstname     string `json:"dstname"`     //ROM custom field
	Srccontact  string `json:"srccontac"`   //ROM custom field
	Dstcontact  string `json:"dstcontact"`  //ROM custom field
	Sequence    string `json:"sequence"`    //ROM custom field

}

type MSG struct {
	Id       string `json:"id"`
	Inboxid  string `json:"inboxid"`
	Created  string `json:"created"`
	Username string `json:"username"`
	Message  string `json:"message"`
	From     string `json:"from"`
	To       string `json:"to"`
	Roomid   string `json:"roomid"`
}

type INBOX struct {
	Id       string `json:"id"`
	Created  string `json:"created"`
	Username string `json:"username"`
	Message  string `json:"message"`
	From     string `json:"from"`
	To       string `json:"to"`
	Fromname string `json:"fromname"`
	Toname   string `json:"toname"`
	Roomid   string `json:"roomid"`
}

type AIATOKEN struct {
	Token     string `json:"token"`
	Created   string `json:"created"`
	Username  string `json:"username"`
	Ipaddress string `json:"ipaddress"`
}

type Config struct {
	AdminUsername            string `json:"admin_username"`
    AdminPassword            string `json:"admin_password"`
    AmiPassword               string `json:"ami_password"`
    AmiUsername               string `json:"ami_username"`
	APIKey                    string `json:"apikey"`
	AsteriskConfPath          string `json:"asterisk_conf_path"`
    AutoJoin                  string `json:"auto_join"`
    CertDir                   string `json:"cert_dir"`
    HomepageURL               string `json:"homepageurl"`
    LogoFilename              string `json:"logo_filename"`
    MaxExtensionCountPerUser  string `json:"max_extension_count_per_user"`
    Name                      string `json:"name"`
    OkBSSID                   string `json:"ok_bssid"`
    OkLogoFilename            string `json:"ok_logo_filename"`
    OkSSID                    string `json:"ok_ssid"`
    PublicHostname            string `json:"public_hostname"`
    PublicIP                  string `json:"public_ip"`
    WifiBSSID                 string `json:"wifi_bssid"`
    WifiSSID                  string `json:"wifi_ssid"`
	SMTPHost     string `json:"smtp_host"`
	SMTPPort     string `json:"smtp_port"` // "465" or "587"
	SMTPUsername string `json:"smtp_username"`
	SMTPPassword string `json:"smtp_password"`
	AdminEmail   string `json:"admin_email"` // recipient
}

var configPath = "/etc/xoftswitch/config.json"

var api_xoftswitch =  "https://api.xoftphone.com/api/xoftswitch"
var api_server_ip = ""

var apikey = ""
var kiv_iv = ""
var kiv_key = ""


var clients = []WSClient{}
var registered = []WSRegister{}
var incomingcalls = []WSIncomingCall{}
var userextensions = []UserExtension{}
var users = []User{}
var joinrequests = []JoinRequest{}
var activateextensionrequests = []ActivateExtensionRequest{}

var roles = make(map[string]string)
var extension_range = make(map[string]string)
var extensions = []Extension{} //make(Map)
var extensions_auth = make(map[string]string)
var endpoints = make(Map) //make(map[string]string)
var contactstatusmaptable = make(Map)
var config = make(map[string]string)
var ping_enabled = false
var hostname = ``        //"fpbx.sharecle.com"
var endpoint = make(Map) //make(map[string]string) //Map
var glo_ami *gami.AMIClient

// This should be in an env file in production
//const cryptosecret string = "abc&1*~#^2^#s0^=)^^7%b34"
//var cryptobytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
//var cryptosecret = ""

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

/*
"user" : {
		"username" : "unknown",
		"name" : "Dr Ziegler",
		"number" : "205"
}
*/

var colorReset = "\033[0m"

var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorYellow = "\033[33m"
var colorBlue = "\033[34m"
var colorPurple = "\033[35m"
var colorCyan = "\033[36m"
var colorWhite = "\033[37m"

type LogLevel int

const (
	LOG_DEBUG LogLevel = iota
	LOG_INFO
	LOG_WARN
	LOG_ERROR
	LOG_NONE
)

var levelNames = map[LogLevel]string{
	LOG_DEBUG: "DEBUG",
	LOG_INFO:  "INFO",
	LOG_WARN:  "WARN",
	LOG_ERROR: "ERROR",
}

var currentLogLevel = LOG_NONE // default to no logs


var db *sql.DB

var isReload bool

//var tmpl = template.Must(template.ParseGlob("templates/*.html"))
var templates = map[string]*template.Template{}

/*
go run main.go
# no output (LOG_NONE)

go run main.go --log=debug
# shows all messages (debug and above)

go run main.go --log=warn
# only shows warn and error
*/

var Version = "2025.10.1"
var lastNetIn uint64
var lastNetOut uint64
var lastNetTime time.Time

func main() {

	//done_agi := make(chan bool)
	//done_ws :=  make(chan bool)
	//readFileJSONToken()
	var err error
	/*
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/xoftswitchdb")

	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	// ✅ Close the DB when main() returns (usually when the program ends)
    defer db.Close()
	*/
	
	logLevelArg := flag.String("log", "", "log level: debug, info, warn, error")

	// Parse the flags
	flag.Parse()


	// Normalize to lowercase
	switch strings.ToLower(*logLevelArg) {
	case "debug":
		currentLogLevel = LOG_DEBUG
	case "info":
		currentLogLevel = LOG_INFO
	case "warn":
		currentLogLevel = LOG_WARN
	case "error":
		currentLogLevel = LOG_ERROR
	default:
		currentLogLevel = LOG_NONE // default
	}

	var workingDir string

	switch runtime.GOOS {
		case "linux":
			workingDir = "/etc/xoftswitch"
		//case "windows":
		//	workingDir = "C:\\xoftswitch"
		default:
			workingDir = "./"
	}

	err = os.Chdir(workingDir)
	if err != nil {
		fmt.Printf("Failed to change working directory to %s: %v\n", workingDir, err)
		os.Exit(1)
	}

	fmt.Printf("474: Working directory changed to: %s\n", workingDir)

	if err := initDB(); err != nil {
        log.Fatalf("❌ Database error: %v", err)
    }
	
	defer db.Close() // only when the program is about to exit

	
	initTables()

	loadTemplates() // Load all templates at startup
	
	setAIPServerIPv4()

	newpath := filepath.Join(".", "public")
	err = os.MkdirAll(newpath, os.ModePerm)

	newpath = filepath.Join(".", "static")
	err = os.MkdirAll(newpath, os.ModePerm)

	newpath = filepath.Join(".", "temp-images")
	err = os.MkdirAll(newpath, os.ModePerm)

	readFileConfig()
	fmt.Println(`643: config=`, config)
	//register server if needed

	registerXoServer()
	if len(apikey) == 0 || len(kiv_key) == 0 || len(kiv_iv) == 0 {
		panic(errors.New(`Failed to register server`))
	}

	
	readFileRoles()

	readFileExtensionRange()

	dbLoadJoinRequests()

	hostname = config[`public_hostname`]
	fmt.Println(`286: *********** hostname=`, hostname)
	
	
	reloadXoftSwitch()

	done := make(chan bool)

	
	
	go startWSServerTLS()

	
	
	go startAMI()

	/*uncomment to enable console*/
	//go startConsole(done, ws_server)

	<-done

}
/*
func initDB() error {
    var err error

    db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/xoftswitchdb")
    if err != nil {
        return fmt.Errorf("failed to open DB: %w", err)
    }

    if err = db.Ping(); err != nil {
        return fmt.Errorf("failed to ping DB: %w", err)
    }

    fmt.Println("✅ Database connection established")
    return nil
}
*/
func initDB() error {
    var err error

    // Step 1: Connect without selecting a specific DB
    tempDB, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/")
    if err != nil {
        return fmt.Errorf("failed to open temp DB connection: %w", err)
    }
    defer tempDB.Close()

    // Step 2: Create database if it doesn't exist
    _, err = tempDB.Exec("CREATE DATABASE IF NOT EXISTS xoftswitchdb")
    if err != nil {
        return fmt.Errorf("failed to create database: %w", err)
    }

    // Step 3: Now connect to the newly created DB
    db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/xoftswitchdb")
    if err != nil {
        return fmt.Errorf("failed to open xoftswitchdb: %w", err)
    }

    // Step 4: Verify the connection
    if err = db.Ping(); err != nil {
        return fmt.Errorf("failed to ping DB: %w", err)
    }

    fmt.Println("✅ Database connection established")
    return nil
}


func initTables(){
	var qry = "CREATE TABLE IF NOT EXISTS inbox(id varchar(75) UNIQUE, created datetime,username varchar(75),`from` varchar(75),`to` varchar(75),fromname varchar(75), toname varchar(75), roomid varchar(30),messageid varchar(48), message varchar(512))"

	_, err := db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`Error creating inbox table!`)
		panic(err.Error())
	}
	//try to connect
	fmt.Println(`inbox table created OK!!!`)

	//create inbox_messages if needed
	qry = "CREATE TABLE IF NOT EXISTS messages(id varchar(48) NOT NULL,inboxid varchar(75) NOT NULL, created datetime,username varchar(75),`from` varchar(75),`to` varchar(75),roomid varchar(30),message varchar(512))"
	_, err = db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`Error creating table inbox_messages!`)
		panic(err.Error())
	}

	fmt.Println(`inbox_messages table created OK!!!`)

	qry = `CREATE TABLE IF NOT EXISTS extensions(
			number varchar(25) NOT NULL PRIMARY KEY,
			displayname varchar(50) DEFAULT '',
			email varchar(50) DEFAULT '',
			status varchar(25) DEFAULT '',
			contacts varchar(100) DEFAULT '',
			source varchar(10) DEFAULT '',
			
			secret varchar(50) DEFAULT '', 
			direct_media varchar(3) DEFAULT 'yes',
			media_encryption varchar(3) DEFAULT 'no', 
			webrtc varchar(3) DEFAULT '', 
			devicetype varchar(10) DEFAULT '', 
			public varchar(3) DEFAULT 'no', 
			
			canblock varchar(3) DEFAULT 'yes', 
			canunlist varchar(3) DEFAULT 'yes', 
			canaddlist varchar(3) DEFAULT 'yes', 
			candial varchar(3) DEFAULT 'yes', 
			ver varchar(25) DEFAULT '')`
	_, err = db.Query(qry)
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`Error creating extensions table!`)
		panic(err.Error())
	}
	fmt.Println(`extensions table created OK!!!`)
/*
	qry = `CREATE TABLE IF NOT EXISTS extensions_config(
		extension varchar(25) NOT NULL PRIMARY KEY,
		displayname varchar(50) DEFAULT '',
		email varchar(50) DEFAULT '',
		secret varchar(50) DEFAULT '', 
		devicetype varchar(10) DEFAULT '', 
		public varchar(3) DEFAULT 'no', 
		webrtc varchar(3) DEFAULT 'yes', 
		canblock varchar(3) DEFAULT 'yes', 
		canunlist varchar(3) DEFAULT 'yes', 
		canaddlist varchar(3) DEFAULT 'yes', 
		candial varchar(3) DEFAULT 'yes')`
_, err = db.Query(qry)
if err != nil {
	//panic(err.Error()) // proper error handling instead of panic in your app
	fmt.Println(`Error creating extensions_config table!`)
	panic(err.Error())
}
fmt.Println(`extensions table created OK!!!`)
*/


	//maps users added contacts to extensions
	qry = `CREATE TABLE IF NOT EXISTS mycontacts(
			username varchar(75) NOT NULL,
			number varchar(25) NOT NULL,
			displayname varchar(75) DEFAULT '',
			PRIMARY KEY (username, number)
			)`
	_, err = db.Query(qry)
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`Error creating mycontacts table!`)
		panic(err.Error())
	}
	fmt.Println(`mycontacts table created OK!!!`)

	qry = `CREATE TABLE IF NOT EXISTS extensiongroup(
			number varchar(25),
			groupid varchar(25),
			primary key (number, groupid) )`
	_, err = db.Query(qry)
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`Error creating extensiongroup table!`)
		panic(err.Error())
	}
	fmt.Println(`extensions table created OK!!!`)

	//contact_transport,register_status,awake_status,date
	qry = `CREATE TABLE IF NOT EXISTS registered(
			exten varchar(25),
			username varchar(75),
			name varchar(75),
			photo varchar(100),
			
			device_token varchar(128),
			device_type varchar(10),
			remote_addr varchar(24),
			contact_name varchar(24),
			contact_host varchar(32),
			contact_transport varchar(8),
			register_status int,
			awake_status varchar(2), 
			date varchar(25) DEFAULT '',
			ver varchar(25) DEFAULT '',
			
			primary key (exten, device_token)
			)`
	_, err = db.Query(qry)
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`Error creating registered table!`)
		panic(err.Error())
	}
	fmt.Println(`registered table created OK!!!`)

	qry = `CREATE TABLE IF NOT EXISTS users(
			username varchar(75) NOT NULL PRIMARY KEY,
			email varchar(75)  DEFAULT '',
			name varchar(50) DEFAULT '', 
			firstname varchar(25) DEFAULT '', 
			lastname varchar(25) DEFAULT '', 
			phonenumber varchar(25) DEFAULT '', 
			photo varchar(100) DEFAULT '', 
			roleid varchar(25) DEFAULT '', 
			status INT DEFAULT 0, 
			sessionid varchar(100)  DEFAULT '',
			missed INT DEFAULT 0, 
			misseduniqueid varchar(25) DEFAULT '',
			unreadmessages INT DEFAULT 0, 
			homepageurl varchar(200) DEFAULT '', 
			ver varchar(25) DEFAULT '' )`

	_, err = db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`Error creating users table!`)
		panic(err.Error())
	}

	fmt.Println(`users table created OK!!!`)

	qry = "CREATE TABLE IF NOT EXISTS userextensions(username varchar(75) NOT NULL,extension varchar(50)  NOT NULL, groupid varchar(256) NOT NULL DEFAULT '')"

	_, err = db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`Error creating extensions table!`)
		panic(err.Error())
	}
	//try to connect
	fmt.Println(`extensions table created OK!!!`)

	qry = "CREATE TABLE IF NOT EXISTS blockedusers(blockedby varchar(75) NOT NULL,blockeduser varchar(75)  NOT NULL)"

	_, err = db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`Error creating extensions table!`)
		panic(err.Error())
	}
	//try to connect
	fmt.Println(`extensions table created OK!!!`)

	/*
				Created string `json:"created"`
		Username string `json:"username"`
		Email string `json:"email"`
		Name string `json:"name"`
	*/
	//190530b3-a446-40f6-9fee-988f30b0095b
	qry = "CREATE TABLE IF NOT EXISTS joinrequests(id varchar(48) PRIMARY KEY, email varchar(50),username varchar(50) NOT NULL DEFAULT '',hostname varchar(50) NOT NULL DEFAULT '',name varchar(50) NOT NULL DEFAULT '',status INT DEFAULT 0,auto_join BOOLEAN DEFAULT FALSE,roleid VARCHAR(30) DEFAULT 'customer', created varchar(30) NOT NULL DEFAULT '')"

	_, err = db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`Error creating joinrequests table!`)
		panic(err.Error())
	}
	//try to connect
	fmt.Println(`joinrequests table created OK!!!`)

	qry = `CREATE TABLE IF NOT EXISTS aiatokens(
			token varchar(128)  NOT NULL PRIMARY KEY,
			created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			username varchar(75) NOT NULL,
			ipaddress varchar(16) DEFAULT '')
			`

	_, err = db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`Error creating aiatokens table!`)
		panic(err.Error())
	}

	fmt.Println(`aiatokens table created OK!!!`)

}

func reloadXoftSwitch(){
	fmt.Println(`1006: reloadXoftSwitch()!`)
	/*
	if err := syncAORConfig(); err != nil {
        fmt.Println("Failed to sync AOR:", err)
    }
    if err := syncAuthConfig(); err != nil {
        fmt.Println("Failed to sync Auth:", err)
    }
    if err := syncEndpointConfig(); err != nil {
        fmt.Println("Failed to sync Endpoint:", err)
    }
		

    // Set ownership to asterisk:asterisk
    chownFiles := []string{
        "/etc/asterisk/pjsip.aor_custom.conf",
        "/etc/asterisk/pjsip.auth_custom.conf",
        "/etc/asterisk/pjsip.endpoint_custom.conf",
    }
    for _, file := range chownFiles {
        if err := chownToAsterisk(file); err != nil {
            fmt.Println("Failed to set ownership for", file, ":", err)
        }
    }
		*/
	loadExtensionsAuth()
	loadCustomExtensionsAuth()
	loadEndpoints()
	loadCustomEndpoints()

	//readFileRegistered()
	dbLoadRegistered()
	//readFileUserExtensions()
	dbLoadUserExtensions()
	//readFileUsers()
	dbLoadUsers()
}

/*
func Log(level LogLevel, v ...interface{}) {
	if level >= currentLogLevel {
		fmt.Println(v...)
	}
}
*/
func Log(level LogLevel, v ...interface{}) {
	if level >= currentLogLevel {
		prefix := levelNames[level]
		fmt.Println(append([]interface{}{prefix + ":"}, v...)...)
	}
}

func setAIPServerIPv4(){
/*
	ips, _ := net.LookupIP("api.xoftphone.com")
    for _, ip := range ips {
        if ipv4 := ip.To4(); ipv4 != nil {
            fmt.Println("IPv4: ", ipv4)
			api_server_ip = string(ipv4)
        }
    }
	*/

	//api_server_ip = "208.76.253.182"
	
	fqdn := "api.xoftswitch.com" // Replace with your FQDN

	ips, err := net.LookupIP(fqdn)
	if err != nil {
		fmt.Printf("Failed to resolve %s: %v\n", fqdn, err)
		return
	}

	/*
	ips, _ := net.LookupIP(fqdn)
	for _, ip := range ips {
		if ip.To4() != nil {
			fmt.Println(ip.String())
			break
		}
	}
	*/
	fmt.Printf("IP addresses for %s:\n", fqdn)
	for _, ip := range ips {
		if ip.To4() != nil {
			api_server_ip = ip.String() //"208.76.253.182"
			fmt.Println("api_server_ip IPv4:", api_server_ip)
			
		} else {
			fmt.Println("IPv6:", ip)
		}
	}

	
}
// --> Admin only
func createUserExtension(username, extension string) *UserExtension {
	//readFileUserExtensions()
	if len(username) == 0 || len(extension) == 0 {
		fmt.Println(`729: ERROR createUserExtension. Invalid username or extension.`)
		return nil
	}
	ret := findOneUserExtensions(username, extension)
	if ret == nil {

		//writeFileUserExtensions()

		new_ue := UserExtension{Username: username, Extension: extension}
		userextensions = append(userextensions, new_ue)
		err := dbUpsertUserExtension(new_ue)
		if err != nil {
			fmt.Println(string(colorRed), err.Error())
			fmt.Println(string(colorReset))
		}

		ret = &userextensions[len(userextensions)-1]
		//readFileUserExtensions()
	}

	return ret

}

func removeUserExtensionIngroup_noneed(username, extension, groupid string) *UserExtension {
	//readFileUserExtensions()
	fmt.Println(`897: START replaceUserExtensionNotIngroup`)
	if len(username) == 0 || len(extension) == 0 {
		fmt.Println(`729: ERROR replaceUserExtensionNotIngroup. Invalid username or extension.`)
		return nil
	}

	removeAllUserExtensionsNotInGroup(extension, groupid)

	ret := findOneUserExtensions(username, extension)
	if ret == nil {

		//writeFileUserExtensions()

		new_ue := UserExtension{Username: username, Extension: extension, Groupid: groupid}
		userextensions = append(userextensions, new_ue)
		err := dbUpsertUserExtension(new_ue)
		if err != nil {
			fmt.Println(string(colorRed), err.Error())
			fmt.Println(string(colorReset))
		}

		ret = &userextensions[len(userextensions)-1]
		//readFileUserExtensions()
	}

	return ret

}

func replaceUserExtensionNotIngroup(username, extension, groupid string) *UserExtension {
	//readFileUserExtensions()
	fmt.Println(`897: START replaceUserExtensionNotIngroup`)
	if len(username) == 0 || len(extension) == 0 {
		fmt.Println(`729: ERROR replaceUserExtensionNotIngroup. Invalid username or extension.`)
		return nil
	}

	removeAllUserExtensionsNotInGroup(extension, groupid)

	ret := findOneUserExtensions(username, extension)
	if ret == nil {

		//writeFileUserExtensions()

		new_ue := UserExtension{Username: username, Extension: extension, Groupid: groupid}
		userextensions = append(userextensions, new_ue)
		err := dbUpsertUserExtension(new_ue)
		if err != nil {
			fmt.Println(string(colorRed), err.Error())
			fmt.Println(string(colorReset))
		}

		ret = &userextensions[len(userextensions)-1]
		//readFileUserExtensions()
	}

	return ret

}

func deleteUserExtension(username, exten string) error {
	fmt.Println(`3500: deleteUserExtension(`, username, `,`, exten, `)`)
	//is_save_file := false
	remove_index := -1
	fmt.Println(`883: Remove user extension start...`)
	if len(userextensions) == 0 {
		fmt.Println(`userextensions is empty. Exit...`)

		return nil
	}

	remove_index, ue, err := findUserExtensions(username, exten)
	if err != nil {
		return err
	}

	if remove_index >= 0 && ue != nil {
		userextensions = append(userextensions[:remove_index], userextensions[remove_index+1:]...)
		//new_ue := UserExtension{Username: username, Extension: extension}
		//userextensions = append(userextensions, new_ue)
		err := dbDeleteUserExtension(username, exten)
		if err != nil {
			fmt.Println(string(colorRed), err.Error())
			fmt.Println(string(colorReset))
		}

		//is_save_file = true

	}

	/*
		if is_save_file{
			writeFileUserExtensions()
		}
	*/
	fmt.Println(`Remove user extension done!`)
	return nil

}

func deleteUserExtensions(username string) {
	fmt.Println(`3500: deleteUserExtensions(username=`, username)
	//is_save_file := false
	remove_index := -1
	fmt.Println(`923: Remove user extension start...`)
	if len(userextensions) == 0 {
		return
	}

	for {

		if len(userextensions) == 0 {
			break
		}

		for index, x := range userextensions {
			if x.Username == username {
				remove_index = index
				//is_save_file = true
				break

			}
		}

		if remove_index >= 0 {
			if len(userextensions) == 1 {
				userextensions = userextensions[:0]

			} else {

				userextensions = append(userextensions[:remove_index], userextensions[remove_index+1:]...)

			}

			remove_index = -1

		} else {
			//done here
			break

		}

	}

	err := dbDeleteAllUserExtensionsByUsername(username)
	if err != nil {
		fmt.Println(string(colorRed), err.Error())
		fmt.Println(string(colorReset))
	}
	/*
		if is_save_file{
			writeFileUserExtensions()
		}
	*/
	fmt.Println(`Remove user extension done!`)

}

func deleteUserExtensionsByExten(exten string) {
	//fmt.Println(`3500: deleteUserExtensions(username=`,username)
	//is_save_file := false
	remove_index := -1
	fmt.Println(`Remove extension start...`)
	/*
		if len(userextensions) == 0{
			return
		}
	*/

	for {

		if len(userextensions) == 0 {
			break
		}

		for index, x := range userextensions {
			if x.Extension == exten {
				remove_index = index
				//is_save_file = true
				break

			}
		}

		if remove_index >= 0 {
			if len(userextensions) == 1 {
				userextensions = userextensions[:0]
			} else {

				userextensions = append(userextensions[:remove_index], userextensions[remove_index+1:]...)

			}

			remove_index = -1

		} else {
			//done here
			break

		}

	}

	err := dbDeleteAllUserExtensionsByExtension(exten)
	if err != nil {
		fmt.Println(string(colorRed), err.Error())
		fmt.Println(string(colorReset))
	}
	/*
		if is_save_file{
			writeFileUserExtensions()
		}
	*/

	fmt.Println(`Remove user extension done!`)

}

func deleteExtension(exten string) {
	//fmt.Println(`3500: deleteUserExtensions(username=`,username)
	is_save_file := false
	remove_index := -1
	fmt.Println(`Remove extension start...`)
	/*
		if len(userextensions) == 0{
			return
		}
	*/

	for {

		if len(extensions) == 0 {
			break
		}

		for index, x := range extensions {
			if x.Number == exten {
				remove_index = index
				is_save_file = true
				break

			}
		}

		if remove_index >= 0 {
			if len(extensions) == 1 {
				extensions = extensions[:0]
			} else {

				extensions = append(extensions[:remove_index], extensions[remove_index+1:]...)

			}

			dbDeleteExtension(exten)
			remove_index = -1

		} else {
			//done here
			break

		}

	}
	if is_save_file {
		//writeFileExtensions()
	}
	fmt.Println(`Remove user extension done!`)

}

func deleteRegistered(username, devicetoken string) {
	//fmt.Println(`3500: deleteRegistered(username=`,username)
	is_save_file := false
	remove_index := -1
	fmt.Println(`Remove registered start...`)
	/*
		if len(userextensions) == 0{
			return
		}
	*/

	for {

		if len(registered) == 0 {
			break
		}

		for index, x := range registered {
			if x.Username == username && x.Device_Token == devicetoken {
				remove_index = index
				is_save_file = true
				break

			}
		}

		if remove_index >= 0 {
			if len(extensions) == 1 {
				registered = registered[:0]
			} else {

				registered = append(registered[:remove_index], registered[remove_index+1:]...)

			}

			remove_index = -1

		} else {
			//done here
			break

		}

	}
	if is_save_file {
		writeFileRegistered()
		dbSaveRegistered()
	}
	fmt.Println(`Remove user extension done!`)

}
func deleteRegisteredInvalidDeviceToken() {
	//fmt.Println(`3500: deleteRegistered(username=`,username)
	is_save_file := false
	remove_index := -1
	fmt.Println(`Remove registered start...`)
	/*
		if len(userextensions) == 0{
			return
		}
	*/

	for {

		if len(registered) == 0 {
			break
		}

		for index, x := range registered {
			if x.Device_Token == "" || len(x.Device_Token) < 32 {
				remove_index = index
				is_save_file = true
				break

			}
		}

		if remove_index >= 0 {
			is_save_file = true
			if len(extensions) == 1 {
				registered = registered[:0]
			} else {

				registered = append(registered[:remove_index], registered[remove_index+1:]...)

			}

			remove_index = -1

		} else {
			//done here
			break

		}

	}
	if is_save_file {
		writeFileRegistered()
		dbSaveRegistered()
	}
	fmt.Println(`deleteRegisteredInvalidDeviceToken() !`)

}

func deleteAuthExtension(exten string) {
	//fmt.Println(`3500: deleteUserExtensions(username=`,username)
	is_save_file := false
	//remove_index := -1
	fmt.Println(`Remove extension start...`)
	/*
		if len(userextensions) == 0{
			return
		}
	*/

	//for{

	/*
		if len(extensions_auth) == 0 {
			break
		}

		for index,x := range extensions_auth{
			if x == exten {
				remove_index = index
				is_save_file = true
				break

			}
		}
	*/

	_, ok_ea := extensions_auth[exten]
	if ok_ea {

		delete(extensions_auth, exten)
		is_save_file = true

	}

	//}
	if is_save_file {
		writeFileExtensionsAuth()
	}
	fmt.Println(`Remove user auth extension done!`)

}

func deleteAllXoExtensions() {
	//fmt.Println(`3500: deleteUserExtensions(username=`,username)
	is_save_file := false
	remove_index := -1
	fmt.Println(`1273: Remove user extension start...`)
	if len(extensions) == 0 {
		return
	}

	for {

		if len(extensions) == 0 {
			break
		}

		for index, x := range extensions {
			//if x.Source == `xo` {
			remove_index = index
			is_save_file = true
			deleteUserExtensionsByExten(x.Number)
			dbDeleteExtension(x.Number)
			break

			//}
		}

		if remove_index >= 0 {
			if len(extensions) == 1 {
				extensions = extensions[:0]
			} else {

				extensions = append(extensions[:remove_index], extensions[remove_index+1:]...)

			}

			remove_index = -1

		} else {
			//done here
			break

		}

	}
	if is_save_file {
		//writeFileExtensions()
	}
	fmt.Println(`Remove user extension done!`)

}

func deleteUser(username string) error {
	//is_save_file := false
	wsnotified := 0
	//notify device ws
	ue, err := findAllUserExtensions(username)
	if err != nil {
		return err
	}
	//if len(ue) > 0{

	for _, e := range ue {
		r := findClientsByExten(e.Extension)
		//if len(r) > 0{
		for _, r_item := range r {
			if r_item.Username == username {
				for _, c := range clients {
					if c.device_token == r_item.Device_Token {
						err := c.conn.WriteMessage(1, []byte(fmt.Sprintf(`{"type":"DISJOINED","hostname":"%s"}`, hostname)))
						if err == nil {
							wsnotified++
							//fmt.Println(`deleteUser err=`,err.Error())
						}

					}

				}
			}
		}
		//}
	}

	//}
	if wsnotified == 0 {
		//try push

		httpPushNotifyData(username, map[string]string{`type`: `DISJOINED`})
	}

	deleteUserExtensions(username)

	if len(users) == 0 {
		return nil
	}

	i, user := findUserByUsername(username)
	if user != nil {
		users = append(users[:i], users[i+1:]...)
		dbDeleteUser(username)
		//is_save_file = true
	}
	/*
		if is_save_file{
			writeFileUsers()
		}
	*/

	return nil

}

func emailInputValidation(email string) bool {
	if len(email) == 0 {
		return false
	}

	/*if len(email) < 5{
		fmt.Println(`603: email address does not pass validation!`)
		return false
	}*/

	email_split := strings.Split(email, `@`)
	if len(email_split) != 2 {
		fmt.Println(`592.1: email address does not pass validation!`)
		return false
	}

	if len(email_split[0]) < 1 || len(email_split[1]) < 2 {
		fmt.Println(`592.2: email address does not pass validation!`)
		return false
	}

	return true

}

func usernameInputValidation(username string) bool {
	if len(username) < 3 {
		fmt.Println(`610: username does not pass validation!`)
		return false
	}

	return true

}

func extensionInputValidation(extension string) bool {
	exten, err := strconv.Atoi(extension)
	//if len(extension) == 0 || len(extension) > 9{
	if err != nil || exten < 0 || exten > 999999999 {
		fmt.Println(`610: extension does not pass validation!`)
		return false
	}

	return true

}

//--> Admin only
/*
func createUser_dep(username,name, email,roleId string) (int, *User){

	//readFileUsers()

	i,user := findUserByUsername(username)
	if user != nil{
		return i,user
	}


	ver := time.Now().Format(time.RFC3339)
	new_user := User{Username: username, Email:email, Name: name,Roleid: roleId, Status: 1, Sessionid: ``,Missed: 0,Misseduniqueid: ``, Unreadmessages: 0, Ver: ver}
	users = append(users, new_user)
	dbUpsertUser(new_user)
	//writeFileUsers()

	i,user = findUserByUsername(username)



	return i,user
}
*/

func createUser(firstname, lastname, name, username, email, roleId, photo, phonenumber, homepageurl string) (int, *User) {

	fmt.Println(`7934: START createUser(`,
		`firstname=`, firstname,
		`lastname=`, lastname,
		`name=`, name,
		`username=`, username,
		`email=`, email,
		`roleId=`, roleId,
		`photo=`, photo,
		`phonenumber=`, phonenumber,
		`homepageurl=`, homepageurl)

	//readFileUsers()
	func_name := `createUser`
	fmt.Println(func_name, "(", username, name, email, roleId, photo, ")")

	i, user := findUserByUsername(username)
	if user != nil {
		return i, user
	}

	/*
		i, user = findUserByEmail(email)
		if user != nil{
			return i,user
		}
	*/
	ver := time.Now().Format(time.RFC3339)
	new_user := User{Username: username, Email: email, Name: name, Firstname: firstname, Lastname: lastname, Phonenumber: phonenumber, Photo: photo, Roleid: roleId, Status: 1, Sessionid: ``, Missed: 0, Misseduniqueid: ``, Unreadmessages: 0, Homepageurl: homepageurl, Ver: ver}
	users = append(users, new_user)
	dbUpsertUser(new_user, func_name)
	//writeFileUsers()

	i, user = findUserByUsername(username)

	return i, user
}

func removeUser(username string) error {

	fmt.Println(`1507: START removeUser(`,
		`username=`, username)

	//readFileUsers()
	func_name := `removeUser`
	//fmt.Println(func_name,"(",username,name,email,roleId,photo,")")
	/*
		_,user := findUserByUsername(username)

		if user == nil{
			return nil
		}
	*/

	//new_user := User{Username: username, Email:email, Name: name, Firstname:firstname,Lastname: lastname,Phonenumber: phonenumber,  Photo: photo, Roleid: roleId, Status: 1, Sessionid: ``,Missed: 0,Misseduniqueid: ``, Unreadmessages: 0,Homepageurl: homepageurl, Ver: ver}
	//users = append(users, new_user)
	err := deleteUser(username)
	//dbUpsertUser(new_user,func_name)
	if err == nil {

		//writeFileUsers()
		fmt.Println(func_name, "deleteUser()", username, "ok")

	}

	err = dbDeleteUser(username)

	if err == nil {
		fmt.Println(func_name, "dbDeleteUser()", username, "ok")
	}

	return nil
}

func adminUpdateUser_dep(username, roleId string) error {

	fmt.Println(`7934: START adminUpdateUser(`,
		//`firstname=`,firstname,
		//`lastname=`,lastname,
		//`name=`, name,
		`username=`, username,
		//`email=`,email,
		`roleId=`, roleId,
	//`photo=`,photo,
	//`phonenumber=`,phonenumber,
	//`homepageurl=`, homepageurl
	)

	//readFileUsers()
	func_name := `adminUpdateUser`
	fmt.Println(func_name, "(", username, roleId, ")")

	_, user := findUserByUsername(username)
	if user == nil {
		return errors.New(fmt.Sprintf(`username %s is invalid.`, username))
	}

	/*
		i, user = findUserByEmail(email)
		if user != nil{
			return i,user
		}
	*/
	/*
		ver := time.Now().Format(time.RFC3339)
		new_user := User{Username: username, Email:email, Name: name, Firstname:firstname,Lastname: lastname,Phonenumber: phonenumber,  Photo: photo, Roleid: roleId, Status: 1, Sessionid: ``,Missed: 0,Misseduniqueid: ``, Unreadmessages: 0,Homepageurl: homepageurl, Ver: ver}
		users = append(users, new_user)
	*/
	user.Ver = time.Now().Format(time.RFC3339)
	user.Roleid = roleId

	err := dbUpsertUser(*user, func_name)
	if err != nil {
		fmt.Println(string(colorRed), err.Error())
		fmt.Println(string(colorReset))
		return errors.New(fmt.Sprintf(`dbUpsertUser failed %s.`, username))

	}

	//writeFileUsers()

	//i,user = findUserByUsername(username)

	//return i,user
	return nil

}

func adminUpdateUser(username, roleId, firstname, lastname string) error {

	fmt.Println(`7934: START adminUpdateUser(`,
		`firstname=`, firstname,
		`lastname=`, lastname,
		//`name=`, name,
		`username=`, username,
		//`email=`,email,
		`roleId=`, roleId,
	//`photo=`,photo,
	//`phonenumber=`,phonenumber,
	//`homepageurl=`, homepageurl
	)

	//readFileUsers()
	func_name := `adminUpdateUser`
	fmt.Println(func_name, "(", username, roleId, ")")

	_, user := findUserByUsername(username)
	if user == nil {
		return errors.New(fmt.Sprintf(`username %s is invalid.`, username))
	}

	/*
		i, user = findUserByEmail(email)
		if user != nil{
			return i,user
		}
	*/
	/*
		ver := time.Now().Format(time.RFC3339)
		new_user := User{Username: username, Email:email, Name: name, Firstname:firstname,Lastname: lastname,Phonenumber: phonenumber,  Photo: photo, Roleid: roleId, Status: 1, Sessionid: ``,Missed: 0,Misseduniqueid: ``, Unreadmessages: 0,Homepageurl: homepageurl, Ver: ver}
		users = append(users, new_user)
	*/
	user.Ver = time.Now().Format(time.RFC3339)
	user.Roleid = roleId
	user.Firstname = firstname
	user.Lastname = lastname
	
	user.Name = firstname + " " + lastname


	err := dbUpsertUser(*user, func_name)
	if err != nil {
		fmt.Println(string(colorRed), err.Error())
		fmt.Println(string(colorReset))
		return errors.New(fmt.Sprintf(`dbUpsertUser failed %s.`, username))

	}

	//writeFileUsers()

	//i,user = findUserByUsername(username)

	//return i,user
	return nil

}

func get_admin_pass_manager_conf() (string, string) {
	//file, err := os.Open("/etc/asterisk/manager.conf")
	file, err := os.Open(fmt.Sprintf(`%s/manager.conf`, config[`asterisk_conf_path`]))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	/* looking for:
	[admin]
	secret = EHn6Yj7z9iz6
	*/
	admin := false
	secret := ``
	for scanner.Scan() {

		//fmt.Println(scanner.Text())
		str := scanner.Text()
		str = strings.Replace(str, ` = `, `=`, 1)
		//fmt.Println(`str=`,str)

		if admin {
			if strings.Contains(str, `secret=`) {
				secret = strings.Replace(str, `secret=`, ``, 1)
				break
			}
		} else {
			if strings.Contains(str, `[admin]`) {
				admin = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(`admin=`, admin, `secret=`, secret)
	return `admin`, secret

}

func addExtensionHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Redirect(w, r, "/extensions", http.StatusSeeOther)
        return
    }

    number := r.FormValue("number")
    displayname := r.FormValue("displayname")
    secret := r.FormValue("secret")

    // Dropdown values: "yes"/"no" with defaults
    public := r.FormValue("public")
    if public == "" {
        public = "yes"
    }
    canBlock := r.FormValue("canblock")
    if canBlock == "" {
        canBlock = "yes"
    }
    canUnlist := r.FormValue("canunlist")
    if canUnlist == "" {
        canUnlist = "yes"
    }
    canAddList := r.FormValue("canaddlist")
    if canAddList == "" {
        canAddList = "yes"
    }
    canDial := r.FormValue("candial")
    if canDial == "" {
        canDial = "yes"
    }

    // Check if extension already exists
    var exists string
    err := db.QueryRow("SELECT number FROM extensions WHERE number = ?", number).Scan(&exists)
    if err != nil && err != sql.ErrNoRows {
        http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
        return
    }
    if exists != "" {
        http.Error(w, "Extension already exists", http.StatusBadRequest)
        return
    }

    
    _, err = db.Exec(`
        INSERT INTO extensions
        (number, displayname, secret, public, canblock, canunlist, canaddlist, candial)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
        number, displayname, secret, public, canBlock, canUnlist, canAddList, canDial, "xoftswitch",
    )
    if err != nil {
        fmt.Println("Failed to insert extension:", err)
        http.Error(w, "Failed to add extension", http.StatusInternalServerError)
        return
    }

    // Sync PJSIP configs
    if err := syncAORConfig(); err != nil {
        fmt.Println("Failed to sync AOR:", err)
    }
    if err := syncAuthConfig(); err != nil {
        fmt.Println("Failed to sync Auth:", err)
    }
    if err := syncEndpointConfig(); err != nil {
        fmt.Println("Failed to sync Endpoint:", err)
    }

    // Set ownership to asterisk:asterisk
    chownFiles := []string{
        "/etc/asterisk/pjsip.aor_custom.conf",
        "/etc/asterisk/pjsip.auth_custom.conf",
        "/etc/asterisk/pjsip.endpoint_custom.conf",
    }
    for _, file := range chownFiles {
        if err := chownToAsterisk(file); err != nil {
            fmt.Println("Failed to set ownership for", file, ":", err)
        }
    }

    // Reload PJSIP
    cmd := exec.Command("asterisk", "-rx", "pjsip reload")
    if err := cmd.Run(); err != nil {
        fmt.Println("Failed to reload PJSIP:", err)
    }

    http.Redirect(w, r, "/extensions", http.StatusSeeOther)
}

func reloadHandler(w http.ResponseWriter, r *http.Request) {
	if !requireAdminOrLogin(w, r) {
        return
    }
/*
	if r.Method != http.MethodPost {
        http.Redirect(w, r, "/extensions", http.StatusSeeOther)
        return
    }
*/
    // Reload PJSIP
/*
	cmd := exec.Command("asterisk", "-rx", "pjsip reload")
    if err := cmd.Run(); err != nil {
        fmt.Println("Failed to reload PJSIP:", err)
    }
*/
	//reloadXoftSwitch()
	amiGetConfigJson(glo_ami);
	//reloadXoftSwitch()

    http.Redirect(w, r, "/extensions", http.StatusSeeOther)
}

// -------------------- Helpers --------------------

// Ensure base [xoftswitch-*](!) contexts exist
func ensureBaseContexts(certPath, keyPath string) error {
	if err := ensureSection("/etc/asterisk/pjsip.endpoint_custom.conf", "[xoftswitch-endpoints](!)", fmt.Sprintf(`[xoftswitch-endpoints](!)
type = endpoint
allow = ulaw,alaw,gsm,g726,g722,h264,mpeg4,opus
allow_subscribe=yes
aggregate_mwi = no
bundle = yes
cos_audio = 5
cos_video = 4
context = from-internal
dtmf_mode = rfc4733
direct_media = yes
dtls_verify = no
dtls_setup = actpass
dtls_rekey = 0
dtls_cert_file = %s
dtls_private_key = %s
ice_support = yes
force_rport = yes
language = en
mwi_subscribe_replaces_unsolicited = yes
max_audio_streams = 1
max_video_streams = 1
media_use_received_transport=yes
media_encryption = dtls
media_use_received_transport = yes
media_encryption_optimistic = yes
one_touch_recording = on
qualify_frequency = 60
record_on_feature = apprecord
record_off_feature = apprecord
refer_blind_progress = yes
rewrite_contact = yes
rtcp_mux = yes
rtp_timeout = 30
rtp_timeout_hold = 300
rtp_keepalive = 0
rtp_symmetric = yes
send_connected_line = yes
send_pai = yes
timers = yes
timers_min_se = 90
tos_audio = ef
tos_video = af41
trust_id_inbound = yes
use_avpf = yes
user_eq_phone = no
`, certPath, keyPath)); err != nil {
		return err
	}

	if err := ensureSection("/etc/asterisk/pjsip.aor_custom.conf", "[xoftswitch-aor](!)", `[xoftswitch-aor](!)
type = aor
max_contacts = 10
remove_existing = yes
maximum_expiration = 7200
minimum_expiration = 60
qualify_frequency = 60
`); err != nil {
		return err
	}

	if err := ensureSection("/etc/asterisk/pjsip.auth_custom.conf", "[xoftswitch-auth](!)", `[xoftswitch-auth](!)
type = auth
auth_type = userpass
`); err != nil {
		return err
	}

	return nil
}

// Generic: ensure section exists
func ensureSection(filePath, sectionHeader, sectionContent string) error {
	data, _ := os.ReadFile(filePath)
	if strings.Contains(string(data), sectionHeader) {
		return nil
	}
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(sectionContent + "\n"); err != nil {
		return err
	}
	fmt.Println("Added section to", filePath, ":", sectionHeader)
	return nil
}

// Check if section exists
func sectionExists(filePath, section string) bool {
	data, _ := os.ReadFile(filePath)
	return strings.Contains(string(data), section)
}

// Append if not exists
func appendIfMissing(filePath, content, sectionHeader string) error {
	if !sectionExists(filePath, sectionHeader) {
		return appendToFile(filePath, content)
	}
	return nil
}

// Simple append helper
/*
func appendToFile_dep(filePath, content string) error {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content + "\n")
	return err
}
*/
// Append content to a file and ensure ownership
func appendToFile(filePath, content string) error {
    f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer f.Close()

    if _, err := f.WriteString(content + "\n"); err != nil {
        return err
    }

    // Ensure file ownership is asterisk:asterisk
    if err := chownAsterisk(filePath); err != nil {
        fmt.Println("Failed to set ownership on", filePath, err)
    }

    return nil
}

// Generic helper to set ownership to asterisk:asterisk
func chownAsterisk(path string) error {
    // Lookup UID and GID for asterisk user
    uid, gid, err := lookupUserGroup("asterisk", "asterisk")
    if err != nil {
        return err
    }

    return os.Chown(path, uid, gid)
}

// Lookup UID and GID by username and group
func lookupUserGroup(user, group string) (int, int, error) {
    u, err := exec.Command("id", "-u", user).Output()
    if err != nil {
        return 0, 0, err
    }
    g, err := exec.Command("getent", "group", group).Output()
    if err != nil {
        return 0, 0, err
    }

    uid, err := strconv.Atoi(strings.TrimSpace(string(u)))
    if err != nil {
        return 0, 0, err
    }

    gidFields := strings.Split(strings.TrimSpace(string(g)), ":")
    if len(gidFields) < 3 {
        return 0, 0, fmt.Errorf("invalid group entry")
    }
    gid, err := strconv.Atoi(gidFields[2])
    if err != nil {
        return 0, 0, err
    }

    return uid, gid, nil
}
// -------------------- Sync helpers --------------------

func chownToAsterisk(path string) error {
    u, err := user.Lookup("asterisk")
    if err != nil {
        return err
    }
    g, err := user.LookupGroup("asterisk")
    if err != nil {
        return err
    }

    uid, _ := strconv.Atoi(u.Uid)
    gid, _ := strconv.Atoi(g.Gid)

    return os.Chown(path, uid, gid)
}

// Sync AOR
func syncAORConfig() error {
    rows, err := db.Query("SELECT number FROM extensions")
    if err != nil {
        return err
    }
    defer rows.Close()

    var dbNumbers []string
    for rows.Next() {
        var n string
        if err := rows.Scan(&n); err != nil {
            return err
        }
        dbNumbers = append(dbNumbers, n)
    }

    for _, n := range dbNumbers {
        entry := fmt.Sprintf("[%s](xoftswitch-aor)\nmailboxes = %s@default\n", n, n)
        appendIfMissing("/etc/asterisk/pjsip.aor_custom.conf", entry, "["+n+"](xoftswitch-aor)")
    }

    //os.Chown("/etc/asterisk/pjsip.aor_custom.conf", 103, 103)
	
	

    return nil
}


// Sync AUTH
func syncAuthConfig() error {
	rows, err := db.Query("SELECT number, secret FROM extensions")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var n, secret string
		rows.Scan(&n, &secret)
		entry := fmt.Sprintf("[%s-auth](xoftswitch-auth)\npassword = %s\nusername = %s\n", n, secret, n)
		appendIfMissing("/etc/asterisk/pjsip.auth_custom.conf", entry, "["+n+"-auth](xoftswitch-auth)")
	}

	return nil
}

// Sync ENDPOINT
func syncEndpointConfig() error {
	rows, err := db.Query("SELECT number, displayname FROM extensions")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var n, displayname string
		rows.Scan(&n, &displayname)
		entry := fmt.Sprintf("[%s](xoftswitch-endpoints)\naors = %s\nauth = %s-auth\ncallerid = %s <%s>\noutbound_auth = %s-auth\nmailboxes = %s@default\n",
			n, n, n, displayname, n, n, n)
		appendIfMissing("/etc/asterisk/pjsip.endpoint_custom.conf", entry, "["+n+"](xoftswitch-endpoints)")
	}

	return nil
}



func loadExtensionsAuth() {
	//file, err := os.Open("/etc/asterisk/pjsip.auth.conf")
	file, err := os.Open(fmt.Sprintf(`%s/pjsip.auth.conf`, config[`asterisk_conf_path`]))

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	/* looking for:
	[2000-auth]
	type=auth
	auth_type=userpass
	password=PXC365
	username=2000
	*/
	//ret := map[string]string{}

	found_item := false
	username := ``
	password := ``

	for scanner.Scan() {

		//fmt.Println(scanner.Text())
		str := scanner.Text()
		str = strings.Replace(str, ` = `, `=`, 1)
		//fmt.Println(`str=`,str)

		if found_item {
			if strings.Contains(str, `password=`) {
				password = strings.Replace(str, `password=`, ``, 1)

			} else if strings.Contains(str, `username=`) {
				username = strings.Replace(str, `username=`, ``, 1)

			}
			if len(username) > 0 && len(password) > 0 {

				//ret[username] = password
				extensions_auth[username] = password

				found_item = false
				username = ``
				password = ``

			}

		} else {
			if strings.Contains(str, `[`) {

				found_item = true

			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//fmt.Println(`admin=`,admin,`secret=`,secret)
	//return ret
	//extensions_auth = ret
	writeFileExtensionsAuth()

}

func loadCustomExtensionsAuth() {
	//file, err := os.Open("/etc/asterisk/pjsip.auth.conf")

	file, err := os.Open(fmt.Sprintf(`%s/pjsip.auth_custom.conf`, config[`asterisk_conf_path`]))

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	/* looking for:
	[2000-auth]
	type=auth
	auth_type=userpass
	password=PXC365
	username=2000
	*/
	//ret := map[string]string{}

	found_key := false
	username := ``
	password := ``

	for scanner.Scan() {

		//fmt.Println(scanner.Text())
		str := scanner.Text()
		strings.Trim(str, ` `)
		if len(str) == 0 {
			continue
		}
		if strings.Contains(str, `(!)`) {
			found_key = false
			continue
		}

		str = strings.Replace(str, ` = `, `=`, 1)
		//fmt.Println(`str=`,str)

		if found_key {
			if strings.Contains(str, `password=`) {
				password = strings.Replace(str, `password=`, ``, 1)

			} else if strings.Contains(str, `username=`) {
				username = strings.Replace(str, `username=`, ``, 1)

			}
			if len(username) > 0 && len(password) > 0 {

				//ret[username] = password
				extensions_auth[username] = password
				found_key = false
				username = ``
				password = ``

			}

		} else {
			if strings.Contains(str, `[`) {

				found_key = true

			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//extensions_auth = ret
	writeFileExtensionsAuth()

}

func loadEndpoints() {
	//file, err := os.Open("/etc/asterisk/pjsip.auth.conf")
	//file, err := os.Open(fmt.Sprintf(`%s/pjsip.endpoint.conf`,`/etc/asterisk`))
	file, err := os.Open(fmt.Sprintf(`%s/pjsip.endpoint.conf`, config[`asterisk_conf_path`]))

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	/* looking for:
	[2000-auth]
	type=auth
	auth_type=userpass
	password=PXC365
	username=2000
	*/
	//ret := make(map[string]interface{})

	key := ``
	var key_data map[string]string

	for scanner.Scan() {

		//fmt.Println(scanner.Text())
		str := scanner.Text()
		str = strings.Replace(str, ` = `, `=`, 1)
		//fmt.Println(`521: str=`,str)

		re_key := regexp.MustCompile(`\[(.*)\]`)

		ma_key := re_key.FindStringSubmatch(str)
		//fmt.Println(`ma_key=`,ma_key)
		if len(ma_key) == 2 {
			//must key line

			key = ma_key[1]
			key_data = make(map[string]string)
			endpoints[key] = key_data
		} else {

			//must name value line
			if len(key) > 0 {
				re_nameval := regexp.MustCompile("(.*)=(.*)")

				ma_nameval := re_nameval.FindStringSubmatch(str)

				if len(ma_nameval) == 3 {
					key_data[ma_nameval[1]] = ma_nameval[2]
				}

			}

		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//fmt.Println(`admin=`,admin,`secret=`,secret)

	//return ret
	//endpoints = ret
	writeFileEndpoints()
}

func loadCustomEndpoints() {
	//file, err := os.Open("/etc/asterisk/pjsip.auth.conf")
	//file, err := os.Open(fmt.Sprintf(`%s/pjsip.endpoint.conf`,`/etc/asterisk`))
	file, err := os.Open(fmt.Sprintf(`%s/pjsip.endpoint_custom.conf`, config[`asterisk_conf_path`]))

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	/* looking for:
	[2000-auth]
	type=auth
	auth_type=userpass
	password=PXC365
	username=2000
	*/
	//ret := make(map[string]interface{})
	re_key := regexp.MustCompile(`\[(.*)\]`)
	found_key := false
	key := ``
	var key_data map[string]string

	for scanner.Scan() {

		//fmt.Println(scanner.Text())
		str := scanner.Text()
		strings.Trim(str, ` `)
		if len(str) == 0 {
			continue
		}

		if strings.Contains(str, `(!)`) {
			found_key = false
			continue
		}

		str = strings.Replace(str, ` = `, `=`, 1)
		//fmt.Println(`521: str=`,str)

		ma_key := re_key.FindStringSubmatch(str)
		//fmt.Println(`ma_key=`,ma_key)
		if len(ma_key) == 2 {
			//must key line
			found_key = true
			key = ma_key[1]
			key_data = make(map[string]string)
			endpoints[key] = key_data
		} else {
			if found_key {
				//must name value line
				if len(key) > 0 {
					re_nameval := regexp.MustCompile("(.*)=(.*)")

					ma_nameval := re_nameval.FindStringSubmatch(str)

					if len(ma_nameval) == 3 {
						key_data[ma_nameval[1]] = ma_nameval[2]
					}

				}
			}

		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//endpoints = ret
	writeFileEndpoints()
}

func upsertExtensionsCustomConf(exten, dialplan string) {
	func_name := `upsertExtensionsCustomConf()`
	//file, err := os.Open("/etc/asterisk/pjsip.auth.conf")
	filename := fmt.Sprintf(`%s/extensions_custom.conf`, config[`asterisk_conf_path`])
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	//defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	/* looking for:
	[2000-auth]
	type=auth
	auth_type=userpass
	password=PXC365
	username=2000
	*/
	//ret := map[string]string{}

	//exten => 400,1,AGI(agi://localhost:8181)

	found_entry := false
	//username := ``
	//password := ``

	//check if entry exists

	for scanner.Scan() {

		//fmt.Println(scanner.Text())
		str := scanner.Text()
		strings.Trim(str, ` `)
		fmt.Println(func_name, `1344: str=`, str)
		if len(str) == 0 {
			continue
		}
		/*
			if strings.Contains(str,`(!)`){
				found_key = false
				continue
			}
		*/
		//remove spaces for easier parsing
		str_no_spaces := strings.ReplaceAll(str, ` `, ``)
		//exten => 400,1,AGI(agi://localhost:8181)
		if strings.Contains(str_no_spaces, fmt.Sprintf("exten=>%s,", exten)) {
			//found_key = false
			//continue
			found_entry = true

			break
		}

		//TODO: check for pattern like 3XX

	}

	if err := scanner.Err(); err != nil {
		//log.Fatal(err)
		fmt.Println(func_name, `err=`, err)
	}

	file.Close()

	//extensions_auth = ret
	//writeFileExtensionsAuth()
	fmt.Println(func_name, `1381: found_entry=`, found_entry)
	if !found_entry {
		//append
		f, err := os.OpenFile(filename,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		dp := fmt.Sprintf("\n%s", dialplan)
		if _, err := f.WriteString(dp); err == nil {
			//fmt.Println(err)
			amiCommand(glo_ami, `module reload`)
		}

	}

}

func get_public_ip_hostname() (string, string) {
	/*
		[0.0.0.0-tls]
		type=transport
		protocol=tls
		bind=0.0.0.0:5061
		external_media_address=136.179.36.235
		external_signaling_address=136.179.36.235
		ca_list_file=/etc/pki/tls/certs/ca-bundle.crt
		cert_file=/etc/asterisk/keys/fpbx15.peerxc.com-fullchain.crt
		priv_key_file=/etc/asterisk/keys/fpbx15.peerxc.com.key
		method=default
		verify_client=yes
		verify_server=yes
		allow_reload=no
		tos=cs3
		cos=3
		local_net=172.16.6.0/23
	*/
	file, err := os.Open(fmt.Sprintf(`%s/pjsip.transports.conf`, config[`asterisk_conf_path`]))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	/* looking for:
	[admin]
	secret = EHn6Yj7z9iz6
	*/
	found := false
	public_ip := ``
	public_hostname := ``
	for scanner.Scan() {

		//fmt.Println(scanner.Text())
		str := scanner.Text()
		str = strings.Replace(str, ` = `, `=`, 1)
		//fmt.Println(`str=`,str)

		if found {

			if strings.Contains(str, `external_signaling_address=`) {
				//external_signaling_address=136.179.36.235
				public_ip = strings.Replace(str, `external_signaling_address=`, ``, 1)
				//break
			} else if strings.Contains(str, `priv_key_file=/etc/asterisk/keys/`) {
				//priv_key_file=/etc/asterisk/keys/fpbx15.peerxc.com.key
				public_hostname = strings.Replace(str, `priv_key_file=/etc/asterisk/keys/`, ``, 1)
				public_hostname = strings.Replace(public_hostname, `.key`, ``, 1)
				//break
			}
		} else {
			if strings.Contains(str, `[`) {
				if strings.Contains(str, `[0.0.0.0-tls]`) {
					found = true

				} else {
					found = false
				}

			}
		}

		if len(public_ip) == 0 {

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(`public_ip=`, public_ip, `public_hostname=`, public_hostname)

	if len(public_ip) == 0 {
		json_ip2 := getip2()
		map_ip2 := map[string]string{}
		//RESULT: {"status":"success","country":"United States","countryCode":"US","region":"NV","regionName":"Nevada","city":"Las Vegas","zip":"89101","lat":36.1685,"lon":-115.1164,"timezone":"America/Los_Angeles","isp":"CenturyLink Communications, LLC","org":"CenturyLink Communications, LLC","as":"AS209 CenturyLink Communications, LLC","query":"71.222.1.109"}
		err := json.Unmarshal([]byte(json_ip2), &map_ip2)
		if err == nil {
			query, ok_query := map_ip2[`query`]
			fmt.Println(`query=`, query, `ok_query=`, ok_query)
			if ok_query {
				public_ip = map_ip2[`query`]
				public_hostname = map_ip2[`query`]
			}
		}
	}

	return public_ip, public_hostname

}

func getip2() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}
	var ip IP
	json.Unmarshal(body, &ip)
	// fmt.Print(ip.Query)
	return ip.Query
}




/*
	 func makeHTTPServer() *http.Server {
	    mux := &http.ServeMux{}
	    //mux.HandleFunc("/", handleIndex)
		mux.HandleFunc("/",homePage)
		mux.HandleFunc("/join",joinPage)
		mux.HandleFunc("/api",apiPage)
		mux.HandleFunc("/logo",logoPage)


		http.HandleFunc("/ws", wsEndpoint)

	    // set timeouts so that a slow or malicious client doesn't
	    // hold resources forever
	    return &http.Server{
	        ReadTimeout:  5 * time.Second,
	        WriteTimeout: 5 * time.Second,
	        IdleTimeout:  120 * time.Second,
	        Handler:      handler,
	    }
	}
*/
func startWSServer(server *http.Server) {
	fmt.Println("3604: START WS :8180")
	setupRoutes()
	//fmt.Println("Listening :8180")
	//log.Fatal(http.ListenAndServe(":8180",nil))
	//log.Fatal(server.ListenAndServe(":8180",nil))
	/*
			err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
		    if err != nil {
		        log.Fatal("ListenAndServe: ", err)
		    }
	*/
	//err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {

		panic(err)
	} else {
		fmt.Println("application stopped gracefully")
	}

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func startWSServer_dep(server *http.Server) {
	fmt.Println("3631: START WS :8180")
	setupRoutes()
	//fmt.Println("Listening :8180")
	log.Fatal(http.ListenAndServe(":8180", nil))

	//log.Fatal(server.ListenAndServe(":8180",nil))
	/*
			err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
		    if err != nil {
		        log.Fatal("ListenAndServe: ", err)
		    }
	*/
	//err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)

	//http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/www.yourdomain.com/fullchain.pem", "/etc/letsencrypt/live/www.yourdomain.com/privkey.pem", nil)

	//if err := http.ListenAndServeTLS(":8180", fmt.Sprintf("/etc/asterisk/keys/%s/fullchain.pem",hostname), fmt.Sprintf("/etc/asterisk/keys/%s/privkey.pem",hostname), nil); err != nil && !errors.Is(err, http.ErrServerClosed) {
	//if err := http.ListenAndServeTLS(":8180", fmt.Sprintf("%s/fullchain.pem",config[`cert_dir`]), fmt.Sprintf("%s/private.pem",config[`cert_dir`]), nil); err != nil && !errors.Is(err, http.ErrServerClosed) {

	//USE THIS FOR HTTPS BUT WILL WORK ONLY FOR WIFI NOT MOD+BILE DATA(5G)
	/*if err := http.ListenAndServeTLS(":8180", fmt.Sprintf("%s/certificate.pem",config[`cert_dir`]), fmt.Sprintf("%s/webserver.key",config[`cert_dir`]), nil); err != nil && !errors.Is(err, http.ErrServerClosed) {

		panic(err)
	} else {
		fmt.Println("application stopped gracefully")
	}
	*/

}

func startWSServerTLS() {
	
	//Log(LOG_INFO, "21529: START WS :8180")
	fmt.Println("3664: START WS :8180")
	setupRoutes()
	
	
	//USE THIS FOR HTTPS BUT WILL WORK ONLY FOR WIFI NOT MOBILE DATA(5G..)
	if err := http.ListenAndServeTLS(":8180", fmt.Sprintf("%s/certificate.pem", config[`cert_dir`]), fmt.Sprintf("%s/webserver.key", config[`cert_dir`]), nil); err != nil && !errors.Is(err, http.ErrServerClosed) {

		panic(err)
	} else {
		fmt.Println("application stopped gracefully")
	}

}

func stopWSServer(ctx context.Context, server *http.Server) {
	fmt.Println("STOP WS")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	} else {
		fmt.Println("application shutdowned")
	}

}

func startWSClient(done chan bool) {
	fmt.Println("START startWS")
	//var addr = flag.String("addr", "localhost:8180", "http service address")
	var addr = flag.String("addr", ":8180", "http service address")
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	//done_ws := make(chan struct{})

	//go func() {
	defer close(done)
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
		//message_str := string(message)

		//fmt.Println("<-", msg_str, ws.RemoteAddr())

		var result Map

		json.Unmarshal([]byte(message), &result)

		switch result["type"] {
		case "PING":
			err := c.WriteMessage(websocket.TextMessage, []byte(`{"type":"REGISTER","exten":"AGISVR"}`))
			if err != nil {
				fmt.Println("write:", err)
				return
			}

		}
	}
	fmt.Println("END startWS")
	//<- done
	//}()

}

func startAMI() {
	fmt.Println("START AMI")
	//fmt.Println("Listening :8181")
	readFileAMIContactStatus()

	//readFileExtensions()
	dbLoadExtensions()
	//extensions = []Extension{}

	dbMyContactsInitIfNeeded()

	done := make(chan bool)
	ami, err := gami.Dial("127.0.0.1:5038")
	if err != nil {
		log.Fatal(err)
	}

	glo_ami = ami
	defer ami.Close()

	ami.Run()

	//install manager
	go func() {
		for {
			select {
			//handle network errors
			case err := <-ami.NetError:
				fmt.Println("24: Network Error:", err)
				//try new connection every second
				<-time.After(time.Second)
				if err := ami.Reconnect(); err == nil {
					//call start actions
					ami.Action(gami.Params{"Action": "Events", "EventMask": "on"})
				}

			case err := <-ami.Error:
				fmt.Println("33: error:", err)
			//wait events and process
			case ev := <-ami.Events:

				//fmt.Sprintf("38: EventType: %v", event.New(ev))
				//fmt.Println("2172.1: EventType:", event.New(ev))
				event_type := event.New(ev)
				if ev.ID != `RTCPReceived` && ev.ID != `RTCPSent` {
					/*
					fmt.Println("3647.1: EventType:", event_type)
					fmt.Println("3647.2: ID=", ev.ID, "Privilege=", ev.Privilege, "Params=", ev.Params)
					*/
					
					Log(LOG_DEBUG, "3647.1: EventType:", event_type)
					Log(LOG_DEBUG, "3647.2: ID=", ev.ID, "Privilege=", ev.Privilege, "Params=", ev.Params)


				}

				switch ev.ID {
				case `SoftHangupRequest`:
					//Detect missed call
					//ex: ID= SoftHangupRequest Privilege= [call all] Params= map[Accountcode: Calleridname:405 Calleridnum:405 Cause:32 Channel:PJSIP/405-0000019a Channelstate:6 Channelstatedesc:Up Connectedlinename:<unknown> Connectedlinenum:<unknown> Context:from-internal Exten:100 Language:en Linkedid:1688876389.416 Priority:1 Uniqueid:1688876389.416]
					calleridnum, ok_calleridnum := ev.Params["Calleridnum"]
					exten, ok_exten := ev.Params["Exten"]
					uniqueid, ok_uniqueid := ev.Params["Uniqueid"]

					if ok_exten && ok_calleridnum && ok_uniqueid {
						//fmt.Println(string(colorRed), "Missed call to", exten,"from",calleridnum, "uniqueid",uniqueid)
						//fmt.Println(string(colorReset))
						dbUpdateUserMissedCall(exten, calleridnum, uniqueid)
					}
				case `ContactStatus`:

					//ex: ID= ContactStatus Privilege= [system all] Params= map[Aor:400 Contactstatus:Reachable Endpointname:400 Roundtripusec:54995 Uri:sip:365m2y45@71.222.1.109:64576;transport=ws]

					aor, ok_aor := ev.Params["Aor"]
					if ok_aor {
						contactstatusmaptable[aor] = ev.Params
						writeFileAMIContactStatus()
						onAMIContactStatus(ev.Params)

					}

				case `AorListComplete`:
					onAMIAorListComplete()

				case `AorList`:
					//NOTE: if the Contacts field is NO empty it's Available or Reachable
					//2023/05/23 03:29:25 456: ID= AorList Privilege= [] Params= map[Actionid:1684812565726458195 Authenticatequalify:false Contacts: Defaultexpiration:3600 Mailboxes: Maxcontacts:1 Maximumexpiration:7200 Minimumexpiration:60 Objectname:405 Objecttype:aor Outboundproxy: Qualifyfrequency:60 Qualifytimeout:3.000000 Removeexisting:true Removeunavailable:false Supportpath:false Voicemailextension:]
					//2023/05/23 03:29:25 456: ID= AorList Privilege= [] Params= map[Actionid:1684812565726458195 Authenticatequalify:false Contacts:400/sip:byceobya@71.222.1.109:57135;transport=ws Defaultexpiration:3600 Mailboxes:400@default Maxcontacts:5 Maximumexpiration:7200 Minimumexpiration:60 Objectname:400 Objecttype:aor Outboundproxy: Qualifyfrequency:60 Qualifytimeout:3.000000 Removeexisting:true Removeunavailable:false Supportpath:false Voicemailextension:]

					onAMIAorList(ev.Params)

				case `CoreShowChannel`:

					/*ex: 2023/06/27 20:25:29 2172.2: ID= CoreShowChannel Privilege= [] Params= map[
					Accountcode:
					Actionid:1687897529123997076
					Application:AppDial
					Applicationdata:(Outgoing Line)
					Bridgeid:fff2a768-016e-4500-af8d-2f175449b3b5
					Calleridname:405
					Calleridnum:405
					Channel:PJSIP/405-000000d4
					Channelstate:6 Channelstatedesc:Up
					Connectedlinename:100 XO
					Connectedlinenum:100
					Context:from-internal
					Duration:00:01:13 Exten:
					Language:en
					Linkedid:1687897455.215
					Priority:1
					Uniqueid:1687897455.216
					]
					*/

					//aor,ok_aor := ev.Params["Aor"]
					//if ok_aor{
					//contactstatusmaptable[aor] = ev.Params
					//writeFileAMIContactStatus()
					onAMICoreShowChannel(ev.Params)

					//}
				case "Reload":
					// Typical params: Module, Status, Message
					mod := ev.Params["Module"]
					status := ev.Params["Status"]
					msg := ev.Params["Message"]
					Log(LOG_INFO, "AMI Reload event: module=", mod, " status=", status, " msg=", msg)
				
					// Call your post-reload hook once you’ve seen what you care about.
					// For example, trigger after core/pbx/pjsip reloads:
					if mod == "" || mod == "pbx" || strings.HasPrefix(mod, "res_pjsip") || mod == "core" {
						//onAMIAsteriskReload(ev.Params) // your function
						amiGetConfigJson(glo_ami);
					}

				}

				/*
					type AMIEvent struct {
						//Identification of event Event: xxxx
						ID        string
						Privilege []string
						// Params  of arguments received
						Params map[string]string
					}
				*/
			}
		}
	}()

	//if err := ami.Login("xoftphone", "625f975db835748deda7d646803684d5"); err != nil {
	if err := ami.Login(config[`ami_username`], config[`ami_password`]); err != nil {
		log.Fatal("44:", err)
	}

	// Action method
	rsPing, rsActioID, rsErr := ami.Action(gami.Params{"Action": "Ping"})
	if rsErr != nil {

		log.Fatal("51: FATAL rsErr:", rsErr)
	}
	fmt.Println("53: rsActioID", rsActioID)

	// Synchronous response
	fmt.Println(<-rsPing)

	// Asynchronous response
	go func() {
		fmt.Println("60: rsActioID", <-rsPing)
	}()

	if _, _, err := ami.Action(gami.Params{"Action": "Events", "EventMask": "on"}); err != nil {
		log.Fatal("53: FATAL  err", err)
	}

	fmt.Println("67: ping:", <-rsPing)
	amiGetConfigJson(ami)

	/* ROM THIS WORKS but maybe hard to parse, use json instead
	rsGetConfig, rsGetConfigActionID, rsGetConfigErr := ami.Action(gami.Params{"Action":"GetConfig","Filename":"extensions.conf","Category":"default"})
	if rsGetConfigErr != nil{
		fmt.Println("76: rsGetConfigErr=",rsGetConfigErr)

	}else{

		fmt.Println("rsGetConfigActionID=",rsGetConfigActionID)
		fmt.Println(<-rsGetConfig)
	}
	*/

	/*
		rsListCategories, rsListCategoriesID, rsListCategoriesErr := ami.Action(gami.Params{"Action":"ListCategories","Filename":"extensions.conf"})
		if rsListCategoriesErr != nil{
			fmt.Println("76: rsGetConfigErr=",rsListCategoriesErr)

		}else{

			fmt.Println("rsListCategoriesID=",rsListCategoriesID)
			fmt.Println(<-rsListCategories)
		}
	*/

	//	rsGetConfig, rsGetConfigActionID, rsGetConfigErr := ami.Action(gami.Params{"Action":"GetConfigJSON","Filename":"extensions.conf","Category":"from-internal"})
	//rsGetConfig, rsGetConfigActionID, rsGetConfigErr := ami.Action(gami.Params{"Action":"GetConfigJSON","Filename":"extensions.conf"})

	//--> We can use this to populate extensions var without the status
	//rsGetConfig, rsGetConfigActionID, rsGetConfigErr := ami.Action(gami.Params{"Action":"GetConfigJSON","Filename":"pjsip.aor.conf"})

	rsCommand, _, rsCommandErr := ami.Action(gami.Params{"Action": "Pjsipshowaors"})

	if rsCommandErr == nil {
		fmt.Println("76: rsCommandErr=", rsCommandErr)
	} else {

		rs := <-rsCommand
		fmt.Println("549: rsCommand", <-rsCommand)
		fmt.Println("550: ++++++++++ ID", rs.ID, "Status=", rs.Status, "Params=", rs.Params)

	}

	<-done
}

// --> Crypto START
func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

//THIS WORKS/TESTED using encrypt_aes.go test

//func AesEncrypt(origData, key []byte) ([]byte, error) {
//func AesEncrypt(origData, key []byte) ([]byte, error) {
//func AesEncrypt(p_origData, p_key, p_iv string) ([]byte, error) {

func AesEncrypt(p_origData, p_key, p_iv string) (string, error) {

	key16 := p_key[0:16]
	block, err := aes.NewCipher([]byte(key16))
	if err != nil {
		return ``, err
	}
	blockSize := block.BlockSize()
	origData := PKCS5Padding([]byte(p_origData), blockSize)
	//iv := []byte("1234567812345678")
	//iv := []byte(p_iv)
	iv16 := p_iv[0:16]
	iv := []byte(iv16)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(origData))
	blockMode.CryptBlocks(encrypted, origData)
	b64str := base64.StdEncoding.EncodeToString(encrypted)

	return b64str, nil
}

// func AesDecrypt(encrypted, key []byte, iv []byte) ([]byte, error) {
// func AesDecrypt(p_encrypted, p_key , p_iv string) ([]byte, error) {
func AesDecrypt(p_encrypted, p_key, p_iv string) (string, error) {
	bytes_encrypted, _ := base64.StdEncoding.DecodeString(p_encrypted)
	//string_encrypted := string(bytes_encrypted)

	key16 := p_key[0:16]

	block, err := aes.NewCipher([]byte(key16))
	if err != nil {
		return ``, err
	}
	//iv := []byte(p_iv) //[]byte("1234567812345678")
	iv16 := p_iv[0:16]
	//fmt.Println(`iv16`,iv16)

	iv := []byte(iv16)
	//fmt.Println(`len(iv)=`,len(iv))
	blockMode := cipher.NewCBCDecrypter(block, iv)

	//encrypted := []byte(string_encrypted)
	origData := make([]byte, len(bytes_encrypted))
	//blockMode.CryptBlocks(origData, encrypted)
	blockMode.CryptBlocks(origData, bytes_encrypted)
	origData = PKCS5UnPadding(origData)
	return string(origData), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func amiGetConfigJson_dep(ami *gami.AMIClient) {
	rsGetConfig, rsGetConfigActionID, rsGetConfigErr := ami.Action(gami.Params{"Action": "GetConfigJSON", "Filename": "pjsip.endpoint.conf"})

	if rsGetConfigErr != nil {
		fmt.Println("76: rsGetConfigErr=", rsGetConfigErr)

	} else {

		fmt.Println("rsGetConfigActionID=", rsGetConfigActionID)
		//fmt.Println(<-rsGetConfig)
		body := fmt.Sprintf("%v", <-rsGetConfig)
		//RESULT: &{1684608737814743495 Success map[Actionid:1684608737814743495 Json:{"from-internal":{"include":"from-internal-noxfer","include":"from-internal-xfer","include":"bad-number","exten":"h,1,Macro(hangupcall)"}}]}
		// use capture groups
		/*
			// use capture groups
			phoneRECaps := `(\d{3})\-(\d{3})\-(\d{4})$`
			re = regexp.MustCompile(phoneRECaps)

			// caps is a slice of strings, where caps[0] matches the whole match
			// caps[1] == "202" etc
			matches := re.FindStringSubmatch(phone)

			// print out: there're 3 capture groups
			assert.Equal(re.NumSubexp(), 3)
			assert.Equal(matches[0], "202-555-0147")
			assert.Equal(matches[1], "202")
			assert.Equal(matches[2], "555")
			assert.Equal(matches[3], "0147")
			assert.ElementsMatch(matches, []string{"202-555-0147", "202", "555", "0147"})
		*/
		jsonRECaps := `Json:(.+)]}$`
		re := regexp.MustCompile(jsonRECaps)

		// caps is a slice of strings, where caps[0] matches the whole match
		// caps[1] == "202" etc
		matches := re.FindStringSubmatch(body)
		if len(matches) == 2 {

			//prettyJSON(matches[1])

			/* pjsip.aor.conf result
			{
			"405": {
				"type": "aor",
				"max_contacts": "1",
				"remove_existing": "yes",
				"maximum_expiration": "7200",
				"minimum_expiration": "60",
				"qualify_frequency": "60"
			},
			"400": {
				"type": "aor",
				"mailboxes": "400@default",
				"max_contacts": "5",
				"remove_existing": "yes",
				"maximum_expiration": "7200",
				"minimum_expiration": "60",
				"qualify_frequency": "60"
			},
			"sharecle": {
				"type": "aor",
				"qualify_frequency": "60",
				"contact": "sip:sharecle@sharecle.pstn.twilio.com"
			}

			*/

			/* pjsip.endpoint.conf result
			{
				"405": {
					"type": "endpoint",
					"aors": "405",
					"auth": "405-auth",
					"tos_audio": "ef",
					"tos_video": "af41",
					"cos_audio": "5",
					"cos_video": "4",
					"allow": "ulaw,alaw,gsm,g726,g722,h264,mpeg4",
					"context": "from-internal",
					"callerid": "405 <405>",
					"dtmf_mode": "rfc4733",
					"direct_media": "yes",
					"outbound_auth": "405-auth",
					"aggregate_mwi": "yes",
					"use_avpf": "no",
					"rtcp_mux": "no",
					"max_audio_streams": "1",
					"max_video_streams": "1",
					"bundle": "no",
					"ice_support": "no",
					"media_use_received_transport": "no",
					"trust_id_inbound": "yes",
					"user_eq_phone": "no",
					"send_connected_line": "yes",
					"media_encryption": "no",
					"timers": "yes",
					"timers_min_se": "90",
					"media_encryption_optimistic": "no",
					"refer_blind_progress": "yes",
					"refer_blind_progress": "yes",
					"rtp_timeout": "30",
					"rtp_timeout_hold": "300",
					"rtp_keepalive": "0",
					"send_pai": "yes",
					"rtp_symmetric": "yes",
					"rewrite_contact": "yes",
					"force_rport": "yes",
					"language": "en",
					"one_touch_recording": "on",
					"record_on_feature": "apprecord",
					"record_off_feature": "apprecord"
				},
				"400": {
					"type": "endpoint",
					"aors": "400",
					"auth": "400-auth",
					"tos_audio": "ef",
					"tos_video": "af41",
					"cos_audio": "5",
					"cos_video": "4",
					"allow": "ulaw,alaw,gsm,g726,g722,h264,mpeg4",
					"context": "from-internal",
					"callerid": "400 App <400>",
					"dtmf_mode": "rfc4733",
					"direct_media": "yes",
					"outbound_auth": "400-auth",
					"mailboxes": "400@default",
					"mwi_subscribe_replaces_unsolicited": "yes",
					"aggregate_mwi": "no",
					"use_avpf": "yes",
					"rtcp_mux": "yes",
					"max_audio_streams": "1",
					"max_video_streams": "1",
					"bundle": "yes",
					"ice_support": "yes",
					"media_use_received_transport": "yes",
					"trust_id_inbound": "yes",
					"user_eq_phone": "no",
					"send_connected_line": "yes",
					"media_encryption": "dtls",
					"timers": "yes",
					"timers_min_se": "90",
					"media_encryption_optimistic": "yes",
					"refer_blind_progress": "yes",
					"refer_blind_progress": "yes",
					"rtp_timeout": "30",
					"rtp_timeout_hold": "300",
					"rtp_keepalive": "0",
					"send_pai": "yes",
					"rtp_symmetric": "yes",
					"rewrite_contact": "yes",
					"force_rport": "yes",
					"language": "en",
					"one_touch_recording": "on",
					"record_on_feature": "apprecord",
					"record_off_feature": "apprecord",
					"dtls_verify": "fingerprint",
					"dtls_setup": "actpass",
					"dtls_rekey": "0",
					"dtls_cert_file": "/etc/asterisk/keys/fpbx.sharecle.com.crt",
					"dtls_private_key": "/etc/asterisk/keys/fpbx.sharecle.com.key"
				},
				"sharecle": {
					"type": "endpoint",
					"transport": "0.0.0.0-udp",
					"context": "from-pstn",
					"disallow": "all",
					"allow": "ulaw,alaw,gsm,g726,g722,h264,mpeg4",
					"aors": "sharecle",
					"send_connected_line": "false",
					"rtp_keepalive": "0",
					"language": "en",
					"outbound_auth": "sharecle",
					"user_eq_phone": "no",
					"t38_udptl": "no",
					"t38_udptl_ec": "none",
					"fax_detect": "no",
					"trust_id_inbound": "no",
					"t38_udptl_nat": "no",
					"direct_media": "no",
					"rtp_symmetric": "yes",
					"dtmf_mode": "auto"
				}
			}
			*/

			/*
				json_content,err := json.Marshal(prettyJSON);
				if err == nil{
					fmt.Println("json_content=",json_content)
				}
			*/
			//var aor Map

			err := json.Unmarshal([]byte(matches[1]), &endpoint)
			if err != nil {
				fmt.Println(err)
			} else {

				i := 0
				ver := time.Now().Format(time.RFC3339)
				for k, v := range endpoint {

					//fmt.Println("k=",k,"v=", v, "i=",i)

					switch v.(type) {
					case map[string]interface{}:
						vmap, ok_vmap := v.(map[string]interface{})
						//fmt.Println("ok_vmap",ok_vmap, "vmap=",vmap)
						if ok_vmap {
							
							displayname := ``
							callerid, ok_callerid := vmap["callerid"].(string)
							//fmt.Println("ok_callerid",ok_callerid, "callerid=",callerid)

							if ok_callerid {
								displayname = callerid
							}

							iexten, exten := findExtension(k)

							// ROM: IMPORTANT implement versioning so we can remove extensions not in configuration, also we should retain missedcall, misscalluniqueid data from dbload
							itemCtx, itemCancel := context.WithTimeout(context.Background(), 5*time.Second)
							if exten == nil {
								new_exten := Extension{Number: k, Displayname: displayname, Status: "", Source: "", Secret: "", Ver: ver}
								extensions = append(extensions, new_exten)
								if err := dbUpsertExtension(itemCtx, new_exten); err != nil {
									fmt.Println("2751:", err)
								}
							} else {
								extensions[iexten].Displayname = callerid
								extensions[iexten].Ver = ver
								if err := dbUpsertExtension(itemCtx, extensions[iexten]); err != nil {
									fmt.Println("2751:", err)
								}
							}
							itemCancel() // <- do NOT defer this inside the loop

							i++
						}

					default:
						fmt.Println("Unexpected!!!")

					}
				}

				dbDeleteExtensionsWithOldVersion_dep(ver)

				fmt.Println("extensions=", extensions)

			}

			//}
		}

	}

}

func amiGetConfigJson(ami *gami.AMIClient) {
	rsGetConfig, rsGetConfigActionID, rsGetConfigErr := ami.Action(
		gami.Params{"Action": "GetConfigJSON", "Filename": "pjsip.endpoint.conf"},
	)
	if rsGetConfigErr != nil {
		fmt.Println("76: rsGetConfigErr=", rsGetConfigErr)
		return
	}
	fmt.Println("rsGetConfigActionID=", rsGetConfigActionID)

	// Avoid blocking forever if AMI doesn't reply.
	var body string
	select {
	case msg := <-rsGetConfig:
		body = fmt.Sprintf("%v", msg)
	case <-time.After(5 * time.Second):
		fmt.Println("AMI GetConfigJSON timed out")
		return
	}

	// Extract the JSON payload using a non-greedy match.
	re := regexp.MustCompile(`Json:(.+?)]}$`)
	matches := re.FindStringSubmatch(body)
	if len(matches) != 2 {
		fmt.Println("could not extract JSON from AMI response")
		return
	}

	// Unmarshal into a generic map[string]interface{} (endpoint → settings map).
	var endpoint map[string]interface{}
	if err := json.Unmarshal([]byte(matches[1]), &endpoint); err != nil {
		fmt.Println("json.Unmarshal:", err)
		return
	}

	ver := time.Now().Format(time.RFC3339)

	// Parent context for the whole batch (optional overall cap).
	batchCtx, batchCancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer batchCancel()

	i := 0
	for k, v := range endpoint {
		select {
		case <-batchCtx.Done():
			fmt.Println("batch canceled:", batchCtx.Err())
			return
		default:
		}

		vmap, ok := v.(map[string]interface{})
		if !ok {
			fmt.Println("Unexpected value for key:", k)
			continue
		}

		// Derive display name from callerid when present.
		displayname := ""
		if callerid, ok := vmap["callerid"].(string); ok {
			displayname = callerid
		}

		iexten, extenPtr := findExtension(k)

		// Per-item timeout to avoid one slow DB op blocking the whole batch.
		itemCtx, itemCancel := context.WithTimeout(batchCtx, 5*time.Second)

		if extenPtr == nil {
			newExt := Extension{
				Number:      k,
				Displayname: displayname,
				Status:      "",
				Source:      "",
				Secret:      "",
				Ver:         ver,
			}
			extensions = append(extensions, newExt)
			if err := dbUpsertExtension(itemCtx, newExt); err != nil {
				fmt.Printf("upsert (new) failed for %s: %v\n", k, err)
			}
		} else {
			// Update in-memory slice and persist
			extensions[iexten].Displayname = displayname
			extensions[iexten].Ver = ver
			if err := dbUpsertExtension(itemCtx, extensions[iexten]); err != nil {
				fmt.Printf("upsert (existing) failed for %s: %v\n", k, err)
			}
		}

		itemCancel()
		i++
	}

	// Clean up any rows not touched in this run (version mismatch).
	if err := dbDeleteExtensionsWithOldVersion(batchCtx, ver); err != nil {
		fmt.Println("dbDeleteExtensionsWithOldVersion:", err)
	}

	fmt.Println("extensions=", extensions)
}

func amiPjsipshowaors(ami *gami.AMIClient) {
	rs, rsActionID, rsErr := ami.Action(gami.Params{"Action": "Pjsipshowaors"})

	if rsErr != nil {
		fmt.Println("76: rsErr=", rsErr)

	} else {

		fmt.Println("rsActionID=", rsActionID)
		//fmt.Println(<-rsGetConfig)
		body := fmt.Sprintf("%v", <-rs)

		jsonRECaps := `Json:(.+)]}$`
		re := regexp.MustCompile(jsonRECaps)

		// caps is a slice of strings, where caps[0] matches the whole match
		// caps[1] == "202" etc
		matches := re.FindStringSubmatch(body)
		if len(matches) == 2 {

			prettyJSON(matches[1])

			err := json.Unmarshal([]byte(matches[1]), &endpoint)
			if err != nil {
				fmt.Println(err)
			} else {

				i := 0
				ver := time.Now().Format(time.RFC3339)
				for k, v := range endpoint {

					//fmt.Println("k=",k,"v=", v, "i=",i)

					switch v.(type) {
					case map[string]interface{}:
						vmap, ok_vmap := v.(map[string]interface{})
						fmt.Println("ok_vmap", ok_vmap, "vmap=", vmap)
						if ok_vmap {
							callerid, ok_callerid := vmap["callerid"].(string)
							fmt.Println("ok_callerid", ok_callerid, "callerid=", callerid)

							displayname := k
							if ok_callerid {
								displayname = callerid
							}
							iexten, exten := findExtension(k)

							if exten == nil {
								//extensions = append(extensions, Extension{Number: k,Displayname: displayname,Status: ``,Source: ``})
								new_exten := Extension{Number: k, Displayname: displayname, Status: ``, Source: ``, Secret: ``, Ver: ver}
								extensions = append(extensions, new_exten)
								ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
    							defer cancel()
								err := dbUpsertExtension(ctx,new_exten)
								if err != nil {
									fmt.Println("2837: ", err)
								}
							} else {
								extensions[iexten].Displayname = callerid
							}

							i++
						}

					default:
						fmt.Println("Unexpected!!!")

					}

				}
				fmt.Println("extensions=", extensions)

			}

			//}
		}

	}

}

func amiUpdateconfigPjsipAor(ami *gami.AMIClient, exten string) bool {
	func_name := `amiUpdateconfigPjsipAor()`
	/* http://www.asteriskdocs.org/en/2nd_Edition/asterisk-book-html-chunk/asterisk-APP-F-42.html

	/*
	[xoftphone-aor](!)
	type = aor
	max_contacts = 5
	remove_existing = yes
	maximum_expiration = 7200
	minimum_expiration = 60
	qualify_frequency = 60

	[500](xoftphone-aor)
	mailboxes = 500@default

	*/

	//Start AOR

	rs, rsActionID, rsErr := ami.Action(gami.Params{
		"Action":      "Updateconfig",
		"SrcFilename": "pjsip.aor_custom.conf",
		"DstFilename": "pjsip.aor_custom.conf",

		"Action-000000":  "NewCat",
		"Cat-000000":     exten,
		"Options-000000": `inherit='xoftphone-aor'`,

		"Action-000001": "append",
		"Cat-000001":    exten,
		"Var-000001":    "mailboxes",
		"Value-000001":  fmt.Sprintf(`%s@default`, exten),
	})

	if rsErr != nil {
		fmt.Println(func_name, "2109: rsErr=", rsErr)

	} else {

		fmt.Println("rsActionID=", rsActionID)
		//fmt.Println(<-rsGetConfig)
		body := fmt.Sprintf("%v", <-rs)
		log.Print(func_name, `2117: amiAddExtension() body=`, body)
		if !strings.Contains(body, `Error`) {
			return true
		}

	}
	return false

}

func amiUpdateconfigPjsipAuth(ami *gami.AMIClient, exten, secret string) bool {
	func_name := `amiUpdateconfigPjsipAuth()`

	//START PJSIP AUTH
	/*
		[xoftphone-auth](!)
		type = auth
		auth_type = userpass

		[500-auth](xoftphone-auth)
		password = 04e5d62b2f71a1cde4e98be2c7ca822a
		username = 500
	*/
	//auth_cat := fmt.Sprintf(`%s-auth`,exten)
	cat_exten := fmt.Sprintf(`%s-auth`, exten)
	rs, rsActionID, rsErr := ami.Action(gami.Params{
		"Action":      "Updateconfig",
		"SrcFilename": "pjsip.auth_custom.conf",
		"DstFilename": "pjsip.auth_custom.conf",

		"Action-000000": "NewCat",
		//"Cat-000000": fmt.Sprintf(`%s-auth`,exten),
		//"Cat-000000": fmt.Sprintf("%s",auth_cat),
		"Cat-000000":     cat_exten,
		"Options-000000": `inherit='xoftphone-auth'`,

		"Action-000001": "append",
		"Cat-000001":    cat_exten,
		"Var-000001":    "password",
		"Value-000001":  secret,

		"Action-000002": "append",
		"Cat-000002":    cat_exten,
		"Var-000002":    "username",
		"Value-000002":  exten,
	})

	if rsErr != nil {
		fmt.Println(func_name, "2162: rsErr=", rsErr)

	} else {

		fmt.Println(func_name, "2166: rsActionID=", rsActionID)
		//fmt.Println(<-rsGetConfig)
		body := fmt.Sprintf("%v", <-rs)

		fmt.Println(func_name, `2169: amiAddExtension() body=`, body)
		if !strings.Contains(body, `Error`) {
			return true
		}

	}

	return false

}

func amiSendFlash(ami *gami.AMIClient, channel string) bool {
	/*
		SendFlash
		Synopsis
		Send a hook flash on a specific channel.
		Description
		Sends a hook flash on the specified channel.
		Syntax
		Action: SendFlash
		ActionID: <value>
		Channel: <value>
		[Receive:] <value>
		Arguments
		ActionID - ActionID for this transaction. Will be returned.
		Channel - Channel name to send hook flash to.
		Receive - Emulate receiving a hook flash on this channel instead of sending it out.

	*/
	func_name := `amiSendFlash`

	rs, rsActionID, rsErr := ami.Action(gami.Params{
		"Action":  "Sendflash",
		"Channel": channel,
	})

	if rsErr != nil {
		fmt.Println(func_name, "2845.1: rsErr=", rsErr)

	} else {

		fmt.Println(func_name, "2845.2: rsActionID=", rsActionID)
		//fmt.Println(<-rsGetConfig)
		body := fmt.Sprintf("%v", <-rs)

		fmt.Println(func_name, `2845.3: amiSendFlash() body=`, body)
		if !strings.Contains(body, `Error`) {
			return true
		}

	}

	return false
}

func amiPark(ami *gami.AMIClient, channel string) bool {
	/*
		Action: Park
		ActionID: <value>
		Channel: <value>
		[TimeoutChannel:] <value>
		[AnnounceChannel:] <value>
		[Timeout:] <value>
		[Parkinglot:] <value>

		Arguments
		ActionID - ActionID for this transaction. Will be returned.
		Channel - Channel name to park.
		TimeoutChannel - Channel name to use when constructing the dial string that will be dialed if the parked channel times out. If TimeoutChannel is in a two party bridge with Channel, then TimeoutChannel will receive an announcement and be treated as having parked Channel in the same manner as the Park Call DTMF feature.
		AnnounceChannel - If specified, then this channel will receive an announcement when Channel is parked if AnnounceChannel is in a state where it can receive announcements (AnnounceChannel must be bridged). AnnounceChannel has no bearing on the actual state of the parked call.
		Timeout - Overrides the timeout of the parking lot for this park action. Specified in milliseconds, but will be converted to seconds. Use a value of 0 to disable the timeout.
		Parkinglot - The parking lot to use when parking the channel


	*/
	func_name := `amiPark`
	fmt.Println("2945: START ", func_name, `channel=`, channel)

	rs, rsActionID, rsErr := ami.Action(gami.Params{
		"Action":  "Park",
		"Channel": channel,
		"Timeout": `120000`})

	if rsErr != nil {
		fmt.Println(func_name, "2845.1: rsErr=", rsErr)

	} else {

		fmt.Println(func_name, "2845.2: rsActionID=", rsActionID)
		//fmt.Println(<-rsGetConfig)
		body := fmt.Sprintf("%v", <-rs)

		fmt.Println(func_name, `2845.3: body=`, body)
		if !strings.Contains(body, `Error`) {
			return true
		}

	}

	return false
}

func amiAtxFer(ami *gami.AMIClient, channel, exten, context string) bool {
	/*
		Atxfer
		Synopsis
		Attended transfer.
		Description
		Attended transfer.

		Syntax
		Action: Atxfer
		ActionID: <value>
		Channel: <value>
		Exten: <value>
		Context: <value>
		Arguments
		ActionID - ActionID for this transaction. Will be returned.
		Channel - Transferer's channel.
		Exten - Extension to transfer to.
		Context - Context to transfer to.

		See Also
		Asterisk 16 ManagerEvent_AttendedTransfer


	*/
	func_name := `amiAtxFer`
	fmt.Println("2945: START ", func_name, `channel=`, channel)

	rs, rsActionID, rsErr := ami.Action(gami.Params{
		"Action":  "Atxfer",
		"Channel": channel,
		"Exten":   exten,
		"Context": context})

	if rsErr != nil {
		fmt.Println(func_name, "2845.1: rsErr=", rsErr)

	} else {

		fmt.Println(func_name, "2845.2: rsActionID=", rsActionID)
		//fmt.Println(<-rsGetConfig)
		body := fmt.Sprintf("%v", <-rs)

		fmt.Println(func_name, `2845.3: body=`, body)
		if !strings.Contains(body, `Error`) {
			return true
		}

	}

	return false
}

func amiMuteAudio(ami *gami.AMIClient, channel, direction, state string) bool {
	/*
		Synopsis
		Mute an audio stream.
		Description
		Mute an incoming or outgoing audio stream on a channel.
		Syntax
		Action: MuteAudio
		ActionID: <value>
		Channel: <value>
		Direction: <value>
		State: <value>
		Arguments
		ActionID - ActionID for this transaction. Will be returned.
		Channel - The channel you want to mute.
		Direction
		in - Set muting on inbound audio stream. (to the PBX)
		out - Set muting on outbound audio stream. (from the PBX)
		all - Set muting on inbound and outbound audio streams.
		State
		on - Turn muting on.
		off - Turn muting off.


	*/
	func_name := `amiMuteAudio`
	fmt.Println("2945: START ", func_name, `channel=`, channel)

	rs, rsActionID, rsErr := ami.Action(gami.Params{
		"Action":    "MuteAudio",
		"Channel":   channel,
		"Direction": direction,
		"State":     state})

	if rsErr != nil {
		fmt.Println(func_name, "2845.1: rsErr=", rsErr)

	} else {

		fmt.Println(func_name, "2845.2: rsActionID=", rsActionID)
		//fmt.Println(<-rsGetConfig)
		body := fmt.Sprintf("%v", <-rs)

		fmt.Println(func_name, `2845.3: body=`, body)
		if !strings.Contains(body, `Error`) {
			return true
		}

	}

	return false
}

func amiRedirect(ami *gami.AMIClient, channel, exten, context, extra_channel, extra_exten, extra_context string) bool {
	/*
		Description
		Redirect (transfer) a call.
		Syntax
		Action: Redirect
		ActionID: <value>
		Channel: <value>
		ExtraChannel: <value>
		Exten: <value>
		ExtraExten: <value>
		Context: <value>
		ExtraContext: <value>
		Priority: <value>
		ExtraPriority: <value>
		Arguments
		ActionID - ActionID for this transaction. Will be returned.
		Channel - Channel to redirect.
		ExtraChannel - Second call leg to transfer (optional).
		Exten - Extension to transfer to.
		ExtraExten - Extension to transfer extrachannel to (optional).
		Context - Context to transfer to.
		ExtraContext - Context to transfer extrachannel to (optional).
		Priority - Priority to transfer to.
		ExtraPriority - Priority to transfer extrachannel to (optional).



	*/

	func_name := `amiRedirect`

	/*
		params := map[string]string{"Action":"Redirect",
		"Channel":channel,
		"Exten":exten,
		"Priority": `1`}
	*/

	params := gami.Params{
		"Action":  "Redirect",
		"Channel": channel,
		//"ExtraChannel":extra_channel,
		"Exten":   exten,
		"Context": context, //`from-internal`,
		//"ExtraExten":extra_exten,
		"Priority": `1`,
		//"ExtraPriority": 1,
	}

	if len(extra_channel) > 0 && len(extra_exten) > 0 {

		/*rs, rsActionID, rsErr  := ami.Action(gami.Params{
			"Action":"Redirect",
			"Channel":channel,
			"ExtraChannel":extra_channel,
			"Exten":exten,
			"ExtraExten":extra_exten,
			"Priority": 1,
			"ExtraPriority": 1,
		})
		*/
		params = gami.Params{
			"Action":        "Redirect",
			"Channel":       channel,
			"ExtraChannel":  extra_channel,
			"Exten":         exten,
			"ExtraExten":    extra_exten,
			"Priority":      `1`,
			"Context":       context, //`from-internal`,
			"ExtraPriority": `1`,
			"ExtraContext":  extra_context,
		}

	} else if len(extra_channel) > 0 {
		/*rs, rsActionID, rsErr  := ami.Action(gami.Params{
			"Action":"Redirect",
			"Channel":channel,
			"ExtraChannel":extra_channel,
			"Exten":exten,
			//"ExtraExten":extra_exten,
			"Priority": 1,
			"ExtraPriority": 1,
		})
		*/
		params = gami.Params{
			"Action":       "Redirect",
			"Channel":      channel,
			"ExtraChannel": extra_channel,
			"Exten":        exten,
			//"ExtraExten":extra_exten,
			"Priority":      `1`,
			"Context":       context, //`from-internal`,
			"ExtraPriority": `1`,
			"ExtraContext":  extra_context,
		}

	} else if len(extra_exten) > 0 {
		/*
			rs, rsActionID, rsErr  := ami.Action(gami.Params{
				"Action":"Redirect",
				"Channel":channel,
				//"ExtraChannel":extra_channel,
				"Exten":exten,
				"ExtraExten":extra_exten,
				"Priority": 1,
				"ExtraPriority": 1,
			})
		*/
		params = gami.Params{
			"Action":  "Redirect",
			"Channel": channel,
			//"ExtraChannel":extra_channel,
			"Exten":         exten,
			"ExtraExten":    extra_exten,
			"Priority":      `1`,
			"Context":       context, //`from-internal`,
			"ExtraPriority": `1`,
			"ExtraContext":  extra_context,
		}
	} /*else{
		rs, rsActionID, rsErr  := ami.Action(gami.Params{
			"Action":"Redirect",
			"Channel":channel,
			//"ExtraChannel":extra_channel,
			"Exten":exten,
			//"ExtraExten":extra_exten,
			"Priority": 1,
			//"ExtraPriority": 1,
		})
	}
	*/

	/*
		gami.Params{
					"Action":"Redirect",
					"Channel":channel,
					"ExtraChannel":extra_channel,
					"Exten":exten,
					"ExtraExten":extra_exten,

				}
	*/

	/*
		rs, rsActionID, rsErr  := ami.Action(gami.Params{
			"Action":"Redirect",
			"Channel":channel,
			"ExtraChannel":extra_channel,
			"Exten":exten,
			"ExtraExten":extra_exten,
			//"Priority": 1,
			//"ExtraPriority": 1,
		})
	*/

	//gami.Params = params
	rs, rsActionID, rsErr := ami.Action(params)

	if rsErr != nil {
		fmt.Println(func_name, "2845.1: rsErr=", rsErr)

	} else {

		fmt.Println(func_name, "2845.2: rsActionID=", rsActionID)
		//fmt.Println(<-rsGetConfig)
		body := fmt.Sprintf("%v", <-rs)

		fmt.Println(func_name, `2845.3: body=`, body)
		if !strings.Contains(body, `Error`) {
			return true
		}

	}

	return false
}

func amiListCommands(ami *gami.AMIClient) bool {
	/*
		SendFlash
		Synopsis
		Send a hook flash on a specific channel.
		Description
		Sends a hook flash on the specified channel.
		Syntax
		Action: SendFlash
		ActionID: <value>
		Channel: <value>
		[Receive:] <value>
		Arguments
		ActionID - ActionID for this transaction. Will be returned.
		Channel - Channel name to send hook flash to.
		Receive - Emulate receiving a hook flash on this channel instead of sending it out.

	*/
	func_name := `amiListCommands`

	rs, rsActionID, rsErr := ami.Action(gami.Params{
		"Action": "ListCommands"})

	if rsErr != nil {
		fmt.Println(func_name, "2897.1: rsErr=", rsErr)

	} else {

		fmt.Println(func_name, "2897.2: rsActionID=", rsActionID)
		//fmt.Println(<-rsGetConfig)
		body := fmt.Sprintf("%v", <-rs)

		fmt.Println(func_name, `2897.3: amiSendFlash() body=`, body)
		if !strings.Contains(body, `Error`) {
			return true
		}

	}

	return false
}

func amiCoreShowChannels(ami *gami.AMIClient) bool {
	/*

		Synopsis
		List currently active channels.
		Description
		List currently defined channels and some information about them.
		Syntax
		Action: CoreShowChannels
		ActionID: <value>
		Arguments
		ActionID - ActionID for this transaction. Will be returned.



	*/
	func_name := `amiCoreShowChannels`

	rs, rsActionID, rsErr := ami.Action(gami.Params{
		"Action": "CoreShowChannels"})

	if rsErr != nil {
		fmt.Println(func_name, "2897.1: rsErr=", rsErr)

	} else {

		fmt.Println(func_name, "2897.2: rsActionID=", rsActionID)
		//fmt.Println(<-rsGetConfig)
		body := fmt.Sprintf("%v", <-rs)

		fmt.Println(func_name, `2897.3: amiCoreShowChannels() body=`, body)
		if !strings.Contains(body, `Error`) {
			return true
		}

	}

	return false
}

func amiReload(ami *gami.AMIClient, module string) bool {
	/*
			[Syntax]
		Action: Reload
		[ActionID:] <value>
		[Module:] <value>

		[Synopsis]
		Send a reload event.

		[Description]
		Send a reload event.

		[Arguments]
		ActionID
		    ActionID for this transaction. Will be returned.
		Module
		    Name of the module to reload.

		[See Also]
		ModuleLoad

	*/
	func_name := `amiReload`

	rs, rsActionID, rsErr := ami.Action(gami.Params{
		"Action": "Reload",
		"Module": module,
	})

	if rsErr != nil {
		fmt.Println(func_name, "2162: rsErr=", rsErr)

	} else {

		fmt.Println(func_name, "2166: rsActionID=", rsActionID)
		//fmt.Println(<-rsGetConfig)
		body := fmt.Sprintf("%v", <-rs)

		fmt.Println(func_name, `2169: amiReload() body=`, body)
		if !strings.Contains(body, `Error`) {
			return true
		}

	}

	return false
}

func amiCommand(ami *gami.AMIClient, cli_command string) bool {
	/*
			[Syntax]
		Action: Command
		[ActionID:] <value>
		Command: <value>

		[Synopsis]
		Execute Asterisk CLI Command.

		[Description]
		Run a CLI command.

		[Arguments]
		ActionID
		    ActionID for this transaction. Will be returned.
		Command
		    Asterisk CLI command to run.

		[See Also]
		Not available

	*/
	func_name := `amiCommand`

	rs, rsActionID, rsErr := ami.Action(gami.Params{
		"Action":  "Command",
		"Command": cli_command,
	})

	if rsErr != nil {
		fmt.Println(func_name, "2162: rsErr=", rsErr)

	} else {

		fmt.Println(func_name, "2166: rsActionID=", rsActionID)
		//fmt.Println(<-rsGetConfig)
		body := fmt.Sprintf("%v", <-rs)

		fmt.Println(func_name, `2169: amiCommand() body=`, body)
		if !strings.Contains(body, `Error`) {
			return true
		}

	}

	return false
}

func amiMailboxStatus(ami *gami.AMIClient, Mailbox string) bool {
	/*
			[Syntax]
		Action: MailboxStatus
		[ActionID:] <value>
		Mailbox: <value>

		[Synopsis]
		Check mailbox.

		[Description]
		Checks a voicemail account for status.
		Returns whether there are messages waiting.
		Message: Mailbox Status.
		Mailbox: <mailboxid>.
		Waiting: '0' if messages waiting, '1' if no messages waiting.

		[Arguments]
		ActionID
		    ActionID for this transaction. Will be returned.
		Mailbox
		    Full mailbox ID <mailbox>@<vm-context>.

		[See Also]
		MailboxCount

		[Privilege]
		call,reporting,all

		[List Responses]
		None

		[Final Response]
		None


	*/
	func_name := `amiMailboxStatus`

	rs, rsActionID, rsErr := ami.Action(gami.Params{
		"Action":  `MailboxStatus`,
		"Mailbox": Mailbox,
	})

	if rsErr != nil {
		fmt.Println(func_name, "2162: rsErr=", rsErr)

	} else {

		fmt.Println(func_name, "2166: rsActionID=", rsActionID)
		//fmt.Println(<-rsGetConfig)
		body := fmt.Sprintf("%v", <-rs)

		fmt.Println(func_name, `2169: amiCommand() body=`, body)
		if !strings.Contains(body, `Error`) {
			return true
		}

	}

	return false
}

func amiUpdateconfigPjsipEndpoint(ami *gami.AMIClient, exten string) bool {
	func_name := `amiUpdateconfigPjsipEndpoint()`

	//START PJSIP ENDPOINT
	/*
		[xoftphone-endpoints](!)
			type=endpoint

			tos_audio=ef
			tos_video=af41
			cos_audio=5
			cos_video=4
			allow=ulaw,alaw,gsm,g726,g722,h264,mpeg4
			context=from-internal
			dtmf_mode=rfc4733
			direct_media=yes
			mwi_subscribe_replaces_unsolicited=yes
			aggregate_mwi=no
			use_avpf=yes
			rtcp_mux=yes
			max_audio_streams=1
			max_video_streams=1
			bundle=yes
			ice_support=yes
			media_use_received_transport=yes
			trust_id_inbound=yes
			user_eq_phone=no
			send_connected_line=yes
			media_encryption=dtls
			timers=yes
			timers_min_se=90
			media_encryption_optimistic=yes
			refer_blind_progress=yes
			refer_blind_progress=yes
			rtp_timeout=30
			rtp_timeout_hold=300
			rtp_keepalive=0
			send_pai=yes
			rtp_symmetric=yes
			rewrite_contact=yes
			force_rport=yes
			language=en
			one_touch_recording=on
			record_on_feature=apprecord
			record_off_feature=apprecord
			dtls_verify=fingerprint
			dtls_setup=actpass
			dtls_rekey=0
			dtls_cert_file=/etc/asterisk/keys/fpbx.sharecle.com.crt
			dtls_private_key=/etc/asterisk/keys/fpbx.sharecle.com.key

		[500](xoftphone-endpoints)
		aors = 500
		auth = 500-auth
		callerid = 500 App <500>
		outbound_auth = 500-auth
		mailboxes = 500@default

	*/
	//auth_exten := fmt.Sprintf(`%s-auth`,exten)
	aut_exten := fmt.Sprintf(`%s-auth`, exten)
	rs, rsActionID, rsErr := ami.Action(gami.Params{
		"Action":      "Updateconfig",
		"SrcFilename": "pjsip.endpoint_custom.conf",
		"DstFilename": "pjsip.endpoint_custom.conf",

		"Action-000000":  "NewCat",
		"Cat-000000":     exten,
		"Options-000000": `inherit='xoftphone-endpoints'`,

		"Action-000001": "append",
		"Cat-000001":    exten,
		"Var-000001":    "aors",
		"Value-000001":  exten,

		"Action-000002": "append",
		"Cat-000002":    exten,
		"Var-000002":    "auth",
		//"Value-000002": fmt.Sprintf(`%s-auth`,exten),
		"Value-000002": aut_exten,
		//"Value-000002": fmt.Sprintf("%s",auth_exten),

		"Action-000003": "append",
		"Cat-000003":    exten,
		"Var-000003":    "callerid",
		"Value-000003":  fmt.Sprintf(`%s XO <%s>`, exten, exten),

		"Action-000004": "append",
		"Cat-000004":    exten,
		"Var-000004":    "outbound_auth",
		//"Value-000004":fmt.Sprintf(`%s-auth`,exten),
		"Value-000004": aut_exten,

		"Action-000005": "append",
		"Cat-000005":    exten,
		"Var-000005":    "mailboxes",
		"Value-000005":  fmt.Sprintf(`%s@default`, exten),
	})

	if rsErr != nil {
		fmt.Println(func_name, "2276: rsErr=", rsErr)

	} else {

		fmt.Println(func_name, "2280: rsActionID=", rsActionID)
		//fmt.Println(<-rsGetConfig)
		body := fmt.Sprintf("%v", <-rs)
		log.Print(func_name, `2283: amiUpdateconfigPjsipEndpoint() body=`, body)
		if !strings.Contains(body, `Error`) {
			return true
		}

	}

	return false

}

func amiAddPJSIPExtension(ami *gami.AMIClient, exten, secret, displayname string) bool {
	func_name := `amiAddPJSIPExtension()`

	if len(displayname) == 0 {
		displayname = exten
	}

	if len(displayname) > 25 {
		/*
					 s := "Hello World"
			    	fmt.Println(s[1:4])  // ell
		*/
		displayname = displayname[:25]
	}

	fmt.Println(func_name, `START 2296: exten=`, exten, `secret=`, secret, `displayname=`, displayname)
	if ami == nil || len(exten) == 0 || len(secret) == 0 {
		fmt.Println(func_name, `!!!!!!!!! 2265: INVALID params ami=`, ami, `exten=`, exten, `secret=`, secret)
		return false
	}
	ret := false
	ok_aor := amiUpdateconfigPjsipAor(ami, exten)

	if ok_aor {
		fmt.Println(func_name, `2308: amiUpdateconfigPjsipAor OK`)
		//ok_auth := amiUpdateconfigPjsipAuth(glo_ami,p_exten,uuid.New().String())
		ok_auth := amiUpdateconfigPjsipAuth(ami, exten, secret)

		if ok_auth {
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
    		defer cancel()
			fmt.Println(func_name, `2308: amiUpdateconfigPjsipAuth OK`)
			ok_endpoint := amiUpdateconfigPjsipEndpoint(ami, exten)
			if ok_endpoint {
				fmt.Println(func_name, `2308: amiUpdateconfigPjsipEndpoint OK`)
				//add dialplan
				dialplan := fmt.Sprintf(`exten => %s,1,AGI(agi://localhost:8181)`, exten)
				upsertExtensionsCustomConf(exten, dialplan)
				amiCommand(glo_ami, `module reload`)

				//set extension Source to xo if needed
				ver := time.Now().Format(time.RFC3339)
				i, ext := findExtension(exten)
				if i >= 0 && ext != nil {
					extensions[i].Source = `xo`
					extensions[i].Secret = secret
					extensions[i].Ver = ver
					fmt.Println(func_name, `2483: update extend=`, extensions[i])
					dbUpsertExtension(ctx,extensions[i])

				} else {
					//lets add now

					//new_exten := Extension{Number: exten, Displayname: fmt.Sprintf(`%s XO <%s>`,exten,exten),Secret: secret, Status: ``,Contacts: ``,Source:`xo`,Ver: ver}
					new_exten := Extension{Number: exten, Displayname: displayname, Secret: secret, Status: ``, Contacts: ``, Source: `xo`, Ver: ver}

					extensions = append(extensions, new_exten)
					err := dbUpsertExtension(ctx,new_exten)
					if err != nil {
						fmt.Println("3731: ", err)
					}
					fmt.Println(func_name, `2490: new exten=`, new_exten)

				}

				//writeFileExtensions()

				amiGetConfigJson(ami)
				loadCustomExtensionsAuth()
				loadCustomEndpoints()
				rsCommand, _, rsCommandErr := ami.Action(gami.Params{"Action": "Pjsipshowaors"})

				if rsCommandErr == nil {
					fmt.Println("76: rsCommandErr=", rsCommandErr)
				} else {

					rs := <-rsCommand
					fmt.Println("549: rsCommand", <-rsCommand)
					fmt.Println("550: ++++++++++ ID", rs.ID, "Status=", rs.Status, "Params=", rs.Params)

				}

				ret = true
			} else {
				fmt.Println(func_name, `2308: amiUpdateconfigPjsipEndpoint NOT OK`)

			}

		} else {
			fmt.Println(func_name, `2308: amiUpdateconfigPjsipAuth NOT OK`)

		}
	} else {
		fmt.Println(func_name, `2343: amiUpdateconfigPjsipAor NOT OK`)
	}

	return ret

}

func onAMIContactStatus(params map[string]string) {
	//Params= map[Aor:400 Contactstatus:Reachable Endpointname:400 Roundtripusec:54995 Uri:sip:365m2y45@71.222.1.109:64576;transport=ws]
	fmt.Println("++++++ START onAMIContactStatus(params=", params)
	json_params, err1 := json.Marshal(params)
	fmt.Println("json_params=", json_params)
	if err1 == nil {
		//ret_json := fmt.Sprintf(`{"total":%v,"page_num":%v,"items":%s}`,total,page_num,json_items)

		//jsonString, err2 := json.Marshal(ret_json)
		//fmt.Println(err)
		//if err2 == nil {
		//fmt.Println("-> CONTACT_STATUS", json_params)
		sendString := fmt.Sprintf(`{"type":"CONTACT_STATUS", "data":%s}`, json_params)
		sendBytes := []byte(sendString)
		fmt.Println("sendString=", sendString)
		//wait 5 sec to make sure clients are connected
		timer1 := time.NewTimer(time.Duration(5) * time.Second)
		<-timer1.C

		for _, c := range clients {
			fmt.Println("=> CONTACT_STATUS", c.conn.RemoteAddr())
			c.conn.WriteMessage(1, sendBytes)
		}

		//}
	}
}

func onAMIAorListComplete() {
	//Params= map[Aor:400 Contactstatus:Reachable Endpointname:400 Roundtripusec:54995 Uri:sip:365m2y45@71.222.1.109:64576;transport=ws]
	fmt.Println("++++++ START onAMIAorListComplete")

	//writeFileExtensions()

	//NOTE: We dont need o include data because this nis paged by the app -> sendString := fmt.Sprintf(`{"type":"AMI_AOR_LIST_COMPLETE", "data":%s}`,json_params)
	sendString := fmt.Sprintf(`{"type":"AMI_AOR_LIST_COMPLETE"}`)
	sendBytes := []byte(sendString)
	//fmt.Println("sendString=", sendString)

	//wait 5 sec to make sure clients are connected
	timer1 := time.NewTimer(time.Duration(5) * time.Second)
	<-timer1.C

	for _, c := range clients {
		fmt.Println("=> AMI_AOR_LIST_COMPLETE", c.conn.RemoteAddr())
		c.conn.WriteMessage(1, sendBytes)
	}
}

func onAMIAorList(params map[string]string) {
	/*fmt.Println("++++++++++++++++++++++++++++")
	fmt.Println("++++++++++++++++++++++++++++")

	fmt.Println("1404: START onAMIAorList params=",params)
	fmt.Println("1404: extensions=",extensions)
	*/
	objname, ok_objname := params["Objectname"]
	//fmt.Println("1404: Objectname=",objname,"ok_objname",ok_objname)

	if ok_objname {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
    	defer cancel()
		for i, exten := range extensions {
			if exten.Number == objname {
				fmt.Println("!!!! found exten ", objname)
				contacts, ok_contacts := params["Contacts"]
				if ok_contacts {
					if len(contacts) > 0 {

						extensions[i].Status = `Reachable`
						extensions[i].Contacts = contacts
						dbUpsertExtension(ctx,extensions[i])

					} else {

						extensions[i].Status = ``
						extensions[i].Contacts = ``
						dbUpsertExtension(ctx,extensions[i])
					}
				}

				break
			}

		}

	}
}

func onAMIAsteriskReload(params map[string]string){
	reloadXoftSwitch()

}

func onAMICoreShowChannel(params map[string]string) {
	/*
		-> coreshowchannels
		2023/06/27 20:25:29 amiCoreShowChannels 2897.2: rsActionID= 1687897529123997076
		2023/06/27 20:25:29 amiCoreShowChannels 2897.3: amiCoreShowChannels() body= &{1687897529123997076 Success map[Actionid:1687897529123997076 Eventlist:start Message:Channels will follow]}
		-> 2023/06/27 20:25:29 2172.1: EventType: {CoreShowChannel [] map[Accountcode: Actionid:1687897529123997076 Application:AppDial Applicationdata:(Outgoing Line) Bridgeid:fff2a768-016e-4500-af8d-2f175449b3b5 Calleridname:405 Calleridnum:405 Channel:PJSIP/405-000000d4 Channelstate:6 Channelstatedesc:Up Connectedlinename:100 XO Connectedlinenum:100 Context:from-internal Duration:00:01:13 Exten: Language:en Linkedid:1687897455.215 Priority:1 Uniqueid:1687897455.216]}
		2023/06/27 20:25:29 2172.2: ID= CoreShowChannel Privilege= [] Params= map[Accountcode: Actionid:1687897529123997076 Application:AppDial Applicationdata:(Outgoing Line) Bridgeid:fff2a768-016e-4500-af8d-2f175449b3b5 Calleridname:405 Calleridnum:405 Channel:PJSIP/405-000000d4 Channelstate:6 Channelstatedesc:Up Connectedlinename:100 XO Connectedlinenum:100 Context:from-internal Duration:00:01:13 Exten: Language:en Linkedid:1687897455.215 Priority:1 Uniqueid:1687897455.216]
		2023/06/27 20:25:29 2172.1: EventType: {CoreShowChannel [] map[Accountcode: Actionid:1687897529123997076 Application:Dial Applicationdata:PJSIP/405/sip:405@71.222.1.109:5060,,HhTtrb(func-apply-sipheaders^s^1) Bridgeid:fff2a768-016e-4500-af8d-2f175449b3b5 Calleridname:100 XO Calleridnum:100 Channel:PJSIP/100-000000d3 Channelstate:6 Channelstatedesc:Up Connectedlinename:405 Connectedlinenum:405 Context:macro-dial-one Duration:00:01:13 Exten:s Language:en Linkedid:1687897455.215 Priority:56 Uniqueid:1687897455.215]}
		2023/06/27 20:25:29 2172.2: ID= CoreShowChannel Privilege= [] Params= map[Accountcode: Actionid:1687897529123997076 Application:Dial Applicationdata:PJSIP/405/sip:405@71.222.1.109:5060,,HhTtrb(func-apply-sipheaders^s^1) Bridgeid:fff2a768-016e-4500-af8d-2f175449b3b5 Calleridname:100 XO Calleridnum:100 Channel:PJSIP/100-000000d3 Channelstate:6 Channelstatedesc:Up Connectedlinename:405 Connectedlinenum:405 Context:macro-dial-one Duration:00:01:13 Exten:s Language:en Linkedid:1687897455.215 Priority:56 Uniqueid:1687897455.215]
		2023/06/27 20:25:29 2172.1: EventType: {CoreShowChannelsComplete [] map[Actionid:1687897529123997076 Eventlist:Complete Listitems:2]}
		2023/06/27 20:25:29 2172.2: ID= CoreShowChannelsComplete Privilege= [] Params= map[Actionid:1687897529123997076 Eventlist:Complete Listitems:2]


		2023/06/27 20:25:29 2172.2: ID= CoreShowChannel Privilege= [] Params= map[
			Accountcode:
			Actionid:1687897529123997076
			Application:AppDial
			Applicationdata:(Outgoing Line)
			Bridgeid:fff2a768-016e-4500-af8d-2f175449b3b5
		Calleridname:405
		Calleridnum:405
		Channel:PJSIP/405-000000d4
		Channelstate:6 Channelstatedesc:Up
		Connectedlinename:100 XO
		Connectedlinenum:100
		Context:from-internal
		Duration:00:01:13 Exten:
		Language:en
		Linkedid:1687897455.215
		Priority:1
		Uniqueid:1687897455.216]


		2023/06/27 20:25:29 2172.2: ID= CoreShowChannel Privilege= [] Params= map[
			Accountcode: Actionid:1687897529123997076
			Application:Dial Applicationdata:PJSIP/405/sip:405@71.222.1.109:5060,,HhTtrb(func-apply-sipheaders^s^1) Bridgeid:fff2a768-016e-4500-af8d-2f175449b3b5
		Calleridname:100 XO
		Calleridnum:100
		Channel:PJSIP/100-000000d3
		Channelstate:6 Channelstatedesc:Up
		Connectedlinename:405
		Connectedlinenum:405
		Context:macro-dial-one
		:00:01:13 Exten:s Language:en Linkedid:1687897455.215 Priority:56 Uniqueid:1687897455.215]



	*/

	//Params= map[Aor:400 Contactstatus:Reachable Endpointname:400 Roundtripusec:54995 Uri:sip:365m2y45@71.222.1.109:64576;transport=ws]
	fmt.Println("++++++ START onAMICoreShowChannel(params=", params)
	json_params, err1 := json.Marshal(params)
	fmt.Println("json_params=", json_params)
	if err1 == nil {
		//ret_json := fmt.Sprintf(`{"total":%v,"page_num":%v,"items":%s}`,total,page_num,json_items)

		//jsonString, err2 := json.Marshal(ret_json)
		//fmt.Println(err)
		//if err2 == nil {
		//fmt.Println("-> CONTACT_STATUS", json_params)
		sendString := fmt.Sprintf(`{"type":"ON_AMI_CORESHOWCHANNEL", "data":%s}`, json_params)
		sendBytes := []byte(sendString)
		fmt.Println("sendString=", sendString)
		//wait 5 sec to make sure clients are connected
		//timer1 := time.NewTimer(time.Duration(5) * time.Second)
		//<-timer1.C

		for _, c := range clients {
			fmt.Println("=> ON_AMI_CORESHOWCHANNEL", c.conn.RemoteAddr())
			c.conn.WriteMessage(1, sendBytes)
		}

		//}
	}
}

func onAMICoreShowChannelsComplete() {
	/*
		2023/06/27 20:25:29 2172.1: EventType: {CoreShowChannelsComplete [] map[Actionid:1687897529123997076 Eventlist:Complete Listitems:2]}
		2023/06/27 20:25:29 2172.2: ID= CoreShowChannelsComplete Privilege= [] Params= map[Actionid:1687897529123997076 Eventlist:Complete Listitems:2]
	*/
	fmt.Println("++++++ START onAMICoreShowChannelsComplete")

	//writeFileExtensions()

	//NOTE: We dont need o include data because this nis paged by the app -> sendString := fmt.Sprintf(`{"type":"AMI_AOR_LIST_COMPLETE", "data":%s}`,json_params)
	sendString := fmt.Sprintf(`{"type":"ON_AMI_CORESHOWCHANNELSCOMPLETE"}`)
	sendBytes := []byte(sendString)
	//fmt.Println("sendString=", sendString)

	//wait 5 sec to make sure clients are connected
	//timer1 := time.NewTimer(time.Duration(5) * time.Second)
	//			<-timer1.C

	for _, c := range clients {
		fmt.Println("=> ON_AMI_CORESHOWCHANNELSCOMPLETE", c.conn.RemoteAddr())
		c.conn.WriteMessage(1, sendBytes)
	}
}



func prettyJSON(str_json string) {
	var b bytes.Buffer
	error := json.Indent(&b, []byte(str_json), "", "\t")
	if error != nil {
		fmt.Println("JSON parse error: ", error)

	} else {
		fmt.Println("prettyJSON:", string(b.String()))
	}
}
func writeFileRegistered() {
	return

	//...................................
	//Writing struct type to a JSON file
	//...................................
	//content, err := json.Marshal(user)
	fmt.Println("START writeFileRegistered() registered=", registered)

	content, err := json.Marshal(registered)
	if err != nil {
		fmt.Println("ERROR writeFileRegistered() registered=", registered)
		fmt.Println(err)
	}
	//fmt.Println("writeFileRegistered() content=",content)
	err = ioutil.WriteFile("registered.json", content, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func writeFileAOR(data Map) {

	content, err := json.Marshal(data)
	if err != nil {
		fmt.Println("ERROR writeFileAOR() data=", data)
		fmt.Println(err)
	}
	//fmt.Println("writeFileRegistered() content=",content)
	err = ioutil.WriteFile("aor.json", content, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

/*
func writeFileRegistered2(data []WSRegister){

	//...................................
	//Writing struct type to a JSON file
	//...................................
	//content, err := json.Marshal(user)
	fmt.Println("START writeFileRegistered() registered=",data)

	content, err := json.Marshal(registered)
	if err != nil {
		fmt.Println("ERROR writeFileRegistered() registered=",registered)
		fmt.Println(err)
	}
	fmt.Println("writeFileRegistered() content=",content)
	err = ioutil.WriteFile("registered.json", content, 0644)
	if err != nil {
		fmt.Println(err)
	}

}
*/
func readFileRegistered_dep() {
	//...................................
	//Reading into struct type from a JSON file
	//...................................

	content, err := ioutil.ReadFile("registered.json")
	if err != nil {
		fmt.Println(err)
	} else {

		err = json.Unmarshal(content, &registered)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("LOADED registered.json registered=", registered)

	}
}

func readFileUsers_dep() {
	//...................................
	//Reading into struct type from a JSON file
	//...................................

	content, err := ioutil.ReadFile("users.json")
	if err != nil {
		fmt.Println(err)
	} else {

		err = json.Unmarshal(content, &users)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("LOADED users.json users=", users)

	}
}
func writeFileUsers_dep() {

	//...................................
	//Writing struct type to a JSON file
	//...................................
	//content, err := json.Marshal(user)
	fmt.Println("START writeFileUsers() users=", users)

	content, err := json.Marshal(users)
	if err != nil {
		fmt.Println("ERROR writeFileUsers() users=", users)
		fmt.Println(err)
	}
	//fmt.Println("writeFileRegistered() content=",content)
	err = ioutil.WriteFile("users.json", content, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func readFileJoinRequests_dep() {
	//...................................
	//Reading into struct type from a JSON file
	//...................................

	content, err := ioutil.ReadFile("joinrequests.json")
	if err != nil {
		fmt.Println(err)
	} else {

		err = json.Unmarshal(content, &joinrequests)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("LOADED joinrequests.json joinrequests=", joinrequests)

	}
}
func writeFileJoinRequests_dep() {

	//...................................
	//Writing struct type to a JSON file
	//...................................
	//content, err := json.Marshal(user)
	fmt.Println("START writeFileJoinRequests() joinrequests=", joinrequests)

	content, err := json.Marshal(joinrequests)
	if err != nil {
		fmt.Println("ERROR writeFileJoinRequests() joinrequests=", joinrequests)
		fmt.Println(err)
	}
	//fmt.Println("writeFileRegistered() content=",content)
	err = ioutil.WriteFile("joinrequests.json", content, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func readFileUserExtensions_dep() {
	//...................................
	//Reading into struct type from a JSON file
	//...................................

	content, err := ioutil.ReadFile("userextensions.json")
	if err != nil {
		fmt.Println(err)
	} else {
		//save_file := false
		err = json.Unmarshal(content, &userextensions)
		if err != nil {
			fmt.Println(err)
		} else {
			//sync passwords with asterisk
			/*if len(userextensions) > 0{
				for i,x := range userextensions{
					auth_password, ok_auth_password := extensions_auth[x.Extension]
					if ok_auth_password && x.Password != auth_password{
						userextensions[i].Password = auth_password
						save_file = true
					}
				}
				if save_file{
					writeFileUserExtensions()
				}
			}
			*/

		}

		fmt.Println("LOADED userextensions.json userextensions=", userextensions)

	}

	// sync passwords with asterisk

}
func writeFileUserExtensions_dep() {

	//...................................
	//Writing struct type to a JSON file
	//...................................
	//content, err := json.Marshal(user)

	content, err := json.Marshal(userextensions)
	if err != nil {
		fmt.Println("ERROR writeFileUserExtensions() userextensions=", userextensions)
		fmt.Println(err)
	} else {
		fmt.Println("1820:************")
		fmt.Println("1820:************")
		fmt.Println("1820:************")
		fmt.Println("START writeFileUserExtensions() userextensions=", userextensions)
		fmt.Println("1820:************")
		fmt.Println("1820:************")
		fmt.Println("1820:************")
		err = ioutil.WriteFile("userextensions.json", content, 0644)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func writeFileEndpoints() {

	//...................................
	//Writing struct type to a JSON file
	//...................................
	//content, err := json.Marshal(user)
	fmt.Println("START writeFileEndpoints() endpoints=", endpoints)

	content, err := json.Marshal(endpoints)
	if err != nil {
		fmt.Println("ERROR writeFileEndpoints() endpoints=", endpoints)
		fmt.Println(err)
	}
	//fmt.Println("writeFileRegistered() content=",content)
	err = ioutil.WriteFile("endpoints.json", content, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
func writeFileExtensionsAuth() {

	//...................................
	//Writing struct type to a JSON file
	//...................................
	//content, err := json.Marshal(user)
	//fmt.Println("START writeFileExtensionsAuth() extensions_auth=",extensions_auth)

	content, err := json.Marshal(extensions_auth)
	if err != nil {
		fmt.Println("ERROR writeFileExtensionsAuth() extensions_auth=", extensions_auth)
		fmt.Println(err)
	}
	//fmt.Println("writeFileRegistered() content=",content)
	err = ioutil.WriteFile("extensions_auth.json", content, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

/*
f, err := os.OpenFile("text.log",
	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
	fmt.Println(err)
}
defer f.Close()
if _, err := f.WriteString("text to append\n"); err != nil {
	fmt.Println(err)
}
*/

func writeFileAMIContactStatus() {

	//...................................
	//Writing struct type to a JSON file
	//...................................
	//content, err := json.Marshal(user)
	//fmt.Println("START writeFileRegistered() registered=",registered)

	content, err := json.Marshal(contactstatusmaptable)
	if err != nil {
		fmt.Println("ERROR writeFileAMIContactStatus() contactstatusmaptable=", contactstatusmaptable)
		fmt.Println(err)
	}
	//fmt.Println("writeFileRegistered() content=",content)
	err = ioutil.WriteFile("contactstatus.json", content, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func readFileAMIContactStatus() {
	//...................................
	//Reading into struct type from a JSON file
	//...................................

	content, err := ioutil.ReadFile("contactstatus.json")
	if err != nil {
		fmt.Println(err)
	} else {

		err = json.Unmarshal(content, &contactstatusmaptable)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("LOADED contactstatus.json contactstatusmaptable=", contactstatusmaptable)

	}
}



func GenerateRandomPassword(length int) (string, error) {
	if length < 8 {
		length = 8
	}

	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+"
	var password []byte

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		password = append(password, chars[n.Int64()])
	}

	return string(password), nil
}

func GenerateRandomUsername(prefix string, length int) (string, error) {
	if length < 4 {
		length = 4 // reasonable minimum
	}

	chars := "abcdefghijklmnopqrstuvwxyz0123456789"
	var sb strings.Builder
	sb.WriteString(prefix)
	if prefix != "" && !strings.HasSuffix(prefix, "_") {
		sb.WriteString("_")
	}

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		sb.WriteByte(chars[n.Int64()])
	}

	return sb.String(), nil
}

func writeFileConfig() {

	//...................................
	//Writing struct type to a JSON file
	//...................................
	//content, err := json.Marshal(user)
	//fmt.Println("START writeFileRegistered() registered=",registered)

	content, err := json.Marshal(config)
	if err != nil {
		fmt.Println("ERROR writeFileConfig() config=", config)
		fmt.Println(err)
	}
	//fmt.Println("writeFileRegistered() content=",content)
	err = ioutil.WriteFile("config.json", content, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func readFileConfig() {
	//...................................
	//Reading into struct type from a JSON file
	//...................................

	has_changed := false

	bytes_content, err := ioutil.ReadFile("config.json")
	if err != nil {
		return
	}

	str_content := string(bytes_content)
	str_content = strings.ReplaceAll(str_content, "\r\n", "")
	str_content = strings.ReplaceAll(str_content, "\n", "")

	content := []byte(str_content)

	if err != nil {
		fmt.Println(err)

	} else {

		err = json.Unmarshal(content, &config)
		if err != nil {
			fmt.Println(err)

		}
	}

	fmt.Println(`1778: config=`, config)

	//public_hostname, ok_public_hostname := config[`public_hostname`]
	admin_username, ok_admin_username := config[`admin_username`]
	//public_ip, ok_public_ip := config[`public_ip`]
	admin_password, ok_admin_password := config[`admin_password`]

	if ok_admin_username {
		config[`admin_username`] = admin_username

	} else {
		username, err := GenerateRandomUsername("admin", 6)
		if err != nil {
			panic(err)
		}
		fmt.Println("Generated username:", username)
		config[`admin_username`] = username
		has_changed = true
	}

	if ok_admin_password {
		config[`admin_password`] = admin_password

	} else {
		pass, err := GenerateRandomPassword(12)
		if err != nil {
			panic(err)
		}
		println("Generated password:", pass)
		config[`admin_password`] = pass
		has_changed = true
	}

	
	
	
	asterisk_conf_path, ok_asterisk_conf_path := config[`asterisk_conf_path`]
	fmt.Println(`1780: asterisk_conf_path=`, asterisk_conf_path, `ok_asterisk_conf_path=`, ok_asterisk_conf_path)

	if ok_asterisk_conf_path == false {
		//fmt.Println(`1784:`)

		config[`asterisk_conf_path`] = `/etc/asterisk`
		asterisk_conf_path = config[`asterisk_conf_path`]
		has_changed = true
	}
	//fmt.Println(`1789: config=`,config)
	/* ex:
	if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
	// path/to/whatever does not exist
	}

	if _, err := os.Stat("/path/to/whatever"); !os.IsNotExist(err) {
		// path/to/whatever exists
	}

	*/

	if _, err := os.Stat(asterisk_conf_path); os.IsNotExist(err) {
		log.Fatal(asterisk_conf_path, " does not exist!")
	}

	public_hostname, ok_public_hostname := config[`public_hostname`]
	public_ip, ok_public_ip := config[`public_ip`]

	if !ok_public_hostname || len(public_hostname) == 0 || !ok_public_ip || len(public_ip) == 0 {
		public_ip, public_hostname = get_public_ip_hostname()
		if len(public_ip) > 0 {
			config[`public_hostname`] = public_hostname
			config[`public_ip`] = public_ip
			has_changed = true
		} else {

			log.Fatal("public_hostname required!")

		}
	}

	cert_dir, ok_cert_dir := config[`cert_dir`]

	if !ok_cert_dir || len(cert_dir) == 0 {
		/*
			if len(public_ip) > 0{
				//config[`cert_dir`] = fmt.Sprintf(`/etc/asterisk/keys/%s`,public_hostname)
				// /etc/asterisk/keys/integration/certificate.pem
				// /etc/asterisk/keys/integration/webserver.key
				config[`cert_dir`] = `/etc/asterisk/keys/integration`

				has_changed = true
			} else{

				log.Fatal("cert_dir required!")

			}
		*/
		config[`cert_dir`] = `/etc/asterisk/keys/integration`
		has_changed = true
	}

	ami_username, ami_password := get_admin_pass_manager_conf()
	if len(ami_username) > 0 && len(ami_password) > 0 {
		config[`ami_username`] = ami_username
		config[`ami_password`] = ami_password

	} else {
		log.Fatal("Invalid AMI login!")
	}

	auto_join, ok_auto_join := config[`auto_join`]
	if ok_auto_join {
		config[`auto_join`] = auto_join

	} else {
		config[`auto_join`] = `no`
		has_changed = true
	}

	logo_filename, ok_logo_filename := config[`logo_filename`]
	if ok_logo_filename {
		config[`ok_logo_filename`] = logo_filename

	} else {
		config[`logo_filename`] = `logo.png`
		has_changed = true
	}

	ssid, ok_ssid := config[`wifi_ssid`]
	if ok_ssid {
		config[`wifi_ssid`] = ssid

	} else {
		config[`wifi_ssid`] = ``
		has_changed = true
	}

	bssid, ok_bssid := config[`wifi_bssid`]
	if ok_bssid {
		config[`wifi_bssid`] = bssid

	} else {
		config[`wifi_bssid`] = ``
		has_changed = true
	}

	if has_changed {
		writeFileConfig()
	}

	fmt.Println(`+++++++ 1773: readFileConfig() config=`, config)
}

func readFileRoles() {
	//...................................
	//Reading into struct type from a JSON file
	//...................................
	roles = map[string]string{`admin`: `Admin`, `employee`: `Employee`, `customer`: `Customer`, `default`: `default`}

	content, err := ioutil.ReadFile("roles.json")

	if err != nil {
		fmt.Println(err)
		writeFileRoles()

	} else {

		err = json.Unmarshal(content, &roles)
		if err != nil {
			fmt.Println(err)

		}
	}

}

func writeFileRoles() {

	//...................................
	//Writing struct type to a JSON file
	//...................................
	//content, err := json.Marshal(user)
	//fmt.Println("START writeFileRegistered() registered=",registered)

	content, err := json.Marshal(roles)
	if err != nil {
		fmt.Println("ERROR writeFileRoles() roles=", roles)
		fmt.Println(err)
	}
	//fmt.Println("writeFileRegistered() content=",content)
	err = ioutil.WriteFile("roles.json", content, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func readFileExtensionRange() {
	func_name := `readFileExtensionRange`
	//...................................
	//Reading into struct type from a JSON file
	//...................................
	// start,end
	//has_changed := false
	ra := map[string]string{
		`admin`:    `100,199`,
		`employee`: `200,299`,
		`customer`: `300,309`,
		`default`:  `310,319`}

	content, err := ioutil.ReadFile("extension_range.json")

	if err != nil {
		fmt.Println(func_name, `3085: err=`, err)

		extension_range = ra
		writeFileExtenionRange()

	} else {

		err = json.Unmarshal(content, &extension_range)
		if err != nil {
			fmt.Println(func_name, `3094: err=`, err)

		} else {
			fmt.Println(func_name, `3097: OK`)
			writeFileExtenionRange()
		}
	}

}

func writeFileExtenionRange() {

	//...................................
	//Writing struct type to a JSON file
	//...................................
	//content, err := json.Marshal(user)
	//fmt.Println("START writeFileRegistered() registered=",registered)

	content, err := json.Marshal(extension_range)
	if err != nil {
		fmt.Println("ERROR writeFileRoles() extension_range=", extension_range)
		fmt.Println(err)
	}
	//fmt.Println("writeFileRegistered() content=",content)
	err = ioutil.WriteFile("extension_range.json", content, 0644)
	if err != nil {
		fmt.Println(err)
	}
}



func testAOrContacts(a *agi.AGI, exten string) {
	fmt.Println("testAOrContacts...")
	aor_contacts, err_aor_contacts := a.Get(fmt.Sprintf(`PJSIP_DIAL_CONTACTS(%s)`, exten))
	fmt.Println("731: ++++++++++++++")

	if err_aor_contacts == nil {
		fmt.Println("current aor_contacts=", aor_contacts)
	
		
	} else {
		fmt.Println("current err_aor_contacts=", err_aor_contacts)
		
	}

	timer1 := time.NewTimer(15 * time.Second)
	<-timer1.C
	aor_contacts, err_aor_contacts = a.Get(fmt.Sprintf(`PJSIP_DIAL_CONTACTS(%s)`, exten))
	fmt.Println("746: ++++++++++++++")

	if err_aor_contacts == nil {
		fmt.Println("final aor_contacts=", aor_contacts)
		
		
	} else {
		fmt.Println("final err_aor_contacts=", err_aor_contacts)
		
	}

}



func pushNotify(notif_to, notif_title, notif_body string) {
	fmt.Println("8429.0: START pushNotify(notif_to=", notif_to, "notif_title=", notif_title, "notif_body=", notif_body)

	
	postBody := []byte(fmt.Sprintf(`{
		"apikey": "%s",
		"action":"PUSH_NOTIFY",

		"notif_to":"%s",
		"notif_title":"%s",
		"notif_body":"%s"}`,
		config["apikey"],
		notif_to,
		notif_title,
		notif_body))
	fmt.Println("8429.1: wakeUpDevice body=", string(postBody))

	resp, err := http.Post(api_xoftswitch, "application/json", bytes.NewBuffer(postBody))
	//Handle Error
	if err != nil {
		//log.Fatalf("511: An Error Occured %v", err)
		fmt.Println("8429.2: An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Fatalln(err)
		fmt.Println(err)
	}
	sb := string(body)
	fmt.Println("8429.3: pushNotify() resp.Body=", sb)
	var sb_map Map
	json.Unmarshal(body, &sb_map)
	if sb_map["statusCode"] == 200 && sb_map["uuid"] != nil {
		//call_uuid := sb_map["uuid"]

	}

}



func findOneUserExtensions(username, extension string) *UserExtension {

	//ret := []UserExtension{}
	for index, exten := range userextensions {
		if exten.Extension == extension && exten.Username == username {
			//ret = append(ret,exten)
			return &userextensions[index]
		}
	}
	return nil
}
func findOneUserExtensionsInGroup(username, extension, groupid string) *UserExtension {

	//ret := []UserExtension{}
	for index, exten := range userextensions {
		if exten.Extension == extension && exten.Username == username && exten.Groupid == groupid {
			//ret = append(ret,exten)
			return &userextensions[index]
		}
	}
	return nil
}

func findUserExtensionsByExtension(extension string) (int, *UserExtension) {
	func_name := `findUserExtensionsByExtension()`
	fmt.Println(func_name, `4100: START extension=`, extension)

	//ret := []UserExtension{}
	for i, x := range userextensions {
		if x.Extension == extension {
			//ret = append(ret,exten)
			fmt.Println(func_name, `4106: FOUND`, extension)
			return i, &x
		}
	}
	fmt.Println(func_name, `4109: NOT FOUND`, extension)
	return -1, nil
}

func findUserExtensions(username, extension string) (int, *UserExtension, error) {
	func_name := `findUserExtensionsByExtension()`
	fmt.Println(func_name, `4100: START extension=`, extension, `username=`, username)

	//ues,err := dbFindUserExtensions(username)
	err := dbLoadUserExtensions()
	//ret := []UserExtension{}
	if err != nil {
		return -1, nil, err
	}
	for i, x := range userextensions {
		if x.Username == username && x.Extension == extension {
			//ret = append(ret,exten)
			fmt.Println(func_name, `FOUND userextensions`, username, extension)
			return i, &x, nil
		}
	}
	fmt.Println(func_name, `4109: NOT FOUND userextensions`, username, extension)
	return -1, nil, nil
}

// this one is use for messagin
func findOneUserExtensionsNotUsername(username, extension string) *UserExtension {

	//ret := []UserExtension{}
	for index, exten := range userextensions {
		if exten.Extension == extension && exten.Username != username {
			//ret = append(ret,exten)
			return &userextensions[index]
		}
	}
	return nil
}

func findAllUserExtensions(username string) ([]UserExtension, error) {
	//fmt.Println(`findAllUserExtensions(`,username,`)`)
	results, err := dbFindUserExtensions(username)

	if err != nil {
		return nil, err
	}

	ret := []UserExtension{}
	/*
		for _,exten := range userextensions{
			if exten.Username == username{
				ret = append(ret,exten)
			}
		}
	*/
	for _, exten := range results {
		ret = append(ret, UserExtension{Username: username, Extension: exten})

	}

	fmt.Println(`findAllUserExtensions(`, username, `) ret=`, ret)

	return ret, nil

}

func findAllUserExtensionsInGroup(username, groupid string) ([]UserExtension, error) {
	//fmt.Println(`findAllUserExtensions(`,username,`)`)
	results, err := dbFindUserExtensionsInGroup(username, groupid)

	if err != nil {
		return nil, err
	}

	ret := []UserExtension{}
	/*
		for _,exten := range userextensions{
			if exten.Username == username{
				ret = append(ret,exten)
			}
		}
	*/
	for _, exten := range results {
		ret = append(ret, UserExtension{Username: username, Extension: exten, Groupid: groupid})

	}

	fmt.Println(`findAllUserExtensions(`, username, `) ret=`, ret)

	return ret, nil

}

func removeAllUserExtensions(username string) bool {

	if len(userextensions) == 0 {
		return true
	}
	found_index := -1
	for {
		for index, ue := range userextensions {

			//fmt.Println("96: c.device_token",r.device_token)
			if ue.Username == username {
				found_index = index
				break

			}
		}
		if found_index >= 0 {
			//fmt.Println("572: findRegistrations xten=",exten,"device_token=", device_token, "FOUND")
			userextensions = append(userextensions[:found_index], userextensions[found_index+1:]...)
			found_index = -1
		} else {
			break
		}

	}

	ues, err := findAllUserExtensions(username)
	if err != nil {
		return false
	}

	if len(ues) > 0 {
		return false
	}

	dbLoadUserExtensions()

	return true
	
}


func unregisterPJSIPExtensionn(ext string) {
	
	
	func_name := `unregisterPJSIPExtensionn`

	rs, rsActionID, rsErr := glo_ami.Action(gami.Params{ //ami.Action(gami.Params{
		"Action":   "PJSIPUnregister",
		"Endpoint": ext,
	})

	if rsErr != nil {
		fmt.Println(func_name, "8397.1: rsErr=", rsErr)

	} else {

		fmt.Println(func_name, "8397.2: rsActionID=", rsActionID)
		//fmt.Println(<-rsGetConfig)
		body := fmt.Sprintf("%v", <-rs)

		fmt.Println(func_name, `8397.3: unregisterPJSIPExtensionn() body=`, body)
		/*if !strings.Contains(body, `Error`) {
			return true
		}
		*/

	}

}

func removeAllUserExtensionsInGroup(groupid string) bool {

	fmt.Println(`8339.1: START removeAllUserExtensionsInGroup`)

	if len(userextensions) == 0 {
		return true
	}

	for _, ue := range userextensions {

		//fmt.Println("96: c.device_token",r.device_token)
		if ue.Groupid == groupid {

			//unregisterPJSIPExtension(a, ext)
			unregisterPJSIPExtensionn(ue.Extension)
		}
	}

	dbDeleteUserExtensionsInGroup(groupid)
	
	
	dbLoadUserExtensions()

	return true
	
}

func removeAllUserExtensionsNotInGroup(exten_num, groupid string) bool {

	fmt.Println(`8339.1: START removeAllUserExtensionsNotInGroup`)

	if len(userextensions) == 0 {
		return true
	}

	
	

	dbDeleteUserExtensionsNotInGroup(exten_num, groupid)
	
	
	dbLoadUserExtensions()

	return true
	
}

func findAllUsersWithExtension(pexten string) []UserExtension {

	ret := []UserExtension{}
	for _, exten := range userextensions {
		if exten.Extension == pexten {
			ret = append(ret, exten)
		}
	}
	return ret
}

func findUserByUsername(username string) (int, *User) {

	for i, user := range users {
		if user.Username == username {
			return i, &user
		}
	}
	return -1, nil
}

// find admins to receive notification for join requests, etc...
func findUserAdmins() []User {

	var ret []User

	for _, user := range users {
		if user.Roleid == `admin` {
			//return &user
			ret = append(ret, user)
		}
	}
	return ret
}

func findUserByEmail(email string) (int, *User) {

	for i, user := range users {
		if user.Email == email {
			return i, &user
		}
	}
	return -1, nil
}

func findUserByUsernameOrEmail(username string) (int, *User) {

	for index, user := range users {
		if strings.Contains(username, `@`) {
			if user.Email == username {
				return index, &user
			}
		} else {
			if user.Username == username {
				return index, &user
			}

		}

	}
	return -1, nil
}

func findClientByExtenForAGI(exten string) *WSRegister {
	//index := -1

	fmt.Println("Find exten", exten)
	fmt.Println("clients", clients)

	for _, r := range registered {
		//fmt.Println("30: idx=",idx,"c= ",c.conn.RemoteAddr())
		//if c == conn{
		//if c.conn.RemoteAddr() == conn.RemoteAddr(){
		fmt.Println("96: c.exten", r.Exten)
		if r.Exten == exten {
			fmt.Println("93: client", exten, "FOUND")
			//index = idx
			//break
			return &r
		}
	}
	fmt.Println("Find exten", exten, "NOT found.")
	return nil
}

func findClientsByExten(exten string) []WSRegister {
	//index := -1

	//fmt.Println("Find exten",exten)
	//fmt.Println("clients",clients)
	ret := []WSRegister{}

	for _, r := range registered {
		//fmt.Println("30: idx=",idx,"c= ",c.conn.RemoteAddr())
		//if c == conn{
		//if c.conn.RemoteAddr() == conn.RemoteAddr(){
		fmt.Println("96: c.exten", r.Exten)
		if r.Exten == exten {
			fmt.Println("93: client", exten, "FOUND")
			//index = idx
			//break
			//return &r
			ret = append(ret, r)
		}
	}
	fmt.Println("Find exten", exten, "NOT found.")
	return nil
}

func findClientByConn(conn *websocket.Conn) (int, *WSClient) {

	for index, c := range clients {
		//fmt.Println("30: idx=",idx,"c= ",c.conn.RemoteAddr())
		//if c == conn{
		if c.conn.RemoteAddr() == conn.RemoteAddr() {
			fmt.Println("75: FOUND")
			//index = idx
			//break
			return index, &c
		}
	}
	return -1, nil
}

func findClientByDeviceToken(device_token string) (int, *WSClient) {

	for index, c := range clients {
		//fmt.Println("30: idx=",idx,"c= ",c.conn.RemoteAddr())
		//if c == conn{
		if c.device_token == device_token {
			fmt.Println("1338: OK findClientByDeviceToken", device_token)
			//index = idx
			//break
			return index, &c
		}
	}
	return -1, nil
}

// func findContactsByConn(conn *websocket.Conn) (int, *WSRegister){
func findContactsByConn(conn *websocket.Conn) []WSRegister {

	contacts := []WSRegister{}

	for _, r := range registered {
		//fmt.Println("30: idx=",idx,"c= ",c.conn.RemoteAddr())
		//if c == conn{
		if r.Remote_Addr == fmt.Sprintf("%s", conn.RemoteAddr()) {
			fmt.Println("544: FOUND")
			//index = idx
			//break
			//return index,&r
			contacts = append(contacts, r)
		}

	}
	return contacts
}

// func findContacts(exten string, caller string) []WSContact{
func findContacts(exten string) []WSContact {
	//index := -1

	//fmt.Println("Find exten",exten)
	//fmt.Println("clients",clients)
	contacts := []WSContact{}

	for _, r := range registered {

		//fmt.Println("96: r.exten",r.exten)
		if r.Exten == exten {
			//fmt.Println("669: registered FOUND", exten, "caller",caller)
			contact := WSContact{contact: &r, client: nil}
			fmt.Println("669: contact registered FOUND", exten)

			for _, c := range clients {
				//fmt.Println("30: idx=",idx,"c= ",c.conn.RemoteAddr())
				//if c == conn{
				//if c.conn.RemoteAddr() == conn.RemoteAddr(){
				//fmt.Println("96: c.exten",c.exten)
				remote_addr := fmt.Sprintf("%s", c.conn.RemoteAddr())
				if remote_addr == r.Remote_Addr {
					//fmt.Println("678: client FOUND", exten,  "caller",caller)
					contact.client = &c
					fmt.Println("678: contact client FOUND", exten)
					//index = idx
					//break

				} else {
					//fmt.Println("683: client NOT FOUND", exten, "caller",caller)
					fmt.Println("683: contact client NOT FOUND", exten)
					//contacts = append(contacts,WSContact{contact: &r, client: nil})
				}
			}
			contacts = append(contacts, contact)

		}
	}

	//fmt.Println("694: contacts=", contacts,"exten=", exten, "caller",caller)
	fmt.Println("694: FINAL contacts=", contacts, "exten=", exten)
	return contacts
}

func findContactsByDeviceToken(device_token string) []WSContact {
	//index := -1

	//fmt.Println("Find exten",exten)
	//fmt.Println("clients",clients)
	contacts := []WSContact{}

	for _, r := range registered {

		//fmt.Println("96: r.exten",r.exten)
		if r.Device_Token == device_token {
			//fmt.Println("669: registered FOUND", exten, "caller",caller)
			contact := WSContact{contact: &r, client: nil}
			fmt.Println("669: registered device_token FOUND", device_token)

			for _, c := range clients {
				//fmt.Println("30: idx=",idx,"c= ",c.conn.RemoteAddr())
				//if c == conn{
				//if c.conn.RemoteAddr() == conn.RemoteAddr(){
				//fmt.Println("96: c.exten",c.exten)
				remote_addr := fmt.Sprintf("%s", c.conn.RemoteAddr())
				if remote_addr == r.Remote_Addr {
					//fmt.Println("678: client FOUND", exten,  "caller",caller)
					contact.client = &c
					fmt.Println("678: client device_token FOUND", device_token)
					//index = idx
					//break

				} else {
					//fmt.Println("683: client NOT FOUND", exten, "caller",caller)
					fmt.Println("683: client device_token NOT FOUND", device_token)
					//contacts = append(contacts,WSContact{contact: &r, client: nil})
				}
			}
			contacts = append(contacts, contact)

		}
	}

	//fmt.Println("694: contacts=", contacts,"exten=", exten, "caller",caller)
	fmt.Println("694: FINAL contacts=", contacts, "exten=", device_token)
	return contacts
}

func findAllRegisteredWithDeviceToken(device_token string) []WSRegister {
	//index := -1

	//fmt.Println("Find exten",exten)
	//fmt.Println("clients",clients)
	results := []WSRegister{}

	for _, r := range registered {

		//fmt.Println("96: r.exten",r.exten)
		if r.Device_Token == device_token {
			//fmt.Println("669: registered FOUND", exten, "caller",caller)
			//contact := WSContact{contact: &r, client: nil}
			fmt.Println("669: registered device_token FOUND", device_token)
			results = append(results, r)

		}
	}

	//fmt.Println("694: contacts=", contacts,"exten=", exten, "caller",caller)
	fmt.Println("p54: FINAL results=", results, "device_token=", device_token)
	return results
}

func findContact(exten string, device_token string) *WSContact {
	//index := -1

	//fmt.Println("Find exten",exten)
	//fmt.Println("clients",clients)

	var contact *WSContact = nil

	for _, r := range registered {

		//fmt.Println("96: r.exten",r.exten)
		if r.Exten == exten {
			fmt.Println("711: registered", exten, "FOUND")

			for _, c := range clients {
				//fmt.Println("30: idx=",idx,"c= ",c.conn.RemoteAddr())
				//if c == conn{
				//if c.conn.RemoteAddr() == conn.RemoteAddr(){
				//fmt.Println("96: c.exten",c.exten)
				remote_addr := fmt.Sprintf("%s", c.conn.RemoteAddr())
				if remote_addr == r.Remote_Addr {
					fmt.Println("720: client", exten, "FOUND")
					//index = idx
					//break
					contact = &WSContact{contact: &r, client: &c}
				} else {
					fmt.Println("720: client", exten, "NOT FOUND")
					contact = &WSContact{contact: &r, client: nil}
				}
			}

		}
	}

	//fmt.Println("Find exten",exten, "NOT found.")
	fmt.Println("737: FINAL contact=", contact, "exten=", exten)
	return contact
}

func findRegistered(exten string, device_token string) (int, *WSRegister) {
	//index := -1

	//fmt.Println("findRegistrations",device_token)
	//fmt.Println("registrations",registrations)

	for i, r := range registered {

		//fmt.Println("96: c.device_token",r.device_token)
		if r.Device_Token == device_token && r.Exten == exten {
			fmt.Println("572: findRegistrations xten=", exten, "device_token=", device_token, "FOUND")
			//index = idx
			//break
			return i, &r
		}
	}

	fmt.Println("9100: findRegistrations xten=", exten, "device_token=", device_token, "NOT FOUND")
	return -1, nil
}
func removeRegisteredWithTokenNotUsername(device_token, username string) {
	//index := -1

	//fmt.Println("findRegistrations",device_token)
	//fmt.Println("registrations",registrations)
	//ret := []WSRegister{}
	found_index := -1
	for {
		for index, r := range registered {

			//fmt.Println("96: c.device_token",r.device_token)
			if r.Device_Token == device_token && r.Username != username {
				found_index = index
				break

			}
		}
		if found_index >= 0 {
			//fmt.Println("572: findRegistrations xten=",exten,"device_token=", device_token, "FOUND")
			registered = append(registered[:found_index], registered[found_index+1:]...)
			found_index = -1
		} else {
			break
		}

	}

}

func findManyRegistered(exten string) []WSRegister {
	//index := -1

	//fmt.Println("findRegistrations",device_token)
	//fmt.Println("registrations",registrations)
	ret := []WSRegister{}

	for _, r := range registered {

		//fmt.Println("96: c.device_token",r.device_token)
		if r.Exten == exten {
			//fmt.Println("572: findRegistrations xten=",exten,"device_token=", device_token, "FOUND")
			//index = idx
			//break
			ret = append(ret, r)
		}
	}

	
	return ret
}

func findManyRegisteredIOS(exten string) []WSRegister {
	//index := -1

	//fmt.Println("findRegistrations",device_token)
	//fmt.Println("registrations",registrations)
	ret := []WSRegister{}

	for _, r := range registered {

		//fmt.Println("96: c.device_token",r.device_token)
		if r.Exten == exten && r.Device_Type == `ios` {
			//fmt.Println("572: findRegistrations xten=",exten,"device_token=", device_token, "FOUND")
			//index = idx
			//break
			isdup := false
			for _, r_exist := range ret {
				if r_exist.Device_Token == r.Device_Token {
					isdup = true
					break
				}
			}
			if !isdup {
				ret = append(ret, r)
			}

		}
	}

	
	return ret
}
func findManyRegisteredMacOS(exten string) []WSRegister {
	//index := -1

	//fmt.Println("findRegistrations",device_token)
	//fmt.Println("registrations",registrations)
	ret := []WSRegister{}

	for _, r := range registered {

		//fmt.Println("96: c.device_token",r.device_token)
		if r.Exten == exten && r.Device_Type == `macos` {
			//fmt.Println("572: findRegistrations xten=",exten,"device_token=", device_token, "FOUND")
			//index = idx
			//break
			ret = append(ret, r)
		}
	}

	
	return ret
}

func findManyRegisteredWithStatus(exten string, status int) []WSRegister {
	//index := -1

	//fmt.Println("findRegistrations",device_token)
	//fmt.Println("registrations",registrations)
	ret := []WSRegister{}

	for _, r := range registered {

		//fmt.Println("96: c.device_token",r.device_token)
		if r.Exten == exten && r.Register_Status == status {
			//fmt.Println("572: findRegistrations xten=",exten,"device_token=", device_token, "FOUND")
			//index = idx
			//break
			ret = append(ret, r)
		}
	}

	
	return ret
}

func findIncomingCallsByUuid(uid string) []WSIncomingCall {
	//index := -1

	//fmt.Println("findRegistrations",device_token)
	//fmt.Println("registrations",registrations)
	ret := []WSIncomingCall{}

	for _, r := range incomingcalls {

		//fmt.Println("96: c.device_token",r.device_token)
		if r.Uuid == uid {
			//fmt.Println("572: findRegistrations xten=",exten,"device_token=", device_token, "FOUND")
			//index = idx
			//break
			ret = append(ret, r)
		}
	}

	
	return ret
}
func findIncomingCallsById(id string) *WSIncomingCall {
	//index := -1

	//fmt.Println("findRegistrations",device_token)
	//fmt.Println("registrations",registrations)
	//var ret *WSIncomingCall

	for _, r := range incomingcalls {

		//fmt.Println("96: c.device_token",r.device_token)
		if r.Id == id {
			//fmt.Println("572: findRegistrations xten=",exten,"device_token=", device_token, "FOUND")
			//index = idx
			//break
			//ret = append(ret,r)
			return &r
		}
	}

	
	return nil
}

func findAllActiveIncomingCalls() []WSIncomingCall {
	//index := -1

	//fmt.Println("findRegistrations",device_token)
	//fmt.Println("registrations",registrations)
	ret := []WSIncomingCall{}

	for _, r := range incomingcalls {

		//fmt.Println("96: c.device_token",r.device_token)
		if r.CKAnswerState == 2 && r.CKHangupState == 0 {
			//fmt.Println("572: findRegistrations xten=",exten,"device_token=", device_token, "FOUND")
			//index = idx
			//break
			ret = append(ret, r)
		}
	}

	
	return ret
}

func findIncomingCallsByCallerDeviceToken(device_token string) []WSIncomingCall {
	//index := -1

	//fmt.Println("findRegistrations",device_token)
	//fmt.Println("registrations",registrations)
	ret := []WSIncomingCall{}

	for _, r := range incomingcalls {

		//fmt.Println("96: c.device_token",r.device_token)
		if r.Device_Token == device_token {
			//fmt.Println("572: findRegistrations xten=",exten,"device_token=", device_token, "FOUND")
			//index = idx
			//break
			ret = append(ret, r)
		}
	}
	return ret

	
}

func findExtension(exten string) (int, *Extension) {
	for i, x := range extensions {
		if x.Number == exten {
			return i, &x

			break
		}

	}
	return -1, nil
}

func pushNotifiyAdmins(title, body string) {
	//send notification to admins
	func_name := `pushNotifiyAdmins()`
	admins := findUserAdmins()
	fmt.Println(func_name, `5000: pushNotifiyAdmins() admins=`, admins)
	if admins != nil && len(admins) > 0 {
		to := ``

		for _, u := range admins {
			to = fmt.Sprintf(`%s%s,`, to, u.Username)

		}

		to = strings.TrimRight(to, `,`)
		fmt.Println(func_name, `5011: pushNotifiyAdmins() to=`, to)
		if len(to) > 0 {
			httpPushNotify(to, title, body)
		} else {
			fmt.Println(func_name, `5015: pushNotifiyAdmins() to is invalid or empty!`)
		}

	}

}



func sendEmail(to, from, title, body string) error {
	funcName := "sendEmail()"

	smtpHost := config["SMTPHost"]
	smtpPort := config["SMTPPort"]
	smtpUser := config["SMTPUsername"]
	smtpPass := config["SMTPPassword"]
	adminEmail := config["AdminEmail"]

	// Validate configuration
	if smtpHost == "" || smtpPort == "" || smtpUser == "" || smtpPass == "" || adminEmail == "" {
		return fmt.Errorf("%s: missing SMTP configuration fields", funcName)
	}

	// Convert port to int
	portNum, err := strconv.Atoi(smtpPort)
	if err != nil {
		return fmt.Errorf("%s: invalid SMTP port: %v", funcName, err)
	}

	// Email message
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + title + "\n\n" +
		body

	// SMTPS (port 465)
	if portNum == 465 {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         smtpHost,
		}

		conn, err := tls.Dial("tcp", smtpHost+":"+smtpPort, tlsConfig)
		if err != nil {
			return fmt.Errorf("%s: TLS dial error: %v", funcName, err)
		}
		defer conn.Close()

		client, err := smtp.NewClient(conn, smtpHost)
		if err != nil {
			return fmt.Errorf("%s: SMTP client creation error: %v", funcName, err)
		}
		defer client.Quit()

		if err = client.Auth(smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)); err != nil {
			return fmt.Errorf("%s: authentication error: %v", funcName, err)
		}

		if err = client.Mail(from); err != nil {
			return fmt.Errorf("%s: MAIL FROM error: %v", funcName, err)
		}
		if err = client.Rcpt(to); err != nil {
			return fmt.Errorf("%s: RCPT TO error: %v", funcName, err)
		}

		wc, err := client.Data()
		if err != nil {
			return fmt.Errorf("%s: DATA error: %v", funcName, err)
		}
		_, err = wc.Write([]byte(msg))
		if err != nil {
			return fmt.Errorf("%s: message write error: %v", funcName, err)
		}
		if err := wc.Close(); err != nil {
			return fmt.Errorf("%s: write closer close error: %v", funcName, err)
		}

		fmt.Println(funcName, "Email sent via SMTPS (SSL)")
		return nil
	}

	// Default: STARTTLS (e.g., port 587)
	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
	if err != nil {
		return fmt.Errorf("%s: SMTP SendMail error: %v", funcName, err)
	}

	fmt.Println(funcName, "Email sent via SMTP (STARTTLS)")
	return nil
}



func getUserExtensions(user *User) []map[string]string {
	ret := []map[string]string{}
	sip_server := config[`public_hostname`]
	for _, x := range userextensions {
		if x.Username == user.Username {

			password, ok_password := extensions_auth[x.Extension]
		
			
			if ok_password {
				ep := endpoints[x.Extension].(map[string]string)
				callerid, ok_callerid := ep[`callerid`]
				if !ok_callerid {
					callerid = x.Extension
				}
				item := map[string]string{
					`sip_server`:    sip_server,
					`sip_portnum`:   `:8089`,
					`auth_user`:     x.Extension,
					`auth_password`: password,
					`display_name`:  callerid,
					`roleid`:        user.Roleid,
				}
				ret = append(ret, item)

			}

		}
	}

	return ret
}

func getUserExtensionsByUsername(username string) []map[string]string {
	ret := []map[string]string{}
	sip_server := config[`public_hostname`]

	_, user := findUserByUsername(username)
	if user == nil {
		return ret
	}

	for _, x := range userextensions {
		if x.Username == user.Username {
			password, ok_password := extensions_auth[x.Extension]
			
			
			if ok_password {
				ep := endpoints[x.Extension].(map[string]string)
				callerid, ok_callerid := ep[`callerid`]
				if !ok_callerid {
					callerid = x.Extension
				}
				item := map[string]string{
					`sip_server`:    sip_server,
					`sip_portnum`:   `:8089`,
					`auth_user`:     x.Extension,
					`auth_password`: password,
					`display_name`:  callerid,
					`roleid`:        user.Roleid,
				}
				ret = append(ret, item)

			}

		}
	}

	return ret
}

func getExtentionRangeByRole(roleid string) (int, int) {
	start_range := 100
	end_range := 10000
	range_val, ok_use_range := extension_range[roleid]
	if ok_use_range {
		start_end_arr := strings.Split(range_val, `,`)
		if len(start_end_arr) >= 2 {
			int_start, err := strconv.Atoi(start_end_arr[0])
			if err == nil {
				int_end, err := strconv.Atoi(start_end_arr[1])
				if err == nil {
					start_range = int_start
					end_range = int_end

				}

			}

		}
	}
	return start_range, end_range
}

func createUserExtensionForUser(user *User) *UserExtension {
	fmt.Println(`4334: START createUserExtensionForUser  user=`, user)

	// try reuseable extension
	reext := findReuseExtensionByRole(user.Roleid)

	if reext != nil {
		return createUserExtension(user.Username, reext.Number)
	}

	
	var ret *UserExtension

	start_range, end_range := getExtentionRangeByRole(user.Roleid)
	
	
	fmt.Println(`4377: start_range=`, start_range, `end_range=`, end_range)

	//get availble extension
	for i := start_range; i <= end_range; i++ {
		str_exten := strconv.Itoa(i)

		_, ext := findExtension(str_exten)

		if ext == nil {

			//create it
			fmt.Println(`4377: Called amiAddPJSIPExtension(str_exten=)`, str_exten)
			if amiAddPJSIPExtension(glo_ami, str_exten, uuid.New().String(), user.Name) {

				fmt.Println(`4377: Called createUserExtension(username=`, user.Username, `str_exten=`, str_exten, `)`)
				ret = createUserExtension(user.Username, str_exten)
				//ret = Map{`statusCode`:200}
				break

			} else {
				fmt.Println(`4377: amiAddPJSIPExtension() FALSE`)
				//return false
				break
			}
		} else {
			//see if available for reuse
			_, ret = findUserExtensionsByExtension(ext.Number)
			if ret == nil {
				//let's reuse
				ret = createUserExtension(user.Username, ext.Number)
				break

			}
		}

	}

	return ret
}

func createUserExtensionIfNeeded(user *User, exten string) *UserExtension {
	fmt.Println(`5501: START createUserExtensionForUser  user=`, user)

	_, x := findExtension(exten)
	if x == nil {
		//ok := amiAddPJSIPExtension(glo_ami,exten,uuid.New().String())
		secret := strings.ReplaceAll(uuid.New().String(), `-`, ``)
		ok := amiAddPJSIPExtension(glo_ami, exten, secret, user.Name)
		if ok {
			_, x = findExtension(exten)
		}
	}

	if x != nil {

		return createUserExtension(user.Username, x.Number)
	}

	return nil
}

func findReuseExtension() *Extension {
	//Search with Source xo
	for _, x := range extensions {
		if x.Source == `xo` {
			_, r := findUserExtensionsByExtension(x.Number)
			if r == nil {
				return &x
			}
		}

	}

	return nil
}

func findReuseExtensionByRole(roleid string) *Extension {
	//start_range := 100
	//end_range := 10000
	//Search with Source xo
	start_range, end_range := getExtentionRangeByRole(roleid)
	for _, x := range extensions {
		if x.Source == `xo` {
			int_exten, err := strconv.Atoi(x.Number)
			if err == nil && int_exten >= start_range && int_exten <= end_range {
				_, r := findUserExtensionsByExtension(x.Number)
				if r == nil {
					return &x
				}
			}

		}

	}

	return nil
}

func addXoManagedExtension(p_start, p_end int) {
	//Search with Source xo
	fmt.Println("START addXoManagedExtension")
	fmt.Println(`p_start=`, p_start)
	fmt.Println(`p_end=`, p_end)

	save_file := false

	if p_start > p_end {
		p_end = p_start
	}

	for n := p_start; n <= p_end; n++ {
		fmt.Println(`9230: n=`, n)

		p := strconv.Itoa(n)
		i, x := findExtension(p)

		if i >= 0 && x != nil && x.Number == p {

			extensions[i].Source = `xo`
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
    		defer cancel()
			dbUpsertExtension(ctx, extensions[i])
			save_file = true
		}

	}
	if save_file {
		//writeFileExtensions()

	}

}

func addXoManagedExtensionFromUserExtensions() {
	//Search with Source xo
	save_file := false

	for _, ue := range userextensions {

		i, x := findExtension(ue.Extension)

		if x != nil && x.Number == ue.Extension {

			extensions[i].Source = `xo`
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
    		defer cancel()
			dbUpsertExtension(ctx,extensions[i])
			save_file = true
		}

	}
	if save_file {
		//writeFileExtensions()
	}

}

func assignExtension(exten *Extension, to_username string) {

	//userextensions = append(userextensions, UserExtension{Username: to_username, Extension: exten.Number})
	//writeFileUserExtensions()
	new_ue := UserExtension{Username: to_username, Extension: exten.Number}
	userextensions = append(userextensions, new_ue)
	err := dbUpsertUserExtension(new_ue)
	if err != nil {
		fmt.Println(string(colorRed), err.Error())
		fmt.Println(string(colorReset))
	}

}



func findActivateExtensionRequest(username string) (int, *JoinRequest) {
	if len(joinrequests) > 0 {
		for i, j := range joinrequests {
			if j.Username == username {
				return i, &j
			}
		}
	}
	return -1, nil
}

func findJoinRequest(username string) (int, *JoinRequest) {
	if len(joinrequests) > 0 {
		for i, j := range joinrequests {
			if j.Username == username {
				return i, &j
			}
		}
	}
	return -1, nil
}

func findJoinRequestByEmail(email string) (int, *JoinRequest) {
	if len(joinrequests) > 0 {
		for i, j := range joinrequests {
			if j.Email == email {
				return i, &j
			}
		}
	}
	return -1, nil
}





func dbGetJoinRequestById(id string) (*JoinRequest, error) {
	qry := `
		SELECT id, hostname, email, username, name, created, status, auto_join, roleid
		FROM joinrequests
		WHERE id = ?
	`

	row := db.QueryRow(qry, id)

	var jr JoinRequest
	err := row.Scan(&jr.Id, &jr.Hostname, &jr.Email, &jr.Username, &jr.Name, &jr.Created, &jr.Status, &jr.Autojoin, &jr.Roleid)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &jr, nil
}


func dbUpdateJoinRequestStatus(username string, status int) error {
	query := `UPDATE joinrequests SET status = ? WHERE username = ?`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("prepare failed: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(status, username)
	if err != nil {
		return fmt.Errorf("exec failed: %w", err)
	}

	return nil
}

func findRequestToJoinDb(username, hostname string) (*JoinRequest, error) {
	query := `
		SELECT id, hostname, email, username, name, created, status, auto_join, roleid 
		FROM joinrequests 
		WHERE username = ? AND hostname = ?
		LIMIT 1`

	row := db.QueryRow(query, username, hostname)

	var jr JoinRequest
	err := row.Scan(&jr.Id, &jr.Hostname, &jr.Email, &jr.Username, &jr.Name, &jr.Created, &jr.Status, &jr.Autojoin, &jr.Roleid)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no match found
		}
		return nil, fmt.Errorf("query error: %w", err)
	}
	return &jr, nil
}
func getRequestToJoinDb(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	funcName := "getRequestToJoinDb()"
	fmt.Println(funcName, "20158: START")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var mapData Map
	if err := json.Unmarshal(reqBody, &mapData); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	strUsername := getStringFromMap(mapData, "username")
	if strUsername == "" {
		http.Error(w, "Missing 'username'", http.StatusBadRequest)
		return
	}
	strHostname := getStringFromMap(mapData, "hostname")
	if strHostname == "" {
		http.Error(w, "Missing 'hostname'", http.StatusBadRequest)
		return
	}

	// Call logic
	result := httpGetRequestToJoinResponse(strUsername,strHostname)
	        

	// Write correct status code based on result["statusCode"]
	statusCode := http.StatusOK
	if val, ok := result["statusCode"].(int); ok {
		statusCode = val
	}
	delete(result, "statusCode") // Optional: don't leak internal field

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(result)
}

func httpGetRequestToJoinResponse(username, hostname string) Map {
	funcName := "httpGetRequestToJoinResponse"
	fmt.Println(funcName, "10156.1: START", "username=", username, "hostname=", hostname)

	jr, err := findRequestToJoinDb(username, hostname)
	if err != nil {
		fmt.Println(funcName, "10156.2 DB lookup error:", err)
		return Map{"statusCode": http.StatusInternalServerError, "error": "Database error"}
	}

	fmt.Println(funcName, "10156.3: ")

	if jr != nil {
		fmt.Println(funcName, "10156.4: ", "js=", jr)
		return Map{
			"statusCode": http.StatusOK,
			"count":      1,
			"data":       jr,
		}
	}
	fmt.Println(funcName, "10156.5: ", "js=", jr)

	// Count = 0, but still valid request — no data found
	return Map{
		"statusCode": http.StatusOK,
		"count":      0,
	}
}



func httpNewJoinRequestDbResponse(
	id, hostname, username, name, email, photo, homepageurl, firstname, lastname, phonenumber string,
	auto_join bool,
) Map {
	func_name := `httpNewJoinRequestDbResponse()`
	fmt.Println(func_name, `START`, `username=`, username, `name=`, name, `email=`, email, `auto_join=`, auto_join)

	ret := Map{
		"statusCode": 500,
		"count":      0,
		"data":       nil,
	}

	autojoinInt := 0
	if auto_join {
		autojoinInt = 1
	}

	new_join := JoinRequest{
		Id:       id,
		Hostname: hostname,
		Created:  time.Now().Format(time.RFC3339),
		Username: username,
		Name:     name,
		Email:    email,
		Status:   1,
		Autojoin: autojoinInt,
		Roleid:   "customer",
	}

	if err := dbUpsertJoinRequest(new_join); err != nil {
		fmt.Println(func_name, "DB upsert error:", err)
		return Map{
			"statusCode": 500,
			"error":      "Failed to create/update join request",
			"count":      0,
			"data":       nil,
		}
	}

	jr, err := findRequestToJoinDb(username, hostname)
	if err != nil || jr == nil {
		fmt.Println(func_name, "DB fetch after upsert error:", err)
		return Map{
			"statusCode": 500,
			"error":      "Failed to retrieve join request after upsert",
			"count":      0,
			"data":       nil,
		}
	}

	
	if auto_join {
		
		//ret = approveJoinRequest(jr.Id)
		err := approveJoinRequest(jr.Id)
		if err != nil {

			jrFinal, err := findRequestToJoinDb(username, hostname)
			if err != nil || jrFinal == nil {
				fmt.Println(func_name, "DB fetch after approval error:", err)
				ret["statusCode"] = 500
				ret["error"] = "Failed to retrieve updated join request"
				ret["count"] = 0
				ret["data"] = nil
				return ret
			}

			ret["count"] = 1
			ret["data"] = *jrFinal // 👈 change from slice to single object
			return ret
		}
		
		
		
	}

	// No auto-join: return unapproved single object
	return Map{
		"statusCode": 202,
		"count":      1,
		"data":       *jr, // 👈 change from slice to single object
	}
}




func httpJoinResponse(username, name, email, photo, homepageurl, firstname, lastname, phonenumber string, auto_join bool) Map {
	func_name := `httpJoinResponse()`
	fmt.Println(func_name, `10340: START httpJoinResponse(username=`, username, `name=`, name, `email=`, email, `auto_join=`, auto_join)
	/*
		if len(roleid) == 0{
			roleid=`default`
		}
	*/
	//ret := []map[string]string{}

	roleid := `default`
	ret := Map{`statusCode`: 500}

	i_jr, jr := findJoinRequest(username)
	//i_jr,jr := findJoinRequestByEmail(email)
	if jr != nil {
		
		
		joinrequests[i_jr].Created = time.Now().Format(time.RFC3339)
		dbUpsertJoinRequest(joinrequests[i_jr])
		
		
	} else {
		new_join := JoinRequest{Created: time.Now().Format(time.RFC3339), Username: username, Name: name, Email: email}
		joinrequests = append(joinrequests, new_join)
		//writeFileJoinRequests()
		dbUpsertJoinRequest(new_join)
	}

	final_extensions := []map[string]string{}

	fmt.Println(func_name, `5543: findUserByUsername(`, username, `)`)
	iuser, user := findUserByUsername(username)

	//iuser,user := findUserByEmail(email)

	fmt.Println(func_name, `5544: user=`, user)

	if user == nil {
		//fmt.Println(func_name, `5547: username `,username,`does NOT exists!`)
		fmt.Println(func_name, `5547: user `, username, `does NOT exists!`)

		if auto_join {
			fmt.Println(func_name, `5550: auto_join is ON`)

			iuser, user = createUser(firstname, lastname, name, username, email, roleid, photo, phonenumber, homepageurl)
			//_,user = findUserByUsername(username)
			if user == nil {
				return Map{`statusCode`: 501}
			}

			//createUserExtensionForUser(user)

		} else {
			fmt.Println(func_name, `9623: auto_join is OFF`)

			//pushNotifiyAdmins(`Join Request`,fmt.Sprintf(`Request to join from %s`,username))
			pushNotifiyAdmins(`Join Request`, fmt.Sprintf(`Request to join from %s`, email))
			return Map{`statusCode`: 502}
		}

	} else {

		fmt.Println(func_name, `9632: User `, user.Username, ` exists.`)
		if user.Username != username {
			users[iuser].Username = username
			err := dbUpdateUsername(email, username)
			if err != nil {
				return Map{`statusCode`: 503}

			}

			//update userextensions
			//ues := findAllUserExtensions(user.Username)
			//if len(userextensions) > 0{
			for iue, ue := range userextensions {
				if ue.Username == user.Username {
					userextensions[iue].Username = username
					dbUpsertUserExtension(userextensions[iue])
				}
			}
			//}

		}
	}

	//fmt.Println(`3500: user found!`)

	//fmt.Println(`3500: userextensions=`,userextensions)

	final_extensions = getUserExtensions(user)
	fmt.Println(func_name, `5579: final_extensions=`, final_extensions)

	if len(final_extensions) > 0 {
		fmt.Println(`5582: userextensions > 0!`)

		ret = Map{`statusCode`: 200, `data`: final_extensions, `roleid`: user.Roleid}
	} else {
		fmt.Println(`5586: userextensions == 0!`)
		if auto_join {
			fmt.Println(`5588: auto_join true`)

			createUserExtensionForUser(user)
			
			
			final_extensions = getUserExtensions(user)
			

			ret = Map{`statusCode`: 200, `data`: final_extensions, `roleid`: user.Roleid}
			

		} else {
			fmt.Println(`5601: auto_join false`)
			
			ret = Map{`statusCode`: 200, `roleid`: user.Roleid}
		}

	}

	fmt.Println(func_name, `5609: ret=`, ret)
	return ret //Map{`statusCode`:201}
}


func httpCheckOutResponse(groupid string) Map {
	func_name := `httpCheckOutResponse()`
	fmt.Println(func_name, `1007: START httpCheckOutResponse(groupid=`, groupid)

	removeAllUserExtensionsInGroup(groupid)

	return Map{`statusCode`: 200} //ret //Map{`statusCode`:201}
}

func httpCheckInResponse(username, name, email, photo, homepageurl, firstname, lastname, phonenumber, exten_num, groupid string) Map {
	func_name := `httpActivateUserExtensionResponse()`
	fmt.Println(func_name, `10016: START httpActivateUserExtensionResponse(username=`, username, `name=`, name, `email=`, email, `exten_num=`, exten_num, `groupid=`, groupid)
	/*
		if len(roleid) == 0{
			roleid=`default`
		}
	*/
	//ret := []map[string]string{}

	roleid := `default`
	ret := Map{`statusCode`: 500}

	i_jr, jr := findJoinRequest(username)
	//i_jr,jr := findJoinRequestByEmail(email)
	if jr != nil {
		
		
		joinrequests[i_jr].Created = time.Now().Format(time.RFC3339)
		//writeFileJoinRequests()
		dbUpsertJoinRequest(joinrequests[i_jr])
		//}
		//}
		//}

	} else {
		new_join := JoinRequest{Created: time.Now().Format(time.RFC3339), Username: username, Name: name, Email: email}
		joinrequests = append(joinrequests, new_join)
		//writeFileJoinRequests()
		dbUpsertJoinRequest(new_join)
	}

	
	final_extensions := []map[string]string{}

	fmt.Println(func_name, `5543: findUserByUsername(`, username, `)`)
	_, user := findUserByUsername(username)

	//iuser,user := findUserByEmail(email)

	fmt.Println(func_name, `5544: user=`, user)

	if user == nil {
		//fmt.Println(func_name, `5547: username `,username,`does NOT exists!`)
		fmt.Println(func_name, `5547: user `, username, `does NOT exists! Create user automatically.`)

		
		
		_, user = createUser(firstname, lastname, name, username, email, roleid, photo, phonenumber, homepageurl)
		//_,user = findUserByUsername(username)
		if user == nil {
			return Map{`statusCode`: 501}
		}

		
		

	} //else {

	fmt.Println(func_name, `9835: User `, user.Email, ` exists.`)
	
	
	replaceUserExtensionNotIngroup(username, exten_num, groupid)
	final_extensions = getUserExtensions(user)

	ret = Map{`statusCode`: 200, `data`: final_extensions, `roleid`: user.Roleid}

	fmt.Println(func_name, `5609: ret=`, ret)
	return ret //Map{`statusCode`:201}
}

func httpActivateExtensionResponse(username, name, email, photo, homepageurl, firstname, lastname, phonenumber string, auto_join bool) Map {
	func_name := `httpActivateResponse()`
	fmt.Println(func_name, `10157: START httpActivateExtensionResponse(username=`, username, `name=`, name, `email=`, email, `auto_join=`, auto_join)
	
	

	roleid := `default`
	ret := Map{`statusCode`: 500}

	i_jr, jr := findActivateExtensionRequest(username)
	//i_jr,jr := findJoinRequestByEmail(email)
	if jr != nil {
		
		
		joinrequests[i_jr].Created = time.Now().Format(time.RFC3339)
		
		dbUpsertJoinRequest(joinrequests[i_jr])
	
		
	} else {
		new_join := JoinRequest{Created: time.Now().Format(time.RFC3339), Username: username, Name: name, Email: email}
		joinrequests = append(joinrequests, new_join)
		//writeFileJoinRequests()
		dbUpsertJoinRequest(new_join)
	}

	final_extensions := []map[string]string{}

	fmt.Println(func_name, `5543: findUserByUsername(`, username, `)`)
	iuser, user := findUserByUsername(username)

	//iuser,user := findUserByEmail(email)

	fmt.Println(func_name, `5544: user=`, user)

	if user == nil {
		//fmt.Println(func_name, `5547: username `,username,`does NOT exists!`)
		fmt.Println(func_name, `5547: user `, username, `does NOT exists!`)

		if auto_join {
			fmt.Println(func_name, `5550: auto_join is ON`)

			iuser, user = createUser(firstname, lastname, name, username, email, roleid, photo, phonenumber, homepageurl)
			//_,user = findUserByUsername(username)
			if user == nil {
				return Map{`statusCode`: 501}
			}

			//createUserExtensionForUser(user)

		} else {
			fmt.Println(func_name, `5563: auto_join is OFF`)

			//pushNotifiyAdmins(`Join Request`,fmt.Sprintf(`Request to join from %s`,username))
			pushNotifiyAdmins(`Join Request`, fmt.Sprintf(`Request to join from %s`, email))
			return Map{`statusCode`: 502}
		}

	} else {

		fmt.Println(func_name, `9975: User `, user.Email, ` exists.`)
		if user.Username != username {
			users[iuser].Username = username
			err := dbUpdateUsername(email, username)
			if err != nil {
				return Map{`statusCode`: 503}

			}

			//update userextensions
			//ues := findAllUserExtensions(user.Username)
			//if len(userextensions) > 0{
			for iue, ue := range userextensions {
				if ue.Username == user.Username {
					userextensions[iue].Username = username
					dbUpsertUserExtension(userextensions[iue])
				}
			}
			//}

		}
	}

	
	final_extensions = getUserExtensions(user)
	fmt.Println(func_name, `10004: final_extensions=`, final_extensions)

	if len(final_extensions) > 0 {
		fmt.Println(`10007: userextensions > 0!`)

		ret = Map{`statusCode`: 200, `data`: final_extensions, `roleid`: user.Roleid}
	} else {
		fmt.Println(`10011: userextensions == 0!`)
		if auto_join {
			fmt.Println(`10013: auto_join true`)

			createUserExtensionForUser(user)
			
			
			final_extensions = getUserExtensions(user)
			
			
			ret = Map{`statusCode`: 200, `data`: final_extensions, `roleid`: user.Roleid}
			
			

		} else {
			fmt.Println(`5601: auto_join false`)
			//pushNotifiyAdmins(`Request extention`,fmt.Sprintf(`Requested to assign extension for %s.`,user.Username))
			//ret = Map{`statusCode`:201}
			ret = Map{`statusCode`: 200, `roleid`: user.Roleid}
		}

	}

	fmt.Println(func_name, `10068: ret=`, ret)
	return ret //Map{`statusCode`:201}
}

func httpDisJoinResponse(username string) Map {
	fmt.Println(`3500: START httpDisJoinResponse(username`, username)

	deleteUser(username)

	return Map{`statusCode`: 200}
}


func httpAdminAddUser(map_data Map) (*User, error) {
	func_name := `httpAdminAddUser`
	admin := ``
	name := ``
	firstname := ``
	lastname := ``
	username := ``
	password := ``
	email := ``
	exten := ``
	homepageurl := fmt.Sprintf(`https://%s:8180`, hostname)
	//str_name := `Unknown`
	roleid := `guest`

	dob := ``
	gender := ``
	address := ``
	phonenumber := ``
	imageurl := ``
	emailsubject := ``
	emailbody := ``
	upsert := false
	isroot := false

	p_admin, ok_admin := map_data[`admin`]
	//p_name,ok_name := map_data[`name`]
	p_username, ok_username := map_data[`username`]
	p_password, ok_password := map_data[`password`]
	p_email, ok_p_email := map_data[`email`]
	p_exten, ok_exten := map_data[`exten`]
	p_roleid, ok_roleid := map_data[`roleid`]
	p_firstname, ok_firstname := map_data[`firstname`]
	p_lastname, ok_lastname := map_data[`lastname`]
	p_dob, ok_dob := map_data[`dob`]
	p_gender, ok_gender := map_data[`gender`]
	p_address, ok_address := map_data[`address`]
	p_phonenumber, ok_phonenumber := map_data[`phonenumber`]
	p_imageurl, ok_imageurl := map_data[`imageurl`]
	p_homepageurl, ok_homepageurl := map_data[`homepageurl`]
	p_emailsubject, ok_emailsubject := map_data[`emailsubject`]
	p_emailbody, ok_emailbody := map_data[`emailbody`]

	p_upsert, ok_upsert := map_data[`upsert`]
	p_isroot, ok_isroot := map_data[`isroot`]

	if ok_admin {
		admin = p_admin.(string)
	}
	/*if ok_name{
	name= p_name.(string)
	}
	*/
	if ok_username {
		username = p_username.(string)
	}
	if ok_password {
		password = p_password.(string)
	}

	if ok_p_email {
		email = p_email.(string)
	}
	email = username
	if ok_exten {
		exten = p_exten.(string)
	}
	if ok_roleid {
		roleid = p_roleid.(string)
	}
	if ok_firstname {
		firstname = p_firstname.(string)
	}
	if ok_lastname {
		lastname = p_lastname.(string)
	}
	if ok_dob {
		dob = p_dob.(string)
	}
	if ok_dob {
		dob = p_dob.(string)
	}

	if ok_gender {
		gender = p_gender.(string)
	}
	if ok_address {
		address = p_address.(string)
	}

	if ok_phonenumber {
		phonenumber = p_phonenumber.(string)
	}
	if ok_imageurl {
		imageurl = p_imageurl.(string)
	}

	if ok_homepageurl {
		homepageurl = p_homepageurl.(string)
	}

	if ok_emailsubject {
		emailsubject = p_emailsubject.(string)
	}
	if ok_emailbody {
		emailbody = p_emailbody.(string)
	}

	if ok_upsert {
		upsert = p_upsert.(bool)
	}

	if ok_isroot {
		isroot = p_isroot.(bool)
	}

	//if len(name) == 0{
	name = fmt.Sprintf(`%s %s`, firstname, lastname)
	name = strings.TrimRight(name, ` `)
	//}

	if len(name) == 0 {
		firstname = `Unknown`
		name = `Unknown`
	}

	//exten_to_username,ok_exten_to_username := map_data[`exten_to_username`]

	//sessionid,ok_sessionid := map_data[`sessionid`]
	//if ok_username && ok_sessionid{
	if !isroot {
		if ok_admin && ok_username && ok_upsert {
			fmt.Println(func_name, `7672: params OK`)

			/////////

			if !usernameInputValidation(admin) {
				//return nil,errors.New(fmt.Sprintf(`500.1 username %s does not pass validation.`,admin))
				return nil, errors.New(fmt.Sprintf(`500.1 username %s does not pass validation.`, admin))
			}

		}
		_, u_admin := findUserByUsername(admin)

		if u_admin == nil || u_admin.Roleid != `admin` {
			//fmt.Println(`500.2: UNAUTHORIZE username`,p_admin)
			//return nil,errors.New(fmt.Sprintf(`username %s is not an admin.`,p_admin))
			return nil, errors.New(fmt.Sprintf(`500.2: UNAUTHORIZE admin `, admin))
		} //else{
	}
	httpXoAdminSignupUserIfNeeded(username, password, email, firstname, lastname, dob, gender, address, phonenumber, imageurl, roleid)
	//return doAdminAddUser(firstname,lastname, name,phonenumber,imageurl, username,email,roleid,exten,homepageurl,upsert)
	ret_user, err := doAdminAddUser(firstname, lastname, name, phonenumber, imageurl, username, email, roleid, exten, homepageurl, upsert)
	if ret_user != nil {
		fmt.Println(`7729: doAdminAddUser OK`)
		if len(emailbody) > 0 || len(emailsubject) > 0 {
			fmt.Println(`7729: doAdminAddUser Send email...`)
			httpEmailNotify(email, emailsubject, emailbody)
			fmt.Println(`7729: doAdminAddUser Send email...done`)
		}
	} else {
		fmt.Println(`7729: doAdminAddUser FAILED`)
	}

	return ret_user, err

}

func httpAdminEditUser(map_data Map) error {
	func_name := `httpAdminEditUser`
	admin := ``
	//name := ``
	firstname := ``
	lastname := ``
	username := ``
	
	exten := ``
	
	roleid := `guest`

	
	emailsubject := ``
	emailbody := ``
	
	isroot := false

	p_admin, ok_admin := map_data[`admin`]
	
	p_firstname, ok_firstname := map_data[`firstname`]
	p_lastname, ok_lastname := map_data[`lastname`]

	p_username, ok_username := map_data[`username`]
	p_exten, ok_exten := map_data[`exten`]
	p_roleid, ok_roleid := map_data[`roleid`]
	
	p_emailsubject, ok_emailsubject := map_data[`emailsubject`]
	p_emailbody, ok_emailbody := map_data[`emailbody`]

	
	
	if ok_admin {
		admin = p_admin.(string)
	}

	if ok_firstname {
		firstname = p_firstname.(string)
	}

	if ok_lastname {
		lastname = p_lastname.(string)
	}

	if ok_username {
		username = p_username.(string)
	}
	
	
	if ok_exten {
		exten = p_exten.(string)
	}

	if ok_roleid {
		roleid = p_roleid.(string)
	}
	
	

	if ok_emailsubject {
		emailsubject = p_emailsubject.(string)
	}
	if ok_emailbody {
		emailbody = p_emailbody.(string)
	}

	

	if admin == `root` {
		isroot = true
	}
	
	
	if !isroot {
		//if ok_admin && ok_username && ok_upsert{
		if ok_admin && ok_username {
			fmt.Println(func_name, `7672: params OK`)

			if !usernameInputValidation(admin) {
				//return nil,errors.New(fmt.Sprintf(`500.1 username %s does not pass validation.`,admin))
				return errors.New(fmt.Sprintf(`500.1 username %s does not pass validation.`, admin))
			}

		}
		_, u_admin := findUserByUsername(admin)

		if u_admin == nil || u_admin.Roleid != `admin` {
			//fmt.Println(`500.2: UNAUTHORIZE username`,p_admin)
			//return nil,errors.New(fmt.Sprintf(`username %s is not an admin.`,p_admin))
			return errors.New(fmt.Sprintf(`500.2: UNAUTHORIZE admin `, admin))
		} //else{
	}

	//httpXoAdminSignupUserIfNeeded(username,password,email,firstname,lastname,dob,gender,address,phonenumber,imageurl,roleid)

	user, err := doAdminEditUser(username, roleid, exten, firstname, lastname)
	if err != nil {
		
		
		return err
	}

	if user != nil {
		if len(emailbody) > 0 || len(emailsubject) > 0 {
			fmt.Println(`7729: doAdminAddUser Send email...`)
			httpEmailNotify(user.Email, emailsubject, emailbody)
			fmt.Println(`7729: doAdminAddUser Send email...done`)
		}

	}

	return nil

}

func httpAdminRemoveUser(map_data Map) error {
	func_name := `httpAdminRemoveUser`
	admin := ``

	username := ``

	isroot := false

	p_admin, ok_admin := map_data[`admin`]

	p_username, ok_username := map_data[`username`]

	if ok_admin {
		admin = p_admin.(string)
	}

	if ok_username {
		username = p_username.(string)
	}

	if admin == `root` {
		isroot = true
	}

	if !isroot {
		//if ok_admin && ok_username && ok_upsert{
		if ok_admin && ok_username {
			fmt.Println(func_name, `7672: params OK`)

			if !usernameInputValidation(admin) {
				//return nil,errors.New(fmt.Sprintf(`500.1 username %s does not pass validation.`,admin))
				return errors.New(fmt.Sprintf(`500.1 username %s does not pass validation.`, admin))
			}

		}
		_, u_admin := findUserByUsername(admin)

		if u_admin == nil || u_admin.Roleid != `admin` {
			//fmt.Println(`500.2: UNAUTHORIZE username`,p_admin)
			//return nil,errors.New(fmt.Sprintf(`username %s is not an admin.`,p_admin))
			return errors.New(fmt.Sprintf(`500.2: UNAUTHORIZE admin `, admin))
		} //else{
	}

	//httpXoAdminSignupUserIfNeeded(username,password,email,firstname,lastname,dob,gender,address,phonenumber,imageurl,roleid)

	err := doAdminRemoveUser(username)
	if err != nil {
		
		
		return err
	}

	return nil

}


func doAdminAddUser(p_firstname, p_lastname, p_name, p_phonenumber, p_photo, p_username, p_email, roleid, p_exten, homepageurl string, upsert bool) (*User, error) {

	log.Println(`10794.1: START doAdminAddUser`,
		`p_firstname=`, p_firstname,
		`p_lastname=`, p_lastname,
		`p_name=`, p_name,
		`p_phonenumber=`, p_phonenumber,
		`p_photo=`, p_photo,
		`p_username=`, p_username,
		`p_email=`, p_email,
		`roleid=`, roleid,
		`p_exten=`, p_exten,
		`homepageurl=`, homepageurl,
		`upsert=`, upsert)
	
	func_name := `doAdminAddUser`
	
	
	if !emailInputValidation(p_username) {
		return nil, errors.New(fmt.Sprintf(`username %s does not pass validation.`, p_username))
	}

	_, new_user := findUserByUsername(p_username)
	log.Println(`10794.2:`, func_name, `findUserByUsername(p_username=`, p_username, `new_user=`, new_user)

	
	
	username := p_username
	
	
	if new_user != nil {
		//if !upsert{
		log.Println(`10794.3:`, func_name, `User exists! findUserByUsername(p_username=`, p_username, `new_user=`, new_user)
		if !upsert {
			return new_user, nil
		}
		//}
	} else {
		log.Println(`10794.4:`, func_name, `User DOES NOT exists! findUserByUsername(p_username=`, p_username, `new_user=`, new_user)
		
		//u_map := httpGetUserNameByEmail(p_email)
		u_map := httpGetUserName(p_username)

		_, ok_username := u_map[`username`]
		//email := p_email
		if ok_username {
			//p_username = musername
			m_email, ok_email := u_map[`email`]
			if ok_email {
				p_email = m_email
			}
			m_firstname, ok_firstname := u_map[`firstName`]
			if ok_firstname {
				p_firstname = m_firstname
			}
			m_lastname, ok_lastname := u_map[`lastName`]
			if ok_lastname {
				p_lastname = m_lastname
			}
			m_name, ok_name := u_map[`fullName`]
			if ok_name {
				p_name = m_name
			}
			m_photo, ok_photo := u_map[`photo`]
			if ok_photo {
				p_photo = m_photo
			}

			
		} else {
			//username = p_email
		}

		_, new_user = createUser(p_firstname, p_lastname, p_name, p_username, p_email, roleid, p_photo, p_phonenumber, homepageurl)
		//log.Println(`4966.0:`,func_name,`createUser(p_username=`,username,`new_user=`,new_user)
		fmt.Println(string(colorGreen), `10794.5:`, `createUser(p_username=`, username, `new_user=`, new_user)
		fmt.Println(string(colorReset))
		//_,user = findUserByUsername(username)
		if new_user == nil {
			return nil, errors.New(fmt.Sprintf(`10794.6: There was an error creating user %s.`, username))
		}
	}

	log.Println(`10794.7:`, func_name, `new_user=`, new_user)

	if len(p_exten) == 0 {
		return new_user, nil
	}

	// let's make sure maximum count is met
	max_extension_count_per_user := 1

	str_max_count, ok_str_max_count := config[`max_extension_count_per_user`]

	if ok_str_max_count {
		max_count, err := strconv.Atoi(str_max_count)
		if err != nil {
			max_extension_count_per_user = max_count
		}
	} else {
		config[`max_extension_count_per_user`] = strconv.Itoa(max_extension_count_per_user)
		writeFileConfig()
	}

	ue, err := findAllUserExtensions(new_user.Username)
	if err != nil {
		return new_user, nil
	}

	if len(ue) >= max_extension_count_per_user {

		//if !upsert{
		return new_user, nil
		//}

	}

	if len(ue) > 0 && upsert {
		deleteUserExtension(username, ue[0].Extension)
	}

	//
	if p_exten == `auto` {
		p_exten = ``
	}

	if len(p_exten) == 0 {
		log.Println(`10794.8:`, func_name, `new_user=`, new_user)
		createUserExtensionForUser(new_user)
		log.Println(`10794.9:`, func_name)

	} else {
		log.Println(`10794.10:`, func_name)
		ue := createUserExtensionIfNeeded(new_user, p_exten)
		if ue == nil {
			log.Println(`10794.11:`, func_name, `createUserExtensionIfNeeded FAILED!`)

		}
		
		

	}

	//} else{

	return new_user, nil

}

func doAdminEditUser(p_username, roleid, p_exten, p_firstname, p_lastname string) (*User, error) {
	func_name := `doAdminEditUser`
	fmt.Println(`10989.1: START`, func_name,

		`p_firstname=`, p_firstname,
		`p_lastname=`, p_lastname,
		
		
		`p_username=`, p_username,
		
		`roleid=`, roleid,
		`p_exten=`, p_exten,
	
	)

	err_load := dbLoadUserExtensions()

	if err_load != nil {
		return nil, err_load
	}

	_, existing_user := findUserByUsername(p_username)

	//fmt.Println(`4966.0:`,func_name,`findUserByUsername(p_username=`,p_username,`existing_user=`,existing_user)
	if existing_user == nil {
		fmt.Println(`10989..1: username %s is invalid.`, p_username)
		return existing_user, errors.New(fmt.Sprintf(`username %s is invalid.`, p_username))
	}

	//if existing_user.Roleid != roleid || existing_user.Firstname {
		adminUpdateUser(existing_user.Username, roleid, p_firstname, p_lastname)

	//}

	p_extens := []string{}

	if len(p_exten) > 0 {

		p_extens = strings.Split(p_exten, "\n")
		fmt.Println(`p_extens=`, p_extens)

	}

	if len(p_extens) == 0 {
		fmt.Println(`10989.2: DEBUG doAdminEditUser.1`)
		//if len(ues) > 0{
		removeAllUserExtensions(existing_user.Username)

		//}
		return existing_user, nil
	}

	fmt.Println(`10989.3: DEBUG doAdminEditUser.1.1`)
	//see if there's a change
	ues, err := findAllUserExtensions(existing_user.Username)
	if err != nil {
		return existing_user, nil
	}

	new_extens := []string{}
	remove_extens := []string{}

	if len(ues) > 0 {
		fmt.Println(`10989.4: DEBUG doAdminEditUser.2`)
		for _, item := range ues {
			isincluded := false
			for _, ext_item := range p_extens {
				fmt.Println(`10989.5: DEBUG doAdminEditUser.2.1 ext_item=`, ext_item, `item.Extension=`, item.Extension)
				if ext_item == item.Extension {
					fmt.Println(`10989.6: DEBUG doAdminEditUser.2.2`)
					isincluded = true
					break
				}
			}
			if !isincluded {
				remove_extens = append(remove_extens, item.Extension)
			}

		}
		for _, item := range p_extens {
			isincluded := false
			for _, ue_item := range ues {
				fmt.Println(`10989.7: DEBUG doAdminEditUser.2.3 item=`, item, `item.Extension=`, ue_item.Extension)

				if item == ue_item.Extension {

					isincluded = true
					break
				}
			}
			if !isincluded {
				new_extens = append(new_extens, item)
			}

		}

	} else {
		fmt.Println(`10989.8: DEBUG doAdminEditUser.3`)
		new_extens = p_extens
	}

	if len(remove_extens) > 0 {
		fmt.Println(`10989.9: DEBUG doAdminEditUser.4`)
		fmt.Println(`10989.10: user exentions to remove`, remove_extens)
		for _, rem_exten := range remove_extens {
			deleteUserExtension(existing_user.Username, rem_exten)

		}
	}
	if len(new_extens) > 0 {
		fmt.Println(`10989.11: DEBUG doAdminEditUser.5`)
		fmt.Println(`10989.12: user exentions to add`, new_extens)
		// let's make sure maximum count is met
		max_extension_count_per_user := 1

		str_max_count, ok_str_max_count := config[`max_extension_count_per_user`]

		if ok_str_max_count {
			max_count, err := strconv.Atoi(str_max_count)
			if err != nil {
				max_extension_count_per_user = max_count
			}
		} else {
			config[`max_extension_count_per_user`] = strconv.Itoa(max_extension_count_per_user)
			writeFileConfig()
		}

		for _, new_exten := range new_extens {

			ok_add := false
			_, ext := findExtension(new_exten)

			if ext == nil {
				secret := strings.ReplaceAll(uuid.New().String(), `-`, ``)
				ok_add = amiAddPJSIPExtension(glo_ami, new_exten, secret, existing_user.Name)

			} else {
				ok_add = true
			}

			if ok_add {
				fmt.Println(`10989.13: DEBUG doAdminEditUser.6`)
				fmt.Println(`10989.14: doAdminEditUser() Add PJSIP SUCCESS!`, new_exten)
				err := dbUpsertUserExtension(UserExtension{Username: existing_user.Username, Extension: new_exten})
				if err != nil {
					fmt.Println(string(colorRed), err.Error())
					fmt.Println(string(colorReset))

				}
			} else {
				fmt.Println(`10989.15: DEBUG doAdminEditUser.7 ok_add is false`)
			}

		}

	}

	fmt.Println(`10989.16: DEBUG doAdminEditUser.8`)

	return existing_user, nil

}

func doAdminRemoveUser(p_username string) error {
	func_name := `doAdminRemoveUser`
	fmt.Println(`1155.1: START`, func_name, `p_username=`, p_username)

	
	
	err := removeUser(p_username)

	return err

}

func httpAddAssignExtenResponse(p_admin, p_exten, p_username string) bool {
	//ret := []map[string]string{}
	//exten_num := p_exten
	func_name := `httpAddAssignExtenResponse`
	fmt.Println(`httpAddAssignExtenResponse(p_admin=`, p_admin, `p_exten=`, p_exten, `p_username=`, p_username)

	if usernameInputValidation(p_admin) {
		return false
	}

	_, user := findUserByUsername(p_admin)
	if user == nil || user.Roleid != `admin` {
		fmt.Println(`3897: UNAUTHORIZE username`, p_admin)
		return false
	}

	if usernameInputValidation(p_username) {
		return false
	}

	_, p_exten_to_user := findUserByUsername(p_username)
	if p_exten_to_user == nil {
		fmt.Println(`3897: Does not exists username`, p_username)
		return false
	}

	var exten *Extension

	if len(p_exten) == 0 {
		exten = findReuseExtension()
		if exten != nil {
			assignExtension(exten, p_username)
			return true
		}

	}

	//index := 0
	if len(p_exten) > 0 {

		fmt.Println(`3899.1: DEBUG:`, func_name)

		// check if exten is already assigned to username
		ue := findOneUserExtensions(p_username, p_exten)
		if ue != nil {
			return true
		}

		_, exten = findExtension(p_exten)
		if exten == nil {
			fmt.Println(`3899.2: DEBUG:`, func_name)
			//create
			if !extensionInputValidation(p_exten) {
				fmt.Println(`399.3: Invalid extension length. max 9 digit required.`)
				return false
			}

			if amiAddPJSIPExtension(glo_ami, p_exten, uuid.New().String(), p_exten_to_user.Name) {

				_, exten = findExtension(p_exten)
				if exten == nil {
					fmt.Println(`3899.4: DEBUG:`, func_name)
					return false
				}
			} else {
				fmt.Println(`3899.5: DEBUG:`, func_name)
				return false
			}

		}

	} else {
		fmt.Println(`3899.6: DEBUG:`, func_name)
		//auto generate
		gen_exten := ``
		start_num := 500
		end_num := 1000000
		for i := start_num; i < end_num; i++ {
			_, ext := findExtension(string(i))
			if ext == nil {
				gen_exten = string(i)
			}
		}

		if len(gen_exten) > 0 {
			if amiAddPJSIPExtension(glo_ami, gen_exten, uuid.New().String(), p_exten_to_user.Name) {

				_, exten = findExtension(gen_exten)
				if exten == nil {
					fmt.Println(`3899.7: DEBUG:`, func_name)
					return false
				}
			} else {
				fmt.Println(`3899.8: DEBUG:`, func_name)
				return false
			}
		}

	}

	if exten != nil {
		fmt.Println(`3946.1: DEBUG:`, func_name)

		_, assign_to_user := findUserByUsername(p_username)

		if assign_to_user != nil {
			ue := findOneUserExtensions(p_username, exten.Number)
			if ue == nil {
				fmt.Println(`3946.2: DEBUG:`, func_name)
				new_ue := UserExtension{Username: p_username, Extension: exten.Number}
				userextensions = append(userextensions, new_ue)
				err := dbUpsertUserExtension(new_ue)
				if err != nil {
					fmt.Println(string(colorRed), err.Error())
					fmt.Println(string(colorReset))
				}
				//writeFileUserExtensions()
			} else {
				fmt.Println(`3946.3: DEBUG:`, func_name)
			}

		} else {
			fmt.Println(`3946.4: DEBUG:`, func_name)
		}

	} else {
		fmt.Println(`3946.5: DEBUG:`, func_name)
	}

	return true
}



func httpAdminAddUserExten(p_username, p_exten string) bool {
	//ret := []map[string]string{}
	//exten_num := p_exten
	func_name := `httpAdminAddUserExten`
	//fmt.Println(`httpAdminAddUserExten(p_admin=`,p_admin,`p_exten=`,p_exten,`p_username=`,p_username)
	/*
		if p_admin != nil || p_admin.Roleid != `admin` {
			return false
		}
	*/

	if usernameInputValidation(p_username) {
		return false
	}

	_, p_exten_to_user := findUserByUsername(p_username)
	if p_exten_to_user == nil {
		fmt.Println(`3897: Does not exists username`, p_username)
		return false
	}

	var exten *Extension

	if len(p_exten) == 0 {
		exten = findReuseExtension()
		if exten != nil {
			assignExtension(exten, p_username)
			return true
		}

	}

	//index := 0
	if len(p_exten) > 0 {

		fmt.Println(`3899.1: DEBUG:`, func_name)

		// check if exten is already assigned to username
		ue := findOneUserExtensions(p_username, p_exten)
		if ue != nil {
			return true
		}

		_, exten = findExtension(p_exten)
		if exten == nil {
			fmt.Println(`3899.2: DEBUG:`, func_name)
			//create
			if !extensionInputValidation(p_exten) {
				fmt.Println(`399.3: Invalid extension length. max 9 digit required.`)
				return false
			}

			if amiAddPJSIPExtension(glo_ami, p_exten, uuid.New().String(), p_exten_to_user.Name) {

				_, exten = findExtension(p_exten)
				if exten == nil {
					fmt.Println(`3899.4: DEBUG:`, func_name)
					return false
				}
			} else {
				fmt.Println(`3899.5: DEBUG:`, func_name)
				return false
			}

		}

	} else {
		fmt.Println(`3899.6: DEBUG:`, func_name)
		//auto generate
		gen_exten := ``
		start_num := 500
		end_num := 1000000
		for i := start_num; i < end_num; i++ {
			_, ext := findExtension(string(i))
			if ext == nil {
				gen_exten = string(i)
			}
		}

		if len(gen_exten) > 0 {
			if amiAddPJSIPExtension(glo_ami, gen_exten, uuid.New().String(), p_exten_to_user.Name) {

				_, exten = findExtension(gen_exten)
				if exten == nil {
					fmt.Println(`3899.7: DEBUG:`, func_name)
					return false
				}
			} else {
				fmt.Println(`3899.8: DEBUG:`, func_name)
				return false
			}
		}

	}

	if exten != nil {
		fmt.Println(`3946.1: DEBUG:`, func_name)

		_, assign_to_user := findUserByUsername(p_username)

		if assign_to_user != nil {
			ue := findOneUserExtensions(p_username, exten.Number)
			if ue == nil {
				fmt.Println(`3946.2: DEBUG:`, func_name)
				new_ue := UserExtension{Username: p_username, Extension: exten.Number}
				userextensions = append(userextensions, new_ue)
				err := dbUpsertUserExtension(new_ue)
				if err != nil {
					fmt.Println(string(colorRed), err.Error())
					fmt.Println(string(colorReset))
				}
				//writeFileUserExtensions()
			} else {
				fmt.Println(`3946.3: DEBUG:`, func_name)
			}

		} else {
			fmt.Println(`3946.4: DEBUG:`, func_name)
		}

	} else {
		fmt.Println(`3946.5: DEBUG:`, func_name)
	}

	return true
}

func httpAdminAddExten(p_exten, displayname string) *Extension {
	//ret := []map[string]string{}
	//exten_num := p_exten
	func_name := `httpAdminAddExten`
	//fmt.Println(`httpAdminAddExten(p_admin=`,p_admin,`p_exten=`,p_exten,`p_username=`,p_username)
	/*
		if p_admin != nil || p_admin.Roleid != `admin` {
			return false
		}
	*/

	var exten *Extension

	if len(p_exten) == 0 {
		//exten = findReuseExtension()

		/*if exten != nil{
			assignExtension(exten,p_username)
			return true
		}
		*/
		//return exten
		return nil

	}

	if !extensionInputValidation(p_exten) {
		fmt.Println(`399.3: Invalid extension length. max 9 digit required.`)
		return nil
	}

	

	_, exten = findExtension(p_exten)
	if exten == nil {
		fmt.Println(`3899.2: DEBUG:`, func_name)
		//create

		if amiAddPJSIPExtension(glo_ami, p_exten, uuid.New().String(), displayname) {

			_, exten = findExtension(p_exten)
			if exten == nil {
				fmt.Println(`3899.4: DEBUG:`, func_name)
				return nil
			}
		} else {
			fmt.Println(`3899.5: DEBUG:`, func_name)
			return nil
		}

	}

	
	
	return nil
}

func httpCDRResponse_dep(username, q string, g string, skip, limit int) Map {
	//func httpCDRResponse(data Map) Map{
	//ret := []map[string]string{}
	//exten_num := p_exten
	func_name := `httpAddCDRResponse`
	
	

	extens, err := findAllUserExtensions(username)

	if err != nil {
		return Map{`statusCode`: 501}
	}

	if extens != nil && len(extens) > 0 {

		
		

		qry := ``

		//qry_count := ``

		qry_update := ``
		
		

		if len(extens) > 1 {
			p := ``

			for _, ext := range extens {
				p = fmt.Sprintf(`%s'%s',`, p, ext.Extension)
			}
			p = strings.TrimRight(p, `,`)

			switch g {
			case `RECENT`:
				//SELECT  COUNT(*) as total  FROM (SELECT COUNT(*) FROM  cdr WHERE src IN('100') OR dst IN('100') GROUP BY src,dst) as A;
				//qry_count = fmt.Sprintf( "SELECT  COUNT(*) as total FROM (SELECT COUNT(*) FROM  cdr WHERE src IN('%s') OR dst IN('%s') GROUP BY src,dst) as A",p,p)

				qry = fmt.Sprintf("SELECT  uniqueid, MAX(calldate) AS calldate ,clid,src,dst,duration,billsec,disposition  FROM  cdr WHERE src IN('%s') OR dst IN('%s') GROUP BY src,dst ORDER by calldate DESC LIMIT %v OFFSET %v", p, p, limit, skip)

			case `MISSED`:
				//qry_count = fmt.Sprintf( "SELECT  COUNT(*) as total FROM (SELECT COUNT(*) FROM  cdr WHERE disposition <> 'ANSWERED' AND (src IN('%s') OR dst IN('%s')) GROUP BY src,dst) as A",p,p)
				qry = fmt.Sprintf("SELECT  uniqueid, MAX(calldate) AS calldate ,clid,src,dst,duration,billsec,disposition FROM cdr WHERE disposition <> 'ANSWERED' AND (src IN('%s') OR dst IN('%s')) GROUP BY src,dst ORDER by calldate DESC LIMIT %v OFFSET %v", p, p, limit, skip)

			default: //ALL

				//qry_count = fmt.Sprintf( "SELECT  COUNT(*) as total FROM (SELECT COUNT(*) FROM  cdr WHERE src IN('%s') OR dst IN('%s')) as A",p,p)

				qry = fmt.Sprintf("SELECT uniqueid, calldate,clid,src,dst,duration,billsec,disposition FROM cdr WHERE src IN(%s) OR dst IN(%s) ORDER BY calldate DESC LIMIT %v OFFSET %v", p, p, limit, skip)
			}
			qry_update = fmt.Sprintf("UPDATE extensions SET missed = 0 WHERE number IN(%s)", p)

			
			

		} else {
			switch g {
			case `RECENT`:
				//qry_count = fmt.Sprintf( "SELECT  COUNT(*) as total FROM (SELECT COUNT(*) FROM  cdr WHERE src = '%s' OR dst = '%s' GROUP BY src,dst) as A",extens[0].Extension,extens[0])

				qry = fmt.Sprintf("SELECT  uniqueid, MAX(calldate) AS calldate ,clid,src,dst,duration,billsec,disposition  FROM cdr WHERE src = '%s' OR dst = '%s' GROUP BY src,dst ORDER by calldate DESC LIMIT %v OFFSET %v", extens[0].Extension, extens[0].Extension, limit, skip)

			case `MISSED`:
				//qry_count = fmt.Sprintf( "SELECT  COUNT(*) as total FROM (SELECT COUNT(*) FROM  cdr WHERE disposition <> 'ANSWERED' AND (src = '%s' OR dst = '%s') GROUP BY src,dst) as A",extens[0].Extension,extens[0])

				qry = fmt.Sprintf("SELECT  uniqueid, MAX(calldate) AS calldate ,clid,src,dst,duration,billsec,disposition  FROM cdr WHERE disposition <> 'ANSWERED' AND (src = '%s' OR dst = '%s') GROUP BY src,dst ORDER BY calldate DESC LIMIT %v OFFSET %v", extens[0].Extension, extens[0].Extension, limit, skip)

			default: //ALL

				//qry_count = fmt.Sprintf( "SELECT  COUNT(*) as total FROM (SELECT COUNT(*) FROM  cdr WHERE src = '%s' OR dst = '%s') as A",extens[0].Extension,extens[0])

				qry = fmt.Sprintf("SELECT uniqueid, calldate ,clid,src,dst,duration,billsec,disposition FROM cdr WHERE src = '%s' OR dst = '%s' ORDER BY calldate DESC LIMIT %v OFFSET %v", extens[0].Extension, extens[0].Extension, limit, skip)
			}
			//qry = fmt.Sprintf("SELECT uniqueid,calldate,clid,src,dst,duration,billsec,disposition FROM cdr WHERE src = '%s' OR dst = '%s'  ORDER BY calldate DESC LIMIT %v OFFSET %v", extens[0].Extension,extens[0].Extension,limit,skip)
			qry_update = fmt.Sprintf("UPDATE extensions SET missed = 0 WHERE number = '%s'", extens[0].Extension)
		}

		fmt.Println(`8299:`, func_name, `qry=`, qry)

		_, err_update := db.Query(qry_update)
		if err_update != nil {
			fmt.Println(`7542:`, func_name, `err_update:=`, err_update.Error())
		}
		results, err := db.Query(qry)

		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(func_name, err.Error())
			return Map{`statusCode`: 502}
		}

		var cdrs []CDR

		for results.Next() {
			var cdr CDR
			// for each row, scan the result into our tag composite object
			err = results.Scan(&cdr.Uniqueid, &cdr.Calldate, &cdr.Clid, &cdr.Dst, &cdr.Src, &cdr.Duration, &cdr.Billsec, &cdr.Disposition)
			if err != nil {
				//panic(err.Error()) // proper error handling instead of panic in your app
				fmt.Println(func_name, err.Error())
				continue
			}
			cdrs = append(cdrs, cdr)

		}

		/* TODO:
		content,err := json.Marshal(cdrs)
		if err == nil{
			fmt.Println(func_name,`SUCCESS!!!`)
			return Map{`statusCode`:200, `data`: content }

		}else{
			fmt.Println(func_name,err.Error())
		}*/
		ret := []Map{}
		for _, cdr := range cdrs {

			item := Map{
				`calldate`:    cdr.Calldate,
				`clid`:        cdr.Clid,
				`disposition`: cdr.Disposition,
				`dst`:         cdr.Dst,
				`src`:         cdr.Src,
				`uniqueid`:    cdr.Uniqueid,
				`billsec`:     cdr.Billsec,
				`duration`:    cdr.Duration,
			}
			ret = append(ret, item)
		}
		return Map{`statusCode`: 200, `data`: ret}

	}
	return Map{`statusCode`: 503}
}

func httpCDRResponse(username, q string, g string, pagenum int) Map {

	//func_name := `httpAddCDRResponse`

	switch g {
	case `RECENTS`:
		return httpCDRRecentsResponse(username, q)

	case `MISSED`:
		return httpCDRMissedResponse(username, q)

	default: //ALL
		return httpCDRAllResponse(username, q, pagenum)

	}

	//return Map{`statusCode`:503}
}

// func httpCDRAllResponse(username,q string,g string, skip,limit int) Map{
func httpCDRAllResponse(username, q string, p_pagenum int) Map {
	//func httpCDRResponse(data Map) Map{
	//ret := []map[string]string{}
	//exten_num := p_exten
	func_name := `httpAddCDRResponse`
	//fmt.Println(func_name,`(p_username=`,p_username,`p_exten=`,p_exten,`p_exten_to_username=`,p_exten_to_username)
	/*
		user := findUserByUsername(p_username)
		if user == nil || user.Roleid != `admin` {
			fmt.Println(`3897: UNAUTHORIZE username`,p_username)
			return Map{`statusCode`:500}
		}
	*/
	/*
		username,ok_username:= data["q"].(string)
		if !ok_username{
			fmt.Println(func_name, `Param username required`)
			return Map{`statusCode`:500}
		}
	*/

	extens, err := findAllUserExtensions(username)

	if err != nil {
		return Map{`statusCode`: 501}
	}

	if extens != nil && len(extens) > 0 {

		

		rows := []Map{}
	
		

		p := ``
		if len(extens) > 1 {

			for _, ext := range extens {
				p = fmt.Sprintf(`%s'%s',`, p, ext.Extension)
			}
			p = strings.TrimRight(p, `,`)

			p = fmt.Sprintf("src IN(%s) OR dst IN(%s)", p, p)

		} else {
			p = extens[0].Extension
			p = fmt.Sprintf("src = '%s' OR dst ='%s'", p, p)

		}

		//qry_count := fmt.Sprintf( "SELECT COUNT(*) as total FROM  cdr WHERE src IN('%s') OR dst IN('%s')",p,p)
		qry_count := fmt.Sprintf("SELECT COUNT(*) as total FROM  cdr WHERE %s", p)

		fmt.Println(`8479:`, func_name, `qry=`, qry_count)

		total := 0
		pagecount := 0
		//pagenum := 0
		rowsperpage := 100
		pagenum := p_pagenum

		row_total := db.QueryRow(qry_count)
		switch err := row_total.Scan(&total); err {
		case sql.ErrNoRows:
			fmt.Println("8471: ", func_name, "No rows were returned!")

			return Map{`statusCode`: 501}

		case nil:
			if total == 0 {
				return Map{`statusCode`: 200, `data`: rows, `total`: 0, `pagenum`: 0, `pagecount`: 0, `rowsperpage`: rowsperpage}
			}

			pagecount = total / rowsperpage
			fmt.Println(`8500.1:`, func_name, `pagecount=`, pagecount)
			if total > (pagecount * rowsperpage) {
				fmt.Println(`8500.2:`, func_name, `pagecount=`, pagecount)
				pagecount = pagecount + 1
				fmt.Println(`8500.3:`, func_name, `pagecount=`, pagecount)
			}
			fmt.Println(`8500.0:`, func_name, `pagecount=`, pagecount)

			if pagenum > pagecount {
				pagenum = pagecount
			}

		}

		
		qry := fmt.Sprintf(`SELECT B.uniqueid, calldate,clid,src,dst,duration,billsec,disposition,B.sequence,
			IFNULL((SELECT username FROM userextensions WHERE extension = A.src LIMIT 1),'') as src_username, 
			IFNULL((SELECT username FROM userextensions WHERE extension = A.dst LIMIT 1),'') as dst_username,
			IFNULL((SELECT name FROM userextensions ue LEFT JOIN users u ON ue.username = u.username WHERE extension = A.src LIMIT 1),'') as src_name, 
			IFNULL((SELECT name FROM userextensions ue LEFT JOIN users u ON ue.username = u.username WHERE extension = A.dst LIMIT 1),'') as dst_name,
			IFNULL((SELECT displayname FROM mycontacts WHERE username = '%s' AND number = A.src LIMIT 1),'') as src_contact, 
			IFNULL((SELECT displayname FROM mycontacts WHERE username = '%s' AND number = A.dst LIMIT 1),'') as dst_contact
			FROM   asteriskcdrdb.cdr A JOIN (SELECT uniqueid,MAX(sequence) as sequence
			FROM   asteriskcdrdb.cdr  WHERE  %s
			GROUP BY uniqueid ORDER BY calldate DESC LIMIT %v OFFSET %v) AS B ON (A.uniqueid = B.uniqueid AND A.sequence = B.sequence)`, username, username, p, rowsperpage, (pagenum-1)*rowsperpage)

		
			
		fmt.Println(string(colorYellow), `8521: CDR ALL`, func_name, `qry=`, qry)
		fmt.Println(string(colorReset))

		
		
		results, err := db.Query(qry)

		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(func_name, err.Error())
			return Map{`statusCode`: 502}
		}

		var cdrs []CDR

		for results.Next() {
			var cdr CDR
			// for each row, scan the result into our tag composite object
			//err = results.Scan(&cdr.Uniqueid,&cdr.Calldate,&cdr.Clid,&cdr.Dst,&cdr.Src,&cdr.Duration,&cdr.Billsec,&cdr.Disposition)
			err = results.Scan(
				&cdr.Uniqueid,
				&cdr.Calldate,
				&cdr.Clid,
				&cdr.Src,
				&cdr.Dst,
				&cdr.Duration,
				&cdr.Billsec,
				&cdr.Disposition,
				&cdr.Sequence,
				&cdr.Srcusername,
				&cdr.Dstusername,
				&cdr.Srcname,
				&cdr.Dstname,
				&cdr.Srccontact,
				&cdr.Dstcontact)
			if err != nil {
				//panic(err.Error()) // proper error handling instead of panic in your app
				fmt.Println(func_name, err.Error())
				continue
			}
			cdrs = append(cdrs, cdr)

		}

		
		
		for _, cdr := range cdrs {
			src_name := cdr.Srcname
			if src_name == `` {
				src_name = cdr.Srccontact
				if src_name != `` && strings.Contains(src_name, `<`) {
					src_name_arr := strings.Split(src_name, " ")
					if len(src_name_arr) > 1 {

						src_name = src_name_arr[0]

					}
				}

			}
			dst_name := cdr.Dstname
			if dst_name == `` {
				dst_name = cdr.Dstcontact
				if dst_name != `` && strings.Contains(dst_name, `<`) {
					dst_name_arr := strings.Split(dst_name, " ")
					if len(dst_name_arr) > 1 {

						dst_name = dst_name_arr[0]

					}
				}
			}
			item := Map{
				`calldate`:     cdr.Calldate,
				`clid`:         cdr.Clid,
				`disposition`:  cdr.Disposition,
				`dst`:          cdr.Dst,
				`src`:          cdr.Src,
				`uniqueid`:     cdr.Uniqueid,
				`billsec`:      cdr.Billsec,
				`duration`:     cdr.Duration,
				`src_username`: cdr.Srcusername,
				`dst_username`: cdr.Dstusername,
				`src_name`:     src_name,
				`dst_name`:     dst_name,
			}
			rows = append(rows, item)
		}
		return Map{`statusCode`: 200, `data`: rows, `total`: total, `pagenum`: pagenum, `pagecount`: pagecount, `rowsperpage`: rowsperpage}

	}
	return Map{`statusCode`: 503}
}

func httpCDRRecentsResponse_dep(username, q string) Map {
	
	func_name := `httpCDRRecentsResponse`
	

	extens, err := findAllUserExtensions(username)

	if err != nil {
		return Map{`statusCode`: 501}
	}

	if extens != nil && len(extens) > 0 {

		
		
		qry := ``

		
		

		ret := []Map{}

		if len(extens) == 0 {
			fmt.Println(`8834:`, func_name, `len(extens) is 0!`)
			return Map{`statusCode`: 200, `data`: ret}

		}

		p := ``

		if len(extens) > 1 {

			for _, ext := range extens {
				p = fmt.Sprintf(`%s'%s',`, p, ext.Extension)
			}
			p = strings.TrimRight(p, `,`)

			//qry = fmt.Sprintf( "SELECT uniqueid, calldate,clid,src,dst,duration,billsec,disposition  FROM   cdr  WHERE (src IN(%s) OR dst IN(%s)) AND  YEARWEEK(`calldate`, 1) = YEARWEEK(CURDATE(), 1) ORDER by calldate DESC LIMIT 100",p,p)
			p = fmt.Sprintf("(src IN(%s) OR dst IN(%s))", p, p)

		} else {

			p = extens[0].Extension

			//qry = fmt.Sprintf( "SELECT uniqueid, calldate,clid,src,dst,duration,billsec,disposition  FROM   cdr  WHERE (src = '%s' OR dst = '%s') AND  YEARWEEK(`calldate`, 1) = YEARWEEK(CURDATE(), 1) ORDER by calldate DESC LIMIT 100",p,p)
			p = fmt.Sprintf("(src = '%s' OR dst ='%s')", p, p)
		}

		
		
		qry = fmt.Sprintf(`SELECT B.uniqueid, calldate,clid,src,dst,duration,billsec,disposition,B.sequence,
			IFNULL((SELECT username FROM userextensions WHERE extension = A.src LIMIT 1),'') as src_username, 
			IFNULL((SELECT username FROM userextensions WHERE extension = A.dst LIMIT 1),'') as dst_username,
			IFNULL((SELECT name FROM userextensions ue LEFT JOIN users u ON ue.username = u.username WHERE extension = A.src LIMIT 1),'') as src_name, 
			IFNULL((SELECT name FROM userextensions ue LEFT JOIN users u ON ue.username = u.username WHERE extension = A.dst LIMIT 1),'') as dst_name,
			IFNULL((SELECT displayname FROM mycontacts WHERE username = '%s' AND number = A.src LIMIT 1),'') as src_contact, 
			IFNULL((SELECT displayname FROM mycontacts WHERE username = '%s' AND number = A.dst LIMIT 1),'') as dst_contact
			FROM   cdr A JOIN (SELECT uniqueid,MAX(sequence) as sequence
			FROM   cdr  WHERE  %s AND  YEARWEEK(calldate,1) = YEARWEEK(CURDATE(),1)
			GROUP BY uniqueid) AS B ON (A.uniqueid = B.uniqueid AND A.sequence = B.sequence)
			ORDER by calldate DESC LIMIT 100`, username, username, p)

			fmt.Println(`12140: `, func_name, `qry=`, qry)


		
			
		fmt.Println(string(colorYellow), `12140: `, func_name, `qry=`, qry)
		fmt.Println(string(colorReset))

		
		
		results, err := db.Query(qry)

		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(func_name, err.Error())
			return Map{`statusCode`: 502}
		}

		var cdrs []CDR

		for results.Next() {
			var cdr CDR
			// for each row, scan the result into our tag composite object
			err = results.Scan(
				&cdr.Uniqueid,
				&cdr.Calldate,
				&cdr.Clid,
				&cdr.Src,
				&cdr.Dst,
				&cdr.Duration,
				&cdr.Billsec,
				&cdr.Disposition,
				&cdr.Sequence,
				&cdr.Srcusername,
				&cdr.Dstusername,
				&cdr.Srcname,
				&cdr.Dstname,
				&cdr.Srccontact,
				&cdr.Dstcontact)

			if err != nil {
				//panic(err.Error()) // proper error handling instead of panic in your app
				fmt.Println(func_name, err.Error())
				continue
			}
			cdrs = append(cdrs, cdr)

		}

		
		

		for _, cdr := range cdrs {
			src_name := cdr.Srcname
			if src_name == `` {
				src_name = cdr.Srccontact
				if src_name != `` && strings.Contains(src_name, `<`) {
					src_name_arr := strings.Split(src_name, " ")
					if len(src_name_arr) > 1 {

						src_name = src_name_arr[0]

					}
				}
			}
			dst_name := cdr.Dstname
			if dst_name == `` {
				dst_name = cdr.Dstcontact
				if dst_name != `` && strings.Contains(dst_name, `<`) {
					dst_name_arr := strings.Split(dst_name, " ")
					if len(dst_name_arr) > 1 {

						dst_name = dst_name_arr[0]

					}
				}
			}

			item := Map{
				`calldate`:     cdr.Calldate,
				`clid`:         cdr.Clid,
				`disposition`:  cdr.Disposition,
				`dst`:          cdr.Dst,
				`src`:          cdr.Src,
				`uniqueid`:     cdr.Uniqueid,
				`billsec`:      cdr.Billsec,
				`duration`:     cdr.Duration,
				`src_username`: cdr.Srcusername,
				`dst_username`: cdr.Dstusername,
				`src_name`:     src_name,
				`dst_name`:     dst_name,
				//`src_name`: cdr.Srcname,
				//`dst_name`: cdr.Dstname,
				//`src_contact`: cdr.Srccontact,
				//`dst_contact`: cdr.Dstcontact
			}
			ret = append(ret, item)
		}
		return Map{`statusCode`: 200, `data`: ret}

	}
	return Map{`statusCode`: 503}
}

func httpCDRRecentsResponse(username, q string) Map {
	func_name := `httpCDRRecentsResponse`

	// Step 1: Get all extensions for this user
	extens, err := findAllUserExtensions(username)
	if err != nil {
		fmt.Println(func_name, "Error in findAllUserExtensions:", err.Error())
		return Map{`statusCode`: 501}
	}

	fmt.Println(func_name, "Found extensions:", extens)

	if extens == nil || len(extens) == 0 {
		fmt.Println(func_name, "No extensions found for user:", username)
		return Map{`statusCode`: 200, `data`: []Map{}}
	}

	// Step 2: Build dynamic WHERE clause for src/dst matching
	var p string
	if len(extens) > 1 {
		for _, ext := range extens {
			p += fmt.Sprintf("'%s',", ext.Extension)
		}
		p = strings.TrimRight(p, ",")
		p = fmt.Sprintf("(src IN(%s) OR dst IN(%s))", p, p)
	} else {
		ext := extens[0].Extension
		p = fmt.Sprintf("(src = '%s' OR dst = '%s')", ext, ext)
	}

	// Step 3: Build SQL query with safer date filter (last 7 days instead of YEARWEEK)
	qry := fmt.Sprintf(`
SELECT 
  B.uniqueid, calldate, clid, src, dst, duration, billsec, disposition, B.sequence,
  IFNULL((SELECT username FROM userextensions WHERE extension = A.src LIMIT 1),'') AS src_username, 
  IFNULL((SELECT username FROM userextensions WHERE extension = A.dst LIMIT 1),'') AS dst_username,
  IFNULL((SELECT name FROM userextensions ue LEFT JOIN users u ON ue.username = u.username WHERE extension = A.src LIMIT 1),'') AS src_name, 
  IFNULL((SELECT name FROM userextensions ue LEFT JOIN users u ON ue.username = u.username WHERE extension = A.dst LIMIT 1),'') AS dst_name,
  IFNULL((SELECT displayname FROM mycontacts WHERE username = '%s' AND number = A.src LIMIT 1),'') AS src_contact, 
  IFNULL((SELECT displayname FROM mycontacts WHERE username = '%s' AND number = A.dst LIMIT 1),'') AS dst_contact
FROM asteriskcdrdb.cdr A 
JOIN (
  SELECT uniqueid, MAX(sequence) AS sequence
  FROM asteriskcdrdb.cdr  
  WHERE %s AND calldate >= NOW() - INTERVAL 7 DAY
  GROUP BY uniqueid
) AS B 
ON A.uniqueid = B.uniqueid AND A.sequence = B.sequence
ORDER BY calldate DESC 
LIMIT 100
`, username, username, p)

	fmt.Println(string(colorYellow), func_name, "SQL Query:\n"+qry, string(colorReset))

	// Step 4: Execute the query
	results, err := db.Query(qry)
	if err != nil {
		fmt.Println(func_name, "Query error:", err.Error())
		return Map{`statusCode`: 502}
	}
	defer results.Close()

	// Step 5: Parse query result rows
	var cdrs []CDR
	rowFound := false

	for results.Next() {
		rowFound = true
		var cdr CDR
		err := results.Scan(
			&cdr.Uniqueid,
			&cdr.Calldate,
			&cdr.Clid,
			&cdr.Src,
			&cdr.Dst,
			&cdr.Duration,
			&cdr.Billsec,
			&cdr.Disposition,
			&cdr.Sequence,
			&cdr.Srcusername,
			&cdr.Dstusername,
			&cdr.Srcname,
			&cdr.Dstname,
			&cdr.Srccontact,
			&cdr.Dstcontact)

		if err != nil {
			fmt.Println(func_name, "Scan error:", err.Error())
			continue
		}
		cdrs = append(cdrs, cdr)
	}

	if !rowFound {
		fmt.Println(func_name, "No recent call records found for user:", username)
		return Map{`statusCode`: 200, `data`: []Map{}}
	}

	// Step 6: Convert to response maps
	var ret []Map
	for _, cdr := range cdrs {
		src_name := cdr.Srcname
		if src_name == "" {
			src_name = cdr.Srccontact
			if strings.Contains(src_name, "<") {
				parts := strings.Split(src_name, " ")
				if len(parts) > 1 {
					src_name = parts[0]
				}
			}
		}

		dst_name := cdr.Dstname
		if dst_name == "" {
			dst_name = cdr.Dstcontact
			if strings.Contains(dst_name, "<") {
				parts := strings.Split(dst_name, " ")
				if len(parts) > 1 {
					dst_name = parts[0]
				}
			}
		}

		ret = append(ret, Map{
			"calldate":     cdr.Calldate,
			"clid":         cdr.Clid,
			"disposition":  cdr.Disposition,
			"dst":          cdr.Dst,
			"src":          cdr.Src,
			"uniqueid":     cdr.Uniqueid,
			"billsec":      cdr.Billsec,
			"duration":     cdr.Duration,
			"src_username": cdr.Srcusername,
			"dst_username": cdr.Dstusername,
			"src_name":     src_name,
			"dst_name":     dst_name,
		})
	}

	return Map{`statusCode`: 200, `data`: ret}
}

func httpCDRMissedResponse(username, q string) Map {
	
	func_name := `httpCDRRecentsResponse`
	
	

	extens, err := findAllUserExtensions(username)

	if err != nil {
		return Map{`statusCode`: 501}
	}

	if extens != nil && len(extens) > 0 {

		
		

		qry := ``

		
		

		ret := []Map{}

		if len(extens) == 0 {
			fmt.Println(`8834:`, func_name, `len(extens) is 0!`)
			return Map{`statusCode`: 200, `data`: ret}

		}

		p := ``
		if len(extens) > 1 {

			for _, ext := range extens {
				p = fmt.Sprintf(`%s'%s',`, p, ext.Extension)
			}
			p = strings.TrimRight(p, `,`)

			p = fmt.Sprintf("(src IN(%s) OR dst IN(%s))", p, p)

		} else {

			p = extens[0].Extension
			//qry = fmt.Sprintf( "SELECT  uniqueid, MAX(calldate) AS calldate ,clid,src,dst,duration,billsec,disposition FROM cdr WHERE (src = '%s' OR dst = '%s') AND disposition <> 'ANSWERED' AND  YEARWEEK(`calldate`, 1) = YEARWEEK(CURDATE(), 1) ORDER by calldate DESC LIMIT 100",p,p)
			p = fmt.Sprintf("(src = '%s' OR dst ='%s')", p, p)
		}

		
		
		qry = fmt.Sprintf(`SELECT B.uniqueid, calldate,clid,src,dst,duration,billsec,disposition,B.sequence,
			IFNULL((SELECT username FROM userextensions WHERE extension = A.src LIMIT 1),'') as src_username, 
			IFNULL((SELECT username FROM userextensions WHERE extension = A.dst LIMIT 1),'') as dst_username,
			IFNULL((SELECT name FROM userextensions ue LEFT JOIN users u ON ue.username = u.username WHERE extension = A.src LIMIT 1),'') as src_name, 
			IFNULL((SELECT name FROM userextensions ue LEFT JOIN users u ON ue.username = u.username WHERE extension = A.dst LIMIT 1),'') as dst_name,
			IFNULL((SELECT displayname FROM mycontacts WHERE username = '%s' AND number = A.src LIMIT 1),'') as src_contact, 
			IFNULL((SELECT displayname FROM mycontacts WHERE username = '%s' AND number = A.dst LIMIT 1),'') as dst_contact
			FROM   asteriskcdrdb.cdr A JOIN (SELECT uniqueid,MAX(sequence) as sequence
			FROM   asteriskcdrdb.cdr  WHERE  %s AND  YEARWEEK(calldate,1) = YEARWEEK(CURDATE(),1)
			GROUP BY uniqueid) AS B ON (A.uniqueid = B.uniqueid AND A.sequence = B.sequence) HAVING disposition <> 'ANSWERED'
			ORDER by calldate DESC LIMIT 100`, username, username, p)

	
			
		fmt.Println(string(colorYellow), `12329: MISSED `, func_name, `qry=`, qry)
		fmt.Println(string(colorReset))

		results, err := db.Query(qry)

		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(func_name, err.Error())
			return Map{`statusCode`: 502}
		}

		var cdrs []CDR

		for results.Next() {
			var cdr CDR
		
			
			err = results.Scan(
				&cdr.Uniqueid,
				&cdr.Calldate,
				&cdr.Clid,
				&cdr.Src,
				&cdr.Dst,
				&cdr.Duration,
				&cdr.Billsec,
				&cdr.Disposition,
				&cdr.Sequence,
				&cdr.Srcusername,
				&cdr.Dstusername,
				&cdr.Srcname,
				&cdr.Dstname,
				&cdr.Srccontact,
				&cdr.Dstcontact)
			if err != nil {
				//panic(err.Error()) // proper error handling instead of panic in your app
				fmt.Println(func_name, err.Error())
				continue
			}
			cdrs = append(cdrs, cdr)

		}

		
		

		for _, cdr := range cdrs {
			src_name := cdr.Srcname
			if src_name == `` {
				src_name = cdr.Srccontact
				if src_name != `` && strings.Contains(src_name, `<`) {
					src_name_arr := strings.Split(src_name, " ")
					if len(src_name_arr) > 1 {

						src_name = src_name_arr[0]

					}
				}
			}
			dst_name := cdr.Dstname
			if dst_name == `` {
				dst_name = cdr.Dstcontact
				if dst_name != `` && strings.Contains(dst_name, `<`) {
					dst_name_arr := strings.Split(dst_name, " ")
					if len(dst_name_arr) > 1 {

						dst_name = dst_name_arr[0]

					}
				}
			}
			item := Map{
				`calldate`:     cdr.Calldate,
				`clid`:         cdr.Clid,
				`disposition`:  cdr.Disposition,
				`dst`:          cdr.Dst,
				`src`:          cdr.Src,
				`uniqueid`:     cdr.Uniqueid,
				`billsec`:      cdr.Billsec,
				`duration`:     cdr.Duration,
				`src_username`: cdr.Srcusername,
				`dst_username`: cdr.Dstusername,
				`src_name`:     src_name,
				`dst_name`:     dst_name,
			}
			ret = append(ret, item)
		}
		return Map{`statusCode`: 200, `data`: ret}

	}
	return Map{`statusCode`: 503}
}

func httpInboxResponse(p_username string) Map {
	//ret := []map[string]string{}
	//exten_num := p_exten
	func_name := `httpInboxResponse`
	fmt.Println("START ", func_name, "username=", p_username)
	
	

	ret := []Map{}

	
	

	
	qry := ``
	
	
	qry = fmt.Sprintf("SELECT id, created, username, `from`,`to`,fromname,toname, roomid, message FROM inbox WHERE id = '%s'  ORDER BY created DESC LIMIT 100", p_username) //extens[0].Extension)
	//}

	fmt.Println(`8959:`, func_name, `qry=`, qry)

	results, err := db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`4342: `, func_name, err.Error())
		return Map{`statusCode`: 502}
	}

	var inboxes []INBOX

	for results.Next() {
		var inbox INBOX
		// for each row, scan the result into our tag composite object
		err = results.Scan(&inbox.Id, &inbox.Created, &inbox.Username, &inbox.From, &inbox.To, &inbox.Fromname, &inbox.Toname, &inbox.Roomid, &inbox.Message)
		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(`4354`, func_name, err.Error())
			continue
		}
		inboxes = append(inboxes, inbox)

	}

	
	for _, msg := range inboxes {

		item := Map{
			`id`:       msg.Id,
			`created`:  msg.Created,
			`username`: msg.Username,
			`from`:     msg.From,
			`to`:       msg.To,
			`fromname`: msg.Fromname,
			`toname`:   msg.Toname,
			`roomid`:   msg.Roomid,
			`message`:  msg.Message,
		}
		ret = append(ret, item)
	}
	return Map{`statusCode`: 200, `data`: ret}

	
	
}

func httpInboxMessagesResponse(p_username string, inboxid string, data Map) Map {
	
	func_name := `httpInboxMessagesResponse`
	
	fmt.Println("START:", func_name)
	fmt.Println("data:", data)

	
	

	q, ok_q := data["q"].(string)

	if !ok_q {
		q = ""
	}

	skip, ok_skip := data["skip"].(float64)

	if !ok_skip {
		skip = 0
	}

	fmt.Println("**** skip:", skip)

	limit := 25

	//var y int = int(x)
	offset := int(skip)

	
	

	qry := ``
	
	
	if len(q) > 0 {
		qry = fmt.Sprintf("SELECT id, inboxid, created, username, `from`,`to`, roomid, message FROM messages WHERE inboxid = '%s'  ORDER BY created DESC LIMIT %v OFFSET %v", inboxid, limit, offset)

	} else {
		qry = fmt.Sprintf("SELECT id, inboxid, created, username, `from`,`to`, roomid, message FROM messages WHERE inboxid = '%s'  ORDER BY created DESC LIMIT %v OFFSET %v", inboxid, limit, offset)

	}

	
	
	fmt.Println(`9094:`, func_name, `qry=`, qry)

	rows, err := db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`12814:`, func_name, err.Error())
		return Map{`statusCode`: 502}
	}

	defer rows.Close()

	var messages []MSG

	for rows.Next() {
		var msg MSG
		// for each row, scan the result into our tag composite object
		err = rows.Scan(&msg.Id, &msg.Inboxid, &msg.Created, &msg.Username, &msg.From, &msg.To, &msg.Roomid, &msg.Message)
		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(`4484:`, func_name, err.Error())
			return Map{`statusCode`: 503}
		}
		messages = append(messages, msg)

	}

	
	
	ret := []Map{}
	for _, msg := range messages {

		item := Map{
			`id`:       msg.Id,
			`inboxid`:  msg.Inboxid,
			`created`:  msg.Created,
			`username`: msg.Username,
			`from`:     msg.From,
			`to`:       msg.To,
			`roomid`:   msg.Roomid,
			`message`:  msg.Message,
		}
		ret = append(ret, item)
	}
	return Map{`statusCode`: 200, `data`: ret}
}

func httpTypeAheadContactsResponse(q, p_username string) Map {
	
	func_name := `httpInboxMessagesResponse`
	fmt.Println(`8282: START`, func_name)
	
	results, err := dbTypeAheadContacts(q, p_username)
	if err != nil {
		return Map{`statusCode`: 500}

	}
	
	
	return Map{`statusCode`: 200, `data`: results}
}

func httpTypeAheadExtensionsResponse(q, p_username string) Map {
	
	func_name := `httpTypeAheadExtensionsResponse`
	fmt.Println(`8282: START`, func_name)
	
	

	results, err := dbTypeAheadExtensions(q, p_username)
	if err != nil {
		return Map{`statusCode`: 500}

	}
	
	
	return Map{`statusCode`: 200, `data`: results}
}




func dbUpsertExtension_dep(exten Extension) error {
	fmt.Println("START dbUpsertExtension")
	
	
	qry := fmt.Sprintf("REPLACE INTO extensions(number, displayname, status, contacts,source, secret, public, ver) VALUES('%s','%s','%s','%s','%s','%s','%s','%s')",
		exten.Number,
		exten.Displayname,
		exten.Status,
		exten.Contacts,
		exten.Source,
		exten.Secret,
		exten.Public,
		
		exten.Ver)

	
		insert, err := db.Query(qry)

	
		if err != nil {
		//panic(err.Error())
		fmt.Println("FAILED %s", qry)
		return err
	}
	defer insert.Close()

	
	
	return nil //Map{`statusCode`:503}
}

func dbUpsertExtension_dep2(exten Extension) error {
    fmt.Println("START dbUpsertExtension")

    const q = `
        INSERT INTO extensions
            (number, displayname, status, contacts, secret, ver)
        VALUES
            (?,?, ?, ?, ?, ?)
        ON DUPLICATE KEY UPDATE
            displayname = VALUES(displayname),
            status      = VALUES(status),
            contacts    = VALUES(contacts),
            
            secret      = VALUES(secret),
            
            ver         = VALUES(ver)
    `

    res, err := db.Exec(q,
        exten.Number,
        exten.Displayname,
        exten.Status,
        exten.Contacts,
        exten.Secret,
        exten.Ver,
    )
    if err != nil {
        return fmt.Errorf("upsert extension %q: %w", exten.Number, err)
    }

    if n, _ := res.RowsAffected(); n == 0 {
        // No-op (values identical). Optional: log if you expect a change.
        // fmt.Printf("No change for extension %s\n", exten.Number)
    }
    return nil
}

func dbUpsertExtension(ctx context.Context, exten Extension) error {
    const q = `
        INSERT INTO extensions
            (number, displayname, status, contacts, ver)
        VALUES
            (?, ?, ?, ?, ?)
        ON DUPLICATE KEY UPDATE
            displayname = VALUES(displayname),
            status      = VALUES(status),
            contacts    = VALUES(contacts),
            
            ver         = VALUES(ver)
    `
    res, err := db.ExecContext(ctx, q,
        exten.Number,
        exten.Displayname,
        exten.Status,
        exten.Contacts,
        
        exten.Ver,
    )
    if err != nil {
        return fmt.Errorf("upsert extension %q: %w", exten.Number, err)
    }
    if n, _ := res.RowsAffected(); n == 0 {
        // no change
    }
    return nil
}




func dbUpsertExtensionAdmin(ctx context.Context, exten Extension) error {
	fmt.Printf("14466: dbUpsertExtensionAdmin() exten: %+v\n", exten)
    const q = `
        INSERT INTO extensions
            (number, displayname, secret, webrtc, media_encryption, direct_media,
             public, canblock, canunlist, canaddlist, candial)
        VALUES
            (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
        ON DUPLICATE KEY UPDATE
            displayname       = VALUES(displayname),
            secret            = VALUES(secret),
            webrtc            = VALUES(webrtc),
            media_encryption  = VALUES(media_encryption),
            direct_media      = VALUES(direct_media),
            public            = VALUES(public),
            canblock          = VALUES(canblock),
            canunlist         = VALUES(canunlist),
            canaddlist        = VALUES(canaddlist),
            candial           = VALUES(candial)
    `
    _, err := db.ExecContext(ctx, q,
        exten.Number,
        exten.Displayname,
        exten.Secret,
        exten.Webrtc,           // "yes"/"no" per your schema
        exten.Mediaencryption,  // e.g. "dtls"
        exten.Directmedia,      // "yes"/"no"
        exten.Public,           // "yes"/"no"
        exten.Canblock,         // "yes"/"no"
        exten.Canunlist,        // "yes"/"no"
        exten.Canaddlist,       // "yes"/"no"
        exten.Candial,          // "yes"/"no"
    )
    if err != nil {
        return fmt.Errorf("upsert extension %q: %w", exten.Number, err)
    }
    return nil
}



func dbUpsertExtensionConfig(exten Extension) error {
	fmt.Println("START dbUpsertExtensionConfig")
	
	

	qry := fmt.Sprintf("REPLACE INTO extensions(extension, displayname, email, secret) VALUES('%s','%s','%s','%s','%s')",
		exten.Number,
		exten.Displayname,
		exten.Email,
		exten.Secret,
		)

		
	insert, err := db.Query(qry)

	// if there is an error inserting, handle it
	if err != nil {
		//panic(err.Error())
		fmt.Println("FAILED %s", qry)
		return err
	}
	defer insert.Close()

	
	
	return nil //Map{`statusCode`:503}
}

func dbUpdateUsername(email, username string) error {
	
	

	qry := fmt.Sprintf("UPDATE users SET username = '%s'  WHERE email = '%s'", username, email)

	
	

	update, err := db.Query(qry)

	
	if err != nil {
	
		fmt.Println("FAILED %s", qry)
	
	}

	defer update.Close()

	return nil //Map{`statusCode`:503}
}

func dbUpdateUserMissedCall(str_exten, calleridnum, misseduniqueid string) error {
	
	
	results := findAllUsersWithExtension(str_exten)

	if results != nil || len(results) == 0 {
		return nil
	}

	for _, u := range results {

		qry := fmt.Sprintf("UPDATE users SET missed = (missed + 1), misseduniqueid = '%s'  WHERE username = '%s' AND misseduniqueid <> '%s'", misseduniqueid, u.Username, misseduniqueid)

		//fmt.Println(string(colorYellow),qry)
		//fmt.Println(string(colorReset))

		update, err := db.Query(qry)

		// if there is an error inserting, handle it
		if err != nil {
			//panic(err.Error())
			fmt.Println("FAILED %s", qry)
			//return err
		}

		defer update.Close()

	}

	return nil //Map{`statusCode`:503}
}

func dbUpdateUserUnreadMessages(p_username string) error {
	
	

	qry := fmt.Sprintf("UPDATE users SET unreadmessages = (unreadmessages + 1)  WHERE username = '%s'", p_username)

	
	
	update, err := db.Query(qry)

	
	if err != nil {
	
		fmt.Println("FAILED %s", qry)
	
	}

	defer update.Close()

	

	return nil 
}

func dbDeleteExtension(exten_num string) error {
	

	qry := fmt.Sprintf("DELETE FROM extensions WHERE number = '%s'", exten_num)

	del, err := db.Query(qry)

	// if there is an error inserting, handle it
	if err != nil {
		//panic(err.Error())
		fmt.Println("FAILED %s", qry)
		return err
	}
	defer del.Close()

	
	return nil //Map{`statusCode`:503}
}

func dbMyContactsInitIfNeeded() error {
	//ret := []map[string]string{}
	//exten_num := p_exten
	func_name := `dbMyContactsInitIfNeeded`
	fmt.Println(`13255: START`, func_name)

	// Open up our database connection.

	

	qry := "SELECT number,displayname FROM extensions WHERE public = '1'"

	results, err := db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(func_name, err.Error())
		return err
	}

	//var cdrs []CDR

	for results.Next() {

		// for each row, scan the result into our tag composite object
		var number string
		var displayname string

		err = results.Scan(
			&number,
			&displayname)

		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(func_name, err.Error())
			continue
		}
		dbMyContactsSyncByExtenNun(number, displayname)

	}

	defer results.Close()

	return nil //Map{`statusCode`:503}
}

func dbMyContactsSyncByExtenNun(exten_num, displayname string) error {
	//ret := []map[string]string{}
	//exten_num := p_exten
	func_name := `dbSetPublicExtension`
	fmt.Println(`13255: START`, func_name)
	
	

	qry := fmt.Sprintf("SELECT username FROM users WHERE username NOT IN(SELECT username FROM mycontacts WHERE number = '%s')", exten_num)

	results, err := db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(func_name, err.Error())
		return err
	}

	//var cdrs []CDR

	for results.Next() {
		var username string
		// for each row, scan the result into our tag composite object
		err = results.Scan(
			&username)

		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(func_name, err.Error())
			continue
		}
		//cdrs = append(cdrs, cdr)
		qry = fmt.Sprintf("INSERT mycontacts(username,number,displayname) VALUES('%s','%s','%s')", username, exten_num, displayname)
		fmt.Println(`13255.2:`, func_name, "qry=", qry)

		res, err := db.Query(qry)

		// if there is an error inserting, handle it
		if err != nil {
			//panic(err.Error())
			fmt.Println("FAILED %s", qry)
			return err
		}
		defer res.Close()

	}

	defer results.Close()

	
	return nil //Map{`statusCode`:503}
}

func dbExtensionSetPublic(exten_num, public_val string) error {
	//ret := []map[string]string{}
	//exten_num := p_exten
	func_name := `dbSetPublicExtension`
	fmt.Println(`13255: START`, func_name)
	
	qry := fmt.Sprintf("SELECT number,displayname FROM extensions WHERE number = '%s'", exten_num)

	results, err := db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(func_name, err.Error())
		return err
	}

	//var cdrs []CDR

	var number string
	var displayname string
	for results.Next() {

		// for each row, scan the result into our tag composite object
		err = results.Scan(
			&number,
			&displayname)

		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(func_name, err.Error())
			continue
		}

	}

	if number != exten_num {
		return errors.New(`13333: number != exten_num !!!`)
	}

	if len(displayname) == 0 {
		displayname = number
	}

	qry = fmt.Sprintf("UPDATE extensions SET public = '%s' WHERE number = '%s'", public_val, exten_num)
	fmt.Println(`13255.2:`, func_name, "qry=", qry)

	res, err := db.Query(qry)

	// if there is an error inserting, handle it
	if err != nil {
		//panic(err.Error())
		fmt.Println("FAILED %s", qry)
		return err
	}
	defer res.Close()

	//propagate changes
	dbMyContactsSyncByExtenNun(exten_num, displayname)
	
	return nil //Map{`statusCode`:503}
}

func dbMyContactSetDisplayname(exten_num, displayname_val string) error {
	
	func_name := `dbMyContactSetDisplayname`
	fmt.Println(`13255: START`, func_name)
	
	
	qry := fmt.Sprintf("UPDATE mycontacts SET displayname = '%s' WHERE number = '%s'", displayname_val, exten_num)
	fmt.Println(`13255.2:`, func_name, "qry=", qry)

	res, err := db.Query(qry)

	// if there is an error inserting, handle it
	if err != nil {
		//panic(err.Error())
		fmt.Println("FAILED %s", qry)
		return err
	}
	defer res.Close()

	
	return nil //Map{`statusCode`:503}
}

func dbDeleteExtensionsWithOldVersion_dep(ver string) error {
	//ret := []map[string]string{}
	//exten_num := p_exten
	//func_name := `dbDeleteExtensionsWithOldVersion`
	//fmt.Println(func_name,`(p_username=`,p_username,`p_exten=`,p_exten,`p_exten_to_username=`,p_exten_to_username)
	/*
		user := findUserByUsername(p_username)
		if user == nil || user.Roleid != `admin` {
			fmt.Println(`3897: UNAUTHORIZE username`,p_username)
			return Map{`statusCode`:500}
		}
	*/

	//extens := findAllUserExtensions(p_username)

	//if extens != nil && len(extens) > 0{

	// Open up our database connection.

	

	qry := fmt.Sprintf("DELETE FROM extensions WHERE ver <> '%s'", ver)

	del, err := db.Query(qry)

	// if there is an error inserting, handle it
	if err != nil {
		//panic(err.Error())
		fmt.Println("FAILED %s", qry)
		return err
	}
	defer del.Close()

	return nil //Map{`statusCode`:503}
}

func dbDeleteExtensionsWithOldVersion(ctx context.Context, ver string) error {
    const q = `DELETE FROM extensions WHERE ver <> ?`
    res, err := db.ExecContext(ctx, q, ver)
    if err != nil {
        return fmt.Errorf("delete stale extensions (ver != %q): %w", ver, err)
    }
    if n, _ := res.RowsAffected(); n > 0 {
        fmt.Printf("deleted %d stale extensions (ver != %s)\n", n, ver)
    }
    return nil
}



func dbLoadExtensions() error {

	
	func_name := `dbLoadExtensions`
	
	

	qry := fmt.Sprintf("SELECT number, displayname, status, contacts,source, secret,public, ver FROM extensions")

	//fmt.Println(`9930:`,func_name,`qry=`,qry)

	results, err := db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`13359.2: `, func_name, err.Error())
		panic(err.Error())
		//return err //Map{`statusCode`:502}
	}

	//var inboxes []INBOX
	var extens []Extension
	for results.Next() {
		//var inbox INBOX
		var exten Extension
		// for each row, scan the result into our tag composite object
		err = results.Scan(&exten.Number, &exten.Displayname, &exten.Status, &exten.Contacts, &exten.Source, &exten.Secret, &exten.Public, &exten.Ver)
		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(`13359.3`, func_name, err.Error())
			panic(err.Error())
			//continue
		}
		extens = append(extens, exten)

	}

	extensions = extens

	return nil

}

func dbLoadRegistered() error {

	//ret := []map[string]string{}
	//exten_num := p_exten
	func_name := `dbLoadRegistered`
	
	qry := "SELECT exten, username, device_token, device_type, remote_addr,contact_name,contact_host,contact_transport,register_status,awake_status,date,ver FROM registered"

	//fmt.Println(`10036:`,func_name,`qry=`,qry)

	results, err := db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`13449.2: `, func_name, err.Error())
		panic(err.Error())
		//return err //Map{`statusCode`:502}
	}

	//var inboxes []INBOX
	var regs []WSRegister
	count := 0
	for results.Next() {
		count++
		
		var reg WSRegister
		
		err = results.Scan(&reg.Exten, &reg.Username, &reg.Device_Token, &reg.Device_Type, &reg.Remote_Addr, &reg.Contact_Name, &reg.Contact_Host, &reg.Contact_Transport, &reg.Register_Status, &reg.Awake_Status, &reg.Date, &reg.Ver)

		if err != nil {
		
			fmt.Println(`13449.3`, func_name, err.Error())
			panic(err.Error())
		}
		regs = append(regs, reg)

	}

	fmt.Println(`****************`)
	fmt.Println(`count=`, count)
	fmt.Println(`regs=`, regs)
	fmt.Println(`****************`)
	registered = regs

	return nil

}

func dbSaveRegistered() {
	
	
	ver := time.Now().Format(time.RFC3339)
	for _, reg := range registered {

		//qry := fmt.Sprintf("SELECT exten, username, device_token, device_type, remote_addr,contact_name,contact_host,contact_transport,register_status,awake_status,date FROM registered")

		qry := fmt.Sprintf(`REPLACE INTO registered(
			exten, 
			username,
			device_token, 
			device_type, 
			remote_addr,
			contact_name,
			contact_host,
			contact_transport,
			register_status,
			awake_status,
			date,
			ver) VALUES('%s','%s','%s','%s','%s','%s','%s','%s','%v','%s','%s','%s')`,
			reg.Exten,
			reg.Username,
			reg.Device_Token,
			reg.Device_Type,
			reg.Remote_Addr,
			reg.Contact_Name,
			reg.Contact_Host,
			reg.Contact_Transport,
			reg.Register_Status,
			reg.Awake_Status,
			reg.Date,
			ver)

		//fmt.Println(string(colorRed), "qry=", qry)
		//fmt.Println(string(colorReset))
		insert, err := db.Query(qry)

		// if there is an error inserting, handle it
		if err != nil {
			//panic(err.Error())
			fmt.Println("FAILED %s", qry)
			return
		}
		defer insert.Close()

	}
	qry := fmt.Sprintf("DELETE FROM registered WHERE ver <> '%s'", ver)

	del, err := db.Query(qry)

	// if there is an error inserting, handle it
	if err != nil {
		//panic(err.Error())
		fmt.Println("FAILED %s", qry)
		return //err
	}
	defer del.Close()

	
	return //nil //Map{`statusCode`:503}
}

func dbLoadUsers() error {

	
	func_name := `dbLoadUsers`

	
	
	qry := fmt.Sprintf("SELECT username, email, name, roleid,status, sessionid,missed, misseduniqueid,unreadmessages,homepageurl,ver FROM users")

	
	
	results, err := db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`13753.2: `, func_name, err.Error())
		return err //Map{`statusCode`:502}
	}

	users = []User{}
	for results.Next() {

		var user User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.Username, &user.Email, &user.Name, &user.Roleid, &user.Status, &user.Sessionid, &user.Missed, &user.Misseduniqueid, &user.Unreadmessages, &user.Homepageurl, &user.Ver)
		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(`13753.3`, func_name, err.Error())
			continue
		}
		users = append(users, user)

	}

	return nil

}

func dbUpsertUser(user User, func_caller string) error {
	
	func_name := `dbUpsertUsert`
	fmt.Println("START:", func_name, "user=", user, "called by", func_caller)
	
	
	
	
	qry := fmt.Sprintf("REPLACE INTO users(username, email, name,firstname,lastname,phonenumber, photo, roleid,status, sessionid, missed,misseduniqueid,unreadmessages, ver) VALUES('%s','%s','%s','%s','%s','%s','%s','%s', %v,'%s',%v,'%s',%v,'%s')",
		user.Username,
		user.Email,
		user.Name,
		user.Firstname,
		user.Lastname,
		user.Phonenumber,
		user.Photo,
		user.Roleid,
		user.Status,
		user.Sessionid,
		user.Missed,
		user.Misseduniqueid,
		user.Unreadmessages,
		user.Ver)

	fmt.Println(string(colorRed), "qry=", qry)
	fmt.Println(string(colorReset))
	insert, err := db.Query(qry)

	// if there is an error inserting, handle it
	if err != nil {
		//panic(err.Error())
		fmt.Println("FAILED %s", qry)
		return err
	}
	defer insert.Close()

	
	return nil //Map{`statusCode`:503}
}

func dbDeleteUser(username string) error {

	//func_name := `dbDeleteUser`

	

	qry := fmt.Sprintf("DELETE FROM userextensions WHERE username = '%s'", username)
	del, err := db.Query(qry)
	fmt.Println(`qry=`, qry)
	if err != nil {
		fmt.Println(`FAILED delete userextensions`, username, `qry=`, qry)

	} else {
		fmt.Println(`OK delete userextensions`, username)
	}

	qry = fmt.Sprintf("DELETE FROM blockedusers WHERE blockedby = '%s' OR blockeduser = '%s'", username, username)
	fmt.Println(`qry=`, qry)
	del, err = db.Query(qry)
	if err != nil {
		fmt.Println(`FAILED delete userextensions`, username, username, `qry=`, qry)

	} else {
		fmt.Println(`OK delete userextensions`, username)
	}

	qry = fmt.Sprintf("DELETE FROM messages WHERE `from` = '%s' OR `to` = '%s'", username, username)
	fmt.Println(`qry=`, qry)
	del, err = db.Query(qry)
	if err != nil {
		fmt.Println(`FAILED delete messages`, username)

	} else {
		fmt.Println(`OK delete messages`, username)
	}

	qry = fmt.Sprintf("DELETE FROM inbox WHERE `from` = '%s' OR `to` = '%s'", username, username)
	fmt.Println(`qry=`, qry)
	del, err = db.Query(qry)
	if err != nil {
		fmt.Println(`FAILED delete inbox`, username)

	} else {
		fmt.Println(`OK delete inbox`, username)
	}

	qry = fmt.Sprintf("DELETE FROM users WHERE username = '%s'", username)
	fmt.Println(`qry=`, qry)
	del, err = db.Query(qry)
	if err != nil {
		fmt.Println("FAILED delete user", username)

	} else {
		fmt.Println(`OK delete user`, username)
	}

	defer del.Close()

	return nil //Map{`statusCode`:503}
}

func dbFindUserExtensions(p_username string) ([]string, error) {

	extens := []string{}
	func_name := `dbDeleteUser`

	

	qry := fmt.Sprintf("SELECT extension FROM userextensions WHERE username = '%s'", p_username)
	fmt.Println("dbFindUserExtensions() qry=", qry)

	results, err := db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`14064.2: `, func_name, err.Error())
		return extens, err //Map{`statusCode`:502}
	}

	for results.Next() {

		var exten string
		// for each row, scan the result into our tag composite object
		err = results.Scan(&exten)
		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(`14064.3`, func_name, err.Error())
			continue
		}

		extens = append(extens, exten)

	}

	return extens, nil //Map{`statusCode`:503}
}

func dbFindUserExtensionsInGroup(username, p_groupid string) ([]string, error) {

	extens := []string{}
	func_name := `dbDeleteUser`

	

	qry := fmt.Sprintf("SELECT extension FROM userextensions WHERE username = '%s' AND groupid = '%s'", username, p_groupid)
	fmt.Println("dbFindUserExtensionsInGroup() qry=", qry)

	results, err := db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`14064.2: `, func_name, err.Error())
		return extens, err //Map{`statusCode`:502}
	}

	for results.Next() {

		var exten string
		// for each row, scan the result into our tag composite object
		err = results.Scan(&exten)
		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(`14064.3`, func_name, err.Error())
			continue
		}

		extens = append(extens, exten)

	}

	return extens, nil //Map{`statusCode`:503}
}

func dbTypeAheadContacts(q, p_username string) ([]map[string]string, error) {

	names := []map[string]string{}
	func_name := `dbTypeAheadContacts`

	

	qry := fmt.Sprintf("SELECT username,name,photo FROM users WHERE username <> '%s' AND name LIKE '%s%%'", p_username, q)
	fmt.Println("dbTypeAheadUser() qry=", qry)

	results, err := db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`14137.2: `, func_name, err.Error())
		return names, err //Map{`statusCode`:502}
	}

	for results.Next() {

		var name, username, photo string
		// for each row, scan the result into our tag composite object
		err = results.Scan(&username, &name, &photo)
		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(`14137.3`, func_name, err.Error())
			continue
		}

		info := ""
		extens, err := dbFindUserExtensions(username)
		fmt.Println(`14137.4`, func_name, `extens=`, extens)
		if err == nil && len(extens) > 0 {
			for _, s := range extens {
				info = fmt.Sprintf("%s#%s ", info, s)
			}
		}

		names = append(names, map[string]string{"username": username, "name": name, "photo": photo, "info": info})

	}

	return names, nil //Map{`statusCode`:503}
}

func dbTypeAheadExtensions(q, p_username string) ([]map[string]string, error) {

	names := []map[string]string{}
	func_name := `dbTypeAheadExtensions`

	

	//qry := fmt.Sprintf("SELECT username,name,photo FROM users WHERE username <> '%s' AND name LIKE '%s%%'",p_username, q)
	//SELECT u.username,name,photo,ue.extension FROM users u INNER JOIN userextensions ue ON u.username = ue.username WHERE u.username <> 'rquidilig@gmail.com' AND name LIKE 'rom%';
	qry := fmt.Sprintf("SELECT u.username,name,photo,ue.extension FROM users u INNER JOIN userextensions ue ON u.username = ue.username WHERE u.username <> '%s' AND name LIKE '%s%%'", p_username, q)
	// SELECT username,name,photo FROM users WHERE username <> 'rquidilig@gmail.com' AND name LIKE 'rom'
	fmt.Println("dbTypeAheadUser() qry=", qry)

	results, err := db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`14184.2: `, func_name, err.Error())
		return names, err //Map{`statusCode`:502}
	}

	for results.Next() {

		var name, username, photo, exten string
		// for each row, scan the result into our tag composite object
		err = results.Scan(&username, &name, &photo, &exten)
		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(`14184.3`, func_name, err.Error())
			continue
		}

		

		names = append(names, map[string]string{"username": username, "name": name, "photo": photo, "extension": exten})

	}

	return names, nil //Map{`statusCode`:503}
}

func dbLoadJoinRequests() error {


	func_name := `dbLoadJoinRequests`

	// Open up our database connection.

	/*
	MariaDB [xoftswitchdb]> describe joinrequests;
+-----------+-------------+------+-----+----------+-------+
| Field     | Type        | Null | Key | Default  | Extra |
+-----------+-------------+------+-----+----------+-------+
| id        | varchar(48) | NO   | PRI | NULL     |       |
| email     | varchar(50) | YES  |     | NULL     |       |
| username  | varchar(50) | NO   |     |          |       |
| hostname  | varchar(50) | NO   |     |          |       |
| name      | varchar(50) | NO   |     |          |       |
| status    | int(11)     | YES  |     | 0        |       |
| auto_join | tinyint(1)  | YES  |     | 0        |       |
| roleid    | varchar(30) | YES  |     | customer |       |
| created   | varchar(30) | NO   |     |          |       |
+-----------+-------------+------+-----+----------+-------+
9 rows in set (0.01 sec)
	*/

	qry := fmt.Sprintf("SELECT id,email,username,  name, status,auto_join,roleid, created FROM joinrequests")

	//fmt.Println(`10757:`,func_name,`qry=`,qry)

	results, err := db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`14780: `, func_name, err.Error())
		return err //Map{`statusCode`:502}
	}

	joinrequests = []JoinRequest{}
	for results.Next() {

		var joinreq JoinRequest
		// for each row, scan the result into our tag composite object
		err = results.Scan(&joinreq.Email, &joinreq.Username, &joinreq.Name,&joinreq.Status,&joinreq.Autojoin,&joinreq.Roleid, &joinreq.Created)
		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(`14255.3`, func_name, err.Error())
			continue
		}
		joinrequests = append(joinrequests, joinreq)

	}

	return nil

}

func dbUpsertJoinRequest(join JoinRequest) error {
	qry := `
		INSERT INTO joinrequests (
			id, hostname, email, username, name, created, status, auto_join, roleid
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
			hostname = VALUES(hostname),
			email = VALUES(email),
			username = VALUES(username),
			name = VALUES(name),
			created = VALUES(created),
			status = VALUES(status),
			auto_join = VALUES(auto_join),
			roleid = VALUES(roleid)
	`

	fmt.Println(string(colorYellow), "Executing UPSERT on joinrequests...")
	fmt.Println(string(colorReset))

	_, err := db.Exec(qry,
		join.Id,
		join.Hostname,
		join.Email,
		join.Username,
		join.Name,
		join.Created,
		join.Status,
		join.Autojoin,
		join.Roleid,
	)

	if err != nil {
		fmt.Printf("FAILED to upsert joinrequest: %v\n", err)
		return err
	}

	return nil
}

func dbRemoveJoinRequestById(id string) error {

	//func_name := `dbDeleteJoinRequest`

	

	qry := fmt.Sprintf("DELETE FROM joinrequests WHERE id = '%s'", id)

	del, err := db.Query(qry)

	// if there is an error inserting, handle it
	if err != nil {
		//panic(err.Error())
		fmt.Println("FAILED %s", qry)
		return err
	}
	defer del.Close()

	return nil //Map{`statusCode`:503}
}


func dbDeleteJoinRequestByEmail(email string) error {

	//func_name := `dbDeleteJoinRequest`

	

	qry := fmt.Sprintf("DELETE FROM joinrequests WHERE email = '%s'", email)

	del, err := db.Query(qry)

	// if there is an error inserting, handle it
	if err != nil {
		//panic(err.Error())
		fmt.Println("FAILED %s", qry)
		return err
	}
	defer del.Close()

	return nil //Map{`statusCode`:503}
}

func dbDeleteJoinRequest(id string) error {
	qry := `DELETE FROM joinrequests WHERE id = ?`

	fmt.Println(string(colorYellow), "Executing delete on joinrequests for id =", id)
	fmt.Println(string(colorReset))

	_, err := db.Exec(qry, id)
	if err != nil {
		fmt.Printf("FAILED to delete joinrequest: %v\n", err)
		return err
	}

	return nil
}


func dbLoadActivateExtensionRequests() error {

	//ret := []map[string]string{}
	//exten_num := p_exten
	func_name := `dbLoadActivateExtensionRequests`

	// Open up our database connection.

	

	qry := fmt.Sprintf("SELECT email,username,  name, created FROM activateextensionrequests")

	//fmt.Println(`10757:`,func_name,`qry=`,qry)

	results, err := db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`14872.2: `, func_name, err.Error())
		return err //Map{`statusCode`:502}
	}

	activateextensionrequests = []ActivateExtensionRequest{}
	for results.Next() {

		var aer ActivateExtensionRequest
		// for each row, scan the result into our tag composite object
		err = results.Scan(&aer.Email, &aer.Username, &aer.Name, &aer.Created)
		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(`14255.3`, func_name, err.Error())
			continue
		}
		activateextensionrequests = append(activateextensionrequests, aer)

	}

	return nil

}

func dbUpsertActivateExtensionRequest(aer ActivateExtensionRequest) error {
	//ret := []map[string]string{}
	//exten_num := p_exten
	//email, username, name string
	//func_name := `dbUpsertActivateExtensionRequest`

	

	//qry := fmt.Sprintf("REPLACE INTO joinrequests(email,username, name,created) VALUES('%s','%s','%s','%s')",
	qry := fmt.Sprintf("REPLACE INTO activateextensionrequests(email,username, name,created) VALUES('%s','%s','%s','%s')",
		aer.Email,
		aer.Username,
		aer.Name,
		aer.Created,
	)

	fmt.Println(string(colorYellow), "qry=", qry)
	fmt.Println(string(colorReset))
	insert, err := db.Query(qry)

	// if there is an error inserting, handle it
	if err != nil {
		//panic(err.Error())
		fmt.Println("FAILED %s", qry)
		return err
	}
	defer insert.Close()

	return nil //Map{`statusCode`:503}
}

func dbDeleteActivateExtensionRequest(email string) error {

	//func_name := `dbDeleteActivateExtensionRequest`

	

	qry := fmt.Sprintf("DELETE FROM activateextensionrequests WHERE email = '%s'", email)

	del, err := db.Query(qry)

	// if there is an error inserting, handle it
	if err != nil {
		//panic(err.Error())
		fmt.Println("FAILED %s", qry)
		return err
	}
	defer del.Close()

	return nil //Map{`statusCode`:503}
}

func dbUpsertBlockedUser(blockedby, blockeduser string, isblocked bool) error {

	//func_name := `dbUpsertUsert`

	if len(blockedby) == 0 || len(blockeduser) == 0 {
		fmt.Println(`11207: ERROR createUserExtension. Invalid username or blockeduser.`)
		return nil
	}

	
	/*
				REPLACE INTO BLOGPOSTs
			(
				postId, postTitle, postPublished
			)
		VALUES
			(5, 'Python Tutorial', '2019-08-04');
	*/

	if isblocked {
		qry := fmt.Sprintf("REPLACE INTO blockedusers(blockedby, blockeduser) VALUES('%s','%s')",
			blockedby,
			blockeduser,
		)

		fmt.Println(string(colorRed), "qry=", qry)
		fmt.Println(string(colorReset))
		insert, err := db.Query(qry)

		// if there is an error inserting, handle it
		if err != nil {
			//panic(err.Error())
			fmt.Println("FAILED %s", qry)
			return err
		}
		defer insert.Close()

	} else {
		qry := fmt.Sprintf("DELETE FROM blockedusers WHERE blockedby = '%s' AND blockeduser = '%s'", blockedby, blockeduser)

		del, err := db.Query(qry)

		// if there is an error inserting, handle it
		if err != nil {
			//panic(err.Error())
			fmt.Println("FAILED %s", qry)
			return err
		}
		defer del.Close()

	}

	return nil //Map{`statusCode`:503}
}

func dbRemoveExtensionFromMyContact(exten, username string) error {

	
	
	qry := fmt.Sprintf("DELETE FROM mycontacts WHERE username = '%s' AND number = '%s'", username, exten)

	del, err := db.Query(qry)

	// if there is an error inserting, handle it
	if err != nil {
		//panic(err.Error())
		fmt.Println("FAILED %s", qry)
		return err
	}
	defer del.Close()

	return nil //Map{`statusCode`:503}
}

func dbLoadUserExtensions() error {

	fmt.Println(`15271.1: START dbLoadUserExtensions()`)
	//ret := []map[string]string{}
	//exten_num := p_exten
	func_name := `dbLoadUserExtensions`

	// Open up our database connection.

	

	qry := fmt.Sprintf("SELECT username, extension,groupid FROM userextensions")

	//fmt.Println(`10900:`,func_name,`qry=`,qry)

	results, err := db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`14472.2: `, func_name, err.Error())
		return err //Map{`statusCode`:502}
	}

	userextensions = []UserExtension{}
	for results.Next() {

		var ue UserExtension
		// for each row, scan the result into our tag composite object
		err = results.Scan(&ue.Username, &ue.Extension, &ue.Groupid)
		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(`14472.3`, func_name, err.Error())
			continue
		}
		userextensions = append(userextensions, ue)

	}

	return nil

}

func dbRemoveUserExtensions(username string) error {

	//ret := []map[string]string{}
	//exten_num := p_exten
	func_name := `dbRemoveUserExtensions`

	// Open up our database connection.

	

	//qry := fmt.Sprintf("SELECT username, extension FROM userextensions")
	qry := fmt.Sprintf("DELETE  FROM userextensions WHERE username = '%s'", username)

	//fmt.Println(`10900:`,func_name,`qry=`,qry)

	results, err := db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`14534.2: `, func_name, err.Error())
		return err //Map{`statusCode`:502}
	}

	userextensions = []UserExtension{}
	for results.Next() {

		var ue UserExtension
		// for each row, scan the result into our tag composite object
		err = results.Scan(&ue.Username, &ue.Extension)
		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(`14534.3`, func_name, err.Error())
			continue
		}
		userextensions = append(userextensions, ue)

	}

	return nil

}

func dbUpsertUserExtension(ue UserExtension) error {

	//func_name := `dbUpsertUsert`

	if len(ue.Username) == 0 || len(ue.Extension) == 0 {
		fmt.Println(`11207: ERROR createUserExtension. Invalid username or extension.`)
	}

	
	/*
				REPLACE INTO BLOGPOSTs
			(
				postId, postTitle, postPublished
			)
		VALUES
			(5, 'Python Tutorial', '2019-08-04');
	*/

	qry := fmt.Sprintf("REPLACE INTO userextensions(username, extension,groupid) VALUES('%s','%s','%s')",
		ue.Username,
		ue.Extension,
		ue.Groupid,
	)

	fmt.Println(string(colorRed), "qry=", qry)
	fmt.Println(string(colorReset))
	insert, err := db.Query(qry)

	// if there is an error inserting, handle it
	if err != nil {
		//panic(err.Error())
		fmt.Println("FAILED %s", qry)
		return err
	}
	defer insert.Close()

	return nil //Map{`statusCode`:503}
}

func dbDeleteUserExtension(username, exten string) error {

	//func_name := `dbDeleteUser`

	

	qry := fmt.Sprintf("DELETE FROM userextensions WHERE username = '%s' AND extension = '%s'", username, exten)

	del, err := db.Query(qry)

	// if there is an error inserting, handle it
	if err != nil {
		//panic(err.Error())
		fmt.Println("FAILED %s", qry)
		return err
	}
	defer del.Close()

	return nil //Map{`statusCode`:503}
}

func dbDeleteUserExtensionsInGroup(groupid string) error {

	//func_name := `dbDeleteUserExtensionsInGroup()`

	

	//qry := fmt.Sprintf("DELETE FROM userextensions WHERE extension = '%s' AND groupid <> '%s'", exten, groupid)
	qry := fmt.Sprintf("DELETE FROM userextensions WHERE groupid = '%s'", groupid)
	del, err := db.Query(qry)

	// if there is an error inserting, handle it
	if err != nil {
		//panic(err.Error())
		fmt.Println("FAILED %s", qry)
		return err
	}
	defer del.Close()

	return nil //Map{`statusCode`:503}
}

func dbDeleteUserExtensionsNotInGroup(exten, groupid string) error {

	//func_name := `dbDeleteUserExtensionsNotInGroup()`

	

	qry := fmt.Sprintf("DELETE FROM userextensions WHERE extension = '%s' AND groupid <> '%s'", exten, groupid)

	del, err := db.Query(qry)

	// if there is an error inserting, handle it
	if err != nil {
		//panic(err.Error())
		fmt.Println("FAILED %s", qry)
		return err
	}
	defer del.Close()

	return nil //Map{`statusCode`:503}
}

func dbDeleteAllUserExtensionsByUsername(username string) error {

	//func_name := `dbDeleteUser`
	fmt.Println(`14601.1: dbDeleteAllUserExtensionsByUsername start...`)

	
	qry := fmt.Sprintf("DELETE FROM userextensions WHERE username = '%s'", username)

	del, err := db.Query(qry)

	// if there is an error inserting, handle it
	if err != nil {
		//panic(err.Error())
		fmt.Println("FAILED %s", qry)
		return err
	}
	defer del.Close()

	return nil //Map{`statusCode`:503}
}

func dbDeleteAllUserExtensionsByExtension(exten string) error {

	//func_name := `dbDeleteUser`

	

	qry := fmt.Sprintf("DELETE FROM userextensions WHERE extension = '%s')", exten)

	del, err := db.Query(qry)

	// if there is an error inserting, handle it
	if err != nil {
		//panic(err.Error())
		fmt.Println("FAILED %s", qry)
		return err
	}
	defer del.Close()

	return nil //Map{`statusCode`:503}
}

func httpGetExtensionResponse(p_exten string) Map {

	//func_name := `httpGetExtensionResponse`

	var exten *Extension
	for _, v := range extensions {

		if v.Number == p_exten {
			exten = &v
			break
		}

	}

	if exten == nil {
		return Map{`statusCode`: 500}
	}


	return Map{`statusCode`: 200, `data`: exten}

}


func httpGetMyExtensionsResponse(p_exten string) Map {
    var myextensions = dbGetMyUserExtensions(p_exten, 10)
    return Map{`statusCode`: 200, `data`: myextensions}
}

func httpGetMyTokenResponse(p_exten string) Map {
    var myextensions = dbGetMyUserExtensions(p_exten, 10)
    return Map{`statusCode`: 200, `data`: myextensions}
}





func httpGetUserMissedUnreadResponse_dep(p_username string) Map {

	//func_name := `httpGetExtensionResponse`

	var u *User
	for _, v := range users {

		if v.Username == p_username {
			u = &v
			break
		}

	}

	if u == nil {
		return Map{`statusCode`: 500}
	}

	//json_item, err1 := json.Marshal(exten)
	//fmt.Println(func_name, "json_item=",json_item)
	//if err1 == nil{

	return Map{`statusCode`: 200, `data`: u}

}

func dbGetUser(username string) *User {
    func_name := "dbGetUser"
    fmt.Println("START:", func_name)

    // Use parameterized query to avoid SQL injection
    query := `
        SELECT username, email, name, roleid, status, sessionid, missed,
               misseduniqueid, unreadmessages, homepageurl, ver
        FROM users
        WHERE username = ?
    `

    fmt.Println(string(colorYellow), "15010.1: qry=", query, "param=", username)
    fmt.Println(string(colorReset))

    var user User
    // QueryRow with placeholder
    row := db.QueryRow(query, username)

    // Scan the result into user struct
    err := row.Scan(
        &user.Username,
        &user.Email,
        &user.Name,
        &user.Roleid,
        &user.Status,
        &user.Sessionid,
        &user.Missed,
        &user.Misseduniqueid,
        &user.Unreadmessages,
        &user.Homepageurl,
        &user.Ver,
    )

    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("15010.2:", func_name, "no rows found for username:", username)
            return nil
        }
        fmt.Println(string(colorYellow), "Scan error:", err.Error(), string(colorReset))
        return nil
    }

    fmt.Println("user=", user)
    return &user
}

func dbGetAIAToken(username, token, ipaddress string) string {
	func_name := "dbGetAIAToken"
	fmt.Println("START:", func_name)
	
	qry := fmt.Sprintf("SELECT TOP 1 token FROM aiatokens WHERE username = '%s' AND token = '%s' AND ipaddress = '%s'", username, token, ipaddress)

	fmt.Println(string(colorYellow), "15078.1: qry=", qry)
	fmt.Println(string(colorReset))
	//fmt.Println(`11191:`,func_name,`qry=`,qry)

	var ret_token string
	row := db.QueryRow(qry)
/*
	if row == nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`15078.2: `, func_name, err.Error())
		return `` //Map{`statusCode`:501}
	}
	*/

	//err = row.Scan(&user.Username,&user.Email,&user.Name,&user.Roleid,&user.Status,&user.Sessionid,&user.Missed,&user.Misseduniqueid,&user.Unreadmessages,&user.Homepageurl,&user.Ver)
	err := row.Scan(&ret_token)
	if err != nil {

		fmt.Println(string(colorYellow), err.Error())
		fmt.Println(string(colorReset))
	}

	fmt.Println("ret_token=", ret_token)
	return ret_token //Map{`statusCode`:200, `data`:user}
}

func isInrole(username, roleid string) bool {
    func_name := "dbIsInrole"
    fmt.Println("START:", func_name)

    // Use a parameterized query to avoid SQL injection
    query := "SELECT username FROM users WHERE username = ? AND roleid = ?"

    fmt.Println(string(colorYellow), "15131.1: qry=", query, "params:", username, roleid)
    fmt.Println(string(colorReset))

    var foundUsername string
    err := db.QueryRow(query, username, roleid).Scan(&foundUsername)
    if err != nil {
        if err == sql.ErrNoRows {
            // No match found
            fmt.Println("15131.2:", func_name, "no match for user/role:", username, roleid)
            return false
        }
        // Some other error occurred
        fmt.Println("15131.3:", func_name, "scan error:", err.Error())
        return false
    }

    // ✅ If we got here, the user with that role exists
    fmt.Println("15131.4:", func_name, "user in role:", foundUsername, roleid)
    return true
}

func httpAuthenticateAIATokenResponse(ipaddress, token string) Map {
    func_name := "httpAuthenticateAIATokenResponse"
    fmt.Println("START:", func_name)

    // ✅ Use parameterized query
    selectQry := "SELECT username FROM aitokens WHERE token = ? AND ipaddress = ?"
    fmt.Println(string(colorYellow), "15198.1: qry=", selectQry, "params:", token, ipaddress)
    fmt.Println(string(colorReset))

    var r_username string
    err := db.QueryRow(selectQry, token, ipaddress).Scan(&r_username)
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("15198.2:", func_name, "no token found for IP:", ipaddress)
            return Map{`statusCode`: 404} // not found
        }
        fmt.Println("15198.3:", func_name, "scan error:", err)
        return Map{`statusCode`: 500}
    }

    fmt.Println("r_username=", r_username)

    // ✅ Delete the token using Exec
    deleteQry := "DELETE FROM aitokens WHERE token = ?"
    res, err := db.Exec(deleteQry, token)
    if err != nil {
        fmt.Printf("FAILED to delete token (%s): %v\n", token, err)
        return Map{`statusCode`: 500}
    }

    // Optional: check rows affected
    if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
        fmt.Println("No rows deleted for token:", token)
    }

    return Map{`statusCode`: 200, `data`: r_username}
}


func httpGetMyHomePageUrlResponse(username, ipaddress string) Map {
	func_name := "httpGetMyHomePageUrlResponse"
	fmt.Println("14699 START:", func_name)
	/*
		username,ok_username:= data["username"].(string)
		if !ok_username{
			return Map{`statusCode`:500}
		}
	*/

	user := dbGetUser(username)

	if user == nil {
		fmt.Println(`httpGetMyHomePageUrlResponse.1: username`, username, `not found`)
		return Map{`statusCode`: 500}
	}

	homepageurl := `https://xoftphone.com/noaia.html`
	aia_url, ok_aia_url := config[`aia_url`]
	if ok_aia_url && len(aia_url) > 0 {
		//homepageurl = fmt.Sprintf(`https://%s`, hostname)
		homepageurl = aia_url
	}

	

	if len(homepageurl) > 0 {
		aia_use_token, ok_aia_use_token := config[`aia_use_token`]
		if ok_aia_use_token && len(aia_use_token) > 0 {

			token := uuid.New().String()
			//nowiso := time.Now().Format(time.RFC3339)

			qry := fmt.Sprintf("DELETE FROM aiatokens WHERE username = '%s' AND ipaddress = '%s'", username, ipaddress)

			del, err := db.Query(qry)

			// if there is an error inserting, handle it
			if err != nil {
				//panic(err.Error())
				fmt.Println("FAILED %s", qry)
				return Map{`statusCode`: 502}
			}
			defer del.Close()

			//}else{
			//qry := fmt.Sprintf("REPLACE INTO aiatokens(token,username,ipaddress) VALUES('%s','%s','%s')",token, username,ipaddress)

			qry = fmt.Sprintf("INSERT INTO aiatokens(token,username,ipaddress) VALUES('%s','%s','%s')", token, username, ipaddress)

			fmt.Println(string(colorYellow), "16248.1: qry=", qry)
			fmt.Println(string(colorReset))
			_, err = db.Query(qry)
			if err != nil {
				fmt.Println(string(colorRed), `16248.2: httpGetMyHomePageUrlResponse.4:`, func_name, `err:=`, err.Error())
				fmt.Println(string(colorReset))
				return Map{`statusCode`: 503}
			}

			aia_use_token = strings.Replace(aia_use_token, `$token`, token, 0)
			homepageurl = fmt.Sprintf(`%s%s`, homepageurl, aia_use_token)
		}
	}

	
	fmt.Println(`16248.3: httpGetMyHomePageUrlResponse.5: Done! homepageurl=`, homepageurl)

	return Map{`statusCode`: 200, `data`: homepageurl}

}


func httpGetMyUserResponse(data Map) Map {
		func_name := "httpGetMyUserResponse"
		fmt.Println("16097.1 START:", func_name)
	
		username, ok := data["username"].(string)
		if !ok {
			return Map{`statusCode`: 500}
		}
	
		resetMissed, okMissed := data["reset_missed"].(bool)
		if !okMissed {
			return Map{`statusCode`: 501}
		}
		resetUnread, okUnread := data["reset_unreadmessages"].(bool)
		if !okUnread {
			return Map{`statusCode`: 502}
		}
	
		// Build update query if needed
		if resetMissed || resetUnread {
			var qryUpdate string
			if resetMissed && resetUnread {
				qryUpdate = "UPDATE users SET missed = 0, unreadmessages = 0 WHERE username = ?"
			} else if resetMissed {
				qryUpdate = "UPDATE users SET missed = 0 WHERE username = ?"
			} else if resetUnread {
				qryUpdate = "UPDATE users SET unreadmessages = 0 WHERE username = ?"
			}
	
			if qryUpdate != "" {
				fmt.Println(string(colorYellow), "16097.2: qry_update=", qryUpdate, username)
				fmt.Println(string(colorReset))
	
				// ✅ Use Exec for UPDATE
				if _, err := db.Exec(qryUpdate, username); err != nil {
					fmt.Println(string(colorRed), "16097.3: err=", err.Error())
					fmt.Println(string(colorReset))
					return Map{`statusCode`: 504}
				}
			}
		}
	
		// Now fetch the user
		selectQry := `
			SELECT username, email, name, roleid, status, sessionid, missed, misseduniqueid,
				   unreadmessages, homepageurl, ver
			FROM users
			WHERE username = ?
		`
		fmt.Println(string(colorYellow), "16097.4: qry=", selectQry, "param=", username)
		fmt.Println(string(colorReset))
	
		var user User
		err := db.QueryRow(selectQry, username).Scan(
			&user.Username,
			&user.Email,
			&user.Name,
			&user.Roleid,
			&user.Status,
			&user.Sessionid,
			&user.Missed,
			&user.Misseduniqueid,
			&user.Unreadmessages,
			&user.Homepageurl,
			&user.Ver,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("16097.5:", func_name, "no user found for username:", username)
				return Map{`statusCode`: 404}
			}
			fmt.Println(string(colorYellow), "16097.6: scan err=", err.Error())
			fmt.Println(string(colorReset))
			return Map{`statusCode`: 500}
		}
	
		fmt.Println("user=", user)
		return Map{`statusCode`: 200, `data`: user}
	}
	
func httpGetMyAPITokenResponse(data Map) Map {
	func_name := "httpGetMyAPITokenResponse"
	fmt.Println("START:", func_name)

	username, ok := data["username"].(string)
	if !ok {
		return Map{"statusCode": 500, "error": "username missing or invalid"}
	}

	

	qry := fmt.Sprintf(`SELECT apitoken, apitokenexpire FROM users WHERE username = '%s'`, username)
	fmt.Println(func_name, "query:", qry)

	var user User
	row := db.QueryRow(qry)
	err := row.Scan(&user.Apitoken, &user.Apitokencreated)
	if err != nil {
		if err == sql.ErrNoRows {
			// No user found
			fmt.Println(func_name, "no rows found for username:", username)
			return Map{"statusCode": 404, "error": "user not found"}
		}
		// Other errors
		fmt.Println(func_name, "scan error:", err.Error())
		return Map{"statusCode": 500, "error": "query scan error"}
	}

	fmt.Println(func_name, "user data:", user)

	// Return the API token data inside 'data' key
	return Map{
		"statusCode": 200,
		"data": Map{
			"apitoken":        user.Apitoken,
			"apitokencreated": user.Apitokencreated,
		},
	}
}


func httpMoreExtensionsResponse(data Map) Map {
	func_name := "httpMoreExtensionsResponse"
	fmt.Println("START:", func_name)
	fmt.Println("data:", data)

	q, ok_q := data["q"].(string)

	if !ok_q {
		q = ""
	}

	skip, ok_skip := data["skip"].(float64)

	if !ok_skip {
		skip = 0
	}

	fmt.Println("**** skip:", skip)

	limit := 200
	
	page_num, ok_page_num := data["page_num"].(int)
	if !ok_page_num {
		page_num = 1
	}
	fmt.Println("q=", q, "ok_q=", ok_q, "skip=", skip, "ok_skip=", ok_skip, "page_num=", page_num, "ok_page_num=", ok_page_num)
	//page_count := 100

	//item_count := 0
	total := len(extensions)
	items := []MoreExtensionInfo{}

	
	qry := fmt.Sprintf(`SELECT DISTINCT ext.number,IFNULL(ue.username,'') as username,IFNULL(ue.name,'') as name,IFNULL(ue.photo,'') as photo,displayname, status,contacts,source,IFNULL(reg.device_type,'') as device_type, public  FROM extensions ext LEFT JOIN (SELECT ue.username,extension,name,photo FROM userextensions ue LEFT JOIN users u ON  ue.username= u.username) ue ON ext.number = ue.extension LEFT JOIN registered reg ON ext.number = reg.exten LIMIT %v OFFSET %v`, limit, skip)
	
	
	fmt.Println(`15229.2:`, func_name, `qry=`, qry)

	results, err := db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`15229.3: `, func_name, err.Error())
		return Map{`statusCode`: 502}
	}

	//var inboxes []INBOX
	//var extens []Extension
	re_remove := regexp.MustCompile(`\<(.*)\>`)

	for results.Next() {
		//var inbox INBOX
		var exten MoreExtensionInfo
		// for each row, scan the result into our tag composite object
		err = results.Scan(&exten.Number, &exten.Username, &exten.Name, &exten.Photo, &exten.Displayname, &exten.Status, &exten.Contacts, &exten.Source, &exten.Devicetype, &exten.Public)
		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(`15229.4`, func_name, err.Error())
			continue
		}

		/*if len(exten.Displayname) == 0{
			exten.Displayname = exten.Number
		}
		*/
		ma_remove := re_remove.FindStringSubmatch(exten.Displayname)
		//fmt.Println(`ma_key=`,ma_key)
		if len(ma_remove) == 2 {
			//must key line

			r := ma_remove[1]

			exten.Displayname = strings.Replace(exten.Displayname, fmt.Sprintf(`<%s>`, r), ``, 1)
		}

		
		
		items = append(items, exten)

	}
	//remove extra <name>

	fmt.Println("10744: len(items) =", len(items))
	fmt.Println("10744: items=", items)

	ret_items := []Map{}
	if len(items) > 0 {
		for _, item := range items {

			exten := Map{
				`number`:   item.Number,
				`username`: item.Username,
				`name`:     item.Name,
				`photo`:    item.Photo,

				`displayname`: item.Displayname,
				`status`:      item.Status,
				`contacts`:    item.Contacts,
				`source`:      item.Source,
				`devicetype`:  item.Devicetype,
				`public`:      item.Public,
			}
			ret_items = append(ret_items, exten)
		}

	}

	return Map{`statusCode`: 200, `data`: Map{`total`: total, `page_num`: page_num, `items`: ret_items}}
}

func httpMoreMyContactExtensionsResponse(data Map) Map {
	func_name := "httpMoreMyContactExtensionsResponse"
	fmt.Println("15516.0: START:", func_name)
	fmt.Println("data:", data)

	username, ok_username := data["username"].(string)

	if !ok_username {

		fmt.Println(string(colorRed), func_name, "15516.1: Username not supplied!!!")
		fmt.Println(string(colorReset))
		return Map{`statusCode`: 500}
	}

	q, ok_q := data["q"].(string)

	if !ok_q {
		q = ""
	}

	skip, ok_skip := data["skip"].(float64)

	if !ok_skip {
		skip = 0
	}

	fmt.Println("**** skip:", skip)

	limit := 200
	/*limit,ok_limit:= data["limit"].(int)

	if !ok_limit{ limit = 0}

	offset,ok_offset:= data["offset"].(int)

	if !ok_offset{ offset = 0}
	*/
	page_num, ok_page_num := data["page_num"].(int)
	if !ok_page_num {
		page_num = 1
	}
	fmt.Println("q=", q, "ok_q=", ok_q, "skip=", skip, "ok_skip=", ok_skip, "page_num=", page_num, "ok_page_num=", ok_page_num)
	//page_count := 100

	//item_count := 0
	total := len(extensions)
	items := []MoreExtensionInfo{}

	
	
		qry := fmt.Sprintf(
			`SELECT DISTINCT ext.number,
			IFNULL(ue.username,'') as username,
			IFNULL(ue.name,'') as name,
			IFNULL(ue.photo,'') as photo,
			IFNULL(ext.displayname, '') as displayname,
			status,
			contacts,
			source,
			IFNULL(reg.device_type,'') as device_type,
			public, canblock,canunlist,candial 
			FROM extensions ext LEFT JOIN mycontacts mc
			ON ext.number = mc.number AND mc.username = '%s' 
			LEFT JOIN (SELECT ue.username,extension,name,photo 
			FROM userextensions ue LEFT JOIN users u ON  ue.username= u.username) ue 
			ON ext.number = ue.extension 
			LEFT JOIN registered reg 
			ON ext.number = reg.exten WHERE ext.displayname IS NOT NULL 
			AND ext.displayname != '' 
			ORDER BY (status = 'Reachable') DESC, displayname ASC  
			LIMIT %v OFFSET %v`, username, limit, skip)
	
			
	fmt.Println(`15516.2:`, func_name, `qry=`, qry)

	results, err := db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`15516.3: `, func_name, err.Error())
		return Map{`statusCode`: 502}
	}

	re_remove := regexp.MustCompile(`\<(.*)\>`)

	//hasrecords := false

	for results.Next() {
		//var inbox INBOX
		//hasrecords = true
		var exten MoreExtensionInfo
		// for each row, scan the result into our tag composite object
		err = results.Scan(
			&exten.Number,
			&exten.Username,
			&exten.Name,
			&exten.Photo,
			&exten.Displayname,
			&exten.Status,
			&exten.Contacts,
			&exten.Source,
			&exten.Devicetype,
			&exten.Public,
			&exten.Canblock,
			&exten.Canunlist,
			&exten.Candial,
		)
		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(`15516.4`, func_name, err.Error())
			continue
		}

		/*if len(exten.Displayname) == 0{
			exten.Displayname = exten.Number
		}
		*/
		ma_remove := re_remove.FindStringSubmatch(exten.Displayname)
		//fmt.Println(`ma_key=`,ma_key)
		if len(ma_remove) == 2 {
			//must key line

			r := ma_remove[1]

			exten.Displayname = strings.Replace(exten.Displayname, fmt.Sprintf(`<%s>`, r), ``, 1)
		}

		
		
		items = append(items, exten)

	}
	//remove extra <name>

	fmt.Println("15516.5: len(items) =", len(items))
	fmt.Println("15516.6: items=", items)

	if len(items) == 0 {

	}

	ret_items := []Map{}
	if len(items) > 0 {
		for _, item := range items {

			exten := Map{
				`number`:   item.Number,
				`username`: item.Username,
				`name`:     item.Name,
				`photo`:    item.Photo,

				`displayname`: item.Displayname,
				`status`:      item.Status,
				`contacts`:    item.Contacts,
				`source`:      item.Source,
				`devicetype`:  item.Devicetype,
				`public`:      item.Public,
				`canblock`:    item.Canblock,
				`canunlist`:   item.Canunlist,
				`candial`:     item.Candial,

				//`canaddlist`:item.Canaddlist,

			}
			ret_items = append(ret_items, exten)
		}

	}

	return Map{`statusCode`: 200, `data`: Map{`total`: total, `page_num`: page_num, `items`: ret_items}}
}

func dbGetUserExtensions(username string, limit int) []string {
	func_name := "httpGetUserExtensionsResponse"

	//limit := 100
	items := []string{}

	

	qry := fmt.Sprintf(`SELECT extension FROM userextensions WHERE username = '%s' LIMIT %v`, username, limit)

	fmt.Println(`10119:`, func_name, `qry=`, qry)

	results, err := db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`8298.2: `, func_name, err.Error())
		//return Map{`statusCode`:502}
		return items
	}

	for results.Next() {
		//var inbox INBOX
		var exten string
		// for each row, scan the result into our tag composite object
		err = results.Scan(&exten)
		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(`8298.3`, func_name, err.Error())
			continue
		}

		items = append(items, exten)

	}
	//remove extra <name>

	fmt.Println("10744: len(items) =", len(items))
	fmt.Println("10744: items=", items)

	
	return items
}


func dbGetMyUserExtensions(username string, limit int) []UserExtension {
    funcName := "dbGetUserExtensionsResponse"
    items := []UserExtension{}

    

    query := `SELECT extension, groupid, exten_password FROM userextensions WHERE username = ? LIMIT ?`
    rows, err := db.Query(query, username, limit)
    if err != nil {
        fmt.Println("Query error:", funcName, err)
        return items
    }
    
	
    for rows.Next() {
        var exten UserExtension
        if err := rows.Scan(&exten.Extension,&exten.Groupid, &exten.Extenpassword); err != nil {
            fmt.Println("Scan error:", funcName, err)
            continue
        }
        items = append(items, exten)
    }

    if err := rows.Err(); err != nil {
        fmt.Println("Row iteration error:", funcName, err)
    }

    fmt.Printf("Found %d extensions\n", len(items))
    return items
}



func httpMoreUsersResponse(data Map) Map {
	func_name := "httpMoreUsersResponse"
	fmt.Println("START:", func_name)
	fmt.Println("data:", data)

	q, ok_q := data["q"].(string)

	me, ok_me := data["me"].(string)

	if !ok_q {
		q = ""
	}

	if !ok_me {
		fmt.Println("httpMoreUsersResponseInvalid me!")
		return Map{`statusCode`: 500}
	}

	skip, ok_skip := data["skip"].(float64)

	if !ok_skip {
		skip = 0
	}

	fmt.Println("**** skip:", skip)

	limit := 100
	/*limit,ok_limit:= data["limit"].(int)

	if !ok_limit{ limit = 0}

	offset,ok_offset:= data["offset"].(int)

	if !ok_offset{ offset = 0}
	*/
	page_num, ok_page_num := data["page_num"].(int)
	if !ok_page_num {
		page_num = 1
	}
	fmt.Println("q=", q, "ok_q=", ok_q, "skip=", skip, "ok_skip=", ok_skip, "page_num=", page_num, "ok_page_num=", ok_page_num)
	//page_count := 100

	//item_count := 0
	total := len(users)
	items := []UserMoreInfo{}

	qry := fmt.Sprintf(`SELECT username,name,firstname, lastname, phonenumber, photo, roleid,status , missed , unreadmessages, homepageurl, ver, IFNULL(blockedby,'') blockedby  FROM users u LEFT JOIN blockedusers bu ON u.username = bu.blockeduser AND bu.blockedby = '%s'  ORDER BY name asc LIMIT %v OFFSET %v`, me, limit, skip)

	fmt.Println(`10119:`, func_name, `qry=`, qry)

	results, err := db.Query(qry)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(`15491.2: `, func_name, err.Error())
		return Map{`statusCode`: 502}
	}

	

	for results.Next() {
		
		var user UserMoreInfo
		
		err = results.Scan(
			&user.Username,
			&user.Name,
			&user.Firstname,
			&user.Lastname,
			&user.Phonenumber,
			&user.Photo,
			&user.Roleid,
			&user.Status,
			&user.Missed,
			&user.Unreadmessages,
			&user.Homepageurl,
			&user.Ver,
			&user.Blockedby,
		)
		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(`15491.3`, func_name, err.Error())
			continue
		}

		items = append(items, user)

	}
	//remove extra <name>

	fmt.Println("10744: len(items) =", len(items))
	fmt.Println("10744: items=", items)

	ret_items := []Map{}
	if len(items) > 0 {
		for _, item := range items {

			extens := dbGetUserExtensions(item.Username, 10)
			u := Map{
				`username`:       item.Username,
				`email`:          item.Email,
				`name`:           item.Name,
				`firstname`:      item.Firstname,
				`lastname`:       item.Lastname,
				`photo`:          item.Photo,
				`phonenumber`:    item.Phonenumber,
				`roleid`:         item.Roleid,
				`status`:         item.Status,
				`missed`:         item.Missed,
				`unreadmessages`: item.Unreadmessages,
				`homepageurl`:    item.Homepageurl,
				`ver`:            item.Ver,
				`blockedby`:      item.Blockedby,
				`extens`:         extens,
			}
			ret_items = append(ret_items, u)
		}

	}

	return Map{`statusCode`: 200, `data`: Map{`total`: total, `page_num`: page_num, `items`: ret_items}}
}

func httpDeleteUserResponse(m Map) Map {
	func_name := "httpMoreUsersResponse"
	fmt.Println("START:", func_name)
	fmt.Println("m:", m)

	admin, ok_admin := m["admin"].(string)
	username, ok_username := m["username"].(string)

	if ok_admin && ok_username {

		if isInrole(admin, `admin`) {

			

			dbDeleteUser(username)

		}
	} else {
		fmt.Println(`httpDeleteUserResponse() admin or data param does not pass validataion!`)
	}

	return Map{`statusCode`: 200}
}

func httpAdminDeleteExtenResponse(m Map) Map {
	func_name := "httpAdminDeleteExtenResponse"
	fmt.Println("START:", func_name)
	fmt.Println("m:", m)

	admin, ok_admin := m["admin"].(string)
	exten, ok_exten := m["exten"].(string)

	if ok_admin && ok_exten {

		if isInrole(admin, `admin`) {

			//delusername,ok_delusername:= m["username"].(string)

			//if !ok_delusername{
			//return Map{`statusCode`:200}
			//}

			dbDeleteExtension(exten)

		}
	} else {
		fmt.Println(`httpDeleteUserResponse() admin or data param does not pass validataion!`)
	}

	return Map{`statusCode`: 200}
}

func httpAdminEditExtenResponse(m Map) Map {
	func_name := "httpAdminEditExtenResponse"
	fmt.Println("START:", func_name)
	fmt.Println("m:", m)

	admin, ok_admin := m["admin"].(string)
	exten, ok_exten := m["exten"].(string)
	public_val, ok_public := m["public_val"].(string)
	displayname_val, ok_displayname := m["displayname_val"].(string)

	isvalidiput := false

	if ok_admin && ok_exten {
		fmt.Println("16290.1: PASSED:", func_name)
		if ok_public || ok_displayname {
			fmt.Println("16290.2: PASSED:", func_name)

			if isInrole(admin, `admin`) {
				fmt.Println("16290.3: PASSED:", func_name)

				isvalidiput = true
			} else {
				fmt.Println("16290.4: FAILED:", func_name)

			}
		} else {
			fmt.Println("16290.5: FAILED:", func_name)
		}
	}

	if !isvalidiput {
		fmt.Println(`16290.6: httpDeleteUserResponse() admin or data param does not pass validataion!`)
		return Map{`statusCode`: 500}
	}

	if ok_public {
		dbExtensionSetPublic(exten, public_val)

	}
	if ok_displayname {
		dbMyContactSetDisplayname(exten, displayname_val)
	}

	return Map{`statusCode`: 200}
}

func httpBlockUserResponse(m Map) Map {
	func_name := "httpMoreUsersResponse"
	fmt.Println("START:", func_name)
	fmt.Println("m:", m)

	//admin, ok_admin := m["admin"].(string)
	blockeduser, ok_blockeduser := m["blockeduser"]
	blockedby, ok_blockedby := m["blockedby"]
	isblocked, ok_isblocked := m["isblocked"]

	fmt.Println("12562: ok_blockeduser:", ok_blockeduser, "ok_blockedby", ok_blockedby, "ok_isblocked", ok_isblocked)

	if ok_blockedby && ok_blockeduser && ok_isblocked {
		fmt.Println(`12347: httpBlockUserResponse blockeduser=`, blockeduser, "blockedby=", blockedby, "isblocked=", isblocked)

		dbUpsertBlockedUser(blockedby.(string), blockeduser.(string), isblocked.(bool))

	} else {
		fmt.Println(`12351: httpBlockUserResponse invalid params`)
	}

	return Map{`statusCode`: 200}
}

func httpUnlistExtensionResponse(m Map) Map {
	func_name := "httpUnlistExtensionResponse"
	fmt.Println("START:", func_name)
	fmt.Println("m:", m)

	//admin, ok_admin := m["admin"].(string)
	//blockeduser,ok_blockeduser := m["blockeduser"]
	//blockedby,ok_blockedby := m["blockedby"]
	//isblocked,ok_isblocked := m["isblocked"]
	username, ok_username := m["username"].(string)
	exten, ok_exten := m["exten"].(string)

	//fmt.Println("12562: ok_blockeduser:", ok_blockeduser,"ok_blockedby",ok_blockedby,"ok_isblocked",ok_isblocked)

	if ok_exten && ok_username {
		//fmt.Println(`12347: httpBlockUserResponse blockeduser=`,blockeduser,"blockedby=",blockedby,"isblocked=",isblocked)

		err := dbRemoveExtensionFromMyContact(exten, username)
		if err != nil {
			return Map{`statusCode`: 500}
		}

	} else {
		fmt.Println(`12351: httpBlockUserResponse invalid params`)
	}

	return Map{`statusCode`: 200}
}


func sqlAddNewMessage(p_username, inboxid, from_user, to_user, from_name, to_name, roomid, message, attachment, messageid string) bool {
	//ret := []map[string]string{}
	//exten_num := p_exten
	func_name := `sqlAddNewMessage`
	fmt.Println(`9882.0 START`, func_name, `p_username=`, p_username,
		`inboxid=`, inboxid,
		`from_user=`, from_user,
		`to_user=`, to_user,
		`from_name=`, from_name,
		`to_name=`, to_name,
		`roomid=`, roomid,
		`message=`, message,
		`attachment=`, attachment,
		`messageid=`, messageid)

	

	//var qry =  fmt.Sprintf("SELECT * FROM inbox WHERE id = '%s-%s') OR id = '%s-%s') ", from_exten,to_exten,to_exten,from_exten)
	//create 2 copies of invox entry one for each extension(from,to)

	// START from inbox copy
	var qry string
	//inbox_id := from_exten
	//var qry =  fmt.Sprintf("SELECT id FROM inbox WHERE id = '%s'"), inbox_id)

	var id string
	qry = fmt.Sprintf("SELECT id FROM inbox WHERE id = '%s'", inboxid)
	fmt.Println(`sqlAddNewMessage`, `qry=`, qry)
	row := db.QueryRow(qry)

	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		fmt.Println("9882.1: ", func_name, "No inbox", inboxid, "returned! Create if needed.")
		//inbox_id = fmt.Sprintf(`%s-%s`,from_exten,to_exten)
		qry = fmt.Sprintf("INSERT INTO inbox(id,created, username, `from`,`to`,fromname,toname, roomid,messageid, message) VALUES('%s',now(),'%s','%s','%s','%s','%s','%s','%s','%s')", inboxid, p_username, from_user, to_user, from_name, to_name, roomid, messageid, message)
		fmt.Println(`9882.2`, func_name, `qry=`, qry)
		// perform a db.Query insert
		rows_insert, err := db.Query(qry)
		if err != nil {
			fmt.Println(`9882.3:`, func_name, `err:=`, err.Error())
			return false
		}
		rows_insert.Close()

	case nil:
		//fmt.Println(id)
		//inbox_id = fmt.Sprintf(`%s-%s`,from_exten,to_exten)
		qry = fmt.Sprintf("UPDATE inbox SET created = now(), username = '%s' , `from` = '%s',`to` = '%s', fromname = '%s',toname = '%s',roomid = '%s', messageid = '%s', message = '%s' WHERE id = '%s'", p_username, from_user, to_user, from_name, to_name, roomid, messageid, message, inboxid)
		// perform a db.Query insert
		fmt.Println(`9882.4:`, func_name, `qry=`, qry)
		rows, err := db.Query(qry)
		if err != nil {
			fmt.Println(`9882.5:`, func_name, `err:=`, err.Error())
			//return Map{`statusCode`:503}
			return false
		}

		defer rows.Close()

	default:
		//return Map{`statusCode`:501}
		fmt.Println("9882.6: ", func_name, err)
		return false
		//panic(err)
	}

	//message_id := uuid.New().String()

	qry = fmt.Sprintf("INSERT INTO messages(id,inboxid,created, username, `from`,`to`, roomid, message) VALUES('%s','%s',now(),'%s','%s','%s','%s','%s')", messageid, inboxid, p_username, from_user, to_user, roomid, message)
	fmt.Println(`9882.7:`, func_name, `qry=`, qry)

	// perform a db.Query insert
	rows, err := db.Query(qry)
	if err != nil {
		fmt.Println(`9882.8:`, func_name, `err=`, err.Error())
		//return Map{`statusCode`:504}
		return false
	}

	defer rows.Close()

	///return Map{`statusCode`:200, `data`:message_id}
	fmt.Println(`9882.9:`, func_name, `SUCCESS!`)
	return true
	//}
	//return Map{`statusCode`:505}
}


func httpAddNewMessageResponse(p_username, p_from, p_to, roomid, message, atachment string) Map {

	func_name := `httpAddNewMessageResponse`
	fmt.Println(`10057.0 START`, func_name)

	if p_from == p_to {
		fmt.Println(`Sending message to same user is not allowed!`)
		return Map{`statusCode`: 500}
	}

	if p_from != p_username {
		fmt.Println(`p_username != from!`)
		return Map{`statusCode`: 501}
	}

	_, me := findUserByUsername(p_username)
	if me == nil {
		return Map{`statusCode`: 502}
	}

	_, from_user := findUserByUsername(p_from)
	if from_user == nil {
		return Map{`statusCode`: 502}
	}

	//from_username := from_user.Username
	//from_name := from_user.Name

	var to_users []User

	to_arr := strings.Split(p_to, ` `)
	for _, to := range to_arr {
		_, existing_to := findUserByUsername(to)
		if existing_to != nil {
			to_users = append(to_users, *existing_to)
		}
	}

	
	if len(to_users) == 0 {
		return Map{`statusCode`: 503}
	}

	messageid := uuid.New().String()
	for _, to_user := range to_users {

		//if sqlAddNewMessage(p_username,from_exten, from_exten,to_exten,roomid,message,atachment) {

		if sqlAddNewMessage(p_username, from_user.Username, from_user.Username, to_user.Username, from_user.Name, to_user.Name, roomid, message, atachment, messageid) {
			//if sqlAddNewMessage(existing_exten_to.Username,to_exten, from_exten,to_exten,roomid,message,atachment) {
			if sqlAddNewMessage(to_user.Username, to_user.Username, from_user.Username, to_user.Username, from_user.Name, to_user.Name, roomid, message, atachment, messageid) {
				
				
				is_online := false

				for _, r := range registered {
				
					if r.Username == to_user.Username {

				
						
						for _, c := range clients {
						
							if fmt.Sprintf("%s", c.conn.RemoteAddr()) == r.Remote_Addr {
								fmt.Println(`4695:`, func_name, to_user.Username, ` wsclient FOUND!`)
								str := fmt.Sprintf(`{"type":"NEW_MESSAGE_AVAILABLE"}`)
								err := c.conn.WriteMessage(1, []byte(str))
								if err != nil {
									fmt.Println("4699:", func_name, "WS FAILED send NEW_MESSAGE_AVAILABLE")
									fmt.Println("4699:", func_name, err)
									//return "", err
									is_online = true
								} else {
									fmt.Println("4699:", func_name, "WS OK send NEW_MESSAGE_AVAILABLE")
								}

							}
						}

					}
				}

				//send remote notif if needed
				if !is_online {

					notify_body := message
					if len(notify_body) > 50 {
						notify_body = notify_body[0:50]
					}
					//fmt.Println("17523.3:",func_name,"New message",notify_body)

					pushNotify(to_user.Username, config["name"], notify_body)

				}

				return Map{`statusCode`: 200}

			} else {
				fmt.Println(`16518.1: Something is wrong!`)
			}
		} else {
			fmt.Println(`16518.2: Something is wrong!`)
		}

		//return Map{`statusCode`:503}

	}

	return Map{`statusCode`: 200}

	//}
	//return Map{`statusCode`:505}
}

func httpPushNotify(username, notif_title, notif_body string) []map[string]string {
	func_name := `httpPushNotify()`
	fmt.Println(`3500: START httpRemoteNotifyResponse(username`, username)
	ret := []map[string]string{}
	_, user := findUserByUsername(username)
	if user == nil {
		return ret
	}

	fmt.Println(`3500: userextensions=`, userextensions)

	
	postBody := []byte(fmt.Sprintf(`{
						"apikey": "%s",
						"action": "PUSH_NOTIFY",
						"notif_to" : "%s",
						"notif_body": "%s",
						"notif_title": "%s"
						}`,
		config["apikey"],
		username,
		notif_body,
		notif_title))

	resp, err := http.Post("https://api.xoftphone.com/api/pbx", "application/json", bytes.NewBuffer(postBody))
	//Handle Error
	if err != nil {
		//log.Fatalf("511: An Error Occured %v", err)
		fmt.Println("511: An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Fatalln(err)
		fmt.Println(err)
	}
	sb := string(body)
	fmt.Println("4856: httpPushNotify() resp.Body=", sb)
	var sb_map Map
	json.Unmarshal(body, &sb_map)
	if sb_map["statusCode"] == 200 && sb_map["uuid"] != nil {
		//call_uuid := sb_map["uuid"]

	}

	//}

	//}
	//}
	//}

	fmt.Println(func_name, `6129: ret=`, ret)
	return ret
}

func httpPushNotifyData(username string, notif_data map[string]string) []map[string]string {
	func_name := `httpPushNotifyData()`
	fmt.Println(`3500: START httpPushNotifyData(username`, username)
	ret := []map[string]string{}
	_, user := findUserByUsername(username)
	if user == nil {
		return ret
	}

	
	
	str_data := `{`

	for k, v := range notif_data {
		str_data = fmt.Sprintf(`%s"%s":"%s"`, str_data, k, v)
	}
	str_data = fmt.Sprintf(`%s}`, str_data)

	postBody := []byte(fmt.Sprintf(`{
		"apikey": "%s",
		"action": "PUSH_NOTIFY_DATA",
		"notif_to" : "%s",
		"notif_data": "%s",
		
		}`,
		config["apikey"],
		username,
		str_data))

	resp, err := http.Post("https://api.xoftphone.com/api/pbx", "application/json", bytes.NewBuffer(postBody))
	//Handle Error
	if err != nil {
		//log.Fatalf("511: An Error Occured %v", err)
		fmt.Println("511: An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Fatalln(err)
		fmt.Println(err)
	}
	sb := string(body)
	fmt.Println("4856: httpPushNotifyData() resp.Body=", sb)
	var sb_map Map
	json.Unmarshal(body, &sb_map)
	if sb_map["statusCode"] == 200 && sb_map["uuid"] != nil {
		//call_uuid := sb_map["uuid"]

	}

	//}

	//}
	//}
	//}

	fmt.Println(func_name, `6220: ret=`, ret)
	return ret
}

func httpEmailNotify(to, subject, body string) []map[string]string {
	func_name := `httpEmailNotify()`
	fmt.Println(`3500: START httpEmailNotify(to`, to)
	ret := []map[string]string{}

	
	fmt.Println(`3500: userextensions=`, userextensions)

	postBody := []byte(fmt.Sprintf(`{
						"action": "XO_EMAIL_NOTIFY",
						"data":{
						"apikey": "%s",
						
						"to" : "%s",
						"body": "%s",
						"subject": "%s"
						}}`,
		config["apikey"],
		to,
		body,
		subject))

	resp, err := http.Post("https://api.xoftphone.com/api/pbx", "application/json", bytes.NewBuffer(postBody))
	//Handle Error
	if err != nil {
		//log.Fatalf("511: An Error Occured %v", err)
		fmt.Println("511: An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Fatalln(err)
		fmt.Println(err)
	}
	sb := string(body)
	fmt.Println("4856: httpPushNotify() resp.Body=", sb)
	var sb_map Map
	json.Unmarshal(resp_body, &sb_map)
	if sb_map["statusCode"] == 200 && sb_map["uuid"] != nil {
		//call_uuid := sb_map["uuid"]

	}

	//}

	//}
	//}
	//}

	fmt.Println(func_name, `6129: ret=`, ret)
	return ret
}

func registerXoServer() {
	func_name := `registerXoServer()`
	fmt.Println(`17292: registerXoServer() config=`, config)
	//fmt.Println(func_name,`3500: START httpRemoteNotifyResponse(username`,username)

	pb := fmt.Sprintf(`{"action": "REGISTER_XOSERVER",
		"hostname" : "%s",
		"name": "%s",
		"description": "%s",
		"metadata": "%s",
		"wifi_ssid" : "%s",
		"wifi_bssid" : "%s",
		"wifi_password" : "%s"
		}`,
		config["public_hostname"],
		config["name"],
		config["description"],
		config["metadata"],

		config["wifi_ssid"],
		config["wifi_bssid"],
		config["wifi_password"],
	//username,
	)
	fmt.Println(`17997: ************ pb=`, pb)
	postBody := []byte(pb)

	resp, err := http.Post(api_xoftswitch, "application/json", bytes.NewBuffer(postBody))
	//Handle Error
	if err != nil {
		//log.Fatalf("511: An Error Occured %v", err)
		fmt.Println(func_name, "511: An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Fatalln(err)
		fmt.Println(err)
	}
	sb := string(body)
	fmt.Println(func_name, "7639: resp.Body=", sb)
	var sb_map Map
	json.Unmarshal(body, &sb_map)
	//status,ok_status := sb_map[`statusCode`]
	m_apikey, ok_apikey := sb_map[`apikey`]
	m_kiv, ok_kiv := sb_map[`kiv`]
	//m_kiv_iv,ok_kiv_iv := sb_map[`kiv_iv`]
	if ok_apikey && ok_kiv {
		//if status == 200{
		//call_uuid := sb_map["uuid"]
		//ret = sb_map["apikey"].(string)
		arr_kiv := strings.Split(m_kiv.(string), " ")
		if len(arr_kiv) == 2 {
			apikey = m_apikey.(string)
			kiv_key = arr_kiv[0]
			kiv_iv = arr_kiv[1]
			config["apikey"] = apikey
		
		}

		//}

	}

	fmt.Println(func_name, `+++++++++++++++++++++++7412: apikey=`, apikey, `kiv_key=`, kiv_key, `kiv_iv=`, kiv_iv)
	//return ret
}

func httpGetUserNameByEmail(email string) map[string]string {
	func_name := `httpGetUserNameByEmail()`
	fmt.Println(func_name, `10154: START httpGetUserNameByEmail(email`, email)
	//ret := ""
	ret := map[string]string{}
	postBody := []byte(fmt.Sprintf(`{
		
		"action": "GET_USERNAME_BY_EMAIL",
		"email" : "%s"}`, email))

	//resp, err := http.Post("https://api.xoftphone.com/api/account", "application/json", bytes.NewBuffer(postBody))

	resp, err := http.Post("https://api.xoftphone.com/api/account", "application/json", bytes.NewBuffer(postBody))

	//Handle Error
	if err != nil {
		//log.Fatalf("511: An Error Occured %v", err)
		//fmt.Println(func_name,"511: An Error Occured %v", err)
		fmt.Println(string(colorRed), func_name, err.Error())
		fmt.Println(string(colorReset))

	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Fatalln(err)
		//fmt.Println(err)
		fmt.Println(string(colorRed), err.Error())
		fmt.Println(string(colorReset))
	}
	sb := string(body)
	//fmt.Println(func_name,"7639: resp.Body=",sb)
	fmt.Println(string(colorYellow), func_name, "7639: resp.Body=", sb)
	fmt.Println(string(colorReset))
	var sb_map Map
	json.Unmarshal(body, &sb_map)
	//status,ok_status := sb_map[`statusCode`]
	musername, ok_username := sb_map[`username`]
	mname, ok_name := sb_map[`name`]
	mphoto, ok_photo := sb_map[`photo`]

	//m_kiv_iv,ok_kiv_iv := sb_map[`kiv_iv`]
	if ok_username {

		ret[`username`] = musername.(string)
		if ok_name {
			ret[`name`] = mname.(string)

		}
		if ok_photo {
			ret[`photo`] = mphoto.(string)

		}

	}

	fmt.Println(string(colorYellow), func_name, "DONE ret=", ret)
	fmt.Println(string(colorReset))

	return ret
}

func httpGetUserName(p_username string) map[string]string {
	func_name := `httpGetUserName()`
	fmt.Println(func_name, `10154: START httpGetUserNameByEmail(email`, p_username)
	//ret := ""
	ret := map[string]string{}
	postBody := []byte(fmt.Sprintf(`{
		
		"action": "GET_USERNAME",
		"username" : "%s"}`, p_username))

	//resp, err := http.Post("https://api.xoftphone.com/api/account", "application/json", bytes.NewBuffer(postBody))

	resp, err := http.Post("https://api.xoftphone.com/api/account", "application/json", bytes.NewBuffer(postBody))

	//Handle Error
	if err != nil {
		//log.Fatalf("511: An Error Occured %v", err)
		//fmt.Println(func_name,"511: An Error Occured %v", err)
		fmt.Println(string(colorRed), func_name, err.Error())
		fmt.Println(string(colorReset))

	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Fatalln(err)
		//fmt.Println(err)
		fmt.Println(string(colorRed), err.Error())
		fmt.Println(string(colorReset))
	}
	sb := string(body)
	//fmt.Println(func_name,"7639: resp.Body=",sb)
	fmt.Println(string(colorYellow), func_name, "7639: resp.Body=", sb)
	fmt.Println(string(colorReset))
	var sb_map Map
	json.Unmarshal(body, &sb_map)
	//status,ok_status := sb_map[`statusCode`]
	memail, ok_email := sb_map[`email`]
	musername, ok_username := sb_map[`username`]
	mname, ok_name := sb_map[`name`]
	mphoto, ok_photo := sb_map[`photo`]

	//m_kiv_iv,ok_kiv_iv := sb_map[`kiv_iv`]
	if ok_username {

		ret[`username`] = musername.(string)
		if ok_email {

			ret[`email`] = memail.(string)

		}
		if ok_name {
			ret[`name`] = mname.(string)

		}
		if ok_photo {
			ret[`photo`] = mphoto.(string)

		}

	}

	fmt.Println(string(colorYellow), func_name, "DONE ret=", ret)
	fmt.Println(string(colorReset))

	return ret
}

func httpXoAdminSignupUserIfNeeded(username, password, email, firstname, lastname, dob, gender, address, phonenumber, imageurl, xoroleid string) map[string]string {
	/// password is optional, will be decided on the serve
	func_name := `httpXoAdminAddUserIfNeeded()`
	//fmt.Println(func_name,`10154: START httpGetUserNameByEmail(email`,p_username)
	//ret := ""
	ret := map[string]string{}
	
	
	postBody := []byte(fmt.Sprintf(`{
		"action": "XOADMIN_SIGNUPUSER_IFNEEDED",
		"data":{"username" : "%s",
		"password" : "%s",
		"email" : "%s",
		"firstName" : "%s",
		"lastName" : "%s",
		"dob" : "%s",
		"sex" : "%s",
		"address" : "%s",
		"phoneNumber" : "%s",
		"imageUrl" : "%s",
		
		"xoroleid" : "%s",
		"hostname" : "%s"}}`,

		username, password, email, firstname, lastname, dob, gender, address, phonenumber, imageurl, xoroleid, hostname))

	//resp, err := http.Post("https://api.xoftphone.com/api/account", "application/json", bytes.NewBuffer(postBody))

	resp, err := http.Post("https://api.xoftphone.com/api/account", "application/json", bytes.NewBuffer(postBody))

	//Handle Error
	if err != nil {
		//log.Fatalf("511: An Error Occured %v", err)
		//fmt.Println(func_name,"511: An Error Occured %v", err)
		fmt.Println(string(colorRed), func_name, err.Error())
		fmt.Println(string(colorReset))

	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Fatalln(err)
		//fmt.Println(err)
		fmt.Println(string(colorRed), err.Error())
		fmt.Println(string(colorReset))
	}
	sb := string(body)
	//fmt.Println(func_name,"7639: resp.Body=",sb)
	fmt.Println(string(colorYellow), func_name, "7639: resp.Body=", sb)
	fmt.Println(string(colorReset))

	return ret
}


func createQRSignin_dep(username, password, email, name, photoUrl string) {
	func_name := `createQRSignin()`
	//fmt.Println(func_name,`3500: START httpRemoteNotifyResponse(username`,username)
	//ret := []map[string]string{}
	//ret := ""

	postBody := []byte(fmt.Sprintf(`{
		
		"action": "CREATE_QR_SIGNIN",
		"data":{
		"hostname" : "%s",
		"username": "%s",
		"password": "%s",
		"email": "%s",
		"displayName": "%s",
		"photoUrl": "%s"
		}}`,
		config["public_hostname"],
		username,
		password,
		email,
		name,
		photoUrl,
	))

	resp, err := http.Post("https://api.xoftphone.com/api/account", "application/json", bytes.NewBuffer(postBody))
	//Handle Error
	if err != nil {
		//log.Fatalf("511: An Error Occured %v", err)
		fmt.Println(func_name, "511: An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Fatalln(err)
		fmt.Println(err)
	}
	sb := string(body)
	fmt.Println(func_name, "7639: resp.Body=", sb)
	var sb_map Map
	json.Unmarshal(body, &sb_map)
	status, ok_status := sb_map[`statusCode`]
	
	fmt.Println(func_name, `+++++++++++++++++++++++7412: status=`, status, `ok_status=`, ok_status)
	//return ret
}

func createQRSignin(username, password, email, roleid, extension, name, photoUrl, homepageurl, firstname, lastname, phonenumber string) {
	func_name := `createQRSignin()`
	//fmt.Println(func_name,`3500: START httpRemoteNotifyResponse(username`,username)
	//ret := []map[string]string{}
	//ret := ""

	postBody := []byte(fmt.Sprintf(`{
		
		"action": "CREATE_QR_SIGNIN",
		"data":{
		"hostname" : "%s",
		"username": "%s",
		"password": "%s",
		"email": "%s",
		"displayName": "%s",
		"photoUrl": "%s",
		"homepageurl": "%s",
		"wifi_ssid": "%s",
		"wifi_bssid": "%s",
		"wifi_username": "%s",
		"wifi_password": "%s",
		
		}}`,
		config["public_hostname"],
		username,
		password,
		email,
		name,
		photoUrl,
		homepageurl,
	))

	resp, err := http.Post("https://api.xoftphone.com/api/account", "application/json", bytes.NewBuffer(postBody))
	//Handle Error
	if err != nil {
		//log.Fatalf("511: An Error Occured %v", err)
		fmt.Println(func_name, "9845.1: An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Fatalln(err)
		fmt.Println(err)
	}

	sb := string(body)
	//fmt.Println(func_name,"7639: resp.Body=",sb)
	fmt.Println(string(colorYellow), func_name, "9845.2: resp.Body=", sb)
	fmt.Println(string(colorReset))
	var sb_map Map
	json.Unmarshal(body, &sb_map)
	status, ok_status := sb_map[`statusCode`]

	if ok_status && status.(float64) == 200 {
		fmt.Println(string(colorYellow), func_name, "9845.3:")
		fmt.Println(string(colorReset))
		//add user rquidilig Rom rquidilig@gmail.com admin 100
		//ret_user,err := doAdminAddUser(name,username, email,roleid,extension,true)
		ret_user, err := doAdminAddUser(firstname, lastname, name, phonenumber, photoUrl, username, email, roleid, extension, homepageurl, true)
		if err != nil {
			fmt.Println(string(colorRed), err.Error())
			fmt.Println(string(colorReset))

		} else {

			fmt.Println(`Created `, ret_user)

		}

	} else {
		fmt.Println(string(colorRed), func_name, "9845.3: FAILED")
		fmt.Println(string(colorReset))
	}

	//fmt.Println(func_name,`+++++++++++++++++++++++7412: apikey=`,apikey,`kiv_key=`,kiv_key,`kiv_iv=`,kiv_iv)
	fmt.Println(func_name, `+++++++++++++++++++++++7412: status=`, status, `ok_status=`, ok_status)
	//return ret
}



func createQRHostnames(to_email, hostnames, wifi_ssid, wifi_bssid, wifi_password string) {
	func_name := `createQRHostnames()`

	postBody := []byte(fmt.Sprintf(`{
		
		"action": "CREATE_QR_HOSTNAMES",
		"data":{
		"hostnames" : "%s",
		"email":"%s",
		"wifi_ssid": "%s",
		"wifi_bssid": "%s",
		"wifi_password": "%s"
		}}`,
		hostnames,
		to_email,
		wifi_ssid,
		wifi_bssid,
		wifi_password,
	))

	resp, err := http.Post("https://api.xoftphone.com/api/account", "application/json", bytes.NewBuffer(postBody))
	//Handle Error
	if err != nil {
		//log.Fatalf("511: An Error Occured %v", err)
		fmt.Println(func_name, "9845.1: An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Fatalln(err)
		fmt.Println(err)
	}

	sb := string(body)
	//fmt.Println(func_name,"7639: resp.Body=",sb)
	fmt.Println(string(colorYellow), func_name, "9845.2: resp.Body=", sb)
	fmt.Println(string(colorReset))
	var sb_map Map
	json.Unmarshal(body, &sb_map)
	status, ok_status := sb_map[`statusCode`]

	if ok_status && status.(float64) == 200 {
		fmt.Println(string(colorYellow), func_name, "QR code created.")
		fmt.Println(string(colorReset))
		
		

	} else {
		fmt.Println(string(colorRed), func_name, "9845.3: FAILED")
		fmt.Println(string(colorReset))
	}

	fmt.Println(func_name, `+++++++++++++++++++++++7412: status=`, status, `ok_status=`, ok_status)
	//return ret
}

func createQRWifi(email, hostname, wifi_ssid, wifi_password, location, wifi_bssid string) error {
	func_name := `createQRWifi()`

	
	pb := fmt.Sprintf(`{"action": "CREATE_QR_WIFI",
		"data":{"email": "%s",
		"hostname" : "%s",
		"wifi_ssid": "%s",
		"wifi_password": "%s",
		"location":"%s",
		"wifi_bssid":"%s"
		}}`,
		email,
		hostname,
		wifi_ssid,
		wifi_password,
		location,
		wifi_bssid)

	fmt.Println(`createQRWifi().162725 pb=`, pb)

	postBody := []byte(pb)

	resp, err := http.Post("https://api.xoftphone.com/api/account", "application/json", bytes.NewBuffer(postBody))
	//Handle Error
	if err != nil {
		//log.Fatalf("511: An Error Occured %v", err)
		fmt.Println(func_name, "9845.1: An Error Occured %v", err)
		return err
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Fatalln(err)
		fmt.Println(err)
		return err
	}

	sb := string(body)
	//fmt.Println(func_name,"7639: resp.Body=",sb)
	fmt.Println(string(colorYellow), func_name, "9845.2: resp.Body=", sb)
	fmt.Println(string(colorReset))
	var sb_map Map
	json.Unmarshal(body, &sb_map)
	status, ok_status := sb_map[`statusCode`]

	if ok_status && status.(float64) == 200 {
		fmt.Println(string(colorYellow), func_name, "QR code created.")
		fmt.Println(string(colorReset))
		
		

	} else {
		fmt.Println(string(colorRed), func_name, "9845.3: FAILED")
		fmt.Println(string(colorReset))
	}

	fmt.Println(func_name, `+++++++++++++++++++++++7412: status=`, status, `ok_status=`, ok_status)
	//return ret
	return nil
}

func wsINVITE(ws *websocket.Conn, result Map) (string, error) {
	fmt.Println("START: wsINVITE")

	
	

	return "", nil
}



func wsCKANSWER(ws *websocket.Conn, data Map) {

	//id,ok_id := data["id"].(string)
	uid, ok_uid := data["uuid"].(string)

	device_token, ok_device_token := data["device_token"].(string)

	fmt.Println("+++++++++++++ START: wsCKANSWER", "device", device_token, "uuid", uid)

	if ok_uid && ok_device_token {
		
		
		for idx2, ic2 := range incomingcalls {
			if ic2.Uuid == uid && ic2.Device_Token == device_token {
				incomingcalls[idx2].CKAnswerState = 1
				send_str := fmt.Sprintf(`{"type":"CK_ANSWER 200","uuid":"%s"}`, uid)
				fmt.Println("-> CK_ANSWER 200", device_token)

				err := ws.WriteMessage(1, []byte(send_str))
				if err != nil {
					fmt.Println("1331: wsCKANSWER", err, device_token)

				}
			}

		}

		
	} else {
		fmt.Println("!!!!!!!!!!! wsCKANSWER uuid and device_token required!")
	}

}

func wsSIPRing(ws *websocket.Conn, data Map) {
	fmt.Println("14280: START: wsSIPRing")
	fmt.Println("SIP RING data=", data)

	

}

func wsAMICoreShowChannels(ws *websocket.Conn) {
	fmt.Println("START: wsAMICoreShowChannels")
	amiCoreShowChannels(glo_ami)
	
	

}

func wsAMIPark(ws *websocket.Conn, data Map) {
	func_name := `wsAMIPark`
	fmt.Println(func_name, "START: data=", data)
	ch, ok_ch := data["channel"]
	if ok_ch {
		amiPark(glo_ami, ch.(string))

	} else {
		fmt.Println(func_name, `param channel required!`)
	}
	
	
}

func wsAMIATXFER(ws *websocket.Conn, data Map) {
	func_name := `wsAMIATXFER`
	fmt.Println(func_name, "START: data=", data)
	ch, ok_ch := data["channel"]
	exten, ok_exten := data["exten"]
	context, ok_context := data["context"]

	if ok_ch && ok_exten && ok_context {
		amiAtxFer(glo_ami, ch.(string), exten.(string), context.(string))

	} else {
		fmt.Println(func_name, `params required!`)
	}

}
func wsAMIMuteAudio(ws *websocket.Conn, data Map) {
	func_name := `wsAMIMuteAudio`
	fmt.Println(func_name, "START: data=", data)
	ch, ok_ch := data["channel"]
	direction, ok_direction := data["direction"]
	state, ok_state := data["state"]
	if ok_ch && ok_direction && ok_state {
		amiMuteAudio(glo_ami, ch.(string), direction.(string), state.(string))

	} else {
		fmt.Println(func_name, `Missing params!`)
	}

}

func wsAMIRedirect(ws *websocket.Conn, data Map) {
	func_name := `wsAMIRedirect`
	fmt.Println(func_name, "START: data=", data)

	channel, ok_channel := data["channel"]
	exten, ok_exten := data["exten"]
	context, ok_context := data["context"]

	extra_channel, ok_extra_ch := data["extra_channel"]
	extra_exten, ok_extra_exten := data["extra_exten"]
	extra_context, ok_extra_context := data["extra_context"]

	//ROM: All params are required, extras are optional but must provide empty string

	if ok_channel && ok_exten && ok_context && ok_extra_ch && ok_extra_exten && ok_extra_context {
		amiRedirect(glo_ami,
			channel.(string),
			exten.(string),
			context.(string),
			extra_channel.(string),
			extra_exten.(string),
			extra_context.(string))

	} else {
		fmt.Println(func_name, `params missing!`)
	}

}
func wsCKHangup(ws *websocket.Conn, data Map) {

	uid, ok_uid := data["uuid"].(string)
	//id,ok_id := data["id"].(string)
	device_token, ok_device_token := data["device_token"].(string)

	fmt.Println("1930 START: wsCKHangup device=", device_token, "uuid=", uid)

	if ok_uid && ok_device_token {
		//fmt.Println("1657 wsCKHangup")
		//results := findIncomingCallsByUuid(uid)

		//if results != nil && len(results) > 0{
		if len(incomingcalls) > 0 {
			fmt.Println("1662 wsCKHangup device=", device_token, "uuid=", uid)
			//for idx,r := range results{
			for idx, r := range incomingcalls {
				fmt.Println("1665 wsCKHangup device=", device_token, "uuid=", uid)
				if r.Uuid == uid {
					// Update state
					fmt.Println("1668: wsCKHangup() update", r.Uuid, "device=", device_token, "uuid=", uid)
					incomingcalls[idx].CKHangupState = 1
					
					

				}

				//}

			}
			fmt.Println("1696 wsCKHangup", "device=", device_token, "uuid=", uid)

			

		} else {
			fmt.Println("No device to HANGUP!", "device=", device_token, "uuid=", uid)
		}
		send_str := fmt.Sprintf(`{"type":"CK_HANGUP 200","uuid":"%s"}`, uid)
		fmt.Println("-> CK_HANGUP 200 device", device_token, "uuid=", uid)

		err := ws.WriteMessage(1, []byte(send_str))
		if err != nil {
			fmt.Println("1331: wsCKHangup", err)

		}

	} else {
		fmt.Println("wsHANGUP uuid and device_token required!")
		send_str := fmt.Sprintf(`{"type":"CK_HANGUP 200","uuid":"%s"}`, uid)
		fmt.Println("-> CK_HANGUP 500", device_token)

		err := ws.WriteMessage(1, []byte(send_str))
		if err != nil {
			fmt.Println("1331: wsCKHangup", err, device_token)

		}
	}

}

func wsAWAKE(ws *websocket.Conn, result Map) (string, error) {
	fmt.Println("START: wsAWAKE")

	if result["to"] == nil {
		return "", errors.New("500.0")
	}

	to := result["to"].(string)
	_, client_from := findClientByConn(ws)

	if client_from == nil {
		return "", errors.New("500.1")
	}

	//from := client_from.exten
	//client_to := findClientByExten(to)
	contacts := findContacts(to)

	for _, c := range contacts {
		/*if client_to == nil{
			return "", errors.New("500.2")
		}
		*/

		//invite_str := fmt.Sprintf(`{"type":"AWAKE","from":%s}`,from)
		if c.client != nil {
			invite_str := fmt.Sprintf(`{"type":"AWAKE"}`)
			err := c.client.conn.WriteMessage(1, []byte(invite_str))
			if err != nil {
				//fmt.Println("261: ",err)
				return "", err
			}

		}

	}

	//fmt.Println(time_now,"-> INVITE 200",ws.RemoteAddr())

	client_from.conn.WriteMessage(1, []byte("AWAKE 200"))

	
	return "", nil
}

func wsAWAKE_STATUS(ws *websocket.Conn, data Map) {
	fmt.Println("START: wsAWAKE_STATUS")
	
	

	status, ok_status := data["status"].(string)
	device_token, ok_device_token := data["device_token"].(string)

	fmt.Println("1206: ok_status", ok_status, "status=", status)
	fmt.Println("1206: ok_device_token=", ok_device_token, "device_token=", device_token)

	if ok_status && ok_device_token {
		println("1151: wsAWAKE_STATUS OK to proceed")
	}

	
	
	results := findAllRegisteredWithDeviceToken(device_token)

	count := len(results)
	fmt.Println("1171: wsAWAKE_STATUS findAllRegisteredWithDeviceToken count=", count)
	if count > 0 {
		fmt.Println("1175: wsAWAKE_STATUS BEFORE registered=", registered)
		for i, r := range results {
			fmt.Println("1175: wsAWAKE_STATUS r=", r.Awake_Status)
			
			registered[i].Awake_Status = status
			
			

		}
		fmt.Println("1175: wsAWAKE_STATUS AFTER registered=", registered)
		writeFileRegistered()
		dbSaveRegistered()
		//fmt.Println("1196: wsAWAKE_STATUS SAVED registered=",registered)

	}

	//fmt.Println(time_now,"-> INVITE 200",ws.RemoteAddr())

	ws.WriteMessage(1, []byte(`{"type":"AWAKE_STATUS 200"}`))

	/*
		err2 := client_from.conn.WriteMessage(1, []byte("AWAKE 200"))

		if err != nil {
			//fmt.Println("270: ",err)
			return "", err2
		}
	*/

	//return "",nil
}

func wsREGISTER(ws *websocket.Conn, data Map) { //(string, error){
	fmt.Println("START: wsREGISTER data=", data)

	
	device_token, ok_device_token := data["device_token"].(string)
	if ok_device_token {
		index_c, c := findClientByConn(ws)
		if c != nil {
			clients[index_c].device_token = device_token
		}

	}

	username, ok_username := data["username"].(string)
	if ok_device_token && ok_username {
		removeRegisteredWithTokenNotUsername(device_token, username)
	}

	_, userinfo := findUserByUsername(username)

	name, ok_name := data["name"].(string)
	if !ok_name {

		if userinfo != nil {
			name = userinfo.Name
		}
	}
	photo, ok_photo := data["photo"].(string)
	if !ok_photo {

		if userinfo != nil {
			photo = userinfo.Photo
		}
	}

	exten, ok_exten := data["exten"].(string)
	device_type, ok_device_type := data["device_type"].(string)

	awake_status, ok_awake_status := data["awake_status"].(string)
	//date,ok_date := data["date"].(string)
	contact_name, ok_contact_name := data["contact_name"].(string)
	contact_host, ok_contact_host := data["contact_host"].(string)
	contact_transport, ok_contact_transport := data["contact_transport"].(string)

	fmt.Println("1206: ok_exten", ok_exten, "exten=", exten)
	fmt.Println("1206: ok_device_token=", ok_device_token, "device_token=", device_token)
	fmt.Println("1206: ok_device_type=", ok_device_type, "device_type=", device_type)
	fmt.Println("1206: ok_username=", ok_username, "username=", username)
	fmt.Println("1209: ok_awake_status=", ok_awake_status, "awake_status=", awake_status)

	if !ok_contact_name {
		contact_name = ""
	}
	if !ok_contact_host {
		contact_host = ""
	}
	if !ok_contact_transport {
		contact_transport = ""
	}

	if ok_exten && ok_device_token && ok_device_type && ok_username {

		deleteRegisteredInvalidDeviceToken()
		index, r := findRegistered(exten, device_token)
		//r := findContact(exten,device_token)
		remote_addr := fmt.Sprintf("%s", ws.RemoteAddr())
		//if remote_addr == nil || len(remote_addr) == 0{}
		if r == nil {
			//create
			registered = append(registered, WSRegister{
				Date:         time.Now().Format(time.RFC3339),
				Exten:        exten,
				Device_Token: device_token,
				Device_Type:  device_type,
				Username:     username,
				Name:         name,
				Photo:        photo,

				Contact_Name:      contact_name,
				Contact_Host:      contact_host,
				Contact_Transport: contact_transport,
				Remote_Addr:       remote_addr,
				Register_Status:   1,
				Awake_Status:      awake_status})

		} else {

			registered[index].Date = time.Now().Format(time.RFC3339)

			registered[index].Username = username
			registered[index].Exten = exten
			registered[index].Device_Type = device_type
			registered[index].Device_Token = device_token
			registered[index].Contact_Name = contact_name
			registered[index].Contact_Host = contact_host
			registered[index].Contact_Transport = contact_transport
			registered[index].Remote_Addr = remote_addr
			registered[index].Register_Status = 1

			registered[index].Awake_Status = awake_status

		}

	} else {
		//return "", errors.New("500.1")
		fmt.Println("-> REGISTER 500", ws.RemoteAddr())
		ws.WriteMessage(1, []byte(`{"type":"REGISTER 200"}`))
	}

	fmt.Println("166: wsREGISTER clients= ", clients, "registrations= ", registered)
	//writeFileRegistered()
	dbSaveRegistered()

	//return "",nil
	fmt.Println("-> REGISTER 200", ws.RemoteAddr())
	ws.WriteMessage(1, []byte(`{"type":"REGISTER 200"}`))
}

func wsUNREGISTER(ws *websocket.Conn, result Map) (string, error) {
	fmt.Println("START: wsUNREGISTER")

	/*
		index,ret_client := findClientByConn(ws)
		if ret_client == nil{
			return "",errors.New("REGISTER 500.2")
		}
		fmt.Println("165: wsREGISTER ret_client= ", ret_client)
	*/
	//Remove old record

	exten, ok_exten := result["exten"].(string)
	device_token, ok_device_token := result["device_token"].(string)
	//device_type, ok_device_type := result["device_type"].(string)
	//username,ok_username := result["username"].(string)

	if ok_exten && ok_device_token {

		index, r := findRegistered(exten, device_token)
		//r := findContact(exten,device_token)
		//remote_addr := fmt.Sprintf("%s",ws.RemoteAddr())
		if r != nil {

			//registered[index].Username = username
			//registered[index].Exten = exten
			//registered[index].Device_Type = device_type
			//registered[index].Device_Token = device_token
			//registered[index].Remote_Addr = ""
			registered[index].Date = time.Now().Format(time.RFC3339)
			registered[index].Register_Status = 0

		}

	} else {
		return "", errors.New("500.1")
	}

	fmt.Println("166: wsUNREGISTER clients= ", clients, "registrations= ", registered)
	writeFileRegistered()
	dbSaveRegistered()

	return "", nil
}

//--> AMI

func wsMOREEXTENSIONS(ws *websocket.Conn, data Map) {
	fmt.Println("START: wsAMIMOREEXTENSIONS")
	q, ok_q := data["q"].(string)
	if !ok_q {
		q = ""
	}
	skip, ok_skip := data["skip"].(int)
	if !ok_skip {
		skip = 0
	}
	page_num, ok_page_num := data["page_num"].(int)
	if !ok_page_num {
		page_num = 1
	}
	fmt.Println("q=", q, "ok_q=", ok_q, "skip=", skip, "ok_skip=", ok_skip, "page_num=", page_num, "ok_page_num=", ok_page_num)
	page_count := 100
	//index := 1
	item_count := 0
	total := len(extensions)
	items := []Extension{}
	
	
	re_remove := regexp.MustCompile(`\<(.*)\>`)

	fmt.Println("extensions=", extensions)
	for i, v := range extensions {

		if item_count >= page_count {
			break
		}

		ma_remove := re_remove.FindStringSubmatch(v.Displayname)
		//fmt.Println(`ma_key=`,ma_key)
		if len(ma_remove) == 2 {
			//must key line

			r := ma_remove[1]

			v.Displayname = strings.Replace(v.Displayname, fmt.Sprintf(`<%s>`, r), ``, 1)
		}

		if skip > 0 {
			if i > skip {
				items = append(items, v)
				item_count++
			}

		} else {

			items = append(items, v)
			item_count++
		}

	}

	//}
	//else{

	//}
	//}

	fmt.Println("items=", items)
	json_items, err1 := json.Marshal(items)
	//fmt.Println("json_items=",json_items)
	if err1 == nil {
		ret_json := fmt.Sprintf(`{"total":%v,"page_num":%v,"items":%s}`, total, page_num, json_items)

		jsonString, err2 := json.Marshal(ret_json)
		//fmt.Println(err)
		if err2 == nil {
			//fmt.Println("-> MORE_EXTENSIONS 200",)
			jsonString = []byte(fmt.Sprintf(`{"type":"MORE_EXTENSIONS 200", "data":%s}`, jsonString))
			fmt.Println("-> MORE_EXTENSIONS 200", string(jsonString))
			ws.WriteMessage(1, []byte(jsonString))
		}
	}

}

func wsGETEXENSION(ws *websocket.Conn, data Map) {
	fmt.Println("START: wsGET_EXENSION")
	q, ok_q := data["q"].(string)
	if !ok_q {

	}

	
	var exten Extension
	for _, v := range extensions {

		if v.Number == q {
			exten = v
			break
		}

		//v.Displayname = strings.Replace(v.Displayname,fmt.Sprintf(`<%s>`,r),``,1)

	}

	
	
	json_item, err1 := json.Marshal(exten)
	//fmt.Println("json_items=",json_items)
	if err1 == nil {
		ret_json := fmt.Sprintf(`{"exten":%s}`, json_item)

		jsonString, err2 := json.Marshal(ret_json)
		//fmt.Println(err)
		if err2 == nil {
			//fmt.Println("-> MORE_EXTENSIONS 200",)
			jsonString = []byte(fmt.Sprintf(`{"type":"GET_EXTENSION 200", "data":%s}`, jsonString))
			fmt.Println("-> GET_EXTENSION 200", string(jsonString))
			ws.WriteMessage(1, []byte(jsonString))
		}
	}

}

func wsJOIN(ws *websocket.Conn, data Map) {
	fmt.Println("START: wsJOIN")
	
	

}

//<-- AMI

func unregister(conn *websocket.Conn) {

	for idx, r := range registered {
		//fmt.Println("30: idx=",idx,"c= ",c.conn.RemoteAddr())
		//if c == conn{
		if r.Remote_Addr == fmt.Sprintf("%s", conn.RemoteAddr()) {
			fmt.Println("31: index", idx, "FOUND")
			//index = idx
			registered[idx].Remote_Addr = ""

			//break
			//return

		}
	}

	writeFileRegistered()

	index := -1
	for idx, c := range clients {
		fmt.Println("30: idx=", idx, "c= ", c.conn.RemoteAddr())
		//if c == conn{
		if c.conn.RemoteAddr() == conn.RemoteAddr() {
			fmt.Println("31: index", idx, "FOUND")
			index = idx
			break

		}
	}

	if index > -1 {
		if len(clients) > 0 {

			fmt.Println("Before remove clients", len(clients))
			if len(clients) == 1 {
				clients = clients[:0]
			} else {

				clients = append(clients[:index], clients[index+1:]...)

			}
			fmt.Println("After remove clients", len(clients))

		}

	} else {
		fmt.Println("54: index", index, "NOT found")
	}
}

func removeClientsByDeviceTokenAndExten(device_token string, exten string) {
	index := -1
	for idx, r := range registered {
		//fmt.Println("30: idx=",idx,"c= ",c.conn.RemoteAddr())
		//if c == conn{
		//if c.conn.RemoteAddr() == conn.RemoteAddr(){
		if r.Device_Token == r.Device_Token && r.Exten == exten {
			fmt.Println("738: index", idx, "FOUND")
			index = idx
			break

		}
	}
	if index > -1 {
		if len(clients) > 0 {

			fmt.Println("Before remove clients", len(clients))
			if len(clients) == 1 {
				clients = clients[:0]
			} else {

				clients = append(clients[:index], clients[index+1:]...)

			}
			fmt.Println("After remove clients", len(clients))

		}

	} else {
		fmt.Println("763: index", index, "NOT found")
	}
}

func wsPING(exten string) {
	fmt.Println("START wsPING", exten)

	for _, r := range registered {

		if r.Exten == exten {
			fmt.Println("wsPING found registered", exten)

			for _, c := range clients {

				if fmt.Sprintf("%s", c.conn.RemoteAddr()) == r.Remote_Addr {
					fmt.Println("wsPING found client", exten)
					err := c.conn.WriteMessage(1, []byte(fmt.Sprintf(`{"type":"PING"}`)))
					if err != nil {
						fmt.Println("87: ", err)
					}
				}

			}

		}
	}

}



func register(ws *websocket.Conn) {
	//clients = append(clients, *ws)

	//c := WSClient{conn: ws, username:"",   exten: "",device_type:"",device_token:""}
	c := WSClient{conn: ws, device_token: ""}

	clients = append(clients, c)
	fmt.Println("224: connected:", c.conn.RemoteAddr())

	/*
		fmt.Println("Client Connected")

		err := ws.WriteMessage(1, []byte("Hi Client!"))
		if err != nil {
			fmt.Println("87: %v",err)
		}
	*/
	//PING := Map{"type":"PING"}
	fmt.Println("-> PING", ws.RemoteAddr())
	//err := ws.WriteMessage(1, []byte("PING"))
	err := ws.WriteMessage(1, []byte(fmt.Sprintf(`{"type":"PING"}`)))
	if err != nil {
		fmt.Println("87:", err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	reader(ws)

}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint

func reader(ws *websocket.Conn) {
	//fmt.Println("23: conn= %v",conn)
	for {
		// read in a message
		msgType, msg, err := ws.ReadMessage()
		//_, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("27: ", err)
			unregister(ws)
			return
		}
		msg_str := string(msg)
		//time_now := time.Now().Format(time.Stamp)
		//fmt.Println(time_now, "<-", msg_str, ws.RemoteAddr())
		fmt.Println("<-", msg_str, ws.RemoteAddr())

		var result Map
		json.Unmarshal([]byte(msg), &result)

		
		

		switch result["type"] {
		case "PONG":

			if ping_enabled {
				timer1 := time.NewTimer(5 * time.Second)
				<-timer1.C
				fmt.Println("-> PING", ws.RemoteAddr())
				//err := ws.WriteMessage(msgType, []byte("PING"))
				err := ws.WriteMessage(msgType, []byte(fmt.Sprintf(`{"type":"PING"}`)))
				if err != nil {
					fmt.Println("87: ", err)
				}
			}

		case "REGISTER":
			wsREGISTER(ws, result)

		case "UNREGISTER":

			//fmt.Println(time_now,"-> REGISTER 200",ws.RemoteAddr())
			_, err := wsUNREGISTER(ws, result)
			if err == nil {
				fmt.Println("-> UNREGISTER 200", ws.RemoteAddr())
				ws.WriteMessage(msgType, []byte(`{"type":"UNREGISTER 200"}`))
			} else {
				fmt.Println("-> UNREGISTER 500", ws.RemoteAddr())
				ws.WriteMessage(msgType, []byte(`{"type":"UNREGISTER 500"}`))

			}

		case "INVITE":
			_, err = wsINVITE(ws, result)

			if err != nil {
				fmt.Println("INVITE ", err)
				ws.WriteMessage(msgType, []byte(`{"type":"INVITE 500"}`))
			} else {
				ws.WriteMessage(msgType, []byte(`{"type":"INVITE 200"}`))
			}

		case "AWAKE_STATUS":
			wsAWAKE_STATUS(ws, result)

		case "CK_ANSWER":

			wsCKANSWER(ws, result)
		case "SIP_RING":

			wsSIPRing(ws, result)
		case "AMI_CORESHOWCHANNELS":

			wsAMICoreShowChannels(ws)

		case "AMI_PARK":

			wsAMIPark(ws, result)

		case "AMI_ATXFER":

			wsAMIATXFER(ws, result)

		case "AMI_MUTEAUDIO":

			wsAMIMuteAudio(ws, result)

		case "AMI_REDIRECT":

			wsAMIRedirect(ws, result)

		case "CK_HANGUP":

			//fmt.Println(time_now,"-> REGISTER 200",ws.RemoteAddr())
			wsCKHangup(ws, result)

		case "MORE_EXTENSIONS":
			wsMOREEXTENSIONS(ws, result)

		case "GET_EXTENSION":
			wsGETEXENSION(ws, result)

			//case "JOIN":
			//	wsJOIN(ws,result)

		}

	}
}

func isAuthenticated(username, sessionid string) bool {
	_, user := findUserByUsername(username)

	if user != nil {
		// let's verify
		if len(user.Sessionid) == 0 {
			postBody := []byte(fmt.Sprintf(`{"action":"XO_GET_SESSIONID", "username":"%s"}`, username))

			resp, err := http.Post("https://api.xoftphone.com/api/account", "application/json", bytes.NewBuffer(postBody))
			//Handle Error
			if err != nil {
				//log.Fatalf("511: An Error Occured %v", err)
				fmt.Println("511: An Error Occured %v", err)
				return false
			}
			defer resp.Body.Close()
			//Read the response body
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				//log.Fatalln(err)
				fmt.Println(err)
			}
			sb := string(body)
			fmt.Println("7422: isAuthenticated() resp.Body=", sb)
			var sb_map Map
			json.Unmarshal(body, &sb_map)
			if sb_map["statusCode"] == 200 && sb_map["session"] != nil {
				//call_uuid := sb_map["uuid"]

			}

		}
		if user != nil && user.Sessionid == sessionid {

			return true
		}
		return true
	}

	return false
}

func remoteDecrypt(encrypted_data string) Map {
	//user := findUserByUsername(username)

	if len(encrypted_data) > 0 {
		//if user != nil {
		// let's verify
		//if len(user.Sessionid) == 0{
		postBody := []byte(fmt.Sprintf(`{"action":"DECRYPT", "data":"%s"}`, encrypted_data))

		resp, err := http.Post("https://api.xoftphone.com/api/xoutils", "application/json", bytes.NewBuffer(postBody))
		//Handle Error
		if err != nil {
			//log.Fatalf("511: An Error Occured %v", err)
			fmt.Println("511: An Error Occured", err)
			return nil
		}
		defer resp.Body.Close()
		//Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//log.Fatalln(err)
			fmt.Println(err)
		}
		//sb := string(body)
		fmt.Println("7466: remoteDecrypt() body=", string(body))
		//4445: remoteDecrypt() resp.Body= {"type":"join","hostname":"fpbx.sharecle.com","username":"rquidilig2","device_token":"fPSwo29OTVKXB02org0A7a:APA91bGijDKntRYq-yJPsMZN6NT-0eyDi7UbxwTflkbn7Z4jiRZ6eLN1jsIBNhn1R7uUXo6cDEnnutmC2m-Z1HfTg82d7oFXqCP4TCtQv_EgNvHcVxZEgS0Eet4oqdBBDdnwYACSXkOY","sessionid":"80b2a8a8-cb28-4531-9894-54a09b5ae1c8"}
		var map_body Map
		json.Unmarshal(body, &map_body)

		return map_body

		/*
			   status_code, ok_status_code := map_body["statusCode"]
			   if ok_status_code && status_code == 200{
					data, ok_data := map_body["data"]
					if ok_data && len(data.(string)) > 0{
						return data.(string)
					}
			   }
		*/

	}

	return nil
}

func homePage_dep(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Println("13949: It works!")
	fmt.Fprintf(w, "<h1>200 OK</h1>")
}



func loadTemplates() {
	funcMap := template.FuncMap{
		"humanTime": humanTime,
	}

	templates = make(map[string]*template.Template)

	pages := []string{
		"dashboard",
		"users", "adduser", "edituser",
		"extensions", "addextension", "editextension",
		"userextensions", "adduserextension", "edituserextension",
		"settings",
		"about",
		"joinrequests", // ← add this
	}

	for _, page := range pages {
		templates[page] = template.Must(template.New(page).Funcs(funcMap).ParseFiles(
			"templates/layout.html",
			"templates/"+page+".html",
		))
	}

	templates["login"] = template.Must(template.ParseFiles("templates/login.html"))
}






func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
    tmpl, ok := templates[name]
    if !ok {
        http.Error(w, "Template not found: "+name, http.StatusInternalServerError)
        return
    }
    err := tmpl.ExecuteTemplate(w, "layout", data) // or "layout.html" depending on your define
    if err != nil {
        http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
    }
}

type DashboardData struct {
	JoinRequests []JoinRequest
    Users    []User
    Page     int
    PrevPage int
    NextPage int
    HasPrev  bool
    HasNext  bool
    Search   string
}

type LoginPageData struct {
    Message  string
    HideMenu bool
}




type DashboardPageData struct {
    Version  string
    HideMenu bool
}




func dashboardPage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("DEBUG: dashboardPage called for", r.URL.Path)

    if !requireAdminOrLogin(w, r) {
        fmt.Println("DEBUG: not authenticated, login page should have been rendered")
        return
    }

    data := DashboardPageData{
        Version:  Version,
        HideMenu: false,
    }

    if tpl, ok := templates["dashboard"]; ok && tpl != nil {
        fmt.Println("DEBUG: Rendering dashboard template")
        if err := tpl.ExecuteTemplate(w, "layout", data); err != nil {
            fmt.Println("TEMPLATE ERROR:", err)
            http.Error(w, "Template render error: "+err.Error(), http.StatusInternalServerError)
        }
    } else {
        fmt.Println("ERROR: templates[\"dashboard\"] is nil or not found")
        http.Error(w, "Dashboard template not loaded", http.StatusInternalServerError)
    }
}


func contactUsPage(w http.ResponseWriter, r *http.Request) {
    /*
	if !requireAdminOrLogin(w, r) {
        return
    }
		*/

    data := struct{}{} // add fields if needed
    err := templates["contactus"].ExecuteTemplate(w, "layout", data)
    if err != nil {
        http.Error(w, "Template render error: "+err.Error(), http.StatusInternalServerError)
    }
}

type AboutData struct {
    AppName   string
    Version   string
    //Commit    string
   // BuildDate string
    GoVersion string
    OSArch    string
    //Uptime    string
    DocsURL   string
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
    /*
    if !requireAdminOrLogin(w, r) {
        return
    }
    */

    data := AboutData{
        AppName:   "XoftSwitch",
        Version:   Version,
        //Commit:    Commit,
        //BuildDate: BuildDate,
        GoVersion: runtime.Version(),
        OSArch:    runtime.GOOS + "/" + runtime.GOARCH,
        //Uptime:    time.Since(startTime).Truncate(time.Second).String(),
        DocsURL:   "https://www.xoftswitch.com",
    }

    // Template key should match whatever you used when parsing (e.g., "about")
    if err := templates["about"].ExecuteTemplate(w, "layout", data); err != nil {
        http.Error(w, "Template render error: "+err.Error(), http.StatusInternalServerError)
        return
    }
}



func statusAPI(w http.ResponseWriter, r *http.Request) {
    cpu, mem, err := getXoftSwitchUsage()

    // Get network rates
    netIn, netOut := getNetworkUsage()

    // Get disk usage
    diskUsed, diskTotal, diskPercent := getDiskUsage("/")

    type Resp struct {
        CPUPercent float64 `json:"cpu"`
        MemMB      float32 `json:"mem"`
        NetInKBs   float64 `json:"net_in"`
        NetOutKBs  float64 `json:"net_out"`
        DiskUsedGB float64 `json:"disk_used"`
        DiskTotalGB float64 `json:"disk_total"`
        DiskUsagePercent float64 `json:"disk_usage"`
        Time       int64   `json:"time"`
        Error      string  `json:"error,omitempty"`
    }

    resp := Resp{
        Time:            time.Now().Unix(),
        NetInKBs:        netIn,
        NetOutKBs:       netOut,
        DiskUsedGB:      diskUsed,
        DiskTotalGB:     diskTotal,
        DiskUsagePercent: diskPercent,
    }

    if err != nil {
        resp.Error = err.Error()
    } else {
        resp.CPUPercent = cpu
        resp.MemMB = mem
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}





// --------- Helper to get CPU/Mem ---------
func getXoftSwitchUsage() (float64, float32, error) {
	procs, err := process.Processes()
	if err != nil {
		return 0, 0, err
	}

	for _, p := range procs {
		name, _ := p.Name()
		if name == "xoftswitch" { // match your binary name
			// warm-up call
			p.CPUPercent()
			time.Sleep(500 * time.Millisecond)
			cpuPercent, _ := p.CPUPercent()
			memInfo, _ := p.MemoryInfo()
			memMB := float32(memInfo.RSS) / 1024.0 / 1024.0
			return cpuPercent, memMB, nil
		}
	}
	return 0, 0, fmt.Errorf("xoftswitch process not found")
}

func getNetworkUsage() (inRateKBs, outRateKBs float64) {
    ioStats, err := gnet.IOCounters(false) // ✅ use gnet here
    if err != nil || len(ioStats) == 0 {
        return 0, 0
    }

    now := time.Now()
    if lastNetTime.IsZero() {
        lastNetIn = ioStats[0].BytesRecv
        lastNetOut = ioStats[0].BytesSent
        lastNetTime = now
        return 0, 0
    }

    delta := now.Sub(lastNetTime).Seconds()
    if delta <= 0 {
        return 0, 0
    }

    inRateKBs = float64(ioStats[0].BytesRecv-lastNetIn) / 1024.0 / delta
    outRateKBs = float64(ioStats[0].BytesSent-lastNetOut) / 1024.0 / delta

    lastNetIn = ioStats[0].BytesRecv
    lastNetOut = ioStats[0].BytesSent
    lastNetTime = now

    return
}

func getDiskUsage(path string) (usedGB, totalGB, usedPercent float64) {
    usage, err := disk.Usage(path)
    if err != nil {
        return 0, 0, 0
    }
    usedGB = float64(usage.Used) / 1024 / 1024 / 1024
    totalGB = float64(usage.Total) / 1024 / 1024 / 1024
    usedPercent = usage.UsedPercent
    return
}



func requireAdminOrLogin(w http.ResponseWriter, r *http.Request) bool {
    // ✅ Check if admin_auth cookie is present
    cookie, err := r.Cookie("admin_auth")
    if err == nil && cookie.Value == "true" {
        return true // Already authenticated
    }

    // ✅ If POST, process login
    if r.Method == http.MethodPost {
        user := r.FormValue("username")
        pass := r.FormValue("password")

        if user == config["admin_username"] && pass == config["admin_password"] {
            Log(LOG_DEBUG, "requireAdminOrLogin() OK!")
            http.SetCookie(w, &http.Cookie{
                Name:     "admin_auth",
                Value:    "true",
                Path:     "/",
                HttpOnly: true,
            })
            http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
            return false
        }

        Log(LOG_DEBUG, "requireAdminOrLogin() FAILED!!!")
        if tmpl, ok := templates["login"]; ok {
            tmpl.ExecuteTemplate(w, "login", struct{ Message string }{
                Message: "Invalid username or password",
            })
        } else {
            http.Error(w, "Login template not found", http.StatusInternalServerError)
        }
        return false
    }

    // ✅ If GET, just show login page
    if tmpl, ok := templates["login"]; ok {
        tmpl.ExecuteTemplate(w, "login", nil)
    } else {
        http.Error(w, "Login template not found", http.StatusInternalServerError)
    }
    return false
}

func loginPage(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        // Show the login page (no layout, standalone)
        err := templates["login"].ExecuteTemplate(w, "login", nil)
        if err != nil {
            http.Error(w, "Template render error: "+err.Error(), http.StatusInternalServerError)
        }
        return
    }

    if r.Method == http.MethodPost {
        user := r.FormValue("username")
        pass := r.FormValue("password")

        if user == config["admin_username"] && pass == config["admin_password"] {
            // Set cookie to mark user as logged in
            http.SetCookie(w, &http.Cookie{
                Name:     "admin_auth",
                Value:    "true",
                Path:     "/",
                HttpOnly: true,
            })
            // Redirect to dashboard or original page
            http.Redirect(w, r, "/", http.StatusSeeOther)
            return
        }

        // Login failed: re-show login page with error message
        err := templates["login"].ExecuteTemplate(w, "login", LoginPageData{
            Message: "Invalid username or password",
        })
        if err != nil {
            http.Error(w, "Template render error: "+err.Error(), http.StatusInternalServerError)
        }
        return
    }

    // For other methods, just respond 405 Method Not Allowed
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}



func usersPage(w http.ResponseWriter, r *http.Request) {
    // Require admin login
    if !requireAdminOrLogin(w, r) {
        return
    }

    // Pagination
    pageSize := 10
    page := 1
    if p := r.URL.Query().Get("page"); p != "" {
        if pi, err := strconv.Atoi(p); err == nil && pi > 0 {
            page = pi
        }
    }
    offset := (page - 1) * pageSize

    // Optional search
    search := r.URL.Query().Get("search")

    var (
        rows *sql.Rows
        err  error
    )

    if search != "" {
        query := fmt.Sprintf(`
            SELECT username, name, firstname, lastname, email, phonenumber, roleid, status
            FROM users
            WHERE username LIKE ?
            LIMIT %d OFFSET %d`, pageSize, offset)
        rows, err = db.Query(query, "%"+search+"%")
    } else {
        query := fmt.Sprintf(`
            SELECT username, name, firstname, lastname, email, phonenumber, roleid, status
            FROM users
            LIMIT %d OFFSET %d`, pageSize, offset)
        rows, err = db.Query(query)
    }

    if err != nil {
        http.Error(w, "DB query error: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var u User
        if err := rows.Scan(&u.Username, &u.Name, &u.Firstname, &u.Lastname, &u.Email, &u.Phonenumber, &u.Roleid, &u.Status); err != nil {
            http.Error(w, "Scan error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        users = append(users, u)
    }

    data := DashboardData{
        Users:    users,
        Page:     page,
        PrevPage: page - 1,
        NextPage: page + 1,
        HasPrev:  page > 1,
        HasNext:  len(users) == pageSize,
        Search:   search,
    }

	/*
    if err := tmpl.ExecuteTemplate(w, "users.html", data); err != nil {
        http.Error(w, "Template render error: "+err.Error(), http.StatusInternalServerError)
    }
	*/
	//renderTemplate(w, "users", data)
	err = templates["users"].ExecuteTemplate(w, "layout", data)
    if err != nil {
        http.Error(w, "Template render error: "+err.Error(), 500)
    }
}

func extensionsPageHandler(w http.ResponseWriter, r *http.Request) {
    if !requireAdminOrLogin(w, r) {
        return
    }

    rows, err := db.Query(`SELECT number, displayname, status, contacts, source, secret, devicetype, public, canblock, canunlist, canaddlist, candial, ver FROM extensions`)
    if err != nil {
        http.Error(w, "Query error: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var list []Extension
    for rows.Next() {
        var e Extension
        if err := rows.Scan(&e.Number, &e.Displayname, &e.Status, &e.Contacts, &e.Source, &e.Secret, &e.Devicetype, &e.Public, &e.Canblock, &e.Canunlist, &e.Canaddlist, &e.Candial, &e.Ver); err != nil {
            http.Error(w, "Scan error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        list = append(list, e)
    }

    renderTemplate(w, "extensions", list)
}

// normalizeBoolField ensures only "1" or "0" are saved
func normalizeBoolField(v string) string {
    if strings.TrimSpace(v) == "1" {
        return "1"
    }
    return "0"
}

func to01(v string, def string) string {
	// normalize to "1" or "0"
	switch strings.ToLower(strings.TrimSpace(v)) {
	case "1", "yes", "true", "on":
		return "1"
	case "0", "no", "false", "off":
		return "0"
	default:
		if def == "1" {
			return "1"
		}
		return "0"
	}
}

func toYesNo(v string, def string) string {
	// normalize to "yes" or "no"
	switch strings.ToLower(strings.TrimSpace(v)) {
	case "yes", "1", "true", "on":
		return "yes"
	case "no", "0", "false", "off":
		return "no"
	default:
		if strings.ToLower(def) == "yes" {
			return "yes"
		}
		return "no"
	}
}
type GeneratedExt struct {
    ExtNumber   int
    DisplayName string
    Email       string
    Secret      string
}

func editExtensionPage(w http.ResponseWriter, r *http.Request) {
	if !requireAdminOrLogin(w, r) {
		return
	}

	extension := r.URL.Query().Get("number")
	if extension == "" {
		http.Error(w, "21379: Missing extension number", http.StatusBadRequest)
		return
	}
	
	

	if r.Method == http.MethodPost {
		display  := strings.TrimSpace(r.FormValue("displayname"))
		email    := strings.TrimSpace(r.FormValue("email"))
		secret   := strings.TrimSpace(r.FormValue("secret"))
		//media_encryption   := strings.TrimSpace(r.FormValue("media_encryption"))
		template_path := "/etc/xoftswitch/csv/extension-tpl.csv"
		media_encryption := "no" //or dtls
		webrtc    := toYesNo(r.FormValue("webrtc"), "no")// default Yes
		if webrtc == "yes"{
			media_encryption = "dtls"
			template_path = "/etc/xoftswitch/csv/extension-webrtc-tpl.csv"
		}

		// Keep select-values consistent with templates:
		public    := toYesNo(r.FormValue("public"), "no")     // default Yes
		direct_media    := toYesNo(r.FormValue("direct_media"), "no")
		canblock  := toYesNo(r.FormValue("canblock"), "no")   // default Yes
		canunlist := toYesNo(r.FormValue("canunlist"), "no")  // default Yes
		canaddlist:= toYesNo(r.FormValue("canaddlist"), "no") // default Yes
		candial   := toYesNo(r.FormValue("candial"), "no")    // default Yes
		

		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Minute)
		defer cancel()

		opts := addexts.Options{
			TemplatePath: template_path,
			//OutPath:      "/etc/xoftswitch/bulk_webrtc_import.csv",
			OutPath:      "/etc/xoftswitch/csv/bulk_import.csv",
			Extension:        extension,
			//End:          extension,
			NamePattern:	display,
			EmailPattern:	email,
			Secret: secret,
			WebRtc: webrtc,
			DirectMedia: direct_media,
			MediaEncryption: media_encryption,

			//NamePattern:  "xo {ext}",
			//EmailPattern: "xo{ext}@xoftphone.com",
			//Secret: secret,
			//SecretBytes:  16,
	
			// side effects:
			FwconsolePath: "fwconsole",
			DoImport:      true,
			DoReload:      true,
			Timeout:       5 * time.Minute, // per fwconsole command
	
			Logger: log.Default(),
		}
		/*
		if err := addexts.Generate(ctx, opts); err != nil {
			http.Error(w, "generation/import failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		*/
		g, err := addexts.Generate(ctx, opts)
		if err != nil {
			http.Error(w, "generation/import failed: "+err.Error(), http.StatusInternalServerError)
			return
		}

		//for _, g := range list {
			//fmt.Printf("%d -> %s, %s, %s\n", g.ExtNumber, g.DisplayName, g.Email, g.Secret)
			_, err_db := db.Exec(`UPDATE extensions
			SET displayname=?, email=?, secret=?, webrtc=?,media_encryption=?, direct_media=?, public=?, canblock=?, canunlist=?, canaddlist=?, candial=? 
			WHERE extension=?`,
			g.DisplayName, g.Email, g.Secret, g.WebRtc,g.MediaEncryption,g.DirectMedia, public, canblock, canunlist, canaddlist, candial, g.Extension)
			if err_db != nil {
				//http.Error(w, "Update error: "+err.Error(), http.StatusInternalServerError)
				//return
				fmt.Printf("%d Update error: %s", g.Extension,err.Error())

			}

			

		//}

		amiGetConfigJson(glo_ami);

		

		// If these side-effects are desired on edit as well, keep them:
		//amiGetConfigJson(glo_ami)
		reloadXoftSwitch()

		http.Redirect(w, r, "/extensions", http.StatusSeeOther)
		return
	}

	

	var e Extension
	err_db := db.QueryRow(`
	SELECT number, displayname, secret, email, devicetype,
			public, webrtc, canblock, canunlist, canaddlist, candial
	FROM extensions WHERE number=?`, extension).
	Scan(&e.Number, &e.Displayname, &e.Secret, &e.Email, &e.Devicetype,
		&e.Public, &e.Webrtc, &e.Canblock, &e.Canunlist, &e.Canaddlist, &e.Candial)

	if err_db == sql.ErrNoRows {
		// 1) Ensure a row exists, using table defaults.
		//    This is atomic and idempotent: if the row already exists, it's a no-op.
		if _, err := db.Exec(`
			INSERT INTO extensions (number)
			VALUES (?)
			ON DUPLICATE KEY UPDATE number=number
		`, extension); err != nil {
			http.Error(w, "DB insert error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// 2) Read the row back (now guaranteed to exist)
		if err := db.QueryRow(`
		SELECT number, displayname, secret,email, devicetype,
				public, webrtc, canblock, canunlist, canaddlist, candial
		FROM extensions WHERE number=?`, extension).
			Scan(&e.Number, &e.Displayname, &e.Secret, &e.Email, &e.Devicetype,
				&e.Webrtc,&e.Directmedia,  &e.Public, &e.Canblock, &e.Canunlist, &e.Canaddlist, &e.Candial); err != nil {
			http.Error(w, "DB read-after-insert error: "+err.Error(), http.StatusInternalServerError)
			return
		}
	} else if err_db != nil {
		http.Error(w, "21661: DB error: "+err_db.Error(), http.StatusInternalServerError)
		return
	}

	// `e` now has the values (existing or just-created defaults).


	renderTemplate(w, "editextension", e)
}




func addExtensionPage(w http.ResponseWriter, r *http.Request) {
    if !requireAdminOrLogin(w, r) {
        return
    }

    if r.Method == http.MethodPost {
        
		extStartStr := strings.TrimSpace(r.FormValue("extstart"))
    if extStartStr == "" {
        http.Error(w, "Starting extension number is required", http.StatusBadRequest)
        return
    }

    extEndStr := strings.TrimSpace(r.FormValue("extend"))
    if extEndStr == "" {
        extEndStr = extStartStr // default end = start
    }

    // Parse to ints
    extStart, err := strconv.Atoi(extStartStr)
    if err != nil || extStart <= 0 {
        http.Error(w, "Invalid starting extension number", http.StatusBadRequest)
        return
    }
    extEnd, err := strconv.Atoi(extEndStr)
    if err != nil || extEnd < extStart {
        http.Error(w, "Invalid ending extension number", http.StatusBadRequest)
        return
    }


		email := strings.TrimSpace(r.FormValue("email"))
		secret   := strings.TrimSpace(r.FormValue("secret"))
		
		
        /*
		if email == "" {
            http.Error(w, "Number is required", http.StatusBadRequest)
            return
        }
			*/

		
		//secret := strings.TrimSpace(r.FormValue("secret"))
		/*
		// Generate a random 16-byte secret if not provided
		
		if secret == "" {
			b := make([]byte, 16)
			if _, err := rand.Read(b); err != nil {
				fmt.Println("Failed to generate secret:", err)
				http.Error(w, "Failed to generate secret", http.StatusInternalServerError)
				return
			}
			secret = hex.EncodeToString(b)
		}
		*/

        display := strings.TrimSpace(r.FormValue("displayname"))
        
		
		// Dropdown values: "yes"/"no" with defaults
		template_path := "/etc/xoftswitch/csv/extension-tpl.csv"
		media_encryption := "no" //or dtls

		webrtc    := toYesNo(r.FormValue("webrtc"), "no")// default Yes
		if webrtc == "yes"{
			media_encryption = "dtls"
			template_path = "/etc/xoftswitch/csv/extension-webrtc-tpl.csv"
		}
		direct_media    := toYesNo(r.FormValue("direct_media"), "yes")
		//media_encryption   := strings.TrimSpace(r.FormValue("media_encryption"))
		public    := toYesNo(r.FormValue("public"), "yes")     // default Yes
		
		canblock  := toYesNo(r.FormValue("canblock"), "yes")   // default Yes
		canunlist := toYesNo(r.FormValue("canunlist"), "no")  // default Yes
		canaddlist:= toYesNo(r.FormValue("canaddlist"), "no") // default Yes
		candial   := toYesNo(r.FormValue("candial"), "no")    // default Yes

		
		
		// Optional: cap total runtime for this request (independent of per-command timeout)
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Minute)
		defer cancel()

		opts := addexts.OptionsRange{
			TemplatePath: template_path,
			//OutPath:      "/etc/xoftswitch/bulk_webrtc_import.csv",
			OutPath:      "/etc/xoftswitch/csv/bulk_import.csv",
			Start:        extStart,
			End:          extEnd,
			Secret: secret,
			NamePattern:	display,
			EmailPattern:	email,
			WebRtc: webrtc,
			DirectMedia: direct_media,
			MediaEncryption: media_encryption,
			//NamePattern:  "xo {ext}",
			//EmailPattern: "xo{ext}@xoftphone.com",
			//Secret: secret,
			//SecretBytes:  16,
	
			// side effects:
			FwconsolePath: "fwconsole",
			DoImport:      true,
			DoReload:      true,
			Timeout:       5 * time.Minute, // per fwconsole command
	
			Logger: log.Default(),
		}
	
		/*
		if err := addexts.Generate(ctx, opts); err != nil {
			http.Error(w, "generation/import failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		*/
		list, err := addexts.GenerateRange(ctx, opts)
		if err != nil { log.Fatal(err) }
		
		
		
		
		for _, g := range list {
			new_exten := Extension{
				Number:          g.Extension,      // string
				Displayname:     g.DisplayName,
				Email:           g.Email,
				Secret:          g.Secret,
				Webrtc:          webrtc,
				Mediaencryption: media_encryption,
				Directmedia:     direct_media,
				Public:          public,
				Canblock:        canblock,
				Canunlist:       canunlist,
				Canaddlist:      canaddlist,
				Candial:         candial,
			}
			fmt.Printf("22112: addExtensionPage() new_exten: %+v\n", new_exten)
		
			if err := dbUpsertExtensionAdmin(r.Context(),new_exten); err != nil {
				// error path: err is non-nil, safe to print .Error()
				fmt.Printf("upsert failed for extension %s: %v\n", g.Extension, err)
				// optionally: http.Error(...)
				//continue
			}else{
			// success path: DO NOT touch err here; it’s nil
				fmt.Printf("22122: addExtensionPage() upsert OK for extension %s (display=%q email=%q)\n",
				g.Extension, g.DisplayName, g.Email)
			}
		}
		
		
		amiGetConfigJson(glo_ami)
		
	
		// If these side-effects are desired on edit as well, keep them:
		//amiGetConfigJson(glo_ami)
		reloadXoftSwitch()
		

		


        http.Redirect(w, r, "/extensions", http.StatusSeeOther)
        return
    }

    renderTemplate(w, "addextension", nil)
}

func deleteExtensionPage_dep(w http.ResponseWriter, r *http.Request) {
	if !requireAdminOrLogin(w, r) {
		return
	}

	number := strings.TrimSpace(r.URL.Query().Get("number"))
	if number == "" {
		http.Error(w, "21689: Missing extension number", http.StatusBadRequest)
		return
	}
	// Accept only digits, reasonable length guard
	if !regexp.MustCompile(`^\d{1,8}$`).MatchString(number) {
		http.Error(w, "21690: Invalid extension number format", http.StatusBadRequest)
		return
	}

	logger := log.New(os.Stdout, "[delexts] ", log.LstdFlags)
	opts := delexts.Options{
		PHPPath:        "php",
		ConfPath:       "/etc/freepbx.conf",
		Reload:         false,          // <— avoid double reload; we reload below
		Parallel:       1,
		Logger:         logger,
		PerCallTimeout: 20 * time.Second,
	}

	// Give the whole operation a ceiling (e.g., 30s)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	exts := []string{number}
	results, err := delexts.Delete(ctx, exts, opts)

	var failed bool
	for _, rres := range results {
		if rres.Err != nil {
			logger.Printf("EXT %s FAILED: %v", rres.Ext, rres.Err)
			failed = true
		} else {
			logger.Printf("EXT %s OK: %s", rres.Ext, rres.Output)
		}
	}
	if err != nil {
		logger.Printf("Batch error: %v", err)
		failed = true
	}

	// Your existing hooks
	amiGetConfigJson(glo_ami)
	reloadXoftSwitch()
	/*
	if reloadErr := reloadXoftSwitch(); reloadErr != nil {
		logger.Printf("reloadXoftSwitch error: %v", reloadErr)
		// still redirect; surface error to UI
		if !failed {
			failed = true
		}
	}
	*/

	// Redirect with a message code the list page can show
	if failed {
		http.Redirect(w, r, "/extensions?msg=delete_failed&ext="+url.QueryEscape(number), http.StatusSeeOther)
		return
	}
	
	http.Redirect(w, r, "/extensions?msg=deleted&ext="+url.QueryEscape(number), http.StatusSeeOther)
}

func deleteExtensionPage(w http.ResponseWriter, r *http.Request) {
	if !requireAdminOrLogin(w, r) { return }

	number := strings.TrimSpace(r.URL.Query().Get("number"))
	if number == "" || !regexp.MustCompile(`^\d{1,8}$`).MatchString(number) {
		http.Error(w, "Invalid extension number", http.StatusBadRequest)
		return
	}

	// 1) Write a minimal spinner page and FLUSH so user sees progress immediately.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `<!doctype html>
<html><head><meta charset="utf-8"><title>Deleting %s…</title>
<style>
body{font-family:system-ui,-apple-system,Segoe UI,Roboto,Helvetica,Arial;margin:0;background:#f6f7f9}
.wrap{display:flex;align-items:center;justify-content:center;min-height:60vh}
.card{background:#fff;padding:18px 22px;border-radius:12px;box-shadow:0 6px 24px rgba(0,0,0,.08);display:flex;gap:12px;align-items:center}
.spin{width:22px;height:22px;border:3px solid #ddd;border-top-color:#555;border-radius:50%;animation:spin .8s linear infinite}
@keyframes spin{to{transform:rotate(360deg)}}
</style></head>
<body><div class="wrap"><div class="card">
<div class="spin"></div><div id="msg">Deleting extension %s…</div>
</div></div>`, number, number)

	// Important: flush
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}

	// 2) Do the work (same as before)
	logger := log.New(os.Stdout, "[delexts] ", log.LstdFlags)
	opts := delexts.Options{
		PHPPath:        "php",
		ConfPath:       "/etc/freepbx.conf",
		Reload:         true,
		Parallel:       1,
		Logger:         logger,
		PerCallTimeout: 20 * time.Second,
		FwconsolePath:  "/usr/sbin/fwconsole",
		ReloadTimeout:  60 * time.Second,
	  }
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	results, err := delexts.Delete(ctx, []string{number}, opts)
	var failed bool
	for _, rres := range results {
		if rres.Err != nil {
			logger.Printf("EXT %s FAILED: %v", rres.Ext, rres.Err)
			failed = true
		} else {
			logger.Printf("EXT %s OK: %s", rres.Ext, rres.Output)
		}
	}
	if err != nil { logger.Printf("Batch error: %v", err); failed = true }

	amiGetConfigJson(glo_ami)
	//_ = reloadXoftSwitch() // keep if you need it

	// 3) Tell the already-rendered page to navigate
	dest := "/extensions?msg="
	if failed { dest += "delete_failed&ext=" + url.QueryEscape(number) 
	}else { dest += "deleted&ext=" + url.QueryEscape(number) }

	fmt.Fprintf(w, `<script>setTimeout(function(){ location.replace(%q); }, 300);</script></body></html>`, dest)
	// no redirect header; we let JS move the user after the work completes
}



func userExtensionsPage(w http.ResponseWriter, r *http.Request) {
    if !requireAdminOrLogin(w, r) {
        return
    }

    pageSize := 10
    page := 1
    if p := r.URL.Query().Get("page"); p != "" {
        if pi, err := strconv.Atoi(p); err == nil && pi > 0 {
            page = pi
        }
    }

    search := r.URL.Query().Get("search")
    usernameParam := r.URL.Query().Get("username") // ✅ capture incoming username param
    offset := (page - 1) * pageSize

    var rows *sql.Rows
    var err error

    if search != "" {
        query := fmt.Sprintf(`
            SELECT username, extension, groupid
            FROM userextensions
            WHERE username LIKE ?
            LIMIT %d OFFSET %d`, pageSize, offset)
        rows, err = db.Query(query, "%"+search+"%")
    } else if usernameParam != "" {
        // ✅ filter by exact username if param given
        query := fmt.Sprintf(`
            SELECT username, extension, groupid
            FROM userextensions
            WHERE username = ?
            LIMIT %d OFFSET %d`, pageSize, offset)
        rows, err = db.Query(query, usernameParam)
    } else {
        query := fmt.Sprintf(`
            SELECT username, extension, groupid
            FROM userextensions
            LIMIT %d OFFSET %d`, pageSize, offset)
        rows, err = db.Query(query)
    }

    if err != nil {
        http.Error(w, "DB query error: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var userExts []UserExtension
    for rows.Next() {
        var ue UserExtension
        if err := rows.Scan(&ue.Username, &ue.Extension, &ue.Groupid); err != nil {
            http.Error(w, "Scan error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        userExts = append(userExts, ue)
    }

    data := struct {
        UserExtensions []UserExtension
        Page           int
        PrevPage       int
        NextPage       int
        HasPrev        bool
        HasNext        bool
        Search         string
        UsernameParam  string // ✅ pass through to template
    }{
        UserExtensions: userExts,
        Page:           page,
        PrevPage:       page - 1,
        NextPage:       page + 1,
        HasPrev:        page > 1,
        HasNext:        len(userExts) == pageSize,
        Search:         search,
        UsernameParam:  usernameParam,
    }

    if err := templates["userextensions"].ExecuteTemplate(w, "layout", data); err != nil {
        http.Error(w, "Template render error: "+err.Error(), http.StatusInternalServerError)
    }
}



func usernamestaAPI(w http.ResponseWriter, r *http.Request) {
    if !requireAdminOrLogin(w, r) {
        return
    }

    query := strings.TrimSpace(r.URL.Query().Get("q"))
    if query == "" {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte("[]"))
        return
    }

    rows, err := db.Query(`SELECT username FROM users WHERE username LIKE ? LIMIT 10`, query+"%")
    if err != nil {
        http.Error(w, "DB error: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var usernames []string
    for rows.Next() {
        var u string
        if err := rows.Scan(&u); err == nil {
            usernames = append(usernames, u)
        }
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(usernames)
}
/*
type EditUserExtensionData struct {
    Errors []string
    Form   map[string]string
}
	*/
	func loadAvailableExtensions() ([]string, error) {
		rows, err := db.Query(`SELECT number FROM extensions ORDER BY number`)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
	
		var exts []string
		for rows.Next() {
			var n string
			if err := rows.Scan(&n); err != nil {
				return nil, err
			}
			exts = append(exts, n)
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}
		return exts, nil
	}

	
type EditUserExtensionData struct {
    Errors              []string
    Form                map[string]string
    AvailableExtensions []string
}

func addUserExtensionPage(w http.ResponseWriter, r *http.Request) {
    if !requireAdminOrLogin(w, r) {
        return
    }

    type AddUserExtensionData struct {
        Errors []string
        Form   map[string]string
    }

    if r.Method == http.MethodPost {
        username := strings.TrimSpace(r.FormValue("username"))
        extension := strings.TrimSpace(r.FormValue("extension"))
        groupid := strings.TrimSpace(r.FormValue("groupid"))

        var errs []string
        if username == "" {
            errs = append(errs, "Username is required.")
        }
        if extension == "" {
            errs = append(errs, "Extension is required.")
        }

        // ✅ uniqueness check
        if extension != "" {
            var count int
            err := db.QueryRow(`SELECT COUNT(*) FROM userextensions WHERE extension = ?`, extension).Scan(&count)
            if err != nil {
                http.Error(w, "DB error: "+err.Error(), http.StatusInternalServerError)
                return
            }
            if count > 0 {
                errs = append(errs, "This extension is already assigned to another username.")
            }
        }

        if len(errs) > 0 {
            renderTemplate(w, "adduserextension", AddUserExtensionData{
                Errors: errs,
                Form: map[string]string{
                    "Username":  username,
                    "Extension": extension,
                    "Groupid":   groupid,
                },
            })
            return
        }

        _, err := db.Exec(`INSERT INTO userextensions (username, extension, groupid) VALUES (?,?,?)`,
            username, extension, groupid)
        if err != nil {
            http.Error(w, "Insert error: "+err.Error(), http.StatusInternalServerError)
            return
        }

		/*
		fmt.Println("Add extension to Asterisk if needed!")
		exten := extension
    	secret := ""
    	displayname := ""
		if err := pjsip.FileBasedAddPJSIPExtension(exten, secret, displayname,false); err != nil {
			fmt.Println("❌ Failed to add extension to Asterisk:", err)
			http.Error(w, "ailed to add extension to Asterisk: "+err.Error(), http.StatusInternalServerError)
            return
			
		}
	
		if err := pjsip.ReloadAsterisk(); err != nil {
			fmt.Println("❌ Failed to reload Asterisk:", err)
			http.Error(w, "Failed to reload Asterisk: "+err.Error(), http.StatusInternalServerError)
			return
		}
			*/
		reloadXoftSwitch()
	
		fmt.Println("✅ Extension added and Asterisk reloaded successfully!")
		
		

        http.Redirect(w, r, "/userextensions", http.StatusSeeOther)
        return
    }

    // GET: prefill username if param present
    prefillUsername := r.URL.Query().Get("username")
    renderTemplate(w, "adduserextension", AddUserExtensionData{
        Form: map[string]string{
            "Username": prefillUsername,
        },
    })
}



func editUserExtensionPage(w http.ResponseWriter, r *http.Request) {
    if !requireAdminOrLogin(w, r) {
        return
    }

    // always have the dropdown data available
    exts, err := loadAvailableExtensions()
    if err != nil {
        http.Error(w, "DB error loading extensions: "+err.Error(), http.StatusInternalServerError)
        return
    }

    if r.Method == http.MethodPost {
        origUsername := r.FormValue("original_username")
        origExtension := r.FormValue("original_extension")
        newUsername := strings.TrimSpace(r.FormValue("username"))
        newExtension := strings.TrimSpace(r.FormValue("extension"))
        newGroupID := strings.TrimSpace(r.FormValue("groupid"))

        var errs []string
        if newUsername == "" {
            errs = append(errs, "Username is required.")
        }
        if newExtension == "" {
            errs = append(errs, "Extension is required.")
        }

        // uniqueness check
        var count int
        err := db.QueryRow(
            `SELECT COUNT(*) FROM userextensions WHERE extension = ? AND NOT (username = ? AND extension = ?)`,
            newExtension, origUsername, origExtension,
        ).Scan(&count)
        if err != nil {
            http.Error(w, "DB error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        if count > 0 {
            errs = append(errs, "This extension is already assigned to another username.")
        }

        if len(errs) > 0 {
            renderTemplate(w, "edituserextension", EditUserExtensionData{
                Errors: errs,
                Form: map[string]string{
                    "OriginalUsername":  origUsername,
                    "OriginalExtension": origExtension,
                    "Username":          newUsername,
                    "Extension":         newExtension,
                    "Groupid":           newGroupID,
                },
                AvailableExtensions: exts, // keep dropdown populated on error
            })
            return
        }

        // perform update
        _, err = db.Exec(`
            UPDATE userextensions
            SET username = ?, extension = ?, groupid = ?
            WHERE username = ? AND extension = ?`,
            newUsername, newExtension, newGroupID,
            origUsername, origExtension,
        )
        if err != nil {
            http.Error(w, "Update error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        reloadXoftSwitch()
        http.Redirect(w, r, "/userextensions", http.StatusSeeOther)
        return
    }

    // GET: load the current record
    username := r.URL.Query().Get("username")
    extension := r.URL.Query().Get("extension")
    var ue UserExtension
    err = db.QueryRow(`SELECT username, extension, groupid FROM userextensions WHERE username=? AND extension=?`,
        username, extension).Scan(&ue.Username, &ue.Extension, &ue.Groupid)
    if err != nil {
        http.Error(w, "Record not found", http.StatusNotFound)
        return
    }

    renderTemplate(w, "edituserextension", EditUserExtensionData{
        Form: map[string]string{
            "OriginalUsername":  ue.Username,
            "OriginalExtension": ue.Extension,
            "Username":          ue.Username,
            "Extension":         ue.Extension,
            "Groupid":           ue.Groupid,
        },
        AvailableExtensions: exts, // dropdown options
    })
}


func deleteUserExtensionPage(w http.ResponseWriter, r *http.Request) {
    if !requireAdminOrLogin(w, r) {
        return
    }

    username := r.URL.Query().Get("username")
    extension := r.URL.Query().Get("extension")

    if username == "" || extension == "" {
        http.Error(w, "Missing username or extension", http.StatusBadRequest)
        return
    }

    _, err := db.Exec(`DELETE FROM userextensions WHERE username=? AND extension=?`, username, extension)
    if err != nil {
        http.Error(w, "Delete error: "+err.Error(), http.StatusInternalServerError)
        return
    }
	reloadXoftSwitch()

    http.Redirect(w, r, "/userextensions", http.StatusSeeOther)
}



func addUserPage(w http.ResponseWriter, r *http.Request) {
    if !requireAdminOrLogin(w, r) {
        return
    }

    type AddUserData struct {
		Errors []string
        Message string
        Form    map[string]string
    }

    if r.Method == http.MethodPost {
        // Parse multipart form to allow file upload
        if err := r.ParseMultipartForm(10 << 20); err != nil { // 10 MB
            http.Error(w, "Error parsing form: "+err.Error(), http.StatusBadRequest)
            return
        }

        // Trim inputs
        username := strings.TrimSpace(r.FormValue("username"))
        name := strings.TrimSpace(r.FormValue("name"))
        first := strings.TrimSpace(r.FormValue("firstname"))
        last := strings.TrimSpace(r.FormValue("lastname"))
        email := strings.TrimSpace(r.FormValue("email"))
        phone := strings.TrimSpace(r.FormValue("phonenumber"))
        roleid := strings.TrimSpace(r.FormValue("roleid"))
        statusStr := strings.TrimSpace(r.FormValue("status"))
		homepageurl := strings.TrimSpace(r.FormValue("homepageurl"))

        statusVal := 0
        if statusStr != "" {
            if parsed, convErr := strconv.Atoi(statusStr); convErr == nil {
                statusVal = parsed
            }
        }

        // Validate required fields
        if username == "" || email == "" {
            renderTemplate(w, "adduser", AddUserData{
                Message: "Username and Email are required.",
                Form: map[string]string{
                    "Username":    username,
                    "Name":        name,
                    "Firstname":   first,
                    "Lastname":    last,
                    "Email":       email,
                    "Phonenumber": phone,
                    "Roleid":      roleid,
                    "Status":      statusStr,
					"Homepageurl":      homepageurl,
					
                },
            })
            return
        }

        // Check for existing username
        var exists int
        err := db.QueryRow(`SELECT COUNT(*) FROM users WHERE username = ?`, username).Scan(&exists)
        if err != nil {
            http.Error(w, "DB check error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        if exists > 0 {
            renderTemplate(w, "adduser", AddUserData{
                Message: "Username already exists.",
                Form: map[string]string{
                    "Username":    username,
                    "Name":        name,
                    "Firstname":   first,
                    "Lastname":    last,
                    "Email":       email,
                    "Phonenumber": phone,
                    "Roleid":      roleid,
                    "Status":      statusStr,
					"Homepageurl":      homepageurl,
                },
            })
            return
        }

        // Handle photo upload
        var photoFilename string
        file, header, err := r.FormFile("photo")
        if err == nil && header != nil && header.Filename != "" {
            defer file.Close()

            // Make sure uploads directory exists
            uploadDir := "./static/uploads"
            if mkErr := os.MkdirAll(uploadDir, os.ModePerm); mkErr != nil {
                http.Error(w, "Error creating upload dir: "+mkErr.Error(), http.StatusInternalServerError)
                return
            }

            // Generate unique filename
            ext := filepath.Ext(header.Filename)
            uniqueName := fmt.Sprintf("%s_%d%s", username, time.Now().Unix(), ext)
            photoFilename = uniqueName

            dstPath := filepath.Join(uploadDir, uniqueName)
            out, err := os.Create(dstPath)
            if err != nil {
                http.Error(w, "Error saving file: "+err.Error(), http.StatusInternalServerError)
                return
            }
            defer out.Close()

            if _, err := io.Copy(out, file); err != nil {
                http.Error(w, "Error writing file: "+err.Error(), http.StatusInternalServerError)
                return
            }
        }

        // Insert new user
        _, err = db.Exec(
            `INSERT INTO users (username, name, firstname, lastname, email, phonenumber, roleid, status, photo, homepageurl)
             VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
            username, name, first, last, email, phone, roleid, statusVal, photoFilename,homepageurl,
        )
        if err != nil {
            http.Error(w, "Insert error: "+err.Error(), http.StatusInternalServerError)
            return
        }

        // Redirect to users list
        http.Redirect(w, r, "/users", http.StatusSeeOther)
        return
    }

    // Render page with no data (GET)
    if err := templates["adduser"].ExecuteTemplate(w, "layout", AddUserData{}); err != nil {
        http.Error(w, "Template render error: "+err.Error(), 500)
    }
}

func editUserPage(w http.ResponseWriter, r *http.Request) {
    if !requireAdminOrLogin(w, r) {
        return
    }

    if r.Method == http.MethodPost {
        username := r.FormValue("username")
        name := r.FormValue("name")
        first := r.FormValue("firstname")
        last := r.FormValue("lastname")
        email := r.FormValue("email")
        phone := r.FormValue("phonenumber")
        roleid := r.FormValue("roleid")
        statusStr := r.FormValue("status")
        status, _ := strconv.Atoi(statusStr)
		homepageurl := r.FormValue("homepageurl")

        // Handle file upload
        file, header, err := r.FormFile("photo")
        var photoPath string
        if err == nil && header != nil && header.Filename != "" {
            defer file.Close()

            // Create uploads dir if not exists
            uploadDir := "./static/uploads"
            if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
                http.Error(w, "Error creating upload dir: "+err.Error(), http.StatusInternalServerError)
                return
            }

            // Generate unique filename
            ext := filepath.Ext(header.Filename)
            uniqueName := username + "_" + fmt.Sprint(time.Now().Unix()) + ext
            dstPath := filepath.Join(uploadDir, uniqueName)

            out, err := os.Create(dstPath)
            if err != nil {
                http.Error(w, "Error saving file: "+err.Error(), http.StatusInternalServerError)
                return
            }
            defer out.Close()

            if _, err := io.Copy(out, file); err != nil {
                http.Error(w, "Error writing file: "+err.Error(), http.StatusInternalServerError)
                return
            }

            // ✅ Store only the filename in DB
            photoPath = uniqueName
        }

        // Update user in DB
        if photoPath != "" {
            _, err = db.Exec(`UPDATE users 
                SET name=?, firstname=?, lastname=?, email=?, phonenumber=?, roleid=?, status=?, photo=?, homepageurl=? 
                WHERE username=?`,
                name, first, last, email, phone, roleid, status, photoPath, username,homepageurl)
        } else {
            _, err = db.Exec(`UPDATE users 
                SET name=?, firstname=?, lastname=?, email=?, phonenumber=?, roleid=?, status=? , homepageurl=? 
                WHERE username=?`,
                name, first, last, email, phone, roleid, status, username, homepageurl)
        }

        if err != nil {
            http.Error(w, "Update error: "+err.Error(), http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/users", http.StatusSeeOther)
        return
    }

    // GET: load user
    username := r.URL.Query().Get("username")
    var u User
    err := db.QueryRow(`
        SELECT username, name, firstname, lastname, email, phonenumber, roleid, status, photo, homepageurl 
        FROM users WHERE username=?`, username).
        Scan(&u.Username, &u.Name, &u.Firstname, &u.Lastname, &u.Email, &u.Phonenumber, &u.Roleid, &u.Status, &u.Photo, &u.Homepageurl)

    if err != nil {
        http.Error(w, "User not found: "+err.Error(), 404)
        return
    }

    err = templates["edituser"].ExecuteTemplate(w, "layout", u)
    if err != nil {
        http.Error(w, "Template render error: "+err.Error(), 500)
    }
}

func deleteUserPage(w http.ResponseWriter, r *http.Request) {
    if !requireAdminOrLogin(w, r) { return }

    if r.Method != http.MethodPost {
        http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
        return
    }
    username := r.FormValue("username")
    if username == "" {
        http.Error(w, "Missing username", 400)
        return
    }
    _, err := db.Exec(`DELETE FROM users WHERE username=?`, username)
    if err != nil {
        http.Error(w, "Delete error: "+err.Error(), 500)
        return
    }
    http.Redirect(w, r, "/users", http.StatusSeeOther)
}



func loadConfig() (*Config, error) {
    data, err := os.ReadFile(configPath)
    if err != nil {
        return nil, err
    }
    var c Config
    if err := json.Unmarshal(data, &c); err != nil {
        return nil, err
    }
    return &c, nil
}

/*
func saveConfig(c *Config) error {
    data, err := json.MarshalIndent(c, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(configPath, data, 0644)
}
*/
func saveConfig(c *Config) error {
    // Marshal struct to pretty JSON
    data, err := json.MarshalIndent(c, "", "  ")
    if err != nil {
        return err
    }

    // Ensure config directory exists
    dir := filepath.Dir(configPath)
    if err := os.MkdirAll(dir, 0755); err != nil {
        return err
    }

    // Backup existing config if present
    if _, err := os.Stat(configPath); err == nil {
        backupPath := configPath + ".bak"
        err = os.Rename(configPath, backupPath)
        if err != nil {
            log.Printf("Warning: failed to create backup %s: %v", backupPath, err)
            // Not fatal, continue anyway
        } else {
            log.Printf("Backup created: %s", backupPath)
        }
    }

    // Write to a temp file first for atomic update
    tmpPath := configPath + ".tmp"
    if err := os.WriteFile(tmpPath, data, 0644); err != nil {
        return err
    }

    // Rename temp file to actual config file (atomic on most OS)
    if err := os.Rename(tmpPath, configPath); err != nil {
        return err
    }

    log.Printf("Config saved successfully to %s", configPath)
    return nil
}

func testEmailHandler(w http.ResponseWriter, r *http.Request) {
	// Validate config keys exist and are not empty
	requiredKeys := []string{"SMTPHost", "SMTPPort", "SMTPUsername", "SMTPPassword", "AdminEmail"}
	for _, key := range requiredKeys {
		if config[key] == "" {
			http.Error(w, "Missing email configuration: "+key, http.StatusBadRequest)
			return
		}
	}

	err := sendEmail(config["SenderEmail"], config["AdminEmail"], "Test Email from XoftSwitch", "This is a test email to confirm SMTP settings are correct.")
	//err := emailNotifiyAdmin(config, "Test Email from XoftSwitch", "This is a test email to confirm SMTP settings are correct.")

	if err != nil {
		http.Error(w, "Error sending test email: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Test email sent successfully"))
}

func timeAgo(t time.Time) string {
	duration := time.Since(t)

	switch {
	case duration < time.Minute:
		return "just now"
	case duration < time.Hour:
		return fmt.Sprintf("%d minutes ago", int(duration.Minutes()))
	case duration < 24*time.Hour:
		return fmt.Sprintf("%d hours ago", int(duration.Hours()))
	default:
		return fmt.Sprintf("%d days ago", int(duration.Hours()/24))
	}
}



func humanTime(t time.Time) string {
	duration := time.Since(t)
	switch {
	case duration < time.Minute:
		return "just now"
	case duration < time.Hour:
		return fmt.Sprintf("%d minutes ago", int(duration.Minutes()))
	case duration < 24*time.Hour:
		return fmt.Sprintf("%d hours ago", int(duration.Hours()))
	default:
		return fmt.Sprintf("%d days ago", int(duration.Hours()/24))
	}
}

func settingsPage(w http.ResponseWriter, r *http.Request) {
    if !requireAdminOrLogin(w, r) {
        return
    }

    type ViewData struct {
        *Config
        Message string
        Error   string
    }

    if r.Method == http.MethodPost {
        c := &Config{
			AdminUsername:            r.FormValue("admin_username"),
			AdminPassword:            r.FormValue("admin_password"),
			AdminEmail:               r.FormValue("admin_email"),
			AmiPassword:              r.FormValue("ami_password"),
			AmiUsername:              r.FormValue("ami_username"),
			AsteriskConfPath:         r.FormValue("asterisk_conf_path"),
			AutoJoin:                 r.FormValue("auto_join"),
			CertDir:                  r.FormValue("cert_dir"),
			HomepageURL:              r.FormValue("homepageurl"),
			LogoFilename:             r.FormValue("logo_filename"),
			MaxExtensionCountPerUser: r.FormValue("max_extension_count_per_user"),
			Name:                     r.FormValue("name"),
			PublicHostname:           r.FormValue("public_hostname"),
			PublicIP:                 r.FormValue("public_ip"),
			WifiBSSID:                r.FormValue("wifi_bssid"),
			WifiSSID:                 r.FormValue("wifi_ssid"),
			SMTPHost:                 r.FormValue("smtp_host"),
			SMTPPort:                 r.FormValue("smtp_port"),
			SMTPUsername:             r.FormValue("smtp_username"),
			SMTPPassword:             r.FormValue("smtp_password"),
			
		}
		

        if c.Name == "" {
            c.Name = c.PublicHostname
        }

        // Optional: Validate important fields
        if c.AdminUsername == "" || c.AdminPassword == "" {
            err := templates["settings"].ExecuteTemplate(w, "layout", ViewData{
                Config: c,
                Error:  "Admin username and password are required.",
            })
            if err != nil {
                http.Error(w, "Template render error: "+err.Error(), 500)
            }
            return
        }

        if err := saveConfig(c); err != nil {
            err = templates["settings"].ExecuteTemplate(w, "layout", ViewData{
                Config: c,
                Error:  "Save failed: " + err.Error(),
            })
            if err != nil {
                http.Error(w, "Template render error: "+err.Error(), 500)
            }
            return
        }

        err := templates["settings"].ExecuteTemplate(w, "layout", ViewData{
            Config:  c,
            Message: "Settings saved successfully.",
        })
        if err != nil {
            http.Error(w, "Template render error: "+err.Error(), 500)
        }
        return
    }

    c, err := loadConfig()
    if err != nil {
        http.Error(w, "Load config error: "+err.Error(), 500)
        return
    }

    readFileConfig()
    registerXoServer()

    if len(apikey) == 0 || len(kiv_key) == 0 || len(kiv_iv) == 0 {
        panic(errors.New(`Failed to register server`))
    }

    err = templates["settings"].ExecuteTemplate(w, "layout", ViewData{
        Config: c,
    })
    if err != nil {
        http.Error(w, "Template render error: "+err.Error(), 500)
    }
}


type UpdateJoinRequestPayload struct {
	Action    string `json:"action"`
	Hostname  string `json:"hostname"`
	Username  string `json:"username"`
	RequestID string `json:"request_id"`
	Apikey    string `json:"apikey"`
	Status    string `json:"status"`
}

func callNodeAPI(payload UpdateJoinRequestPayload) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %v", err)
	}

	//resp, err := http.Post("https://api.xoftswitch.com/api", "application/json", bytes.NewBuffer(jsonData))
	resp, err := http.Post(api_xoftswitch, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to call Node API: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Node API returned status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}

func handleApproveJoinRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleApproveJoinRequest: received request")

	if !requireAdminOrLogin(w, r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	id := data.ID
	if id == "" {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	err := approveJoinRequest(id)
	if err != nil {
		http.Error(w, "Failed to approve", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}


func approveJoinRequest(id string) error {
    funcName := "approveJoinRequest"
    fmt.Println(funcName, "START — id =", id)

    // Fetch join request
    fmt.Println(funcName, "calling dbGetJoinRequestById with id:", id)
    jr, err := dbGetJoinRequestById(id)
    if err != nil {
        fmt.Println(funcName, "ERROR — fetching join request:", err)
        return err
    }
    fmt.Println(funcName, "dbGetJoinRequestById returned:", jr)
    if jr == nil {
        fmt.Println(funcName, "no join request found for id:", id)
        return fmt.Errorf("join request not found")
    }

    // Extract fields
    username := jr.Username
    name := jr.Name
    email := jr.Email
    roleid := jr.Roleid
    fmt.Printf("%s — extracted fields: username=%s, name=%s, email=%s, roleid=%v\n", funcName, username, name, email, roleid)

    // Check if user exists
    fmt.Println(funcName, "checking if user exists:", username)
    iuser, user := findUserByUsername(username)
    fmt.Printf("%s — findUserByUsername returned: iuser=%v, user=%+v\n", funcName, iuser, user)

    if user == nil {
        fmt.Println(funcName, "user not found — creating new user:", username)
        iuser, user = createUser("", "", name, username, email, roleid, "", "", "")
        fmt.Printf("%s — createUser returned: iuser=%v, user=%+v\n", funcName, iuser, user)
        if user == nil {
            fmt.Println(funcName, "ERROR — failed to create user")
            return fmt.Errorf("failed to create user")
        }
    }

    // Update local user data
    fmt.Println(funcName, "updating in-memory user username:", username)
    users[iuser].Username = username

    // Update DB username
    fmt.Println(funcName, "updating DB username for email:", email, "to:", username)
    if err := dbUpdateUsername(email, username); err != nil {
        fmt.Println(funcName, "ERROR — failed to update username:", err)
        return err
    }

    // Update userextensions if needed
    fmt.Println(funcName, "checking userextensions for:", username)
    for i, ue := range userextensions {
        fmt.Printf("%s — checking userextensions[%d] = %+v\n", funcName, i, ue)
        if ue.Username == user.Username {
            fmt.Println(funcName, "found matching userextension — updating username to:", username)
            userextensions[i].Username = username
            fmt.Println(funcName, "calling dbUpsertUserExtension for:", userextensions[i])
            dbUpsertUserExtension(userextensions[i])
        }
    }

    // Ensure user has extensions
    fmt.Println(funcName, "fetching final user extensions")
    finalExtensions := getUserExtensions(user)
    fmt.Println(funcName, "finalExtensions count =", len(finalExtensions))
    if len(finalExtensions) == 0 {
        fmt.Println(funcName, "no extensions found — creating user extension")
        createUserExtensionForUser(user)
    }

    // Prepare payload for Node API
    payload := UpdateJoinRequestPayload{
        Action:    "UPDATE_JOINREQUEST_STATUS",
        Hostname:  jr.Hostname,
        Username:  username,
        RequestID: jr.Id,
        Apikey:    config["apikey"],
        Status:    "APPROVED",
    }
    fmt.Println(funcName, "prepared Node API payload:", payload)

    // Call Node API
    fmt.Println(funcName, "calling callNodeAPI")
    if err := callNodeAPI(payload); err != nil {
        fmt.Println(funcName, "ERROR — Node API call failed:", err)
        return err
    }

    // Update join request status in DB
    fmt.Println(funcName, "updating join request status in DB for username:", username)
    if err := dbUpdateJoinRequestStatus(username, 2); err != nil {
        fmt.Println(funcName, "ERROR — failed to update join request status:", err)
        return err
    }

    fmt.Println(funcName, "COMPLETE — join request approved for username:", username)
    return nil
}


func handleDenyJoinRequest(w http.ResponseWriter, r *http.Request) {
	if !requireAdminOrLogin(w, r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	id := data.ID
	if id == "" {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	err := denyJoinRequest(id)
	if err != nil {
		http.Error(w, "Failed to deny", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func denyJoinRequest(id string) error {
	// Example: update the join request's status to -1 (Denied)
	query := "UPDATE join_requests SET status = -1 WHERE id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("denyJoinRequest: %v", err)
	}
	return nil
}


func joinRequestsPage(w http.ResponseWriter, r *http.Request) {
	if !requireAdminOrLogin(w, r) {
		return
	}

	const pageSize = 10
	page := 1
	if p := r.URL.Query().Get("page"); p != "" {
		if pi, err := strconv.Atoi(p); err == nil && pi > 0 {
			page = pi
		}
	}
	offset := (page - 1) * pageSize
	search := strings.TrimSpace(r.URL.Query().Get("search"))

	baseQuery := `
		SELECT id, name, username, email, hostname, status, created
		FROM joinrequests
	`
	orderLimit := fmt.Sprintf("ORDER BY status ASC, created ASC LIMIT %d OFFSET %d", pageSize, offset)

	var (
		rows *sql.Rows
		err  error
	)

	if search != "" {
		query := baseQuery + ` WHERE username LIKE ? ` + orderLimit
		rows, err = db.Query(query, "%"+search+"%")
	} else {
		query := baseQuery + orderLimit
		rows, err = db.Query(query)
	}

	if err != nil {
		http.Error(w, "DB query error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var requests []JoinRequest
	for rows.Next() {
		var jr JoinRequest
		if err := rows.Scan(&jr.Id, &jr.Name, &jr.Username, &jr.Email, &jr.Hostname, &jr.Status, &jr.Created); err != nil {
			http.Error(w, "Scan error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		requests = append(requests, jr)
	}

	data := DashboardData{
		JoinRequests: requests,
		Page:         page,
		PrevPage:     page - 1,
		NextPage:     page + 1,
		HasPrev:      page > 1,
		HasNext:      len(requests) == pageSize,
		Search:       search,
	}

	if err := templates["joinrequests"].ExecuteTemplate(w, "layout", data); err != nil {
		http.Error(w, "Template render error: "+err.Error(), http.StatusInternalServerError)
	}
}

func approveOrDenyJoinRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("approveOrDenyJoinRequestHandler: received request")

	if r.Method != http.MethodPost {
		fmt.Println("approveOrDenyJoinRequestHandler: invalid method", r.Method)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if !requireAdminOrLogin(w, r) {
		fmt.Println("approveOrDenyJoinRequestHandler: unauthorized access attempt")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Println("approveOrDenyJoinRequestHandler: error parsing form:", err)
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	id := r.FormValue("id")
	action := r.FormValue("action") // JOIN_APPROVED or JOIN_DENIED
	fmt.Println("approveOrDenyJoinRequestHandler: id =", id, ", action =", action)

	if id == "" || (action != "JOIN_APPROVED" && action != "JOIN_DENIED") {
		fmt.Println("approveOrDenyJoinRequestHandler: missing or invalid parameters")
		http.Error(w, "Missing or invalid parameters", http.StatusBadRequest)
		return
	}

	var newStatus int
	if action == "JOIN_APPROVED" {
		newStatus = 2
	} else {
		newStatus = -1
	}
	fmt.Println("approveOrDenyJoinRequestHandler: setting newStatus =", newStatus, "for id =", id)

	_, err := db.Exec("UPDATE joinrequests SET status = ? WHERE id = ?", newStatus, id)
	if err != nil {
		fmt.Println("approveOrDenyJoinRequestHandler: database update error:", err)
		http.Error(w, "Database update error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("approveOrDenyJoinRequestHandler: successfully updated join request with id =", id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success"))
}


func joinRequestsAPI(w http.ResponseWriter, r *http.Request) {
	if !requireAdminOrLogin(w, r) {
		return
	}

	pageSize := 10
	page := 1
	if p := r.URL.Query().Get("page"); p != "" {
		if pi, err := strconv.Atoi(p); err == nil && pi > 0 {
			page = pi
		}
	}
	offset := (page - 1) * pageSize
	search := r.URL.Query().Get("search")

	baseQuery := `
		SELECT id, username, email, hostname, status, created
		FROM joinrequests
	`
	orderLimit := fmt.Sprintf("ORDER BY status ASC, created ASC LIMIT %d OFFSET %d", pageSize, offset)

	var (
		rows *sql.Rows
		err  error
	)

	if search != "" {
		query := baseQuery + ` WHERE username LIKE ? ` + orderLimit
		rows, err = db.Query(query, "%"+search+"%")
	} else {
		query := baseQuery + orderLimit
		rows, err = db.Query(query)
	}

	if err != nil {
		http.Error(w, "DB query error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var requests []JoinRequest
	for rows.Next() {
		var jr JoinRequest
		if err := rows.Scan(&jr.Id, &jr.Username, &jr.Email, &jr.Hostname, &jr.Status, &jr.Created); err != nil {
			http.Error(w, "Scan error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		requests = append(requests, jr)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(requests)
}


func approveJoinHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    id := r.FormValue("id")
    if id == "" {
        http.Error(w, "Missing ID", http.StatusBadRequest)
        return
    }

    _, err := db.Exec("UPDATE joinrequests SET status = 2 WHERE id = ?", id)
    if err != nil {
        http.Error(w, "Failed to approve: "+err.Error(), http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/joinrequests", http.StatusSeeOther)
}

func denyJoinHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    id := r.FormValue("id")
    if id == "" {
        http.Error(w, "Missing ID", http.StatusBadRequest)
        return
    }

    _, err := db.Exec("UPDATE joinrequests SET status = -1 WHERE id = ?", id)
    if err != nil {
        http.Error(w, "Failed to deny: "+err.Error(), http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/joinrequests", http.StatusSeeOther)
}
func approveJoin(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		http.Error(w, "Missing ID", 400)
		return
	}

	_, err := db.Exec("UPDATE joinrequests SET status = 1 WHERE id = ?", id)
	if err != nil {
		http.Error(w, "DB error: "+err.Error(), 500)
		return
	}

	// Optional API call
	http.PostForm("http://xoftswitch.com/api", url.Values{
		"action": {"JOIN_APPROVED"},
		"id":     {id},
	})

	setFlash(w, "Request approved.")
	http.Redirect(w, r, "/joinrequests", http.StatusSeeOther)
}

func denyJoin(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		http.Error(w, "Missing ID", 400)
		return
	}

	_, err := db.Exec("UPDATE joinrequests SET status = -1 WHERE id = ?", id)
	if err != nil {
		http.Error(w, "DB error: "+err.Error(), 500)
		return
	}

	setFlash(w, "Request denied.")
	http.Redirect(w, r, "/joinrequests", http.StatusSeeOther)
}

func deleteJoin(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		http.Error(w, "Missing ID", 400)
		return
	}

	_, err := db.Exec("DELETE FROM joinrequests WHERE id = ?", id)
	if err != nil {
		http.Error(w, "DB error: "+err.Error(), 500)
		return
	}

	setFlash(w, "Request deleted.")
	http.Redirect(w, r, "/joinrequests", http.StatusSeeOther)
}

func setFlash(w http.ResponseWriter, msg string) {
	cookie := &http.Cookie{
		Name:  "flash",
		Value: url.QueryEscape(msg),
		Path:  "/",
	}
	http.SetCookie(w, cookie)
}

func getFlash(w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie("flash")
	if err != nil {
		return ""
	}
	msg, _ := url.QueryUnescape(cookie.Value)

	// clear it
	http.SetCookie(w, &http.Cookie{
		Name:   "flash",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	return msg
}


func findJoinRequestById(id string) (*JoinRequest, error) {
	qry := `
		SELECT id, hostname, email, username, name, created, status, auto_join, roleid
		FROM joinrequests
		WHERE id = ?
	`

	row := db.QueryRow(qry, id)

	var jr JoinRequest
	err := row.Scan(&jr.Id, &jr.Hostname, &jr.Email, &jr.Username, &jr.Name, &jr.Created, &jr.Status, &jr.Autojoin, &jr.Roleid)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no result found
		}
		return nil, err
	}

	return &jr, nil
}

func updateJoinRequestStatus(id string, status int) error {
	//db := GetDB() // or your actual DB connection reference
	query := `UPDATE joinrequests SET status = ? WHERE id = ?`
	_, err := db.Exec(query, status, id)
	return err
}



func selected(current, option string) string {
    if current == option {
        return "selected"
    }
    return ""
}


// login page
func renderLogin(w http.ResponseWriter, errorMsg string) {
	tmpl := `
	<!DOCTYPE html>
	<html>
	<head><title>Admin Login</title></head>
	<body>
	<h2>Admin Login</h2>
	{{if .Error}}<p style="color:red;">{{.Error}}</p>{{end}}
	<form method="POST" action="/">
	  <label>Username:</label><input name="username"><br>
	  <label>Password:</label><input name="password" type="password"><br>
	  <button type="submit">Login</button>
	</form>
	</body>
	</html>`

	t := template.Must(template.New("login").Parse(tmpl))
	t.Execute(w, struct{ Error string }{Error: errorMsg})
}


func logoutPage(w http.ResponseWriter, r *http.Request) {
    // clear cookie
    http.SetCookie(w, &http.Cookie{
        Name:   "admin_auth",
        Value:  "",
        Path:   "/",
        MaxAge: -1, // expire immediately
    })
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func joinPageServerDecrypt(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	//fmt.Fprintf(w, "It works!")
	func_name := `joinPage()`
	fmt.Println(func_name, `20039: START`)

	reqBody, _ := ioutil.ReadAll(r.Body)

	/* EXAMPLE

	   var article Article
	   json.Unmarshal(reqBody, &article)
	   // update our global Articles array to include
	   // our new Article
	   Articles = append(Articles, article)

	   json.NewEncoder(w).Encode(article)
	*/
	fmt.Println(func_name, `>>>>>>> reqBody=`, string(reqBody))
	var map_req Map
	json.Unmarshal(reqBody, &map_req)
	fmt.Println(func_name, `+++++++ 19082.1: map_req=`, map_req)
	encrypted_data, ok_encrypted_data := map_req[`data`]
	if !ok_encrypted_data {
		return
	}

	/* TODO: this currently DOES NOT work so we use remote decrypt
	decText, err := Decrypt(encrypted_data, `romthegreat1qwertyasdfghzxcvbn`)
	if err != nil {
	fmt.Println("error decrypting your encrypted text: ", err)
	}
	fmt.Println("=========== decText=",decText)
	*/

	map_data := remoteDecrypt(encrypted_data.(string))
	if map_data != nil {
		////4445: remoteDecrypt() resp.Body= {"type":"join","hostname":"fpbx.sharecle.com","username":"rquidilig2","device_token":"fPSwo29OTVKXB02org0A7a:APA91bGijDKntRYq-yJPsMZN6NT-0eyDi7UbxwTflkbn7Z4jiRZ6eLN1jsIBNhn1R7uUXo6cDEnnutmC2m-Z1HfTg82d7oFXqCP4TCtQv_EgNvHcVxZEgS0Eet4oqdBBDdnwYACSXkOY","sessionid":"80b2a8a8-cb28-4531-9894-54a09b5ae1c8"}

		t, ok_type := map_data[`type`]
		fmt.Println(func_name, `7527: type=`, t, `ok_type=`, ok_type)
		if ok_type && t == `join` {
			str_name := ``
			str_email := ``
			str_photo := ``
			str_homepageurl := ``
			firstname := ``
			lastname := ``
			phonenumber := ``
			//str_roleid := `customer`
			bol_auto_join := false

			username, ok_username := map_data[`username`]
			name, ok_name := map_data[`name`]
			photo, ok_photo := map_data[`photo`]
			homepageurl, ok_homepageurl := map_data[`homepageurl`]

			p_firstname, ok_firstname := map_data[`firstname`]
			if ok_firstname {
				firstname = p_firstname.(string)
			}

			p_lastname, ok_lastname := map_data[`lastname`]
			if ok_lastname {
				lastname = p_lastname.(string)
			}

			p_phonenumber, ok_phonenumber := map_data[`phonenumber`]
			if ok_phonenumber {
				phonenumber = p_phonenumber.(string)
			}

			if ok_name {
				str_name = name.(string)
			}

			email, ok_email := map_data[`email`]
			if ok_email {
				str_email = email.(string)
			}

			if ok_photo {
				str_photo = photo.(string)
			}

			if ok_homepageurl {
				str_homepageurl = homepageurl.(string)
			}

			auto_join, ok_auto_join := map_data[`auto_join`]
			if ok_auto_join {
				bol_auto_join = auto_join.(bool)
			} else {
				cfg_auto_join, ok_cfg_auto_join := config[`auto_join`]
				if ok_cfg_auto_join {
					if cfg_auto_join == `yes` {
						bol_auto_join = true
					}
				}

			}
			//sessionid,ok_sessionid := map_data[`sessionid`]
			//if ok_username && ok_sessionid{
			if ok_username {
				//fmt.Println(`username=`,username, `ok_username=`,ok_username,`sessionid=`,sessionid,`ok_sessionid=`,ok_sessionid )
				fmt.Println(func_name, `20211: username=`, username, `ok_username=`, ok_username)
				//if isAuthenticated(username,sessionid){
				ret := httpJoinResponse(username.(string), str_name, str_email, str_photo, str_homepageurl, firstname, lastname, phonenumber, bol_auto_join)
				fmt.Println(func_name, `6968: map_response=`, ret)
				//fmt.Println(`4493: map_response=`,data)
				//ret := Map{`statusCode`:200,`data`:data}
				json.NewEncoder(w).Encode(ret)
				//}
			}

		}
	}
}


func joinPage(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "It works!")
	enableCors(&w)
	func_name := `joinPage()`
	fmt.Println(func_name, `20158: START`)

	enableCors(&w)

	reqBody, _ := ioutil.ReadAll(r.Body)

	
	fmt.Println(func_name, `>>>>>>> reqBody=`, string(reqBody))
	var map_req Map
	json.Unmarshal(reqBody, &map_req)
	fmt.Println(func_name, `+++++++ 19203.1: map_req=`, map_req)
	encrypted_data, ok_encrypted_data := map_req[`data`]
	fmt.Println(func_name, `+++++++ 19203.1: ok_encrypted_data=`, ok_encrypted_data)
	if !ok_encrypted_data {
		
		return
	}

	//map_data := remoteDecrypt(encrypted_data.(string))
	encrypted := encrypted_data.(string)
	fmt.Println(func_name, `+++++++ 19203.2: encrypted=`, encrypted, `kiv_key=`, kiv_key, `kiv_iv=`, kiv_iv)

	decText, err := AesDecrypt(encrypted, kiv_key, kiv_iv)
	if err != nil {
		fmt.Println("error decrypting your encrypted text: ", err)
		return
	}

	fmt.Println("=========== decText=", decText)

	var map_data Map
	json.Unmarshal([]byte(decText), &map_data)

	/*
	final toJsonResult = {
        //'data':{
        'type': 'join',
        'hostname': hostname,
        'username': Globals.session.user.username,
        'email': Globals.session.user.profile.email.address,

        'device_token': Globals.session.device.getDeviceToken(),
        'sessionid': Globals.session.sessionId

        //}
      };
	*/
	if map_data != nil {
		////4445: remoteDecrypt() resp.Body= {"type":"join","hostname":"fpbx.sharecle.com","username":"rquidilig2","device_token":"fPSwo29OTVKXB02org0A7a:APA91bGijDKntRYq-yJPsMZN6NT-0eyDi7UbxwTflkbn7Z4jiRZ6eLN1jsIBNhn1R7uUXo6cDEnnutmC2m-Z1HfTg82d7oFXqCP4TCtQv_EgNvHcVxZEgS0Eet4oqdBBDdnwYACSXkOY","sessionid":"80b2a8a8-cb28-4531-9894-54a09b5ae1c8"}

		t, ok_type := map_data[`type`]
		fmt.Println(func_name, `7527: type=`, t, `ok_type=`, ok_type)
		if ok_type && t == `join` {
			str_name := ``
			str_email := ``
			firstname := ``
			lastname := ``
			phonenumber := ``

			//str_roleid := `customer`
			bol_auto_join := false

			username, ok_username := map_data[`username`]

			name, ok_name := map_data[`name`]
			if ok_name {
				str_name = name.(string)
			}

			p_firstname, ok_firstname := map_data[`firstname`]
			if ok_firstname {
				firstname = p_firstname.(string)
			}

			p_lastname, ok_lastname := map_data[`lastname`]
			if ok_lastname {
				lastname = p_lastname.(string)
			}

			p_phonenumber, ok_phonenumber := map_data[`phonenumber`]
			if ok_phonenumber {
				phonenumber = p_phonenumber.(string)
			}

			email, ok_email := map_data[`email`]
			if ok_email {
				str_email = email.(string)
			}

			str_photo := ``
			photo, ok_photo := map_data[`photo`]
			if ok_photo {
				str_photo = photo.(string)
			}
			str_homepageurl := ``
			homepageurl, ok_homepageurl := map_data[`homepageurl`]
			if ok_homepageurl {
				str_homepageurl = homepageurl.(string)
			}
			/*
				roleid,ok_roleid := map_data[`roleid`]
				if ok_roleid{
					str_roleid = roleid.(string)
				}
			*/

			//auto_join := false

			auto_join, ok_auto_join := map_data[`auto_join`]
			if ok_auto_join {
				bol_auto_join = auto_join.(bool)
			} else {
				cfg_auto_join, ok_cfg_auto_join := config[`auto_join`]
				if ok_cfg_auto_join {
					if cfg_auto_join == `yes` {
						bol_auto_join = true
					}
				}

			}
			//sessionid,ok_sessionid := map_data[`sessionid`]
			//if ok_username && ok_sessionid{
			if ok_username {
				//fmt.Println(`username=`,username, `ok_username=`,ok_username,`sessionid=`,sessionid,`ok_sessionid=`,ok_sessionid )
				fmt.Println(func_name, `username=`, username, `ok_username=`, ok_username)
				//if isAuthenticated(username,sessionid){
				ret := httpJoinResponse(username.(string), str_name, str_email, str_photo, str_homepageurl, firstname, lastname, phonenumber, bol_auto_join)
				fmt.Println(func_name, `6968: map_response=`, ret)
				//fmt.Println(`4493: map_response=`,data)
				//ret := Map{`statusCode`:200,`data`:data}
				json.NewEncoder(w).Encode(ret)
				//}
			}

		}
	}
}


func getStringFromMap(m map[string]interface{}, key string) string {
	if val, ok := m[key]; ok {
		if strVal, ok := val.(string); ok {
			return strVal
		}
	}
	return ""
}

func humanizeTime(t time.Time) string {
	duration := time.Since(t)

	switch {
	case duration < time.Minute:
		return "just now"
	case duration < time.Hour:
		return fmt.Sprintf("%d minutes ago", int(duration.Minutes()))
	case duration < 24*time.Hour:
		return fmt.Sprintf("%d hours ago", int(duration.Hours()))
	case duration < 48*time.Hour:
		return "yesterday"
	default:
		return fmt.Sprintf("%d days ago", int(duration.Hours()/24))
	}
}

// Define a simple response struct
type JoinResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
}

type JoinRequest struct {
    Id           string `json:"id"`
    Hostname     string `json:"hostname"`
    Email        string `json:"email"`
    Username     string `json:"username"`
    Name         string `json:"name"`
    Created      string `json:"created"`
    CreatedHuman string `json:"created_human"`
    Status       int    `json:"status"`
    Autojoin     int    `json:"autojoin"`
    Roleid       string `json:"roleid"`
    Phonenumber  string `json:"phonenumber"` 
}

type RemoveJoinRequest struct {
    Id           string `json:"id"`
    Hostname     string `json:"hostname"`
    
	Username     string `json:"username"`
    
}




func newJoinRequestDb(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	func_name := `requesToJoinPage()`
	fmt.Println(func_name, `20158: START`)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(JoinResponse{Status: -1, Msg: "Failed to read request body"})
		return
	}

	var map_data Map
	if err := json.Unmarshal(reqBody, &map_data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(JoinResponse{Status: -1, Msg: "Invalid JSON format"})
		return
	}

	fmt.Println(`22517: map_data= `, map_data)

	str_id := getStringFromMap(map_data, "id")
	if str_id == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(JoinResponse{Status: -1, Msg: "Missing 'id'"})
		return
	}

	str_username := getStringFromMap(map_data, "username")
	if str_username == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(JoinResponse{Status: -1, Msg: "Missing 'username'"})
		return
	}

	str_hostname := getStringFromMap(map_data, "hostname")
	if str_hostname == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(JoinResponse{Status: -1, Msg: "Missing 'str_hostname'"})
		return
	}

	// Extract other optional fields
	str_email := getStringFromMap(map_data, "email")
	str_name := getStringFromMap(map_data, "name")
	firstname := getStringFromMap(map_data, "firstname")
	lastname := getStringFromMap(map_data, "lastname")
	phonenumber := getStringFromMap(map_data, "phonenumber")
	str_photo := getStringFromMap(map_data, "photo")
	str_homepageurl := getStringFromMap(map_data, "homepageurl")

	// Call internal join logic
	result := httpNewJoinRequestDbResponse(str_id,str_hostname,str_username, str_name, str_email, str_photo, str_homepageurl, firstname, lastname, phonenumber, false)

	// Assume status is an integer in result["status"]
	statusInt := 1
	if val, ok := result["status"].(int); ok {
		statusInt = val
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(JoinResponse{
		Status: statusInt,
		Msg:    fmt.Sprintf("User %s processed", str_username),
		Data:   result,
	})
}


func removeJoinRequestDbResponse(id string) Map {
	func_name := `removeJoinRequestDbResponse()`
	fmt.Println(func_name, `START`, `id=`, id)

	if err := dbRemoveJoinRequestById(id); err != nil {
		fmt.Println(func_name, "DB delete error:", err)
		return Map{
			"statusCode": 500,
			"error":      "Failed to remove join request",
		}
	}

	return Map{
		"statusCode": 200,
	}
}


func handelerRemoveJoinRequest(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

/*
curl -X POST https://localhost:8180/remove-joinrequest \
-H "Content-Type: application/json" \
-d '{"id": "abc123"}'

*/

	func_name := `handelerRemoveJoinRequest()`
	fmt.Println(func_name, `23173: START`)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(JoinResponse{Status: -1, Msg: "Failed to read request body"})
		return
	}

	var map_data Map
	if err := json.Unmarshal(reqBody, &map_data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(JoinResponse{Status: -1, Msg: "Invalid JSON format"})
		return
	}

	fmt.Println("22517: map_data =", map_data)

	/*
	str_id := getStringFromMap(map_data, "id")
	
	if str_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(JoinResponse{Status: -1, Msg: "Missing 'id'"})
		return
	}
	*/

	val, ok := map_data["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(JoinResponse{Status: -1, Msg: "Missing 'id' key"})
		return
	}

	str_id, _ := val.(string) // will be "" if empty, and that's allowed
	
	
	// Call internal logic
	result := removeJoinRequestDbResponse(str_id)

	// Extract status code
	statusInt := 500
	if val, ok := result["statusCode"].(int); ok {
		statusInt = val
	}

	// Get message from result["error"], fallback to success msg
	msg := fmt.Sprintf("ID %s processed", str_id)
	if errMsg, ok := result["error"].(string); ok {
		msg = errMsg
	}

	// Use actual statusInt as HTTP response code (optional)
	w.WriteHeader(statusInt)
	json.NewEncoder(w).Encode(JoinResponse{
		Status: statusInt,
		Msg:    msg,
		Data:   result,
	})
}




// func activateUserExtensionPage(w http.ResponseWriter, r *http.Request) {
func checkinPage(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "It works!")
	ip := getIPAddress(r)
	//fmt.Fprintf(w, "Client IP: %s\n", ip)
	//if ip != "67.227.0.62" {
	if ip != api_server_ip {
		
		fmt.Println("checkinPage: IP ", ip, " is NOT ", api_server_ip, " access denied.")
		return
	}
	 
	fmt.Println("20148: checkinPage:  IP ", ip, " is VALID ")
		
	
	enableCors(&w)
	func_name := `joinPage()`
	fmt.Println(func_name, `20320: START`)

	enableCors(&w)

	reqBody, _ := ioutil.ReadAll(r.Body)

	/* EXAMPLE

	   var article Article
	   json.Unmarshal(reqBody, &article)
	   // update our global Articles array to include
	   // our new Article
	   Articles = append(Articles, article)

	   json.NewEncoder(w).Encode(article)
	*/
	fmt.Println(func_name, `>>>>>>> reqBody=`, string(reqBody))
	
	

	var map_data Map
	json.Unmarshal(reqBody, &map_data)

	if map_data != nil {
		
		
		username := ``
		exten_num := ``
		groupid := ``

		p_username, ok_username := map_data[`username`]

		p_exten_num, ok_exten_num := map_data[`exten_num`]

		p_groupid, ok_groupid := map_data[`groupid`]

		if ok_username && ok_exten_num && ok_groupid {
			username = p_username.(string)

			exten_num = p_exten_num.(string)

			groupid = p_groupid.(string)

			name := ``
			p_name, ok_name := map_data[`name`]
			if ok_name {
				name = p_name.(string)
			}

			firstname := ``
			p_firstname, ok_firstname := map_data[`firstname`]
			if ok_firstname {
				firstname = p_firstname.(string)
			}

			lastname := ``
			p_lastname, ok_lastname := map_data[`lastname`]
			if ok_lastname {
				lastname = p_lastname.(string)
			}

			phonenumber := ``
			p_phonenumber, ok_phonenumber := map_data[`phonenumber`]
			if ok_phonenumber {
				phonenumber = p_phonenumber.(string)
			}

			email := ``
			p_email, ok_email := map_data[`email`]
			if ok_email {
				email = p_email.(string)
			}

			photo := ``
			p_photo, ok_photo := map_data[`photo`]
			if ok_photo {
				photo = p_photo.(string)
			}

			homepageurl := ``
			p_homepageurl, ok_homepageurl := map_data[`homepageurl`]
			if ok_homepageurl {
				homepageurl = p_homepageurl.(string)
			}

			fmt.Println(func_name, `username=`, username, `ok_username=`, ok_username)

			ret := httpCheckInResponse(username, name, email, photo, homepageurl, firstname, lastname, phonenumber, exten_num, groupid)
			fmt.Println(func_name, `6968: map_response=`, ret)

			json.NewEncoder(w).Encode(ret)

		}

		//}
	}
}

func checkoutPage(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "It works!")
	ip := getIPAddress(r)
	//fmt.Fprintf(w, "Client IP: %s\n", ip)
	if ip != api_server_ip {

		fmt.Println("checkoutPage: IP ", ip, " is NOT ", api_server_ip, " access denied.")
		return
	}

	fmt.Println("20148: checkoutPage:  IP ", ip, " is VALID ")
	enableCors(&w)
	func_name := `joinPage()`
	fmt.Println(func_name, `20496: START`)

	enableCors(&w)

	reqBody, _ := ioutil.ReadAll(r.Body)

	
	
	fmt.Println(func_name, `>>>>>>> reqBody=`, string(reqBody))
	
	
	var map_data Map
	json.Unmarshal(reqBody, &map_data)

	if map_data != nil {
		
		
		username := ``
		exten_num := ``
		groupid := ``

		p_username, ok_username := map_data[`username`]
		if ok_username {
			username = p_username.(string)
		}
		p_exten_num, ok_exten_num := map_data[`exten_num`]
		if ok_exten_num {
			exten_num = p_exten_num.(string)
		}

		p_groupid, ok_groupid := map_data[`groupid`]
		if ok_groupid {
			groupid = p_groupid.(string)
		}

		if len(groupid) > 0 && len(username) > 0 && len(exten_num) > 0 {

			fmt.Println(func_name, `groupid=`, groupid, `ok_groupid=`, ok_groupid)

			ret := httpCheckOutResponse(groupid)
			fmt.Println(func_name, `6968: map_response=`, ret)

			json.NewEncoder(w).Encode(ret)

		}

		//}
	}
}


func apiPage(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	ip, ip_err := getIP(r)
	if ip_err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//fmt.Fprintf(w, "It works!")
	func_name := `apiPage`
	statusCode := 500
	error_message := ``

	enableCors(&w)
	reqBody, _ := ioutil.ReadAll(r.Body)

	
	fmt.Println(`>>>>>>> reqBody=`, string(reqBody))
	var map_req Map
	json.Unmarshal(reqBody, &map_req)
	fmt.Println(`+++++++ 13044: map_req=`, map_req)
	encrypted_data, ok_encrypted_data := map_req[`data`]
	if !ok_encrypted_data {
		fmt.Println(`+++++++ 13046: map_req is invalid!`)

		json.NewEncoder(w).Encode(Map{`statusCode`: 501})

		return
	}

	
	encrypted := encrypted_data.(string)
	//fmt.Println(func_name,`+++++++ 19371: encrypted=`,encrypted, `kiv_key=`,kiv_key,`kiv_iv=`,kiv_iv)
	fmt.Println(string(colorYellow), func_name, `+++++++ 19372.1: encrypted=`, encrypted, `kiv_key=`, kiv_key, `kiv_iv=`, kiv_iv)

	decText, err := AesDecrypt(encrypted, kiv_key, kiv_iv)
	if err != nil {
		//fmt.Println("error decrypting your encrypted text: ", err)
		fmt.Println(string(colorRed), func_name, "19372.2: decText=", decText)
		fmt.Println(string(colorReset))
		return
	}
	//fmt.Println("=========== decText=",decText)
	fmt.Println(string(colorYellow), func_name, "19372.3: decText=", decText)

	fmt.Println(string(colorReset))

	var map_data Map
	json.Unmarshal([]byte(decText), &map_data)

	if map_data != nil {
		////4445: remoteDecrypt() resp.Body= {"type":"join","hostname":"fpbx.sharecle.com","username":"rquidilig2","device_token":"fPSwo29OTVKXB02org0A7a:APA91bGijDKntRYq-yJPsMZN6NT-0eyDi7UbxwTflkbn7Z4jiRZ6eLN1jsIBNhn1R7uUXo6cDEnnutmC2m-Z1HfTg82d7oFXqCP4TCtQv_EgNvHcVxZEgS0Eet4oqdBBDdnwYACSXkOY","sessionid":"80b2a8a8-cb28-4531-9894-54a09b5ae1c8"}

		t, ok_type := map_data[`type`]
		str_type := t.(string)

		fmt.Println(func_name, `9797: str_type=`, str_type)
		if ok_type {
			switch str_type {
			case `assign_or_add_extension`:

				//if ok_type && t == `assign_or_add_extenstion`  {
				username, ok_username := map_data[`username`]
				exten, ok_exten := map_data[`exten`]
				exten_to_username, ok_exten_to_username := map_data[`exten_to_username`]

				//sessionid,ok_sessionid := map_data[`sessionid`]
				//if ok_username && ok_sessionid{
				if ok_username && ok_exten && ok_exten_to_username {
					//fmt.Println(`username=`,username, `ok_username=`,ok_username,`sessionid=`,sessionid,`ok_sessionid=`,ok_sessionid )
					fmt.Println(`username=`, username, `ok_username=`, ok_username)
					//if isAuthenticated(username,sessionid){
					is_ok := httpAddAssignExtenResponse(username.(string), exten.(string), exten_to_username.(string))

					//fmt.Println(`4493: map_response=`,data)
					fmt.Println(`5241: httpAddAssignExtenResponse returned %v`, is_ok)
					if is_ok {

						//json.NewEncoder(w).Encode(Map{`statusCode`:200})
						statusCode = 200

					}

				} else {
					fmt.Println(`5241: assign_or_add_extenstion. Invalid or missing parameters!`)

				}

			case `admin_add_user`:
				fmt.Println(func_name, `7654: admin_add_user mapdata=`, map_data)
				u, err := httpAdminAddUser(map_data)

				//fmt.Println(`4493: map_response=`,data)
				//fmt.Println(`5241: httpAddAssignExtenResponse returned %v`,is_ok)
				if err == nil {

					//json.NewEncoder(w).Encode(Map{`statusCode`:200})
					fmt.Println(`Admin add user`, u.Username, `OK.`)
					statusCode = 200

				} else {
					//statusCode = 500
					error_message = err.Error()
					//fmt.Println(`Admin add user`, username, `FAILED.`,error_message)
					fmt.Println(`Admin add user FAILED.`, error_message)

				}
				
				
				json.NewEncoder(w).Encode(fmt.Sprintf(`{"statusCode":%v,"error":%s}`, statusCode, error_message))
				return

			case `admin_edit_user`:
				fmt.Println(func_name, `7654: admin_edit_user mapdata=`, map_data)
				err := httpAdminEditUser(map_data)

				//fmt.Println(`4493: map_response=`,data)
				//fmt.Println(`5241: httpAddAssignExtenResponse returned %v`,is_ok)
				if err == nil {

					//json.NewEncoder(w).Encode(Map{`statusCode`:200})
					//fmt.Println(`Admin edit user`, u.Username, `OK.`)
					fmt.Println(`Admin edit user OK.`)
					statusCode = 200

				} else {
					//statusCode = 500
					error_message = err.Error()
					//fmt.Println(`Admin add user`, username, `FAILED.`,error_message)
					fmt.Println(`Admin add user FAILED.`, error_message)

				}
				/*
					} else{
						fmt.Println(func_name,`7672: params INVALID OR MISSING`)
						//fmt.Println(`5241: assign_or_add_extenstion. Invalid or missing parameters!`)
						error_message = `One or more input parameters is invalid.`
					}
				*/
				json.NewEncoder(w).Encode(fmt.Sprintf(`{"statusCode":%v,"error":%s}`, statusCode, error_message))
				return

			case `CDR_MORE`:

				username := ""
				q := ""
				g := ""
				//skip := 0
				pagenum := 0
				//limit := 100

				m_username, ok_username := map_data[`username`]
				if ok_username {
					username = m_username.(string)
					m_q, ok_q := map_data["q"]

					if ok_q {
						q = m_q.(string)
					}

					m_g, ok_g := map_data["g"]

					if ok_g {
						g = m_g.(string)
					}

					/*
						m_skip,ok_skip:= map_data["skip"]

						if ok_skip{ skip = int(m_skip.(float64))}
					*/
					m_pagenum, ok_pagenum := map_data["pagenum"]

					if ok_pagenum {
						pagenum = int(m_pagenum.(float64))
					}

					//ret := httpCDRResponse(username,q,g,skip)
					ret := httpCDRResponse(username, q, g, pagenum)
					json.NewEncoder(w).Encode(ret)
					return
				}

			case `GET_EXTENSION`:
				exten, ok_exten := map_data[`exten`]
				if ok_exten {
					ret := httpGetExtensionResponse(exten.(string))
					json.NewEncoder(w).Encode(ret)
					return
				}

			case `GET_MY_EXTENSIONS`:
				exten, ok_exten := map_data[`exten`]
				if ok_exten {
					//ret := httpGetExtensionResponse(exten.(string))
					ret := httpGetMyExtensionsResponse(exten.(string))
					json.NewEncoder(w).Encode(ret)
					return
				}

			case `GET_MYUSER`:
				//username,ok_username := map_data[`username`]
				//if ok_username{
				//ret := httpGetMyUserResponse(username.(string))
				ret := httpGetMyUserResponse(map_data)
				json.NewEncoder(w).Encode(ret)
				return
				//}

			case `GET_MY_API_TOKEN`:
				//username,ok_username := map_data[`username`]
				//if ok_username{
				//ret := httpGetMyUserResponse(username.(string))
				ret := httpGetMyAPITokenResponse(map_data)
				json.NewEncoder(w).Encode(ret)
				return

			case `GET_MYHOMEPAGEURL`:
				username, ok_username := map_data[`username`]
				if ok_username {
					ret := httpGetMyHomePageUrlResponse(username.(string), ip)
					json.NewEncoder(w).Encode(ret)
					return
				}

			case `AUTHENTICATE_MYAIATOKEN`:
				username, ok_username := map_data[`username`]
				if ok_username {
					ret := httpAuthenticateAIATokenResponse(username.(string), ip)
					json.NewEncoder(w).Encode(ret)
					return
				}

			/*
			case `GET_MY_USEREXTENSIONS`:
				q, ok_q := map_data[`q`]
				if ok_q {
					ret := httpGetMyExtensionsResponse(q.(string))
					json.NewEncoder(w).Encode(ret)
					return
				}
				*/

			case `MORE_EXTENSIONS`:

				ret := httpMoreExtensionsResponse(map_data)
				json.NewEncoder(w).Encode(ret)
				return
			case `MORE_MYCONTACT_EXTENSIONS`:

				ret := httpMoreMyContactExtensionsResponse(map_data)
				json.NewEncoder(w).Encode(ret)
				return
			case `MORE_USERS`:

				ret := httpMoreUsersResponse(map_data)
				json.NewEncoder(w).Encode(ret)
				return

			case `BLOCK_USER`:

				ret := httpBlockUserResponse(map_data)

				json.NewEncoder(w).Encode(ret)
				return

			case `ADMIN_DELETE_USER`:

				ret := httpDeleteUserResponse(map_data)

				json.NewEncoder(w).Encode(ret)
				return

			case `ADMIN_DELETE_EXTENSION`:

				ret := httpAdminDeleteExtenResponse(map_data)

				json.NewEncoder(w).Encode(ret)
				return

			case `UNLIST_EXTENSION`:

				ret := httpUnlistExtensionResponse(map_data)

				json.NewEncoder(w).Encode(ret)
				return

			case `ADMIN_EDIT_EXTENSION`:

				ret := httpAdminEditExtenResponse(map_data) //httpAdminSetPublicExtenResponse(map_data)

				json.NewEncoder(w).Encode(ret)
				return

			case `INBOX_MORE`:
				username, ok_username := map_data[`username`]
				if ok_username {
					ret := httpInboxResponse(username.(string))
					json.NewEncoder(w).Encode(ret)
					return
				}

			case `MESSAGES_MORE`:
				username, ok_username := map_data[`username`]
				inboxid, ok_inboxid := map_data[`inboxid`]

				if ok_username && ok_inboxid {
					ret := httpInboxMessagesResponse(username.(string), inboxid.(string), map_data)
					//ret := httpInboxMessagesResponse(username.(string), map_data)
					json.NewEncoder(w).Encode(ret)
					return
				}

			case `NEW_MESSAGE`:

				username, ok_username := map_data[`username`]
				from, ok_from := map_data[`from`]
				to, ok_to := map_data[`to`]
				roomid, ok_roomid := map_data[`roomid`]
				message, ok_message := map_data[`message`]
				attachment, ok_attachment := map_data[`attachment`]

				if ok_username && ok_from && ok_to && ok_roomid && ok_message && ok_attachment {
					if len(message.(string)) > 0 && len(from.(string)) > 0 && len(to.(string)) > 0 {
						ret := httpAddNewMessageResponse(username.(string), from.(string), to.(string), roomid.(string), message.(string), attachment.(string))
						json.NewEncoder(w).Encode(ret)
						return
					}

				} else {
					fmt.Println(`12924: NEW_MESSAGE did not passed validation!`)
				}

			case `TYPEAHEAD_CONTACTS`: //For messagging
				username, ok_username := map_data[`username`]
				q, ok_q := map_data[`q`]

				if ok_username && ok_q {
					ret := httpTypeAheadContactsResponse(q.(string), username.(string))
					json.NewEncoder(w).Encode(ret)
					return
				}

			case `TYPEAHEAD_EXTENSIONS`: //For calls
				username, ok_username := map_data[`username`]
				q, ok_q := map_data[`q`]

				if ok_username && ok_q {
					ret := httpTypeAheadExtensionsResponse(q.(string), username.(string))
					json.NewEncoder(w).Encode(ret)
					return
				}

			case `DISJOIN`:
				username, ok_username := map_data[`username`]
				//inboxid,ok_inboxid := map_data[`inboxid`]

				if ok_username {
					ret := httpDisJoinResponse(username.(string))
					json.NewEncoder(w).Encode(ret)
					return
				}

			case `PING`:
				//_,ok_username := map_data[`username`]

				//if ok_username{

				ret := Map{`statusCode`: 200}
				json.NewEncoder(w).Encode(ret)
				return
				//}

			default:
				fmt.Println(func_name, `5247: Invalid or missing type! the type is %s`, str_type)

			}
		}

	}

	json.NewEncoder(w).Encode(Map{`statusCode`: statusCode})
}


func logoPage(w http.ResponseWriter, r *http.Request) {
	//reqBody, _ := ioutil.ReadAll(r.Body)
	enableCors(&w)
	//buf, err := ioutil.ReadFile("logo.png")
	buf, err := ioutil.ReadFile(config[`logo_filename`])

	if err != nil {

		//log.Fatal(err)
		fmt.Println(`imagesPage() err=`, err.Error())
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(buf)
}
func uploadFile(w http.ResponseWriter, r *http.Request) {
	//https://tutorialedge.net/golang/go-file-upload-tutorial/

	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	/*
		File Upload Endpoint Hit
		Uploaded File: rquidilig.PNG
		File Size: 6215822
		MIME Header: map[Content-Disposition:[form-data; name="myFile"; filename="rquidilig.PNG"] Content-Type:[image/png]]
	*/
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)

	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	fmt.Printf("Content-Type: %+v\n", handler.Header[`Content-Type`])

	contenttypes := handler.Header[`Content-Type`]

	ok_contenttype := false

	contenttype := ``
	for _, x := range contenttypes {
		if x == `image/png` || x == `image/jpg` || x == `image/jpeg` {
			ok_contenttype = true
			contenttype = x
			//is_save_file = true
			break

		}
	}

	if ok_contenttype {

		// Create a temporary file within our temp-images directory that follows
		// a particular naming pattern

		temppattern := "upload-*.png"

		if contenttype == `image/jpg` {
			temppattern = "upload-*.jpg"
		} else if contenttype == `image/jpeg` {
			temppattern = "upload-*.jpeg"
		}

		tempFile, err := ioutil.TempFile("temp-images", temppattern)
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		// write this byte array to our temporary file
		tempFile.Write(fileBytes)
		fmt.Println("Successfully Uploaded File ", tempFile.Name())
		// return that we have successfully uploaded our file!
		fmt.Fprintf(w, "Successfully Uploaded File\n")
		// Rename and Remove a file
		// Using Rename() function

		/*Original_Path := "GeeksforGeeks.txt"
		New_Path := "gfg.txt"
		e := os.Rename(Original_Path, New_Path)
		if e != nil {
			log.Fatal(e)
		} */
	} else {
		fmt.Fprintf(w, "Upload failed. Invalid.\n")
	}

}

func uploadFileExtensionAvatar(w http.ResponseWriter, r *http.Request) {
	//https://tutorialedge.net/golang/go-file-upload-tutorial/

	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()

	var exten_num = r.FormValue("exten_num")
	fmt.Println("exten_num=", exten_num)
	if len(exten_num) == 0 {
		return
	}

	/*
		File Upload Endpoint Hit
		Uploaded File: rquidilig.PNG
		File Size: 6215822
		MIME Header: map[Content-Disposition:[form-data; name="myFile"; filename="rquidilig.PNG"] Content-Type:[image/png]]
	*/
	fmt.Println(string(colorYellow))
	fmt.Println("Uploaded File: %+v", handler.Filename)
	file_name := handler.Filename

	fmt.Println("File Size: %+v", handler.Size)
	fmt.Println("MIME Header: %+v", handler.Header)
	fmt.Println("Content-Type: %+v", handler.Header[`Content-Type`])

	contenttypes := handler.Header[`Content-Type`]

	ok_contenttype := false

	contenttype := ``
	for _, x := range contenttypes {
		if x == `image/png` || x == `image/jpg` || x == `image/jpeg` {
			ok_contenttype = true
			contenttype = x
			//is_save_file = true
			break

		}
	}

	isvalidfile := false

	fexten := path.Ext(file_name)

	if ok_contenttype {
		isvalidfile = true
	} else {
		switch strings.ToLower(fexten) {
		case ".png":
			fallthrough
		case ".jpg":
			fallthrough
		case ".jpeg":
			isvalidfile = true

		}
	}

	if isvalidfile {
		fmt.Println("isvalidfile = true")

		// Create a temporary file within our temp-images directory that follows
		// a particular naming pattern

		temppattern := "upload-*.png"

		if contenttype == `image/jpg` {
			temppattern = "upload-*.jpg"
		} else if contenttype == `image/jpeg` {
			temppattern = "upload-*.jpeg"
		}

		tempFile, err := ioutil.TempFile("temp-images", temppattern)
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		// write this byte array to our temporary file
		tempFile.Write(fileBytes)
		fmt.Println("Successfully Uploaded File ", tempFile.Name())
		// return that we have successfully uploaded our file!
		fmt.Fprintf(w, "Successfully Uploaded File\n")
		// Rename and Remove a file
		// Using Rename() function

		/*Original_Path := "GeeksforGeeks.txt"
		New_Path := "gfg.txt"
		e := os.Rename(Original_Path, New_Path)
		if e != nil {
			log.Fatal(e)
		} */

		Original_Path := tempFile.Name()

		fmt.Println(`19812.1: Original_Path=`, Original_Path)
		newpath := filepath.Join(".", "static/images")

		err1 := os.MkdirAll(newpath, os.ModePerm)

		if err1 != nil {
			return
		}

		New_Path := fmt.Sprintf(`static/images/exten-%s%s`, exten_num, fexten)
		fmt.Println(`19812.2: New_Path=`, New_Path)

		e := os.Rename(Original_Path, New_Path)
		if e != nil {
			log.Fatal(e)
		}

		fmt.Println(string(colorGreen), "19812.3: Upload success!", New_Path)

	} else {
		fmt.Println(string(colorRed), "19812.3: Upload failed. Invalid.")
		//fmt.Fprintf(w, "19812.3: Upload failed. Invalid.\n")
	}

	fmt.Println(string(colorReset))

}

func serveFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path
	if p == "./" {
		p = "./static/index.html"
	}
	http.ServeFile(w, r, p)
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("79: %v", err)

	} else {
		register(ws)
	}
	/*
		clients = append(clients, *ws)

		fmt.Println("Client Connected")
		err = ws.WriteMessage(1, []byte("Hi Client!"))
		if err != nil {
			fmt.Println("87: %v",err)
		}
		// listen indefinitely for new messages coming
		// through on our WebSocket connection
		reader(ws)
	*/
}

// getIP returns the ip address from the http request
func getIP(r *http.Request) (string, error) {
	ips := r.Header.Get("X-Forwarded-For")
	splitIps := strings.Split(ips, ",")

	if len(splitIps) > 0 {
		// get last IP in list since ELB prepends other user defined IPs, meaning the last one is the actual client IP.
		netIP := net.ParseIP(splitIps[len(splitIps)-1])
		if netIP != nil {
			return netIP.String(), nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	netIP := net.ParseIP(ip)
	if netIP != nil {
		ip := netIP.String()
		if ip == "::1" {
			return "127.0.0.1", nil
		}
		return ip, nil
	}

	return "", errors.New("IP not found")
}
func getIPAddress(r *http.Request) string {
	// Check X-Forwarded-For header (comma-separated list)
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		ips := strings.Split(xff, ",")
		return strings.TrimSpace(ips[0]) // First IP is the client's IP
	}

	// Check X-Real-IP header
	xRealIP := r.Header.Get("X-Real-IP")
	if xRealIP != "" {
		return xRealIP
	}

	// Fallback to RemoteAddr
	ip := r.RemoteAddr
	if colon := strings.LastIndex(ip, ":"); colon != -1 {
		return ip[:colon] // Remove port if present
	}

	return ip
}

func logRequests(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("📥 %s %s\n", r.Method, r.URL.Path)
        next(w, r)
    }
}

func setupRoutes() {
	fmt.Println("21530: START setupRoutes() ")

	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/logout", logoutPage)
	//http.HandleFunc("/contactus", contactUsPage)
	http.HandleFunc("/about", aboutPage)
	

	http.HandleFunc("/users", usersPage)
	http.HandleFunc("/adduser", addUserPage)
	http.HandleFunc("/edituser", editUserPage)
	http.HandleFunc("/deleteuser", deleteUserPage)

	http.HandleFunc("/extensions", extensionsPageHandler)
	http.HandleFunc("/addextension", addExtensionPage)
	//http.HandleFunc("/addextension", addExtensionHandler)
	http.HandleFunc("/reload", reloadHandler)
	
	http.HandleFunc("/editextension", editExtensionPage)
	http.HandleFunc("/deleteextension", deleteExtensionPage)

	http.HandleFunc("/userextensions", userExtensionsPage)
	http.HandleFunc("/adduserextension", addUserExtensionPage)
	http.HandleFunc("/edituserextension", editUserExtensionPage)
	http.HandleFunc("/deleteuserextension", deleteUserExtensionPage)

	http.HandleFunc("/settings", settingsPage)

	http.HandleFunc("/join", joinPage)
	http.HandleFunc("/join-approve", logRequests(handleApproveJoinRequest))
	http.HandleFunc("/join-deny", handleDenyJoinRequest)
	

	http.HandleFunc("/new-joinrequest", newJoinRequestDb)
	http.HandleFunc("/remove-joinrequest", handelerRemoveJoinRequest)
	
	http.HandleFunc("/get-joinrequest", getRequestToJoinDb)
	http.HandleFunc("/joinrequests", joinRequestsPage)
	http.HandleFunc("/deletejoin", deleteJoin)
	http.HandleFunc("/join/action", approveOrDenyJoinRequestHandler)

	http.HandleFunc("/api", apiPage)
	http.HandleFunc("/logo", logoPage)
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/uploadextensionavatar", uploadFileExtensionAvatar)

	http.HandleFunc("/ws", wsEndpoint)

	http.HandleFunc("/checkin", checkinPage)
	http.HandleFunc("/checkout", checkoutPage)
	http.HandleFunc("/dashboardstatus", statusAPI)
	http.HandleFunc("/usernamesta", usernamestaAPI)
	http.HandleFunc("/settings/test-email", testEmailHandler)

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 👇 Place this LAST so it doesn't override specific routes
	http.HandleFunc("/", dashboardPage)

	fmt.Println("21978: END setupRoutes() ")
}
