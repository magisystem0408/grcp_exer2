package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"
	"grcp_exer2/pb"
	"io/ioutil"
	"log"
)

func main() {
	employee := &pb.Employee{
		Id:    1,
		Name:  "mamushi",
		Email: "test@example.com",
		//pbと打つと、候補が出てくる
		Oppuration:  pb.Occpation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "080-1234-2114"},
		Project:     map[string]*pb.Company_Project{"project": &pb.Company_Project{}},
		Profile: &pb.Employee_Text{
			Text: "My name is suzuki",
		},
		Birthday: &pb.Date{
			Year:  2000,
			Month: 20,
			Day:   1,
		},
	}

	//これをバイナリデータとして書き出し
	// シリアライズ化
	binData, err := proto.Marshal(employee)
	if err != nil {
		log.Printf("Can't serialize", err)

	}
	if err := ioutil.WriteFile("test.bin", binData, 0666); err != nil {
		log.Println("Can't write", err)
	}

	in, err := ioutil.ReadFile("test.bin")
	if err != nil {
		log.Fatalln("Can't read", err)
	}

	//デシリアラウズされたものを空の構造体に入る
	readEmployee := &pb.Employee{}
	err = proto.Unmarshal(in, readEmployee)
	if err != nil {
		log.Fatalln("Can't unmarshal'", err)
	}

	fmt.Println(readEmployee)

	//jsonへの変換
	m := jsonpb.Marshaler{}
	out, err := m.MarshalToString(employee)
	if err != nil {
		log.Fatalln("Can't json", err)
	}
	fmt.Println()
	fmt.Println(out)

	//jsonから構造体への変換
	readJsonEmployee := &pb.Employee{}
	if err := jsonpb.UnmarshalString(out, readJsonEmployee); err != nil {
		log.Fatalln("Can't unmarshal from JSON", err)
	}

	fmt.Println()
	fmt.Println(readJsonEmployee)
}
