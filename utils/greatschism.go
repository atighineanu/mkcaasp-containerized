package utils

const RegCode = "INTERNAL-USE-ONLY-2afd635738257c4a"

/*func Hashinator(pass string, key string, homedir string, caasporsespdir string) {
	v := OSAPI{}
	var tempkey string
	//----------if there isn't key.json we create one basing on *hash string...
	out, err := exec.Command("ls", "-alh", homedir).CombinedOutput()
	if strings.Contains(fmt.Sprintf("%s", string(out)), "key.json") && !strings.Contains(key, "default") {
		f, err := os.Create(homedir + "/key.json")
		if err != nil {
			log.Fatalf("func Hashinator: couldn't create the file...%s", err)
		}
		f.Write([]byte("\"" + key + "\""))
		f.Close()
	}

	file, err := os.Open(homedir + "/key.json")
	if err != nil {
		log.Fatalf("couldn't open the file \"key.json\": %s", err)
		//	return err
	}

	decoder := json.NewDecoder(file)
	defer file.Close()
	err = decoder.Decode(&tempkey)
	if err != nil {
		fmt.Printf("This is bad! .json decoding didn't work at opening %s key.json:  %s\n", homedir, err)
	}

	file, _ = os.Open(homedir + "/" + caasporsespdir + "/openstack.json")
	decoder = json.NewDecoder(file)
	defer file.Close()
	err = decoder.Decode(&v)
	if err != nil {
		fmt.Printf("This is bad! .json opening at %s openstack.json didn't work:  %s\n", homedir+"/"+caasporsespdir, err)
	}
	block, err := aes.NewCipher([]byte(Hasher(tempkey)))
	if err != nil {
		log.Fatalf("cipher block wasn't created properly: %s\n", err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("new GCM block wasn't created properly: %s\n", err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("", err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, []byte(pass), nil)
	f, err := ioutil.ReadFile(homedir + "/" + caasporsespdir + "/openstack.json")
	if err != nil {
		fmt.Printf("Error at opening json file!...%s\n", err)
	}
	v.Password = ciphertext[:]
	f, _ = json.MarshalIndent(v, "", " ")
	err = ioutil.WriteFile(homedir+"/"+caasporsespdir+"/openstack.json", f, 0644)
	if err != nil {
		fmt.Printf("Error at writing to json file!...%s\n", err)
	}
}

func Dehashinator(homedir string, caasporsespdir string) string {
	v := OSAPI{}
	file, _ := os.Open(caasporsespdir + "/openstack.json")
	decoder := json.NewDecoder(file)
	defer file.Close()
	err := decoder.Decode(&v)
	if err != nil {
		fmt.Printf("This is bad! .json decoding didn't work @opening openstack.json: %s", err)
	}

	var tempkey string
	file, _ = os.Open(homedir + "/key.json")
	decoder = json.NewDecoder(file)
	defer file.Close()
	err = decoder.Decode(&tempkey)
	if err != nil {
		fmt.Printf("This is bad! .json decoding didn't work @opening key.json  %s:", err)
	}

	block, err := aes.NewCipher([]byte(Hasher(tempkey)))
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := v.Password[:nonceSize], v.Password[nonceSize:]
	decoded, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return string(decoded)
}

func Hasher(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
*/
