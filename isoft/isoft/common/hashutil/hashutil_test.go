package hashutil

import (
	"fmt"
	"testing"
)


func Test_CalculateHashWithString(t *testing.T)  {
	fmt.Println(CalculateHashWithString("admin"))
}

func Test_CalculateHashWithFile(t *testing.T)  {
	hash1, _ := CalculateHashWithFile("D:/Elasticsearch.docx")
	fmt.Println(hash1)

	hash2, _ := CalculateHashWithBigFile("D:/Elasticsearch.docx")
	fmt.Println(hash2)

	hash3, _ := CalculateHashWithFileS("D:/Elasticsearch.docx")
	fmt.Println(hash3)
}