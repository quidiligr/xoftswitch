package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/googolgl/gami"
	"github.com/googolgl/gami/event"
	"regexp"
	"io/ioutil"
)

type Map map[string]interface{}

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

	Displayname string `json:"displayname"`
	Number string `json:"number"`
	Status string `json:"status"`
	Contacts string `json:"contacts"`
	Source string `json:"source"`
	//Index	int `json:"index"`
	Secret      string `json:"secret"`
    Devicetype  string `json:"devicetype"`
    Public      string `json:"public"`
    Canblock    string `json:"canblock"`
    Canunlist   string `json:"canunlist"`
    Canaddlist  string `json:"canaddList"`
    Candial     string `json:"candial"`
    Ver         string `json:"ver"`
}

var endpoint = Map {}
var extensions = []Extension{}

func main() {
	done := make(chan bool)
	ami, err := gami.Dial("127.0.0.1:5038")
	if err != nil {
		log.Fatal(err)
	}
	defer ami.Close()

	ami.Run()
	
	//install manager
	go func() {
		for {
			select {
			//handle network errors
			case err := <-ami.NetError:
				log.Println("24: Network Error:", err)
				//try new connection every second
				<-time.After(time.Second)
				if err := ami.Reconnect(); err == nil {
					//call start actions
					ami.Action(gami.Params{"Action":"Events","EventMask":"on"})
				}
				
			case err := <-ami.Error:
				log.Println("33: error:", err)
			//wait events and process
			case ev := <-ami.Events:
				log.Println("36: Event Detect:", *ev)
				//if want type of events
				//log.Println("38: EventType:", event.New(ev))
				ev_type := event.New(ev)
				//log.Println("38: EventType:", event.New(ev))
				log.Println("455: EventType:", ev_type)
				log.Println("456: ID=",ev.ID,"Privilege=",ev.Privilege,"Params=", ev.Params)
				
			}
		}
	}()
	
	if err := ami.Login("xoftphone", "625f975db835748deda7d646803684d5"); err != nil {
		log.Fatal("44:",err)
	}
	
	// Action method
	rsPing, rsActioID, rsErr := ami.Action(gami.Params{"Action":"Ping"})
	if rsErr != nil {
		
		log.Fatal("51: FATAL rsErr:",rsErr)
	}
	log.Println("53: rsActioID",rsActioID)

	// Synchronous response
	log.Println(<-rsPing)

	// Asynchronous response
	/*
	go func() {
		log.Println("60: rsActioID",<-rsPing)
	}()
	*/
						
	if _, _, err := ami.Action(gami.Params{"Action":"Events","EventMask":"on"}); err != nil {
		log.Fatal("53: FATAL  err",err)
	}

	amiGetConfigJson(ami)

	rsEndpoint, rsEndpointActionID, rsEndpointErr := ami.Action(gami.Params{"Action":"GetConfig","Filename":"pjsip.endpoint.conf"})
	if rsEndpointErr != nil{
		log.Println("76: rsEndpoint=",rsEndpointErr)
		
	}else{
		
		log.Println("rsEndpointActionID=",rsEndpointActionID)
		//log.Println(<-rsGetConfig)
		rs := <- rsEndpoint
		log.Println("93: ++++++++++ rsEndpoint ID",rs.ID, "Status=",rs.Status,"Params=",rs.Params)
	}
	
	
/*
	rsListCategories, rsListCategoriesID, rsListCategoriesErr := ami.Action(gami.Params{"Action":"ListCategories","Filename":"extensions.conf"})
	if rsListCategoriesErr != nil{
		log.Println("76: rsGetConfigErr=",rsListCategoriesErr)
		
	}else{
		
		log.Println("rsListCategoriesID=",rsListCategoriesID)
		log.Println(<-rsListCategories)
	}
	*/
	
//	rsGetConfig, rsGetConfigActionID, rsGetConfigErr := ami.Action(gami.Params{"Action":"GetConfigJSON","Filename":"extensions.conf","Category":"from-internal"})
	//rsGetConfig, rsGetConfigActionID, rsGetConfigErr := ami.Action(gami.Params{"Action":"GetConfigJSON","Filename":"extensions.conf"})
	//rsGetConfig, rsGetConfigActionID, rsGetConfigErr := ami.Action(gami.Params{"Action":"GetConfigJSON","Filename":"pjsip.aor.conf"})
	

	//rsCommand, rsCommandActionID, rsCommandErr := ami.Action(gami.Params{"Action":"Command","Command":"pjsip show aors"})
	//rsCommand, _, _ := ami.Action(gami.Params{"Action":"Command","Command":"pjsip show aors"})
	//rsCommand, _, _ := ami.Action(gami.Params{"Action":"CoreShowChannels"})
	//rsCommand, _, _ := ami.Action(gami.Params{"Action":"ListCommands"})
	//rsCommand, _, _ := ami.Action(gami.Params{"Action":"Pjsipshowendpoints"})
	//rsCommand, _, _ := ami.Action(gami.Params{"Action":"Pjsipshowendpoints"})
	rsCommand, _, _ := ami.Action(gami.Params{"Action":"Pjsipshowaors"})
	//
	
	
	//go func() {
		rs := <-rsCommand
		log.Println("123: rsCommand", <-rsCommand)

		//log.Println("ID=",rsCommand.ID,"Status=",rsCommand.Status,"Params=",rsCommand.Params)
		log.Println("127: ++++++++++ ID",rs.ID, "Status=",rs.Status,"Params=",rs.Params)
	//}()
		
/*
	if rsCommandErr != nil{
		log.Println("76: rsCommandErr=",rsCommandErr)
		
	}else{
		
		log.Println("rsCommandActionID=",rsCommandActionID)
		
		body := fmt.Sprintf("%v",<-rsCommand)
		
		
		log.Println("body=",body)
	}
	*/
	
	log.Println("72.1: ping:", <-rsPing)
	<-done
}

