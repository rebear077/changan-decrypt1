package server

// type Server struct {
// 	encrypte     *encrypter.Encrypter
// 	sql          *sql.SqlCtr
// 	contractAddr string
// 	symKey       []byte
// 	pubKey       []byte
// 	priKey       []byte
// }

// func NewServer() *Server {
// 	en := encrypter.NewEncrypter()
// 	addr, err := getContractAddr("configs/contractAddress.txt")
// 	if err != nil {
// 		logrus.Fatalln(err)
// 	}
// 	symkey, err := getSymKey("configs/symPri.txt")
// 	if err != nil {
// 		logrus.Fatalln(err)
// 	}
// 	pubkey, err := getRSAPublicKey("configs/public.pem")
// 	if err != nil {
// 		logrus.Fatalln(err)
// 	}
// 	prikey, err := getRSAPrivateKey("configs/private.pem")
// 	if err != nil {
// 		logrus.Fatalln(err)
// 	}

// 	return &Server{
// 		encrypte:     en,
// 		sql:          sql.NewSqlCtr(),
// 		contractAddr: addr,
// 		symKey:       symkey,
// 		pubKey:       pubkey,
// 		priKey:       prikey,
// 	}
// }

// func (s *Server) ValidateHash(hash []byte, plain []byte) bool {
// 	resHash := s.encrypte.Signature(plain)
// 	if string(resHash) == string(hash) {
// 		return true
// 	} else {
// 		return false
// 	}
// }
// func getContractAddr(path string) (string, error) {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return "", err
// 	}
// 	stat, _ := file.Stat()
// 	addr := make([]byte, stat.Size())
// 	_, err = file.Read(addr)
// 	if err != nil {
// 		return "", err
// 	}
// 	err = file.Close()
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(addr), nil
// }
// func getSymKey(path string) ([]byte, error) {
// 	filesymPrivate, err := os.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	stat, err := filesymPrivate.Stat()
// 	if err != nil {
// 		return nil, err
// 	}
// 	symkey := make([]byte, stat.Size())
// 	filesymPrivate.Read(symkey)
// 	filesymPrivate.Close()
// 	return symkey, nil
// }
// func getRSAPublicKey(path string) ([]byte, error) {
// 	pubKey, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return pubKey, nil
// }
// func getRSAPrivateKey(path string) ([]byte, error) {
// 	privateKey, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return privateKey, err
// }

// func (s *Server) InsertLog(level string, info string) {
// 	time := time.Now().String()[0:19]
// 	err := s.sql.InsertLogs(time, level, info)
// 	if err != nil {
// 		fmt.Println("....", err)
// 	}
// }

// func (s *Server) DecryptSymkey(ensymkey []byte) ([]byte, error) {
// 	symkey, err := s.encrypte.AsymDecrypt(ensymkey, s.priKey)
// 	// fmt.Println("私钥： ", string(s.priKey))
// 	// fmt.Println(err)
// 	return symkey, err
// }

// func (s *Server) DecryptData(endata string, symkey []byte) ([]byte, error) {
// 	data, err := s.encrypte.SymDecrypt([]byte(endata), symkey)
// 	return data, err
// }
