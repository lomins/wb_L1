package main

import "fmt"

//------------------------

type AnaliticalDataService interface {
	SendXmlData()
}

type XmlDocument struct {
}

func (doc XmlDocument) SendXmlData() {
	fmt.Println("Отправка xml документа...")
}

//------------------------

type JsonDocument struct {
}

func (doc JsonDocument) ConvertToXml() string {
	return "<xml></xml>"
}

type JsonDocumentaAdapter struct {
	jsonDocument *JsonDocument
}

func (adapter JsonDocumentaAdapter) SendXmlData() {
	data := adapter.jsonDocument.ConvertToXml()
	fmt.Println("Отправка Xml данных: ", data)
}

//------------------------

func main() {
	jsonDoc := &JsonDocument{}

	jsonDocAdapter := &JsonDocumentaAdapter{
		jsonDocument: jsonDoc,
	}

	jsonDocAdapter.SendXmlData()
}