func writeFileAOR(data Map){
	
	
	content, err := json.Marshal(data)
	if err != nil {
		log.Println("ERROR writeFileAOR() data=",data)
		fmt.Println(err)
	}
	//log.Println("writeFileRegistered() content=",content)
	err = ioutil.WriteFile("aor.json", content, 0644)
	if err != nil {
		log.Println(err)
	}	
}

func amiGetConfigJson(ami *gami.AMIClient){
	rsGetConfig, rsGetConfigActionID, rsGetConfigErr := ami.Action(gami.Params{"Action":"GetConfigJSON","Filename":"pjsip.endpoint.conf"})
	
	if rsGetConfigErr != nil{
		log.Println("76: rsGetConfigErr=",rsGetConfigErr)
		
	}else{
		
		log.Println("rsGetConfigActionID=",rsGetConfigActionID)
		//log.Println(<-rsGetConfig)
		body := fmt.Sprintf("%v",<-rsGetConfig)
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
		if len(matches) == 2{
			//log.Println("matches=",matches[1])
			var prettyJSON bytes.Buffer
			error := json.Indent(&prettyJSON, []byte(matches[1]), "", "\t")
			if error != nil {
				log.Println("JSON parse error: ", error)
			
			
			} else{
				log.Println("prettyJSON:", string(prettyJSON.String()))
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
					log.Println("json_content=",json_content)
				}
				*/
				//var aor Map
				
				err := json.Unmarshal([]byte(matches[1]), &endpoint)
				if err != nil {
					log.Println(err)
				}else{
					log.Println("endpoint=",endpoint)
					//extensions = Map{}
					i := 0
					for k, v := range endpoint { 
						
						log.Println("k=",k,"v=", v, "i=",i)
						
						
							
							switch v.(type) {
							case map[string]interface{}:
								vmap,ok_vmap := v.(map[string]interface{})
								log.Println("ok_vmap",ok_vmap, "vmap=",vmap)
								if ok_vmap {
									callerid, ok_callerid := vmap["callerid"].(string)
									log.Println("ok_callerid",ok_callerid, "callerid=",callerid)
									
									if ok_callerid{
										log.Println("callerid=",callerid)
										//extensions[k] = Extension{Number: k,Displayname: callerid,Status: ``, Index:i}
										extensions = append(extensions, Extension{Number: k,Displayname: callerid,Status: ``, Index:i})
									}else{
										//extensions[k] = Extension{Number: k,Displayname: k,Status: ``, Index:i}
										extensions = append(extensions, Extension{Number: k,Displayname: k,Status: ``, Index:i})
									}
									i++
								}
								
								
							default:
								log.Println("Unexpected!!!")
								
							}
							
							
						
						
						
						
					}
					log.Println("extensions=",extensions)
					
				}
				
			}
		}
		/*
		body := fmt.Sprintf("%v",<-rsGetConfig)
		log.Println("body=",body)
		var prettyJSON bytes.Buffer
		error := json.Indent(&prettyJSON, []byte(body), "", "\t")
		if error != nil {
			log.Println("JSON parse error: ", error)
			
			return
		}

		//log.Println("prettyJSON:", string(prettyJSON.Bytes()))
		log.Println("prettyJSON:", string(prettyJSON.String()))
		*/
	}
	
}
