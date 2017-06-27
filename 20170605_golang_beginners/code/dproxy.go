// START OMIT
		json.Unmarshal(f.([]byte), &conf)
		monitorName, err := dproxy.New(conf).M("name").String()
// END OMIT
