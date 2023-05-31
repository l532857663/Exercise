package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
)

func main() {
	// 定义原始字符串
	str := "Hello, world! This is a test string."

	// 使用gzip对字符串进行压缩
	var compressed bytes.Buffer
	gzipWriter := gzip.NewWriter(&compressed)
	_, err := gzipWriter.Write([]byte(str))
	if err != nil {
		panic(err)
	}
	gzipWriter.Close()

	compressedBytes := compressed.Bytes()
	fmt.Printf("Compressed size: %d %x\n", len(compressedBytes), compressedBytes)

	// 解压缩字符串
	gzipReader, err := gzip.NewReader(bytes.NewReader(compressedBytes))
	if err != nil {
		panic(err)
	}
	uncompressedBytes := new(bytes.Buffer)
	_, err = uncompressedBytes.ReadFrom(gzipReader)
	if err != nil {
		panic(err)
	}
	str2 := uncompressedBytes.String()

	fmt.Println("Original string: ", str)
	fmt.Println("Uncompressed string: ", str2)
}
